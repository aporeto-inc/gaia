# Model
model:
  rest_name: pctimerange
  resource_name: pctimeranges
  entity_name: PCTimeRange
  package: karl
  group: core/rql
  description: Represents the time range parameter of PC.

# Attributes
attributes:
  v1:
  - name: relativeTimeType
    description: The type of relative time.
    type: string
    exposed: true
    read_only: true

  - name: type
    description: The type of time range.
    type: string
    exposed: true
    read_only: true

  - name: value
    description: The value of time range.
    type: external
    exposed: true
    subtype: pctimevalue
    read_only: true
