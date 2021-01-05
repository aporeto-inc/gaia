# Model
model:
  rest_name: vpcdata
  resource_name: vpcdata
  entity_name: VPCData
  package: pcn
  group: prisma/infrastructure
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
    example_value: 10.0.0.0/8
    validations:
    - $cidr
