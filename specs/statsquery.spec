# Model
model:
  rest_name: statsquery
  resource_name: statsqueries
  entity_name: StatsQuery
  package: jenova
  description: |-
    StatsQuery is a generic API to retrieve time series data stored by the Aporeto
    system. The API allows different types of queries that are all protected within
    the namespace of the user.
  aliases:
  - sq

# Attributes
attributes:
  v1:
  - name: fields
    description: |-
      List of fields to extract. If you don't pass anything, all available fields will
      be returned. It is also possible to use function like `sum(value)`.
    type: list
    exposed: true
    subtype: string

  - name: groups
    description: |-
      Group results by the provided values. Note that not all fields can be used to
      group the results.
    type: list
    exposed: true
    subtype: string

  - name: limit
    description: Limits the number of results. -1 means no limit.
    type: integer
    exposed: true
    default_value: -1

  - name: measurement
    description: Name of the measurement.
    type: enum
    exposed: true
    allowed_choices:
    - Flows
    - Audit
    - Enforcers
    - Files
    - Eventlogs
    default_value: Flows

  - name: offset
    description: Offsets the of results. -1 means no offset.
    type: integer
    exposed: true
    default_value: -1

  - name: results
    description: Results contains the result of the query.
    type: external
    exposed: true
    subtype: time_series_results
    read_only: true
    autogenerated: true
