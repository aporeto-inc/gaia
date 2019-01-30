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
  - name: IPTables
    description: IPTables instructs the enforcers to provide an iptables trace for
      a PU.
    type: boolean
    exposed: true
    stored: true

  - name: applicationConnections
    description: |-
      Instructs the enforcer to send records for all
      application initiated connections.
    type: boolean
    exposed: true
    stored: true

  - name: interval
    description: |-
      Determines the length of the time interval that the trace must be
      enabled.
    type: string
    exposed: true
    stored: true
    default_value: 10s

  - name: networkConnections
    description: |-
      Instructs the enforcer to send records for all network
      initiated connections.
    type: boolean
    exposed: true
    stored: true
