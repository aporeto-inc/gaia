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
    description: ID unique to a single origin and reply report.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: destinationID
    description: ID of the destination PU.
    type: string
    exposed: true

  - name: destinationNamespace
    description: Namespace of the destination PU.
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

  - name: flowTuple
    description: Flow tuple in the format <sip:dip:spt:dpt>.
    type: string
    exposed: true

  - name: latency
    description: Time taken for a single request to complete.
    type: string
    exposed: true

  - name: payloadSize
    description: Size of the payload attached to the packet.
    type: integer
    exposed: true

  - name: pingType
    description: Represents the ping type used.
    type: string
    exposed: true

  - name: protocol
    description: Protocol used for the communication.
    type: integer
    exposed: true

  - name: request
    description: Request represents the request number.
    type: integer
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

  - name: sourceNamespace
    description: Namespace of the source PU.
    type: string
    exposed: true
    required: true
    example_value: /my/ns

  - name: stage
    description: Stage when the packet is received.
    type: string
    exposed: true

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
