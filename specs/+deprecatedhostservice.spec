# Model
model:
  rest_name: deprecatedhostservice
  resource_name: deprecatedhostservices
  entity_name: DeprecatedHostService
  package: squall
  group: deprecated
  description: Represents a service of the enforcer's host.
  detached: true

# Attributes
attributes:
  v1:
  - name: associatedTags
    description: AssociatedTags are the list of tags attached to an entity.
    type: list
    exposed: true
    subtype: string
    stored: true
    getter: true
    setter: true

  - name: name
    description: Name of the service.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: networkonly
    description: networkonly indicate the host service is of type network only.
    type: boolean
    exposed: true
    stored: true
    read_only: true
    default_value: true
    deprecated: true

  - name: services
    description: Services lists all protocols and ports a service is running.
    type: refList
    exposed: true
    subtype: processingunitservice
    stored: true
    orderable: true
    extensions:
      refMode: pointer
