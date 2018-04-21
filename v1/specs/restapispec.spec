# Model
model:
  rest_name: restapispec
  resource_name: restapispecs
  entity_name: RESTAPISpec
  package: squall
  description: |-
    RESTAPISpec descibes a L4/L7 service and the corresponding implementation. It
    allows users to define their services, the APIs that they expose, the
    implementation of the service. These definitions can be used by network policy
    in order to define advanced controls based on the APIs.
  get: true
  update: true
  delete: true
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: endpoints
    description: EndPoints is a list of API endpoints that are exposed for the service.
    type: external
    exposed: true
    subtype: exposed_api_list
    stored: true
