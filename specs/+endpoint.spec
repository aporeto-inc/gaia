# Model
model:
  rest_name: endpoint
  resource_name: endpoints
  entity_name: Endpoint
  package: squall
  group: policy/services
  description: Represents an HTTP endpoint.
  detached: true

# Attributes
attributes:
  v1:
  - name: URI
    description: URI of the exposed API.
    type: string
    exposed: true
    stored: true

  - name: allowedScopes
    description: AllowedScopes authorized to access the API.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    orderable: true
    validations:
    - $tagsExpression

  - name: methods
    description: methods exposed to access the API.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $httpMethods

  - name: public
    description: public defines if the api is public or not.
    type: boolean
    exposed: true
    stored: true

  - name: scopes
    description: Scopes is deprecated.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    deprecated: true
    orderable: true
