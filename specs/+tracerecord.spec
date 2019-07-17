# Model
model:
  rest_name: tracerecord
  resource_name: tracerecords
  entity_name: TraceRecord
  package: zack
  group: core/enforcer
  description: Represents a single trace record from the enforcer.
  detached: true

# Attributes
attributes:
  v1:
  - name: TTL
    description: The time to live (TTL) value of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 64
    max_value: 255

  - name: chain
    description: Chain that the trace was collected from.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: PREROUTING

  - name: destinationIP
    description: The destination IP.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: destinationInterface
    description: The destination interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: destinationPort
    description: The destination UPD or TCP port of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536
    min_value: 1

  - name: length
    description: Length of the observed packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 98
    max_value: 65536

  - name: packetID
    description: The IP packet header ID.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 10

  - name: protocol
    description: The protocol of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536

  - name: ruleID
    description: Priority index of the iptables entry that was hit.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 10

  - name: sourceIP
    description: Source IP of the packet.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: sourceInterface
    description: Source interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: sourcePort
    description: Source TCP or UDP port of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536
    min_value: 1

  - name: tableName
    description: The iptables name that the trace collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: raw

  - name: timestamp
    description: The time-date stamp of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
