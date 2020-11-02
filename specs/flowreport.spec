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
  - '@zoned'
  - '@migratable'
  - '@flow'

# Indexes
indexes:
- - remotenamespace
  - timestamp
