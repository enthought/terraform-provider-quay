package provider

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

// notFoundBody is the shape the sync-now endpoint returns while a sync is
// already in progress.
const notFoundBody = `{"detail": "Not Found", "error_message": "Not Found", "error_type": "not_found", "title": "not_found", "type": "https://quay.example.com/api/v1/error/not_found", "status": 404}`

func newTestQuayClient(t *testing.T, statusCode int, body string) *quay_api.APIClient {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	configuration := &quay_api.Configuration{
		DefaultHeader: map[string]string{"Authorization": "Bearer test"},
		Servers: quay_api.ServerConfigurations{
			{URL: server.URL, Description: "test"},
		},
	}
	return quay_api.NewAPIClient(configuration)
}

func TestIsSyncAlreadyRunning(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		body       string
		want       bool
	}{
		{
			name:       "404 means sync already in progress",
			statusCode: http.StatusNotFound,
			body:       notFoundBody,
			want:       true,
		},
		{
			name:       "400 is a genuine error",
			statusCode: http.StatusBadRequest,
			body:       `{"detail": "bad request", "error_message": "bad request", "error_type": "bad_request", "title": "bad_request", "type": "https://example.com/api/v1/error/bad_request", "status": 400}`,
			want:       false,
		},
		{
			name:       "500 is a genuine error",
			statusCode: http.StatusInternalServerError,
			body:       `{"detail": "boom", "error_message": "boom", "error_type": "internal", "title": "internal", "type": "https://example.com/api/v1/error/internal", "status": 500}`,
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newTestQuayClient(t, tt.statusCode, tt.body)
			_, err := client.RepositoryAPI.SyncRepoMirrorNow(context.Background(), "org/repo").Execute()
			if err == nil {
				t.Fatalf("expected an error from sync-now with status %d", tt.statusCode)
			}
			if got := isSyncAlreadyRunning(err); got != tt.want {
				t.Fatalf("isSyncAlreadyRunning() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("non-api error", func(t *testing.T) {
		if got := isSyncAlreadyRunning(errors.New("network down")); got {
			t.Fatal("isSyncAlreadyRunning() = true for a plain error, want false")
		}
	})
}

func TestTriggerMirrorSyncIgnoresInflightSync(t *testing.T) {
	client := newTestQuayClient(t, http.StatusNotFound, notFoundBody)
	r := &repositoryResource{client: client}
	if diags := r.triggerMirrorSync(context.Background(), "org/repo"); diags.HasError() {
		t.Fatalf("expected a 404 from sync-now (sync in progress) to be ignored, got diagnostics: %v", diags)
	}
}

func TestTriggerMirrorSyncSurfacesOtherErrors(t *testing.T) {
	badRequestBody := `{"detail": "bad request", "error_message": "bad request", "error_type": "bad_request", "title": "bad_request", "type": "https://example.com/api/v1/error/bad_request", "status": 400}`
	client := newTestQuayClient(t, http.StatusBadRequest, badRequestBody)
	r := &repositoryResource{client: client}
	diags := r.triggerMirrorSync(context.Background(), "org/repo")
	if !diags.HasError() {
		t.Fatal("expected a 400 from sync-now to surface as an error")
	}
}

func TestBuildMirrorValuePreservesSyncTimeoutWhenUnstored(t *testing.T) {
	ctx := context.Background()

	cases := []struct {
		name      string
		stored    *int32
		preserved types.Int64
		want      int64
		wantNull  bool
	}{
		{name: "platform returns the value", stored: ptrInt32(900), preserved: types.Int64Value(600), want: 900},
		{name: "platform omits it, keep prior", stored: nil, preserved: types.Int64Value(600), want: 600},
		{name: "platform omits it, no prior (import)", stored: nil, preserved: types.Int64Null(), wantNull: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := quay_api.NewMirrorWithDefaults()
			if tc.stored != nil {
				m.SetSkopeoTimeoutInterval(*tc.stored)
			}
			mv, d := buildMirrorValueFromAPI(ctx, m, types.StringNull(), true, tc.preserved)
			if d.HasError() {
				t.Fatalf("unexpected diagnostics: %v", d)
			}
			if tc.wantNull {
				if !mv.SyncTimeout.IsNull() {
					t.Fatalf("sync_timeout = %d, want null", mv.SyncTimeout.ValueInt64())
				}
				return
			}
			if mv.SyncTimeout.ValueInt64() != tc.want {
				t.Fatalf("sync_timeout = %d, want %d", mv.SyncTimeout.ValueInt64(), tc.want)
			}
		})
	}
}

func ptrInt32(v int32) *int32 { return &v }

func TestIsRepoAlreadyExists(t *testing.T) {
	existsBody := `{"detail": "Repository already exists", "error_message": "Repository already exists", "error_type": "invalid_request", "title": "invalid_request", "type": "https://example.com/api/v1/error/invalid_request", "status": 400}`
	tests := []struct {
		name       string
		statusCode int
		body       string
		want       bool
	}{
		{name: "400 already exists", statusCode: http.StatusBadRequest, body: existsBody, want: true},
		{name: "400 other error", statusCode: http.StatusBadRequest, body: `{"detail": "invalid name", "error_message": "invalid name", "error_type": "invalid_request", "title": "invalid_request", "type": "https://example.com/api/v1/error/invalid_request", "status": 400}`, want: false},
		{name: "500 error", statusCode: http.StatusInternalServerError, body: `{"detail": "boom", "error_message": "boom", "error_type": "internal", "title": "internal", "type": "https://example.com/api/v1/error/internal", "status": 500}`, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newTestQuayClient(t, tt.statusCode, tt.body)
			_, err := client.RepositoryAPI.CreateRepo(context.Background()).Body(quay_api.NewRepo{Repository: "r", Visibility: "public"}).Execute()
			if err == nil {
				t.Fatalf("expected an error with status %d", tt.statusCode)
			}
			if got := isRepoAlreadyExists(err); got != tt.want {
				t.Fatalf("isRepoAlreadyExists() = %v, want %v", got, tt.want)
			}
		})
	}
	t.Run("non-api error", func(t *testing.T) {
		if isRepoAlreadyExists(errors.New("network down")) {
			t.Fatal("isRepoAlreadyExists() = true for a plain error, want false")
		}
	})
}

func TestRepoState(t *testing.T) {
	t.Run("returns mapped state", func(t *testing.T) {
		client := newTestQuayClient(t, http.StatusOK, `{"name": "r", "namespace": "main", "state": "MIRROR", "is_public": true}`)
		got, ok := repoState(client, context.Background(), "main/r")
		if !ok {
			t.Fatal("repoState() ok = false, want true")
		}
		if got != repoStateMirror {
			t.Fatalf("repoState() = %q, want %q", got, repoStateMirror)
		}
	})
	t.Run("not found", func(t *testing.T) {
		client := newTestQuayClient(t, http.StatusNotFound, notFoundBody)
		if got, ok := repoState(client, context.Background(), "main/missing"); ok {
			t.Fatalf("repoState() ok = true for 404, want false (state=%q)", got)
		}
	})
}

func newTestQuayClientHandler(t *testing.T, h http.HandlerFunc) *quay_api.APIClient {
	t.Helper()
	server := httptest.NewServer(h)
	t.Cleanup(server.Close)
	configuration := &quay_api.Configuration{
		DefaultHeader: map[string]string{"Authorization": "Bearer test"},
		Servers: quay_api.ServerConfigurations{
			{URL: server.URL, Description: "test"},
		},
	}
	return quay_api.NewAPIClient(configuration)
}

func discoverySpec(withSkopeoTimeout bool) string {
	props := `{"external_reference":{"type":"string"},"sync_interval":{"type":"integer"}}`
	if withSkopeoTimeout {
		props = `{"external_reference":{"type":"string"},"sync_interval":{"type":"integer"},"skopeo_timeout_interval":{"type":"integer"}}`
	}
	return `{"definitions":{"CreateMirrorConfig":{"properties":` + props + `},"UpdateMirrorConfig":{"properties":` + props + `}}}`
}

func TestProbeSyncTimeoutSupport(t *testing.T) {
	cases := []struct {
		name   string
		status int
		body   string
		want   bool
	}{
		{name: "spec documents the field", status: http.StatusOK, body: discoverySpec(true), want: true},
		{name: "spec lacks the field", status: http.StatusOK, body: discoverySpec(false), want: false},
		{name: "spec has no mirror definitions", status: http.StatusOK, body: `{"definitions":{}}`, want: true},
		{name: "spec invalid json", status: http.StatusOK, body: `not json`, want: true},
		{name: "probe error", status: http.StatusInternalServerError, body: `boom`, want: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestQuayClientHandler(t, func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			})
			if got := probeSyncTimeoutSupport(client, context.Background()); got != tc.want {
				t.Fatalf("probeSyncTimeoutSupport() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSyncTimeoutSupportedCachesResult(t *testing.T) {
	calls := 0
	client := newTestQuayClientHandler(t, func(w http.ResponseWriter, _ *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(discoverySpec(true)))
	})
	if !syncTimeoutSupported(client, context.Background()) {
		t.Fatal("syncTimeoutSupported() = false, want true")
	}
	if !syncTimeoutSupported(client, context.Background()) {
		t.Fatal("syncTimeoutSupported() = false, want true (cached)")
	}
	if calls != 1 {
		t.Fatalf("discovery endpoint called %d times, want 1", calls)
	}
}
