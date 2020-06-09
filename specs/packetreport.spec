# Model
model:
  rest_name: packetreport
  resource_name: packetreports
  entity_name: PacketReport
  package: zack
  group: core/enforcer
  description: Post a new packet tracing report.

# Attributes
attributes:
  v1:
  - name: TCPFlags
    description: Flags are the TCP flags of the packet.
    type: integer
    exposed: true
    stored: true

  - name: claims
    description: Claims is the list of claims detected for the packet.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: destinationIP
    description: The destination IP address of the packet.
    type: string
    exposed: true
    stored: true

  - name: destinationPort
    description: The destination port of a TCP or UDP packet.
    type: integer
    exposed: true
    stored: true
    example_value: 11000
    max_value: 65536

  - name: dropReason
    description: |-
      If `event` is set to `Dropped`, contains the reason that the packet was dropped.
      Otherwise empty.
    type: string
    exposed: true
    stored: true

  - name: encrypt
    description: Set to `true` if the packet was encrypted.
    type: boolean
    exposed: true
    stored: true

  - name: enforcerID
    description: Identifier of the defender sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxx-xxx-xxxx

  - name: enforcerNamespace
    description: Namespace of the defender sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: event
    description: The event that triggered the report.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Received
    - Transmitted
    - Dropped
    example_value: Rcv

  - name: length
    description: Length is the length of the packet.
    type: integer
    stored: true
    example_value: 94
    max_value: 65536

  - name: mark
    description: Mark is the mark value of the packet.
    type: integer
    exposed: true
    stored: true
    example_value: 123123

  - name: namespace
    description: Namespace of the processing unit reporting the packet.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    filterable: true

  - name: packetID
    description: The ID of the IP header of the reported packet.
    type: integer
    exposed: true
    stored: true
    example_value: 12333

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    example_value: 6
    max_value: 255

  - name: puID
    description: The ID of the processing unit reporting the packet.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: rawPacket
    description: The first 64 bytes of the packet.
    type: string
    exposed: true
    stored: true
    default_value: abcd

  - name: sourceIP
    description: The source IP address of the packet.
    type: string
    exposed: true
    stored: true

  - name: sourcePort
    description: The source port of the packet.
    type: integer
    exposed: true
    stored: true
    example_value: 80
    max_value: 65536

  - name: timestamp
    description: The time-date stamp of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: triremePacket
    description: Set to `true` if the packet arrived with the Trireme options (default).
    type: boolean
    exposed: true
    stored: true
    default_value: true
