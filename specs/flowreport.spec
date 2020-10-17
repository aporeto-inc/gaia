# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@flow'
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'
  
# Indexes
indexes:
- - remotenamespace
  - timestamp
