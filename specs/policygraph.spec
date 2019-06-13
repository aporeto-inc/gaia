# Model
model:
  rest_name: policygraph
  resource_name: policygraphs
  entity_name: PolicyGraph
  package: yeul
  group: visualization/depmaps
  description: |-
    This api returns a data structure representing the policy graph of all selected
    processing units and their possible connectivity based on the current policies
    associated with the namespace. Users can define a selector of processing units
    for which theyare interestedor define the identity tags of a virtual processing
    unit that is not yetactivated.
  aliases:
  - polgraph

# Attributes
attributes:
  v1:
  - name: PUIdentity
    description: |-
      puIdentity is the set of tags that a future activated PU will have for which the
      user wants to evaluate policies and understand its connectivity options.
    type: list
    exposed: true
    subtype: string

  - name: dependencyMap
    description: |-
      The dependencyMap contains the output of the policy evalation, and it is the
      same
      type of dependency map as created by other APIs.
    type: ref
    exposed: true
    subtype: dependencymap
    extensions:
      refMode: pointer

  - name: selectors
    description: |-
      Selectors contains the tag expression that an a processing unit
      must match in order to evaluate policy for it.
    type: external
    exposed: true
    subtype: '[][]string'
    example_value:
    - - $identity=processingunit
    validations:
    - $tagsExpression
