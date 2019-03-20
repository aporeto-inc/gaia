# Model
model:
  rest_name: recipecook
  resource_name: recipecooks
  entity_name: RecipeCook
  package: ignis
  group: workflow
  description: A RecipeCook cooks a recipe based on parameters.
  aliases:
  - rcpck
  - cook

# Attributes
attributes:
  v1:
  - name: parameters
    description: Parameters contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: template
    description: Template of the recipe.
    type: string
    exposed: true
