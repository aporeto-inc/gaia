# Model
model:
  rest_name: enforcerprofilemappingpolicy
  resource_name: enforcerprofilemappingpolicies
  entity_name: EnforcerProfileMappingPolicy
  package: squall
  group: policy/enforcerconfig
  description: |-
    Allows you to map a defender profile to one or more defenders.
    The mapping can also be propagated down to the child namespace.
  aliases:
  - enfpols
  - enfpol
  - epm
  get:
    description: Retrieves the mapping with the given ID.
  update:
    description: Updates the mapping with the given ID.
  delete:
    description: Deletes the mapping with the given ID.
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
  - name: object
    description: |-
      The tag or tag expression that identifies the defender profile to 
      be mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      The tag or tag expression that identifies the defenders that should 
      implement the mapped profile.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: enforcer
  get:
    description: Returns the list of defenders affected by a defender profile mapping.

- rest_name: enforcerprofile
  get:
    description: |-
      Returns the list of defender profiles that a defender profile mapping
      matches.
