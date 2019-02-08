# Model
model:
  rest_name: hostservicemappingpolicy
  resource_name: hostservicemappingpolicies
  entity_name: HostServiceMappingPolicy
  package: squall
  group: policy/hosts
  description: |-
    Defines a host service mapping policy that provides the relation between
    enforcers and host services that they must implement.
  aliases:
  - hostsrvmappol
  - hostsrvmappols
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
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: object
    description: |-
      Object of the policy is the selector for the host services that must be applied
      to this enforcer.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    orderable: true

  - name: subject
    description: |-
      Subject of the policy is the selector of the enforcers that the list of host
      services must apply to.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    orderable: true

# Relations
relations:
- rest_name: enforcer
  get:
    description: Returns the list of enforcers that are affected by this poliy.

- rest_name: hostservice
  get:
    description: Returns the list of host services that are referred by this policy.
