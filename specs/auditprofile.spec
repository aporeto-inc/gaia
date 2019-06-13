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
  - '@zoned'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: rules
    description: Rules is the list of audit policy rules associated with this policy.
    type: external
    exposed: true
    subtype: _audit_profile_rule_list
    stored: true
