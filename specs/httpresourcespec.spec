# Model
model:
  rest_name: httpresourcespec
  resource_name: httpresourcespecs
  entity_name: HTTPResourceSpec
  package: squall
  group: policy/services
  description: |-
    HTTPResourceSpec descibes an HTTP resource exposed by a service. These APIs
    can be associated with one or more services.
  aliases:
  - httpresource
  - resource
  - httpspec
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
  - '@identifiable-stored'
  - '@named'
  - '@propagated'
  - '@metadatable'
  - '@zonable'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: refList
    exposed: true
    subtype: endpoint
    stored: true
    extensions:
      refMode: pointer
