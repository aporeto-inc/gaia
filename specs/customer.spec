# Model
model:
  rest_name: customer
  resource_name: customers
  entity_name: Customer
  package: bill
  group: core/billing
  description: |-
    This API allows you to view and manage basic information about customer profiles for
    billing purposes.
  private: true
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-stored'
  - '@timeable'

# Indexes
indexes:
- - providerCustomerID

# Attributes
attributes:
  v1:
  - name: committedUseLimit
    description: |-
      Holds the customer's use limit provided by a contract.
    type: integer
    exposed: true
    stored: true
    orderable: true

  - name: lastReportTime
    description: |-
      Holds the last time the customer was reported to their
      provider.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: provider
    description: Holds the name of the provider to be billed for this service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - AWS
    default_value: Aporeto
    filterable: true

  - name: providerCustomerID
    description: |-
      Holds the customer ID as used by the provider for this
      customer to enable provider billing.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: providerProductID
    description: |-
      Holds the product ID as used by the provider for this
      customer to enable provider billing.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: state
    description: Holds the status of the customer with the provider.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - SubscribePending
    - SubscribeFailed
    - SubscribeSuccess
    - UnsubscribePending
    - UnsubscribeSuccess
    default_value: SubscribePending
