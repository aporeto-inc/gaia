# Model
model:
  rest_name: cloudnetworkruleset
  resource_name: cloudnetworkrulesets
  entity_name: CloudNetworkRuleSet
  package: yeul
  group: pcn/infrastructure
  description: |-
    A CloudNetworkRuleSet represents a set of cloud network security groups or
    firewall rules as they apply to the infrastructure.
  aliases:
  - crules
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Cloud network ruleset data.
    type: ref
    exposed: true
    subtype: cloudnetworkrulesetdata
    stored: true
    extensions:
      refMode: pointer
