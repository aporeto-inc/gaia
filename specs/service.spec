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
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
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

  - name: authorizationID
    description: |-
      authorizationID is only valid for OIDC authorization and defines the
      issuer ID of the OAUTH token.
    type: string
    exposed: true
    stored: true

  - name: authorizationProvider
    description: |-
      authorizationProvider is only valid for OAUTH authorization and defines the
      URL to the OAUTH provider that must be used.
    type: string
    exposed: true
    stored: true

  - name: authorizationSecret
    description: |-
      authorizationSecret is only valid for OIDC authorization and defines the
      secret that should be used with the OAUTH provider to validate tokens.
    type: string
    exposed: true
    stored: true

  - name: authorizationType
    description: |-
      AuthorizationType defines the user authorization type that should be used.
      Currently supporting PKI, and OIDC.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - PKI
    - OIDC
    - None
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

  - name: redirectOnFail
    description: |-
      RedirectOnFail is a boolean that forces a redirect response if an API request
      arrives and the user authorization information is not valid. This only applies
      to HTTP services and it is only send for APIs that are not public.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    orderable: true

  - name: redirectOnNoToken
    description: |-
      RedirectOnNoToken is a boolean that forces a redirect response if an API request
      arrives and there is no user authorization information. This only applies to
      HTTP services and it is only send for APIs that are not public.
    type: boolean
    exposed: true
    stored: true
    default_value: false
    orderable: true

  - name: redirectURL
    description: |-
      RedirectURL is the URL that will be send back to the user to
      redirect for authentication if there is no user authorization information in
      the API request. If the redirect flag is not set, this field has no meaning.The
      template is a Go Lang template where specific functions are supported.
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
