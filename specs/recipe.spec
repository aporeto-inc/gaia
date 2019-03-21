# Model
model:
  rest_name: recipe
  resource_name: recipes
  entity_name: Recipe
  package: ignis
  group: workflow
  description: A Recipe defines a list of steps to define a workflow.
  aliases:
  - rcp
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@metadatable'
  - '@disabled'
  - '@propagated'

# Attributes
attributes:
  v1:
  - name: category
    description: Category indicates the category of the recipe.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Official
    - Users
    - Network
    - Service

  - name: icon
    description: Icon contains a base64 image for the app.
    type: string
    exposed: true
    stored: true

  - name: steps
    description: Steps contains all the steps with parameters to follow for the recipe.
    type: refList
    exposed: true
    subtype: uistep
    stored: true
    extensions:
      refMode: pointer

  - name: template
    description: Template of the recipe to import.
    type: string
    exposed: true
    stored: true
