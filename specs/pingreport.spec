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
  - name: ID
    description: ID unique to a single request and response report.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

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

  - name: exchange
    description: Exchange represents request/response this report has been generated.
    type: string
    exposed: true

  - name: iteration
    description: Request represents the iteration number.
    type: integer
    exposed: true

  - name: latency
    description: Time taken for a single request to complete.
    type: string
    exposed: true

  - name: namespace
    description: Namespace of the reporting processing unit.
    type: string
    exposed: true

  - name: payloadSize
    description: Size of the payload attached to the packet.
    type: integer
    exposed: true

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

  - name: rxFourTuple
    description: Receiver four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true

  - name: seqNumMatching
    description: If true, transmitter sequence number matches the receiver sequence
      number.
    type: string
    exposed: true

  - name: serviceType
    description: Type of the service.
    type: string
    exposed: true

  - name: sourceID
    description: ID of the source PU.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true

  - name: txFourTuple
    description: Transmiter four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true
