# Model
model:
  rest_name: graphnode
  resource_name: graphnodes
  entity_name: GraphNode
  package: jenova
  description: Represents an node from the dependency map.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of object represented by the node.
    type: string
    exposed: true

  - name: description
    description: Description of object represented by the node.
    type: string
    exposed: true

  - name: groupID
    description: ID of the group the node is eventually part of.
    type: string
    exposed: true

  - name: name
    description: Name of object represented by the node.
    type: string
    exposed: true

  - name: namespace
    description: Namespace of object represented by the node.
    type: string
    exposed: true

  - name: status
    description: Status of object represented by the node.
    type: string
    exposed: true

  - name: tags
    description: Tags of object represented by the node.
    type: list
    exposed: true
    subtype: string

  - name: type
    description: Type of object represented by the node.
    type: enum
    exposed: true
    allowed_choices:
    - Docker
    - ExternalNetwork
    - Volume
    - Claim
