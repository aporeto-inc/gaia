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
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: namespace
    description: Namespace of the PU where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: puID
    description: ID of the pu where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx

  - name: records
    description: List of iptables trace records collected.
    type: refList
    subtype: tracerecord
    stored: true
    extensions:
      refMode: pointer
