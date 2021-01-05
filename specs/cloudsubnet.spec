# Model
model:
  rest_name: cloudsubnet
  resource_name: cloudsubnets
  entity_name: CloudSubnet
  package: pcn
  group: prisma/infrastructure
  description: Manages the list of subnets associated with a deployment.
  get:
    description: Retrieves the subnet with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the subnet with the given ID.
  delete:
    description: Deletes the subnet with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@prismabase'

# Indexes
indexes:
- - nativeID
- - namespace
  - nativeID

# Attributes
attributes:
  v1:
  - name: parameters
    description: Subnet related parameters.
    type: ref
    exposed: true
    subtype: subnetdata
    stored: true
    extensions:
      refMode: pointer
