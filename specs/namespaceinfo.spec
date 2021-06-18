# Model
model:
  rest_name: namespaceinfo
  resource_name: namespaceinfo
  entity_name: NamespaceInfo
  package: squall
  group: core/namespace
  description: Returns the information of the specified namespace.
  extends:
  - '@described'
  - '@named'

# Attributes
attributes:
  v1:
  - name: PUIncomingTrafficAction
    description: The processing unit action for incoming traffic for the namespace.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit

  - name: PUOutgoingTrafficAction
    description: The processing unit action for outgoing traffic for the namespace.
    type: enum
    exposed: true
    read_only: true
    allowed_choices:
    - Allow
    - Reject
    - Inherit

  - name: prefixes
    description: List of tag prefixes that will be used to suggest policies.
    type: list
    exposed: true
    subtype: string
    read_only: true

  - name: protected
    description: Defines if the object is protected.
    type: boolean
    exposed: true
    stored: true
    getter: true
    setter: true
    orderable: true
