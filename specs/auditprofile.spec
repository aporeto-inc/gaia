# Model
model:
  rest_name: auditprofile
  resource_name: auditprofiles
  entity_name: AuditProfile
  package: squall
  group: policy/audit
  description: |-
    AuditProfile is an audit policy that consists of a set of audit rules. An audit
    policy will determine that types of events that must be captured in the kernel.
  aliases:
  - ap
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: propagated
    description: Propagated indicates if the audit profile is propagated.
    type: boolean
    exposed: true
    stored: true

  - name: rules
    description: Rules is the list of audit policy rules associated with this policy.
    type: external
    exposed: true
    subtype: _audit_profile_rule_list
    stored: true
