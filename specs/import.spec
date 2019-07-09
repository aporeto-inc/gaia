# Model
model:
  rest_name: import
  resource_name: import
  entity_name: Import
  package: squall
  group: core
  description: Imports an export of policies and related objects into the namespace.

# Attributes
attributes:
  v1:
  - name: data
    description: Data to import.
    type: ref
    exposed: true
    subtype: export
    required: true
    example_value:
      externalnetworks:
      - associatedTags:
        - ext:net=tcp
        description: Represents all TCP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-tcp
        protocols:
        - tcp
      - associatedTags:
        - ext:net=udp
        description: Represents all UDP traffic on any port
        entries:
        - 0.0.0.0/0
        name: all-udp
        protocols:
        - udp
      networkaccesspolicies:
      - action: Allow
        description: Allows all communication from pu to pu, tcp and udp
        logsEnabled: true
        name: allow-all-communication
        object:
        - - $identity=processingunit
        - - ext:net=tcp
        - - ext:net=udp
        subject:
        - - $identity=processingunit
    extensions:
      refMode: pointer

  - name: mode
    description: |-
      How to import the data: `ReplacePartial`, `Import` (default), or `Remove`. `ReplacePartial` 
      is deprecated. Use `Import` instead. While you can use `ReplacePartial` it will be interpreted 
      as `Import`.
    type: enum
    exposed: true
    allowed_choices:
    - ReplacePartial
    - Import
    - Remove
    default_value: Import
