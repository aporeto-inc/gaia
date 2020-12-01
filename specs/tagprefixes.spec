# Model
model:
  rest_name: tagprefixes
  resource_name: tagprefixes
  entity_name: TagPrefixes
  package: squall
  group: core/namespace
  description: Returns the tag prefixes of the specified namespace.

# Attributes
attributes:
  v1:
  - name: prefixes
    description: List of tag prefixes that will be used to suggest policies.
    type: list
    exposed: true
    subtype: string
    stored: true
