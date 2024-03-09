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

// Runtime transient server properties
type RuntimeServer struct {
	Address          string `json:"address,omitempty"`
	AdminState       string `json:"admin_state,omitempty"`
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	OperationalState string `json:"operational_state,omitempty"`
	Port             int32  `json:"port,omitempty"`
}
