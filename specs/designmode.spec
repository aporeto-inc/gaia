# Model
model:
  rest_name: designmode
  resource_name: designmodes
  entity_name: DesignMode
  package: yuna
  group: core
  description: |-
    When design mode is enabled, all flows are accepted. Flows which do not match an
    existing network policy will be represented by a dotted line in your Platform
    view.
  delete:
    description: Remove the design mode assests with the specified namespace ID.
  extends:
  - '@identifiable-not-stored'
  - '@propagated'
