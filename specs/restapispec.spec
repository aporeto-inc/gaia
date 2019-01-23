# Model
model:
  rest_name: restapispec
  resource_name: restapispecs
  entity_name: RESTAPISpec
  package: squall
  description: This is deprecated. Use HTTPResourceSpec instead.
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - archived
  - - namespace
    - normalizedTags
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@propagated'
  - '@metadatable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: external
    exposed: true
    subtype: _exposed_api_list
    stored: true

  - name: migrated
    description: Migrated indicated if the object has been migrated to an HTTPResourceSpec.
    type: boolean
    stored: true
