# Model
model:
  rest_name: packetreport
  resource_name: packetreports
  entity_name: PacketReport
  package: zack
  group: core/enforcer
  description: Post a new packet tracing report.
  extends:
  - '@identifiable-stored'
  - '@zoned-monotonic'
  - '@migratable'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

# Attributes
attributes:
  v1:
  - name: TCPFlags
    description: Flags are the TCP flags of the packet.
    type: integer
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: a

  - name: claims
    description: Claims is the list of claims detected for the packet.
    type: list
    exposed: true
    subtype: string
    stored: true
    omit_empty: true
    extensions:
      bson_name: b

  - name: destinationIP
    description: The destination IP address of the packet.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: c

  - name: destinationPort
    description: The destination port of a TCP or UDP packet.
    type: integer
    exposed: true
    stored: true
    example_value: 11000
    max_value: 65536
    omit_empty: true
    extensions:
      bson_name: d

  - name: dropReason
    description: |-
      If `event` is set to `Dropped`, contains the reason that the packet was dropped.
      Otherwise empty.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: e

  - name: encrypt
    description: Set to `true` if the packet was encrypted.
    type: boolean
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: f

  - name: enforcerID
    description: Identifier of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxx-xxx-xxxx
    omit_empty: true
    extensions:
      bson_name: g

  - name: enforcerNamespace
    description: Namespace of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    omit_empty: true
    extensions:
      bson_name: h

  - name: event
    description: The event that triggered the report.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Received
    - Transmitted
    - Dropped
    example_value: Rcv
    omit_empty: true
    extensions:
      bson_name: i

  - name: length
    description: Length is the length of the packet.
    type: integer
    stored: true
    example_value: 94
    max_value: 65536
    omit_empty: true
    extensions:
      bson_name: j

  - name: mark
    description: Mark is the mark value of the packet.
    type: integer
    exposed: true
    stored: true
    example_value: 123123
    omit_empty: true
    extensions:
      bson_name: k

  - name: namespace
    description: Namespace of the processing unit reporting the packet.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    filterable: true
    omit_empty: true
    extensions:
      bson_name: l

  - name: packetID
    description: The ID of the IP header of the reported packet.
    type: integer
    exposed: true
    stored: true
    example_value: 12333
    omit_empty: true
    extensions:
      bson_name: m

  - name: protocol
    description: Protocol number.
    type: integer
    exposed: true
    stored: true
    example_value: 6
    max_value: 255
    omit_empty: true
    extensions:
      bson_name: "n"

  - name: puID
    description: The ID of the processing unit reporting the packet.
    type: string
    exposed: true
    stored: true
    example_value: xxx-xxx-xxx
    filterable: true
    omit_empty: true
    extensions:
      bson_name: o

  - name: rawPacket
    description: The first 64 bytes of the packet.
    type: string
    exposed: true
    stored: true
    default_value: abcd
    omit_empty: true
    extensions:
      bson_name: p

  - name: sourceIP
    description: The source IP address of the packet.
    type: string
    exposed: true
    stored: true
    omit_empty: true
    extensions:
      bson_name: q

  - name: sourcePort
    description: The source port of the packet.
    type: integer
    exposed: true
    stored: true
    example_value: 80
    max_value: 65536
    omit_empty: true
    extensions:
      bson_name: r

  - name: timestamp
    description: The time-date stamp of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
    orderable: true
    omit_empty: true
    extensions:
      bson_name: s

  - name: triremePacket
    description: Set to `true` if the packet arrived with the Trireme options (default).
    type: boolean
    exposed: true
    stored: true
    default_value: true
    omit_empty: true
    extensions:
      bson_name: t
