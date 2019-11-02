# Model
model:
  rest_name: enforcerlog
  resource_name: enforcerlog
  entity_name: EnforcerLog
  package: ifrit
  group: core/enforcer
  description: |-
    An enforcer log represents the log collected by an enforcer. Each enforcer log
    can have partial or complete data. The collectionID is used to aggregate the
    multipart data into one.
  get:
    description: Retrieves the enforcerlog with the given ID.
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
      CollectionID is the ID of the enforcer log. CollectionID is used to
      aggregate the multipart data.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: data
    description: Represents the data collected by the enforcer.
    type: string
    exposed: true
    stored: true

  - name: enforcerID
    description: ID of the enforcer.
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
