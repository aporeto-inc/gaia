# Model
model:
  rest_name: containerimage
  resource_name: containerimages
  entity_name: ContainerImage
  package: aki
  group: integration/vulnerability
  description: A container image can be affected by vulnerabilities and compliance
    issues.
  get:
    description: Retrieves a container image with a given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the container image with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@timeable'

# Indexes
indexes:
- - dockerImages
- - externalID

# Attributes
attributes:
  v1:
  - name: complianceRiskScore
    description: Score of the compliance.
    type: integer
    exposed: true
    stored: true
    example_value: 16

  - name: criticalComplianceIssueCount
    description: Number of critical compliance issue.
    type: integer
    exposed: true
    stored: true
    example_value: 0

  - name: criticalVulnerabilityCount
    description: Number of critical vulnerabilities.
    type: integer
    exposed: true
    stored: true
    example_value: 119

  - name: distro
    description: Full name of the distribution.
    type: string
    exposed: true
    stored: true
    example_value: Debian GNU/Linux 9 (stretch)
    filterable: true

  - name: externalID
    description: External Identifier of the container image scan.
    type: string
    exposed: true
    stored: true
    example_value: sha256:4635a5562b040fd83ec821bb885405587a52cfef898ffb7402649005dfda75ff
    filterable: true

  - name: highComplianceIssueCount
    description: Number of high compliance issue.
    type: integer
    exposed: true
    stored: true
    example_value: 0

  - name: highVulnerabilityCount
    description: Number of high vulnerabilities.
    type: integer
    exposed: true
    stored: true
    example_value: 16

  - name: lowComplianceIssueCount
    description: Number of low compliance issue.
    type: integer
    exposed: true
    stored: true
    example_value: 0

  - name: lowVulnerabilityCount
    description: Number of low vulnerabilities.
    type: integer
    exposed: true
    stored: true
    example_value: 7

  - name: mediumComplianceIssueCount
    description: Number of medium compliance issue.
    type: integer
    exposed: true
    stored: true
    example_value: 0

  - name: mediumVulnerabilityCount
    description: Number of medium vulnerabilities.
    type: integer
    exposed: true
    stored: true
    example_value: 20

  - name: osDistro
    description: Name of the os distribution.
    type: string
    exposed: true
    stored: true
    example_value: debian
    filterable: true

  - name: osDistroRelease
    description: Name of the release.
    type: string
    exposed: true
    stored: true
    example_value: stretch
    filterable: true

  - name: totalComplianceIssueCount
    description: Number of total compliance issue.
    type: integer
    exposed: true
    stored: true
    example_value: 0

  - name: totalVulnerabilityCount
    description: Number of total vulnerabilities.
    type: integer
    exposed: true
    stored: true
    example_value: 145

  - name: vulnerabilitiesCount
    description: Number of vulnerabilities affecting this image.
    type: integer
    exposed: true
    stored: true
    example_value: 119

  - name: vulnerabilityRiskScore
    description: Score of the vulnerability.
    type: integer
    exposed: true
    stored: true
    example_value: 16

# Relations
relations:
- rest_name: complianceissue
  get:
    description: Returns the list of compliance issues of a container image.

- rest_name: vulnerability
  get:
    description: Returns the list of vulnerabilities of a container image.
