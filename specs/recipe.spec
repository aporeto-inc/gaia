# Model
model:
  rest_name: recipe
  resource_name: recipes
  entity_name: Recipe
  package: ignis
  group: core/workflow
  description: A Recipe defines a list of steps to define a workflow.
  aliases:
  - rcp
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
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
  - '@propagated'
  - '@timeable'

# Indexes
indexes:
- - :unique
  - namespace
  - key

# Attributes
attributes:
  v1:
  - name: icon
    description: Icon contains a base64 image for the recipe.
    type: string
    exposed: true
    stored: true

  - name: key
    description: Key is the unique key of the recipe.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: label
    description: Label defines the recipe.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    default_value: magicpanda

  - name: longDescription
    description: LongDescription provides a long description of the recipe.
    type: string
    exposed: true
    stored: true

  - name: options
    description: Options of the recipe.
    type: ref
    exposed: true
    subtype: recipeoptions
    stored: true
    extensions:
      refMode: pointer

  - name: steps
    description: Steps contains all the steps with parameters to follow for the recipe.
    type: refList
    exposed: true
    subtype: uistep
    stored: true
    extensions:
      refMode: pointer

  - name: successfullMessage
    description: successfullMessage is presented if present and success.
    type: string
    exposed: true
    stored: true

  - name: template
    description: Template of the recipe to import.
    type: string
    exposed: true
    stored: true

  - name: templateHash
    description: templateHash is a hash of the template.
    type: string
    exposed: true
    stored: true
    read_only: true
