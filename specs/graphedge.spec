# Model
model:
  rest_name: graphedge
  resource_name: graphedges
  entity_name: GraphEdge
  package: meteor
  group: visualization/depmaps
  description: Represents an edge from the dependency map.
  private: true
  extends:
  - '@zoned'
  extensions:
    forceDocumentation: true

# Indexes
indexes:
- - namespace
  - lastSeen
  - firstSeen
- - lastSeen
  - firstSeen
- - lastSeen
- - firstSeen

# Attributes
attributes:
  v1:
  - name: ID
    description: DB Identifier of the edge.
    type: string
    stored: true
    identifier: true

  - name: acceptedFlows
    description: Number of accepted flows in the edge.
    type: integer
    exposed: true
    stored: true

  - name: createTime
    description: Date on which the edge has been inserted.
    type: time
    stored: true

  - name: destinationID
    description: ID of the destination `GraphNode` of the edge.
    type: string
    exposed: true
    stored: true

  - name: destinationType
    description: Type of the destination `GraphNode` of the edge.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Node

  - name: encrypted
    description: The number of encrypted flows in the edge.
    type: integer
    exposed: true
    stored: true

  - name: firstSeen
    description: Contains the date when the edge was first seen.
    type: time
    exposed: true
    stored: true

  - name: flowID
    exposed_name: ID
    description: Identifier of the edge.
    type: string
    exposed: true
    stored: true

  - name: lastSeen
    description: Contains the date when the edge was last seen.
    type: time
    exposed: true
    stored: true

  - name: namespace
    description: Namespace of object represented by the node.
    type: string
    exposed: true
    stored: true

  - name: observedAcceptedFlows
    description: Number of accepted observed flows.
    type: integer
    exposed: true
    stored: true

  - name: observedEncrypted
    description: Number of encrypted observed flows.
    type: integer
    exposed: true
    stored: true

  - name: observedRejectedFlows
    description: Number of rejected observed flows.
    type: integer
    exposed: true
    stored: true

  - name: rejectedFlows
    description: Number of rejected flows in the edge.
    type: integer
    exposed: true
    stored: true

  - name: sourceID
    description: ID of the source `GraphNode` of the edge.
    type: string
    exposed: true
    stored: true

  - name: sourceType
    description: Type of the source `GraphNode` of the edge.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Node
