# Model
model:
  rest_name: invoice
  resource_name: invoices
  entity_name: Invoice
  package: bill
  group: core/billing
  description: Provides access to Segment customer invoices.
  get:
    description: Retrieves the invoice with the given ID.
  update:
    description: Updates the invoice with the given ID.
  delete:
    description: Deletes the invoice with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: ID
    description: The ID of the invoice.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: accountID
    description: The ID of the customer that this invoice belongs to.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: billedToProvider
    description: The name of the provider that this invoice was billed to.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - AWS
    default_value: Aporeto

  - name: endDate
    description: The end date of the invoice.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: startDate
    description: The start date of this invoice.
    type: time
    exposed: true
    stored: true
    orderable: true
