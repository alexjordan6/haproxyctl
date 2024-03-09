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

// Define a pattern used to create an entry in a stickiness table or matching condition or associate a user to a server.
type StickRule struct {
	Cond string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Index int32 `json:"index"`
	Pattern string `json:"pattern"`
	Table string `json:"table,omitempty"`
	Type_ string `json:"type"`
}