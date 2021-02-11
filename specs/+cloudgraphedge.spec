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

# Indexes
indexes:
- - namespace

# Attributes
attributes:
  v1:
  - name: destinationID
    description: ID of the destination `cloud node` of the edge.
    type: string
    exposed: true
    stored: true

  - name: sourceID
    description: ID of the source `cloud node` of the edge.
    type: string
    exposed: true
    stored: true
