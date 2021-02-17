# Model
model:
  rest_name: pcsearchresult
  resource_name: pcsearchresults
  entity_name: PCSearchResults
  package: karl
  group: core/rql
  description: Represents the result data of RQL search.

# Attributes
attributes:
  v1:
  - name: items
    description: The payload of the search result.
    type: ref
    exposed: true
    subtype: reportsquery
    read_only: true
    extensions:
      refMode: pointer

  - name: nextPageToken
    description: The pagination token for next page.
    type: string
    exposed: true
    read_only: true

  - name: totalRows
    description: The total number of result items.
    type: integer
    exposed: true
    read_only: true
