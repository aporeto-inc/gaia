# Model
model:
  rest_name: pingresult
  resource_name: pingresults
  entity_name: PingResult
  package: guy
  group: core/enforcer
  description: Represents the results of a ping request.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@timeable'
  - '@identifiable-stored'

# Indexes
indexes:
- - pingID
- - namespace
  - pingID

# Attributes
attributes:
  v1:
  - name: errors
    description: May contain a list of errors that have happened during the collection.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: pingID
    description: Contains the Ping ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: pingPairs
    description: Contains the result of aggregated ping pairs.
    type: refList
    exposed: true
    subtype: pingpair
    stored: true
    extensions:
      refMode: pointer
