# Model
model:
  rest_name: cloudinterfacedata
  resource_name: cloudinterfacedata
  entity_name: CloudInterfaceData
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

  - name: attachmentType
    description: |-
      Attachment type describes where this interface is attached to (Instance, Load
      Balancer, Gateway, etc).
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Instance
    - LoadBalancer
    - Gateway
    - Service
    - TransitGatewayVPCAttachment
    - NetworkLoadBalancer
    - Lambda
    - GatewayLoadBalancer
    - GatewayLoadBalancerEndpoint
    - VPCEndpoint
    - APIGatewayManaged
    - EFA

  - name: relatedObjectID
    description: |-
      If the interface is of type or external, the relatedObjectID identifies the
      related service or gateway.
    type: string
    exposed: true
    stored: true

  - name: routeTableID
    description: |-
      The route table that must be used for this interface. Applies to TransitGateways
      and other special types.
    type: string
    exposed: true
    stored: true
    example_value:
    - rt1233

  - name: securityTags
    description: Security tags associated with the instance.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: subnets
    description: ID of subnet associated with this interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - subnet-074c152ae45ea0c73
