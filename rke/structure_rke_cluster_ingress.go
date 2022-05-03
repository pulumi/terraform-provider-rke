package rke

import (
	rancher "github.com/rancher/rke/types"
)

// Flatteners

func flattenRKEClusterIngress(in rancher.IngressConfig) []interface{} {
	obj := make(map[string]interface{})

	if len(in.DNSPolicy) > 0 {
		obj["dns_policy"] = in.DNSPolicy
	}

	if len(in.ExtraArgs) > 0 {
		obj["extra_args"] = toMapInterface(in.ExtraArgs)
	}

	if in.HTTPPort > 0 {
		obj["http_port"] = in.HTTPPort
	}

	if in.HTTPSPort > 0 {
		obj["https_port"] = in.HTTPSPort
	}

	if len(in.NetworkMode) > 0 {
		obj["network_mode"] = in.NetworkMode
	}

	if len(in.NodeSelector) > 0 {
		obj["node_selector"] = toMapInterface(in.NodeSelector)
	}

	if len(in.Options) > 0 {
		obj["options"] = toMapInterface(in.Options)
	}

	if len(in.Provider) > 0 {
		obj["provider"] = in.Provider
	}

	if in.DefaultBackend != nil {
		obj["default_backend"] = *in.DefaultBackend
	}

	return []interface{}{obj}
}

// Expanders

func expandRKEClusterIngress(p []interface{}) rancher.IngressConfig {
	obj := rancher.IngressConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["dns_policy"].(string); ok && len(v) > 0 {
		obj.DNSPolicy = v
	}

	if v, ok := in["extra_args"].(map[string]interface{}); ok && len(v) > 0 {
		obj.ExtraArgs = toMapString(v)
	}

	if v, ok := in["http_port"].(int); ok && v > 0 {
		obj.HTTPPort = v
	}

	if v, ok := in["https_port"].(int); ok && v > 0 {
		obj.HTTPSPort = v
	}

	if v, ok := in["network_mode"].(string); ok && len(v) > 0 {
		obj.NetworkMode = v
	}

	if v, ok := in["node_selector"].(map[string]interface{}); ok && len(v) > 0 {
		obj.NodeSelector = toMapString(v)
	}

	if v, ok := in["options"].(map[string]interface{}); ok && len(v) > 0 {
		obj.Options = toMapString(v)
	}

	if v, ok := in["provider"].(string); ok && len(v) > 0 {
		obj.Provider = v
	}

	if v, ok := in["default_backend"].(bool); ok {
		obj.DefaultBackend = &v
	}

	return obj
}
