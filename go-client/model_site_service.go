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

type SiteService struct {
	HttpConnectionMode string `json:"http_connection_mode,omitempty"`
	Listeners []Bind `json:"listeners,omitempty"`
	Maxconn int32 `json:"maxconn,omitempty"`
	Mode string `json:"mode,omitempty"`
}
