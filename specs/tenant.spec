# Model
model:
  rest_name: tenant
  resource_name: tenants
  entity_name: Tenant
  package: karl
  group: core/tenant
  description: |-
    Can be used to create a tenant's namespace and API authorization policy to grant
    access.
  delete:
    description: Delete the tenant with the given namespace ID.
  extends:
  - '@identifiable-not-stored'

# Attributes
attributes:
  v1:
  - name: externalID
    description: The external ID of the tenant.
    type: string
    exposed: true
    required: true
    example_value: customer-123
    transient: true

  - name: name
    description: The name of the tenant.
    type: string
    exposed: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/]+$
    allowed_chars_message: must only contain alpha numerical characters, '-' or '_'
    example_value: acme
    transient: true
