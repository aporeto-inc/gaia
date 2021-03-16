# Model
model:
  rest_name: tagprefix
  resource_name: tagprefixes
  entity_name: TagPrefix
  package: squall
  group: core/namespace
  description: Returns the tag prefixes of the specified namespace.

# Attributes
attributes:
  v1:
  - name: prefixes
    description: |-
      List of tag prefixes that will be used to suggest policies. Only these tags will
      be transmitted on the wire.
    type: list
    exposed: true
    subtype: string
    stored: true
