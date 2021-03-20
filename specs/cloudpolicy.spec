# Model
model:
  rest_name: cloudpolicy
  resource_name: cloudpolicies
  entity_name: CloudPolicy
  package: yeul
  group: pcn/infrastructure
  description: Creates a Prisma Cloud policy and corresponding alert rules.
  get:
    description: Retrieves the Prisma Cloud policy with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the Prisma Cloud policy with the given ID.
  delete:
    description: Deletes the the Prisma Cloud policy with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: prismaCloudPolicyID
    description: Reference to the corresponding Prisma Cloud Policy ID.
    type: string
    exposed: true
    stored: true

  - name: queryID
    description: |-
      The query ID that this policy refers to. This is auto-calculated since it is
      derived from the parent.
    type: string
    exposed: true
    read_only: true

  - name: severity
    description: The severity of a policy violation.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Low
    - Medium
    - High
