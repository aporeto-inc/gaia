# Model
model:
  rest_name: cloudsubnet
  resource_name: cloudsubnets
  entity_name: CloudSubnet
  package: yeul
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
  - '@identifiable-stored'
  - '@prismabase'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Subnet related parameters.
    type: ref
    exposed: true
    subtype: cloudsubnetdata
    stored: true
    extensions:
      refMode: pointer
