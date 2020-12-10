# Model
model:
  rest_name: cnssuggest
  resource_name: cnssuggests
  entity_name: CNSSuggest
  package: karl
  group: core/cnssuggest
  description: Provide search suggestion query for Primsa Cloud's investigate page.

# Attributes
attributes:
  v1:
  - name: needsOffsetUpdate
    description: Rquired by Prisma Cloud. Always set to true.
    type: boolean
    exposed: true
    default_value: true

  - name: offset
    description: The length of the rql query part that is valid.
    type: integer
    exposed: true
    default_value: 0

  - name: query
    description: Prisma Cloud's rql query.
    type: string
    exposed: true
    read_only: true
    example_value: network from DNS where id == 1

  - name: suggestions
    description: List of query suggestions.
    type: list
    exposed: true
    subtype: string
    example_value:
    - id
    - action

  - name: translate
    description: Rquired by Prisma Cloud. Always set to false.
    type: boolean
    exposed: true
    default_value: false

  - name: valid
    description: The validity of the rql query.
    type: boolean
    exposed: true
    default_value: false
