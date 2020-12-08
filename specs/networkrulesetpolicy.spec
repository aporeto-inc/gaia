# Model
model:
  rest_name: networkrulesetpolicy
  resource_name: networkrulesetpolicies
  entity_name: NetworkRuleSetPolicy
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
    type: refList
    exposed: true
    subtype: networkrule
    extensions:
      refMode: pointer

  - name: ingressRules
    description: The set of ingress rules that comprise this rule set.
    type: refList
    exposed: true
    subtype: networkrule
    extensions:
      refMode: pointer

  - name: subject
    description: |-
      A tag or tag expression identifying the set of workloads where this policy
      applies to.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $tagsExpression
