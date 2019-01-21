# Model
model:
  rest_name: tracemode
  resource_name: tracemode
  entity_name: TraceMode
  package: squall
  description: TraceMode is the tracing mode that must be applied to a PU.
  detached: true

# Attributes
attributes:
  v1:
  - name: ApplicationConnections
    description: |-
      ApplicationConnections instructs the enforcer to send records for all
      application initiated connections.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: IPTables
    description: IPTables instructs the enforcers to provide an iptables trace for
      a PU.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: NetworkConnections
    description: |-
      NetworkConnections instructs the enforcer to send records for all network
      initiated connections.
    type: boolean
    exposed: true
    stored: true
    default_value: false

  - name: TimeInterval
    description: |-
      TimeInterval determins the length of the time interval that the trace must be
      enabled.
    type: string
    exposed: true
    stored: true
    default_value: 10s
