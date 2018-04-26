# Model
model:
  rest_name: policyrule
  resource_name: policyrules
  entity_name: PolicyRule
  package: squall
  description: |-
    PolicyRule is an internal policy resolution API. Services can use this API to
    retrieve a policy resolution.
  get: true
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
    subtype: actions_list

  - name: enforcerProfiles
    description: EnforcerProfiles provides the information about the server profile.
    type: external
    exposed: true
    subtype: enforcerprofiles_list

  - name: externalServices
    description: Policy target networks.
    type: external
    exposed: true
    subtype: network_entities

  - name: filePaths
    description: Policy target file paths.
    type: external
    exposed: true
    subtype: file_entities

  - name: isolationProfiles
    description: IsolationProfiles are the isolation profiles of the rule.
    type: external
    exposed: true
    subtype: isolation_profile_entities

  - name: namespaces
    description: Policy target namespaces.
    type: external
    exposed: true
    subtype: namespace_entities

  - name: passthroughExternalServices
    description: |-
      List of external services the policy mandate to pass through before reaching the
      destination.
    type: external
    exposed: true
    subtype: network_entities

  - name: policyNamespace
    description: PolicyNamespace is the namespace of the policy that created this
      rule.
    type: string
    exposed: true

  - name: propagated
    description: Propagated indicates if the policy is propagated.
    type: boolean
    exposed: true

  - name: relation
    description: |-
      Relation describes the required operation to be performed between subjects and
      objects.
    type: external
    exposed: true
    subtype: relations_list

  - name: services
    description: Services provides the services of this policy rule.
    type: external
    exposed: true
    subtype: api_services_entities

  - name: tagClauses
    description: Policy target tags.
    type: external
    exposed: true
    subtype: target_tags
