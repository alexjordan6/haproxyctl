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

// HAProxy global configuration
type Global struct {
	Anonkey int32 `json:"anonkey,omitempty"`
	BusyPolling bool `json:"busy_polling,omitempty"`
	CaBase string `json:"ca_base,omitempty"`
	Chroot string `json:"chroot,omitempty"`
	CloseSpreadTime int32 `json:"close_spread_time,omitempty"`
	ClusterSecret string `json:"cluster_secret,omitempty"`
	CpuMaps []GlobalCpuMaps `json:"cpu_maps,omitempty"`
	CrtBase string `json:"crt_base,omitempty"`
	Daemon string `json:"daemon,omitempty"`
	DefaultPath *GlobalDefaultPath `json:"default_path,omitempty"`
	Description string `json:"description,omitempty"`
	DeviceAtlasOptions *GlobalDeviceAtlasOptions `json:"device_atlas_options,omitempty"`
	ExposeExperimentalDirectives bool `json:"expose_experimental_directives,omitempty"`
	ExternalCheck bool `json:"external_check,omitempty"`
	FiftyOneDegreesOptions *GlobalFiftyOneDegreesOptions `json:"fifty_one_degrees_options,omitempty"`
	Gid int32 `json:"gid,omitempty"`
	Grace int32 `json:"grace,omitempty"`
	Group string `json:"group,omitempty"`
	H1CaseAdjust []GlobalH1CaseAdjust `json:"h1_case_adjust,omitempty"`
	H1CaseAdjustFile string `json:"h1_case_adjust_file,omitempty"`
	H2WorkaroundBogusWebsocketClients bool `json:"h2_workaround_bogus_websocket_clients,omitempty"`
	HardStopAfter int32 `json:"hard_stop_after,omitempty"`
	HttpclientResolversDisabled string `json:"httpclient_resolvers_disabled,omitempty"`
	HttpclientResolversId string `json:"httpclient_resolvers_id,omitempty"`
	HttpclientResolversPrefer string `json:"httpclient_resolvers_prefer,omitempty"`
	HttpclientRetries int32 `json:"httpclient_retries,omitempty"`
	HttpclientSslCaFile string `json:"httpclient_ssl_ca_file,omitempty"`
	HttpclientSslVerify string `json:"httpclient_ssl_verify,omitempty"`
	HttpclientTimeoutConnect int32 `json:"httpclient_timeout_connect,omitempty"`
	InsecureForkWanted bool `json:"insecure_fork_wanted,omitempty"`
	InsecureSetuidWanted bool `json:"insecure_setuid_wanted,omitempty"`
	IssuersChainPath string `json:"issuers_chain_path,omitempty"`
	LimitedQuic bool `json:"limited_quic,omitempty"`
	Localpeer string `json:"localpeer,omitempty"`
	LogSendHostname *GlobalLogSendHostname `json:"log_send_hostname,omitempty"`
	LuaLoadPerThread string `json:"lua_load_per_thread,omitempty"`
	LuaLoads []GlobalLuaLoads `json:"lua_loads,omitempty"`
	LuaPrependPath []GlobalLuaPrependPath `json:"lua_prepend_path,omitempty"`
	MasterWorker bool `json:"master-worker,omitempty"`
	MaxSpreadChecks int32 `json:"max_spread_checks,omitempty"`
	Maxcompcpuusage int32 `json:"maxcompcpuusage,omitempty"`
	Maxcomprate int32 `json:"maxcomprate,omitempty"`
	Maxconn int32 `json:"maxconn,omitempty"`
	Maxconnrate int32 `json:"maxconnrate,omitempty"`
	Maxpipes int32 `json:"maxpipes,omitempty"`
	Maxsessrate int32 `json:"maxsessrate,omitempty"`
	Maxsslconn int32 `json:"maxsslconn,omitempty"`
	Maxsslrate int32 `json:"maxsslrate,omitempty"`
	Maxzlibmem int32 `json:"maxzlibmem,omitempty"`
	MworkerMaxReloads int32 `json:"mworker_max_reloads,omitempty"`
	Nbproc int32 `json:"nbproc,omitempty"`
	Nbthread int32 `json:"nbthread,omitempty"`
	NoQuic bool `json:"no-quic,omitempty"`
	Node string `json:"node,omitempty"`
	Noepoll bool `json:"noepoll,omitempty"`
	Noevports bool `json:"noevports,omitempty"`
	Nogetaddrinfo bool `json:"nogetaddrinfo,omitempty"`
	Nokqueue bool `json:"nokqueue,omitempty"`
	Nopoll bool `json:"nopoll,omitempty"`
	Noreuseport bool `json:"noreuseport,omitempty"`
	Nosplice bool `json:"nosplice,omitempty"`
	NumaCpuMapping string `json:"numa_cpu_mapping,omitempty"`
	Pidfile string `json:"pidfile,omitempty"`
	Pp2NeverSendLocal bool `json:"pp2_never_send_local,omitempty"`
	PreallocFd bool `json:"prealloc-fd,omitempty"`
	Presetenv []GlobalPresetenv `json:"presetenv,omitempty"`
	ProfilingTasks string `json:"profiling_tasks,omitempty"`
	Quiet bool `json:"quiet,omitempty"`
	Resetenv string `json:"resetenv,omitempty"`
	RuntimeApis []interface{} `json:"runtime_apis,omitempty"`
	ServerStateBase string `json:"server_state_base,omitempty"`
	ServerStateFile string `json:"server_state_file,omitempty"`
	SetDumpable bool `json:"set_dumpable,omitempty"`
	SetVar []GlobalSetVar `json:"set_var,omitempty"`
	SetVarFmt []GlobalSetVarFmt `json:"set_var_fmt,omitempty"`
	Setcap string `json:"setcap,omitempty"`
	Setenv []GlobalSetenv `json:"setenv,omitempty"`
	SpreadChecks int32 `json:"spread_checks,omitempty"`
	SslDefaultBindCiphers string `json:"ssl_default_bind_ciphers,omitempty"`
	SslDefaultBindCiphersuites string `json:"ssl_default_bind_ciphersuites,omitempty"`
	SslDefaultBindClientSigalgs string `json:"ssl_default_bind_client_sigalgs,omitempty"`
	SslDefaultBindCurves string `json:"ssl_default_bind_curves,omitempty"`
	SslDefaultBindOptions string `json:"ssl_default_bind_options,omitempty"`
	SslDefaultBindSigalgs string `json:"ssl_default_bind_sigalgs,omitempty"`
	SslDefaultServerCiphers string `json:"ssl_default_server_ciphers,omitempty"`
	SslDefaultServerCiphersuites string `json:"ssl_default_server_ciphersuites,omitempty"`
	SslDefaultServerClientSigalgs string `json:"ssl_default_server_client_sigalgs,omitempty"`
	SslDefaultServerCurves string `json:"ssl_default_server_curves,omitempty"`
	SslDefaultServerOptions string `json:"ssl_default_server_options,omitempty"`
	SslDefaultServerSigalgs string `json:"ssl_default_server_sigalgs,omitempty"`
	SslDhParamFile string `json:"ssl_dh_param_file,omitempty"`
	SslEngines []GlobalSslEngines `json:"ssl_engines,omitempty"`
	SslLoadExtraFiles string `json:"ssl_load_extra_files,omitempty"`
	SslModeAsync string `json:"ssl_mode_async,omitempty"`
	SslPropquery string `json:"ssl_propquery,omitempty"`
	SslProvider string `json:"ssl_provider,omitempty"`
	SslProviderPath string `json:"ssl_provider_path,omitempty"`
	SslServerVerify string `json:"ssl_server_verify,omitempty"`
	SslSkipSelfIssuedCa bool `json:"ssl_skip_self_issued_ca,omitempty"`
	StatsMaxconn int32 `json:"stats_maxconn,omitempty"`
	StatsTimeout int32 `json:"stats_timeout,omitempty"`
	StrictLimits bool `json:"strict_limits,omitempty"`
	ThreadGroupLines []GlobalThreadGroupLines `json:"thread_group_lines,omitempty"`
	ThreadGroups int32 `json:"thread_groups,omitempty"`
	TuneOptions *GlobalTuneOptions `json:"tune_options,omitempty"`
	TuneSslDefaultDhParam int32 `json:"tune_ssl_default_dh_param,omitempty"`
	Uid int32 `json:"uid,omitempty"`
	UlimitN int32 `json:"ulimit_n,omitempty"`
	Unsetenv string `json:"unsetenv,omitempty"`
	User string `json:"user,omitempty"`
	WurflOptions *GlobalWurflOptions `json:"wurfl_options,omitempty"`
	ZeroWarning bool `json:"zero_warning,omitempty"`
}