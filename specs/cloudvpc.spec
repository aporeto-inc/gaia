# Model
model:
  rest_name: cloudvpc
  resource_name: cloudvpcs
  entity_name: CloudVPC
  package: pcn
  group: prisma/infrastructure
  description: |-
    A CloudVPC represents a VPC as defined in an cloud provider (AWS/Azure/GCP etc).
    The VPC is essentially an L3 routing domain with at least one subnet attached
    and it defines an isolated network.
  aliases:
  - vpc
  - vpcs
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
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
- - namespace
  - nativeID
- - nativeID

# Attributes
attributes:
  v1:
  - name: address
    description: The IP CIDR associated with the VPC.
    type: string
    exposed: true
    stored: true
    example_value: 10.0.0.0/16
    validations:
    - $cidr
