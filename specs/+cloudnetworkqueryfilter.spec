# Model
model:
  rest_name: cloudnetworkqueryfilter
  resource_name: cloudnetworkqueryfilters
  entity_name: CloudNetworkQueryFilter
  package: yeul
  group: prisma/infrastructure
  description: |-
    Captures the parameters allowed in a query filter for a net effective
    permissions request.
  detached: true

# Attributes
attributes:
  v1:
  - name: AccountIDs
    description: |-
      The accounts that the search must apply to. These are the actually IDs of the
      account as provided by the cloud provider. One or more IDs can be included.
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
      the fields are ignored. An object ID can refer to an instance, VPC endpoint, or
      network interface.
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
    description: |-
      The type of endpoint resource. The resource type is a mandatory field and a
      query cannot span multiple resource types.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Instance
    - Interface
    - Service
    - ProcessingUnit
    default_value: Instance

  - name: SecurityTags
    description: |-
      The list of security tags associated with the targets of the query. Security
      tags refer to security groups in AWS or network tags in GCP. So they can have
      different meaning depending on the target cloud.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: ServiceOwners
    description: |-
      Identifies the owner of the service that the resource is attached to. Field is
      not valid if the resource type is not an interface.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: ServiceTypes
    description: |-
      Identifies the type of service that the interface is attached to. Field is not
      valid if the resource type is not an
      interface.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: Subnets
    description: |-
      The subnets where the resources must reside. A subnet parameter can only be
      provided for a network interface resource type.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: Tags
    description: |-
      A list of tags that select the same of endpoints for the query. These tags refer
      to the tags attached to the resources in the cloud provider definitions.
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
