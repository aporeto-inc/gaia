attributes:
- description: Object is the list of tags to use to find a enforcer profile.
  exposed: true
  name: object
  required: true
  stored: true
  subtype: policies_list
  type: external
- description: Subject is the subject of the policy.
  exposed: true
  name: subject
  required: true
  stored: true
  subtype: policies_list
  type: external
children:
- rest_name: enforcerprofile
  get: true
- rest_name: enforcer
  get: true
model:
  aliases:
  - srvpols
  - srvpol
  delete: true
  get: true
  update: true
  description: A Enforcer Profile Mapping Policy will tell what Enforcer Profile should
    be used by and Aporeto Agent based on the Enforcer that have been used during
    the registration. The policy can also be propagated down to the child namespace.
  entity_name: EnforcerProfileMappingPolicy
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  package: squall
  resource_name: enforcerprofilemappingpolicies
  rest_name: enforcerprofilemappingpolicy
