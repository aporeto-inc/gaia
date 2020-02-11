# Model
model:
  rest_name: complianceissue
  resource_name: complianceissues
  entity_name: ComplianceIssue
  package: aki
  group: aki/vulnerabilities
  description: Represents a compliance issue.
  get:
    description: Retrieves a compliance issue with a given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the compliance issue with the given ID.
  delete:
    description: Deletes the compliance issue with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@timeable'

# Indexes
indexes:
- - severity
  - namespace

# Attributes
attributes:
  v1:
  - name: severity
    description: Refers to the security vulnerability level.
    type: external
    exposed: true
    subtype: _vulnerability_level
    stored: true
    required: true
    creation_only: true
    example_value: 3
