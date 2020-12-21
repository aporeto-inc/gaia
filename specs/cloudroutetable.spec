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
  - name: gatewayID
    description: The gateway that this route table is associated with.
    type: string
    exposed: true
    stored: true
    example_value: tgw-009251c49cf46d940

  - name: mainTable
    description: Indicates that this is the default route table for the VPC.
    type: boolean
    exposed: true
    stored: true
    example_value: true

  - name: routelist
    description: Routes associated with this route table.
    type: refList
    exposed: true
    subtype: cloudroute
    stored: true

  - name: subnetAssociations
    description: The list of subnets that this route table is associated with.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value: '[subnet-096bb677ed112475d]'
