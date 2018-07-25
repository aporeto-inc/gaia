# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  description: Post a new flow statistics report.

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the flow.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: destinationID
    description: ID of the destination.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: destinationIP
    description: Type of the destination.
    type: string
    exposed: true

  - name: destinationNamespace
    description: Namespace of the receiver.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true

  - name: destinationType
    description: Type of the source.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalService
    - Claims
    example_value: ProcessingUnit

  - name: dropReason
    description: Reason for the rejection.
    type: string
    exposed: true

  - name: encrypted
    description: Tells is the flow has been encrypted.
    type: boolean
    exposed: true

  - name: namespace
    description: This is here for backward compatibility.
    type: string
    exposed: true
    required: true
    deprecated: true
    example_value: /my/namespace

  - name: observed
    description: Tells if the flow is from design mode.
    type: boolean
    exposed: true

  - name: observedAction
    description: Action observed on the flow.
    type: enum
    exposed: true
    allowed_choices:
    - Accept
    - Reject
    - NotApplicable
    default_value: NotApplicable

  - name: observedPolicyID
    description: ID of the network policy that observed the flow.
    type: string
    example_value: xxx-xxx-xxx

  - name: observedPolicyNamespace
    description: Namespace of the network policy that observed the flow.
    type: string
    example_value: /my/namespace

  - name: policyID
    description: ID of the network policy that accepted the flow.
    type: string
    required: true
    example_value: xxx-xxx-xxx

  - name: policyNamespace
    description: Namespace of the network policy that accepted the flow.
    type: string
    required: true
    example_value: /my/namespace

  - name: protocol
    description: protocol number.
    type: integer
    exposed: true
    required: true
    example_value: 6

  - name: serviceClaimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true

  - name: serviceID
    description: ID of the service.
    type: string
    exposed: true

  - name: serviceNamespace
    description: Service URL accessed.
    type: string
    exposed: true

  - name: serviceType
    description: ID of the service.
    type: enum
    exposed: true
    allowed_choices:
    - L3
    - HTTP
    - TCP

  - name: serviceURL
    description: Service URL accessed.
    type: string
    exposed: true

  - name: sourceID
    description: ID of the source.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true

  - name: sourceNamespace
    description: Namespace of the receiver.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: sourceType
    description: Type of the source.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalService
    - Claims
    example_value: ProcessingUnit

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: value
    description: Number of flows in the report.
    type: integer
    exposed: true
    required: true
    example_value: 1
