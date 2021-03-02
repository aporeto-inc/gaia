# Model
model:
  rest_name: cloudattachment
  resource_name: cloudattachments
  entity_name: CloudAttachment
  package: pcn
  group: prisma/infrastructure
  description: |-
    Cloud attachments provide a means to model relationships between cloud objects.
    For example relationships between Transit Gateways and Transit Gateway
    Attachments.
  get:
    description: Retrieves the cloud attachment with the given ID.
  update:
    description: Updates the cloud attachment with the given ID.
  delete:
    description: Deletes the cloud attachment with the given ID.
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
  - name: attachments
    description: List of attached objects.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: nodeID
    description: The node that the attachments are connected to.
    type: string
    exposed: true
    stored: true

  - name: type
    description: Type of attachment.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - TransitGatewayVPCAttachment
    - TransitGatewayAttachment
