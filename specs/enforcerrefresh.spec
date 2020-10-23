# Model
model:
  rest_name: enforcerrefresh
  resource_name: enforcerrefreshes
  entity_name: EnforcerRefresh
  package: gaga
  group: core/policy
  description: |-
    Sent to enforcers when a poke has been triggered using the
    parameter `?notify=true`. This is used to notify an enforcer of an
    external change on the processing unit that must be processed.

# Attributes
attributes:
  v1:
  - name: ID
    description: Contains the ID of the target enforcer.
    type: string
    exposed: true
    read_only: true
    getter: true
    setter: true
    identifier: true

  - name: debug
    description: Set the debug information collected by the enforcer.
    type: enum
    exposed: true
    allowed_choices:
    - Counters
    - Logs
    - Packets
    - PUState
    - Pcap
    - CoreDump
    default_value: Counters
    omit_empty: true

  - name: debugID
    description: Can be used to correlate with a DebugBundle.
    type: string
    exposed: true
    omit_empty: true

  - name: debugPcapFilter
    description: Packet capture filter, syntax varying by platform.
    type: string
    exposed: true
    omit_empty: true

  - name: debugProcessingUnitID
    description: Isolates debug information to a given processing unit, where possible.
    type: string
    exposed: true
    omit_empty: true

  - name: namespace
    description: Contains the original namespace of the enforcer.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: upgradeEnforcerSelector
    description: Request an upgrade on the enforcer matching the following tag expression.
    type: external
    exposed: true
    subtype: '[][]string'
    read_only: true
    autogenerated: true
    example_value:
    - - $namespace=/a/b
    omit_empty: true
    validations:
    - $tagsExpression

  - name: upgradeRecursive
    description: Indicates if the upgrade should be done recursively in all namespaces.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true

  - name: upgradeTargetVersion
    description: Defines the version to upgrade the enforcers.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    omit_empty: true
    validations:
    - $semver
