# Model
model:
  rest_name: cloudgraph
  resource_name: cloudgraphs
  entity_name: CloudGraph
  package: pcn/infrastructure
  group: pcn/infrastructure
  description: "Returns a data structure representing the graph of all cloud nodes
    \nand their connections in a particular namespace."

# Attributes
attributes:
  v1:
  - name: edges
    description: The edges of the map.
    type: refMap
    exposed: true
    subtype: cloudedge
    read_only: true
    extensions:
      refMode: pointer

  - name: nodes
    description: Refers to the nodes of the map.
    type: refMap
    exposed: true
    subtype: cloudnode
    read_only: true
    extensions:
      refMode: pointer

  - name: requestType
    description: The type of request/calculation that must be performedn.
    type: enum
    exposed: true
    allowed_choices:
    - Topology
    - TraceRoute
    default_value: Topology

  - name: targetVPCs
    description: The VPCs that should be captured in the map.
    type: list
    exposed: true
    subtype: string
    stored: true
