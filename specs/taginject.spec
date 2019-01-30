# Model
model:
  rest_name: taginject
  resource_name: taginjects
  entity_name: TagInject
  package: tagle
  description: This internal api is used to inject a new tag in the database.
  private: true

# Attributes
attributes:
  v1:
  - name: addedTags
    description: List of tags to be added.
    type: external
    exposed: true
    subtype: map_of_string_of_integers

  - name: removedTags
    description: List of tags to be removed.
    type: external
    exposed: true
    subtype: map_of_string_of_integers

  - name: targetNamespace
    description: List of tags to inject.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace
