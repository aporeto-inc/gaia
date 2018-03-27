# Model
model:
  rest_name: policy
  resource_name: policies
  entity_name: Policy
  package: squall
  description: Policy represents the policy primitive used by all aporeto policies.
  get: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'

# Attributes
attributes:
- name: action
  description: Action defines set of actions that must be enforced when a dependency
    is met.
  type: external
  exposed: true
  subtype: actions_list
  stored: true

- name: allObjectTags
  description: This is a set of all object tags for matching in the DB.
  type: external
  subtype: tags_list
  stored: true

- name: allSubjectTags
  description: This is a set of all subject tags for matching in the DB.
  type: external
  subtype: tags_list
  stored: true

- name: object
  description: |-
    Object represents set of entities that another entity depends on. As subjects,
    objects are identified as logical operations on tags when a policy is defined.
  type: external
  exposed: true
  subtype: policies_list
  stored: true

- name: relation
  description: |-
    Relation describes the required operation to be performed between subjects and
    objects.
  type: external
  exposed: true
  subtype: relations_list
  stored: true

- name: subject
  description: |-
    Subject represent sets of entities that will have a dependency other entities.
    Subjects are defined as logical operations on tags. Logical operations can
    includes AND/OR.
  type: external
  exposed: true
  subtype: policies_list
  stored: true

- name: type
  description: Type of the policy.
  type: enum
  exposed: true
  stored: true
  creation_only: true
  allowed_choices:
  - APIAuthorization
  - EnforcerProfile
  - File
  - Hook
  - NamespaceMapping
  - Network
  - ProcessingUnit
  - Quota
  - Syscall
  - TokenScope
  filterable: true
  primary_key: true
