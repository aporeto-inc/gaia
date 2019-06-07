# Model
model:
  rest_name: infrastructurepolicy
  resource_name: infrastructurepolicies
  entity_name: InfrastructurePolicy
  package: squall
  group: policy/networking
  description: |-
    Infrastructure policies capture the network access rules of the underlying
    infrastructure and can be used to model cloud security groups, firewalls or
    other ACL based mechanisms. They are not used in the identity-based network
    authorization of Aporeto, but they can affect traffic flows in the underlying
    infrastructure.
  aliases:
  - infrapol
  - infrapols
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
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
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: action
    description: Action defines the action to apply to a flow.
    type: enum
    exposed: true
    allowed_choices:
    - Allow
    - Reject
    default_value: Allow
    orderable: true

  - name: applyPolicyMode
    description: |-
      applyPolicyMode determines if the policy has to be applied to the
      outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
      Default is both directions.
    type: enum
    exposed: true
    allowed_choices:
    - OutgoingTraffic
    - IncomingTraffic
    default_value: OutgoingTraffic
    orderable: true

  - name: expirationTime
    description: If set the policy will be auto deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: Subject of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: externalnetwork
  get:
    description: Returns the list of external networks affected by an infrastructure
      policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: processingunit
  get:
    description: Returns the list of Processing Units affected by an infrastructure
      policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object

- rest_name: service
  get:
    description: Returns the list of services affected by an infrastructure policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object
