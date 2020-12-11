# Model
model:
  rest_name: reportsquery
  resource_name: reportsqueries
  entity_name: ReportsQuery
  package: jenova
  group: visualization/reportsquery
  description: |-
    Supports querying Aporeto reports. All queries are protected within the
    namespace of the user.
  aliases:
  - rq

# Attributes
attributes:
  v1:
  - name: report
    description: Name of the report type to query.
    type: enum
    exposed: true
    allowed_choices:
    - Flows
    - Enforcers
    - EventLogs
    - Packets
    - Counters
    - DNSLookups
    default_value: Flows
