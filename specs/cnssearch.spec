# Model
model:
  rest_name: cnssearch
  resource_name: cnssearches
  entity_name: CNSSearch
  package: karl
  group: core/rql
  description: Provide search results for Primsa Cloud's investigate page.

# Attributes
attributes:
  v1:
  - name: data
    description: The payload of the search results.
    type: ref
    exposed: true
    subtype: pcsearchresult
    extensions:
      refMode: pointer

  - name: description
    description: Description of the search.
    type: string
    exposed: true
    omit_empty: true

  - name: id
    description: ID of the search request.
    type: string
    exposed: true
    omit_empty: true

  - name: limit
    description: The number of items to fetch.
    type: integer
    exposed: true
    omit_empty: true

  - name: name
    description: Name of the rql search request. Should set to be empty.
    type: string
    exposed: true
    omit_empty: true

  - name: pageToken
    description: Represents the token to fetch next page.
    type: string
    exposed: true
    omit_empty: true

  - name: query
    description: The rql query.
    type: string
    exposed: true
    omit_empty: true

  - name: saved
    description: Indicates if the search has been saved.
    type: boolean
    exposed: true
    omit_empty: true

  - name: searchType
    description: Type of search request. Should set to be network.
    type: string
    exposed: true
    omit_empty: true

  - name: timeRange
    description: Time range of the search.
    type: ref
    exposed: true
    subtype: pctimerange
    omit_empty: true
    extensions:
      noInit: true
      refMode: pointer
