# Model
model:
  rest_name: rendertemplate
  resource_name: rendertemplates
  entity_name: RenderTemplate
  package: ignis
  group: core/workflow
  description: Cooks a template based some parameters.
  aliases:
  - cook
  - rtpl

# Attributes
attributes:
  v1:
  - name: output
    description: Holds the rendered template.
    type: string
    exposed: true

  - name: parameters
    description: Contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: template
    description: Template of the recipe.
    type: string
    exposed: true
