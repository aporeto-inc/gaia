# Model
model:
  rest_name: defaultenforcerversion
  resource_name: defaultenforcerversion
  entity_name: DefaultEnforcerVersion
  package: squall
  group: core/namespace
  description: Returns the default enforcer version of the specified namespace.

# Attributes
attributes:
  v1:
  - name: defaultVersion
    description: The default enforcer version for the namespace.
    type: string
    exposed: true
    validations:
    - $semver
