# Model
model:
  rest_name: enforcertracereport
  resource_name: enforcertracereports
  entity_name: EnforcerTraceReport
  package: zack
  description: Post a new enforcer trace that determines how packets are.

# Attributes
attributes:
  v1:
  - name: enforcerID
    description: EnforcerID of the enforcer where the trace was collected.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: namespace
    description: Namespace of the PU where the trace was collected.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: puID
    description: ID of the pu where the trace was collected.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx
    filterable: true

  - name: traceList
    description: TraceList is the list of iptables trace records collected.
    type: refList
    subtype: tracerecord
    stored: true
    extensions:
      refMode: pointer
