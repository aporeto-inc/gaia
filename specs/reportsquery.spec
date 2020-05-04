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
  - name: descending
    description: If set, the results will be ordered by time from the most recent
      to the oldest.
    type: boolean
    exposed: true

  - name: filter
    description: Apply a filter to the query.
    type: string
    exposed: true

  - name: limit
    description: Limits the number of results. `-1` means no limit.
    type: integer
    exposed: true
    default_value: -1

  - name: offset
    description: Offsets the results. -1 means no offset.
    type: integer
    exposed: true
    default_value: -1

  - name: report
    description: Name of the report type to query.
    type: enum
    exposed: true
    allowed_choices:
    - Flows
    - Audit
    - Enforcers
    - Files
    - EventLogs
    - Packets
    - EnforcerTraces
    - Counters
    - Accesses
    - DNSLookups
    - PingReports
    default_value: Flows
