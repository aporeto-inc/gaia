# Model
model:
  rest_name: cloudroutetable
  resource_name: cloudroutetables
  entity_name: CloudRouteTable
  package: pcn
  group: prisma/infrastructure
  description: Manages the list of route tables available in a cloud deployment.
  get:
    description: Retrieves the route table with the given ID.
  update:
    description: Updates the route table with the given ID.
  delete:
    description: Deletes the route table with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Indexes
indexes:
- - nativeID
- - namespace
  - nativeID

# Attributes
attributes:
  v1:
  - name: parameters
    description: Route table related parameters.
    type: ref
    exposed: true
    subtype: cloudroutedata
    stored: true
    extensions:
      refMode: pointer
