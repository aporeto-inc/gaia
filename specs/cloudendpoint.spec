# Model
model:
  rest_name: cloudendpoint
  resource_name: cloudendpoints
  entity_name: CloudEndpoint
  package: pcn
  group: prisma/infrastructure
  description: Manages the list of endpoints available in a cloud deployment.
  get:
    description: Retrieves the endpoint with the given ID.
  update:
    description: Updates the endpoint with the given ID.
  delete:
    description: Deletes the endpoint with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'

# Indexes
indexes:
- - nativeID
- - namespace
  - nativeID
- - vpcid

# Attributes
attributes:
  v1:
  - name: parameters
    description: Endpoint related parameters.
    type: ref
    exposed: true
    subtype: cloudendpointdata
    stored: true
    extensions:
      refMode: pointer
