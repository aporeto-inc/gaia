# Model
model:
  rest_name: tracerecord
  resource_name: tracerecord
  entity_name: TraceRecord
  package: zack
  description: Represents a single trace record from the enforcer.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the packet ID.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 10
    filterable: true

  - name: Length
    description: Length of the observed packet.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 98
    max_value: 65536

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

  - name: destIP
    description: DestIP is the destination IP.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: destInterface
    description: DestInterface is the destination interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: dstPort
    description: DstPort is the destination Port.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536

  - name: policy
    description: Policy is the index of the iptables entry that was hit.
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

  - name: srcIP
    description: SrcIP is the source IP.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.1.1.30

  - name: srcInterface
    description: SrcInterface is the source interface of the packet.
    type: string
    exposed: true
    stored: true
    example_value: en0

  - name: srcPort
    description: SrcPort is the source Port.
    type: integer
    exposed: true
    stored: true
    required: true
    example_value: 80
    max_value: 65536

  - name: table
    description: Table is the iptable that the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: raw

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
