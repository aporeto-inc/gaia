# Model
model:
  rest_name: connectionexceptionreport
  resource_name: connectionexceptionreports
  entity_name: ConnectionExceptionReport
  package: zack
  group: policy/networking
  description: Post a new flow log.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Attributes
attributes:
  v1:
  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true
    stored: true
    omit_empty: true

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: processingUnitID
    description: ID of the processing unit encountered this exception.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true

  - name: processingUnitNamespace
    description: Namespace of the processing unit encountered this exception.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: /my/namespace
    omit_empty: true

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 6
    omit_empty: true

  - name: reason
    description: It specifies the reason for the exception.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: remoteController
    description: Identifier of the remote controller.
    type: string
    exposed: true
    stored: true
    deprecated: true
    example_value: api.west.acme.com
    omit_empty: true

  - name: remoteProcessingUnitID
    description: ID of the remote processing unit.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx
    omit_empty: true

  - name: sourceIP
    description: Source IP address.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: state
    description: Represents the current state this report was generated.
    type: enum
    exposed: true
    subtype: string
    stored: true
    required: true
    allowed_choices:
    - SynTransmitted
    - SynAckTransmitted
    - AckTransmitted
    - Unknown
    example_value:
    - Unknown

  - name: timestamp
    description: Time and date of the report.
    type: time
    exposed: true
    stored: true
    omit_empty: true

  - name: value
    description: Number of packets hit.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 1
    omit_empty: true
