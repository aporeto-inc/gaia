# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'
  - '@flow'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - remotenamespace
  - timestamp
- - sourceID
- - destinationID
