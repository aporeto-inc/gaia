# Model
model:
  rest_name: cachedflowrecord
  resource_name: cachedflowrecords
  entity_name: CachedFlowRecord
  package: zack
  group: policy/networking
  description: Post a new cached flow record.
  extends:
  - '@flow'
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'
  validations:
  - $cachedflowrecord

# Attributes
attributes:
  v1:
  - name: isLocalDestinationID
    description: Indicates if the destination endpoint is an enforcer-local processing unit.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ai

  - name: isLocalSourceID
    description: Indicates if the source endpoint is an enforcer-local processing unit.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: aj
