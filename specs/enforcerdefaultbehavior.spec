# Model
model:
  rest_name: enforcerdefaultbehavior
  resource_name: enforcerdefaultbehavior
  entity_name: EnforcerDefaultBehavior
  package: squall
  group: core/namespace
  description: Returns the default enforcer behavior of the specified namespace.

# Attributes
attributes:
  v1:
  - name: behavior
    description: The default enforcer behavior for the namespace.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit
