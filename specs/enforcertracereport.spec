# Model
model:
  rest_name: enforcertracereport
  resource_name: enforcertracereports
  entity_name: EnforcerTraceReport
  package: zack
  group: core/enforcer
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
    example_value: 5c6cce207ddf1fc159a104bf

  - name: enforcerNamespace
    description: Namespace of the enforcer where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /acme/prod

  - name: namespace
    description: Namespace of the PU where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /acme/prod/database

  - name: puID
    description: ID of the pu where the trace was collected.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 5c6ccd947ddf1fc159a104b7

  - name: records
    description: List of iptables trace records collected.
    type: refList
    subtype: tracerecord
    stored: true
    extensions:
      refMode: pointer
