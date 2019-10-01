# Model
model:
  rest_name: hit
  resource_name: hits
  entity_name: Hit
  package: minwu
  group: core
  description: This API allows to retrieve a generic hit counter for a given object.

# Attributes
attributes:
  v1:
  - name: name
    description: name of the counter.
    type: string
    exposed: true
    required: true
    default_value: counter

  - name: targetID
    description: The ID of the referenced object..
    type: string
    exposed: true

  - name: targetIdentity
    description: The identity of the referenced object.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: networkaccesspolicy
    validations:
    - $identity

  - name: value
    description: The value of the hit.
    type: integer
    exposed: true
    read_only: true
