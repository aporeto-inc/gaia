# Model
model:
  rest_name: policyrule
  resource_name: policyrules
  entity_name: PolicyRule
  package: squall
  group: core/policy
  description: |-
    PolicyRule is an internal policy resolution API. Services can use this API to
    retrieve a policy resolution.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-not-stored'
  - '@named'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: Action defines set of actions that must be enforced when a dependency
      is met.
    type: external
    exposed: true
    subtype: map[string]map[string]interface{}

  - name: auditProfiles
    description: AuditProfiles provides the audit profiles that must be applied.
    type: refList
    exposed: true
    subtype: auditprofile
    omit_empty: true

  - name: enforcerProfiles
    description: EnforcerProfiles provides the information about the server profile.
    type: refList
    exposed: true
    subtype: enforcerprofile
    omit_empty: true
    extensions:
      noInit: true

  - name: externalNetworks
    description: Policy target networks.
    type: refList
    exposed: true
    subtype: externalnetwork
    omit_empty: true
    extensions:
      noInit: true

  - name: filePaths
    description: Policy target file paths.
    type: refList
    exposed: true
    subtype: filepath
    omit_empty: true
    extensions:
      noInit: true

  - name: hostServices
    description: HostServices provides the list of host services that must be instantiated.
    type: refList
    exposed: true
    subtype: hostservice
    omit_empty: true

  - name: isolationProfiles
    description: IsolationProfiles are the isolation profiles of the rule.
    type: refList
    exposed: true
    subtype: isolationprofile
    omit_empty: true
    extensions:
      noInit: true

  - name: namespaces
    description: Policy target namespaces.
    type: refList
    exposed: true
    subtype: namespace
    omit_empty: true
    extensions:
      noInit: true

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
    omit_empty: true
    extensions:
      noInit: true

  - name: tagClauses
    description: Policy target tags.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $tagsExpression
