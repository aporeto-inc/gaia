# Model
model:
  rest_name: tracerecord
  resource_name: tracerecords
  entity_name: TraceRecord
  package: zack
  description: Represents a single trace record from the enforcer.
  detached: true

# Attributes
attributes:
  v1:
  - name: TTL
    description: TTL is the TTL value of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 64
    max_value: 255

  - name: chain
    description: Chain is the chain that the trace was collected from.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: PREROUTING

  - name: destinationIP
    description: DestinationIP is the destination IP.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: destinationInterface
    description: DestinationInterface is the destination interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: destinationPort
    description: DestinationPort is the destination UPD or TCP port of the packet.
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
    description: PacketID is the IP packet header ID.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 10

  - name: protocol
    description: Protocol is the protocol of the packets.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536

  - name: ruleID
    description: ruleID is the priority index of the iptables entry that was hit.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 10

  - name: sourceIP
    description: SourceIP is the source IP of the packet.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: sourceInterface
    description: SourceInterface is the source interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: sourcePort
    description: SourcePort is the source TCP or UDP Port of the packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536
    min_value: 1

  - name: tableName
    description: TableName is the iptable name that the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: raw

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
