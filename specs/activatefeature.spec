# Model
model:
  rest_name: activatefeature
  resource_name: activatefeatures
  entity_name: ActivateFeature
  package: karl
  group: core/tenant
  description: Activates one or multiple features per tenant's Prisma ID.

# Attributes
attributes:
  v1:
  - name: feature
    description: Name of the feature to activate for the specified tenant.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - NetworkEffectivePermission
    - NetworkSecurity
    example_value: NetworkSecurity
