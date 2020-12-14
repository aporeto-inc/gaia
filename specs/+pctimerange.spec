# Model
model:
  rest_name: pctimerange
  resource_name: pctimeranges
  entity_name: PCTimeRange
  package: karl
  group: core/rql
  description: Represents the time range option of rql search request.

# Attributes
attributes:
  v1:
  - name: type
    description: Type of the time range.
    type: string
    exposed: true

  - name: value
    description: Value of the time range.
    type: ref
    exposed: true
    subtype: pctimerangevalue
    extensions:
      refMode: pointer
