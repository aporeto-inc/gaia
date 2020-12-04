# Model
model:
  rest_name: networkruleset
  resource_name: networkrulesets
  entity_name: NetworkRuleSet
  package: squall
  group: policy/networking
  description: |-
    Allows you to define network rule sets to allow or prevent processing units
    identified by their tags to talk to other processing units or external networks
    (also identified by their tags).
  aliases:
  - netruleset
  - netset
  get:
    description: Retrieves the policy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the policy with the given ID.
  delete:
    description: Deletes the policy with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: egressRules
    description: The set of egress rules that comprise this rule set.
    type: list
    exposed: true
    subtype: networkrule

  - name: ingressRules
    description: The set of ingress rules that comprise this rule set.
    type: list
    exposed: true
    subtype: networkrule

  - name: selector
    description: |-
      A tag or tag expression identifying the set of workloads where this policy
      applies to.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression
