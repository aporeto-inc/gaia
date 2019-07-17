# Model
model:
  rest_name: httpresourcespec
  resource_name: httpresourcespecs
  entity_name: HTTPResourceSpec
  package: squall
  group: policy/services
  description: |-
    Describes an HTTP resource exposed by one or more services.
  aliases:
  - httpresource
  - resource
  - httpspec
  get:
    description: Retrieves the HTTP resource with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the HTTP resource with the given ID.
  delete:
    description: Deletes the HTTP resource with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@base'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@propagated'
  - '@metadatable'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: A list of API endpoints that are exposed for the service.
    type: refList
    exposed: true
    subtype: endpoint
    stored: true
    extensions:
      refMode: pointer
