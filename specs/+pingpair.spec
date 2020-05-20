# Model
model:
  rest_name: pingpair
  resource_name: pingpairs
  entity_name: PingPair
  package: squall
  group: policy/processingunits
  description: Represents a service attached to a processing unit.
  detached: true

# Attributes
attributes:
  v1:
  - name: request
    description: Contains the request probe information.
    type: ref
    exposed: true
    subtype: pingprobe
    stored: true
    extensions:
      refMode: pointer

  - name: response
    description: Contains the response probe information.
    type: ref
    exposed: true
    subtype: pingprobe
    stored: true
    extensions:
      refMode: pointer
