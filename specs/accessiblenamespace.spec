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
  extends:
  - '@namespaced'

# Ordering
default_order:
- name

# Indexes
indexes:
- - name
- - namespace
  - name

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of the namespace that is accessible.
    type: string
    exposed: true
    read_only: true
    example_value: 123-4343-54343

  - name: name
    description: Name of the namespace that is accessible.
    type: string
    exposed: true
    read_only: true
