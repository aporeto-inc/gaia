# Model
model:
  rest_name: packetreport
  resource_name: packetreports
  entity_name: PacketReport
  package: zack
  description: Post a new packet tracing report.

# Attributes
attributes:
  v1:
  - name: DstPort
    description: DstPort is the destination port of the packet.
    type: integer
    exposed: true
    example_value: 11000
    max_value: 65536

  - name: ID
    description: ID is the packet ID from the IP header.
    type: integer
    required: true
    example_value: 12333

  - name: SrcPort
    description: SrcPort is the source port of the packet.
    type: integer
    exposed: true
    example_value: 80
    max_value: 65536

  - name: claims
    description: Claims is the list of claims detected for the packet.
    type: external
    subtype: enforcer_claims

  - name: destinationIP
    description: Type of the destination.
    type: string
    exposed: true

  - name: dropReason
    description: |-
      This field is only set if 'action' is set to 'Reject' and specifies the reason
      for the rejection.
    type: string
    exposed: true

  - name: encrypt
    description: Encrypt indicates that the packet was encrypted.
    type: boolean
    exposed: true
    required: true
    example_value: false

  - name: event
    description: Event is the event that triggered the report.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Rcv
    - Txt
    - Drop
    example_value: Rcv

  - name: flags
    description: Flags are the TCP flags of the packet.
    type: integer

  - name: length
    description: Length is the length of the packet.
    type: integer
    example_value: 94
    max_value: 65536

  - name: mark
    description: Mark is the mark value of the packet.
    type: integer
    required: true
    example_value: 123123

  - name: namespace
    description: Namespace of the PU reporting the packet.
    type: string
    exposed: true
    example_value: /my/namespace
    filterable: true

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    required: true
    example_value: 6
    max_value: 255

  - name: puID
    description: ID of the PU reporting the packet.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: triremePacket
    description: TriremePacket is set if the packet arrived with the Trireme options.
    type: boolean
    exposed: true
    stored: true
    required: true
    default_value: true
