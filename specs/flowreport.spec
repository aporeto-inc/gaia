# Model
model:
  rest_name: flowreport
  resource_name: flowreports
  entity_name: FlowReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@flowedge'

# Attributes
attributes:
  v1:
  - name: destinationController
    description: Identifier of the destination controller.
    type: string
    exposed: true
    example_value: api.east.acme.com

  - name: destinationID
    description: ID of the destination.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true

  - name: destinationNamespace
    description: |-
      Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true

  - name: destinationType
    description: Destination type.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit

  - name: sourceController
    description: Identifier of the source controller.
    type: string
    exposed: true
    example_value: api.west.acme.com

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
    description: |-
      Namespace of the source. This is deprecated. Use `remoteNamespace`. This
      property does nothing.
    type: string
    exposed: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: sourceType
    description: Type of the source.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - ProcessingUnit
    - ExternalNetwork
    - Claims
    example_value: ProcessingUnit
