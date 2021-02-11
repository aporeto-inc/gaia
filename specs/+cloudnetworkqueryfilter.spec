# Model
model:
  rest_name: cloudnetworkqueryfilter
  resource_name: cloudnetworkqueryfilters
  entity_name: CloudNetworkQueryFilter
  package: pcn
  group: prisma/infrastructure
  description: Parameters associated with a cloud endpoint.
  detached: true

# Attributes
attributes:
  v1:
  - name: AccountIDs
    description: The accounts that the search must apply to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: CloudTypes
    description: The cloud types that the search must apply to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: ObjectIDs
    description: |-
      The exact object that the search applies. If ObjectIDs are defined, the rest of
      the fields are ignored.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: Regions
    description: The region that the search must apply to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: ResourceType
    description: The type of endpoint resource.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Instance
    - Interface
    - Service
    - ProcessingUnit
    default_value: Instance

  - name: SecurityTags
    description: The list of security tags associated with the targets of the query.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: Subnets
    description: The subnets where the resources must reside.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: Tags
    description: A list of tags that select the same of endpoints for the query.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: VPCIDs
    description: The VPC ID of the target resources.
    type: list
    exposed: true
    subtype: string
    stored: true
