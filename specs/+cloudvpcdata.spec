# Model
model:
  rest_name: cloudvpcdata
  resource_name: cloudvpcdata
  entity_name: CloudVPCData
  package: yeul
  group: pcn/infrastructure
  description: Managed the list of IP addresses associated with an interface.
  detached: true

# Attributes
attributes:
  v1:
  - name: address
    description: Address CIDR of the VPC.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 10.0.0.0/8
    validations:
    - $cidr
