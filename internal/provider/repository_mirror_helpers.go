// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Modifications copyright (c) Enthought, Inc.
// SPDX-License-Identifier:	BSD-3-Clause

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/enthought/terraform-provider-quay/quay_api"
	"terraform-provider-quay/internal/datasource_repository"
	"terraform-provider-quay/internal/resource_repository"
)

const (
	repoStateNormal   = "normal"
	repoStateReadonly = "readonly"
	repoStateMirror   = "mirror"
)

// mirrorStateToAPI maps a provider state value to the Quay API value.
func mirrorStateToAPI(state string) string {
	switch state {
	case repoStateMirror:
		return "MIRROR"
	case repoStateReadonly:
		// This platform's changestate endpoint spells the value READ_ONLY (with an
		// underscore); READONLY is rejected by the request schema.
		return "READ_ONLY"
	default:
		return "NORMAL"
	}
}

// mirrorStateFromAPI maps a Quay API state value to a provider state value.
func mirrorStateFromAPI(state string) string {
	switch state {
	case "MIRROR":
		return repoStateMirror
	case "READ_ONLY", "READONLY":
		return repoStateReadonly
	default:
		return repoStateNormal
	}
}

// validateMirrorConsistency enforces the rule that a mirror block is required when the
// repository state is "mirror" and forbidden otherwise, and that the block's required
// attributes are present. The platform mandates upstream_repository, robot_user and a
// root_rule (tag patterns) on mirror create, but these are declared optional in the schema
// because the framework validates required attributes even inside an absent block, which
// would make "transition out of mirror" (removing the block) impossible. Their presence is
// therefore enforced here at plan/apply time instead.
func validateMirrorConsistency(state types.String, mirror resource_repository.MirrorValue) diag.Diagnostics {
	var diags diag.Diagnostics
	// state can be unknown during config validation (e.g. derived from for_each),
	// in which case ValueString() returns "" and the consistency rule cannot be
	// evaluated. The check is re-run in Create/Update where state is known.
	if state.IsUnknown() {
		return diags
	}
	switch state.ValueString() {
	case repoStateMirror:
		if mirror.IsNull() {
			diags.AddError(
				"Missing mirror configuration",
				"A \"mirror\" block must be provided when state is \"mirror\".",
			)
			return diags
		}
		diags.Append(validateMirrorRequiredAttributes(mirror)...)
	default:
		if !mirror.IsNull() {
			diags.AddError(
				"Invalid mirror configuration",
				"The \"mirror\" block is only valid when state is \"mirror\". Remove it or set state to \"mirror\".",
			)
		}
	}
	return diags
}

// validateMirrorRequiredAttributes checks the mirror block for the attributes the platform
// requires. It reports each missing attribute so the operator can fix them in one pass.
func validateMirrorRequiredAttributes(mirror resource_repository.MirrorValue) diag.Diagnostics {
	var diags diag.Diagnostics

	if mirror.UpstreamRepository.IsNull() {
		diags.AddAttributeError(
			path.Root("mirror").AtName("upstream_repository"),
			"Missing mirror attribute",
			"\"upstream_repository\" is required when state is \"mirror\".",
		)
	}
	if mirror.RobotUser.IsNull() {
		diags.AddAttributeError(
			path.Root("mirror").AtName("robot_user"),
			"Missing mirror attribute",
			"\"robot_user\" is required when state is \"mirror\".",
		)
	}
	if mirror.RootRule.IsNull() {
		diags.AddAttributeError(
			path.Root("mirror").AtName("root_rule"),
			"Missing mirror attribute",
			"\"root_rule\" is required when state is \"mirror\".",
		)
		return diags
	}
	kind, ok := mirror.RootRule.Attributes()["rule_kind"].(basetypes.StringValue)
	if !ok || kind.IsNull() {
		diags.AddAttributeError(
			path.Root("mirror").AtName("root_rule").AtName("rule_kind"),
			"Missing mirror attribute",
			"\"root_rule.rule_kind\" is required when state is \"mirror\".",
		)
	}
	return diags
}

// buildRootRuleObject builds the root_rule object attribute value from a rule kind and values.
// The attribute types must match the type of the enclosing schema (resource or data source).
func buildRootRuleObject(ctx context.Context, attrTypes map[string]attr.Type, kind *string, values *[]string) basetypes.ObjectValue {
	attrs := map[string]attr.Value{}
	if kind == nil {
		attrs["rule_kind"] = types.StringNull()
	} else {
		attrs["rule_kind"] = types.StringValue(*kind)
	}
	if values == nil {
		attrs["rule_value"] = types.ListNull(types.StringType)
	} else {
		elems := make([]attr.Value, len(*values))
		for i, v := range *values {
			elems[i] = types.StringValue(v)
		}
		lv, _ := types.ListValue(types.StringType, elems)
		attrs["rule_value"] = lv
	}
	return types.ObjectValueMust(attrTypes, attrs)
}

// mirrorRecordExists reports whether the repository already has a mirror record.
// On this platform the record persists once created (DELETE is a no-op), so it
// determines whether a mirror configuration change must be a PUT (existing) or a
// POST (new record).
func mirrorRecordExists(client *quay_api.APIClient, ctx context.Context, repoRef string) bool {
	httpRes, err := client.RepositoryAPI.GetRepoMirror(ctx, repoRef).Execute()
	if err != nil {
		// A 404 means no record was ever created; any other error is treated as
		// "exists" so the caller attempts a PUT and surfaces the real error.
		return !strings.Contains(err.Error(), "404")
	}
	defer httpRes.Body.Close()
	return true
}

// mirrorValueToAPI converts a provider MirrorValue into the Quay API mirror record.
// Returns nil when the value is null.
func mirrorValueToAPI(ctx context.Context, m resource_repository.MirrorValue, includeSyncTimeout bool) (*quay_api.Mirror, diag.Diagnostics) {
	var diags diag.Diagnostics
	if m.IsNull() {
		return nil, diags
	}

	// mirror_type is intentionally left empty: the platform assigns it (always PULL) and
	// rejects create/update bodies that carry it.
	mirror := quay_api.NewMirrorWithDefaults()
	mirror.SetExternalReference(m.UpstreamRepository.ValueString())
	mirror.SetRobotUsername(m.RobotUser.ValueString())

	if !m.UpstreamUsername.IsNull() {
		v := m.UpstreamUsername.ValueString()
		mirror.SetExternalRegistryUsername(v)
	}
	if !m.UpstreamPassword.IsNull() {
		v := m.UpstreamPassword.ValueString()
		mirror.SetExternalRegistryPassword(v)
	}
	if !m.SyncInterval.IsNull() {
		v := m.SyncInterval.ValueInt64()
		mirror.SetSyncInterval(int32(v))
	}
	// skopeo_timeout_interval is mandatory on mirror create in newer Quay versions
	// (the schema default 600 covers the null case), but older versions (e.g. 3.13)
	// crash when it is present. The caller retries without it when the registry
	// rejects the request.
	if includeSyncTimeout && !m.SyncTimeout.IsNull() {
		v := m.SyncTimeout.ValueInt64()
		mirror.SetSkopeoTimeoutInterval(int32(v))
	}

	cfg := quay_api.NewMirrorExternalRegistryConfig()
	cfg.SetVerifyTls(m.VerifyTls.ValueBool())
	cfg.SetUnsignedImages(m.UnsignedImages.ValueBool())
	var proxy quay_api.MirrorProxy
	proxySet := false
	if !m.HttpProxy.IsNull() {
		v := m.HttpProxy.ValueString()
		proxy.SetHttpProxy(v)
		proxySet = true
	}
	if !m.HttpsProxy.IsNull() {
		v := m.HttpsProxy.ValueString()
		proxy.SetHttpsProxy(v)
		proxySet = true
	}
	if !m.NoProxy.IsNull() {
		v := m.NoProxy.ValueString()
		proxy.SetNoProxy(v)
		proxySet = true
	}
	if proxySet {
		cfg.SetProxy(proxy)
	}
	mirror.SetExternalRegistryConfig(*cfg)

	if !m.RootRule.IsNull() {
		rr, d := resource_repository.NewRootRuleValue(resource_repository.RootRuleValue{}.AttributeTypes(ctx), m.RootRule.Attributes())
		diags.Append(d...)
		if !diags.HasError() {
			var values []string
			rr.RuleValue.ElementsAs(ctx, &values, false)
			var kind *string
			if !rr.RuleKind.IsNull() {
				k := rr.RuleKind.ValueString()
				kind = &k
			}
			var valuesPtr *[]string
			if !rr.RuleValue.IsNull() {
				valuesPtr = &values
			}
			mirrorRootRule := quay_api.NewMirrorRootRule()
			if kind != nil {
				mirrorRootRule.SetRuleKind(*kind)
			}
			if valuesPtr != nil {
				mirrorRootRule.SetRuleValue(*valuesPtr)
			}
			mirror.SetRootRule(*mirrorRootRule)
		}
	}

	return mirror, diags
}

// syncOnChangeFromMirror returns the previously-configured sync_on_change flag, or the
// documented default (true) when there is no previous mirror block (e.g. during import).
func syncOnChangeFromMirror(prev resource_repository.MirrorValue) bool {
	if !prev.IsNull() && !prev.SyncOnChange.IsNull() {
		return prev.SyncOnChange.ValueBool()
	}
	return true
}

// buildMirrorValueFromAPI converts a Quay API mirror record into a provider MirrorValue.
// preservePassword is the previously-configured upstream password (the API never returns it),
// so it is kept in state instead of being read from the response. syncOnChange is a
// provider-side flag likewise not stored by the platform, so the previous value is carried
// forward. preserveSyncTimeout is the previously-configured sync timeout; Quay versions
// that do not store skopeo_timeout_interval omit it from the response, so the previous
// value is kept to avoid a permanent plan diff.
func buildMirrorValueFromAPI(ctx context.Context, mirror *quay_api.Mirror, preservePassword types.String, syncOnChange bool, preserveSyncTimeout types.Int64) (resource_repository.MirrorValue, diag.Diagnostics) {
	attrs := map[string]attr.Value{}

	attrs["upstream_repository"] = types.StringValue(mirror.GetExternalReference())
	if mirror.ExternalRegistryUsername != nil {
		attrs["upstream_username"] = types.StringValue(*mirror.ExternalRegistryUsername)
	} else {
		attrs["upstream_username"] = types.StringNull()
	}
	attrs["upstream_password"] = preservePassword
	attrs["sync_on_change"] = types.BoolValue(syncOnChange)
	if mirror.RobotUsername != nil {
		attrs["robot_user"] = types.StringValue(*mirror.RobotUsername)
	} else {
		attrs["robot_user"] = types.StringNull()
	}

	if mirror.SyncInterval != nil {
		attrs["sync_interval"] = types.Int64Value(int64(*mirror.SyncInterval))
	} else {
		attrs["sync_interval"] = types.Int64Null()
	}
	if mirror.SkopeoTimeoutInterval != nil {
		attrs["sync_timeout"] = types.Int64Value(int64(*mirror.SkopeoTimeoutInterval))
	} else if !preserveSyncTimeout.IsNull() && !preserveSyncTimeout.IsUnknown() {
		// The platform (e.g. Quay < 3.14) does not store skopeo_timeout_interval,
		// so keep the previously-configured value to avoid a permanent plan diff.
		attrs["sync_timeout"] = preserveSyncTimeout
	} else {
		attrs["sync_timeout"] = types.Int64Null()
	}

	verifyTls := true
	unsignedImages := false
	if mirror.ExternalRegistryConfig != nil {
		if mirror.ExternalRegistryConfig.VerifyTls != nil {
			verifyTls = *mirror.ExternalRegistryConfig.VerifyTls
		}
		if mirror.ExternalRegistryConfig.UnsignedImages != nil {
			unsignedImages = *mirror.ExternalRegistryConfig.UnsignedImages
		}
	}
	attrs["verify_tls"] = types.BoolValue(verifyTls)
	attrs["unsigned_images"] = types.BoolValue(unsignedImages)

	attrs["http_proxy"] = types.StringNull()
	attrs["https_proxy"] = types.StringNull()
	attrs["no_proxy"] = types.StringNull()
	var kind *string
	var values *[]string
	if mirror.ExternalRegistryConfig != nil && mirror.ExternalRegistryConfig.Proxy != nil {
		p := *mirror.ExternalRegistryConfig.Proxy
		if p.HttpProxy != nil {
			attrs["http_proxy"] = types.StringValue(*p.HttpProxy)
		}
		if p.HttpsProxy != nil {
			attrs["https_proxy"] = types.StringValue(*p.HttpsProxy)
		}
		if p.NoProxy != nil {
			attrs["no_proxy"] = types.StringValue(*p.NoProxy)
		}
	}
	if mirror.RootRule != nil {
		kind = mirror.RootRule.RuleKind
		values = mirror.RootRule.RuleValue
	}
	attrs["root_rule"] = buildRootRuleObject(ctx, resource_repository.RootRuleValue{}.AttributeTypes(ctx), kind, values)

	return resource_repository.NewMirrorValueMust(resource_repository.MirrorValue{}.AttributeTypes(ctx), attrs), nil
}

// buildDataSourceMirrorValue converts a Quay API mirror record into the data source MirrorValue.
// The data source does not expose the upstream password (the platform never returns it).
func buildDataSourceMirrorValue(ctx context.Context, mirror *quay_api.Mirror) (datasource_repository.MirrorValue, diag.Diagnostics) {
	attrs := map[string]attr.Value{}

	attrs["upstream_repository"] = types.StringValue(mirror.GetExternalReference())
	if mirror.RobotUsername != nil {
		attrs["robot_user"] = types.StringValue(*mirror.RobotUsername)
	} else {
		attrs["robot_user"] = types.StringNull()
	}

	if mirror.SyncInterval != nil {
		attrs["sync_interval"] = types.Int64Value(int64(*mirror.SyncInterval))
	} else {
		attrs["sync_interval"] = types.Int64Null()
	}
	if mirror.SkopeoTimeoutInterval != nil {
		attrs["sync_timeout"] = types.Int64Value(int64(*mirror.SkopeoTimeoutInterval))
	} else {
		attrs["sync_timeout"] = types.Int64Null()
	}

	verifyTls := true
	unsignedImages := false
	if mirror.ExternalRegistryConfig != nil {
		if mirror.ExternalRegistryConfig.VerifyTls != nil {
			verifyTls = *mirror.ExternalRegistryConfig.VerifyTls
		}
		if mirror.ExternalRegistryConfig.UnsignedImages != nil {
			unsignedImages = *mirror.ExternalRegistryConfig.UnsignedImages
		}
	}
	attrs["verify_tls"] = types.BoolValue(verifyTls)
	attrs["unsigned_images"] = types.BoolValue(unsignedImages)

	attrs["http_proxy"] = types.StringNull()
	attrs["https_proxy"] = types.StringNull()
	attrs["no_proxy"] = types.StringNull()
	var kind *string
	var values *[]string
	if mirror.ExternalRegistryConfig != nil && mirror.ExternalRegistryConfig.Proxy != nil {
		p := *mirror.ExternalRegistryConfig.Proxy
		if p.HttpProxy != nil {
			attrs["http_proxy"] = types.StringValue(*p.HttpProxy)
		}
		if p.HttpsProxy != nil {
			attrs["https_proxy"] = types.StringValue(*p.HttpsProxy)
		}
		if p.NoProxy != nil {
			attrs["no_proxy"] = types.StringValue(*p.NoProxy)
		}
	}
	if mirror.RootRule != nil {
		kind = mirror.RootRule.RuleKind
		values = mirror.RootRule.RuleValue
	}
	attrs["root_rule"] = buildRootRuleObject(ctx, datasource_repository.RootRuleValue{}.AttributeTypes(ctx), kind, values)

	return datasource_repository.NewMirrorValueMust(datasource_repository.MirrorValue{}.AttributeTypes(ctx), attrs), nil
}
