# Model
model:
  rest_name: accessiblenamespace
  resource_name: accessiblenamespaces
  entity_name: AccessibleNamespace
  package: squall
  group: core/accessiblenamespace
  description: |-
    An Accessible Namespace represents a namespace that can be accessed by a given
    user.
  aliases:
  - accns

# Ordering
default_order:
- name

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: name
    description: Name of the namespace that is accessible.
    type: string
    exposed: true
    read_only: true
