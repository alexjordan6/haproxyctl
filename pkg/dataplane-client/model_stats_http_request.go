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

type StatsHttpRequest struct {
	Cond string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	Realm string `json:"realm,omitempty"`
	Type_ string `json:"type"`
}