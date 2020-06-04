# Model
model:
  rest_name: pingpair
  resource_name: pingpairs
  entity_name: PingPair
  package: guy
  group: core/enforcer
  description: Represents a pair of ping probes.
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
      noInit: true
      refMode: pointer

  - name: response
    description: Contains the response probe information.
    type: ref
    exposed: true
    subtype: pingprobe
    stored: true
    extensions:
      noInit: true
      refMode: pointer
