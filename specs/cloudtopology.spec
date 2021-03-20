# Model
model:
  rest_name: cloudtopology
  resource_name: cloudtopologies
  entity_name: CloudTopology
  package: yeul
  group: pcn/infrastructure
  description: Returns the full topology of all nodes and their relationships.

# Attributes
attributes:
  v1:
  - name: edges
    description: The edges of the map.
    type: refMap
    exposed: true
    subtype: cloudgraphedge
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

  - name: targetVPCs
    description: |-
      The target VPCs for the topology. If empty, it will return the topology for all
      VPCs.
    type: list
    exposed: true
    subtype: string
