# Model
model:
  rest_name: timeseriesrow
  resource_name: timeseriesrows
  entity_name: TimeSeriesRow
  package: jenova
  description: Represent a time series row.
  detached: true

# Attributes
attributes:
  v1:
  - name: columns
    description: colums of the row.
    type: list
    exposed: true
    subtype: string

  - name: name
    description: the name of row.
    type: string
    exposed: true

  - name: tags
    description: List of tags.
    type: external
    exposed: true
    subtype: map_of_string_of_strings

  - name: values
    description: List of tags.
    type: external
    exposed: true
    subtype: list_of_lists_of_objects
