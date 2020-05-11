# Model
model:
  rest_name: networkaccesspolicy
  resource_name: networkaccesspolicies
  entity_name: NetworkAccessPolicy
  package: squall
  group: policy/networking
  description: |-
    Allows you to define network policies to allow or prevent processing units
    identified by their tags to talk to other processing units or external networks
    (also identified by their tags).
  aliases:
  - netpol
  - netpols
  get:
    description: Retrieves the policy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the policy with the given ID.
  delete:
    description: Deletes the policy with the given ID.
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
    description: |-
      Defines the action to apply to a flow.

      - `Allow`: allows the defined traffic.
      - `Reject`: rejects the defined traffic; useful in conjunction with an allow all
      policy.
      - `Continue`: neither allows or rejects the traffic; useful for applying another
      property to the traffic.
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
      Sets three different types of policies. `IncomingTraffic`: applies the policy to
      all
      processing units that match the `object` and allows them to *accept* connections
      from
      processing units or external networks that match the `subject`.
      `OutgoingTraffic`: applies
      the policy to all processing units that match the `subject` and allows them to
      *initiate*
      connections with processing units or external networks that match the `object`.
      `Bidirectional` (default): applies the policy to all processing units that match
      the `object`
      and allows them to *accept* connections from processing units that match the
      `subject`.
      Also applies the policy to all processing units that match the `subject` and
      allows them
      to *initiate* connections with processing units that match the `object`.
    type: enum
    exposed: true
    allowed_choices:
    - OutgoingTraffic
    - IncomingTraffic
    - Bidirectional
    default_value: Bidirectional
    orderable: true

  - name: encryptionEnabled
    description: |-
      Defines if the flow has to be encrypted. This property is deprecated and have no
      incidence.
    type: boolean
    exposed: true
    deprecated: true
    orderable: true

  - name: expirationTime
    description: If set the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: logsEnabled
    description: |-
      If `true`, the relevant flows are logged and available from the Aporeto control
      plane.
      Under some advanced scenarios you may wish to set this to `false`, such as to
      save space or
      improve performance.
    type: boolean
    exposed: true
    orderable: true

  - name: object
    description: A tag or tag expression identifying the object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: observationEnabled
    description: If set to `true`, the flow will be in observation mode.
    type: boolean
    exposed: true
    orderable: true

  - name: observedTrafficAction
    description: |-
      If `observationEnabled` is set to `true`, this defines the final action taken
      on the packets: `Apply` or `Continue` (default).
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
    description: A tag or tag expression identifying the subject of the policy.
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
    description: Returns the list of external networks affected by a network policy.
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
    description: Returns the list of processing units affected by a network policy.
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
    description: Returns the list of services affected by a network policy.
    parameters:
      entries:
      - name: mode
        description: Matching mode.
        type: enum
        allowed_choices:
        - subject
        - object
        default_value: object
