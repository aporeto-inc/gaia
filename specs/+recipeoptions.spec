# Model
model:
  rest_name: recipeoptions
  resource_name: recipeoptions
  entity_name: RecipeOptions
  package: ignis
  group: core/workflow
  description: Represents recipe options.
  detached: true

# Attributes
attributes:
  v1:
  - name: appCrendentialFormat
    description: Indicates the format of the app credential.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - JSON
    - YAML
    default_value: JSON
