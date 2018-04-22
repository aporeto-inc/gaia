# Model
model:
  rest_name: restapispec
  resource_name: restapispecs
  entity_name: RESTAPISpec
  package: squall
  description: |-
    RESTAPISpec descibes the REST APIs exposed by a service. These APIs
    can be associated with one or more services.
  get: true
  update: true
  delete: true
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  - '@propagated'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: external
    exposed: true
    subtype: exposed_api_list
    stored: true
