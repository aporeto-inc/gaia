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

# Attributes
attributes:
  v1:
  - name: externalID
    description: The external ID of the tenant.
    type: string
    exposed: true
    transient: true

  - name: name
    description: The name of the tenant.
    type: string
    exposed: true
    transient: true
