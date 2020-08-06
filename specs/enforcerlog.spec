# Model
model:
  rest_name: enforcerlog
  resource_name: enforcerlog
  entity_name: EnforcerLog
  package: ifrit
  group: core/enforcer
  description: |-
    A Defender log represents the log collected by a Defender. Each Defender log
    can have partial or complete data. The `collectionID` is used to aggregate the
    multipart data into one.
  get:
    description: Retrieves the `enforcerlog` with the given ID.
  extends:
  - '@zoned'
  - '@base'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'

# Indexes
indexes:
- - enforcerID
- - namespace
  - enforcerID
- - collectionID
- - namespace
  - collectionID

# Attributes
attributes:
  v1:
  - name: collectionID
    description: |-
      Contains the ID of the Defender log. `CollectionID` is used to
      aggregate the multipart data.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: data
    description: Represents the data collected by the Defender.
    type: string
    exposed: true
    stored: true

  - name: enforcerID
    description: ID of the Defender.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: page
    description: Number assigned to each log in the increasing order.
    type: integer
    exposed: true
    stored: true

  - name: title
    description: Title of the log.
    type: string
    exposed: true
    stored: true
