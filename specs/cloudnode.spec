# Model
model:
  rest_name: cloudnode
  resource_name: cloudnodes
  entity_name: CloudNode
  package: pcn
  group: prisma/infrastructure
  description: Manages the list of cloud nodes available in a cloud deployment.
  get:
    description: Retrieves the cloud node with the given ID.
  update:
    description: Updates the cloud node with the given ID.
  delete:
    description: Deletes the cloud node with the given ID.
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
- - vpcID

# Attributes
attributes:
  v1:
  - name: attachments
    description: The list of attachments for this node.
    type: list
    exposed: true
    subtype: string
    stored: true
    getter: true
    setter: true

  - name: parameters
    description: The cloud attributes of the object.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true
    getter: true
    setter: true

  - name: type
    description: Type of the endpoint.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Endpoint
    - Subnet
    - VPC
    - Interface
    - RouteTable
    example_value: Interface
