# Model
model:
  rest_name: systemcall
  resource_name: systemcalls
  entity_name: SystemCall
  package: squall
  description: '[nodoc]'
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
