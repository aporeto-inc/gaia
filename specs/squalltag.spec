# Model
model:
  rest_name: squalltag
  resource_name: squalltags
  entity_name: SquallTag
  package: squall
  description: |-
    Internal api that retrieve all tags from squall and their count for given
    `?identity=<identity>` parameter with their counts.
  private: true

# Attributes
attributes:
  v1:
  - name: count
    description: Number of time this tag is used.
    type: integer
    exposed: true
    stored: true

  - name: namespace
    description: namespace containing these tags.
    type: string
    exposed: true
    stored: true

  - name: value
    description: Value of the tag.
    type: string
    exposed: true
    stored: true
