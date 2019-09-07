# Model
model:
  rest_name: authn
  resource_name: authn
  entity_name: Authn
  package: midgard
  group: core/authentication
  description: |-
    Verifies if the given token is valid or not. If it is valid it will
    return the claims of the token.

# Attributes
attributes:
  v1:
  - name: claims
    description: The claims in the token.
    type: external
    exposed: true
    subtype: _claims
    read_only: true
    autogenerated: true

  - name: token
    description: The token to verify. This is only used is a POST request is used.
    type: string
    exposed: true
