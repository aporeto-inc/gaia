# Model
model:
  rest_name: cloudsnapshotaccount
  resource_name: cloudsnapshotaccounts
  entity_name: CloudSnapshotAccount
  package: pandemona
  group: pcn/infrastructure
  description: |-
    Initiates a poll for a particular account. Data are stored in the current
    namespace.
  extends:
  - '@base'
  - '@namespaced'
  - '@zoned'
  - '@migratable'
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: cloudType
    description: The cloud type for the account.
    type: enum
    exposed: true
    allowed_choices:
    - AWS
    - GCP
    default_value: AWS

  - name: customerName
    description: The Customer name of the tenant where the account is onboarded.
    type: string
    exposed: true
    required: false
    example_value: customer-name

  - name: name
    description: The name of the account as onboarded in Prisma Cloud.
    type: string
    exposed: true
    required: true
    example_value: account-foo

  - name: tenantID
    description: The Prisma ID of the tenant where the account is onboarded.
    type: string
    exposed: true
    required: true
    example_value: tenant-foo
