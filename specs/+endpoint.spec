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
    description: The scopes authorized to access the API.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    orderable: true

  - name: methods
    description: Methods exposed to access the API.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $httpMethods

  - name: public
    description: If `true`, the API is public.
    type: boolean
    exposed: true
    stored: true

  - name: scopes
    description: Use `allowedScopes`.
    type: list
    exposed: true
    subtype: string
    read_only: true
    deprecated: true
    orderable: true
