# Config

Config is the configuration object used by Tyk to set up various parameters.

**Field: `hostname`** (HostName, `string`)

Force your Gateway to work only on a specific domain name. Can be overridden by API custom domain.

**Field: `listen_address`** (ListenAddress, `string`)

If your machine has multiple network devices or IPs you can force the Gateway to use the IP address you want.

**Field: `listen_port`** (ListenPort, `int`)

Setting this value will change the port that Tyk listens on. Default: 8080.

**Field: `control_api_hostname`** (ControlAPIHostname, `string`)

Custom hostname for the Control API

**Field: `control_api_port`** (ControlAPIPort, `int`)

Set to run your Gateway Control API on a separate port, and protect it behind a firewall if needed. Please make sure you follow this guide when setting the control port https://tyk.io/docs/planning-for-production/#change-your-control-port.

**Field: `secret`** (Secret, `string`)

This should be changed as soon as Tyk is installed on your system.
This value is used in every interaction with the Tyk Gateway API. It should be passed along as the X-Tyk-Authorization header in any requests made.
Tyk assumes that you are sensible enough not to expose the management endpoints publicly and to keep this configuration value to yourself.

**Field: `node_secret`** (NodeSecret, `string`)

The shared secret between the Gateway and the Dashboard to ensure that API Definition downloads, heartbeat and Policy loads are from a valid source.

**Field: `pid_file_location`** (PIDFileLocation, `string`)

Linux PID file location. Do not change unless you know what you are doing. Default: /var/run/tyk/tyk-gateway.pid

**Field: `allow_insecure_configs`** (AllowInsecureConfigs, `bool`)

Can be set to disable Dashboard message signature verification. When set to `true`, `public_key_path` can be ignored.

**Field: `public_key_path`** (PublicKeyPath, `string`)

While communicating with the Dashboard. By default, all messages are signed by a private/public key pair. Set path to public key.

**Field: `allow_remote_config`** (AllowRemoteConfig, `bool`)

Allow your Dashboard to remotely set Gateway configuration via the Nodes screen.

**Field: `security`** (Security, [SecurityConfig](#SecurityConfig))

Global Certificate configuration

**Field: `http_server_options`** (HttpServerOptions, [HttpServerOptionsConfig](#HttpServerOptionsConfig))

Gateway HTTP server configuration

**Field: `version_header`** (VersionHeader, `string`)

Expose version header with a given name. Works only for versioned APIs.

**Field: `suppress_redis_signal_reload`** (SuppressRedisSignalReload, `bool`)

Disable dynamic API and Policy reloads, e.g. it will load new changes only on procecss start.

**Field: `hash_keys`** (HashKeys, `bool`)

Enable Key hashing

**Field: `hash_key_function`** (HashKeyFunction, `string`)

Specify the Key hashing algorithm. Possible values: murmur64, murmur128, sha256.

**Field: `basic_auth_hash_key_function`** (BasicAuthHashKeyFunction, `string`)

Specify the Key hashing algorithm for "basic auth". Possible values: murmur64, murmur128, sha256, bcrypt.
Will default to "bcrypt" if not set.

**Field: `hash_key_function_fallback`** (HashKeyFunctionFallback, `[]string`)

Specify your previous key hashing algorithm if you migrated from one algorithm to another.

**Field: `enable_hashed_keys_listing`** (EnableHashedKeysListing, `bool`)

Allows the listing of hashed API keys

**Field: `min_token_length`** (MinTokenLength, `int`)

Minimum API token length

**Field: `template_path`** (TemplatePath, `string`)

Path to error and webhook templates. Defaults to the current binary path.

**Field: `policies`** (Policies, [PoliciesConfig](#PoliciesConfig))

The policies section allows you to define where Tyk can find its policy templates. Policy templates are similar to key definitions in that they allow you to set quotas, access rights and rate limits for keys.
Policies are loaded when Tyk starts and if changed require a hot-reload so they are loaded into memory.
A policy can be defined in a file (Open Source installations) or from the same database as the Dashboard.

**Field: `ports_whitelist`** (PortWhiteList, [PortsWhiteList](#PortsWhiteList))

Defines the ports that will be available for the API services to bind to in the following format: `"{“":“”}"`. Remember to escape JSON strings.
This is a map of protocol to PortWhiteList. This allows per protocol
configurations.

**Field: `disable_ports_whitelist`** (DisablePortWhiteList, `bool`)

Disable port whilisting, essentially allowing you to use any port for your API.

**Field: `app_path`** (AppPath, `string`)

If Tyk is being used in its standard configuration (Open Source installations), then API definitions are stored in the apps folder (by default in /opt/tyk-gateway/apps).
This location is scanned for .json files and re-scanned at startup or reload.
See the API section of the Tyk Gateway API for more details.

**Field: `use_db_app_configs`** (UseDBAppConfigs, `bool`)

If you are a Tyk Pro user, this option will enable polling the Dashboard service for API definitions.
On startup Tyk will attempt to connect and download any relevant application configurations from from your Dashboard instance.
The files are exactly the same as the JSON files on disk with the exception of a BSON ID supplied by the Dashboard service.

**Field: `db_app_conf_options`** (DBAppConfOptions, [DBAppConfOptionsConfig](#DBAppConfOptionsConfig))

This section defines API loading and shard options. Enable these settings to selectively load API definitions on a node from your Dashboard service.

**Field: `storage`** (Storage, [StorageOptionsConf](#StorageOptionsConf))

This section defines your Redis configuration.

**Field: `disable_dashboard_zeroconf`** (DisableDashboardZeroConf, `bool`)

Disable the capability of the Gateway to `autodiscover` the Dashboard through heartbeat messages via Redis.
The goal of zeroconf is auto-discovery, so you do not have to specify the Tyk Dashboard address in your Gateway`tyk.conf` file.
In some specific cases, for example, when the Dashboard is bound to a public domain, not accessible inside an internal network, or similar, `disable_dashboard_zeroconf` can be set to `true`, in favour of directly specifying a Tyk Dashboard address.

**Field: `slave_options`** (SlaveOptions, [SlaveOptionsConfig](#SlaveOptionsConfig))

The `slave_options` allow you to configure the RPC slave connection required for MDCB installations.
These settings must be configured for every RPC slave/worker node.

**Field: `management_node`** (ManagementNode, `bool`)

If set to `true`, distributed rate limiter will be disabled for this node, and it will be excluded from any rate limit calculation.

Note:
  If you set `db_app_conf_options.node_is_segmented` to `true` for multiple Gateway nodes, you should ensure that `management_node` is set to `false`.
  This is to ensure visibility for the management node across all APIs.

**Field: `auth_override`** (AuthOverride, [AuthOverrideConf](#AuthOverrideConf))

This is used as part of the RPC / Hybrid back-end configuration in a Tyk Enterprise installation and isn’t used anywhere else.

**Field: `enable_redis_rolling_limiter`** (EnableRedisRollingLimiter, `bool`)

Redis based rate limiter with fixed window. Provides 100% rate limiting accuracy, but require two additional Redis roundtrip for each request.

**Field: `enable_sentinel_rate_limiter`** (EnableSentinelRateLimiter, `bool`)

To enable, set to `true`. The sentinel-based rate limiter delivers a smoother performance curve as rate-limit calculations happen off-thread, but a stricter time-out based cool-down for clients. For example, when a throttling action is triggered, they are required to cool-down for the period of the rate limit.
Disabling the sentinel based rate limiter will make rate-limit calculations happen on-thread and therefore offers a staggered cool-down and a smoother rate-limit experience for the client.
For example, you can slow your connection throughput to regain entry into your rate limit. This is more of a “throttle” than a “block”.
The standard rate limiter offers similar performance as the sentinel-based limiter. This is disabled by default.

**Field: `enable_non_transactional_rate_limiter`** (EnableNonTransactionalRateLimiter, `bool`)

An enhancement for the Redis and Sentinel rate limiters, that offers a significant improvement in performance by not using transactions on Redis rate-limit buckets.

**Field: `drl_notification_frequency`** (DRLNotificationFrequency, `int`)

How frequently a distributed rate limiter synchronises information between the Gateway nodes. Default: 2 seconds.

**Field: `drl_threshold`** (DRLThreshold, `float64`)

A distributed rate limiter is inaccurate on small rate limits, and it will fallback to a Redis or Sentinel rate limiter on an individual user basis, if its rate limiter lower then threshold.
A Rate limiter threshold calculated using the following formula: `rate_threshold = drl_threshold * number_of_gateways`.
So you have 2 Gateways, and your threshold is set to 5, if a user rate limit is larger than 10, it will use the distributed rate limiter algorithm.
Default: 5

**Field: `drl_enable_sentinel_rate_limiter`** (DRLEnableSentinelRateLimiter, `bool`)

Controls which algorthm to use as a fallback when your distributed rate limiter can't be used.

**Field: `enforce_org_data_age`** (EnforceOrgDataAge, `bool`)

Allows you to dynamically configure analytics expiration on a per organisation level

**Field: `enforce_org_data_detail_logging`** (EnforceOrgDataDetailLogging, `bool`)

Allows you to dynamically configure detailed logging on a per organisation level

**Field: `enforce_org_quotas`** (EnforceOrgQuotas, `bool`)

Allows you to dynamically configure organisation quotas on a per organisation level

**Field: `experimental_process_org_off_thread`** (ExperimentalProcessOrgOffThread, `bool`)



**Field: `monitor`** (Monitor, [MonitorConfig](#MonitorConfig))

The monitor section is useful if you wish to enforce a global trigger limit on organisation and user quotas.
This feature will trigger a webhook event to fire when specific triggers are reached.
Triggers can be global (set in the node), by organisation (set in the organisation session object) or by key (set in the key session object)

While Organisation-level and Key-level triggers can be tiered (e.g. trigger at 10%, trigger at 20%, trigger at 80%), in the node-level configuration only a global value can be set.
If a global value and specific trigger level are the same the trigger will only fire once:

```
"monitor": {
  "enable_trigger_monitors": true,
  "configuration": {
   "method": "POST",
   "target_path": "http://domain.com/notify/quota-trigger",
   "template_path": "templates/monitor_template.json",
   "header_map": {
     "some-secret": "89787855"
   },
   "event_timeout": 10
 },
 "global_trigger_limit": 80.0,
 "monitor_user_keys": false,
 "monitor_org_keys": true
},
```

**Field: `max_idle_connections`** (MaxIdleConns, `int`)

Maximum idle connections, per API, between Tyk and Upstream. By default not limited.

**Field: `max_idle_connections_per_host`** (MaxIdleConnsPerHost, `int`)

Maximum idle connections, per API, per upstream, between Tyk and Upstream. Default:100

**Field: `max_conn_time`** (MaxConnTime, `int64`)

Maximum connection time. If set it will force gateway reconnect to the upstream.

**Field: `close_connections`** (CloseConnections, `bool`)

If set, disable keepalive between User and Tyk

**Field: `enable_custom_domains`** (EnableCustomDomains, `bool`)

Allows you to use custom domains

**Field: `allow_master_keys`** (AllowMasterKeys, `bool`)

If AllowMasterKeys is set to true, session objects (key definitions) that do not have explicit access rights set
will be allowed by Tyk. This means that keys that are created have access to ALL APIs, which in many cases is
unwanted behaviour unless you are sure about what you are doing.

**Field: `service_discovery`** (ServiceDiscovery, [ServiceDiscoveryConf](#ServiceDiscoveryConf))



**Field: `proxy_ssl_insecure_skip_verify`** (ProxySSLInsecureSkipVerify, `bool`)

Globally ignore TLS verification between Tyk and your Upstream services

**Field: `proxy_enable_http2`** (ProxyEnableHttp2, `bool`)

Enable HTTP2 support between Tyk and your upstream service. Required for gRPC.

**Field: `proxy_ssl_min_version`** (ProxySSLMinVersion, `uint16`)

Minimum TLS version for connection between Tyk and your upstream service.

**Field: `proxy_ssl_max_version`** (ProxySSLMaxVersion, `uint16`)

Maximum TLS version for connection between Tyk and your upstream service.

**Field: `proxy_ssl_ciphers`** (ProxySSLCipherSuites, `[]string`)

Whitelist ciphers for connection between Tyk and your upstream service.

**Field: `proxy_default_timeout`** (ProxyDefaultTimeout, `float64`)

This can specify a default timeout in seconds for upstream API requests.

**Field: `proxy_ssl_disable_renegotiation`** (ProxySSLDisableRenegotiation, `bool`)

Disable TLS renegotiation.

**Field: `proxy_close_connections`** (ProxyCloseConnections, `bool`)

Disable keepalives between Tyk and your upstream service.
Set this value to `true` to force Tyk to close the connection with the server, otherwise the connections will remain open for as long as your OS keeps TCP connections open.
This can cause a file-handler limit to be exceeded. Setting to false can have performance benefits as the connection can be reused.

**Field: `uptime_tests`** (UptimeTests, [UptimeTestsConfig](#UptimeTestsConfig))

Tyk nodes can provide uptime awareness, uptime testing and analytics for your underlying APIs uptime and availability.
Tyk can also notify you when a service goes down.

**Field: `health_check`** (HealthCheck, [HealthCheckConfig](#HealthCheckConfig))

This section enables the configuration of the health-check API endpoint and the size of the sample data cache (in seconds).

**Field: `health_check_endpoint_name`** (HealthCheckEndpointName, `string`)

Enables you to rename the /hello endpoint

**Field: `oauth_refresh_token_expire`** (OauthRefreshExpire, `int64`)

Change the expiry time of a refresh token. By default 14 days (in seconds).

**Field: `oauth_token_expire`** (OauthTokenExpire, `int32`)

Change the expiry time of OAuth tokens (in seconds).

**Field: `oauth_token_expired_retain_period`** (OauthTokenExpiredRetainPeriod, `int32`)

Specifies how long expired tokens are stored in Redis. The value is in seconds and the default is 0. Using the default means expired tokens are never removed from Redis.

**Field: `oauth_redirect_uri_separator`** (OauthRedirectUriSeparator, `string`)

Character which should be used as a separator for OAuth redirect URI URLs. Default: ;.

**Field: `oauth_error_status_code`** (OauthErrorStatusCode, `int`)

Configures the OAuth error status code returned. If not set, it defaults to a 403 error.

**Field: `enable_key_logging`** (EnableKeyLogging, `bool`)

By default all key IDs in logs are hidden. Set to `true` if you want to see them for debugging reasons.

**Field: `ssl_force_common_name_check`** (SSLForceCommonNameCheck, `bool`)

Force the validation of the hostname against the common name, even if TLS verification is disabled.

**Field: `enable_analytics`** (EnableAnalytics, `bool`)

Tyk is capable of recording every hit to your API to a database with various filtering parameters. Set this value to `true` and fill in the sub-section below to enable logging.

Note:
  For performance reasons, Tyk will store traffic data to Redis initially and then purge the data from Redis to MongoDB or other data stores on a regular basis as determined by the purge_delay setting in your Tyk Pump configuration.

**Field: `analytics_config`** (AnalyticsConfig, [AnalyticsConfigConfig](#AnalyticsConfigConfig))

This section defines options on what analytics data to store.

**Field: `enable_separate_analytics_store`** (EnableSeperateAnalyticsStore, `bool`)

Enable separate analytics storage. Used together with `analytics_storage`.

**Field: `analytics_storage`** (AnalyticsStorage, [StorageOptionsConf](#StorageOptionsConf))



**Field: `liveness_check`** (LivenessCheck, [LivenessCheckConfig](#LivenessCheckConfig))



**Field: `dns_cache`** (DnsCache, [DnsCacheConfig](#DnsCacheConfig))

This section enables the global configuration of the expireable DNS records caching for your Gateway API endpoints.
By design caching affects only http(s), ws(s) protocols APIs and doesn’t affect any plugin/middleware DNS queries.

```
"dns_cache": {
  "enabled": true, //Turned off by default
  "ttl": 60, //Time in seconds before the record will be removed from cache
  "multiple_ips_handle_strategy": "random" //A strategy, which will be used when dns query will reply with more than 1 ip address per single host.
}
```

**Field: `disable_regexp_cache`** (DisableRegexpCache, `bool`)

If set to `true` this allows you to disable the regular expression cache. The default setting is `false`.

**Field: `regexp_cache_expire`** (RegexpCacheExpire, `int32`)

If you set `disable_regexp_cache` to `false`, you can use this setting to limit how long the regular expression cache is kept for in seconds.
The default is 60 seconds. This must be a positive value. If you set to 0 this uses the default value.

**Field: `local_session_cache`** (LocalSessionCache, [LocalSessionCacheConf](#LocalSessionCacheConf))

Tyk can cache some data locally, this can speed up lookup times on a single node and lower the number of connections and operations being done on Redis. It will however introduce a slight delay when updating or modifying keys as the cache must expire.
This does not affect rate limiting.

**Field: `enable_separate_cache_store`** (EnableSeperateCacheStore, `bool`)

Enable to use a separate Redis for cache storage

**Field: `cache_storage`** (CacheStorage, [StorageOptionsConf](#StorageOptionsConf))



**Field: `enable_bundle_downloader`** (EnableBundleDownloader, `bool`)

Enable downloading Plugin bundles
Example:
```
"enable_bundle_downloader": true,
"bundle_base_url": "http://my-bundle-server.com/bundles/",
"public_key_path": "/path/to/my/pubkey",
```

**Field: `bundle_base_url`** (BundleBaseURL, `string`)

Is a base URL that will be used to download the bundle. In this example we have `bundle-latest.zip` specified in the API settings, Tyk will fetch the following URL: http://my-bundle-server.com/bundles/bundle-latest.zip (see the next section for details).

**Field: `bundle_insecure_skip_verify`** (BundleInsecureSkipVerify, `bool`)

Disable TLS validation for bundle URLs

**Field: `enable_jsvm`** (EnableJSVM, `bool`)

Set to true if you are using JSVM custom middleware or virtual endpoints.

**Field: `jsvm_timeout`** (JSVMTimeout, `int`)

Set the execution timeout for JSVM plugins and virtal endpoints

**Field: `disable_virtual_path_blobs`** (DisableVirtualPathBlobs, `bool`)

Disable virtual endpoints and the code will not be loaded into the VM when the API definition initialises.
This is useful for systems where you want to avoid having third-party code run.

**Field: `tyk_js_path`** (TykJSPath, `string`)

Path to the JavaScript file which will be pre-loaded for any JSVM middleware or virtual endpoint. Useful for defining global shared functions.

**Field: `middleware_path`** (MiddlewarePath, `string`)

Path to the plugins dirrectory. By default is ``./middleware`.

**Field: `coprocess_options`** (CoProcessOptions, [CoProcessConfig](#CoProcessConfig))

Configuration options for Python and gRPC plugins.

**Field: `ignore_endpoint_case`** (IgnoreEndpointCase, `bool`)

Ignore the case of any endpoints for APIs managed by Tyk. Setting this to `true` will override any individual API and Ignore, Blacklist and Whitelist plugin endpoint settings.

**Field: `ignore_canonical_mime_header_key`** (IgnoreCanonicalMIMEHeaderKey, `bool`)

When enabled Tyk ignores the canonical format of the MIME header keys.

For example when a request header with a “my-header” key is injected using “global_headers”, the upstream would typically get it as “My-Header”. When this flag is enabled it will be sent as “my-header” instead.

Current support is limited to JavaScript plugins, global header injection, virtual endpoint and JQ transform header rewrites.
This functionality doesn’t affect headers that are sent by the HTTP client and the default formatting will apply in this case.

For technical details refer to the [CanonicalMIMEHeaderKey](https://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey) functionality in the Go documentation.

**Field: `log_level`** (LogLevel, `string`)

You can now set a logging level (log_level). The following levels can be set: debug, info, warn, error.
If not set or left empty, it will default to `info`.

**Field: `tracing`** (Tracer, [Tracer](#Tracer))

Section for configuring OpenTracing support

**Field: `newrelic`** (NewRelic, [NewRelicConfig](#NewRelicConfig))



**Field: `enable_http_profiler`** (HTTPProfile, `bool`)

Enable debugging of your Tyk Gateway by exposing profiling information through https://tyk.io/docs/troubleshooting/tyk-gateway/profiling/

**Field: `use_redis_log`** (UseRedisLog, `bool`)

Enables the real-time Gateway log view in the Dashboard.

**Field: `use_sentry`** (UseSentry, `bool`)

Enable Sentry logging

**Field: `sentry_code`** (SentryCode, `string`)

Sentry API code

**Field: `sentry_log_level`** (SentryLogLevel, `string`)

Log verbosity for Sentry logging

**Field: `use_syslog`** (UseSyslog, `bool`)

Enable Syslog log output

**Field: `syslog_transport`** (SyslogTransport, `string`)

Syslong transport to use. Values: tcp or udp.

**Field: `syslog_network_addr`** (SyslogNetworkAddr, `string`)

Graylog server address

**Field: `use_graylog`** (UseGraylog, `bool`)

Use Graylog log output

**Field: `graylog_network_addr`** (GraylogNetworkAddr, `string`)

Graylog server address

**Field: `use_logstash`** (UseLogstash, `bool`)

Use logstash log output

**Field: `logstash_transport`** (LogstashTransport, `string`)

Logstash network transport. Values: tcp or udp.

**Field: `logstash_network_addr`** (LogstashNetworkAddr, `string`)

Logstash server address

**Field: `track_404_logs`** (Track404Logs, `bool`)

Show 404 HTTP errors in your Gateway application logs

**Field: `statsd_connection_string`** (StatsdConnectionString, `string`)

Address of StatsD server. If set enable statsd monitoring.

**Field: `statsd_prefix`** (StatsdPrefix, `string`)

StatsD prefix

**Field: `event_handlers`** (EventHandlers, `apidef.EventHandlerMetaConfig`)

Event System

**Field: `event_trigers_defunct`** (EventTriggers, `any`)



> Deprecated: Config.GetEventTriggers instead.

**Field: `event_triggers_defunct`** (EventTriggersDefunct, `any`)



> Deprecated: Config.GetEventTriggers instead.

**Field: `hide_generator_header`** (HideGeneratorHeader, `bool`)

HideGeneratorHeader will mask the 'X-Generator' and 'X-Mascot-...' headers, if set to true.

**Field: `suppress_default_org_store`** (SupressDefaultOrgStore, `bool`)



**Field: `legacy_enable_allowance_countdown`** (LegacyEnableAllowanceCountdown, `bool`)



**Field: `force_global_session_lifetime`** (ForceGlobalSessionLifetime, `bool`)

Enable global API token expiration. Can be needed if all your APIs using JWT or oAuth 2.0 auth methods with dynamically generated keys.

**Field: `session_lifetime_respects_key_expiration`** (SessionLifetimeRespectsKeyExpiration, `bool`)

SessionLifetimeRespectsKeyExpiration respects the key expiration time when the session lifetime is less than the key expiration. That is, Redis waits the key expiration for physical removal.

**Field: `global_session_lifetime`** (GlobalSessionLifetime, `int64`)

global session lifetime, in seconds.

**Field: `kv`** (KV, `struct{}`)

This section enables the use of the KV capabilities to substitute configuration values.
See more details https://tyk.io/docs/tyk-configuration-reference/kv-store/

**Field: `secrets`** (Secrets, `map[string]string`)

Secrets are key-value pairs that can be accessed in the dashboard via "secrets://"

**Field: `override_messages`** (OverrideMessages, `map[string]TykError`)

Override the default error code and or message returned by middleware.
The following message IDs can be used to override the message and error codes:

AuthToken message IDs
* `auth.auth_field_missing`
* `auth.key_not_found`

OIDC message IDs
* `oauth.auth_field_missing`
* `oauth.auth_field_malformed`
* `oauth.key_not_found`
* `oauth.client_deleted`

Sample Override Message Setting
```
"override_messages": {
  "oauth.auth_field_missing" : {
   "code": 401,
   "message": "Token is not authorised"
 }
}
```

**Field: `cloud`** (Cloud, `bool`)

Cloud flag shows the Gateway runs in Tyk-cloud.

**Field: `jwt_ssl_insecure_skip_verify`** (JWTSSLInsecureSkipVerify, `bool`)

Skip TLS verification for JWT JWKs url validation

# SecurityConfig

**Field: `private_certificate_encoding_secret`** (PrivateCertificateEncodingSecret, `string`)

Set the AES256 secret which is used to encode certificate private keys when they uploaded via certificate storage

**Field: `control_api_use_mutual_tls`** (ControlAPIUseMutualTLS, `bool`)

Enable Gateway Control API to use Mutual TLS. Certificates can be set via `security.certificates.control_api` section

**Field: `pinned_public_keys`** (PinnedPublicKeys, `map[string]string`)

Specify public keys used for Certificate Pinning on global level.

**Field: `certificates`** (Certificates, [CertificatesConfig](#CertificatesConfig))



# HttpServerOptionsConfig

**Field: `read_timeout`** (ReadTimeout, `int`)

API Consumer -> Gateway network read timeout. Not setting this config, or setting this to 0, defaults to 120 seconds

**Field: `write_timeout`** (WriteTimeout, `int`)

API Consumer -> Gateway network write timeout. Not setting this config, or setting this to 0, defaults to 120 seconds

**Field: `use_ssl`** (UseSSL, `bool`)

Set to true to enable SSL connections

**Field: `enable_http2`** (EnableHttp2, `bool`)

Enable HTTP2 protocol handling

**Field: `enable_strict_routes`** (EnableStrictRoutes, `bool`)

EnableStrictRoutes changes the routing to avoid nearest-neighbour requests on overlapping routes

- if disabled, `/apple` will route to `/app`, the current default behavior,
- if enabled, `/app` only responds to `/app`, `/app/` and `/app/*` but not `/apple`

Regular expressions and parameterized routes will be left alone regardless of this setting.

**Field: `ssl_insecure_skip_verify`** (SSLInsecureSkipVerify, `bool`)

Disable TLS verification. Required if you are using self-signed certificates.

**Field: `enable_websockets`** (EnableWebSockets, `bool`)

Enabled WebSockets and server side events support

**Field: `certificates`** (Certificates, [CertsData](#CertsData))

Deprecated. SSL certificates used by Gateway server.

**Field: `ssl_certificates`** (SSLCertificates, `[]string`)

SSL certificates used by your Gateway server. A list of certificate IDs or path to files.

**Field: `server_name`** (ServerName, `string`)

Start your Gateway HTTP server on specific server name

**Field: `min_version`** (MinVersion, `uint16`)

Minimum TLS version. Possible values: https://tyk.io/docs/basic-config-and-security/security/tls-and-ssl/#values-for-tls-versions

**Field: `max_version`** (MaxVersion, `uint16`)

Maximum TLS version.

**Field: `skip_client_ca_announcement`** (SkipClientCAAnnouncement, `bool`)

When mTLS enabled, this option allows to skip client CA announcement in the TLS handshake.
This option is useful when you have a lot of ClientCAs and you want to reduce the handshake overhead, as some clients can hit TLS handshake limits.
This option does not give any hints to the client, on which certificate to pick (but this is very rare situation when it is required)

**Field: `flush_interval`** (FlushInterval, `int`)

Set this to the number of seconds that Tyk uses to flush content from the proxied upstream connection to the open downstream connection.
This option needed be set for streaming protocols like Server Side Events, or gRPC streaming.

**Field: `skip_url_cleaning`** (SkipURLCleaning, `bool`)

Allow the use of a double slash in a URL path. This can be useful if you need to pass raw URLs to your API endpoints.
For example: `http://myapi.com/get/http://example.com`.

**Field: `skip_target_path_escaping`** (SkipTargetPathEscaping, `bool`)

Disable automatic character escaping, allowing to path original URL data to the upstream.

**Field: `ssl_ciphers`** (Ciphers, `[]string`)

Custom SSL ciphers. See list of ciphers here https://tyk.io/docs/basic-config-and-security/security/tls-and-ssl/#specify-tls-cipher-suites-for-tyk-gateway--tyk-dashboard

**Field: `max_request_body_size`** (MaxRequestBodySize, `int64`)

MaxRequestBodySize configures the maximum request body size in bytes.

This option evaluates the `Content-Length` header and responds with
a HTTP 413 status code if larger than the defined size. If the header
is not provided, the request body is read up to the defined size.
If the request body is larger than the defined size, then we respond
with HTTP 413 status code.

See more information about setting request size limits here:
https://tyk.io/docs/basic-config-and-security/control-limit-traffic/request-size-limits/#maximum-request-sizes

# PoliciesConfig

**Field: `policy_source`** (PolicySource, `string`)

Set this value to `file` to look in the file system for a definition file. Set to `service` to use the Dashboard service.

**Field: `policy_connection_string`** (PolicyConnectionString, `string`)

This option is required if `policies.policy_source` is set to `service`.
Set this to the URL of your Tyk Dashboard installation. The URL needs to be formatted as: http://dashboard_host:port.

**Field: `policy_record_name`** (PolicyRecordName, `string`)

This option is required if `policies.policy_source` is set to `file`.
Specifies the path of your JSON file containing the available policies.

**Field: `allow_explicit_policy_id`** (AllowExplicitPolicyID, `bool`)

In a Pro installation, Tyk will load Policy IDs and use the internal object-ID as the ID of the policy.
This is not portable in cases where the data needs to be moved from installation to installation.

If you set this value to `true`, then the id parameter in a stored policy (or imported policy using the Dashboard API), will be used instead of the internal ID.

This option should only be used when moving an installation to a new database.

**Field: `policy_path`** (PolicyPath, `string`)

This option is used for storing a policies  if `policies.policy_source` is set to `file`.
it should be some existing file path on hard drive

# PortsWhiteList

Type defined as `map[string]PortWhiteList`, see [map[string]PortWhiteList](map[string]PortWhiteList) definition.

# DBAppConfOptionsConfig

**Field: `connection_string`** (ConnectionString, `string`)

Set the URL to your Dashboard instance (or a load balanced instance). The URL needs to be formatted as: `http://dashboard_host:port`

**Field: `node_is_segmented`** (NodeIsSegmented, `bool`)

Set to `true` to enable filtering (sharding) of APIs.

**Field: `tags`** (Tags, `[]string`)

The tags to use when filtering (sharding) Tyk Gateway nodes. Tags are processed as `OR` operations.
If you include a non-filter tag (e.g. an identifier such as `node-id-1`, this will become available to your Dashboard analytics).

# StorageOptionsConf

**Field: `type`** (Type, `string`)

This should be set to `redis` (lowercase)

**Field: `host`** (Host, `string`)

The Redis host, by default this is set to `localhost`, but for production this should be set to a cluster.

**Field: `port`** (Port, `int`)

The Redis instance port.

**Field: `hosts`** (Hosts, `map[string]string`)



> Deprecated: Addrs instead.

**Field: `addrs`** (Addrs, `[]string`)

If you have multi-node setup, you should use this field instead. For example: ["host1:port1", "host2:port2"].

**Field: `master_name`** (MasterName, `string`)

Redis sentinel master name

**Field: `sentinel_password`** (SentinelPassword, `string`)

Redis sentinel password

**Field: `username`** (Username, `string`)

Redis user name

**Field: `password`** (Password, `string`)

If your Redis instance has a password set for access, you can set it here.

**Field: `database`** (Database, `int`)

Redis database

**Field: `optimisation_max_idle`** (MaxIdle, `int`)

Set the number of maximum idle connections in the Redis connection pool, which defaults to 100. Set to a higher value if you are expecting more traffic.

**Field: `optimisation_max_active`** (MaxActive, `int`)

Set the number of maximum connections in the Redis connection pool, which defaults to 500. Set to a higher value if you are expecting more traffic.

**Field: `timeout`** (Timeout, `int`)

Set a custom timeout for Redis network operations. Default value 5 seconds.

**Field: `enable_cluster`** (EnableCluster, `bool`)

Enable Redis Cluster support

**Field: `use_ssl`** (UseSSL, `bool`)

Enable SSL/TLS connection between your Tyk Gateway & Redis.

**Field: `ssl_insecure_skip_verify`** (SSLInsecureSkipVerify, `bool`)

Disable TLS verification

# SlaveOptionsConfig

**Field: `use_rpc`** (UseRPC, `bool`)

Set to `true` to connect a worker Gateway using RPC.

**Field: `use_ssl`** (UseSSL, `bool`)

Set this option to `true` to use an SSL RPC connection.

**Field: `ssl_insecure_skip_verify`** (SSLInsecureSkipVerify, `bool`)

Set this option to `true` to allow the certificate validation (certificate chain and hostname) to be skipped.
This can be useful if you use a self-signed certificate.

**Field: `connection_string`** (ConnectionString, `string`)

Use this setting to add the URL for your MDCB or load balancer host.

**Field: `rpc_key`** (RPCKey, `string`)

Your organisation ID to connect to the MDCB installation.

**Field: `api_key`** (APIKey, `string`)

This the API key of a user used to authenticate and authorise the Gateway’s access through MDCB.
The user should be a standard Dashboard user with minimal privileges so as to reduce any risk if the user is compromised.
The suggested security settings are read for Real-time notifications and the remaining options set to deny.

**Field: `enable_rpc_cache`** (EnableRPCCache, `bool`)

Set this option to `true` to enable RPC caching for keys.

**Field: `bind_to_slugs`** (BindToSlugsInsteadOfListenPaths, `bool`)

For an Self-Managed installation this can be left at `false` (the default setting). For Legacy Cloud Gateways it must be set to ‘true’.

**Field: `disable_keyspace_sync`** (DisableKeySpaceSync, `bool`)

Set this option to `true` if you don’t want to monitor changes in the keys from a master Gateway.

**Field: `group_id`** (GroupID, `string`)

This is the `zone` that this instance inhabits, e.g. the cluster/data-centre the Gateway lives in.
The group ID must be the same across all the Gateways of a data-centre/cluster which are also sharing the same Redis instance.
This ID should also be unique per cluster (otherwise another Gateway cluster can pick up your keyspace events and your cluster will get zero updates).

**Field: `call_timeout`** (CallTimeout, `int`)

Call Timeout allows to specify a time in seconds for the maximum allowed duration of a RPC call.

**Field: `ping_timeout`** (PingTimeout, `int`)

The maximum time in seconds that a RPC ping can last.

**Field: `rpc_pool_size`** (RPCPoolSize, `int`)

The number of RPC connections in the pool. Basically it creates a set of connections that you can re-use as needed. Defaults to 5.

**Field: `key_space_sync_interval`** (KeySpaceSyncInterval, `float32`)

You can use this to set a period for which the Gateway will check if there are changes in keys that must be synchronized. If this value is not set then it will default to 10 seconds.

**Field: `rpc_cert_cache_expiration`** (RPCCertCacheExpiration, `float32`)

RPCCertCacheExpiration defines the expiration time of the rpc cache that stores the certificates, defined in seconds

**Field: `rpc_global_cache_expiration`** (RPCGlobalCacheExpiration, `float32`)

RPCKeysCacheExpiration defines the expiration time of the rpc cache that stores the keys, defined in seconds

**Field: `synchroniser_enabled`** (SynchroniserEnabled, `bool`)

SynchroniserEnabled enable this config if MDCB has enabled the synchoniser. If disabled then it will ignore signals to synchonise recources

# AuthOverrideConf

**Field: `force_auth_provider`** (ForceAuthProvider, `bool`)



**Field: `auth_provider`** (AuthProvider, `apidef.AuthProviderMeta`)



**Field: `force_session_provider`** (ForceSessionProvider, `bool`)



**Field: `session_provider`** (SessionProvider, `apidef.SessionProviderMeta`)



# MonitorConfig

**Field: `enable_trigger_monitors`** (EnableTriggerMonitors, `bool`)

Set this to `true` to have monitors enabled in your configuration for the node.

**Field: `configuration`** (Config, [WebHookHandlerConf](#WebHookHandlerConf))



**Field: `global_trigger_limit`** (GlobalTriggerLimit, `float64`)

The trigger limit, as a percentage of the quota that must be reached in order to trigger the event, any time the quota percentage is increased the event will trigger.

**Field: `monitor_user_keys`** (MonitorUserKeys, `bool`)

Apply the monitoring subsystem to user keys.

**Field: `monitor_org_keys`** (MonitorOrgKeys, `bool`)

Apply the monitoring subsystem to organisation keys.

# ServiceDiscoveryConf

**Field: `default_cache_timeout`** (DefaultCacheTimeout, `int`)

Service discovery cache timeout

# UptimeTestsConfig

**Field: `disable`** (Disable, `bool`)

To disable uptime tests on this node, set this value to `true`.

**Field: `poller_group`** (PollerGroup, `string`)

If you have multiple Gateway clusters connected to the same Redis instance, you need to set a unique poller group for each cluster.

**Field: `config`** (Config, [UptimeTestsConfigDetail](#UptimeTestsConfigDetail))



# HealthCheckConfig

**Field: `enable_health_checks`** (EnableHealthChecks, `bool`)

Setting this value to `true` will enable the health-check endpoint on /Tyk/health.

**Field: `health_check_value_timeouts`** (HealthCheckValueTimeout, `int64`)

This setting defaults to 60 seconds. This is the time window that Tyk uses to sample health-check data.
You can set a higher value for more accurate data (a larger sample period), or a lower value for less accurate data.
The reason this value is configurable is because sample data takes up space in your Redis DB to store the data to calculate samples. On high-availability systems this may not be desirable and smaller values may be preferred.

# AnalyticsConfigConfig

**Field: `type`** (Type, `string`)

Set empty for a Self-Managed installation or `rpc` for multi-cloud.

**Field: `ignored_ips`** (IgnoredIPs, `[]string`)

Adding IP addresses to this list will cause Tyk to ignore these IPs in the analytics data. These IP addresses will not produce an analytics log record.
This is useful for health checks and other samplers that might skew usage data.
The IP addresses must be provided as a JSON array, with the values being single IPs. CIDR values are not supported.

**Field: `enable_detailed_recording`** (EnableDetailedRecording, `bool`)

Set this value to `true` to have Tyk store the inbound request and outbound response data in HTTP Wire format as part of the Analytics data.
Please note, this will greatly increase your analytics DB size and can cause performance degradation on analytics processing by the Dashboard.
This setting can be overridden with an organisation flag, enabed at an API level, or on individual Key level.

**Field: `enable_geo_ip`** (EnableGeoIP, `bool`)

Tyk can store GeoIP information based on MaxMind DB’s to enable GeoIP tracking on inbound request analytics. Set this value to `true` and assign a DB using the `geo_ip_db_path` setting.

**Field: `geo_ip_db_path`** (GeoIPDBLocation, `string`)

Path to a MaxMind GeoIP database
The analytics GeoIP DB can be replaced on disk. It will cleanly auto-reload every hour.

**Field: `normalise_urls`** (NormaliseUrls, [NormalisedURLConfig](#NormalisedURLConfig))

This section describes methods that enable you to normalise inbound URLs in your analytics to have more meaningful per-path data.

**Field: `pool_size`** (PoolSize, `int`)

Number of workers used to process analytics. Defaults to number of CPU cores.

**Field: `records_buffer_size`** (RecordsBufferSize, `uint64`)

Number of records in analytics queue, per worker. Default: 1000.

**Field: `storage_expiration_time`** (StorageExpirationTime, `int`)

You can set a time (in seconds) to configure how long analytics are kept if they are not processed. The default is 60 seconds.
This is used to prevent the potential infinite growth of Redis analytics storage.

**Field: `enable_multiple_analytics_keys`** (EnableMultipleAnalyticsKeys, `bool`)

Set this to `true` to have Tyk automatically divide the analytics records in multiple analytics keys.
This is especially useful when `storage.enable_cluster` is set to `true` since it will distribute the analytic keys across all the cluster nodes.

**Field: `purge_interval`** (PurgeInterval, `float32`)

You can set the interval length on how often the tyk Gateway will purge analytics data. This value is in seconds and defaults to 10 seconds.

**Field: `serializer_type`** (SerializerType, `string`)

Determines the serialization engine for analytics. Available options: msgpack, and protobuf. By default, msgpack.

# StorageOptionsConf

**Field: `type`** (Type, `string`)

This should be set to `redis` (lowercase)

**Field: `host`** (Host, `string`)

The Redis host, by default this is set to `localhost`, but for production this should be set to a cluster.

**Field: `port`** (Port, `int`)

The Redis instance port.

**Field: `hosts`** (Hosts, `map[string]string`)



> Deprecated: Addrs instead.

**Field: `addrs`** (Addrs, `[]string`)

If you have multi-node setup, you should use this field instead. For example: ["host1:port1", "host2:port2"].

**Field: `master_name`** (MasterName, `string`)

Redis sentinel master name

**Field: `sentinel_password`** (SentinelPassword, `string`)

Redis sentinel password

**Field: `username`** (Username, `string`)

Redis user name

**Field: `password`** (Password, `string`)

If your Redis instance has a password set for access, you can set it here.

**Field: `database`** (Database, `int`)

Redis database

**Field: `optimisation_max_idle`** (MaxIdle, `int`)

Set the number of maximum idle connections in the Redis connection pool, which defaults to 100. Set to a higher value if you are expecting more traffic.

**Field: `optimisation_max_active`** (MaxActive, `int`)

Set the number of maximum connections in the Redis connection pool, which defaults to 500. Set to a higher value if you are expecting more traffic.

**Field: `timeout`** (Timeout, `int`)

Set a custom timeout for Redis network operations. Default value 5 seconds.

**Field: `enable_cluster`** (EnableCluster, `bool`)

Enable Redis Cluster support

**Field: `use_ssl`** (UseSSL, `bool`)

Enable SSL/TLS connection between your Tyk Gateway & Redis.

**Field: `ssl_insecure_skip_verify`** (SSLInsecureSkipVerify, `bool`)

Disable TLS verification

# LivenessCheckConfig

**Field: `check_duration`** (CheckDuration, `time.Duration`)

Frequencies of performing interval healthchecks for Redis, Dashboard, and RPC layer. Default: 10 seconds.

# DnsCacheConfig

**Field: `enabled`** (Enabled, `bool`)

Setting this value to `true` will enable caching of DNS queries responses used for API endpoint’s host names. By default caching is disabled.

**Field: `ttl`** (TTL, `int64`)

This setting allows you to specify a duration in seconds before the record will be removed from cache after being added to it on the first DNS query resolution of API endpoints.
Setting `ttl` to `-1` prevents record from being expired and removed from cache on next check interval.

**Field: `multiple_ips_handle_strategy`** (MultipleIPsHandleStrategy, [IPsHandleStrategy](#IPsHandleStrategy))

A strategy which will be used when a DNS query will reply with more than 1 IP Address per single host.
As a DNS query response IP Addresses can have a changing order depending on DNS server balancing strategy (eg: round robin, geographically dependent origin-ip ordering, etc) this option allows you to not to limit the connection to the first host in a cached response list or prevent response caching.

* `pick_first` will instruct your Tyk Gateway to connect to the first IP in a returned IP list and cache the response.
* `random` will instruct your Tyk Gateway to connect to a random IP in a returned IP list and cache the response.
* `no_cache` will instruct your Tyk Gateway to connect to the first IP in a returned IP list and fetch each addresses list without caching on each API endpoint DNS query.

# LocalSessionCacheConf

**Field: `disable_cached_session_state`** (DisableCacheSessionState, `bool`)

By default sessions are set to cache. Set this to `true` to stop Tyk from caching keys locally on the node.

**Field: `cached_session_timeout`** (CachedSessionTimeout, `int`)



**Field: `cached_session_eviction`** (CacheSessionEviction, `int`)



# StorageOptionsConf

**Field: `type`** (Type, `string`)

This should be set to `redis` (lowercase)

**Field: `host`** (Host, `string`)

The Redis host, by default this is set to `localhost`, but for production this should be set to a cluster.

**Field: `port`** (Port, `int`)

The Redis instance port.

**Field: `hosts`** (Hosts, `map[string]string`)



> Deprecated: Addrs instead.

**Field: `addrs`** (Addrs, `[]string`)

If you have multi-node setup, you should use this field instead. For example: ["host1:port1", "host2:port2"].

**Field: `master_name`** (MasterName, `string`)

Redis sentinel master name

**Field: `sentinel_password`** (SentinelPassword, `string`)

Redis sentinel password

**Field: `username`** (Username, `string`)

Redis user name

**Field: `password`** (Password, `string`)

If your Redis instance has a password set for access, you can set it here.

**Field: `database`** (Database, `int`)

Redis database

**Field: `optimisation_max_idle`** (MaxIdle, `int`)

Set the number of maximum idle connections in the Redis connection pool, which defaults to 100. Set to a higher value if you are expecting more traffic.

**Field: `optimisation_max_active`** (MaxActive, `int`)

Set the number of maximum connections in the Redis connection pool, which defaults to 500. Set to a higher value if you are expecting more traffic.

**Field: `timeout`** (Timeout, `int`)

Set a custom timeout for Redis network operations. Default value 5 seconds.

**Field: `enable_cluster`** (EnableCluster, `bool`)

Enable Redis Cluster support

**Field: `use_ssl`** (UseSSL, `bool`)

Enable SSL/TLS connection between your Tyk Gateway & Redis.

**Field: `ssl_insecure_skip_verify`** (SSLInsecureSkipVerify, `bool`)

Disable TLS verification

# CoProcessConfig

**Field: `enable_coprocess`** (EnableCoProcess, `bool`)

Enable gRPC and Python plugins

**Field: `coprocess_grpc_server`** (CoProcessGRPCServer, `string`)

Address of gRPC user

**Field: `grpc_recv_max_size`** (GRPCRecvMaxSize, `int`)

Maximum message which can be received from a gRPC server

**Field: `grpc_send_max_size`** (GRPCSendMaxSize, `int`)

Maximum message which can be sent to gRPC server

**Field: `grpc_authority`** (GRPCAuthority, `string`)

Authority used in GRPC connection

**Field: `python_path_prefix`** (PythonPathPrefix, `string`)

Sets the path to built-in Tyk modules. This will be part of the Python module lookup path. The value used here is the default one for most installations.

**Field: `python_version`** (PythonVersion, `string`)

If you have multiple Python versions installed you can specify your version.

# Tracer

**Field: `name`** (Name, `string`)

The name of the tracer to initialize. For instance appdash, to use appdash tracer

**Field: `enabled`** (Enabled, `bool`)

Enable tracing

**Field: `options`** (Options, `map[string]interface{}`)

Tracing configuration. Refer to the Tracing Docs for the full list of options.

# NewRelicConfig

**Field: `app_name`** (AppName, `string`)

New Relic Application name

**Field: `license_key`** (LicenseKey, `string`)

New Relic License key

**Field: `enable_distributed_tracing`** (EnableDistributedTracing, `bool`)

Enable distributed tracing

# CertificatesConfig

**Field: `apis`** (API, `[]string`)



**Field: `upstream`** (Upstream, `map[string]string`)

Specify upstream mutual TLS certificates at a global level in the following format: `{ "<host>": "<cert>" }``

**Field: `control_api`** (ControlAPI, `[]string`)

Certificates used for Control API Mutual TLS

**Field: `dashboard_api`** (Dashboard, `[]string`)

Used for communicating with the Dashboard if it is configured to use Mutual TLS

**Field: `mdcb_api`** (MDCB, `[]string`)

Certificates used for MDCB Mutual TLS

# CertsData

Type defined as `[]CertData`, see [CertData](CertData) definition.

# WebHookHandlerConf

**Field: `method`** (Method, `string`)

The method to use for the webhook.

**Field: `target_path`** (TargetPath, `string`)

The target path on which to send the request.

**Field: `template_path`** (TemplatePath, `string`)

The template to load in order to format the request.

**Field: `header_map`** (HeaderList, `map[string]string`)

Headers to set when firing the webhook.

**Field: `event_timeout`** (EventTimeout, `int64`)

The cool-down for the event so it does not trigger again (in seconds).

# UptimeTestsConfigDetail

**Field: `failure_trigger_sample_size`** (FailureTriggerSampleSize, `int`)

The sample size to trigger a `HostUp` or `HostDown` event. For example, a setting of 3 will require at least three failures to occur before the uptime test is triggered.

**Field: `time_wait`** (TimeWait, `int`)

The value in seconds between tests runs. All tests will run simultaneously. This value will set the time between those tests. So a value of 60 will run all uptime tests every 60 seconds.

**Field: `checker_pool_size`** (CheckerPoolSize, `int`)

The goroutine pool size to keep idle for uptime tests. If you have many uptime tests running at a high time period, then increase this value.

**Field: `enable_uptime_analytics`** (EnableUptimeAnalytics, `bool`)

Set this value to `true` to have the node capture and record analytics data regarding the uptime tests.

# NormalisedURLConfig

**Field: `enabled`** (Enabled, `bool`)

Set this to `true` to enable normalisation.

**Field: `normalise_uuids`** (NormaliseUUIDs, `bool`)

Each UUID will be replaced with a placeholder {uuid}

**Field: `normalise_numbers`** (NormaliseNumbers, `bool`)

Set this to true to have Tyk automatically match for numeric IDs, it will match with a preceding slash so as not to capture actual numbers:

**Field: `custom_patterns`** (Custom, `[]string`)

This is a list of custom patterns you can add. These must be valid regex strings. Tyk will replace these values with a {var} placeholder.

# IPsHandleStrategy

Type defined as `string`, see [string](string) definition.

# CertData

**Field: `domain_name`** (Name, `string`)

Domain name

**Field: `cert_file`** (CertFile, `string`)

Path to certificate file

**Field: `key_file`** (KeyFile, `string`)

Path to private key file

# NormaliseURLPatterns

**Field: `UUIDs`** (UUIDs, `regexp.Regexp`)



**Field: `IDs`** (IDs, `regexp.Regexp`)



**Field: `Custom`** (Custom, `[]*regexp.Regexp`)



# TykError

**Field: `message`** (Message, `string`)



**Field: `code`** (Code, `int`)



# Reporter

**Field: `url`** (URL, `string`)

URL connection url to the zipkin server

**Field: `batch_size`** (BatchSize, `int`)



**Field: `max_backlog`** (MaxBacklog, `int`)



# Sampler

**Field: `name`** (Name, `string`)

Name is the name of the sampler to use. Options are

	"boundary"
is appropriate for high-traffic instrumentation who
provision random trace ids, and make the sampling decision only once.
It defends against nodes in the cluster selecting exactly the same ids.

	"count"
is appropriate for low-traffic instrumentation or
those who do not provision random trace ids. It is not appropriate for
collectors as the sampling decision isn't idempotent (consistent based
on trace id).

"mod"
provides a generic type Sampler

**Field: `rate`** (Rate, `float64`)

Rate is used by both "boundary" and "count" samplers

**Field: `salt`** (Salt, `int64`)

Salt is used by "boundary" sampler

**Field: `mod`** (Mod, `uint64`)

Mod is only used when sampler is mod

# ConsulConfig

ConsulConfig is used to configure the creation of a client
This is a stripped down version of the Config struct in consul's API client

**Field: `address`** (Address, `string`)

Address is the address of the Consul server

**Field: `scheme`** (Scheme, `string`)

Scheme is the URI scheme for the Consul server

**Field: `datacenter`** (Datacenter, `string`)

The datacenter to use. If not provided, the default agent datacenter is used.

**Field: `http_auth`** (HttpAuth, `struct{}`)

HttpAuth is the auth info to use for http access.

**Field: `wait_time`** (WaitTime, `time.Duration`)

WaitTime limits how long a Watch will block. If not provided,
the agent default values will be used.

**Field: `token`** (Token, `string`)

Token is used to provide a per-request ACL token
which overrides the agent's default token.

**Field: `tls_config`** (TLSConfig, `struct{}`)

TLS configuration

# EventMessage

EventMessage is a standard form to send event data to handlers

**Field: `Type`** (Type, `apidef.TykEvent`)



**Field: `Meta`** (Meta, ``)



**Field: `TimeStamp`** (TimeStamp, `string`)



# PortRange

PortRange defines a range of ports inclusively.

**Field: `from`** (From, `int`)



**Field: `to`** (To, `int`)



# PortWhiteList

PortWhiteList defines ports that will be allowed by the Gateway.

**Field: `ranges`** (Ranges, [[]PortRange](#PortRange))



**Field: `ports`** (Ports, `[]int`)



# ServicePort

ServicePort defines a protocol and port on which a service can bind to.

**Field: `protocol`** (Protocol, `string`)



**Field: `port`** (Port, `int`)



# TykEventHandler

TykEventHandler defines an event handler, e.g. LogMessageEventHandler will handle an event by logging it to stdout.

Type defined as ``, see []() definition.

# VaultConfig

VaultConfig is used to configure the creation of a client
This is a stripped down version of the config structure in vault's API client

**Field: `address`** (Address, `string`)

Address is the address of the Vault server. This should be a complete
URL such as "http://vault.example.com".

**Field: `agent_address`** (AgentAddress, `string`)

AgentAddress is the address of the local Vault agent. This should be a
complete URL such as "http://vault.example.com".

**Field: `max_retries`** (MaxRetries, `int`)

MaxRetries controls the maximum number of times to retry when a vault
serer occurs

**Field: `timeout`** (Timeout, `time.Duration`)



**Field: `token`** (Token, `string`)

Token is the vault root token

**Field: `kv_version`** (KVVersion, `int`)

KVVersion is the version number of Vault. Usually defaults to 2

# ZipkinConfig

ZipkinConfig configuration options used to initialize openzipkin opentracing
client.

**Field: `reporter`** (Reporter, [Reporter](#Reporter))



**Field: `sampler`** (Sampler, [Sampler](#Sampler))



