# XTykAPIGateway

XTykAPIGateway contains custom Tyk API extensions for the OAS definition.

**Info** (JSON: `info`)

Info contains the main metadata about the API definition.

> required

**Upstream** (JSON: `upstream`)

Upstream contains the configurations related to the upstream.

> required

**Server** (JSON: `server`)

Server contains the configurations related to the server.

> required

**Middleware** (JSON: `middleware`)

Middleware contains the configurations related to the proxy middleware.

# Info

Info contains the main metadata about the API definition.

**ID** (JSON: `id`)

ID is the unique ID of the API.
Tyk classic API definition: `api_id`

**DBID** (JSON: `dbId`)

DBID is the unique database ID of the API.
Tyk classic API definition: `id`

**OrgID** (JSON: `orgId`)

OrgID is the ID of the organisation which the API belongs to.
Tyk classic API definition: `org_id`

**Name** (JSON: `name`)

Name is the name of the API.
Tyk classic API definition: `name`

> required

**Expiration** (JSON: `expiration`)

Expiration date.

**State** (JSON: `state`)

State holds configuration about API definition states (active, internal).

> required

**Versioning** (JSON: `versioning`)

Versioning holds configuration for API versioning.

# Upstream

Upstream holds configuration for an upstream server.

**URL** (JSON: `url`)

URL defines the target URL that the request should be proxied to.
Tyk classic API definition: `proxy.target_url`

> required

**ServiceDiscovery** (JSON: `serviceDiscovery`)

ServiceDiscovery contains the configuration related to Service Discovery.
Tyk classic API definition: `proxy.service_discovery`

**Test** (JSON: `test`)

Test contains the configuration related to uptime tests.

**MutualTLS** (JSON: `mutualTLS`)

MutualTLS contains the configuration related to upstream mutual TLS.

**CertificatePinning** (JSON: `certificatePinning`)

CertificatePinning contains the configuration related to certificate pinning.

# Server

Server contains the configuration related to the OAS API definition.

**ListenPath** (JSON: `listenPath`)

ListenPath represents the path to listen on. Any requests coming into the host, on the port that Tyk is configured to run on,
that match this path will have the rules defined in the API definition applied.

> required

**Slug** (JSON: `slug`)

Slug is the Tyk Cloud equivalent of listen path.
Tyk classic API definition: `slug`

**Authentication** (JSON: `authentication`)

Authentication contains the configurations related to authentication to the API.

**ClientCertificates** (JSON: `clientCertificates`)

ClientCertificates contains the configurations related to static mTLS.

**GatewayTags** (JSON: `gatewayTags`)

GatewayTags contains segment tags to configure which GWs your APIs connect to.

**CustomDomain** (JSON: `customDomain`)

CustomDomain is the domain to bind this API to.

Tyk classic API definition: `domain`

# Middleware

Middleware holds configuration for middleware.

**Global** (JSON: `global`)

Global contains the configurations related to the global middleware.

**Operations** (JSON: `operations`)

Operations configuration.

# EndpointPostPlugins

No exposed fields available.

# OAuthProvider

**JWT** (JSON: `jwt`)



**Introspection** (JSON: `introspection`)



# JWTValidation

**Enabled** (JSON: `enabled`)

Enabled enables OAuth access token validation by introspection to a third party.

**SigningMethod** (JSON: `signingMethod`)

SigningMethod to verify signing method used in jwt - allowed values HMAC/RSA/ECDSA.

**Source** (JSON: `source`)

Source is the secret to verify signature, it could be one among:
- a base64 encoded static secret,
- a valid JWK url in plain text,
- a valid JWK url in base64 encoded format.

**IdentityBaseField** (JSON: `identityBaseField`)

IdentityBaseField is the identity claim name.

**IssuedAtValidationSkew** (JSON: `issuedAtValidationSkew`)

IssuedAtValidationSkew is the clock skew to be considered while validating iat claim.

**NotBeforeValidationSkew** (JSON: `notBeforeValidationSkew`)

NotBeforeValidationSkew is the clock skew to be considered while validating nbf claim.

**ExpiresAtValidationSkew** (JSON: `expiresAtValidationSkew`)

ExpiresAtValidationSkew is the clock skew to be considered while validating exp claim.

# Introspection

**Enabled** (JSON: `enabled`)

Enabled enables OAuth access token validation by introspection to a third party.

**URL** (JSON: `url`)

URL is the URL of the third party provider's introspection endpoint.

**ClientID** (JSON: `clientId`)

ClientID is the public identifier for the client, acquired from the third party.

**ClientSecret** (JSON: `clientSecret`)

ClientSecret is a secret known only to the client and the authorization server, acquired from the third party.

**IdentityBaseField** (JSON: `identityBaseField`)

IdentityBaseField is the key showing where to find the user id in the claims. If it is empty, the `sub` key is looked at.

**Cache** (JSON: `cache`)

Cache is the caching mechanism for introspection responses.

# IntrospectionCache

**Enabled** (JSON: `enabled`)

Enabled enables the caching mechanism for introspection responses.

**Timeout** (JSON: `timeout`)

Timeout is the duration in seconds of how long the cached value stays.
For introspection caching, it is suggested to use a short interval.

# ExternalOAuth

**Enabled** (JSON: `enabled`)



> required

**Providers** (JSON: `providers`)



> required

# APIDef

APIDef is struct to hold both OAS and Classic forms of an API definition.

**OAS** (JSON: `OAS`)



**Classic** (JSON: `Classic`)



# Allowance

Allowance describes allowance actions and behaviour.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag, if set to `true`, then individual allowances (allow, block, ignore) will be enforced.

**IgnoreCase** (JSON: `ignoreCase`)

IgnoreCase is a boolean flag, If set to `true`, checks for requests allowance will be case insensitive.

# AllowanceType

AllowanceType holds the valid allowance types values.

No exposed fields available.

# AuthSource

AuthSource defines an authentication source.

**Enabled** (JSON: `enabled`)

Enabled enables the auth source.
Tyk classic API definition: `auth_configs[X].use_param/use_cookie`

> required

**Name** (JSON: `name`)

Name is the name of the auth source.
Tyk classic API definition: `auth_configs[X].param_name/cookie_name`

# AuthSources

AuthSources defines authentication source configuration: headers, cookies and query parameters.

Tyk classic API definition: `auth_configs{}`.

**Header** (JSON: `header`)

Header contains configurations for the header value auth source, it is enabled by default.

Tyk classic API definition: `auth_configs[x].header`

**Cookie** (JSON: `cookie`)

Cookie contains configurations for the cookie value auth source.

Tyk classic API definition: `auth_configs[x].cookie`

**Query** (JSON: `query`)

Query contains configurations for the query parameters auth source.

Tyk classic API definition: `auth_configs[x].query`

# Authentication

Authentication types contains configuration about the authentication methods and security policies applied to requests.

**Enabled** (JSON: `enabled`)

Enabled makes the API protected when one of the authentication modes is enabled.

Tyk classic API definition: `!use_keyless`.

> required

**StripAuthorizationData** (JSON: `stripAuthorizationData`)

StripAuthorizationData ensures that any security tokens used for accessing APIs are stripped and not leaked to the upstream.

Tyk classic API definition: `strip_auth_data`.

**BaseIdentityProvider** (JSON: `baseIdentityProvider`)

BaseIdentityProvider enables multi authentication mechanism and provides the session object that determines rate limits, ACL rules and quotas.
It should be set to one of the following:

- `auth_token`
- `hmac_key`
- `basic_auth_user`
- `jwt_claim`
- `oidc_user`
- `oauth_key`
- `custom_auth`

Tyk classic API definition: `base_identity_provided_by`.

**HMAC** (JSON: `hmac`)

HMAC contains the configurations related to HMAC authentication mode.

Tyk classic API definition: `auth_configs["hmac"]`

**OIDC** (JSON: `oidc`)

OIDC contains the configurations related to OIDC authentication mode.

Tyk classic API definition: `auth_configs["oidc"]`

**Custom** (JSON: `custom`)

Custom contains the configurations related to Custom authentication mode.

Tyk classic API definition: `auth_configs["coprocess"]`

**SecuritySchemes** (JSON: `securitySchemes`)

SecuritySchemes contains security schemes definitions.

# AuthenticationPlugin

AuthenticationPlugin holds the configuration for custom authentication plugin.

**Enabled** (JSON: `enabled`)

Enabled enables custom authentication plugin.

> required.

**FunctionName** (JSON: `functionName`)

FunctionName is the name of authentication method.

> required.

**Path** (JSON: `path`)

Path is the path to shared object file in case of gopluign mode or path to js code in case of otto auth plugin.

**RawBodyOnly** (JSON: `rawBodyOnly`)

RawBodyOnly if set to true, do not fill body in request or response object.

**IDExtractor** (JSON: `idExtractor`)

IDExtractor configures ID extractor with coprocess custom authentication.

# Basic

Basic type holds configuration values related to http basic authentication.

**Enabled** (JSON: `enabled`)

Enabled enables the basic authentication mode.
Tyk classic API definition: `use_basic_auth`

> required

**DisableCaching** (JSON: `disableCaching`)

DisableCaching disables the caching of basic authentication key.
Tyk classic API definition: `basic_auth.disable_caching`

**CacheTTL** (JSON: `cacheTTL`)

CacheTTL is the TTL for a cached basic authentication key in seconds.
Tyk classic API definition: `basic_auth.cache_ttl`

**ExtractCredentialsFromBody** (JSON: `extractCredentialsFromBody`)

ExtractCredentialsFromBody helps to extract username and password from body. In some cases, like dealing with SOAP,
user credentials can be passed via request body.

# CORS

CORS holds configuration for cross-origin resource sharing.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag, if set to `true`, this option enables CORS processing.

Tyk classic API definition: `CORS.enable`.

> required

**MaxAge** (JSON: `maxAge`)

MaxAge indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.

Tyk classic API definition: `CORS.max_age`.

**AllowCredentials** (JSON: `allowCredentials`)

AllowCredentials indicates whether the request can include user credentials like cookies,
HTTP authentication or client side SSL certificates.

Tyk classic API definition: `CORS.allow_credentials`.

**ExposedHeaders** (JSON: `exposedHeaders`)

ExposedHeaders indicates which headers are safe to expose to the API of a CORS API specification.

Tyk classic API definition: `CORS.exposed_headers`.

**AllowedHeaders** (JSON: `allowedHeaders`)

AllowedHeaders holds a list of non simple headers the client is allowed to use with cross-domain requests.

Tyk classic API definition: `CORS.allowed_headers`.

**OptionsPassthrough** (JSON: `optionsPassthrough`)

OptionsPassthrough is a boolean flag. If set to `true`, it will proxy the CORS OPTIONS pre-flight
request directly to upstream, without authentication and any CORS checks. This means that pre-flight
requests generated by web-clients such as SwaggerUI or the Tyk Portal documentation system
will be able to test the API using trial keys.

If your service handles CORS natively, then enable this option.

Tyk classic API definition: `CORS.options_passthrough`.

**Debug** (JSON: `debug`)

Debug is a boolean flag, If set to `true`, this option produces log files for the CORS middleware.

Tyk classic API definition: `CORS.debug`.

**AllowedOrigins** (JSON: `allowedOrigins`)

AllowedOrigins holds a list of origin domains to allow access from. Wildcards are also supported, e.g. `http://*.foo.com`

Tyk classic API definition: `CORS.allowed_origins`.

**AllowedMethods** (JSON: `allowedMethods`)

AllowedMethods holds a list of methods to allow access via.

Tyk classic API definition: `CORS.allowed_methods`.

# Cache

Cache holds configuration for caching the requests.

**Enabled** (JSON: `enabled`)

Enabled turns global cache middleware on or off. It is still possible to enable caching on a per-path basis
by explicitly setting the endpoint cache middleware.

Tyk classic API definition: `cache_options.enable_cache`

> required

**Timeout** (JSON: `timeout`)

Timeout is the TTL for a cached object in seconds.

Tyk classic API definition: `cache_options.cache_timeout`

**CacheAllSafeRequests** (JSON: `cacheAllSafeRequests`)

CacheAllSafeRequests caches responses to (`GET`, `HEAD`, `OPTIONS`) requests overrides per-path cache settings in versions,
applies across versions.

Tyk classic API definition: `cache_options.cache_all_safe_requests`

**CacheResponseCodes** (JSON: `cacheResponseCodes`)

CacheResponseCodes is an array of response codes which are safe to cache e.g. `404`.

Tyk classic API definition: `cache_options.cache_response_codes`

**CacheByHeaders** (JSON: `cacheByHeaders`)

CacheByHeaders allows header values to be used as part of the cache key.

Tyk classic API definition: `cache_options.cache_by_headers`

**EnableUpstreamCacheControl** (JSON: `enableUpstreamCacheControl`)

EnableUpstreamCacheControl instructs Tyk Cache to respect upstream cache control headers.

Tyk classic API definition: `cache_options.enable_upstream_cache_control`

**ControlTTLHeaderName** (JSON: `controlTTLHeaderName`)

ControlTTLHeaderName is the response header which tells Tyk how long it is safe to cache the response for.

Tyk classic API definition: `cache_options.cache_control_ttl_header`

# CachePlugin

CachePlugin holds the configuration for the cache plugins.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag. If set to `true`, the advanced caching plugin will be enabled.

**CacheByRegex** (JSON: `cacheByRegex`)

CacheByRegex defines a regular expression used against the request body to produce a cache key.

Example value: `\"id\":[^,]*` (quoted json value).

**CacheResponseCodes** (JSON: `cacheResponseCodes`)

CacheResponseCodes contains a list of valid response codes for responses that are okay to add to the cache.

# CertificatePinning

CertificatePinning holds the configuration about mapping of domains to pinned public keys.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag, if set to `true`, it enables certificate pinning for the API.

Tyk classic API definition: `certificate_pinning_disabled`

**DomainToPublicKeysMapping** (JSON: `domainToPublicKeysMapping`)

DomainToPublicKeysMapping maintains the mapping of domain to pinned public keys.

Tyk classic API definition: `pinned_public_keys`

# ClientCertificates

ClientCertificates holds a list of client certificates which are allowed to make requests against the server.

**Enabled** (JSON: `enabled`)

Enabled enables static mTLS for the API.

**Allowlist** (JSON: `allowlist`)

Allowlist is the list of client certificates which are allowed.

# ClientToPolicy

ClientToPolicy contains a 1-1 mapping between Client ID and Policy ID.

**ClientID** (JSON: `clientId`)

ClientID contains a Client ID.

**PolicyID** (JSON: `policyId`)

PolicyID contains a Policy ID.

# CustomPlugin

CustomPlugin configures custom plugin.

**Enabled** (JSON: `enabled`)

Enabled enables the custom pre plugin.

> required.

**FunctionName** (JSON: `functionName`)

FunctionName is the name of authentication method.

> required.

**Path** (JSON: `path`)

Path is the path to shared object file in case of gopluign mode or path to js code in case of otto auth plugin.

**RawBodyOnly** (JSON: `rawBodyOnly`)

RawBodyOnly if set to true, do not fill body in request or response object.

**RequireSession** (JSON: `requireSession`)

RequireSession if set to true passes down the session information for plugins after authentication.
RequireSession is used only with JSVM custom middleware.

# CustomPluginAuthentication

CustomPluginAuthentication holds configuration for custom plugins.

**Enabled** (JSON: `enabled`)

Enabled enables the CustomPluginAuthentication authentication mode.

Tyk classic API definition: `enable_coprocess_auth`/`use_go_plugin_auth`.

> required

**Config** (JSON: `config`)

Config contains configuration related to custom authentication plugin.
Tyk classic API definition: `custom_middleware.auth_check`.

# CustomPlugins

CustomPlugins is a list of CustomPlugin.

No exposed fields available.

# Domain

Domain holds the configuration of the domain name the server should listen on.

**Enabled** (JSON: `enabled`)

Enabled allow/disallow the usage of the domain.

**Name** (JSON: `name`)

Name is the name of the domain.

# DomainToCertificate

DomainToCertificate holds a single mapping of domain name into a certificate.

**Domain** (JSON: `domain`)

Domain contains the domain name.

**Certificate** (JSON: `certificate`)

Certificate contains the certificate mapped to the domain.

# EndpointPostPlugin

EndpointPostPlugin contains endpoint level post plugin configuration.

**Enabled** (JSON: `enabled`)

Enabled enables post plugin.

> required.

**Name** (JSON: `name`)

Name is the name of plugin function to be executed.

> required.

**Path** (JSON: `path`)

Path is the path to plugin.

> required.

# EnforceTimeout

EnforceTimeout holds the configuration for enforcing request timeouts.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag. If set to `true`, requests will enforce a configured timeout.

**Value** (JSON: `value`)

Value is the configured timeout in seconds.

# ExtractCredentialsFromBody

ExtractCredentialsFromBody configures extracting credentials from the request body.

**Enabled** (JSON: `enabled`)

Enabled enables extracting credentials from body.
Tyk classic API definition: `basic_auth.extract_from_body`

> required

**UserRegexp** (JSON: `userRegexp`)

UserRegexp is the regex for username e.g. `<User>(.*)</User>`.
Tyk classic API definition: `basic_auth.userRegexp`

**PasswordRegexp** (JSON: `passwordRegexp`)

PasswordRegexp is the regex for password e.g. `<Password>(.*)</Password>`.
Tyk classic API definition: `basic_auth.passwordRegexp`

# FieldDocError

FieldDocError holds a list of errors.

No exposed fields available.

# FieldInfo

FieldInfo holds details about a field.

**Doc** (JSON: `doc`)

Doc is field docs. comments that are not part of docs are excluded.

**JSONName** (JSON: `json_name`)

JSONName is the corresponding json name of the field.

**JSONType** (JSON: `json_type`)

JSONType valid json type if it was found

**GoPath** (JSON: `go_path`)

GoPath is the go path of this field starting from root object

**MapKey** (JSON: `map_key`)

MapKey is the map key type, if this field is a map

**IsArray** (JSON: `is_array`)

IsArray reports if this field is an array.

# FromOASExamples

FromOASExamples configures mock responses should be returned from OAS example responses.

**Enabled** (JSON: `enabled`)

Enabled enables getting a mock response from OAS examples or schemas documented in OAS.

**Code** (JSON: `code`)

Code is the default HTTP response code that the gateway reads from the path responses documented in OAS.

**ContentType** (JSON: `contentType`)

ContentType is the default HTTP response body type that the gateway reads from the path responses documented in OAS.

**ExampleName** (JSON: `exampleName`)

ExampleName is the default example name among multiple path response examples documented in OAS.

# GatewayTags

GatewayTags holds a list of segment tags that should apply for a gateway.

**Enabled** (JSON: `enabled`)

Enabled enables use of segment tags.

**Tags** (JSON: `tags`)

Tags is a list of segment tags

# Global

Global holds configuration applies globally: CORS and caching.

**PluginConfig** (JSON: `pluginConfig`)

PluginConfig contains the configuration related custom plugin bundles/driver.

**CORS** (JSON: `cors`)

CORS contains the configuration related to cross origin resource sharing.
Tyk classic API definition: `CORS`.

**PrePlugin** (JSON: `prePlugin`)

PrePlugin contains configuration related to custom pre-authentication plugin.
Tyk classic API definition: `custom_middleware.pre`.

**PostAuthenticationPlugin** (JSON: `postAuthenticationPlugin`)

PostAuthenticationPlugin contains configuration related to custom post authentication plugin.
Tyk classic API definition: `custom_middleware.post_key_auth`.

**PostPlugin** (JSON: `postPlugin`)

PostPlugin contains configuration related to custom post plugin.
Tyk classic API definition: `custom_middleware.post`.

**ResponsePlugin** (JSON: `responsePlugin`)

ResponsePlugin contains configuration related to custom post plugin.
Tyk classic API definition: `custom_middleware.response`.

**Cache** (JSON: `cache`)

Cache contains the configurations related to caching.
Tyk classic API definition: `cache_options`.

# HMAC

HMAC holds the configuration for the HMAC authentication mode.

**Enabled** (JSON: `enabled`)

Enabled enables the HMAC authentication mode.
Tyk classic API definition: `enable_signature_checking`

> required

**AllowedAlgorithms** (JSON: `allowedAlgorithms`)

AllowedAlgorithms is the array of HMAC algorithms which are allowed. Tyk supports the following HMAC algorithms:

- `hmac-sha1`
- `hmac-sha256`
- `hmac-sha384`
- `hmac-sha512`

and reads the value from algorithm header.

Tyk classic API definition: `hmac_allowed_algorithms`

**AllowedClockSkew** (JSON: `allowedClockSkew`)

AllowedClockSkew is the amount of milliseconds that will be tolerated for clock skew. It is used against replay attacks.
The default value is `0`, which deactivates clock skew checks.
Tyk classic API definition: `hmac_allowed_clock_skew`

# Header

Header holds a header name and value pair.

**Name** (JSON: `name`)

Name is the name of the header.

**Value** (JSON: `value`)

Value is the value of the header.

# IDExtractor

IDExtractor configures ID Extractor.

**Enabled** (JSON: `enabled`)

Enabled enables ID extractor with coprocess authentication.

> required

**Source** (JSON: `source`)

Source is the source from which ID to be extracted from.

> required

**With** (JSON: `with`)

With is the type of ID extractor to be used.

> required

**Config** (JSON: `config`)

Config holds the configuration specific to ID extractor type mentioned via With.

> required

# IDExtractorConfig

IDExtractorConfig specifies the configuration for ID extractor.

**HeaderName** (JSON: `headerName`)

HeaderName is the header name to extract ID from.

**FormParamName** (JSON: `formParamName`)

FormParamName is the form parameter name to extract ID from.

**Regexp** (JSON: `regexp`)

Regexp is the regular expression to match ID.

**RegexpMatchIndex** (JSON: `regexpMatchIndex`)

RegexpMatchIndex is the index from which ID to be extracted after a match.
Default value is 0, ie if regexpMatchIndex is not provided ID is matched from index 0.

**XPathExp** (JSON: `xPathExp`)

XPathExp is the xpath expression to match ID.

# JWT

JWT holds the configuration for the JWT middleware.

**Enabled** (JSON: `enabled`)



> required

**Source** (JSON: `source`)



**SigningMethod** (JSON: `signingMethod`)



**IdentityBaseField** (JSON: `identityBaseField`)



**SkipKid** (JSON: `skipKid`)



**PolicyFieldName** (JSON: `policyFieldName`)



**ClientBaseField** (JSON: `clientBaseField`)



**Scopes** (JSON: `scopes`)



**DefaultPolicies** (JSON: `defaultPolicies`)



**IssuedAtValidationSkew** (JSON: `issuedAtValidationSkew`)



**NotBeforeValidationSkew** (JSON: `notBeforeValidationSkew`)



**ExpiresAtValidationSkew** (JSON: `expiresAtValidationSkew`)



# ListenPath

ListenPath represents the path the server should listen on.

**Value** (JSON: `value`)

Value is the value of the listen path e.g. `/api/` or `/` or `/httpbin/`.
Tyk classic API definition: `proxy.listen_path`

> required

**Strip** (JSON: `strip`)

Strip removes the inbound listen path in the outgoing request. e.g. `http://acme.com/httpbin/get` where `httpbin`
is the listen path. The `httpbin` listen path which is used to identify the API loaded in Tyk is removed,
and the outbound request would be `http://httpbin.org/get`.
Tyk classic API definition: `proxy.strip_listen_path`

# MockResponse

MockResponse configures the mock responses.

**Enabled** (JSON: `enabled`)

Enabled enables the mock response middleware.

**Code** (JSON: `code`)

Code is the HTTP response code that will be returned.

**Body** (JSON: `body`)

Body is the HTTP response body that will be returned.

**Headers** (JSON: `headers`)

Headers are the HTTP response headers that will be returned.

**FromOASExamples** (JSON: `fromOASExamples`)

FromOASExamples is the configuration to extract a mock response from OAS documentation.

# MutualTLS

MutualTLS holds configuration related to mTLS on APIs, domain to certificate mappings.

**Enabled** (JSON: `enabled`)

Enabled enables/disables upstream mutual TLS auth for the API.
Tyk classic API definition: `upstream_certificates_disabled`

**DomainToCertificates** (JSON: `domainToCertificateMapping`)

DomainToCertificates maintains the mapping of domain to certificate.
Tyk classic API definition: `upstream_certificates`

# Notifications

Notifications holds configuration for updates to keys.

**SharedSecret** (JSON: `sharedSecret`)

SharedSecret is the shared secret used in the notification request.

**OnKeyChangeURL** (JSON: `onKeyChangeUrl`)

OnKeyChangeURL is the URL a request will be triggered against.

# OAS

OAS holds the upstream OAS definition as well as adds functionality like custom JSON marshalling.

No exposed fields available.

# OAuth

OAuth configures the OAuth middleware.

**Enabled** (JSON: `enabled`)



> required

**AllowedAuthorizeTypes** (JSON: `allowedAuthorizeTypes`)



**RefreshToken** (JSON: `refreshToken`)



**AuthLoginRedirect** (JSON: `authLoginRedirect`)



**Notifications** (JSON: `notifications`)



# OIDC

OIDC contains configuration for the OIDC authentication mode.

**Enabled** (JSON: `enabled`)

Enabled enables the OIDC authentication mode.

Tyk classic API definition: `use_openid`

> required

**SegregateByClientId** (JSON: `segregateByClientId`)

SegregateByClientId is a boolean flag. If set to `true, the policies will be applied to a combination of Client ID and User ID.

Tyk classic API definition: `openid_options.segregate_by_client`.

**Providers** (JSON: `providers`)

Providers contains a list of authorised providers and their Client IDs, and matched policies.

Tyk classic API definition: `openid_options.providers`.

**Scopes** (JSON: `scopes`)

Scopes contains the defined scope claims.

# Operation

Operation holds a request operation configuration, allowances, tranformations, caching, timeouts and validation.

**Allow** (JSON: `allow`)

Allow request by allowance.

**Block** (JSON: `block`)

Block request by allowance.

**IgnoreAuthentication** (JSON: `ignoreAuthentication`)

IgnoreAuthentication ignores authentication on request by allowance.

**TransformRequestMethod** (JSON: `transformRequestMethod`)

TransformRequestMethod allows you to transform the method of a request.

**TransformRequestBody** (JSON: `transformRequestBody`)

TransformRequestBody allows you to transform request body.
When both `path` and `body` are provided, body would take precedence.

**Cache** (JSON: `cache`)

Cache contains the caching plugin configuration.

**EnforceTimeout** (JSON: `enforceTimeout`)

EnforceTimeout contains the request timeout configuration.

**ValidateRequest** (JSON: `validateRequest`)

ValidateRequest contains the request validation configuration.

**MockResponse** (JSON: `mockResponse`)

MockResponse contains the mock response configuration.

**VirtualEndpoint** (JSON: `virtualEndpoint`)

VirtualEndpoint contains virtual endpoint configuration.

**PostPlugins** (JSON: `postPlugins`)

PostPlugins contains endpoint level post plugins configuration.

# Operations

Operations holds Operation definitions.

No exposed fields available.

# Path

Path holds plugin configurations for HTTP method verbs.

**Delete** (JSON: `DELETE`)



**Get** (JSON: `GET`)



**Head** (JSON: `HEAD`)



**Options** (JSON: `OPTIONS`)



**Patch** (JSON: `PATCH`)



**Post** (JSON: `POST`)



**Put** (JSON: `PUT`)



**Trace** (JSON: `TRACE`)



**Connect** (JSON: `CONNECT`)



# Paths

Paths is a mapping of API endpoints to Path plugin configurations.

No exposed fields available.

# PinnedPublicKey

PinnedPublicKey contains a mapping from the domain name into a list of public keys.

**Domain** (JSON: `domain`)

Domain contains the domain name.

**PublicKeys** (JSON: `publicKeys`)

PublicKeys contains a list of the public keys pinned to the domain name.

# PinnedPublicKeys

PinnedPublicKeys is a list of domains and pinned public keys for them.

No exposed fields available.

# PluginBundle

PluginBundle holds configuration for custom plugins.

**Enabled** (JSON: `enabled`)

Enabled enables the custom plugin bundles.

Tyk classic API definition: `custom_middleware_bundle_disabled`

> required.

**Path** (JSON: `path`)

Path is the path suffix to construct the URL to fetch plugin bundle from.
Path will be suffixed to `bundle_base_url` in gateway config.

> required.

# PluginConfig

PluginConfig holds configuration for custom plugins.

**Driver** (JSON: `driver`)

Driver configures which custom plugin to be used.
It's value should be set to one of the following:

- `otto`,
- `python`,
- `lua`,
- `grpc`,
- `goplugin`.

Tyk classic API definition: `custom_middleware.driver`.

**Bundle** (JSON: `bundle`)

Bundle configures custom plugin bundles.

**Data** (JSON: `data`)

Data configures custom plugin data.

# PluginConfigData

PluginConfigData configures config data for custom plugins.

**Enabled** (JSON: `enabled`)

Enabled enables custom plugin config data.

> required.

**Value** (JSON: `value`)

Value is the value of custom plugin config data.

> required.

# Plugins

Plugins configures common settings for each plugin, allowances, transforms, caching and timeouts.

**Allow** (JSON: `allow`)

Allow request by allowance.

**Block** (JSON: `block`)

Block request by allowance.

**IgnoreAuthentication** (JSON: `ignoreAuthentication`)

Ignore authentication on request by allowance.

**TransformRequestMethod** (JSON: `transformRequestMethod`)

TransformRequestMethod allows you to transform the method of a request.

**Cache** (JSON: `cache`)

Cache allows you to cache the server side response.

**EnforceTimeout** (JSON: `enforcedTimeout`)

EnforceTimeout allows you to configure a request timeout.

# PostAuthenticationPlugin

PostAuthenticationPlugin configures post authentication plugins.

**Plugins** (JSON: `plugins`)

Plugins configures custom plugins to be run on pre authentication stage.
The plugins would be executed in the order of configuration in the list.

# PostPlugin

PostPlugin configures post plugins.

**Plugins** (JSON: `plugins`)

Plugins configures custom plugins to be run on post stage.
The plugins would be executed in the order of configuration in the list.

# PrePlugin

PrePlugin configures pre stage plugins.

**Plugins** (JSON: `plugins`)

Plugins configures custom plugins to be run on pre authentication stage.
The plugins would be executed in the order of configuration in the list.

# Provider

Provider defines an issuer to validate and the Client ID to Policy ID mappings.

**Issuer** (JSON: `issuer`)

Issuer contains a validation value for the issuer claim, usually a domain name e.g. `accounts.google.com` or similar.

**ClientToPolicyMapping** (JSON: `clientToPolicyMapping`)

ClientToPolicyMapping contains mappings of Client IDs to Policy IDs.

# ResponsePlugin

ResponsePlugin configures response plugins.

**Plugins** (JSON: `plugins`)

Plugins configures custom plugins to be run on post stage.
The plugins would be executed in the order of configuration in the list.

# ScopeToPolicy

ScopeToPolicy contains a single scope to policy ID mapping.

**Scope** (JSON: `scope`)

Scope contains the scope name.

**PolicyID** (JSON: `policyId`)

PolicyID contains the Policy ID.

# Scopes

Scopes holds the scope to policy mappings for a claim name.

**ClaimName** (JSON: `claimName`)

ClaimName contains the claim name.

**ScopeToPolicyMapping** (JSON: `scopeToPolicyMapping`)

ScopeToPolicyMapping contains the mappings of scopes to policy IDs.

# SecurityScheme

SecurityScheme defines an Importer interface for security schemes.

No exposed fields available.

# SecuritySchemes

SecuritySchemes holds security scheme values, filled with Import().

No exposed fields available.

# ServiceDiscovery

ServiceDiscovery holds configuration required for service discovery.

**Enabled** (JSON: `enabled`)

Enabled enables Service Discovery.

Tyk classic API definition: `service_discovery.use_discovery_service`

> required

**QueryEndpoint** (JSON: `queryEndpoint`)

QueryEndpoint is the endpoint to call, this would usually be Consul, etcd or Eureka K/V store.
Tyk classic API definition: `service_discovery.query_endpoint`

**DataPath** (JSON: `dataPath`)

DataPath is the namespace of the data path - where exactly in your service response the namespace can be found.
For example, if your service responds with:

```
{
 "action": "get",
 "node": {
   "key": "/services/single",
   "value": "http://httpbin.org:6000",
   "modifiedIndex": 6,
   "createdIndex": 6
 }
}
```

then your namespace would be `node.value`.

Tyk classic API definition: `service_discovery.data_path`

**UseNestedQuery** (JSON: `useNestedQuery`)

UseNestedQuery enables using a combination of `dataPath` and `parentDataPath`.
It is necessary when the data lives within this string-encoded JSON object.

```
{
 "action": "get",
 "node": {
   "key": "/services/single",
   "value": "{"hostname": "http://httpbin.org", "port": "80"}",
   "modifiedIndex": 6,
   "createdIndex": 6
 }
}
```

Tyk classic API definition: `service_discovery.use_nested_query`

**ParentDataPath** (JSON: `parentDataPath`)

ParentDataPath is the namespace of the where to find the nested
value, if `useNestedQuery` is `true`. In the above example, it
would be `node.value`. You would change the `dataPath` setting
to be `hostname`, since this is where the host name data
resides in the JSON string. Tyk automatically assumes that
`dataPath` in this case is in a string-encoded JSON object and
will try to deserialize it.

Tyk classic API definition: `service_discovery.parent_data_path`

**PortDataPath** (JSON: `portDataPath`)

PortDataPath is the port of the data path. In the above nested example, we can see that there is a separate `port` value
for the service in the nested JSON. In this case, you can set the `portDataPath` value and Tyk will treat `dataPath` as
the hostname and zip them together (this assumes that the hostname element does not end in a slash or resource identifier
such as `/widgets/`). In the above example, the `portDataPath` would be `port`.

Tyk classic API definition: `service_discovery.port_data_path`

**UseTargetList** (JSON: `useTargetList`)

UseTargetList should be set to `true`, if you are using load balancing. Tyk will treat the data path as a list and
inject it into the target list of your API definition.

Tyk classic API definition: `service_discovery.use_target_list`

**CacheTimeout** (JSON: `cacheTimeout`)

CacheTimeout is the timeout of a cache value when a new data is loaded from a discovery service.
Setting it too low will cause Tyk to call the SD service too often, setting it too high could mean that
failures are not recovered from quickly enough.

Deprecated: The field is deprecated, usage needs to be updated to configure caching.

Tyk classic API definition: `service_discovery.cache_timeout`

**Cache** (JSON: `cache`)

Cache holds cache related flags.

Tyk classic API definition:
- `service_discovery.cache_disabled`
- `service_discovery.cache_timeout`

**TargetPath** (JSON: `targetPath`)

TargetPath is to set a target path to append to the discovered endpoint, since many SD services
only provide host and port data. It is important to be able to target a specific resource on that host.
Setting this value will enable that.

Tyk classic API definition: `service_discovery.target_path`

**EndpointReturnsList** (JSON: `endpointReturnsList`)

EndpointReturnsList is set `true` when the response type is a list instead of an object.

Tyk classic API definition: `service_discovery.endpoint_returns_list`

# ServiceDiscoveryCache

ServiceDiscoveryCache holds configuration for caching ServiceDiscovery data.

**Enabled** (JSON: `enabled`)

Enabled turns service discovery cache on or off.

Tyk classic API definition: `service_discovery.cache_disabled`

> required

**Timeout** (JSON: `timeout`)

Timeout is the TTL for a cached object in seconds.

Tyk classic API definition: `service_discovery.cache_timeout`

# Signature

Signature holds the configuration for signature validation.

**Enabled** (JSON: `enabled`)



> required

**Algorithm** (JSON: `algorithm`)



**Header** (JSON: `header`)



**Query** (JSON: `query`)



**Secret** (JSON: `secret`)



**AllowedClockSkew** (JSON: `allowedClockSkew`)



**ErrorCode** (JSON: `errorCode`)



**ErrorMessage** (JSON: `errorMessage`)



# State

State holds configuration about API definition states (active, internal).

**Active** (JSON: `active`)

Active enables the API.
Tyk classic API definition: `active`

> required

**Internal** (JSON: `internal`)

Internal makes the API accessible only internally.
Tyk classic API definition: `internal`

# StructInfo

StructInfo holds ast field information for the docs generator.

**Name** (JSON: `Name`)

Name is struct go name.

**Fields** (JSON: `fields`)

Fields holds information of the fields, if this object is a struct.

# Test

Test holds the test configuration for service discovery.

**ServiceDiscovery** (JSON: `serviceDiscovery`)

ServiceDiscovery contains the configuration related to test Service Discovery.
Tyk classic API definition: `proxy.service_discovery`

# Token

Token holds the values related to authentication tokens.

**Enabled** (JSON: `enabled`)

Enabled enables the token based authentication mode.

Tyk classic API definition: `auth_configs["authToken"].use_standard_auth`

> required

**EnableClientCertificate** (JSON: `enableClientCertificate`)

EnableClientCertificate allows to create dynamic keys based on certificates.

Tyk classic API definition: `auth_configs["authToken"].use_certificate`

**Signature** (JSON: `signatureValidation`)

Signature holds the configuration for verifying the signature of the token.

Tyk classic API definition: `auth_configs["authToken"].use_certificate`

# TransformRequestBody

TransformRequestBody holds configuration about body request transformations.

**Enabled** (JSON: `enabled`)

Enabled enables transform request body middleware.

**Format** (JSON: `format`)

Format of the request body, xml or json.

**Path** (JSON: `path`)

Path file path for the template.

**Body** (JSON: `body`)

Body base64 encoded representation of the template.

# TransformRequestMethod

TransformRequestMethod holds configuration for rewriting request methods.

**Enabled** (JSON: `enabled`)

Enabled enables Method Transform for the given path and method.

**ToMethod** (JSON: `toMethod`)

ToMethod is the http method value to which the method of an incoming request will be transformed.

# TykExtensionConfigParams

TykExtensionConfigParams holds the essential configuration required for the Tyk Extension schema.

**UpstreamURL** (JSON: `UpstreamURL`)



**ListenPath** (JSON: `ListenPath`)



**CustomDomain** (JSON: `CustomDomain`)



**ApiID** (JSON: `ApiID`)



**Authentication** (JSON: `Authentication`)



**AllowList** (JSON: `AllowList`)



**ValidateRequest** (JSON: `ValidateRequest`)



**MockResponse** (JSON: `MockResponse`)



# ValidateRequest

ValidateRequest holds configuration required for validating requests.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag, if set to `true`, it enables request validation.

**ErrorResponseCode** (JSON: `errorResponseCode`)

ErrorResponseCode is the error code emitted when the request fails validation.
If unset or zero, the response will returned with http status 422 Unprocessable Entity.

# VersionToID

VersionToID contains a single mapping from a version name into an API ID.

**Name** (JSON: `name`)

Name contains the user chosen version name, e.g. `v1` or similar.

**ID** (JSON: `id`)

ID is the API ID for the version set in Name.

# Versioning

Versioning holds configuration for API versioning.

Tyk classic API definition: `version_data`.

**Enabled** (JSON: `enabled`)

Enabled is a boolean flag, if set to `true` it will enable versioning of an API.

> required

**Name** (JSON: `name`)

Name contains the name of the version as entered by the user ("v1" or similar).

**Default** (JSON: `default`)

Default contains the default version name if a request is issued without a version.

> required

**Location** (JSON: `location`)

Location contains versioning location information. It can be one of the following:

- `header`,
- `url-param`,
- `url`.

> required

**Key** (JSON: `key`)

Key contains the name of the key to check for versioning information.

> required

**Versions** (JSON: `versions`)

Versions contains a list of versions that map to individual API IDs.

> required

**StripVersioningData** (JSON: `stripVersioningData`)

StripVersioningData is a boolean flag, if set to `true`, the API responses will be stripped of versioning data.

# VirtualEndpoint

VirtualEndpoint contains virtual endpoint configuration.

**Enabled** (JSON: `enabled`)

Enabled enables virtual endpoint.

> required.

**Name** (JSON: `name`)

Name is the name of js function.

> required.

**Path** (JSON: `path`)

Path is the path to js file.

**Body** (JSON: `body`)

Body is the js function to execute encoded in base64 format.

**ProxyOnError** (JSON: `proxyOnError`)

ProxyOnError proxies if virtual endpoint errors out.

**RequireSession** (JSON: `requireSession`)

RequireSession if enabled passes session to virtual endpoint.

# XTykDoc

XTykDoc is a list of information for exported struct type info,
starting from the root struct declaration(XTykGateway).

No exposed fields available.

