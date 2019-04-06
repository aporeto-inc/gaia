# Model
model:
  rest_name: graphgroup
  resource_name: graphgroups
  entity_name: GraphGroup
  package: jenova
  group: visualization/depmaps
  description: Represents an group of nodes from the dependency map.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of the group.
    type: string
    exposed: true

  - name: color
    description: Color to use for the group.
    type: string
    exposed: true

  - name: match
    description: List of tag that was used to create this group.
    type: external
    exposed: true
    subtype: '[][]string'
    validations:
    - $tagsExpression

  - name: name
    description: Name of the group.
    type: string
    exposed: true

  - name: parentID
    description: ID of the parent group if any.
    type: string
    exposed: true
