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

type HttpCheck struct {
	Addr string `json:"addr,omitempty"`
	Alpn string `json:"alpn,omitempty"`
	Body string `json:"body,omitempty"`
	BodyLogFormat string `json:"body_log_format,omitempty"`
	CheckComment string `json:"check_comment,omitempty"`
	Default_ bool `json:"default,omitempty"`
	ErrorStatus string `json:"error_status,omitempty"`
	ExclamationMark bool `json:"exclamation_mark,omitempty"`
	Headers []ReturnHeader `json:"headers,omitempty"`
	Index int32 `json:"index"`
	Linger bool `json:"linger,omitempty"`
	Match string `json:"match,omitempty"`
	Method string `json:"method,omitempty"`
	MinRecv int32 `json:"min_recv,omitempty"`
	OkStatus string `json:"ok_status,omitempty"`
	OnError string `json:"on_error,omitempty"`
	OnSuccess string `json:"on_success,omitempty"`
	Pattern string `json:"pattern,omitempty"`
	Port int32 `json:"port,omitempty"`
	PortString string `json:"port_string,omitempty"`
	Proto string `json:"proto,omitempty"`
	SendProxy bool `json:"send_proxy,omitempty"`
	Sni string `json:"sni,omitempty"`
	Ssl bool `json:"ssl,omitempty"`
	StatusCode string `json:"status-code,omitempty"`
	ToutStatus string `json:"tout_status,omitempty"`
	Type_ string `json:"type"`
	Uri string `json:"uri,omitempty"`
	UriLogFormat string `json:"uri_log_format,omitempty"`
	VarExpr string `json:"var_expr,omitempty"`
	VarFormat string `json:"var_format,omitempty"`
	VarName string `json:"var_name,omitempty"`
	VarScope string `json:"var_scope,omitempty"`
	Version string `json:"version,omitempty"`
	ViaSocks4 bool `json:"via_socks4,omitempty"`
}
