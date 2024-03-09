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

type ClusterControllerInformation struct {
	Address string `json:"address,omitempty"`
	ApiBasePath string `json:"api_base_path,omitempty"`
	ClusterId string `json:"cluster_id,omitempty"`
	Description string `json:"description,omitempty"`
	LogTargets []ClusterControllerInformationLogTargets `json:"log_targets,omitempty"`
	Name string `json:"name,omitempty"`
	Port int32 `json:"port,omitempty"`
}