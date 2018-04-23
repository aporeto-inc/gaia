# Model
model:
  rest_name: service
  resource_name: services
  entity_name: Service
  package: squall
  description: |-
    A Service defines a generic service object at L4 or L7 that encapsulates the
    description of a micro-service. A service exposes APIs and can be implemented
    through third party entities (such as a cloud provider) or through  processign
    units.
  aliases:
  - srv
  get: true
  update: true
  delete: true
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'

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
    format: free

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
      APIs a servie is exposing. The APIs can be defined as the RESTAPISpec or
      similar specifications for other L7 protocols.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
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
    min_value: 1

  - name: external
    description: External is a boolean that indicates if this is an external service.
    type: boolean
    exposed: true
    stored: true
    default_value: "false"
    filterable: true
    orderable: true

  - name: hosts
    description: Hosts are the names that the service can be accessed with.
    type: list
    exposed: true
    subtype: string
    stored: true
    format: free
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
    default_value: 80
    example_value: 443
    max_value: 65535
    min_value: 1

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
    format: free

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
  descriptions:
    get: Retrieves the REST APIs exposed by this service.
  get: true
