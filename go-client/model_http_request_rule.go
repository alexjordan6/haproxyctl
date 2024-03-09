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

// HAProxy HTTP request rule configuration (corresponds to http-request directives)
type HttpRequestRule struct {
	AclFile string `json:"acl_file,omitempty"`
	AclKeyfmt string `json:"acl_keyfmt,omitempty"`
	AuthRealm string `json:"auth_realm,omitempty"`
	BandwidthLimitLimit string `json:"bandwidth_limit_limit,omitempty"`
	BandwidthLimitName string `json:"bandwidth_limit_name,omitempty"`
	BandwidthLimitPeriod string `json:"bandwidth_limit_period,omitempty"`
	CacheName string `json:"cache_name,omitempty"`
	CaptureId int32 `json:"capture_id,omitempty"`
	CaptureLen int32 `json:"capture_len,omitempty"`
	CaptureSample string `json:"capture_sample,omitempty"`
	Cond string `json:"cond,omitempty"`
	CondTest string `json:"cond_test,omitempty"`
	DenyStatus int32 `json:"deny_status,omitempty"`
	Expr string `json:"expr,omitempty"`
	HdrFormat string `json:"hdr_format,omitempty"`
	HdrMatch string `json:"hdr_match,omitempty"`
	HdrMethod string `json:"hdr_method,omitempty"`
	HdrName string `json:"hdr_name,omitempty"`
	HintFormat string `json:"hint_format,omitempty"`
	HintName string `json:"hint_name,omitempty"`
	Index int32 `json:"index"`
	LogLevel string `json:"log_level,omitempty"`
	LuaAction string `json:"lua_action,omitempty"`
	LuaParams string `json:"lua_params,omitempty"`
	MapFile string `json:"map_file,omitempty"`
	MapKeyfmt string `json:"map_keyfmt,omitempty"`
	MapValuefmt string `json:"map_valuefmt,omitempty"`
	MarkValue string `json:"mark_value,omitempty"`
	MethodFmt string `json:"method_fmt,omitempty"`
	NiceValue int32 `json:"nice_value,omitempty"`
	Normalizer string `json:"normalizer,omitempty"`
	NormalizerFull bool `json:"normalizer_full,omitempty"`
	NormalizerStrict bool `json:"normalizer_strict,omitempty"`
	PathFmt string `json:"path_fmt,omitempty"`
	PathMatch string `json:"path_match,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	QueryFmt string `json:"query-fmt,omitempty"`
	RedirCode int32 `json:"redir_code,omitempty"`
	RedirOption string `json:"redir_option,omitempty"`
	RedirType string `json:"redir_type,omitempty"`
	RedirValue string `json:"redir_value,omitempty"`
	Resolvers string `json:"resolvers,omitempty"`
	ReturnContent string `json:"return_content,omitempty"`
	ReturnContentFormat string `json:"return_content_format,omitempty"`
	ReturnContentType string `json:"return_content_type,omitempty"`
	ReturnHdrs []ReturnHeader `json:"return_hdrs,omitempty"`
	ReturnStatusCode int32 `json:"return_status_code,omitempty"`
	ScExpr string `json:"sc_expr,omitempty"`
	ScId int32 `json:"sc_id,omitempty"`
	ScIdx int32 `json:"sc_idx,omitempty"`
	ScInt int32 `json:"sc_int,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	SpoeEngine string `json:"spoe_engine,omitempty"`
	SpoeGroup string `json:"spoe_group,omitempty"`
	StrictMode string `json:"strict_mode,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	TimeoutType string `json:"timeout_type,omitempty"`
	TosValue string `json:"tos_value,omitempty"`
	TrackSc0Key string `json:"track-sc0-key,omitempty"`
	TrackSc0Table string `json:"track-sc0-table,omitempty"`
	TrackSc1Key string `json:"track-sc1-key,omitempty"`
	TrackSc1Table string `json:"track-sc1-table,omitempty"`
	TrackSc2Key string `json:"track-sc2-key,omitempty"`
	TrackSc2Table string `json:"track-sc2-table,omitempty"`
	TrackScKey string `json:"track_sc_key,omitempty"`
	TrackScStickCounter int32 `json:"track_sc_stick_counter,omitempty"`
	TrackScTable string `json:"track_sc_table,omitempty"`
	Type_ string `json:"type"`
	UriFmt string `json:"uri-fmt,omitempty"`
	UriMatch string `json:"uri-match,omitempty"`
	VarExpr string `json:"var_expr,omitempty"`
	VarFormat string `json:"var_format,omitempty"`
	VarName string `json:"var_name,omitempty"`
	VarScope string `json:"var_scope,omitempty"`
	WaitAtLeast int32 `json:"wait_at_least,omitempty"`
	WaitTime int32 `json:"wait_time,omitempty"`
}
