# Model
model:
  rest_name: reportsqueryresults
  resource_name: reportsqueryresults
  entity_name: ReportsQueryResults
  package: jenova
  group: visualization/reportsquery
  description: Represent the results of a reports query.
  detached: true

# Attributes
attributes:
  v1:
  - name: fields
    description: List of projected fields.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: groups
    description: List of projected fields.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: values
    description: List of values associated with the projected fields.
    type: external
    exposed: true
    subtype: '[][]interface{}'
    stored: true
