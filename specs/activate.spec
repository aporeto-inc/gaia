# Model
model:
  rest_name: activate
  resource_name: activate
  entity_name: Activate
  package: vince
  description: Used to activate a pending account.

# Attributes
attributes:
  v1:
  - name: token
    description: Token contains the activation token.
    type: string
    exposed: true
    creation_only: true
    format: free
