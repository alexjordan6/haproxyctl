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

// Nameserver used in Runtime DNS configuration
type Nameserver struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Port    int32  `json:"port,omitempty"`
}
