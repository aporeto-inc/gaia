# Model
model:
  rest_name: servicepublication
  resource_name: servicepublications
  entity_name: ServicePublication
  package: squall
  group: core/service
  description: TODO.
  update:
    description: TODO.
  delete:
    description: TODO.

# Attributes
attributes:
  v1:
  - name: service
    description: TODO.
    type: ref
    exposed: true
    subtype: service
    extensions:
      refMode: pointer
