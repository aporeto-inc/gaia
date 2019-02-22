# Model
model:
  rest_name: service
  resource_name: services
  entity_name: Service
  package: squall
  group: policy/services
  description: |-
    A Service defines a generic service object at L4 or L7 that encapsulates the
    description of a micro-service. A service exposes APIs and can be implemented
    through third party entities (such as a cloud provider) or through  processing
    units.
  aliases:
  - srv
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - archived
  - - namespace
    - normalizedTags
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
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@disabled'
  - '@zonable'
  validations:
  - $serviceEntity

# Attributes
attributes:
  v1:
  - name: IPs
    description: |-
      IPs is the list of IP addresses where the service can be accessed.
      This is an optional attribute and is only required if no host names are
      provided.
      The system will automatically resolve IP addresses from host names otherwise.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: JWTSigningCertificate
    description: |-
      PEM encoded certificate that will be used to validate user JWT in HTTP requests.
      This is an optional field, needed only if the `authorizationType`
      is set to `JWT`.
    type: string
    exposed: true
    stored: true

  - name: MTLSCertificateAuthority
    description: |-
      PEM encoded Certificate Authority to use to verify client
      certificates. This only applies if `authorizationType` is set to
      `MTLS`. If it is not set, Aporeto's Public Signing Certificate Authority will
      be used.
    type: string
    exposed: true
    stored: true

  - name: OIDCCallbackURL
    description: |-
      This is an advanced setting. Optional OIDC callback URL. If you don't set it,
      Aporeto will autodiscover it. It will be
      `https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback`.
    type: string
    exposed: true
    stored: true

  - name: OIDCClientID
    description: OIDC Client ID. Only has effect if the `authorizationType` is set
      to `OIDC`.
    type: string
    exposed: true
    stored: true

  - name: OIDCClientSecret
    description: OIDC Client Secret. Only has effect if the `authorizationType` is
      set to `OIDC`.
    type: string
    exposed: true
    stored: true

  - name: OIDCProviderURL
    description: OIDC Provider URL. Only has effect if the `authorizationType` is
      set to `OIDC`.
    type: string
    exposed: true
    stored: true
    example_value: https://accounts.google.com

  - name: OIDCScopes
    description: |-
      Configures the scopes you want to add to the OIDC provider. Only has effect if
      `authorizationType` is set to `OIDC`.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile

  - name: TLSCertificate
    description: |-
      PEM encoded certificate to expose to the clients for TLS. Only has effect and
      required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true

  - name: TLSCertificateKey
    description: |-
      PEM encoded certificate key associated to `TLSCertificate`. Only has effect and
      required if `TLSType` is set to `External`.
    type: string
    exposed: true
    stored: true

  - name: TLSType
    description: |-
      Set how to provide a server certificate to the service.

      - `Aporeto`: Generate a certificate issued from Aporeto public CA.
      - `LetsEncrypt`: Issue a certificate from letsencrypt.
      - `External`: : Let you define your own certificate and key to use.
      - `None`: : TLS is disabled (not recommended).
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
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: allServiceTags
    description: This is a set of all selector tags for matching in the DB.
    type: list
    subtype: string
    stored: true
    read_only: true

  - name: authorizationType
    description: |-
      AuthorizationType defines the user authorization type that should be used.

      - `None`: No auhtorization.
      - `JWT`:  Configures a simple JWT verification from the HTTP `Auhorization`
      Header
      - `OIDC`: Configures OIDC authorization. You must then set `OIDCClientID`,
      `OIDCClientSecret`, OIDCProviderURL`.
      - `MTLS`: Configures Client Certificate authorization. Then you can optionaly
      `MTLSCertificateAuthority` otherwise Aporeto Public Signing Certificate will be
      used.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - None
    - JWT
    - OIDC
    - MTLS
    default_value: None

  - name: claimsToHTTPHeaderMappings
    description: |-
      Defines a list of mappings between claims and
      HTTP headers. When these mappings are defined, the enforcer will copy the
      values of the claims to the corresponding HTTP headers.
    type: refList
    exposed: true
    subtype: claimmapping
    stored: true
    extensions:
      refMode: pointer

  - name: endpoints
    description: |-
      Endpoints is a read only attribute that actually resolves the API
      endpoints that the service is exposing. Only valid during policy rendering.
    type: refList
    exposed: true
    subtype: endpoint
    read_only: true
    extensions:
      refMode: pointer

  - name: exposedAPIs
    description: |-
      ExposedAPIs contains a tag expression that will determine which
      APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
      similar specifications for other L7 protocols.
    type: external
    exposed: true
    subtype: '[][]string'
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

  - name: exposedServiceIsTLS
    description: |-
      ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
      the enforcer has to initiate a TLS session in order to forrward traffic to the
      service.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    filterable: true
    orderable: true

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

  - name: redirectURLOnAuthorizationFailure
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
    subtype: '[][]string'
    stored: true
    example_value:
    - - $identity=processingunit

  - name: trustedCertificateAuthorities
    description: |-
      PEM encoded Certificate Authorities to trust when additional hops are needed. It
      must be set if the service must reach a Service marked as `external` or must go
      through an additional TLS termination point like a L7 Load Balancer.
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
- rest_name: httpresourcespec
  get:
    description: Retrieves the HTTP Resource exposed by this service.

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
