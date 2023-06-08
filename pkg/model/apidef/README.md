# APIDefinition

APIDefinition represents the configuration for a single proxied API and it's versions.

swagger:model

**Id** (JSON: `id`)



**Name** (JSON: `name`)



**Expiration** (JSON: `expiration`)



**Slug** (JSON: `slug`)



**ListenPort** (JSON: `listen_port`)



**Protocol** (JSON: `protocol`)



**EnableProxyProtocol** (JSON: `enable_proxy_protocol`)



**APIID** (JSON: `api_id`)



**OrgID** (JSON: `org_id`)



**UseKeylessAccess** (JSON: `use_keyless`)



**UseOauth2** (JSON: `use_oauth2`)



**ExternalOAuth** (JSON: `external_oauth`)



**UseOpenID** (JSON: `use_openid`)



**OpenIDOptions** (JSON: `openid_options`)



**Oauth2Meta** (JSON: `oauth_meta`)



**Auth** (JSON: `auth`)



> Deprecated: Use AuthConfigs instead.

**AuthConfigs** (JSON: `auth_configs`)



**UseBasicAuth** (JSON: `use_basic_auth`)



**BasicAuth** (JSON: `basic_auth`)



**UseMutualTLSAuth** (JSON: `use_mutual_tls_auth`)



**ClientCertificates** (JSON: `client_certificates`)



**UpstreamCertificates** (JSON: `upstream_certificates`)

UpstreamCertificates stores the domain to certificate mapping for upstream mutualTLS

**UpstreamCertificatesDisabled** (JSON: `upstream_certificates_disabled`)

UpstreamCertificatesDisabled disables upstream mutualTLS on the API

**PinnedPublicKeys** (JSON: `pinned_public_keys`)

PinnedPublicKeys stores the public key pinning details

**CertificatePinningDisabled** (JSON: `certificate_pinning_disabled`)

CertificatePinningDisabled disables public key pinning

**EnableJWT** (JSON: `enable_jwt`)



**UseStandardAuth** (JSON: `use_standard_auth`)



**UseGoPluginAuth** (JSON: `use_go_plugin_auth`)



> Deprecated. Use CustomPluginAuthEnabled instead.

**EnableCoProcessAuth** (JSON: `enable_coprocess_auth`)



> Deprecated. Use CustomPluginAuthEnabled instead.

**CustomPluginAuthEnabled** (JSON: `custom_plugin_auth_enabled`)



**JWTSigningMethod** (JSON: `jwt_signing_method`)



**JWTSource** (JSON: `jwt_source`)



**JWTIdentityBaseField** (JSON: `jwt_identity_base_field`)



**JWTClientIDBaseField** (JSON: `jwt_client_base_field`)



**JWTPolicyFieldName** (JSON: `jwt_policy_field_name`)



**JWTDefaultPolicies** (JSON: `jwt_default_policies`)



**JWTIssuedAtValidationSkew** (JSON: `jwt_issued_at_validation_skew`)



**JWTExpiresAtValidationSkew** (JSON: `jwt_expires_at_validation_skew`)



**JWTNotBeforeValidationSkew** (JSON: `jwt_not_before_validation_skew`)



**JWTSkipKid** (JSON: `jwt_skip_kid`)



**Scopes** (JSON: `scopes`)



**JWTScopeToPolicyMapping** (JSON: `jwt_scope_to_policy_mapping`)



> Deprecated: use Scopes.JWT.ScopeToPolicy or Scopes.OIDC.ScopeToPolicy

**JWTScopeClaimName** (JSON: `jwt_scope_claim_name`)



> Deprecated: use Scopes.JWT.ScopeClaimName or Scopes.OIDC.ScopeClaimName

**NotificationsDetails** (JSON: `notifications`)



**EnableSignatureChecking** (JSON: `enable_signature_checking`)



**HmacAllowedClockSkew** (JSON: `hmac_allowed_clock_skew`)



**HmacAllowedAlgorithms** (JSON: `hmac_allowed_algorithms`)



**RequestSigning** (JSON: `request_signing`)



**BaseIdentityProvidedBy** (JSON: `base_identity_provided_by`)



**VersionDefinition** (JSON: `definition`)



**VersionData** (JSON: `version_data`)



> Deprecated. Use VersionDefinition instead.

**UptimeTests** (JSON: `uptime_tests`)



**Proxy** (JSON: `proxy`)



**DisableRateLimit** (JSON: `disable_rate_limit`)



**DisableQuota** (JSON: `disable_quota`)



**CustomMiddleware** (JSON: `custom_middleware`)



**CustomMiddlewareBundle** (JSON: `custom_middleware_bundle`)



**CustomMiddlewareBundleDisabled** (JSON: `custom_middleware_bundle_disabled`)



**CacheOptions** (JSON: `cache_options`)



**SessionLifetimeRespectsKeyExpiration** (JSON: `session_lifetime_respects_key_expiration`)



**SessionLifetime** (JSON: `session_lifetime`)



**Active** (JSON: `active`)



**Internal** (JSON: `internal`)



**AuthProvider** (JSON: `auth_provider`)



**SessionProvider** (JSON: `session_provider`)



**EventHandlers** (JSON: `event_handlers`)



**EnableBatchRequestSupport** (JSON: `enable_batch_request_support`)



**EnableIpWhiteListing** (JSON: `enable_ip_whitelisting`)



**AllowedIPs** (JSON: `allowed_ips`)



**EnableIpBlacklisting** (JSON: `enable_ip_blacklisting`)



**BlacklistedIPs** (JSON: `blacklisted_ips`)



**DontSetQuotasOnCreate** (JSON: `dont_set_quota_on_create`)



**ExpireAnalyticsAfter** (JSON: `expire_analytics_after`)



> must have an expireAt TTL index set (http://docs.mongodb.org/manual/tutorial/expire-data/)

**ResponseProcessors** (JSON: `response_processors`)



**CORS** (JSON: `CORS`)



**Domain** (JSON: `domain`)



**DomainDisabled** (JSON: `domain_disabled`)



**Certificates** (JSON: `certificates`)



**DoNotTrack** (JSON: `do_not_track`)



**EnableContextVars** (JSON: `enable_context_vars`)



**ConfigData** (JSON: `config_data`)



**ConfigDataDisabled** (JSON: `config_data_disabled`)



**TagHeaders** (JSON: `tag_headers`)



**GlobalRateLimit** (JSON: `global_rate_limit`)



**StripAuthData** (JSON: `strip_auth_data`)



**EnableDetailedRecording** (JSON: `enable_detailed_recording`)



**GraphQL** (JSON: `graphql`)



**AnalyticsPlugin** (JSON: `analytics_plugin`)



**TagsDisabled** (JSON: `tags_disabled`)

Gateway segment tags

**Tags** (JSON: `tags`)



**IsOAS** (JSON: `is_oas`)

IsOAS is set to true when API has an OAS definition (created in OAS or migrated to OAS)

# ExternalOAuth

**Enabled** (JSON: `enabled`)



**Providers** (JSON: `providers`)



# OpenIDOptions

**Providers** (JSON: `providers`)



**SegregateByClient** (JSON: `segregate_by_client`)



# AuthConfig

**Name** (JSON: `name`)



**UseParam** (JSON: `use_param`)



**ParamName** (JSON: `param_name`)



**UseCookie** (JSON: `use_cookie`)



**CookieName** (JSON: `cookie_name`)



**DisableHeader** (JSON: `disable_header`)



**AuthHeaderName** (JSON: `auth_header_name`)



**UseCertificate** (JSON: `use_certificate`)



**ValidateSignature** (JSON: `validate_signature`)



**Signature** (JSON: `signature`)



# Scopes

**JWT** (JSON: `jwt`)



**OIDC** (JSON: `oidc`)



# NotificationsManager

NotificationsManager handles sending notifications to OAuth endpoints to notify the provider of key changes.
TODO: Make this more generic

**SharedSecret** (JSON: `shared_secret`)



**OAuthKeyChangeURL** (JSON: `oauth_on_keychange_url`)



# RequestSigningMeta

**IsEnabled** (JSON: `is_enabled`)



**Secret** (JSON: `secret`)



**KeyId** (JSON: `key_id`)



**Algorithm** (JSON: `algorithm`)



**HeaderList** (JSON: `header_list`)



**CertificateId** (JSON: `certificate_id`)



**SignatureHeader** (JSON: `signature_header`)



# AuthTypeEnum

No exposed fields available.

# VersionDefinition

**Enabled** (JSON: `enabled`)



**Name** (JSON: `name`)



**Default** (JSON: `default`)



**Location** (JSON: `location`)



**Key** (JSON: `key`)



**StripPath** (JSON: `strip_path`)



> Deprecated. Use StripVersioningData instead.

**StripVersioningData** (JSON: `strip_versioning_data`)



**Versions** (JSON: `versions`)



# VersionData

**NotVersioned** (JSON: `not_versioned`)



**DefaultVersion** (JSON: `default_version`)



**Versions** (JSON: `versions`)



# UptimeTests

**CheckList** (JSON: `check_list`)



**Config** (JSON: `config`)



# ProxyConfig

**PreserveHostHeader** (JSON: `preserve_host_header`)



**ListenPath** (JSON: `listen_path`)



**TargetURL** (JSON: `target_url`)



**DisableStripSlash** (JSON: `disable_strip_slash`)



**StripListenPath** (JSON: `strip_listen_path`)



**EnableLoadBalancing** (JSON: `enable_load_balancing`)



**Targets** (JSON: `target_list`)



**CheckHostAgainstUptimeTests** (JSON: `check_host_against_uptime_tests`)



**ServiceDiscovery** (JSON: `service_discovery`)



**Transport** (JSON: `transport`)



# MiddlewareSection

**Pre** (JSON: `pre`)



**Post** (JSON: `post`)



**PostKeyAuth** (JSON: `post_key_auth`)



**AuthCheck** (JSON: `auth_check`)



**Response** (JSON: `response`)



**Driver** (JSON: `driver`)



**IdExtractor** (JSON: `id_extractor`)



# CacheOptions

**CacheTimeout** (JSON: `cache_timeout`)



**EnableCache** (JSON: `enable_cache`)



**CacheAllSafeRequests** (JSON: `cache_all_safe_requests`)



**CacheOnlyResponseCodes** (JSON: `cache_response_codes`)



**EnableUpstreamCacheControl** (JSON: `enable_upstream_cache_control`)



**CacheControlTTLHeader** (JSON: `cache_control_ttl_header`)



**CacheByHeaders** (JSON: `cache_by_headers`)



# AuthProviderMeta

**Name** (JSON: `name`)



**StorageEngine** (JSON: `storage_engine`)



**Meta** (JSON: `meta`)



# SessionProviderMeta

**Name** (JSON: `name`)



**StorageEngine** (JSON: `storage_engine`)



**Meta** (JSON: `meta`)



# EventHandlerMetaConfig

**Events** (JSON: `events`)



# CORSConfig

**Enable** (JSON: `enable`)



**AllowedOrigins** (JSON: `allowed_origins`)



**AllowedMethods** (JSON: `allowed_methods`)



**AllowedHeaders** (JSON: `allowed_headers`)



**ExposedHeaders** (JSON: `exposed_headers`)



**AllowCredentials** (JSON: `allow_credentials`)



**MaxAge** (JSON: `max_age`)



**OptionsPassthrough** (JSON: `options_passthrough`)



**Debug** (JSON: `debug`)



# GlobalRateLimit

**Rate** (JSON: `rate`)



**Per** (JSON: `per`)



# GraphQLConfig

GraphQLConfig is the root config object for a GraphQL API.

**Enabled** (JSON: `enabled`)

Enabled indicates if GraphQL should be enabled.

**ExecutionMode** (JSON: `execution_mode`)

ExecutionMode is the mode to define how an api behaves.

**Version** (JSON: `version`)

Version defines the version of the GraphQL config and engine to be used.

**Schema** (JSON: `schema`)

Schema is the GraphQL Schema exposed by the GraphQL API/Upstream/Engine.

**LastSchemaUpdate** (JSON: `last_schema_update`)

LastSchemaUpdate contains the date and time of the last triggered schema update to the upstream.

**TypeFieldConfigurations** (JSON: `type_field_configurations`)

TypeFieldConfigurations is a rule set of data source and mapping of a schema field.

**GraphQLPlayground** (JSON: `playground`)

GraphQLPlayground is the Playground specific configuration.

**Engine** (JSON: `engine`)

Engine holds the configuration for engine v2 and upwards.

**Proxy** (JSON: `proxy`)

Proxy holds the configuration for a proxy only api.

**Subgraph** (JSON: `subgraph`)

Subgraph holds the configuration for a GraphQL federation subgraph.

**Supergraph** (JSON: `supergraph`)

Supergraph holds the configuration for a GraphQL federation supergraph.

# AnalyticsPluginConfig

**Enabled** (JSON: `enable`)



**PluginPath** (JSON: `plugin_path`)



**FuncName** (JSON: `func_name`)



# ValidationResult

**IsValid** (JSON: `IsValid`)



**Errors** (JSON: `Errors`)



# ValidationRuleSet

No exposed fields available.

# ValidationRule

No exposed fields available.

# RuleUniqueDataSourceNames

No exposed fields available.

# RuleAtLeastEnableOneAuthSource

No exposed fields available.

# RuleValidateIPList

No exposed fields available.

# AuthProviderCode

No exposed fields available.

# SessionProviderCode

No exposed fields available.

# StorageEngineCode

No exposed fields available.

# TykEvent

No exposed fields available.

# TykEventHandlerName

No exposed fields available.

# EndpointMethodAction

No exposed fields available.

# SourceMode

No exposed fields available.

# MiddlewareDriver

No exposed fields available.

# IdExtractorSource

No exposed fields available.

# IdExtractorType

No exposed fields available.

# RoutingTriggerOnType

No exposed fields available.

# SubscriptionType

No exposed fields available.

# IDExtractor

No exposed fields available.

# EndpointMethodMeta

**Action** (JSON: `action`)



**Code** (JSON: `code`)



**Data** (JSON: `data`)



**Headers** (JSON: `headers`)



# MockResponseMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**IgnoreCase** (JSON: `ignore_case`)



**Code** (JSON: `code`)



**Body** (JSON: `body`)



**Headers** (JSON: `headers`)



# EndPointMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**IgnoreCase** (JSON: `ignore_case`)



**MethodActions** (JSON: `method_actions`)

Deprecated. Use Method instead.

# CacheMeta

**Disabled** (JSON: `disabled`)



**Method** (JSON: `method`)



**Path** (JSON: `path`)



**CacheKeyRegex** (JSON: `cache_key_regex`)



**CacheOnlyResponseCodes** (JSON: `cache_response_codes`)



# RequestInputType

No exposed fields available.

# TemplateData

**Input** (JSON: `input_type`)



**Mode** (JSON: `template_mode`)



**EnableSession** (JSON: `enable_session`)



**TemplateSource** (JSON: `template_source`)



# TemplateMeta

**Disabled** (JSON: `disabled`)



**TemplateData** (JSON: `template_data`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



# TransformJQMeta

**Filter** (JSON: `filter`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



# HeaderInjectionMeta

**DeleteHeaders** (JSON: `delete_headers`)



**AddHeaders** (JSON: `add_headers`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**ActOnResponse** (JSON: `act_on`)



# HardTimeoutMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**TimeOut** (JSON: `timeout`)



# TrackEndpointMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



# InternalMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



# RequestSizeMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



**SizeLimit** (JSON: `size_limit`)



# CircuitBreakerMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



**ThresholdPercent** (JSON: `threshold_percent`)



**Samples** (JSON: `samples`)



**ReturnToServiceAfter** (JSON: `return_to_service_after`)



**DisableHalfOpenState** (JSON: `disable_half_open_state`)



# StringRegexMap

**MatchPattern** (JSON: `match_rx`)



**Reverse** (JSON: `reverse`)



# RoutingTriggerOptions

**HeaderMatches** (JSON: `header_matches`)



**QueryValMatches** (JSON: `query_val_matches`)



**PathPartMatches** (JSON: `path_part_matches`)



**SessionMetaMatches** (JSON: `session_meta_matches`)



**RequestContextMatches** (JSON: `request_context_matches`)



**PayloadMatches** (JSON: `payload_matches`)



# RoutingTrigger

**On** (JSON: `on`)



**Options** (JSON: `options`)



**RewriteTo** (JSON: `rewrite_to`)



# URLRewriteMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



**MatchPattern** (JSON: `match_pattern`)



**RewriteTo** (JSON: `rewrite_to`)



**Triggers** (JSON: `triggers`)



# VirtualMeta

**Disabled** (JSON: `disabled`)



**ResponseFunctionName** (JSON: `response_function_name`)



**FunctionSourceType** (JSON: `function_source_type`)



**FunctionSourceURI** (JSON: `function_source_uri`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**UseSession** (JSON: `use_session`)



**ProxyOnError** (JSON: `proxy_on_error`)



# MethodTransformMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**ToMethod** (JSON: `to_method`)



# ValidatePathMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**Schema** (JSON: `schema`)



**SchemaB64** (JSON: `schema_b64`)



**ErrorResponseCode** (JSON: `error_response_code`)

Allows override of default 422 Unprocessible Entity response code for validation errors.

# ValidateRequestMeta

**Enabled** (JSON: `enabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**ErrorResponseCode** (JSON: `error_response_code`)

Allows override of default 422 Unprocessible Entity response code for validation errors.

# PersistGraphQLMeta

**Path** (JSON: `path`)



**Method** (JSON: `method`)



**Operation** (JSON: `operation`)



**Variables** (JSON: `variables`)



# GoPluginMeta

**Disabled** (JSON: `disabled`)



**Path** (JSON: `path`)



**Method** (JSON: `method`)



**PluginPath** (JSON: `plugin_path`)



**SymbolName** (JSON: `func_name`)



# ExtendedPathsSet

**Ignored** (JSON: `ignored`)



**WhiteList** (JSON: `white_list`)



**BlackList** (JSON: `black_list`)



**MockResponse** (JSON: `mock_response`)



**Cached** (JSON: `cache`)



**AdvanceCacheConfig** (JSON: `advance_cache_config`)



**Transform** (JSON: `transform`)



**TransformResponse** (JSON: `transform_response`)



**TransformJQ** (JSON: `transform_jq`)



**TransformJQResponse** (JSON: `transform_jq_response`)



**TransformHeader** (JSON: `transform_headers`)



**TransformResponseHeader** (JSON: `transform_response_headers`)



**HardTimeouts** (JSON: `hard_timeouts`)



**CircuitBreaker** (JSON: `circuit_breakers`)



**URLRewrite** (JSON: `url_rewrites`)



**Virtual** (JSON: `virtual`)



**SizeLimit** (JSON: `size_limits`)



**MethodTransforms** (JSON: `method_transforms`)



**TrackEndpoints** (JSON: `track_endpoints`)



**DoNotTrackEndpoints** (JSON: `do_not_track_endpoints`)



**ValidateJSON** (JSON: `validate_json`)



**ValidateRequest** (JSON: `validate_request`)



**Internal** (JSON: `internal`)



**GoPlugin** (JSON: `go_plugin`)



**PersistGraphQL** (JSON: `persist_graphql`)



# VersionInfo

**Name** (JSON: `name`)



**Expires** (JSON: `expires`)



**Paths** (JSON: `paths`)



**UseExtendedPaths** (JSON: `use_extended_paths`)



**ExtendedPaths** (JSON: `extended_paths`)



**GlobalHeaders** (JSON: `global_headers`)



**GlobalHeadersRemove** (JSON: `global_headers_remove`)



**GlobalResponseHeaders** (JSON: `global_response_headers`)



**GlobalResponseHeadersRemove** (JSON: `global_response_headers_remove`)



**IgnoreEndpointCase** (JSON: `ignore_endpoint_case`)



**GlobalSizeLimit** (JSON: `global_size_limit`)



**OverrideTarget** (JSON: `override_target`)



# EventHandlerTriggerConfig

**Handler** (JSON: `handler_name`)



**HandlerMeta** (JSON: `handler_meta`)



# MiddlewareDefinition

**Disabled** (JSON: `disabled`)



**Name** (JSON: `name`)



**Path** (JSON: `path`)



**RequireSession** (JSON: `require_session`)



**RawBodyOnly** (JSON: `raw_body_only`)



# MiddlewareIdExtractor

**Disabled** (JSON: `disabled`)



**ExtractFrom** (JSON: `extract_from`)



**ExtractWith** (JSON: `extract_with`)



**ExtractorConfig** (JSON: `extractor_config`)



# ResponseProcessor

**Name** (JSON: `name`)



**Options** (JSON: `options`)



# HostCheckObject

**CheckURL** (JSON: `url`)



**Protocol** (JSON: `protocol`)



**Timeout** (JSON: `timeout`)



**EnableProxyProtocol** (JSON: `enable_proxy_protocol`)



**Commands** (JSON: `commands`)



**Method** (JSON: `method`)



**Headers** (JSON: `headers`)



**Body** (JSON: `body`)



# CheckCommand

**Name** (JSON: `name`)



**Message** (JSON: `message`)



# ServiceDiscoveryConfiguration

**UseDiscoveryService** (JSON: `use_discovery_service`)



**QueryEndpoint** (JSON: `query_endpoint`)



**UseNestedQuery** (JSON: `use_nested_query`)



**ParentDataPath** (JSON: `parent_data_path`)



**DataPath** (JSON: `data_path`)



**PortDataPath** (JSON: `port_data_path`)



**TargetPath** (JSON: `target_path`)



**UseTargetList** (JSON: `use_target_list`)



**CacheDisabled** (JSON: `cache_disabled`)



**CacheTimeout** (JSON: `cache_timeout`)



**EndpointReturnsList** (JSON: `endpoint_returns_list`)



# OIDProviderConfig

**Issuer** (JSON: `issuer`)



**ClientIDs** (JSON: `client_ids`)



# ScopeClaim

**ScopeClaimName** (JSON: `scope_claim_name`)



**ScopeToPolicy** (JSON: `scope_to_policy`)



# UptimeTestsConfig

**ExpireUptimeAnalyticsAfter** (JSON: `expire_utime_after`)



> must have an expireAt TTL index set (http://docs.mongodb.org/manual/tutorial/expire-data/)

**ServiceDiscovery** (JSON: `service_discovery`)



**RecheckWait** (JSON: `recheck_wait`)



# SignatureConfig

**Algorithm** (JSON: `algorithm`)



**Header** (JSON: `header`)



**UseParam** (JSON: `use_param`)



**ParamName** (JSON: `param_name`)



**Secret** (JSON: `secret`)



**AllowedClockSkew** (JSON: `allowed_clock_skew`)



**ErrorCode** (JSON: `error_code`)



**ErrorMessage** (JSON: `error_message`)



# BundleManifest

**FileList** (JSON: `file_list`)



**CustomMiddleware** (JSON: `custom_middleware`)



**Checksum** (JSON: `checksum`)



**Signature** (JSON: `signature`)



# GraphQLConfigVersion

No exposed fields available.

# GraphQLResponseExtensions

**OnErrorForwarding** (JSON: `on_error_forwarding`)



# GraphQLProxyConfig

**AuthHeaders** (JSON: `auth_headers`)



**SubscriptionType** (JSON: `subscription_type`)



**RequestHeaders** (JSON: `request_headers`)



**UseResponseExtensions** (JSON: `use_response_extensions`)



# GraphQLSubgraphConfig

**SDL** (JSON: `sdl`)



# GraphQLSupergraphConfig

**UpdatedAt** (JSON: `updated_at`)

UpdatedAt contains the date and time of the last update of a supergraph API.

**Subgraphs** (JSON: `subgraphs`)



**MergedSDL** (JSON: `merged_sdl`)



**GlobalHeaders** (JSON: `global_headers`)



**DisableQueryBatching** (JSON: `disable_query_batching`)



# GraphQLSubgraphEntity

**APIID** (JSON: `api_id`)



**Name** (JSON: `name`)



**URL** (JSON: `url`)



**SDL** (JSON: `sdl`)



**Headers** (JSON: `headers`)



**SubscriptionType** (JSON: `subscription_type`)



# GraphQLEngineConfig

**FieldConfigs** (JSON: `field_configs`)



**DataSources** (JSON: `data_sources`)



# GraphQLFieldConfig

**TypeName** (JSON: `type_name`)



**FieldName** (JSON: `field_name`)



**DisableDefaultMapping** (JSON: `disable_default_mapping`)



**Path** (JSON: `path`)



# GraphQLEngineDataSourceKind

No exposed fields available.

# GraphQLEngineDataSource

**Kind** (JSON: `kind`)



**Name** (JSON: `name`)



**Internal** (JSON: `internal`)



**RootFields** (JSON: `root_fields`)



**Config** (JSON: `config`)



# GraphQLTypeFields

**Type** (JSON: `type`)



**Fields** (JSON: `fields`)



# GraphQLEngineDataSourceConfigREST

**URL** (JSON: `url`)



**Method** (JSON: `method`)



**Headers** (JSON: `headers`)



**Query** (JSON: `query`)



**Body** (JSON: `body`)



# GraphQLEngineDataSourceConfigGraphQL

**URL** (JSON: `url`)



**Method** (JSON: `method`)



**Headers** (JSON: `headers`)



**SubscriptionType** (JSON: `subscription_type`)



**HasOperation** (JSON: `has_operation`)



**Operation** (JSON: `operation`)



**Variables** (JSON: `variables`)



# GraphQLEngineDataSourceConfigKafka

**BrokerAddresses** (JSON: `broker_addresses`)



**Topics** (JSON: `topics`)



**GroupID** (JSON: `group_id`)



**ClientID** (JSON: `client_id`)



**KafkaVersion** (JSON: `kafka_version`)



**StartConsumingLatest** (JSON: `start_consuming_latest`)



**BalanceStrategy** (JSON: `balance_strategy`)



**IsolationLevel** (JSON: `isolation_level`)



**SASL** (JSON: `sasl`)



# QueryVariable

**Name** (JSON: `name`)



**Value** (JSON: `value`)



# Provider

**JWT** (JSON: `jwt`)



**Introspection** (JSON: `introspection`)



# JWTValidation

**Enabled** (JSON: `enabled`)



**SigningMethod** (JSON: `signing_method`)



**Source** (JSON: `source`)



**IssuedAtValidationSkew** (JSON: `issued_at_validation_skew`)



**NotBeforeValidationSkew** (JSON: `not_before_validation_skew`)



**ExpiresAtValidationSkew** (JSON: `expires_at_validation_skew`)



**IdentityBaseField** (JSON: `identity_base_field`)



# Introspection

**Enabled** (JSON: `enabled`)



**URL** (JSON: `url`)



**ClientID** (JSON: `client_id`)



**ClientSecret** (JSON: `client_secret`)



**IdentityBaseField** (JSON: `identity_base_field`)



**Cache** (JSON: `cache`)



# IntrospectionCache

**Enabled** (JSON: `enabled`)



**Timeout** (JSON: `timeout`)



# HostList

No exposed fields available.

# InboundData

**KeyName** (JSON: `KeyName`)



**Value** (JSON: `Value`)



**SessionState** (JSON: `SessionState`)



**Timeout** (JSON: `Timeout`)



**Per** (JSON: `Per`)



**Expire** (JSON: `Expire`)



# DefRequest

**OrgId** (JSON: `OrgId`)



**Tags** (JSON: `Tags`)



**LoadOAS** (JSON: `LoadOAS`)



# GroupLoginRequest

**UserKey** (JSON: `UserKey`)



**GroupID** (JSON: `GroupID`)



**ForceSync** (JSON: `ForceSync`)



# GroupKeySpaceRequest

**OrgID** (JSON: `OrgID`)



**GroupID** (JSON: `GroupID`)



# KeysValuesPair

**Keys** (JSON: `Keys`)



**Values** (JSON: `Values`)



# GraphQLExecutionMode

GraphQLExecutionMode is the mode in which the GraphQL Middleware should operate.

No exposed fields available.

# GraphQLPlayground

GraphQLPlayground represents the configuration for the public playground which will be hosted alongside the api.

**Enabled** (JSON: `enabled`)

Enabled indicates if the playground should be enabled.

**Path** (JSON: `path`)

Path sets the path on which the playground will be hosted if enabled.

# IDExtractorConfig

IDExtractorConfig specifies the configuration for ID extractor

**HeaderName** (JSON: `header_name`)

HeaderName is the header name to extract ID from.

**FormParamName** (JSON: `param_name`)

FormParamName is the form parameter name to extract ID from.

**RegexExpression** (JSON: `regex_expression`)

RegexExpression is the regular expression to match ID.

**RegexMatchIndex** (JSON: `regex_match_index`)

RegexMatchIndex is the index from which ID to be extracted after a match.

**XPathExpression** (JSON: `xpath_expression`)

XPathExp is the xpath expression to match ID.

