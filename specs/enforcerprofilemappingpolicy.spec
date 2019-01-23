# Model
model:
  rest_name: enforcerprofilemappingpolicy
  resource_name: enforcerprofilemappingpolicies
  entity_name: EnforcerProfileMappingPolicy
  package: squall
  description: |-
    A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used
    by and Aporeto Agent based on the Enforcer that have been used during the
    registration. The policy can also be propagated down to the child namespace.
  aliases:
  - enfpols
  - enfpol
  - epm
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: object
    description: Object is the list of tags to use to find a enforcer profile.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c

  - name: subject
    description: Subject is the subject of the policy.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    stored: true
    example_value:
    - - a=a
      - b=b
    - - c=c

# Relations
relations:
- rest_name: enforcerprofile
  get:
    description: |-
      Returns the list of enforcer profiles that an enforcer profile mapping policy
      matches.

- rest_name: enforcer
  get:
    description: Returns the list of enforcers affected by an enforcer profile mapping
      policy.
