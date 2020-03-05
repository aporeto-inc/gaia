# Model
model:
  rest_name: image
  resource_name: images
  entity_name: Image
  package: aki
  group: core/processingunit
  description: A container image can be affected by vulnerabilities.
  get:
    description: Retrieves a container image with a given ID.
  update:
    description: Updates the container image with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'
  - '@propagated'

# Indexes
indexes:
- - hash
- - namespace
  - hash
- - severity
- - namespace
  - severity

# Attributes
attributes:
  v1:
  - name: hash
    description: Hash of the image.
    type: string
    exposed: true
    stored: true
    example_value: sha256:4635a5562b040fd83ec821bb885405587a52cfef898ffb7402649005dfda75ff

  - name: severity
    description: Overall severity of the container image.
    type: external
    exposed: true
    subtype: _vulnerability_level
    stored: true

  - name: vulnerabilities
    description: List of vulnerabilities affecting this image.
    type: list
    exposed: true
    subtype: string
    stored: true
