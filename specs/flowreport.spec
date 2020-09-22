# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Indexes
indexes:
- - namespace
  - timestamp
- - remotenamespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the flow.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept
    omit_empty: true
    extensions:
      bson_name: a

  - name: destinationController
    description: Identifier of the destination controller.
    type: string
    exposed: true
    stored: true
    example_value: api.east.acme.com
    omit_empty: true
    extensions:
      bson_name: b

  - name: destinationID
    description: ID of the destination.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: c

  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: d

  - name: destinationNamespace
    description: |-
      Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: e

  - name: destinationPlatform
    description: Identifier of the destination platform.
    type: string
    exposed: true
    stored: true
    example_value: api.east.acme.com
    omit_empty: true
    extensions:
      bson_name: f

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: g

  - name: destinationType
    description: Destination type.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit
    omit_empty: true
    extensions:
      bson_name: h

  - name: dropReason
    description: |-
      This field is only set if `action` is set to `Reject`. It specifies the reason
      for the rejection.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: i

  - name: encrypted
    description: If `true`, the flow was encrypted.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: j

  - name: namespace
    description: This is here for backward compatibility.
    type: string
    exposed: true
    stored: true
    required: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: k

  - name: observed
    description: If `true`, design mode is on.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: l

  - name: observedAction
    description: Action observed on the flow.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Accept
    - Reject
    - NotApplicable
    default_value: NotApplicable
    omit_empty: true
    extensions:
      bson_name: m

  - name: observedDropReason
    description: |-
      Specifies the reason for a rejection. Only set if `observedAction` is set
      to `Reject`.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: "n"

  - name: observedEncrypted
    description: Value of the encryption of the network policy that observed the flow.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: o

  - name: observedPolicyID
    description: ID of the network policy that observed the flow.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: p

  - name: observedPolicyNamespace
    description: Namespace of the network policy that observed the flow.
    type: string
    exposed: true
    stored: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: q

  - name: policyID
    description: ID of the network policy that accepted the flow.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: r

  - name: policyNamespace
    description: Namespace of the network policy that accepted the flow.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: s

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 6
    omit_empty: true
    extensions:
      bson_name: t

  - name: remoteNamespace
    description: Namespace of the object at the other end of the flow.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: u

  - name: serviceClaimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: v

  - name: serviceID
    description: ID of the service.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: w

  - name: serviceNamespace
    description: Namespace of Service accessed.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: x

  - name: serviceType
    description: ID of the service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - L3
    - HTTP
    - TCP
    - NotApplicable
    default_value: NotApplicable
    omit_empty: true
    extensions:
      bson_name: "y"

  - name: serviceURL
    description: Service URL accessed.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: z

  - name: sourceController
    description: Identifier of the source controller.
    type: string
    exposed: true
    stored: true
    example_value: api.west.acme.com
    omit_empty: true
    extensions:
      bson_name: aa

  - name: sourceID
    description: ID of the source.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: ab

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ac

  - name: sourceNamespace
    description: |-
      Namespace of the source. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: ad

  - name: sourcePlatform
    description: Identifier of the source platform.
    type: string
    exposed: true
    stored: true
    example_value: api.west.acme.com
    omit_empty: true
    extensions:
      bson_name: ae

  - name: sourceType
    description: Type of the source.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit
    omit_empty: true
    extensions:
      bson_name: af

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: ag

  - name: value
    description: Number of flows in the log.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
    omit_empty: true
    extensions:
      bson_name: ah
