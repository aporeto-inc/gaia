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
  - name: address
    description: The IP subnet address.
    type: string
    exposed: true
    stored: true
    example_value: 10.1.1.0/24

  - name: zoneId
    description: The availability zone of the subnet.
    type: string
    exposed: true
    stored: true
    example_value: aws-east

  - name: zoneName
    description: The availability zone of the subnet.
    type: string
    exposed: true
    stored: true
    example_value: aws-east
