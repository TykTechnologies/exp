# APIDefinition

APIDefinition represents the configuration for a single proxied API and it's versions.

swagger:model

**Field: `id`** (Id, `model.ObjectID`)



**Field: `name`** (Name, `string`)



**Field: `expiration`** (Expiration, `string`)



**Field: `slug`** (Slug, `string`)



**Field: `listen_port`** (ListenPort, `int`)



**Field: `protocol`** (Protocol, `string`)



**Field: `enable_proxy_protocol`** (EnableProxyProtocol, `bool`)



**Field: `api_id`** (APIID, `string`)



**Field: `org_id`** (OrgID, `string`)



**Field: `use_keyless`** (UseKeylessAccess, `bool`)



**Field: `use_oauth2`** (UseOauth2, `bool`)



**Field: `external_oauth`** (ExternalOAuth, [ExternalOAuth](#ExternalOAuth))



**Field: `use_openid`** (UseOpenID, `bool`)



**Field: `openid_options`** (OpenIDOptions, [OpenIDOptions](#OpenIDOptions))



**Field: `oauth_meta`** (Oauth2Meta, `struct{}`)



**Field: `auth`** (Auth, [AuthConfig](#AuthConfig))



> Deprecated: Use AuthConfigs instead.

**Field: `auth_configs`** (AuthConfigs, `map[string]AuthConfig`)



**Field: `use_basic_auth`** (UseBasicAuth, `bool`)



**Field: `basic_auth`** (BasicAuth, `struct{}`)



**Field: `use_mutual_tls_auth`** (UseMutualTLSAuth, `bool`)



**Field: `client_certificates`** (ClientCertificates, `[]string`)



**Field: `upstream_certificates`** (UpstreamCertificates, `map[string]string`)

UpstreamCertificates stores the domain to certificate mapping for upstream mutualTLS

**Field: `upstream_certificates_disabled`** (UpstreamCertificatesDisabled, `bool`)

UpstreamCertificatesDisabled disables upstream mutualTLS on the API

**Field: `pinned_public_keys`** (PinnedPublicKeys, `map[string]string`)

PinnedPublicKeys stores the public key pinning details

**Field: `certificate_pinning_disabled`** (CertificatePinningDisabled, `bool`)

CertificatePinningDisabled disables public key pinning

**Field: `enable_jwt`** (EnableJWT, `bool`)



**Field: `use_standard_auth`** (UseStandardAuth, `bool`)



**Field: `use_go_plugin_auth`** (UseGoPluginAuth, `bool`)



> Deprecated. Use CustomPluginAuthEnabled instead.

**Field: `enable_coprocess_auth`** (EnableCoProcessAuth, `bool`)



> Deprecated. Use CustomPluginAuthEnabled instead.

**Field: `custom_plugin_auth_enabled`** (CustomPluginAuthEnabled, `bool`)



**Field: `jwt_signing_method`** (JWTSigningMethod, `string`)



**Field: `jwt_source`** (JWTSource, `string`)



**Field: `jwt_identity_base_field`** (JWTIdentityBaseField, `string`)



**Field: `jwt_client_base_field`** (JWTClientIDBaseField, `string`)



**Field: `jwt_policy_field_name`** (JWTPolicyFieldName, `string`)



**Field: `jwt_default_policies`** (JWTDefaultPolicies, `[]string`)



**Field: `jwt_issued_at_validation_skew`** (JWTIssuedAtValidationSkew, `uint64`)



**Field: `jwt_expires_at_validation_skew`** (JWTExpiresAtValidationSkew, `uint64`)



**Field: `jwt_not_before_validation_skew`** (JWTNotBeforeValidationSkew, `uint64`)



**Field: `jwt_skip_kid`** (JWTSkipKid, `bool`)



**Field: `scopes`** (Scopes, [Scopes](#Scopes))



**Field: `jwt_scope_to_policy_mapping`** (JWTScopeToPolicyMapping, `map[string]string`)



> Deprecated: use Scopes.JWT.ScopeToPolicy or Scopes.OIDC.ScopeToPolicy

**Field: `jwt_scope_claim_name`** (JWTScopeClaimName, `string`)



> Deprecated: use Scopes.JWT.ScopeClaimName or Scopes.OIDC.ScopeClaimName

**Field: `notifications`** (NotificationsDetails, [NotificationsManager](#NotificationsManager))



**Field: `enable_signature_checking`** (EnableSignatureChecking, `bool`)



**Field: `hmac_allowed_clock_skew`** (HmacAllowedClockSkew, `float64`)



**Field: `hmac_allowed_algorithms`** (HmacAllowedAlgorithms, `[]string`)



**Field: `request_signing`** (RequestSigning, [RequestSigningMeta](#RequestSigningMeta))



**Field: `base_identity_provided_by`** (BaseIdentityProvidedBy, [AuthTypeEnum](#AuthTypeEnum))



**Field: `definition`** (VersionDefinition, [VersionDefinition](#VersionDefinition))



**Field: `version_data`** (VersionData, [VersionData](#VersionData))



> Deprecated. Use VersionDefinition instead.

**Field: `uptime_tests`** (UptimeTests, [UptimeTests](#UptimeTests))



**Field: `proxy`** (Proxy, [ProxyConfig](#ProxyConfig))



**Field: `disable_rate_limit`** (DisableRateLimit, `bool`)



**Field: `disable_quota`** (DisableQuota, `bool`)



**Field: `custom_middleware`** (CustomMiddleware, [MiddlewareSection](#MiddlewareSection))



**Field: `custom_middleware_bundle`** (CustomMiddlewareBundle, `string`)



**Field: `custom_middleware_bundle_disabled`** (CustomMiddlewareBundleDisabled, `bool`)



**Field: `cache_options`** (CacheOptions, [CacheOptions](#CacheOptions))



**Field: `session_lifetime_respects_key_expiration`** (SessionLifetimeRespectsKeyExpiration, `bool`)



**Field: `session_lifetime`** (SessionLifetime, `int64`)



**Field: `active`** (Active, `bool`)



**Field: `internal`** (Internal, `bool`)



**Field: `auth_provider`** (AuthProvider, [AuthProviderMeta](#AuthProviderMeta))



**Field: `session_provider`** (SessionProvider, [SessionProviderMeta](#SessionProviderMeta))



**Field: `event_handlers`** (EventHandlers, [EventHandlerMetaConfig](#EventHandlerMetaConfig))



**Field: `enable_batch_request_support`** (EnableBatchRequestSupport, `bool`)



**Field: `enable_ip_whitelisting`** (EnableIpWhiteListing, `bool`)



**Field: `allowed_ips`** (AllowedIPs, `[]string`)



**Field: `enable_ip_blacklisting`** (EnableIpBlacklisting, `bool`)



**Field: `blacklisted_ips`** (BlacklistedIPs, `[]string`)



**Field: `dont_set_quota_on_create`** (DontSetQuotasOnCreate, `bool`)



**Field: `expire_analytics_after`** (ExpireAnalyticsAfter, `int64`)



> must have an expireAt TTL index set (http://docs.mongodb.org/manual/tutorial/expire-data/)

**Field: `response_processors`** (ResponseProcessors, [[]ResponseProcessor](#ResponseProcessor))



**Field: `CORS`** (CORS, [CORSConfig](#CORSConfig))



**Field: `domain`** (Domain, `string`)



**Field: `domain_disabled`** (DomainDisabled, `bool`)



**Field: `certificates`** (Certificates, `[]string`)



**Field: `do_not_track`** (DoNotTrack, `bool`)



**Field: `enable_context_vars`** (EnableContextVars, `bool`)



**Field: `config_data`** (ConfigData, `map[string]interface{}`)



**Field: `config_data_disabled`** (ConfigDataDisabled, `bool`)



**Field: `tag_headers`** (TagHeaders, `[]string`)



**Field: `global_rate_limit`** (GlobalRateLimit, [GlobalRateLimit](#GlobalRateLimit))



**Field: `strip_auth_data`** (StripAuthData, `bool`)



**Field: `enable_detailed_recording`** (EnableDetailedRecording, `bool`)



**Field: `graphql`** (GraphQL, [GraphQLConfig](#GraphQLConfig))



**Field: `analytics_plugin`** (AnalyticsPlugin, [AnalyticsPluginConfig](#AnalyticsPluginConfig))



**Field: `tags_disabled`** (TagsDisabled, `bool`)

Gateway segment tags

**Field: `tags`** (Tags, `[]string`)



**Field: `is_oas`** (IsOAS, `bool`)

IsOAS is set to true when API has an OAS definition (created in OAS or migrated to OAS)

# ExternalOAuth

**Field: `enabled`** (Enabled, `bool`)



**Field: `providers`** (Providers, [[]Provider](#Provider))



# OpenIDOptions

**Field: `providers`** (Providers, [[]OIDProviderConfig](#OIDProviderConfig))



**Field: `segregate_by_client`** (SegregateByClient, `bool`)



# AuthConfig

**Field: `name`** (Name, `string`)



**Field: `use_param`** (UseParam, `bool`)



**Field: `param_name`** (ParamName, `string`)



**Field: `use_cookie`** (UseCookie, `bool`)



**Field: `cookie_name`** (CookieName, `string`)



**Field: `disable_header`** (DisableHeader, `bool`)



**Field: `auth_header_name`** (AuthHeaderName, `string`)



**Field: `use_certificate`** (UseCertificate, `bool`)



**Field: `validate_signature`** (ValidateSignature, `bool`)



**Field: `signature`** (Signature, [SignatureConfig](#SignatureConfig))



# Scopes

**Field: `jwt`** (JWT, [ScopeClaim](#ScopeClaim))



**Field: `oidc`** (OIDC, [ScopeClaim](#ScopeClaim))



# NotificationsManager

NotificationsManager handles sending notifications to OAuth endpoints to notify the provider of key changes.
TODO: Make this more generic

**Field: `shared_secret`** (SharedSecret, `string`)



**Field: `oauth_on_keychange_url`** (OAuthKeyChangeURL, `string`)



# RequestSigningMeta

**Field: `is_enabled`** (IsEnabled, `bool`)



**Field: `secret`** (Secret, `string`)



**Field: `key_id`** (KeyId, `string`)



**Field: `algorithm`** (Algorithm, `string`)



**Field: `header_list`** (HeaderList, `[]string`)



**Field: `certificate_id`** (CertificateId, `string`)



**Field: `signature_header`** (SignatureHeader, `string`)



# AuthTypeEnum

Type defined as `string`, see [string](string) definition.

# VersionDefinition

**Field: `enabled`** (Enabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `default`** (Default, `string`)



**Field: `location`** (Location, `string`)



**Field: `key`** (Key, `string`)



**Field: `strip_path`** (StripPath, `bool`)



> Deprecated. Use StripVersioningData instead.

**Field: `strip_versioning_data`** (StripVersioningData, `bool`)



**Field: `versions`** (Versions, `map[string]string`)



# VersionData

**Field: `not_versioned`** (NotVersioned, `bool`)



**Field: `default_version`** (DefaultVersion, `string`)



**Field: `versions`** (Versions, `map[string]VersionInfo`)



# UptimeTests

**Field: `check_list`** (CheckList, [[]HostCheckObject](#HostCheckObject))



**Field: `config`** (Config, [UptimeTestsConfig](#UptimeTestsConfig))



# ProxyConfig

**Field: `preserve_host_header`** (PreserveHostHeader, `bool`)



**Field: `listen_path`** (ListenPath, `string`)



**Field: `target_url`** (TargetURL, `string`)



**Field: `disable_strip_slash`** (DisableStripSlash, `bool`)



**Field: `strip_listen_path`** (StripListenPath, `bool`)



**Field: `enable_load_balancing`** (EnableLoadBalancing, `bool`)



**Field: `target_list`** (Targets, `[]string`)



**Field: `check_host_against_uptime_tests`** (CheckHostAgainstUptimeTests, `bool`)



**Field: `service_discovery`** (ServiceDiscovery, [ServiceDiscoveryConfiguration](#ServiceDiscoveryConfiguration))



**Field: `transport`** (Transport, `struct{}`)



# MiddlewareSection

**Field: `pre`** (Pre, [[]MiddlewareDefinition](#MiddlewareDefinition))



**Field: `post`** (Post, [[]MiddlewareDefinition](#MiddlewareDefinition))



**Field: `post_key_auth`** (PostKeyAuth, [[]MiddlewareDefinition](#MiddlewareDefinition))



**Field: `auth_check`** (AuthCheck, [MiddlewareDefinition](#MiddlewareDefinition))



**Field: `response`** (Response, [[]MiddlewareDefinition](#MiddlewareDefinition))



**Field: `driver`** (Driver, [MiddlewareDriver](#MiddlewareDriver))



**Field: `id_extractor`** (IdExtractor, [MiddlewareIdExtractor](#MiddlewareIdExtractor))



# CacheOptions

**Field: `cache_timeout`** (CacheTimeout, `int64`)



**Field: `enable_cache`** (EnableCache, `bool`)



**Field: `cache_all_safe_requests`** (CacheAllSafeRequests, `bool`)



**Field: `cache_response_codes`** (CacheOnlyResponseCodes, `[]int`)



**Field: `enable_upstream_cache_control`** (EnableUpstreamCacheControl, `bool`)



**Field: `cache_control_ttl_header`** (CacheControlTTLHeader, `string`)



**Field: `cache_by_headers`** (CacheByHeaders, `[]string`)



# AuthProviderMeta

**Field: `name`** (Name, [AuthProviderCode](#AuthProviderCode))



**Field: `storage_engine`** (StorageEngine, [StorageEngineCode](#StorageEngineCode))



**Field: `meta`** (Meta, `map[string]interface{}`)



# SessionProviderMeta

**Field: `name`** (Name, [SessionProviderCode](#SessionProviderCode))



**Field: `storage_engine`** (StorageEngine, [StorageEngineCode](#StorageEngineCode))



**Field: `meta`** (Meta, `map[string]interface{}`)



# EventHandlerMetaConfig

**Field: `events`** (Events, `map[TykEvent]interface{}`)



# ResponseProcessor

**Field: `name`** (Name, `string`)



**Field: `options`** (Options, ``)



# CORSConfig

**Field: `enable`** (Enable, `bool`)



**Field: `allowed_origins`** (AllowedOrigins, `[]string`)



**Field: `allowed_methods`** (AllowedMethods, `[]string`)



**Field: `allowed_headers`** (AllowedHeaders, `[]string`)



**Field: `exposed_headers`** (ExposedHeaders, `[]string`)



**Field: `allow_credentials`** (AllowCredentials, `bool`)



**Field: `max_age`** (MaxAge, `int`)



**Field: `options_passthrough`** (OptionsPassthrough, `bool`)



**Field: `debug`** (Debug, `bool`)



# GlobalRateLimit

**Field: `rate`** (Rate, `float64`)



**Field: `per`** (Per, `float64`)



# GraphQLConfig

GraphQLConfig is the root config object for a GraphQL API.

**Field: `enabled`** (Enabled, `bool`)

Enabled indicates if GraphQL should be enabled.

**Field: `execution_mode`** (ExecutionMode, [GraphQLExecutionMode](#GraphQLExecutionMode))

ExecutionMode is the mode to define how an api behaves.

**Field: `version`** (Version, [GraphQLConfigVersion](#GraphQLConfigVersion))

Version defines the version of the GraphQL config and engine to be used.

**Field: `schema`** (Schema, `string`)

Schema is the GraphQL Schema exposed by the GraphQL API/Upstream/Engine.

**Field: `last_schema_update`** (LastSchemaUpdate, `time.Time`)

LastSchemaUpdate contains the date and time of the last triggered schema update to the upstream.

**Field: `type_field_configurations`** (TypeFieldConfigurations, `[]datasource.TypeFieldConfiguration`)

TypeFieldConfigurations is a rule set of data source and mapping of a schema field.

**Field: `playground`** (GraphQLPlayground, [GraphQLPlayground](#GraphQLPlayground))

GraphQLPlayground is the Playground specific configuration.

**Field: `engine`** (Engine, [GraphQLEngineConfig](#GraphQLEngineConfig))

Engine holds the configuration for engine v2 and upwards.

**Field: `proxy`** (Proxy, [GraphQLProxyConfig](#GraphQLProxyConfig))

Proxy holds the configuration for a proxy only api.

**Field: `subgraph`** (Subgraph, [GraphQLSubgraphConfig](#GraphQLSubgraphConfig))

Subgraph holds the configuration for a GraphQL federation subgraph.

**Field: `supergraph`** (Supergraph, [GraphQLSupergraphConfig](#GraphQLSupergraphConfig))

Supergraph holds the configuration for a GraphQL federation supergraph.

# AnalyticsPluginConfig

**Field: `enable`** (Enabled, `bool`)



**Field: `plugin_path`** (PluginPath, `string`)



**Field: `func_name`** (FuncName, `string`)



# Provider

**Field: `jwt`** (JWT, [JWTValidation](#JWTValidation))



**Field: `introspection`** (Introspection, [Introspection](#Introspection))



# OIDProviderConfig

**Field: `issuer`** (Issuer, `string`)



**Field: `client_ids`** (ClientIDs, `map[string]string`)



# SignatureConfig

**Field: `algorithm`** (Algorithm, `string`)



**Field: `header`** (Header, `string`)



**Field: `use_param`** (UseParam, `bool`)



**Field: `param_name`** (ParamName, `string`)



**Field: `secret`** (Secret, `string`)



**Field: `allowed_clock_skew`** (AllowedClockSkew, `int64`)



**Field: `error_code`** (ErrorCode, `int`)



**Field: `error_message`** (ErrorMessage, `string`)



# ScopeClaim

**Field: `scope_claim_name`** (ScopeClaimName, `string`)



**Field: `scope_to_policy`** (ScopeToPolicy, `map[string]string`)



# ScopeClaim

**Field: `scope_claim_name`** (ScopeClaimName, `string`)



**Field: `scope_to_policy`** (ScopeToPolicy, `map[string]string`)



# HostCheckObject

**Field: `url`** (CheckURL, `string`)



**Field: `protocol`** (Protocol, `string`)



**Field: `timeout`** (Timeout, `time.Duration`)



**Field: `enable_proxy_protocol`** (EnableProxyProtocol, `bool`)



**Field: `commands`** (Commands, [[]CheckCommand](#CheckCommand))



**Field: `method`** (Method, `string`)



**Field: `headers`** (Headers, `map[string]string`)



**Field: `body`** (Body, `string`)



# UptimeTestsConfig

**Field: `expire_utime_after`** (ExpireUptimeAnalyticsAfter, `int64`)



> must have an expireAt TTL index set (http://docs.mongodb.org/manual/tutorial/expire-data/)

**Field: `service_discovery`** (ServiceDiscovery, [ServiceDiscoveryConfiguration](#ServiceDiscoveryConfiguration))



**Field: `recheck_wait`** (RecheckWait, `int`)



# ServiceDiscoveryConfiguration

**Field: `use_discovery_service`** (UseDiscoveryService, `bool`)



**Field: `query_endpoint`** (QueryEndpoint, `string`)



**Field: `use_nested_query`** (UseNestedQuery, `bool`)



**Field: `parent_data_path`** (ParentDataPath, `string`)



**Field: `data_path`** (DataPath, `string`)



**Field: `port_data_path`** (PortDataPath, `string`)



**Field: `target_path`** (TargetPath, `string`)



**Field: `use_target_list`** (UseTargetList, `bool`)



**Field: `cache_disabled`** (CacheDisabled, `bool`)



**Field: `cache_timeout`** (CacheTimeout, `int64`)



**Field: `endpoint_returns_list`** (EndpointReturnsList, `bool`)



# MiddlewareDefinition

**Field: `disabled`** (Disabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `path`** (Path, `string`)



**Field: `require_session`** (RequireSession, `bool`)



**Field: `raw_body_only`** (RawBodyOnly, `bool`)



# MiddlewareDefinition

**Field: `disabled`** (Disabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `path`** (Path, `string`)



**Field: `require_session`** (RequireSession, `bool`)



**Field: `raw_body_only`** (RawBodyOnly, `bool`)



# MiddlewareDefinition

**Field: `disabled`** (Disabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `path`** (Path, `string`)



**Field: `require_session`** (RequireSession, `bool`)



**Field: `raw_body_only`** (RawBodyOnly, `bool`)



# MiddlewareDefinition

**Field: `disabled`** (Disabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `path`** (Path, `string`)



**Field: `require_session`** (RequireSession, `bool`)



**Field: `raw_body_only`** (RawBodyOnly, `bool`)



# MiddlewareDefinition

**Field: `disabled`** (Disabled, `bool`)



**Field: `name`** (Name, `string`)



**Field: `path`** (Path, `string`)



**Field: `require_session`** (RequireSession, `bool`)



**Field: `raw_body_only`** (RawBodyOnly, `bool`)



# MiddlewareDriver

Type defined as `string`, see [string](string) definition.

# MiddlewareIdExtractor

**Field: `disabled`** (Disabled, `bool`)



**Field: `extract_from`** (ExtractFrom, [IdExtractorSource](#IdExtractorSource))



**Field: `extract_with`** (ExtractWith, [IdExtractorType](#IdExtractorType))



**Field: `extractor_config`** (ExtractorConfig, `map[string]interface{}`)



# AuthProviderCode

Type defined as `string`, see [string](string) definition.

# StorageEngineCode

Type defined as `string`, see [string](string) definition.

# SessionProviderCode

Type defined as `string`, see [string](string) definition.

# GraphQLExecutionMode

GraphQLExecutionMode is the mode in which the GraphQL Middleware should operate.

Type defined as `string`, see [string](string) definition.

# GraphQLConfigVersion

Type defined as `string`, see [string](string) definition.

# GraphQLPlayground

GraphQLPlayground represents the configuration for the public playground which will be hosted alongside the api.

**Field: `enabled`** (Enabled, `bool`)

Enabled indicates if the playground should be enabled.

**Field: `path`** (Path, `string`)

Path sets the path on which the playground will be hosted if enabled.

# GraphQLEngineConfig

**Field: `field_configs`** (FieldConfigs, [[]GraphQLFieldConfig](#GraphQLFieldConfig))



**Field: `data_sources`** (DataSources, [[]GraphQLEngineDataSource](#GraphQLEngineDataSource))



# GraphQLProxyConfig

**Field: `auth_headers`** (AuthHeaders, `map[string]string`)



**Field: `subscription_type`** (SubscriptionType, [SubscriptionType](#SubscriptionType))



**Field: `request_headers`** (RequestHeaders, `map[string]string`)



**Field: `use_response_extensions`** (UseResponseExtensions, [GraphQLResponseExtensions](#GraphQLResponseExtensions))



# GraphQLSubgraphConfig

**Field: `sdl`** (SDL, `string`)



# GraphQLSupergraphConfig

**Field: `updated_at`** (UpdatedAt, `time.Time`)

UpdatedAt contains the date and time of the last update of a supergraph API.

**Field: `subgraphs`** (Subgraphs, [[]GraphQLSubgraphEntity](#GraphQLSubgraphEntity))



**Field: `merged_sdl`** (MergedSDL, `string`)



**Field: `global_headers`** (GlobalHeaders, `map[string]string`)



**Field: `disable_query_batching`** (DisableQueryBatching, `bool`)



# JWTValidation

**Field: `enabled`** (Enabled, `bool`)



**Field: `signing_method`** (SigningMethod, `string`)



**Field: `source`** (Source, `string`)



**Field: `issued_at_validation_skew`** (IssuedAtValidationSkew, `uint64`)



**Field: `not_before_validation_skew`** (NotBeforeValidationSkew, `uint64`)



**Field: `expires_at_validation_skew`** (ExpiresAtValidationSkew, `uint64`)



**Field: `identity_base_field`** (IdentityBaseField, `string`)



# Introspection

**Field: `enabled`** (Enabled, `bool`)



**Field: `url`** (URL, `string`)



**Field: `client_id`** (ClientID, `string`)



**Field: `client_secret`** (ClientSecret, `string`)



**Field: `identity_base_field`** (IdentityBaseField, `string`)



**Field: `cache`** (Cache, [IntrospectionCache](#IntrospectionCache))



# CheckCommand

**Field: `name`** (Name, `string`)



**Field: `message`** (Message, `string`)



# IdExtractorSource

Type defined as `string`, see [string](string) definition.

# IdExtractorType

Type defined as `string`, see [string](string) definition.

# GraphQLFieldConfig

**Field: `type_name`** (TypeName, `string`)



**Field: `field_name`** (FieldName, `string`)



**Field: `disable_default_mapping`** (DisableDefaultMapping, `bool`)



**Field: `path`** (Path, `[]string`)



# GraphQLEngineDataSource

**Field: `kind`** (Kind, [GraphQLEngineDataSourceKind](#GraphQLEngineDataSourceKind))



**Field: `name`** (Name, `string`)



**Field: `internal`** (Internal, `bool`)



**Field: `root_fields`** (RootFields, [[]GraphQLTypeFields](#GraphQLTypeFields))



**Field: `config`** (Config, `json.RawMessage`)



# SubscriptionType

Type defined as `string`, see [string](string) definition.

# GraphQLResponseExtensions

**Field: `on_error_forwarding`** (OnErrorForwarding, `bool`)



# GraphQLSubgraphEntity

**Field: `api_id`** (APIID, `string`)



**Field: `name`** (Name, `string`)



**Field: `url`** (URL, `string`)



**Field: `sdl`** (SDL, `string`)



**Field: `headers`** (Headers, `map[string]string`)



**Field: `subscription_type`** (SubscriptionType, [SubscriptionType](#SubscriptionType))



# IntrospectionCache

**Field: `enabled`** (Enabled, `bool`)



**Field: `timeout`** (Timeout, `int64`)



# GraphQLEngineDataSourceKind

Type defined as `string`, see [string](string) definition.

# GraphQLTypeFields

**Field: `type`** (Type, `string`)



**Field: `fields`** (Fields, `[]string`)



# InboundData

**Field: `KeyName`** (KeyName, `string`)



**Field: `Value`** (Value, `string`)



**Field: `SessionState`** (SessionState, `string`)



**Field: `Timeout`** (Timeout, `int64`)



**Field: `Per`** (Per, `int64`)



**Field: `Expire`** (Expire, `int64`)



# DefRequest

**Field: `OrgId`** (OrgId, `string`)



**Field: `Tags`** (Tags, `[]string`)



**Field: `LoadOAS`** (LoadOAS, `bool`)



# GroupLoginRequest

**Field: `UserKey`** (UserKey, `string`)



**Field: `GroupID`** (GroupID, `string`)



**Field: `ForceSync`** (ForceSync, `bool`)



# GroupKeySpaceRequest

**Field: `OrgID`** (OrgID, `string`)



**Field: `GroupID`** (GroupID, `string`)



# KeysValuesPair

**Field: `Keys`** (Keys, `[]string`)



**Field: `Values`** (Values, `[]string`)



# ValidationResult

**Field: `IsValid`** (IsValid, `bool`)



**Field: `Errors`** (Errors, `[]error`)



# ValidationRuleSet

Type defined as `[]ValidationRule`, see [ValidationRule](ValidationRule) definition.

# ValidationRule

Type defined as ``, see []() definition.

# RuleUniqueDataSourceNames

Type defined as ``, see []() definition.

# RuleAtLeastEnableOneAuthSource

Type defined as ``, see []() definition.

# RuleValidateIPList

Type defined as ``, see []() definition.

# TykEvent

Type defined as `string`, see [string](string) definition.

# TykEventHandlerName

Type defined as `string`, see [string](string) definition.

# EndpointMethodAction

Type defined as `string`, see [string](string) definition.

# SourceMode

Type defined as `string`, see [string](string) definition.

# RoutingTriggerOnType

Type defined as `string`, see [string](string) definition.

# IDExtractor

Type defined as ``, see []() definition.

# EndpointMethodMeta

**Field: `action`** (Action, [EndpointMethodAction](#EndpointMethodAction))



**Field: `code`** (Code, `int`)



**Field: `data`** (Data, `string`)



**Field: `headers`** (Headers, `map[string]string`)



# MockResponseMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `ignore_case`** (IgnoreCase, `bool`)



**Field: `code`** (Code, `int`)



**Field: `body`** (Body, `string`)



**Field: `headers`** (Headers, `map[string]string`)



# EndPointMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `ignore_case`** (IgnoreCase, `bool`)



**Field: `method_actions`** (MethodActions, `map[string]EndpointMethodMeta`)

Deprecated. Use Method instead.

# CacheMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `method`** (Method, `string`)



**Field: `path`** (Path, `string`)



**Field: `cache_key_regex`** (CacheKeyRegex, `string`)



**Field: `cache_response_codes`** (CacheOnlyResponseCodes, `[]int`)



# RequestInputType

Type defined as `string`, see [string](string) definition.

# TemplateData

**Field: `input_type`** (Input, [RequestInputType](#RequestInputType))



**Field: `template_mode`** (Mode, [SourceMode](#SourceMode))



**Field: `enable_session`** (EnableSession, `bool`)



**Field: `template_source`** (TemplateSource, `string`)



# TemplateMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `template_data`** (TemplateData, [TemplateData](#TemplateData))



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



# TransformJQMeta

**Field: `filter`** (Filter, `string`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



# HeaderInjectionMeta

**Field: `delete_headers`** (DeleteHeaders, `[]string`)



**Field: `add_headers`** (AddHeaders, `map[string]string`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `act_on`** (ActOnResponse, `bool`)



# HardTimeoutMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `timeout`** (TimeOut, `int`)



# TrackEndpointMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



# InternalMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



# RequestSizeMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `size_limit`** (SizeLimit, `int64`)



# CircuitBreakerMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `threshold_percent`** (ThresholdPercent, `float64`)



**Field: `samples`** (Samples, `int64`)



**Field: `return_to_service_after`** (ReturnToServiceAfter, `int`)



**Field: `disable_half_open_state`** (DisableHalfOpenState, `bool`)



# StringRegexMap

**Field: `match_rx`** (MatchPattern, `string`)



**Field: `reverse`** (Reverse, `bool`)



# RoutingTriggerOptions

**Field: `header_matches`** (HeaderMatches, `map[string]StringRegexMap`)



**Field: `query_val_matches`** (QueryValMatches, `map[string]StringRegexMap`)



**Field: `path_part_matches`** (PathPartMatches, `map[string]StringRegexMap`)



**Field: `session_meta_matches`** (SessionMetaMatches, `map[string]StringRegexMap`)



**Field: `request_context_matches`** (RequestContextMatches, `map[string]StringRegexMap`)



**Field: `payload_matches`** (PayloadMatches, [StringRegexMap](#StringRegexMap))



# RoutingTrigger

**Field: `on`** (On, [RoutingTriggerOnType](#RoutingTriggerOnType))



**Field: `options`** (Options, [RoutingTriggerOptions](#RoutingTriggerOptions))



**Field: `rewrite_to`** (RewriteTo, `string`)



# URLRewriteMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `match_pattern`** (MatchPattern, `string`)



**Field: `rewrite_to`** (RewriteTo, `string`)



**Field: `triggers`** (Triggers, [[]RoutingTrigger](#RoutingTrigger))



# VirtualMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `response_function_name`** (ResponseFunctionName, `string`)



**Field: `function_source_type`** (FunctionSourceType, [SourceMode](#SourceMode))



**Field: `function_source_uri`** (FunctionSourceURI, `string`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `use_session`** (UseSession, `bool`)



**Field: `proxy_on_error`** (ProxyOnError, `bool`)



# MethodTransformMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `to_method`** (ToMethod, `string`)



# ValidatePathMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `schema`** (Schema, `map[string]interface{}`)



**Field: `schema_b64`** (SchemaB64, `string`)



**Field: `error_response_code`** (ErrorResponseCode, `int`)

Allows override of default 422 Unprocessible Entity response code for validation errors.

# ValidateRequestMeta

**Field: `enabled`** (Enabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `error_response_code`** (ErrorResponseCode, `int`)

Allows override of default 422 Unprocessible Entity response code for validation errors.

# PersistGraphQLMeta

**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `operation`** (Operation, `string`)



**Field: `variables`** (Variables, `map[string]interface{}`)



# GoPluginMeta

**Field: `disabled`** (Disabled, `bool`)



**Field: `path`** (Path, `string`)



**Field: `method`** (Method, `string`)



**Field: `plugin_path`** (PluginPath, `string`)



**Field: `func_name`** (SymbolName, `string`)



# ExtendedPathsSet

**Field: `ignored`** (Ignored, [[]EndPointMeta](#EndPointMeta))



**Field: `white_list`** (WhiteList, [[]EndPointMeta](#EndPointMeta))



**Field: `black_list`** (BlackList, [[]EndPointMeta](#EndPointMeta))



**Field: `mock_response`** (MockResponse, [[]MockResponseMeta](#MockResponseMeta))



**Field: `cache`** (Cached, `[]string`)



**Field: `advance_cache_config`** (AdvanceCacheConfig, [[]CacheMeta](#CacheMeta))



**Field: `transform`** (Transform, [[]TemplateMeta](#TemplateMeta))



**Field: `transform_response`** (TransformResponse, [[]TemplateMeta](#TemplateMeta))



**Field: `transform_jq`** (TransformJQ, [[]TransformJQMeta](#TransformJQMeta))



**Field: `transform_jq_response`** (TransformJQResponse, [[]TransformJQMeta](#TransformJQMeta))



**Field: `transform_headers`** (TransformHeader, [[]HeaderInjectionMeta](#HeaderInjectionMeta))



**Field: `transform_response_headers`** (TransformResponseHeader, [[]HeaderInjectionMeta](#HeaderInjectionMeta))



**Field: `hard_timeouts`** (HardTimeouts, [[]HardTimeoutMeta](#HardTimeoutMeta))



**Field: `circuit_breakers`** (CircuitBreaker, [[]CircuitBreakerMeta](#CircuitBreakerMeta))



**Field: `url_rewrites`** (URLRewrite, [[]URLRewriteMeta](#URLRewriteMeta))



**Field: `virtual`** (Virtual, [[]VirtualMeta](#VirtualMeta))



**Field: `size_limits`** (SizeLimit, [[]RequestSizeMeta](#RequestSizeMeta))



**Field: `method_transforms`** (MethodTransforms, [[]MethodTransformMeta](#MethodTransformMeta))



**Field: `track_endpoints`** (TrackEndpoints, [[]TrackEndpointMeta](#TrackEndpointMeta))



**Field: `do_not_track_endpoints`** (DoNotTrackEndpoints, [[]TrackEndpointMeta](#TrackEndpointMeta))



**Field: `validate_json`** (ValidateJSON, [[]ValidatePathMeta](#ValidatePathMeta))



**Field: `validate_request`** (ValidateRequest, [[]ValidateRequestMeta](#ValidateRequestMeta))



**Field: `internal`** (Internal, [[]InternalMeta](#InternalMeta))



**Field: `go_plugin`** (GoPlugin, [[]GoPluginMeta](#GoPluginMeta))



**Field: `persist_graphql`** (PersistGraphQL, [[]PersistGraphQLMeta](#PersistGraphQLMeta))



# VersionInfo

**Field: `name`** (Name, `string`)



**Field: `expires`** (Expires, `string`)



**Field: `paths`** (Paths, `struct{}`)



**Field: `use_extended_paths`** (UseExtendedPaths, `bool`)



**Field: `extended_paths`** (ExtendedPaths, [ExtendedPathsSet](#ExtendedPathsSet))



**Field: `global_headers`** (GlobalHeaders, `map[string]string`)



**Field: `global_headers_remove`** (GlobalHeadersRemove, `[]string`)



**Field: `global_response_headers`** (GlobalResponseHeaders, `map[string]string`)



**Field: `global_response_headers_remove`** (GlobalResponseHeadersRemove, `[]string`)



**Field: `ignore_endpoint_case`** (IgnoreEndpointCase, `bool`)



**Field: `global_size_limit`** (GlobalSizeLimit, `int64`)



**Field: `override_target`** (OverrideTarget, `string`)



# EventHandlerTriggerConfig

**Field: `handler_name`** (Handler, [TykEventHandlerName](#TykEventHandlerName))



**Field: `handler_meta`** (HandlerMeta, `map[string]interface{}`)



# BundleManifest

**Field: `file_list`** (FileList, `[]string`)



**Field: `custom_middleware`** (CustomMiddleware, [MiddlewareSection](#MiddlewareSection))



**Field: `checksum`** (Checksum, `string`)



**Field: `signature`** (Signature, `string`)



# GraphQLEngineDataSourceConfigREST

**Field: `url`** (URL, `string`)



**Field: `method`** (Method, `string`)



**Field: `headers`** (Headers, `map[string]string`)



**Field: `query`** (Query, [[]QueryVariable](#QueryVariable))



**Field: `body`** (Body, `string`)



# GraphQLEngineDataSourceConfigGraphQL

**Field: `url`** (URL, `string`)



**Field: `method`** (Method, `string`)



**Field: `headers`** (Headers, `map[string]string`)



**Field: `subscription_type`** (SubscriptionType, [SubscriptionType](#SubscriptionType))



**Field: `has_operation`** (HasOperation, `bool`)



**Field: `operation`** (Operation, `string`)



**Field: `variables`** (Variables, `json.RawMessage`)



# GraphQLEngineDataSourceConfigKafka

**Field: `broker_addresses`** (BrokerAddresses, `[]string`)



**Field: `topics`** (Topics, `[]string`)



**Field: `group_id`** (GroupID, `string`)



**Field: `client_id`** (ClientID, `string`)



**Field: `kafka_version`** (KafkaVersion, `string`)



**Field: `start_consuming_latest`** (StartConsumingLatest, `bool`)



**Field: `balance_strategy`** (BalanceStrategy, `string`)



**Field: `isolation_level`** (IsolationLevel, `string`)



**Field: `sasl`** (SASL, `kafka_datasource.SASL`)



# QueryVariable

**Field: `name`** (Name, `string`)



**Field: `value`** (Value, `string`)



# HostList

Type defined as ``, see []() definition.

# IDExtractorConfig

IDExtractorConfig specifies the configuration for ID extractor

**Field: `header_name`** (HeaderName, `string`)

HeaderName is the header name to extract ID from.

**Field: `param_name`** (FormParamName, `string`)

FormParamName is the form parameter name to extract ID from.

**Field: `regex_expression`** (RegexExpression, `string`)

RegexExpression is the regular expression to match ID.

**Field: `regex_match_index`** (RegexMatchIndex, `int`)

RegexMatchIndex is the index from which ID to be extracted after a match.

**Field: `xpath_expression`** (XPathExpression, `string`)

XPathExp is the xpath expression to match ID.

