# Model
model:
  rest_name: networkaccesspolicy
  resource_name: networkaccesspolicies
  entity_name: NetworkAccessPolicy
  package: squall
  group: policy/networking
  description: |-
    Allows to define networking policies to allow or prevent processing units
    identitied by their tags to talk to other processing units or external services
    (also identified by their tags).
  aliases:
  - netpol
  - netpols
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
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@negatable-object'
  - '@negatable-subject'
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
    - Continue
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
    - Bidirectional
    default_value: Bidirectional
    orderable: true

  - name: encryptionEnabled
    description: EncryptionEnabled defines if the flow has to be encrypted.
    type: boolean
    exposed: true
    orderable: true

  - name: expirationTime
    description: If set the policy will be auto deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: logsEnabled
    description: LogsEnabled defines if the flow has to be logged.
    type: boolean
    exposed: true
    orderable: true

  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: observationEnabled
    description: If set to true, the flow will be in observation mode.
    type: boolean
    exposed: true
    orderable: true

  - name: observedTrafficAction
    description: |-
      If observationEnabled is set to true, this will be the final action taken on the
      packets.
    type: enum
    exposed: true
    allowed_choices:
    - Apply
    - Continue
    default_value: Continue
    orderable: true

  - name: ports
    description: Represents the ports and protocols this policy applies to.
    type: list
    exposed: true
    subtype: string
    orderable: true
    validations:
    - $protoports

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
    description: Returns the list of external networks affected by a network access
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
    description: Returns the list of Processing Units affected by a network access
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
    description: Returns the list of services affected by a network access policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object
