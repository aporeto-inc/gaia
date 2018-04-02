# Model
model:
  rest_name: isolationprofile
  resource_name: isolationprofiles
  entity_name: IsolationProfile
  package: squall
  description: An IsolationProfile needs documentation.
  aliases:
  - ip
  create: true
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: capabilitiesActions
    description: |-
      CapabilitiesActions identifies the capabilities that should be added or removed
      from the processing unit.
    type: external
    exposed: true
    subtype: cap_map
    stored: true
    filterable: true
    orderable: true

  - name: defaultSyscallAction
    description: |-
      DefaultAction is the default action applied to all syscalls of this profile.
      Default is "Allow".
    type: external
    exposed: true
    subtype: syscall_action
    stored: true
    filterable: true
    orderable: true

  - name: syscallRules
    description: |-
      SyscallRules is a list of syscall rules that identify actions for particular
      syscalls.
    type: external
    exposed: true
    subtype: syscall_rules
    stored: true
    filterable: true
    orderable: true

  - name: targetArchitectures
    description: |-
      TargetArchitectures is the target processor architectures where this profile can
      be applied. Default all.
    type: external
    exposed: true
    subtype: arch_list
    stored: true
    filterable: true
    orderable: true
