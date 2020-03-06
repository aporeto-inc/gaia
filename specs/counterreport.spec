# Model
model:
  rest_name: counterreport
  resource_name: counterreports
  entity_name: CounterReport
  package: zack
  group: core/enforcer
  description: Post a new counter tracing report.

# Attributes
attributes:
  v1:
  - name: connectionsDropped
    description: Counter for connections dropped.
    type: integer
    exposed: true
    default_value: 0

  - name: connectionsExpired
    description: Counter for connections expired.
    type: integer
    exposed: true
    default_value: 0

  - name: connectionsProcessed
    description: Counter for connections processed.
    type: integer
    exposed: true
    default_value: 0

  - name: droppedPackets
    description: Counter for dropped packets.
    type: integer
    exposed: true
    default_value: 0

  - name: encryptionFailures
    description: Counter for encryption failures.
    type: integer
    exposed: true
    default_value: 0

  - name: enforcerID
    description: Identifier of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxxx-xxx-xxxx

  - name: enforcerNamespace
    description: Namespace of the enforcer sending the report.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace

  - name: externalNetworkConnections
    description: Counter for external network connections.
    type: integer
    exposed: true
    default_value: 0

  - name: policyDrops
    description: Counter for policy drops.
    type: integer
    exposed: true
    default_value: 0

  - name: processingUnitID
    description: PUID is the ID of the PU reporting the counter.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: processingUnitNamespace
    description: Namespace of the PU reporting the counter.
    type: string
    exposed: true
    example_value: /my/namespace
    filterable: true

  - name: timestamp
    description: Timestamp is the date of the report.
    type: time
    exposed: true
    example_value: "2018-06-14T23:10:46.420397985Z"

  - name: tokenDrops
    description: Counter for token drops.
    type: integer
    exposed: true
    default_value: 0
