# Model
model:
  rest_name: invoice
  resource_name: invoices
  entity_name: Invoice
  package: bill
  description: This api allows to view invoices for Aporeto customers.
  get: true
  update: true
  delete: true
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: AccountID
    description: AccountID references the id of the customer that this invoice belongs
      to.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: BilledToProvider
    description: BilledToProvider holds the name of the provider that this invoice
      was billed to.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - AWS
    default_value: Aporeto

  - name: EndDate
    description: EndDate holds the end date for this invoice.
    type: time
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: ID
    description: ID is the id of the invoice.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: StartDate
    description: StartDate holds the start date for this invoice.
    type: time
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true
