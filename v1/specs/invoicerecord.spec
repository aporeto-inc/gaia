# Model
model:
  rest_name: invoicerecord
  resource_name: invoicerecords
  entity_name: InvoiceRecord
  package: bill
  description: This api allows to view detailed records of invoices for Aporeto customers.
  get: true
  update: true
  delete: true
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
    filterable: true
    format: free
    orderable: true

  - name: InvoiceID
    description: |-
      InvoiceID references the id of the invoice that this invoice record provides
      details for.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: InvoiceRecords
    description: InvoiceRecords provides details about billing units.
    type: external
    exposed: true
    subtype: invoicerecord_list
    stored: true
    filterable: true
    format: free
    orderable: true
