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
    description: |-
      How to import the data. ReplacePartial is deprecated and should be Import. Right
      now the API considers both equivalent.
    type: enum
    exposed: true
    allowed_choices:
    - ReplacePartial
    - Import
    - Remove
    default_value: Import
