# Config

Config is the configuration object used by Tyk to set up various parameters.

**HostName** (JSON: `hostname`)

Force your Gateway to work only on a specific domain name. Can be overridden by API custom domain.

**ListenAddress** (JSON: `listen_address`)

If your machine has multiple network devices or IPs you can force the Gateway to use the IP address you want.

**ListenPort** (JSON: `listen_port`)

Setting this value will change the port that Tyk listens on. Default: 8080.

**ControlAPIHostname** (JSON: `control_api_hostname`)

Custom hostname for the Control API

**ControlAPIPort** (JSON: `control_api_port`)

Set to run your Gateway Control API on a separate port, and protect it behind a firewall if needed. Please make sure you follow this guide when setting the control port https://tyk.io/docs/planning-for-production/#change-your-control-port.

**Secret** (JSON: `secret`)

This should be changed as soon as Tyk is installed on your system.
This value is used in every interaction with the Tyk Gateway API. It should be passed along as the X-Tyk-Authorization header in any requests made.
Tyk assumes that you are sensible enough not to expose the management endpoints publicly and to keep this configuration value to yourself.

**NodeSecret** (JSON: `node_secret`)

The shared secret between the Gateway and the Dashboard to ensure that API Definition downloads, heartbeat and Policy loads are from a valid source.

**PIDFileLocation** (JSON: `pid_file_location`)

Linux PID file location. Do not change unless you know what you are doing. Default: /var/run/tyk/tyk-gateway.pid

**AllowInsecureConfigs** (JSON: `allow_insecure_configs`)

Can be set to disable Dashboard message signature verification. When set to `true`, `public_key_path` can be ignored.

**PublicKeyPath** (JSON: `public_key_path`)

While communicating with the Dashboard. By default, all messages are signed by a private/public key pair. Set path to public key.

**AllowRemoteConfig** (JSON: `allow_remote_config`)

Allow your Dashboard to remotely set Gateway configuration via the Nodes screen.

**Security** (JSON: `security`)

Global Certificate configuration

**HttpServerOptions** (JSON: `http_server_options`)

Gateway HTTP server configuration

**VersionHeader** (JSON: `version_header`)

Expose version header with a given name. Works only for versioned APIs.

**SuppressRedisSignalReload** (JSON: `suppress_redis_signal_reload`)

Disable dynamic API and Policy reloads, e.g. it will load new changes only on procecss start.

**HashKeys** (JSON: `hash_keys`)

Enable Key hashing

**HashKeyFunction** (JSON: `hash_key_function`)

Specify the Key hashing algorithm. Possible values: murmur64, murmur128, sha256.

**BasicAuthHashKeyFunction** (JSON: `basic_auth_hash_key_function`)

Specify the Key hashing algorithm for "basic auth". Possible values: murmur64, murmur128, sha256, bcrypt.
Will default to "bcrypt" if not set.

**HashKeyFunctionFallback** (JSON: `hash_key_function_fallback`)

Specify your previous key hashing algorithm if you migrated from one algorithm to another.

**EnableHashedKeysListing** (JSON: `enable_hashed_keys_listing`)

Allows the listing of hashed API keys

**MinTokenLength** (JSON: `min_token_length`)

Minimum API token length

**TemplatePath** (JSON: `template_path`)

Path to error and webhook templates. Defaults to the current binary path.

**Policies** (JSON: `policies`)

The policies section allows you to define where Tyk can find its policy templates. Policy templates are similar to key definitions in that they allow you to set quotas, access rights and rate limits for keys.
Policies are loaded when Tyk starts and if changed require a hot-reload so they are loaded into memory.
A policy can be defined in a file (Open Source installations) or from the same database as the Dashboard.

**PortWhiteList** (JSON: `ports_whitelist`)

Defines the ports that will be available for the API services to bind to in the following format: `"{“":“”}"`. Remember to escape JSON strings.
This is a map of protocol to PortWhiteList. This allows per protocol
configurations.

**DisablePortWhiteList** (JSON: `disable_ports_whitelist`)

Disable port whilisting, essentially allowing you to use any port for your API.

**AppPath** (JSON: `app_path`)

If Tyk is being used in its standard configuration (Open Source installations), then API definitions are stored in the apps folder (by default in /opt/tyk-gateway/apps).
This location is scanned for .json files and re-scanned at startup or reload.
See the API section of the Tyk Gateway API for more details.

**UseDBAppConfigs** (JSON: `use_db_app_configs`)

If you are a Tyk Pro user, this option will enable polling the Dashboard service for API definitions.
On startup Tyk will attempt to connect and download any relevant application configurations from from your Dashboard instance.
The files are exactly the same as the JSON files on disk with the exception of a BSON ID supplied by the Dashboard service.

**DBAppConfOptions** (JSON: `db_app_conf_options`)

This section defines API loading and shard options. Enable these settings to selectively load API definitions on a node from your Dashboard service.

**Storage** (JSON: `storage`)

This section defines your Redis configuration.

**DisableDashboardZeroConf** (JSON: `disable_dashboard_zeroconf`)

Disable the capability of the Gateway to `autodiscover` the Dashboard through heartbeat messages via Redis.
The goal of zeroconf is auto-discovery, so you do not have to specify the Tyk Dashboard address in your Gateway`tyk.conf` file.
In some specific cases, for example, when the Dashboard is bound to a public domain, not accessible inside an internal network, or similar, `disable_dashboard_zeroconf` can be set to `true`, in favour of directly specifying a Tyk Dashboard address.

**SlaveOptions** (JSON: `slave_options`)

The `slave_options` allow you to configure the RPC slave connection required for MDCB installations.
These settings must be configured for every RPC slave/worker node.

**ManagementNode** (JSON: `management_node`)

If set to `true`, distributed rate limiter will be disabled for this node, and it will be excluded from any rate limit calculation.

Note:
  If you set `db_app_conf_options.node_is_segmented` to `true` for multiple Gateway nodes, you should ensure that `management_node` is set to `false`.
  This is to ensure visibility for the management node across all APIs.

**AuthOverride** (JSON: `auth_override`)

This is used as part of the RPC / Hybrid back-end configuration in a Tyk Enterprise installation and isn’t used anywhere else.

**EnableRedisRollingLimiter** (JSON: `enable_redis_rolling_limiter`)

Redis based rate limiter with fixed window. Provides 100% rate limiting accuracy, but require two additional Redis roundtrip for each request.

**EnableSentinelRateLimiter** (JSON: `enable_sentinel_rate_limiter`)

To enable, set to `true`. The sentinel-based rate limiter delivers a smoother performance curve as rate-limit calculations happen off-thread, but a stricter time-out based cool-down for clients. For example, when a throttling action is triggered, they are required to cool-down for the period of the rate limit.
Disabling the sentinel based rate limiter will make rate-limit calculations happen on-thread and therefore offers a staggered cool-down and a smoother rate-limit experience for the client.
For example, you can slow your connection throughput to regain entry into your rate limit. This is more of a “throttle” than a “block”.
The standard rate limiter offers similar performance as the sentinel-based limiter. This is disabled by default.

**EnableNonTransactionalRateLimiter** (JSON: `enable_non_transactional_rate_limiter`)

An enhancement for the Redis and Sentinel rate limiters, that offers a significant improvement in performance by not using transactions on Redis rate-limit buckets.

**DRLNotificationFrequency** (JSON: `drl_notification_frequency`)

How frequently a distributed rate limiter synchronises information between the Gateway nodes. Default: 2 seconds.

**DRLThreshold** (JSON: `drl_threshold`)

A distributed rate limiter is inaccurate on small rate limits, and it will fallback to a Redis or Sentinel rate limiter on an individual user basis, if its rate limiter lower then threshold.
A Rate limiter threshold calculated using the following formula: `rate_threshold = drl_threshold * number_of_gateways`.
So you have 2 Gateways, and your threshold is set to 5, if a user rate limit is larger than 10, it will use the distributed rate limiter algorithm.
Default: 5

**DRLEnableSentinelRateLimiter** (JSON: `drl_enable_sentinel_rate_limiter`)

Controls which algorthm to use as a fallback when your distributed rate limiter can't be used.

**EnforceOrgDataAge** (JSON: `enforce_org_data_age`)

Allows you to dynamically configure analytics expiration on a per organisation level

**EnforceOrgDataDetailLogging** (JSON: `enforce_org_data_detail_logging`)

Allows you to dynamically configure detailed logging on a per organisation level

**EnforceOrgQuotas** (JSON: `enforce_org_quotas`)

Allows you to dynamically configure organisation quotas on a per organisation level

**ExperimentalProcessOrgOffThread** (JSON: `experimental_process_org_off_thread`)



**Monitor** (JSON: `monitor`)

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

**MaxIdleConns** (JSON: `max_idle_connections`)

Maximum idle connections, per API, between Tyk and Upstream. By default not limited.

**MaxIdleConnsPerHost** (JSON: `max_idle_connections_per_host`)

Maximum idle connections, per API, per upstream, between Tyk and Upstream. Default:100

**MaxConnTime** (JSON: `max_conn_time`)

Maximum connection time. If set it will force gateway reconnect to the upstream.

**CloseConnections** (JSON: `close_connections`)

If set, disable keepalive between User and Tyk

**EnableCustomDomains** (JSON: `enable_custom_domains`)

Allows you to use custom domains

**AllowMasterKeys** (JSON: `allow_master_keys`)

If AllowMasterKeys is set to true, session objects (key definitions) that do not have explicit access rights set
will be allowed by Tyk. This means that keys that are created have access to ALL APIs, which in many cases is
unwanted behaviour unless you are sure about what you are doing.

**ServiceDiscovery** (JSON: `service_discovery`)



**ProxySSLInsecureSkipVerify** (JSON: `proxy_ssl_insecure_skip_verify`)

Globally ignore TLS verification between Tyk and your Upstream services

**ProxyEnableHttp2** (JSON: `proxy_enable_http2`)

Enable HTTP2 support between Tyk and your upstream service. Required for gRPC.

**ProxySSLMinVersion** (JSON: `proxy_ssl_min_version`)

Minimum TLS version for connection between Tyk and your upstream service.

**ProxySSLMaxVersion** (JSON: `proxy_ssl_max_version`)

Maximum TLS version for connection between Tyk and your upstream service.

**ProxySSLCipherSuites** (JSON: `proxy_ssl_ciphers`)

Whitelist ciphers for connection between Tyk and your upstream service.

**ProxyDefaultTimeout** (JSON: `proxy_default_timeout`)

This can specify a default timeout in seconds for upstream API requests.

**ProxySSLDisableRenegotiation** (JSON: `proxy_ssl_disable_renegotiation`)

Disable TLS renegotiation.

**ProxyCloseConnections** (JSON: `proxy_close_connections`)

Disable keepalives between Tyk and your upstream service.
Set this value to `true` to force Tyk to close the connection with the server, otherwise the connections will remain open for as long as your OS keeps TCP connections open.
This can cause a file-handler limit to be exceeded. Setting to false can have performance benefits as the connection can be reused.

**UptimeTests** (JSON: `uptime_tests`)

Tyk nodes can provide uptime awareness, uptime testing and analytics for your underlying APIs uptime and availability.
Tyk can also notify you when a service goes down.

**HealthCheck** (JSON: `health_check`)

This section enables the configuration of the health-check API endpoint and the size of the sample data cache (in seconds).

**HealthCheckEndpointName** (JSON: `health_check_endpoint_name`)

Enables you to rename the /hello endpoint

**OauthRefreshExpire** (JSON: `oauth_refresh_token_expire`)

Change the expiry time of a refresh token. By default 14 days (in seconds).

**OauthTokenExpire** (JSON: `oauth_token_expire`)

Change the expiry time of OAuth tokens (in seconds).

**OauthTokenExpiredRetainPeriod** (JSON: `oauth_token_expired_retain_period`)

Specifies how long expired tokens are stored in Redis. The value is in seconds and the default is 0. Using the default means expired tokens are never removed from Redis.

**OauthRedirectUriSeparator** (JSON: `oauth_redirect_uri_separator`)

Character which should be used as a separator for OAuth redirect URI URLs. Default: ;.

**OauthErrorStatusCode** (JSON: `oauth_error_status_code`)

Configures the OAuth error status code returned. If not set, it defaults to a 403 error.

**EnableKeyLogging** (JSON: `enable_key_logging`)

By default all key IDs in logs are hidden. Set to `true` if you want to see them for debugging reasons.

**SSLForceCommonNameCheck** (JSON: `ssl_force_common_name_check`)

Force the validation of the hostname against the common name, even if TLS verification is disabled.

**EnableAnalytics** (JSON: `enable_analytics`)

Tyk is capable of recording every hit to your API to a database with various filtering parameters. Set this value to `true` and fill in the sub-section below to enable logging.

Note:
  For performance reasons, Tyk will store traffic data to Redis initially and then purge the data from Redis to MongoDB or other data stores on a regular basis as determined by the purge_delay setting in your Tyk Pump configuration.

**AnalyticsConfig** (JSON: `analytics_config`)

This section defines options on what analytics data to store.

**EnableSeperateAnalyticsStore** (JSON: `enable_separate_analytics_store`)

Enable separate analytics storage. Used together with `analytics_storage`.

**AnalyticsStorage** (JSON: `analytics_storage`)



**LivenessCheck** (JSON: `liveness_check`)



**DnsCache** (JSON: `dns_cache`)

This section enables the global configuration of the expireable DNS records caching for your Gateway API endpoints.
By design caching affects only http(s), ws(s) protocols APIs and doesn’t affect any plugin/middleware DNS queries.

```
"dns_cache": {
  "enabled": true, //Turned off by default
  "ttl": 60, //Time in seconds before the record will be removed from cache
  "multiple_ips_handle_strategy": "random" //A strategy, which will be used when dns query will reply with more than 1 ip address per single host.
}
```

**DisableRegexpCache** (JSON: `disable_regexp_cache`)

If set to `true` this allows you to disable the regular expression cache. The default setting is `false`.

**RegexpCacheExpire** (JSON: `regexp_cache_expire`)

If you set `disable_regexp_cache` to `false`, you can use this setting to limit how long the regular expression cache is kept for in seconds.
The default is 60 seconds. This must be a positive value. If you set to 0 this uses the default value.

**LocalSessionCache** (JSON: `local_session_cache`)

Tyk can cache some data locally, this can speed up lookup times on a single node and lower the number of connections and operations being done on Redis. It will however introduce a slight delay when updating or modifying keys as the cache must expire.
This does not affect rate limiting.

**EnableSeperateCacheStore** (JSON: `enable_separate_cache_store`)

Enable to use a separate Redis for cache storage

**CacheStorage** (JSON: `cache_storage`)



**EnableBundleDownloader** (JSON: `enable_bundle_downloader`)

Enable downloading Plugin bundles
Example:
```
"enable_bundle_downloader": true,
"bundle_base_url": "http://my-bundle-server.com/bundles/",
"public_key_path": "/path/to/my/pubkey",
```

**BundleBaseURL** (JSON: `bundle_base_url`)

Is a base URL that will be used to download the bundle. In this example we have `bundle-latest.zip` specified in the API settings, Tyk will fetch the following URL: http://my-bundle-server.com/bundles/bundle-latest.zip (see the next section for details).

**BundleInsecureSkipVerify** (JSON: `bundle_insecure_skip_verify`)

Disable TLS validation for bundle URLs

**EnableJSVM** (JSON: `enable_jsvm`)

Set to true if you are using JSVM custom middleware or virtual endpoints.

**JSVMTimeout** (JSON: `jsvm_timeout`)

Set the execution timeout for JSVM plugins and virtal endpoints

**DisableVirtualPathBlobs** (JSON: `disable_virtual_path_blobs`)

Disable virtual endpoints and the code will not be loaded into the VM when the API definition initialises.
This is useful for systems where you want to avoid having third-party code run.

**TykJSPath** (JSON: `tyk_js_path`)

Path to the JavaScript file which will be pre-loaded for any JSVM middleware or virtual endpoint. Useful for defining global shared functions.

**MiddlewarePath** (JSON: `middleware_path`)

Path to the plugins dirrectory. By default is ``./middleware`.

**CoProcessOptions** (JSON: `coprocess_options`)

Configuration options for Python and gRPC plugins.

**IgnoreEndpointCase** (JSON: `ignore_endpoint_case`)

Ignore the case of any endpoints for APIs managed by Tyk. Setting this to `true` will override any individual API and Ignore, Blacklist and Whitelist plugin endpoint settings.

**IgnoreCanonicalMIMEHeaderKey** (JSON: `ignore_canonical_mime_header_key`)

When enabled Tyk ignores the canonical format of the MIME header keys.

For example when a request header with a “my-header” key is injected using “global_headers”, the upstream would typically get it as “My-Header”. When this flag is enabled it will be sent as “my-header” instead.

Current support is limited to JavaScript plugins, global header injection, virtual endpoint and JQ transform header rewrites.
This functionality doesn’t affect headers that are sent by the HTTP client and the default formatting will apply in this case.

For technical details refer to the [CanonicalMIMEHeaderKey](https://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey) functionality in the Go documentation.

**LogLevel** (JSON: `log_level`)

You can now set a logging level (log_level). The following levels can be set: debug, info, warn, error.
If not set or left empty, it will default to `info`.

**Tracer** (JSON: `tracing`)

Section for configuring OpenTracing support

**NewRelic** (JSON: `newrelic`)



**HTTPProfile** (JSON: `enable_http_profiler`)

Enable debugging of your Tyk Gateway by exposing profiling information through https://tyk.io/docs/troubleshooting/tyk-gateway/profiling/

**UseRedisLog** (JSON: `use_redis_log`)

Enables the real-time Gateway log view in the Dashboard.

**UseSentry** (JSON: `use_sentry`)

Enable Sentry logging

**SentryCode** (JSON: `sentry_code`)

Sentry API code

**SentryLogLevel** (JSON: `sentry_log_level`)

Log verbosity for Sentry logging

**UseSyslog** (JSON: `use_syslog`)

Enable Syslog log output

**SyslogTransport** (JSON: `syslog_transport`)

Syslong transport to use. Values: tcp or udp.

**SyslogNetworkAddr** (JSON: `syslog_network_addr`)

Graylog server address

**UseGraylog** (JSON: `use_graylog`)

Use Graylog log output

**GraylogNetworkAddr** (JSON: `graylog_network_addr`)

Graylog server address

**UseLogstash** (JSON: `use_logstash`)

Use logstash log output

**LogstashTransport** (JSON: `logstash_transport`)

Logstash network transport. Values: tcp or udp.

**LogstashNetworkAddr** (JSON: `logstash_network_addr`)

Logstash server address

**Track404Logs** (JSON: `track_404_logs`)

Show 404 HTTP errors in your Gateway application logs

**StatsdConnectionString** (JSON: `statsd_connection_string`)

Address of StatsD server. If set enable statsd monitoring.

**StatsdPrefix** (JSON: `statsd_prefix`)

StatsD prefix

**EventHandlers** (JSON: `event_handlers`)

Event System

**EventTriggers** (JSON: `event_trigers_defunct`)



> Deprecated: Config.GetEventTriggers instead.

**EventTriggersDefunct** (JSON: `event_triggers_defunct`)



> Deprecated: Config.GetEventTriggers instead.

**HideGeneratorHeader** (JSON: `hide_generator_header`)

HideGeneratorHeader will mask the 'X-Generator' and 'X-Mascot-...' headers, if set to true.

**SupressDefaultOrgStore** (JSON: `suppress_default_org_store`)



**LegacyEnableAllowanceCountdown** (JSON: `legacy_enable_allowance_countdown`)



**ForceGlobalSessionLifetime** (JSON: `force_global_session_lifetime`)

Enable global API token expiration. Can be needed if all your APIs using JWT or oAuth 2.0 auth methods with dynamically generated keys.

**SessionLifetimeRespectsKeyExpiration** (JSON: `session_lifetime_respects_key_expiration`)

SessionLifetimeRespectsKeyExpiration respects the key expiration time when the session lifetime is less than the key expiration. That is, Redis waits the key expiration for physical removal.

**GlobalSessionLifetime** (JSON: `global_session_lifetime`)

global session lifetime, in seconds.

**KV** (JSON: `kv`)

This section enables the use of the KV capabilities to substitute configuration values.
See more details https://tyk.io/docs/tyk-configuration-reference/kv-store/

**Secrets** (JSON: `secrets`)

Secrets are key-value pairs that can be accessed in the dashboard via "secrets://"

**OverrideMessages** (JSON: `override_messages`)

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

**Cloud** (JSON: `cloud`)

Cloud flag shows the Gateway runs in Tyk-cloud.

**JWTSSLInsecureSkipVerify** (JSON: `jwt_ssl_insecure_skip_verify`)

Skip TLS verification for JWT JWKs url validation

# SecurityConfig

**PrivateCertificateEncodingSecret** (JSON: `private_certificate_encoding_secret`)

Set the AES256 secret which is used to encode certificate private keys when they uploaded via certificate storage

**ControlAPIUseMutualTLS** (JSON: `control_api_use_mutual_tls`)

Enable Gateway Control API to use Mutual TLS. Certificates can be set via `security.certificates.control_api` section

**PinnedPublicKeys** (JSON: `pinned_public_keys`)

Specify public keys used for Certificate Pinning on global level.

**Certificates** (JSON: `certificates`)



# HttpServerOptionsConfig

**ReadTimeout** (JSON: `read_timeout`)

API Consumer -> Gateway network read timeout. Not setting this config, or setting this to 0, defaults to 120 seconds

**WriteTimeout** (JSON: `write_timeout`)

API Consumer -> Gateway network write timeout. Not setting this config, or setting this to 0, defaults to 120 seconds

**UseSSL** (JSON: `use_ssl`)

Set to true to enable SSL connections

**EnableHttp2** (JSON: `enable_http2`)

Enable HTTP2 protocol handling

**EnableStrictRoutes** (JSON: `enable_strict_routes`)

EnableStrictRoutes changes the routing to avoid nearest-neighbour requests on overlapping routes

- if disabled, `/apple` will route to `/app`, the current default behavior,
- if enabled, `/app` only responds to `/app`, `/app/` and `/app/*` but not `/apple`

Regular expressions and parameterized routes will be left alone regardless of this setting.

**SSLInsecureSkipVerify** (JSON: `ssl_insecure_skip_verify`)

Disable TLS verification. Required if you are using self-signed certificates.

**EnableWebSockets** (JSON: `enable_websockets`)

Enabled WebSockets and server side events support

**Certificates** (JSON: `certificates`)

Deprecated. SSL certificates used by Gateway server.

**SSLCertificates** (JSON: `ssl_certificates`)

SSL certificates used by your Gateway server. A list of certificate IDs or path to files.

**ServerName** (JSON: `server_name`)

Start your Gateway HTTP server on specific server name

**MinVersion** (JSON: `min_version`)

Minimum TLS version. Possible values: https://tyk.io/docs/basic-config-and-security/security/tls-and-ssl/#values-for-tls-versions

**MaxVersion** (JSON: `max_version`)

Maximum TLS version.

**SkipClientCAAnnouncement** (JSON: `skip_client_ca_announcement`)

When mTLS enabled, this option allows to skip client CA announcement in the TLS handshake.
This option is useful when you have a lot of ClientCAs and you want to reduce the handshake overhead, as some clients can hit TLS handshake limits.
This option does not give any hints to the client, on which certificate to pick (but this is very rare situation when it is required)

**FlushInterval** (JSON: `flush_interval`)

Set this to the number of seconds that Tyk uses to flush content from the proxied upstream connection to the open downstream connection.
This option needed be set for streaming protocols like Server Side Events, or gRPC streaming.

**SkipURLCleaning** (JSON: `skip_url_cleaning`)

Allow the use of a double slash in a URL path. This can be useful if you need to pass raw URLs to your API endpoints.
For example: `http://myapi.com/get/http://example.com`.

**SkipTargetPathEscaping** (JSON: `skip_target_path_escaping`)

Disable automatic character escaping, allowing to path original URL data to the upstream.

**Ciphers** (JSON: `ssl_ciphers`)

Custom SSL ciphers. See list of ciphers here https://tyk.io/docs/basic-config-and-security/security/tls-and-ssl/#specify-tls-cipher-suites-for-tyk-gateway--tyk-dashboard

**MaxRequestBodySize** (JSON: `max_request_body_size`)

MaxRequestBodySize configures the maximum request body size in bytes.

This option evaluates the `Content-Length` header and responds with
a HTTP 413 status code if larger than the defined size. If the header
is not provided, the request body is read up to the defined size.
If the request body is larger than the defined size, then we respond
with HTTP 413 status code.

See more information about setting request size limits here:
https://tyk.io/docs/basic-config-and-security/control-limit-traffic/request-size-limits/#maximum-request-sizes

# PoliciesConfig

**PolicySource** (JSON: `policy_source`)

Set this value to `file` to look in the file system for a definition file. Set to `service` to use the Dashboard service.

**PolicyConnectionString** (JSON: `policy_connection_string`)

This option is required if `policies.policy_source` is set to `service`.
Set this to the URL of your Tyk Dashboard installation. The URL needs to be formatted as: http://dashboard_host:port.

**PolicyRecordName** (JSON: `policy_record_name`)

This option is required if `policies.policy_source` is set to `file`.
Specifies the path of your JSON file containing the available policies.

**AllowExplicitPolicyID** (JSON: `allow_explicit_policy_id`)

In a Pro installation, Tyk will load Policy IDs and use the internal object-ID as the ID of the policy.
This is not portable in cases where the data needs to be moved from installation to installation.

If you set this value to `true`, then the id parameter in a stored policy (or imported policy using the Dashboard API), will be used instead of the internal ID.

This option should only be used when moving an installation to a new database.

**PolicyPath** (JSON: `policy_path`)

This option is used for storing a policies  if `policies.policy_source` is set to `file`.
it should be some existing file path on hard drive

# PortsWhiteList

No exposed fields available.

# DBAppConfOptionsConfig

**ConnectionString** (JSON: `connection_string`)

Set the URL to your Dashboard instance (or a load balanced instance). The URL needs to be formatted as: `http://dashboard_host:port`

**NodeIsSegmented** (JSON: `node_is_segmented`)

Set to `true` to enable filtering (sharding) of APIs.

**Tags** (JSON: `tags`)

The tags to use when filtering (sharding) Tyk Gateway nodes. Tags are processed as `OR` operations.
If you include a non-filter tag (e.g. an identifier such as `node-id-1`, this will become available to your Dashboard analytics).

# StorageOptionsConf

**Type** (JSON: `type`)

This should be set to `redis` (lowercase)

**Host** (JSON: `host`)

The Redis host, by default this is set to `localhost`, but for production this should be set to a cluster.

**Port** (JSON: `port`)

The Redis instance port.

**Hosts** (JSON: `hosts`)



> Deprecated: Addrs instead.

**Addrs** (JSON: `addrs`)

If you have multi-node setup, you should use this field instead. For example: ["host1:port1", "host2:port2"].

**MasterName** (JSON: `master_name`)

Redis sentinel master name

**SentinelPassword** (JSON: `sentinel_password`)

Redis sentinel password

**Username** (JSON: `username`)

Redis user name

**Password** (JSON: `password`)

If your Redis instance has a password set for access, you can set it here.

**Database** (JSON: `database`)

Redis database

**MaxIdle** (JSON: `optimisation_max_idle`)

Set the number of maximum idle connections in the Redis connection pool, which defaults to 100. Set to a higher value if you are expecting more traffic.

**MaxActive** (JSON: `optimisation_max_active`)

Set the number of maximum connections in the Redis connection pool, which defaults to 500. Set to a higher value if you are expecting more traffic.

**Timeout** (JSON: `timeout`)

Set a custom timeout for Redis network operations. Default value 5 seconds.

**EnableCluster** (JSON: `enable_cluster`)

Enable Redis Cluster support

**UseSSL** (JSON: `use_ssl`)

Enable SSL/TLS connection between your Tyk Gateway & Redis.

**SSLInsecureSkipVerify** (JSON: `ssl_insecure_skip_verify`)

Disable TLS verification

# SlaveOptionsConfig

**UseRPC** (JSON: `use_rpc`)

Set to `true` to connect a worker Gateway using RPC.

**UseSSL** (JSON: `use_ssl`)

Set this option to `true` to use an SSL RPC connection.

**SSLInsecureSkipVerify** (JSON: `ssl_insecure_skip_verify`)

Set this option to `true` to allow the certificate validation (certificate chain and hostname) to be skipped.
This can be useful if you use a self-signed certificate.

**ConnectionString** (JSON: `connection_string`)

Use this setting to add the URL for your MDCB or load balancer host.

**RPCKey** (JSON: `rpc_key`)

Your organisation ID to connect to the MDCB installation.

**APIKey** (JSON: `api_key`)

This the API key of a user used to authenticate and authorise the Gateway’s access through MDCB.
The user should be a standard Dashboard user with minimal privileges so as to reduce any risk if the user is compromised.
The suggested security settings are read for Real-time notifications and the remaining options set to deny.

**EnableRPCCache** (JSON: `enable_rpc_cache`)

Set this option to `true` to enable RPC caching for keys.

**BindToSlugsInsteadOfListenPaths** (JSON: `bind_to_slugs`)

For an Self-Managed installation this can be left at `false` (the default setting). For Legacy Cloud Gateways it must be set to ‘true’.

**DisableKeySpaceSync** (JSON: `disable_keyspace_sync`)

Set this option to `true` if you don’t want to monitor changes in the keys from a master Gateway.

**GroupID** (JSON: `group_id`)

This is the `zone` that this instance inhabits, e.g. the cluster/data-centre the Gateway lives in.
The group ID must be the same across all the Gateways of a data-centre/cluster which are also sharing the same Redis instance.
This ID should also be unique per cluster (otherwise another Gateway cluster can pick up your keyspace events and your cluster will get zero updates).

**CallTimeout** (JSON: `call_timeout`)

Call Timeout allows to specify a time in seconds for the maximum allowed duration of a RPC call.

**PingTimeout** (JSON: `ping_timeout`)

The maximum time in seconds that a RPC ping can last.

**RPCPoolSize** (JSON: `rpc_pool_size`)

The number of RPC connections in the pool. Basically it creates a set of connections that you can re-use as needed. Defaults to 5.

**KeySpaceSyncInterval** (JSON: `key_space_sync_interval`)

You can use this to set a period for which the Gateway will check if there are changes in keys that must be synchronized. If this value is not set then it will default to 10 seconds.

**RPCCertCacheExpiration** (JSON: `rpc_cert_cache_expiration`)

RPCCertCacheExpiration defines the expiration time of the rpc cache that stores the certificates, defined in seconds

**RPCGlobalCacheExpiration** (JSON: `rpc_global_cache_expiration`)

RPCKeysCacheExpiration defines the expiration time of the rpc cache that stores the keys, defined in seconds

**SynchroniserEnabled** (JSON: `synchroniser_enabled`)

SynchroniserEnabled enable this config if MDCB has enabled the synchoniser. If disabled then it will ignore signals to synchonise recources

# AuthOverrideConf

**ForceAuthProvider** (JSON: `force_auth_provider`)



**AuthProvider** (JSON: `auth_provider`)



**ForceSessionProvider** (JSON: `force_session_provider`)



**SessionProvider** (JSON: `session_provider`)



# MonitorConfig

**EnableTriggerMonitors** (JSON: `enable_trigger_monitors`)

Set this to `true` to have monitors enabled in your configuration for the node.

**Config** (JSON: `configuration`)



**GlobalTriggerLimit** (JSON: `global_trigger_limit`)

The trigger limit, as a percentage of the quota that must be reached in order to trigger the event, any time the quota percentage is increased the event will trigger.

**MonitorUserKeys** (JSON: `monitor_user_keys`)

Apply the monitoring subsystem to user keys.

**MonitorOrgKeys** (JSON: `monitor_org_keys`)

Apply the monitoring subsystem to organisation keys.

# ServiceDiscoveryConf

**DefaultCacheTimeout** (JSON: `default_cache_timeout`)

Service discovery cache timeout

# UptimeTestsConfig

**Disable** (JSON: `disable`)

To disable uptime tests on this node, set this value to `true`.

**PollerGroup** (JSON: `poller_group`)

If you have multiple Gateway clusters connected to the same Redis instance, you need to set a unique poller group for each cluster.

**Config** (JSON: `config`)



# HealthCheckConfig

**EnableHealthChecks** (JSON: `enable_health_checks`)

Setting this value to `true` will enable the health-check endpoint on /Tyk/health.

**HealthCheckValueTimeout** (JSON: `health_check_value_timeouts`)

This setting defaults to 60 seconds. This is the time window that Tyk uses to sample health-check data.
You can set a higher value for more accurate data (a larger sample period), or a lower value for less accurate data.
The reason this value is configurable is because sample data takes up space in your Redis DB to store the data to calculate samples. On high-availability systems this may not be desirable and smaller values may be preferred.

# AnalyticsConfigConfig

**Type** (JSON: `type`)

Set empty for a Self-Managed installation or `rpc` for multi-cloud.

**IgnoredIPs** (JSON: `ignored_ips`)

Adding IP addresses to this list will cause Tyk to ignore these IPs in the analytics data. These IP addresses will not produce an analytics log record.
This is useful for health checks and other samplers that might skew usage data.
The IP addresses must be provided as a JSON array, with the values being single IPs. CIDR values are not supported.

**EnableDetailedRecording** (JSON: `enable_detailed_recording`)

Set this value to `true` to have Tyk store the inbound request and outbound response data in HTTP Wire format as part of the Analytics data.
Please note, this will greatly increase your analytics DB size and can cause performance degradation on analytics processing by the Dashboard.
This setting can be overridden with an organisation flag, enabed at an API level, or on individual Key level.

**EnableGeoIP** (JSON: `enable_geo_ip`)

Tyk can store GeoIP information based on MaxMind DB’s to enable GeoIP tracking on inbound request analytics. Set this value to `true` and assign a DB using the `geo_ip_db_path` setting.

**GeoIPDBLocation** (JSON: `geo_ip_db_path`)

Path to a MaxMind GeoIP database
The analytics GeoIP DB can be replaced on disk. It will cleanly auto-reload every hour.

**NormaliseUrls** (JSON: `normalise_urls`)

This section describes methods that enable you to normalise inbound URLs in your analytics to have more meaningful per-path data.

**PoolSize** (JSON: `pool_size`)

Number of workers used to process analytics. Defaults to number of CPU cores.

**RecordsBufferSize** (JSON: `records_buffer_size`)

Number of records in analytics queue, per worker. Default: 1000.

**StorageExpirationTime** (JSON: `storage_expiration_time`)

You can set a time (in seconds) to configure how long analytics are kept if they are not processed. The default is 60 seconds.
This is used to prevent the potential infinite growth of Redis analytics storage.

**EnableMultipleAnalyticsKeys** (JSON: `enable_multiple_analytics_keys`)

Set this to `true` to have Tyk automatically divide the analytics records in multiple analytics keys.
This is especially useful when `storage.enable_cluster` is set to `true` since it will distribute the analytic keys across all the cluster nodes.

**PurgeInterval** (JSON: `purge_interval`)

You can set the interval length on how often the tyk Gateway will purge analytics data. This value is in seconds and defaults to 10 seconds.

**SerializerType** (JSON: `serializer_type`)

Determines the serialization engine for analytics. Available options: msgpack, and protobuf. By default, msgpack.

# LivenessCheckConfig

**CheckDuration** (JSON: `check_duration`)

Frequencies of performing interval healthchecks for Redis, Dashboard, and RPC layer. Default: 10 seconds.

# DnsCacheConfig

**Enabled** (JSON: `enabled`)

Setting this value to `true` will enable caching of DNS queries responses used for API endpoint’s host names. By default caching is disabled.

**TTL** (JSON: `ttl`)

This setting allows you to specify a duration in seconds before the record will be removed from cache after being added to it on the first DNS query resolution of API endpoints.
Setting `ttl` to `-1` prevents record from being expired and removed from cache on next check interval.

**MultipleIPsHandleStrategy** (JSON: `multiple_ips_handle_strategy`)

A strategy which will be used when a DNS query will reply with more than 1 IP Address per single host.
As a DNS query response IP Addresses can have a changing order depending on DNS server balancing strategy (eg: round robin, geographically dependent origin-ip ordering, etc) this option allows you to not to limit the connection to the first host in a cached response list or prevent response caching.

* `pick_first` will instruct your Tyk Gateway to connect to the first IP in a returned IP list and cache the response.
* `random` will instruct your Tyk Gateway to connect to a random IP in a returned IP list and cache the response.
* `no_cache` will instruct your Tyk Gateway to connect to the first IP in a returned IP list and fetch each addresses list without caching on each API endpoint DNS query.

# LocalSessionCacheConf

**DisableCacheSessionState** (JSON: `disable_cached_session_state`)

By default sessions are set to cache. Set this to `true` to stop Tyk from caching keys locally on the node.

**CachedSessionTimeout** (JSON: `cached_session_timeout`)



**CacheSessionEviction** (JSON: `cached_session_eviction`)



# CoProcessConfig

**EnableCoProcess** (JSON: `enable_coprocess`)

Enable gRPC and Python plugins

**CoProcessGRPCServer** (JSON: `coprocess_grpc_server`)

Address of gRPC user

**GRPCRecvMaxSize** (JSON: `grpc_recv_max_size`)

Maximum message which can be received from a gRPC server

**GRPCSendMaxSize** (JSON: `grpc_send_max_size`)

Maximum message which can be sent to gRPC server

**GRPCAuthority** (JSON: `grpc_authority`)

Authority used in GRPC connection

**PythonPathPrefix** (JSON: `python_path_prefix`)

Sets the path to built-in Tyk modules. This will be part of the Python module lookup path. The value used here is the default one for most installations.

**PythonVersion** (JSON: `python_version`)

If you have multiple Python versions installed you can specify your version.

# Tracer

**Name** (JSON: `name`)

The name of the tracer to initialize. For instance appdash, to use appdash tracer

**Enabled** (JSON: `enabled`)

Enable tracing

**Options** (JSON: `options`)

Tracing configuration. Refer to the Tracing Docs for the full list of options.

# NewRelicConfig

**AppName** (JSON: `app_name`)

New Relic Application name

**LicenseKey** (JSON: `license_key`)

New Relic License key

**EnableDistributedTracing** (JSON: `enable_distributed_tracing`)

Enable distributed tracing

# IPsHandleStrategy

No exposed fields available.

# NormalisedURLConfig

**Enabled** (JSON: `enabled`)

Set this to `true` to enable normalisation.

**NormaliseUUIDs** (JSON: `normalise_uuids`)

Each UUID will be replaced with a placeholder {uuid}

**NormaliseNumbers** (JSON: `normalise_numbers`)

Set this to true to have Tyk automatically match for numeric IDs, it will match with a preceding slash so as not to capture actual numbers:

**Custom** (JSON: `custom_patterns`)

This is a list of custom patterns you can add. These must be valid regex strings. Tyk will replace these values with a {var} placeholder.

# NormaliseURLPatterns

**UUIDs** (JSON: `UUIDs`)



**IDs** (JSON: `IDs`)



**Custom** (JSON: `Custom`)



# WebHookHandlerConf

**Method** (JSON: `method`)

The method to use for the webhook.

**TargetPath** (JSON: `target_path`)

The target path on which to send the request.

**TemplatePath** (JSON: `template_path`)

The template to load in order to format the request.

**HeaderList** (JSON: `header_map`)

Headers to set when firing the webhook.

**EventTimeout** (JSON: `event_timeout`)

The cool-down for the event so it does not trigger again (in seconds).

# CertsData

No exposed fields available.

# UptimeTestsConfigDetail

**FailureTriggerSampleSize** (JSON: `failure_trigger_sample_size`)

The sample size to trigger a `HostUp` or `HostDown` event. For example, a setting of 3 will require at least three failures to occur before the uptime test is triggered.

**TimeWait** (JSON: `time_wait`)

The value in seconds between tests runs. All tests will run simultaneously. This value will set the time between those tests. So a value of 60 will run all uptime tests every 60 seconds.

**CheckerPoolSize** (JSON: `checker_pool_size`)

The goroutine pool size to keep idle for uptime tests. If you have many uptime tests running at a high time period, then increase this value.

**EnableUptimeAnalytics** (JSON: `enable_uptime_analytics`)

Set this value to `true` to have the node capture and record analytics data regarding the uptime tests.

# CertificatesConfig

**API** (JSON: `apis`)



**Upstream** (JSON: `upstream`)

Specify upstream mutual TLS certificates at a global level in the following format: `{ "<host>": "<cert>" }``

**ControlAPI** (JSON: `control_api`)

Certificates used for Control API Mutual TLS

**Dashboard** (JSON: `dashboard_api`)

Used for communicating with the Dashboard if it is configured to use Mutual TLS

**MDCB** (JSON: `mdcb_api`)

Certificates used for MDCB Mutual TLS

# TykError

**Message** (JSON: `message`)



**Code** (JSON: `code`)



# CertData

**Name** (JSON: `domain_name`)

Domain name

**CertFile** (JSON: `cert_file`)

Path to certificate file

**KeyFile** (JSON: `key_file`)

Path to private key file

# Reporter

**URL** (JSON: `url`)

URL connection url to the zipkin server

**BatchSize** (JSON: `batch_size`)



**MaxBacklog** (JSON: `max_backlog`)



# Sampler

**Name** (JSON: `name`)

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

**Rate** (JSON: `rate`)

Rate is used by both "boundary" and "count" samplers

**Salt** (JSON: `salt`)

Salt is used by "boundary" sampler

**Mod** (JSON: `mod`)

Mod is only used when sampler is mod

# ConsulConfig

ConsulConfig is used to configure the creation of a client
This is a stripped down version of the Config struct in consul's API client

**Address** (JSON: `address`)

Address is the address of the Consul server

**Scheme** (JSON: `scheme`)

Scheme is the URI scheme for the Consul server

**Datacenter** (JSON: `datacenter`)

The datacenter to use. If not provided, the default agent datacenter is used.

**HttpAuth** (JSON: `http_auth`)

HttpAuth is the auth info to use for http access.

**WaitTime** (JSON: `wait_time`)

WaitTime limits how long a Watch will block. If not provided,
the agent default values will be used.

**Token** (JSON: `token`)

Token is used to provide a per-request ACL token
which overrides the agent's default token.

**TLSConfig** (JSON: `tls_config`)

TLS configuration

# EventMessage

EventMessage is a standard form to send event data to handlers

**Type** (JSON: `Type`)



**Meta** (JSON: `Meta`)



**TimeStamp** (JSON: `TimeStamp`)



# PortRange

PortRange defines a range of ports inclusively.

**From** (JSON: `from`)



**To** (JSON: `to`)



# PortWhiteList

PortWhiteList defines ports that will be allowed by the Gateway.

**Ranges** (JSON: `ranges`)



**Ports** (JSON: `ports`)



# ServicePort

ServicePort defines a protocol and port on which a service can bind to.

**Protocol** (JSON: `protocol`)



**Port** (JSON: `port`)



# TykEventHandler

TykEventHandler defines an event handler, e.g. LogMessageEventHandler will handle an event by logging it to stdout.

No exposed fields available.

# VaultConfig

VaultConfig is used to configure the creation of a client
This is a stripped down version of the config structure in vault's API client

**Address** (JSON: `address`)

Address is the address of the Vault server. This should be a complete
URL such as "http://vault.example.com".

**AgentAddress** (JSON: `agent_address`)

AgentAddress is the address of the local Vault agent. This should be a
complete URL such as "http://vault.example.com".

**MaxRetries** (JSON: `max_retries`)

MaxRetries controls the maximum number of times to retry when a vault
serer occurs

**Timeout** (JSON: `timeout`)



**Token** (JSON: `token`)

Token is the vault root token

**KVVersion** (JSON: `kv_version`)

KVVersion is the version number of Vault. Usually defaults to 2

# ZipkinConfig

ZipkinConfig configuration options used to initialize openzipkin opentracing
client.

**Reporter** (JSON: `reporter`)



**Sampler** (JSON: `sampler`)



