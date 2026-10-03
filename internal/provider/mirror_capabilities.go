package provider

import (
	"context"
	"encoding/json"
	"io"
	"sync"

	"github.com/enthought/terraform-provider-quay/quay_api"
)

// syncTimeoutSupportCache maps an API client to whether its registry supports
// skopeo_timeout_interval, so the discovery spec is fetched at most once per
// registry per process.
var syncTimeoutSupportCache sync.Map

// syncTimeoutSupported reports whether the registry accepts the mirror
// skopeo_timeout_interval field (exposed as sync_timeout). Some Quay versions
// do not support it and crash the mirror create/update endpoints with an
// internal server error when it is present, so the provider omits the field
// for those versions. The capability is read from the OpenAPI spec served at
// /api/v1/discovery, which both standard Quay and the in-house fork publish.
func syncTimeoutSupported(client *quay_api.APIClient, ctx context.Context) bool {
	if v, ok := syncTimeoutSupportCache.Load(client); ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	supported := probeSyncTimeoutSupport(client, ctx)
	syncTimeoutSupportCache.Store(client, supported)
	return supported
}

// probeSyncTimeoutSupport fetches the registry's OpenAPI spec and inspects the
// mirror create/update request schemas. When the spec is unavailable or does
// not document the mirror endpoints, it assumes the field is supported,
// matching standard Quay, where it is required.
func probeSyncTimeoutSupport(client *quay_api.APIClient, ctx context.Context) bool {
	httpRes, err := client.DiscoveryAPI.Discovery(ctx).Execute()
	if err != nil {
		return true
	}
	defer httpRes.Body.Close()
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return true
	}
	var spec struct {
		Definitions map[string]struct {
			Properties map[string]json.RawMessage `json:"properties"`
		} `json:"definitions"`
	}
	if err := json.Unmarshal(body, &spec); err != nil {
		return true
	}
	documented := false
	for _, name := range []string{"CreateMirrorConfig", "UpdateMirrorConfig"} {
		def, ok := spec.Definitions[name]
		if !ok {
			continue
		}
		documented = true
		if _, present := def.Properties["skopeo_timeout_interval"]; present {
			return true
		}
	}
	if !documented {
		return true
	}
	return false
}
