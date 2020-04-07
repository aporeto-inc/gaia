# Model
model:
  rest_name: organizationalmetadata
  resource_name: organizationalmetadata
  entity_name: OrganizationalMetadata
  package: squall
  group: core/namespace
  description: |-
    Can be used to retrieve the organization metadata of the namespace and the
    organization metadata of its direct parent namespaces.

# Attributes
attributes:
  v1:
  - name: metadata
    description: |-
      List of organization metadata for the namespace and its direct parent
      namespaces.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $metadata
