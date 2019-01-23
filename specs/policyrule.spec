# Model
model:
  rest_name: policyrule
  resource_name: policyrules
  entity_name: PolicyRule
  package: squall
  description: |-
    PolicyRule is an internal policy resolution API. Services can use this API to
    retrieve a policy resolution.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-nopk-nostored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: action
    description: Action defines set of actions that must be enforced when a dependency
      is met.
    type: external
    exposed: true
    subtype: map_of_string_of_maps_of_string_of_objects

  - name: enforcerProfiles
    description: EnforcerProfiles provides the information about the server profile.
    type: refList
    exposed: true
    subtype: enforcerprofile
    extensions:
      refMode: pointer

  - name: externalNetworks
    description: Policy target networks.
    type: refList
    exposed: true
    subtype: externalnetwork
    extensions:
      refMode: pointer

  - name: externalServices
    description: Policy target networks.
    type: refList
    exposed: true
    subtype: externalservice
    deprecated: true
    extensions:
      refMode: pointer

  - name: filePaths
    description: Policy target file paths.
    type: refList
    exposed: true
    subtype: filepath
    extensions:
      refMode: pointer

  - name: isolationProfiles
    description: IsolationProfiles are the isolation profiles of the rule.
    type: refList
    exposed: true
    subtype: isolationprofile
    extensions:
      refMode: pointer

  - name: namespaces
    description: Policy target namespaces.
    type: refList
    exposed: true
    subtype: namespace
    extensions:
      refMode: pointer

  - name: policyNamespace
    description: PolicyNamespace is the namespace of the policy that created this
      rule.
    type: string
    exposed: true

  - name: policyUpdateTime
    description: Last time the policy was updated.
    type: time
    exposed: true

  - name: propagated
    description: Propagated indicates if the policy is propagated.
    type: boolean
    exposed: true

  - name: relation
    description: |-
      Relation describes the required operation to be performed between subjects and
      objects.
    type: list
    exposed: true
    subtype: string

  - name: services
    description: Services provides the services of this policy rule.
    type: refList
    exposed: true
    subtype: service
    extensions:
      refMode: pointer

  - name: tagClauses
    description: Policy target tags.
    type: external
    exposed: true
    subtype: list_of_list_of_strings
