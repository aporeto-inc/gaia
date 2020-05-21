# Model
model:
  rest_name: pingresult
  resource_name: pingresults
  entity_name: PingResult
  package: guy
  group: core/enforcer
  description: Post a new pu pingresult.

# Attributes
attributes:
  v1:
  - name: pingPairs
    description: Contains the result of aggregated ping pairs.
    type: refList
    exposed: true
    subtype: pingpair
    extensions:
      refMode: pointer
