# Model
model:
  rest_name: customer
  resource_name: customers
  entity_name: Customer
  package: bill
  description: |-
    This api allows to view and manage basic information about customer profile for
    billing purposes.
  private: true
  get: true
  update: true
  delete: true
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: provider
    description: Provider holds the name of the provider to be billed for this service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - AWS
    default_value: Aporeto

  - name: providerCustomerID
    description: |-
      ProviderCustomerID holds the customer id as used by the provider for this
      customer to enable provider billing.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: state
    description: State holds the status of the customer with the provider.
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
