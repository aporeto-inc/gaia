# Model
model:
  rest_name: cloudgraphedge
  resource_name: cloudgraphedges
  entity_name: CloudGraphEdge
  package: yeul
  group: pcn/infrastructure
  description: Represents an edge in the configuration map.
  private: true
  detached: true

# Attributes
attributes:
  v1:
  - name: level
    description: |-
      Provides the level of the tree that this edge belongs in order to assist with
      ordering.
    type: integer
    exposed: true
    stored: true
