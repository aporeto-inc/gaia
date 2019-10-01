# Model
model:
  rest_name: hit
  resource_name: hits
  entity_name: Hit
  package: minwu
  group: core
  description: This API allows to retrieve a generic hit counter for a given object.
  get:
    description: Returns the hit with the given ID.
  delete:
    description: Resets the hit counter with the given ID.
  extends:
  - '@zoned'
  - '@base'
  - '@namespaced'
  - '@identifiable-stored'

# Indexes
indexes:
- - targetID
- - targetIdentity
- - targetIdentity
  - targetID
- - hash

# Attributes
attributes:
  v1:
  - name: hash
    description: Internal hash of the hit.
    type: integer
    stored: true

  - name: targetID
    description: The ID of the referenced object..
    type: string
    exposed: true
    stored: true

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
    description: The hit value.
    type: integer
    exposed: true
    stored: true
    read_only: true
