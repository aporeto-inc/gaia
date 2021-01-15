# Model
model:
  rest_name: cloudnetworkrulesetdata
  resource_name: cloudnetworkrulesetdata
  entity_name: CloudNetworkRuleSetData
  package: pcn
  group: prisma/infrastructure
  description: Parameters associated with a cloud network rule set.
  detached: true

# Attributes
attributes:
  v1:
  - name: incomingRules
    description: |-
      The set of rules to apply to incoming traffic (traffic coming to the Processing
      Unit matching the subject).
    type: refList
    exposed: true
    subtype: cloudnetworkrule
    stored: true
    extensions:
      refMode: pointer

  - name: outgoingRules
    description: |-
      The set of rules to apply to outgoing traffic (traffic coming from the
      Processing Unit matching the subject).
    type: refList
    exposed: true
    subtype: cloudnetworkrule
    stored: true
    extensions:
      refMode: pointer

  - name: subject
    description: |-
      A tag expression identifying used to match processing units to which this policy
      applies to.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression

  - name: type
    description: Type identifies if this is a security group rule set or ACL.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - SecurityGroup
    - ACL
    example_value: SecurityGroup
