# Model
model:
  rest_name: comment
  resource_name: comments
  entity_name: Comment
  package: vivi
  group: core
  description: Represents a comment from a user.
  detached: true

# Attributes
attributes:
  v1:
  - name: claims
    description: The claims of the author.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: content
    description: The content of the comment.
    type: string
    exposed: true
    stored: true

  - name: date
    description: The date of the comment.
    type: time
    exposed: true
    stored: true
