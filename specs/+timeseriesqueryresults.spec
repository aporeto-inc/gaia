# Model
model:
  rest_name: timeseriesqueryresults
  resource_name: timeseriesqueryresults
  entity_name: TimeSeriesQueryResults
  package: jenova
  description: Represent the results of a stats query.
  detached: true

# Attributes
attributes:
  v1:
  - name: rows
    description: List of rows.
    type: refList
    exposed: true
    subtype: timeseriesrow
    extensions:
      refMode: pointer
