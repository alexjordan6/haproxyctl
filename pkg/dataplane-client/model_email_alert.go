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

// Send emails for important log messages.
type EmailAlert struct {
	From       string `json:"from"`
	Level      string `json:"level,omitempty"`
	Mailers    string `json:"mailers"`
	Myhostname string `json:"myhostname,omitempty"`
	To         string `json:"to"`
}
