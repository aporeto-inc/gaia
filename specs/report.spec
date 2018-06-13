# Model
model:
  rest_name: report
  resource_name: reports
  entity_name: Report
  package: zack
  description: Post a new statistics report.

# Attributes
attributes:
  v1:
  - name: fields
    description: TSDB Fields to set for the report.
    type: external
    exposed: true
    subtype: tsdb_fields

  - name: kind
    description: Kind contains the kind of report.
    type: enum
    exposed: true
    allowed_choices:
    - Audit
    - Enforcer
    - FileAccess
    - Flow
    - ProcessingUnit
    - Syscall
    - Claims

  - name: tags
    description: Tags contains the tags associated to the data point.
    type: external
    exposed: true
    subtype: tags_map

  - name: timestamp
    description: Timestamp contains the time for the report.
    type: time
    exposed: true

  - name: value
    description: Value contains the value for the report.
    type: float
    exposed: true
