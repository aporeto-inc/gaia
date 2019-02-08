# Model
model:
  rest_name: auditprofilemappingpolicy
  resource_name: auditprofilemappingpolicies
  entity_name: AuditProfileMappingPolicy
  package: squall
  group: policy/audit
  description: |-
    Defines an audit policy that determine the sets of enforcers that must implement
    a specific audit profile.
  aliases:
  - audpol
  - audpols
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
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: object
    description: |-
      Object of the policy is the selector of the audit profiles that must be applied
      based on this policy.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings

  - name: subject
    description: |-
      Subject of the policy is a selector of the enforcers that must implement the
      policy.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings

# Relations
relations:
- rest_name: enforcer
  get:
    description: Returns the list of enforcers that are affected by this poliy.

- rest_name: auditprofile
  get:
    description: Returns the list of audit profiles that are referred by this policy.
