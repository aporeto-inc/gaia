# Model
model:
  rest_name: service
  resource_name: services
  entity_name: Service
  package: squall
  description: |-
    A Service allows to declare what Services processing
    units are exposing or consuming.
  aliases:
  - srv
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-nopk-nostored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: IPs
    description: |-
      ExposedIPs is the list of IP addresses or subnets of the servers behind the
      services
      matched by object. This is an optional field and it can be
      automatically populated at runtime by the enforcers if DNS resolution is
      available.
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

  - name: exposedAPIs
    description: |-
      ExposedAPIs contains the tag expression that an object must match in order to
      trigger the hook.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
    example_value:
    - - package=p1

  - name: exposedPort
    description: |-
      Ports is a list of the public ports for the service. Ports are either
      exact match, or a range portMin:portMax. For HTTP services only exact match
      ports aresupported. These should be the ports that are used by other services
      to communicate with the defined service.
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
    description: |-
      Hosts  is the fully qualified domain name of the the servers behind the
      services
      matched by object. FQDN must match the host part
      of the URI that is used to call a service. It will be used for automatically
      generating service certificates for internal services.
    type: list
    exposed: true
    subtype: string
    stored: true
    format: free
    orderable: true

  - name: port
    description: |-
      Port is the port that the application is listening to and
      it can be different than the ports describing the service. This is needed for
      port mapping use cases where there is private and public ports.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 443
    max_value: 65535
    min_value: 1

  - name: selectors
    description: |-
      Selectors contains the tag expression that an a processing unit
      must match in order to
      trigger the hook.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
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
