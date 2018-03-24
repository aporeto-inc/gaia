# Model
model:
  rest_name: availableservice
  resource_name: availableservices
  entity_name: AvailableService
  package: highwind
  description: AvailableService represents a service that is available for launching
  aliases:
  - asrv
  extends:
  - '@described'
  - '@named'

# Attributes
attributes:
- name: beta
  description: Beta indicates if the service is in a beta version.
  type: boolean
  exposed: true
  read_only: true

- name: categoryID
  description: CategoryID of the service.
  type: string
  exposed: true
  read_only: true
  filterable: true
  format: free

- name: icon
  description: Icon contains a base64 image for the available service.
  type: string
  exposed: true
  read_only: true
  format: free

- name: longDescription
  description: LongDescription contains a more detailed description of the service.
  type: string
  exposed: true
  format: free

- name: parameters
  description: Parameters of the service the user can or has to specify
  type: external
  exposed: true
  subtype: service_parameters

- name: title
  description: Title represents the title of the service.
  type: string
  exposed: true
  format: free
