package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"
)

type mockPrototype struct {
	ID             string                  `json:"id"`
	Role           string                  `json:"role"`
	Delegate       prototypeUserModelJSON  `json:"delegate"`
	ActivatingUser *prototypeUserModelJSON `json:"activating_user,omitempty"`
}

type mockQuay struct {
	server       *httptest.Server
	mu           sync.Mutex
	protoSeq     int
	prototypes   map[string]map[string]mockPrototype   // org -> id -> prototype
	teams        map[string]map[string]map[string]bool // org -> team -> member -> true
	deletedTeams map[string]map[string]bool            // org -> team -> true
}

func newMockQuay(t *testing.T) *mockQuay {
	t.Helper()
	m := &mockQuay{
		prototypes:   make(map[string]map[string]mockPrototype),
		teams:        make(map[string]map[string]map[string]bool),
		deletedTeams: make(map[string]map[string]bool),
	}
	server := httptest.NewServer(m.handler())
	m.server = server
	t.Cleanup(server.Close)
	t.Setenv("QUAY_URL", server.URL)
	t.Setenv("QUAY_TOKEN", "mock-token")
	return m
}

func (m *mockQuay) writeNotFound(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status": 404,
		"type":   "https://quay.io/api/v1/error/not_found",
		"title":  "not_found",
		"detail": detail,
	})
}

func (m *mockQuay) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.Trim(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		// Base check: /api/v1/organization/{org}/...
		if len(parts) < 4 || parts[0] != "api" || parts[1] != "v1" || parts[2] != "organization" {
			m.writeNotFound(w, "Not found")
			return
		}

		org := parts[3]

		// /api/v1/organization/{org}/prototypes[/{id}]
		if parts[4] == "prototypes" {
			m.handlePrototypes(w, r, org, parts[5:])
			return
		}

		// /api/v1/organization/{org}/team/{team}/members[/{member}]
		if parts[4] == "team" && len(parts) >= 7 && parts[6] == "members" {
			team := parts[5]
			m.handleTeamMembers(w, r, org, team, parts[7:])
			return
		}

		m.writeNotFound(w, "Not found")
	})
}

func (m *mockQuay) handlePrototypes(w http.ResponseWriter, r *http.Request, org string, subParts []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(subParts) == 0 {
		switch r.Method {
		case http.MethodGet:
			var list []mockPrototype
			if orgProtos, ok := m.prototypes[org]; ok {
				for _, p := range orgProtos {
					list = append(list, p)
				}
			}
			sort.Slice(list, func(i, j int) bool {
				return list[i].ID < list[j].ID
			})
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"prototypes": list,
			})
			return

		case http.MethodPost:
			var req struct {
				Role           string                  `json:"role"`
				Delegate       prototypeUserModelJSON  `json:"delegate"`
				ActivatingUser *prototypeUserModelJSON `json:"activating_user"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			m.protoSeq++
			protoID := fmt.Sprintf("proto-%d", m.protoSeq)
			proto := mockPrototype{
				ID:             protoID,
				Role:           req.Role,
				Delegate:       req.Delegate,
				ActivatingUser: req.ActivatingUser,
			}
			if m.prototypes[org] == nil {
				m.prototypes[org] = make(map[string]mockPrototype)
			}
			m.prototypes[org][protoID] = proto
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(proto)
			return

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	if len(subParts) == 1 {
		protoID := subParts[0]
		orgProtos, ok := m.prototypes[org]
		if !ok {
			m.writeNotFound(w, "Prototype not found")
			return
		}
		proto, ok := orgProtos[protoID]
		if !ok {
			m.writeNotFound(w, "Prototype not found")
			return
		}

		switch r.Method {
		case http.MethodPut:
			var req struct {
				Role string `json:"role"`
			}
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			proto.Role = req.Role
			m.prototypes[org][protoID] = proto
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(proto)
			return

		case http.MethodDelete:
			delete(m.prototypes[org], protoID)
			w.WriteHeader(http.StatusNoContent)
			return

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	m.writeNotFound(w, "Not found")
}

func (m *mockQuay) handleTeamMembers(w http.ResponseWriter, r *http.Request, org, team string, subParts []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(subParts) == 0 {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if m.deletedTeams[org] != nil && m.deletedTeams[org][team] {
			m.writeNotFound(w, "Team not found")
			return
		}

		membersMap, exists := m.teams[org][team]
		if !exists {
			m.writeNotFound(w, "Team not found")
			return
		}

		members := make([]teamMemberModelJSON, 0, len(membersMap))
		for name := range membersMap {
			members = append(members, teamMemberModelJSON{
				Name:    name,
				Kind:    "user",
				IsRobot: strings.Contains(name, "+"),
				Invited: false,
			})
		}
		sort.Slice(members, func(i, j int) bool {
			return members[i].Name < members[j].Name
		})

		resp := teamModelJSON{
			Name:    team,
			Members: members,
			CanEdit: true,
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	if len(subParts) == 1 {
		member := subParts[0]

		switch r.Method {
		case http.MethodPut:
			if m.teams[org] == nil {
				m.teams[org] = make(map[string]map[string]bool)
			}
			if m.teams[org][team] == nil {
				m.teams[org][team] = make(map[string]bool)
			}
			m.teams[org][team][member] = true
			if m.deletedTeams[org] != nil {
				delete(m.deletedTeams[org], team)
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{"name": member})
			return

		case http.MethodDelete:
			if m.deletedTeams[org] != nil && m.deletedTeams[org][team] {
				m.writeNotFound(w, "Team not found")
				return
			}
			if m.teams[org] != nil && m.teams[org][team] != nil {
				if _, ok := m.teams[org][team][member]; !ok {
					m.writeNotFound(w, "Member not found")
					return
				}
				delete(m.teams[org][team], member)
			} else {
				m.writeNotFound(w, "Team not found")
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}

	m.writeNotFound(w, "Not found")
}

// Test hook methods

func (m *mockQuay) AddTeamMember(org, team, member string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.teams[org] == nil {
		m.teams[org] = make(map[string]map[string]bool)
	}
	if m.teams[org][team] == nil {
		m.teams[org][team] = make(map[string]bool)
	}
	m.teams[org][team][member] = true
	if m.deletedTeams[org] != nil {
		delete(m.deletedTeams[org], team)
	}
}

func (m *mockQuay) RemoveTeamMember(org, team, member string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.teams[org] != nil && m.teams[org][team] != nil {
		delete(m.teams[org][team], member)
	}
}

func (m *mockQuay) DeleteTeam(org, team string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.teams[org] != nil {
		delete(m.teams[org], team)
	}
	if m.deletedTeams[org] == nil {
		m.deletedTeams[org] = make(map[string]bool)
	}
	m.deletedTeams[org][team] = true
}

func (m *mockQuay) HasTeamMember(org, team, member string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.teams[org] != nil && m.teams[org][team] != nil {
		return m.teams[org][team][member]
	}
	return false
}

func (m *mockQuay) GetTeamMembers(org, team string) []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.teams[org] == nil || m.teams[org][team] == nil {
		return nil
	}
	members := make([]string, 0, len(m.teams[org][team]))
	for name := range m.teams[org][team] {
		members = append(members, name)
	}
	sort.Strings(members)
	return members
}

func (m *mockQuay) DeletePrototype(org, id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.prototypes[org] != nil {
		delete(m.prototypes[org], id)
	}
}

func (m *mockQuay) GetPrototypes(org string) []mockPrototype {
	m.mu.Lock()
	defer m.mu.Unlock()
	var list []mockPrototype
	if orgProtos, ok := m.prototypes[org]; ok {
		for _, p := range orgProtos {
			list = append(list, p)
		}
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}
