# Model
model:
  rest_name: hostservice
  resource_name: hostservices
  entity_name: HostService
  package: squall
  description: Represents a service of the enforcer's host.
  detached: true

# Attributes
attributes:
  v1:
  - name: associatedTags
    description: AssociatedTags are the list of tags attached to an entity.
    type: external
    exposed: true
    subtype: tags_list
    stored: true
    getter: true
    setter: true

  - name: name
    description: Name of the service.
    type: string
    exposed: true

  - name: services
    description: Services lists all protocols and ports a service is running.
    type: external
    exposed: true
    subtype: processing_unit_services_list
    stored: true
    orderable: true
