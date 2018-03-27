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
  - srvpols
  - srvpol
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'

# Attributes
attributes:
- name: object
  description: Object is the list of tags to use to find a enforcer profile.
  type: external
  exposed: true
  subtype: policies_list
  stored: true
  required: true
  example_value:
  - - a=a
    - b=b
  - - c=c

- name: subject
  description: Subject is the subject of the policy.
  type: external
  exposed: true
  subtype: policies_list
  stored: true
  required: true
  example_value:
  - - a=a
    - b=b
  - - c=c

# Relations
relations:
- rest_name: enforcerprofile
  descriptions:
    get: |-
      Returns the list of enforcer profiles that an enforcer profile mapping policy
      matches.
  get: true

- rest_name: enforcer
  descriptions:
    get: Returns the list of enforcers affected by an enforcer profile mapping policy.
  get: true
