# Model
model:
  rest_name: auditprofile
  resource_name: auditprofiles
  entity_name: AuditProfile
  package: squall
  group: policy/audit
  description: |-
    A set of audit rules that determine the types of events that must be captured in 
    the kernel.
  aliases:
  - ap
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the profile with the given ID.
  delete:
    description: Deletes the profile with the given ID.
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
    description: List of audit rules associated with this profile.
    type: external
    exposed: true
    subtype: _audit_profile_rule_list
    stored: true
