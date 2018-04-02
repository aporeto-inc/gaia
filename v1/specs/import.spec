# Model
model:
  rest_name: import
  resource_name: import
  entity_name: Import
  package: squall
  description: Imports an export of policies and related objects into the namespace.

# Attributes
attributes:
  v1:
  - name: data
    description: The data to import.
    type: external
    exposed: true
    subtype: exported_data
    required: true
    example_value: previous output of export

  - name: mode
    description: How to import the data.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Append
    - Replace
    default_value: Replace
