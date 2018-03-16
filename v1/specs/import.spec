attributes:
- description: The data to import.
  exposed: true
  name: data
  required: true
  subtype: exported_data
  type: external
- allowed_choices:
  - Append
  - Replace
  default_value: Replace
  description: How to import the data.
  exposed: true
  name: mode
  required: true
  type: enum
model:
  description: Imports an export of policies and related objects into the namespace.
  entity_name: Import
  package: squall
  resource_name: import
  rest_name: import
