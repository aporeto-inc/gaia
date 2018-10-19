# Model
model:
  rest_name: service
  resource_name: services
  entity_name: Service
  package: squall
  description: |-
    A Service defines a generic service object at L4 or L7 that encapsulates the
    description of a micro-service. A service exposes APIs and can be implemented
    through third party entities (such as a cloud provider) or through  processing
    units.
  aliases:
  - srv
  indexes:
  - - namespace
  - - namespace
    - archived
  - - namespace
    - normalizedtags
  - - allAPITags
  - - namespace
    - allAPITags
  - - allServiceTags
  - - namespace
    - allServiceTags
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@metadatable'

# Attributes
attributes:
  v1:
  - name: IPs
    description: |-
      IPs is the list of IP addresses where the service can be accessed.
      This is an optional attribute and is only required if no host names are
      provided.
      The system will automatically resolve IP addresses from  host names otherwise.
    type: external
    exposed: true
    subtype: ip_list
    stored: true

  - name: JWTSigningCertificate
    description: |-
      JWTSigningCertificate is a certificate that can be used to validate user JWT in
      HTTP requests. This is an optional field, needed only if user JWT validation is
      required for this service. The certificate must be in PEM format.
    type: string
    exposed: true
    stored: true

  - name: MTLSCertificateAuthority
    description: |-
      Base64 encoded version of the Certificate Authority to use to verify client
      certificates. This only applies if `MTLSType` is set to
      `VerifyClientCertIfGiven` or `RequireAndVerifyClientCert`. If it is not set,
      Aporeto own Authority will be used.
    type: string
    exposed: true
    stored: true

  - name: MTLSType
    description: |-
      Set how to perform mtls authorization. This is only applicable it
      `authorizationType` is set to `MTLS` otherwise this property has no effect.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - RequireAnyClientCert
    - RequireAndVerifyClientCert
    default_value: RequireAndVerifyClientCert

  - name: OIDCClientID
    description: |-
      authorizationID is only valid for OIDC authorization and defines the
      issuer ID of the OAUTH token.
    type: string
    exposed: true
    stored: true

  - name: OIDCClientSecret
    description: |-
      authorizationSecret is only valid for OIDC authorization and defines the
      secret that should be used with the OAUTH provider to validate tokens.
    type: string
    exposed: true
    stored: true

  - name: OIDCProviderURL
    description: |-
      authorizationProvider is only valid for OAUTH authorization and defines the
      URL to the OAUTH provider that must be used.
    type: string
    exposed: true
    stored: true

  - name: OIDCScopes
    description: Configures the scopes you want to add to the OIDC provider.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: TLSCertificate
    description: |-
      If `TLSType` is set to `External`, this property sets the base64 encoded
      certificate to expose to the client for TLS.
    type: string
    exposed: true
    stored: true

  - name: TLSCertificateKey
    description: |-
      If `TLSType` is set to `External`, this property sets the base64 encoded
      certificate key associated to `TLSCertificate`.
    type: string
    exposed: true
    stored: true

  - name: TLSType
    description: |-
      Set how to provide a server certificate to the service.

      * `Aporeto`: Generate a certificate issued from Aporeto public CA.
      * `LetsEncrypt`: Issue a certificate from letsencrypt.
      * `External`: : Let you define your own certificate and key to use.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - LetsEncrypt
    - External
    - None
    default_value: Aporeto

  - name: allAPITags
    description: This is a set of all API tags for matching in the DB.
    type: external
    subtype: tags_list
    stored: true
    read_only: true

  - name: allServiceTags
    description: This is a set of all selector tags for matching in the DB.
    type: external
    subtype: tags_list
    stored: true
    read_only: true

  - name: authorizationClaimMappings
    description: |-
      authorizationClaimMappings defines a list of mappings between incoming and
      HTTP headers. When these mappings are defined, the enforcer will copy the
      values of the claims to the corresponding HTTP headers.
    type: refList
    exposed: true
    subtype: claimmapping
    stored: true
    extensions:
      refMode: pointer

  - name: authorizationType
    description: |-
      AuthorizationType defines the user authorization type that should be used.

      * `None`: No auhtorization. The service will only provide server TLS.
      * `JWT`:  Configures a simple JWT verification from the `Auhorization` Header
      * `OIDC`: Configures OIDC authorization. You must then set `OIDCClientID`,
      `OIDCClientSecret`, OIDCProviderURL`.
      * `MTLS`: Configures Client Certificate authorization. You must then set
      `MTLSType` and eventually `MTLSCertificateAuthority`.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - None
    - JWT
    - OIDC
    - MTLS
    default_value: None

  - name: endpoints
    description: |-
      Endpoints is a read only attribute that actually resolves the API
      endpoints that the service is exposing. Only valid during policy rendering.
    type: external
    exposed: true
    subtype: exposed_api_list
    read_only: true

  - name: exposedAPIs
    description: |-
      ExposedAPIs contains a tag expression that will determine which
      APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
      similar specifications for other L7 protocols.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    example_value:
    - - package=p1

  - name: exposedPort
    description: |-
      ExposedPort is the port that the service can be accessed. Note that
      this is different from the Port attribute that describes the port that the
      service is actually listening. For example if a load balancer is used, the
      ExposedPort is the port that the load balancer is listening for the service,
      whereas the port that the implementation is listening can be different.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 443
    max_value: 65535

  - name: external
    description: External is a boolean that indicates if this is an external service.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    filterable: true
    orderable: true

  - name: hosts
    description: Hosts are the names that the service can be accessed with.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: port
    description: |-
      Port is the port that the implementation of the service is listening to and
      it can be different than the exposedPorts describing the service. This is needed
      for port mapping use cases where there is private and public ports.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 443
    max_value: 65535

  - name: publicApplicationPort
    description: |-
      PublicApplicationPort is a new virtual port that the service can
      be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
      application path, you might want to declare a new port where the enforcer
      listens for TLS. However, the application does not need to be modified and
      the enforcer will map the traffic to the correct application port. This useful
      when an application is being accessed from a public network.
    type: integer
    exposed: true
    stored: true
    example_value: 443
    max_value: 65535

  - name: redirectOnAuthorizationFailure
    description: |-
      If this is set, the user will be redirected to that URL in case of any
      authorization failure to let you chance to provide a nice message to the user.
      The query parameter `?failure_message=<message>` will be added to that url
      explaining the possible reasons of the failure.
    type: string
    exposed: true
    stored: true

  - name: selectors
    description: |-
      Selectors contains the tag expression that an a processing unit
      must match in order to implement this particular service.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    example_value:
    - - $identity=processingunit

  - name: serviceCA
    description: |-
      ServiceCA  is the certificate authority that the service is using. This
      is needed for external services with private certificate authorities. The
      field is optional. If provided, this must be a valid PEM CA file.
    type: string
    exposed: true
    stored: true

  - name: type
    description: Type is the type of the service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - HTTP
    - TCP
    - KubernetesSecrets
    - VaultSecrets
    default_value: HTTP

# Relations
relations:
- rest_name: restapispec
  get:
    description: Retrieves the REST APIs exposed by this service.

- rest_name: processingunit
  get:
    description: Retrieves the Processing Units that implement this service.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subjects
        - object
        default_value: objects
