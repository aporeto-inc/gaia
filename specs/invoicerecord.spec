# Model
model:
  rest_name: invoicerecord
  resource_name: invoicerecords
  entity_name: InvoiceRecord
  package: bill
  description: This api allows to view detailed records of invoices for Aporeto customers.
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
    description: ID is the id of this invoice record.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: invoiceID
    description: |-
      InvoiceID references the id of the invoice that this invoice record provides
      details for.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: invoiceRecords
    description: InvoiceRecords provides details about billing units.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true
