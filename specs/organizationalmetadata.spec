# Model
model:
  rest_name: organizationalmetadata
  resource_name: organizationalmetadata
  entity_name: OrganizationalMetadata
  package: squall
  group: core/namespace
  description: Can be used to retrieve the organizational metadata of the namespace.
  aliases:
  - om
  extends:
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: metadata
    description: List of organizational metadata for the namespace.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $metadata
