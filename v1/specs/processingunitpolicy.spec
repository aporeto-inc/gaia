attributes:
- allowed_choices:
  - Delete
  - Enforce
  - LogCompliance
  - Reject
  - Snapshot
  - Stop
  description: Action determines the action to take while enforcing the isolation
    profile.
  exposed: true
  filterable: true
  name: action
  orderable: true
  stored: true
  type: enum
- description: IsolationProfileSelector are the profiles that must be applied when
    this policy matches. Only applies to Enforce and LogCompliance actions.
  exposed: true
  filterable: true
  name: isolationProfileSelector
  orderable: true
  stored: true
  subtype: policies_list
  type: external
- description: Subject defines the tag selectors that identitfy the processing units
    to which this policy applies.
  exposed: true
  filterable: true
  name: subject
  orderable: true
  stored: true
  subtype: policies_list
  type: external
model:
  aliases:
  - pup
  create: true
  delete: true
  get: true
  update: true
  description: A ProcessingUnitPolicies needs a better description.
  entity_name: ProcessingUnitPolicy
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'
  package: squall
  resource_name: processingunitpolicies
  rest_name: processingunitpolicy
