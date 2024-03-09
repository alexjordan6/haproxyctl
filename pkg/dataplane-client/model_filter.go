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

// HAProxy filters
type Filter struct {
	// Name of the fcgi-app section this filter will use.
	AppName string `json:"app_name,omitempty"`
	// Filter name that will be used by 'set-bandwidth-limit' actions to reference a specific bandwidth limitation filter
	BandwidthLimitName string `json:"bandwidth_limit_name,omitempty"`
	CacheName string `json:"cache_name,omitempty"`
	// The max number of bytes that can be forwarded over the period. The value must be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes.
	DefaultLimit int32 `json:"default_limit,omitempty"`
	// The default time period used to evaluate the bandwidth limitation rate. It can be specified for per-stream bandwidth limitation filters only. It follows the HAProxy time format and is expressed in milliseconds.
	DefaultPeriod int32 `json:"default_period,omitempty"`
	Index int32 `json:"index"`
	// A sample expression rule. It describes what elements will be analyzed, extracted, combined, and used to select which table entry to update the counters. It must be specified for shared bandwidth limitation filters only.
	Key string `json:"key,omitempty"`
	// The max number of bytes that can be forwarded over the period. The value must be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes.
	Limit int32 `json:"limit,omitempty"`
	// The optional minimum number of bytes forwarded at a time by a stream excluding the last packet that may be smaller. This value can be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes.
	MinSize int32 `json:"min_size,omitempty"`
	SpoeConfig string `json:"spoe_config,omitempty"`
	SpoeEngine string `json:"spoe_engine,omitempty"`
	// An optional table to be used instead of the default one, which is the stick-table declared in the current proxy. It can be specified for shared bandwidth limitation filters only.
	Table string `json:"table,omitempty"`
	TraceHexdump bool `json:"trace_hexdump,omitempty"`
	TraceName string `json:"trace_name,omitempty"`
	TraceRndForwarding bool `json:"trace_rnd_forwarding,omitempty"`
	TraceRndParsing bool `json:"trace_rnd_parsing,omitempty"`
	Type_ string `json:"type"`
}