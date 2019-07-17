# Model
model:
  rest_name: activate
  resource_name: activate
  entity_name: Activate
  package: vince
  group: core/account
  description: Used to activate a pending account.

# Attributes
attributes:
  v1:
  - name: token
    description: Contains the activation token.
    type: string
    exposed: true
    creation_only: true
    example_value: 2BB3D52C-DE26-406A-8821-613F102282B0
