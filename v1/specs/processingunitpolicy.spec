# Model
model:
  rest_name: processingunitpolicy
  resource_name: processingunitpolicies
  entity_name: ProcessingUnitPolicy
  package: squall
  description: A ProcessingUnitPolicies needs a better description.
  aliases:
  - pup
  create: true
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
  filterable: true
  orderable: true

- name: isolationProfileSelector
  description: |-
    IsolationProfileSelector are the profiles that must be applied when this policy
    matches. Only applies to Enforce and LogCompliance actions.
  type: external
  exposed: true
  subtype: policies_list
  stored: true

- name: subject
  description: |-
    Subject defines the tag selectors that identitfy the processing units to which
    this policy applies.
  type: external
  exposed: true
  subtype: policies_list
  stored: true
