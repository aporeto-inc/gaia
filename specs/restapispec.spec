# Model
model:
  rest_name: restapispec
  resource_name: restapispecs
  entity_name: RESTAPISpec
  package: squall
  description: |-
    RESTAPISpec descibes the REST APIs exposed by a service. These APIs
    can be associated with one or more services.
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@propagated'
  - '@metadatable'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: external
    exposed: true
    subtype: exposed_api_list
    stored: true
