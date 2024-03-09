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

// General use file
type GeneralFile struct {
	Description string `json:"description,omitempty"`
	File string `json:"file,omitempty"`
	Id string `json:"id,omitempty"`
	// File size in bytes.
	Size int32 `json:"size,omitempty"`
	StorageName string `json:"storage_name,omitempty"`
}