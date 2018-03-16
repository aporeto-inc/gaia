attributes:
- description: FQDN is the fully qualified domain name of the service. It is required
    for external API services. It can be deduced from a service discovery system in
    advanced environments.
  exposed: true
  filterable: true
  format: free
  name: FQDN
  orderable: true
  stored: true
  type: string
- description: IPList is the list of ip address or subnets of the service if available.
  exposed: true
  name: IPList
  stored: true
  subtype: ip_list
  type: external
- description: JWTSigningCertificate is a certificate that can be used to validate
    user JWT in HTTP requests. This is an optional field, needed only if user JWT
    validation is required for this service. The certificate must be in PEM format.
  exposed: true
  format: free
  name: JWTSigningCertificate
  stored: true
  type: string
- description: AllServiceTags is an internal object that summarizes all the implementedBy
    tags to accelerate database searches. It is not exposed.
  name: allServiceTags
  stored: true
  subtype: tags_list
  type: external
- description: ExposedAPIs is a list of API endpoints that are exposed for the service.
  exposed: true
  name: exposedAPIs
  stored: true
  subtype: exposed_api_list
  type: external
- default_value: "false"
  description: External is a boolean that indicates if this is an external service.
  exposed: true
  filterable: true
  name: external
  orderable: true
  stored: true
  type: boolean
- description: ExternalServiceCA is the certificate authority that the service is
    using. This is needed for external API services with private certificate authorities.
    The field is optional. If provided, this must be a valid PEM CA file.
  exposed: true
  format: free
  name: externalServiceCA
  stored: true
  type: string
- default_value: "6"
  description: 'NetworkProtocol is the network protocol of the service. Default is
    TCP. '
  exposed: true
  filterable: true
  max_value: 255
  min_value: 1
  name: networkProtocol
  orderable: true
  stored: true
  type: integer
- description: Ports is a list of ports for the service. Ports are either exact match,
    or a range portMin:portMax.
  exposed: true
  name: ports
  required: true
  stored: true
  subtype: port_list
  type: external
- description: RuntimeSelectors is a list of tag selectors that identifies that Processing
    Units that will implement this service.
  exposed: true
  name: runtimeSelectors
  required: true
  stored: true
  subtype: target_tags
  type: external
- allowed_choices:
  - HTTP
  - L3
  - TCP
  default_value: L3
  description: Type is the type of the service (HTTP, TCP, etc). More types will be
    added to the system.
  exposed: true
  filterable: true
  name: type
  orderable: true
  required: true
  stored: true
  type: enum
model:
  delete: true
  get: true
  update: true
  description: APIService descibes a L4/L7 service and the corresponding implementation.
    It allows users to define their services, the APIs that they expose, the implementation
    of the service. These definitions can be used by network policy in order to define
    advanced controls based on the APIs.
  entity_name: APIService
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
  package: squall
  resource_name: apiservices
  rest_name: apiservice
