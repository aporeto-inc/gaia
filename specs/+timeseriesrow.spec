# Model
model:
  rest_name: timeseriesrow
  resource_name: timeseriesrows
  entity_name: TimeSeriesRow
  package: jenova
  group: visualization/statsquery
  description: Represents a time-series row.
  detached: true

# Attributes
attributes:
  v1:
  - name: columns
    description: Columns of the row.
    type: list
    exposed: true
    subtype: string

  - name: name
    description: Name of the row.
    type: string
    exposed: true

  - name: tags
    description: List of tags.
    type: external
    exposed: true
    subtype: map[string]string

  - name: values
    description: List of tags.
    type: external
    exposed: true
    subtype: '[][]interface{}'
