/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package dataplane

type ProcessInfo struct {
	Error_     string           `json:"error,omitempty"`
	Info       *ProcessInfoItem `json:"info,omitempty"`
	RuntimeAPI string           `json:"runtimeAPI,omitempty"`
}
