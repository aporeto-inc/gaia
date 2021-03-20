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
    description: The customer name of the tenant.
    type: string
    exposed: true
    required: true
    example_value: customer-name

  - name: name
    description: The name of the account.
    type: string
    exposed: true
    example_value: account-foo
