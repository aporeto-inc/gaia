# Model
model:
  rest_name: recipeoptions
  resource_name: recipeoptions
  entity_name: RecipeOptions
  package: ignis
  group: recipes
  description: Represents a Recipe Options.
  detached: true

# Attributes
attributes:
  v1:
  - name: appCrendentialFormat
    description: AppCrendentialFormat indicates the format of the AppCredential.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - JSON
    - YAML
    default_value: JSON
