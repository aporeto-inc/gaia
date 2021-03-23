# Model
model:
  rest_name: cloudendpoint
  resource_name: cloudendpoints
  entity_name: CloudEndpoint
  package: yeul
  group: pcn/infrastructure
  description: Manages the list of endpoints available in a cloud deployment.
  get:
    description: Retrieves the endpoint with the given ID.
  update:
    description: Updates the endpoint with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Endpoint related parameters.
    type: ref
    exposed: true
    subtype: cloudendpointdata
    stored: true
    filterable: true
    extensions:
      refMode: pointer
