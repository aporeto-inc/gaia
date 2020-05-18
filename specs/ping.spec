# Model
model:
  rest_name: ping
  resource_name: ping
  entity_name: Ping
  package: guy
  group: core/enforcer
  description: Post a new pu ping.
  extends:
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@timeable'
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: RTT
    description: Time taken for a single request-response to complete.
    type: string
    exposed: true
    stored: true

  - name: TXController
    description: Controller of the transmitter.
    type: string
    exposed: true
    stored: true

  - name: TXType
    description: Type of the transmitter.
    type: string
    exposed: true
    stored: true

  - name: applicationListening
    description: If true, application responded to the request.
    type: boolean
    exposed: true
    stored: true

  - name: claims
    description: Claims carried on the wire.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: destinationID
    description: ID of the destination processing unit.
    type: string
    exposed: true
    stored: true

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns

  - name: enforcerVersion
    description: Semantic version of the enforcer.
    type: string
    exposed: true
    stored: true

  - name: error
    description: A non-empty error indicates a failure.
    type: string
    exposed: true
    stored: true

  - name: excludedNetworks
    description: If true, destination IP is in excludedNetworks.
    type: boolean
    exposed: true
    stored: true

  - name: fourTuple
    description: Four tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true
    stored: true

  - name: iterationID
    description: IterationID unique to a single ping request-response.
    type: string
    exposed: true
    stored: true

  - name: payloadSize
    description: Size of the payload attached to the packet.
    type: integer
    exposed: true
    stored: true

  - name: pingID
    description: PingID unique to a single ping control.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: policyAction
    description: Action of the policy.
    type: string
    exposed: true
    stored: true

  - name: policyID
    description: ID of the policy.
    type: string
    exposed: true
    stored: true

  - name: processingUnitNamespace
    description: Namespace of the reporting processing unit.
    type: string
    exposed: true
    stored: true

  - name: protocol
    description: Protocol used for the communication.
    type: integer
    exposed: true
    stored: true

  - name: seqNum
    description: Sequence number of the TCP packet. number.
    type: integer
    exposed: true
    stored: true

  - name: serviceType
    description: Type of the service.
    type: string
    exposed: true
    stored: true

  - name: sourceID
    description: ID of the source PU.
    type: string
    exposed: true
    stored: true

  - name: targetTCPNetworks
    description: If true, destination IP is in targetTCPNetworks.
    type: boolean
    exposed: true
    stored: true

  - name: type
    description: Type of the report.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Request
    - Response
