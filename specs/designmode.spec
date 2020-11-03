# Model
model:
  rest_name: designmode
  resource_name: designmode
  entity_name: DesignMode
  package: yuna
  group: core
  description: |-
    When design mode is enabled, all flows are accepted. Flows which do not match an
    existing network policy will be represented by a dotted line in your Platform
    view.
  get:
    description: Retrieve the design mode with the given import reference ID.
  delete:
    description: Remove the design mode assets with the given import reference ID.
  extends:
  - '@identifiable-not-stored'
  - '@propagated'
