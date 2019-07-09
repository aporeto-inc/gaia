# Model
model:
  rest_name: invoicerecord
  resource_name: invoicerecords
  entity_name: InvoiceRecord
  package: bill
  group: core/billing
  description: Provides detailed records of invoices for Aporeto customers.
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

# Attributes
attributes:
  v1:
  - name: ID
    description: The ID of the invoice record.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: invoiceID
    description: |-
      The ID of the invoice associated with the invoice record.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: invoiceRecords
    description: Details about billing units.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true
