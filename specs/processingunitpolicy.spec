# Model
model:
  rest_name: processingunitpolicy
  resource_name: processingunitpolicies
  entity_name: ProcessingUnitPolicy
  package: squall
  group: policy/processingunits
  description: Allows you to map isolation profiles to processing units.
  aliases:
  - pup
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
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: Action determines the action to take while enforcing the isolation
      profile.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Delete
    - Enforce
    - LogCompliance
    - Reject
    - Snapshot
    - Stop
    orderable: true

  - name: isolationProfileSelector
    description: |-
      The isolation profiles to be mapped. Only applies to `Enforce` and
      `LogCompliance` actions.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      A tag or tag expression identifying the processing unit(s) to which the
      isolation profile should be mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: isolationprofile
  get:
    description: Returns the list of isolation profiles associated with the mapping.

- rest_name: processingunit
  get:
    description: Returns the list of processing units referenced by the mapping.
