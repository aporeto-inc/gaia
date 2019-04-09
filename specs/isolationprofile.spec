# Model
model:
  rest_name: isolationprofile
  resource_name: isolationprofiles
  entity_name: IsolationProfile
  package: squall
  group: policy/processingunits
  description: An IsolationProfile needs documentation.
  aliases:
  - ip
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'
  - '@propagated'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: capabilitiesActions
    description: |-
      CapabilitiesActions identifies the capabilities that should be added or removed
      from the processing unit.
    type: external
    exposed: true
    subtype: _cap_map
    stored: true
    orderable: true

  - name: defaultSyscallAction
    description: |-
      DefaultAction is the default action applied to all syscalls of this profile.
      Default is "Allow".
    type: external
    exposed: true
    subtype: _syscall_action
    stored: true

  - name: syscallRules
    description: |-
      SyscallRules is a list of syscall rules that identify actions for particular
      syscalls.
    type: external
    exposed: true
    subtype: _syscall_rules
    stored: true
    orderable: true

  - name: targetArchitectures
    description: |-
      TargetArchitectures is the target processor architectures where this profile can
      be applied. Default all.
    type: external
    exposed: true
    subtype: _arch_list
    stored: true
    orderable: true
