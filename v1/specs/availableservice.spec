attributes:
- description: Beta indicates if the service is in a beta version.
  exposed: true
  name: beta
  read_only: true
  type: boolean
- description: CategoryID of the service.
  exposed: true
  filterable: true
  format: free
  name: categoryID
  read_only: true
  type: string
- description: Description of the service
  exposed: true
  format: free
  name: description
  type: string
- description: 'Icon contains a base64 image for the available service. '
  exposed: true
  format: free
  name: icon
  read_only: true
  type: string
- description: LongDescription contains a more detailed description of the service.
  exposed: true
  format: free
  name: longDescription
  type: string
- description: Name of the Service
  exposed: true
  format: free
  name: name
  type: string
- description: Parameters of the service the user can or has to specify
  exposed: true
  name: parameters
  subtype: service_parameters
  type: external
- description: Title represents the title of the service.
  exposed: true
  format: free
  name: title
  type: string
model:
  aliases:
  - asrv
  description: AvailableService represents a service that is available for launching
  entity_name: AvailableService
  package: highwind
  resource_name: availableservices
  rest_name: availableservice
