# Model
model:
  rest_name: tokenscopepolicy
  resource_name: tokenscopepolicies
  entity_name: TokenScopePolicy
  package: squall
  description: |-
    The TokenScopePolicy defines a set of policies that allow customization of the
    authorization tokens issued by the Aporeto service. This allows Aporeto
    generated tokens to be used by external applications.
  aliases:
  - tsp
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
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: assignedScopes
    description: AssignedScopes is the the list of scopes that the policiy will assigns.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: subject
    description: |-
      Subject defines the selection criteria that this policy must match on identiy
      and scope request information.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    stored: true
    orderable: true
