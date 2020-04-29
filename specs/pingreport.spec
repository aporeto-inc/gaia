# Model
model:
  rest_name: pingreport
  resource_name: pingreports
  entity_name: PingReport
  package: zack
  group: core/enforcer
  description: Post a new pu diagnostics report.

# Attributes
attributes:
  v1:
  - name: RTT
    description: Time taken for a single request-response to complete.
    type: string
    exposed: true

  - name: RXFourTuple
    description: Receiver four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true

  - name: TXController
    description: Controller of the transmitter.
    type: string
    exposed: true

  - name: TXFourTuple
    description: Transmitter four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true

  - name: TXType
    description: Type of the transmitter.
    type: string
    exposed: true

  - name: applicationListening
    description: If true, application responded to the request.
    type: boolean
    exposed: true

  - name: destinationID
    description: ID of the destination processing unit.
    type: string
    exposed: true

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: enforcerVersion
    description: Semantic version of the enforcer.
    type: string
    exposed: true

  - name: iterationID
    description: IterationID unique to a single ping request-response.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: namespace
    description: Namespace of the reporting processing unit.
    type: string
    exposed: true

  - name: payloadSize
    description: Size of the payload attached to the packet.
    type: integer
    exposed: true

  - name: pingID
    description: PingID unique to a single ping control.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: policyAction
    description: Action of the policy.
    type: string
    exposed: true

  - name: policyID
    description: ID of the policy.
    type: string
    exposed: true

  - name: protocol
    description: Protocol used for the communication.
    type: integer
    exposed: true

  - name: request
    description: Request represents the current request.
    type: integer
    exposed: true

  - name: seqNumMatching
    description: If Equal, transmitter sequence number matches the receiver sequence
      number.
    type: enum
    exposed: true
    allowed_choices:
    - Equal
    - Unequal
    - Noop
    default_value: Noop

  - name: serviceType
    description: Type of the service.
    type: string
    exposed: true

  - name: sourceID
    description: ID of the source PU.
    type: string
    exposed: true

  - name: stage
    description: Current stage when this report has been generated.
    type: string
    exposed: true

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
