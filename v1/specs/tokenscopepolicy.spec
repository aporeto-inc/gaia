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
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'

# Attributes
attributes:
- name: assignedScopes
  description: AssignedScopes is the the list of scopes that the policiy will assigns.
  type: external
  exposed: true
  subtype: tags_list
  stored: true
  filterable: true
  orderable: true

- name: subject
  description: |-
    Subject defines the selection criteria that this policy must match on identiy
    and scope request information.
  type: external
  exposed: true
  subtype: policies_list
  stored: true
  filterable: true
  orderable: true
