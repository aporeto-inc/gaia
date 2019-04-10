# Model
model:
  rest_name: graphnode
  resource_name: graphnodes
  entity_name: GraphNode
  package: jenova
  group: visualization/depmaps
  description: Represents an node from the dependency map.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of object represented by the node.
    type: string
    exposed: true

  - name: enforcementStatus
    description: Enforcement status of processing unit represented by the node.
    type: string
    exposed: true

  - name: groupID
    description: ID of the group the node is eventually part of.
    type: string
    exposed: true

  - name: images
    description: List of images.
    type: list
    exposed: true
    subtype: string

  - name: lastUpdate
    description: Last update of the node.
    type: time
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

  - name: unreachable
    description: If true the node is marked as unreachable.
    type: boolean
    exposed: true

  - name: vulnerabilityLevel
    description: Tags of object represented by the node.
    type: string
    exposed: true
