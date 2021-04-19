# Model
model:
  rest_name: cloudendpointdataproductinfo
  resource_name: cloudendpointdataproductinfo
  entity_name: CloudEndpointDataProductInfo
  package: yeul
  group: pcn/infrastructure
  description: Parameters associated with a cloud endpoint data product.
  detached: true

# Attributes
attributes:
  v1:
  - name: productID
    description: The ID of the corresponding product.
    type: string
    exposed: true
    stored: true
    omit_empty: true

  - name: type
    description: The type of the product.
    type: string
    exposed: true
    stored: true
    omit_empty: true
