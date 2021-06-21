# Model
model:
  rest_name: servicepublication
  resource_name: servicepublications
  entity_name: ServicePublication
  package: squall
  group: core/service
  description: |-
    Encapsulates a service object that is ought to be published so it can be used
    in a sibling namespace.

# Attributes
attributes:
  v1:
  - name: service
    description: The service object that will be published.
    type: ref
    exposed: true
    required: true
    subtype: service
    extensions:
      refMode: pointer
    example_value:
      hosts: ["localhost"]
      name: "referenced-service"
      port: 443
      exposedPort: 443
      propagate: true
