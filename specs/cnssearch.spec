# Model
model:
  rest_name: cnssearch
  resource_name: cnssearches
  entity_name: CNSSearch
  package: karl
  group: core/cnsearch
  description: Provide search results for Primsa Cloud's investigate page.

# Attributes
attributes:
  v1:
  - name: data
    description: The payload of the search results.
    type: external
    exposed: true
    subtype: _pc_rql_table_data

  - name: description
    description: Description of the search.
    type: string
    exposed: true

  - name: id
    description: ID of the search request.
    type: string
    exposed: true

  - name: limit
    description: The number of items to fetch.
    type: integer
    exposed: true

  - name: pageToken
    description: Represents the token to fetch next page.
    type: string
    exposed: true

  - name: query
    description: The rql query.
    type: string
    exposed: true

  - name: saved
    description: Indicates if the search has been saved.
    type: boolean
    exposed: true

  - name: searchType
    description: Type of search request. Always set to be network.
    type: string
    exposed: true

  - name: timeRange
    description: Time range of the search.
    type: external
    exposed: true
    subtype: _pc_rql_timerange
