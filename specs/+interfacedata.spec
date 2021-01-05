# Model
model:
  rest_name: interfacedata
  resource_name: interfacedata
  entity_name: InterfaceData
  package: pcn
  group: prisma/infrastructure
  description: Parameters associated with a cloud interface.
  detached: true

# Attributes
attributes:
  v1:
  - name: addresses
    description: |-
      List of IP addresses/subnets (IPv4 or IPv6) associated with the
      interface.
    type: refList
    exposed: true
    subtype: cloudaddress
    stored: true

  - name: relatedObjectID
    description: |-
      If the interface is of type or external, the relatedObjectID identifies the
      related service or gateway.
    type: string
    exposed: true
    stored: true

  - name: securityTags
    description: Security tags associated with the instance.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: subnet
    description: ID of subnet associated with this interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value: subnet-074c152ae45ea0c73

  - name: type
    description: Interface type (Instance, Load Balancer, Gateway, etc).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Instance
    - LoadBalancer
    - Gateway
    - Service
    - TransitGateway
