/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs. 
 *
 * API version: 3.0
 * Contact: support@haproxy.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// HAProxy log forward configuration
type LogForward struct {
	Backlog int32 `json:"backlog,omitempty"`
	Maxconn int32 `json:"maxconn,omitempty"`
	Name string `json:"name"`
	TimeoutClient int32 `json:"timeout_client,omitempty"`
}