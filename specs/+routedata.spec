# Model
model:
  rest_name: routedata
  resource_name: routedata
  entity_name: RouteData
  package: pcn
  group: prisma/infrastructure
  description: Parameters associated with a cloud route table.
  detached: true

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
