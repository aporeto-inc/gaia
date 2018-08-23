# Model
model:
  rest_name: invoice
  resource_name: invoices
  entity_name: Invoice
  package: bill
  description: This api allows to view invoices for Aporeto customers.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the id of the invoice.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: accountID
    description: AccountID references the id of the customer that this invoice belongs
      to.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: billedToProvider
    description: BilledToProvider holds the name of the provider that this invoice
      was billed to.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - AWS
    default_value: Aporeto

  - name: endDate
    description: EndDate holds the end date for this invoice.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: startDate
    description: StartDate holds the start date for this invoice.
    type: time
    exposed: true
    stored: true
    orderable: true
