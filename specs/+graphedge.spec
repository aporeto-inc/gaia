# Model
model:
  rest_name: graphedge
  resource_name: graphedges
  entity_name: GraphEdge
  package: jenova
  group: visualization/depmaps
  description: Represents an edge from the dependency map.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of the edge.
    type: string
    exposed: true

  - name: acceptedFlows
    description: Number of accepted flows in the edge.
    type: integer
    exposed: true

  - name: destinationID
    description: ID of the destination GraphNode of the edge.
    type: string
    exposed: true

  - name: destinationType
    description: Type of the destination GraphNode of the edge.
    type: enum
    exposed: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Node

  - name: encrypted
    description: Tells the number of encrypted flows in the edge.
    type: integer
    exposed: true

  - name: observedAcceptedFlows
    description: Number of accepted observed flows.
    type: integer
    exposed: true

  - name: observedEncrypted
    description: Number of encrypted observed flows.
    type: integer
    exposed: true

  - name: observedPolicyIDs
    description: |-
      Information about the observation policies that was hit in the flows
      represented by that edge.
    type: refMap
    exposed: true
    subtype: graphpolicyinfo
    extensions:
      refMode: pointer

  - name: observedRejectedFlows
    description: Number of rejected observed flows.
    type: integer
    exposed: true

  - name: observedServiceIDs
    description: Map of ints...
    type: external
    exposed: true
    subtype: map[string]int

  - name: policyIDs
    description: |-
      Information about the policies that was hit in the flows represented by that
      edge.
    type: refMap
    exposed: true
    subtype: graphpolicyinfo
    extensions:
      refMode: pointer

  - name: rejectedFlows
    description: Number of rejected flows in the edge.
    type: integer
    exposed: true

  - name: sourceID
    description: ID of the source GraphNode of the edge.
    type: string
    exposed: true

  - name: sourceType
    description: Type of the source GraphNode of the edge.
    type: enum
    exposed: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Node
