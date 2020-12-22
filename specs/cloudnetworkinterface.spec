# Model
model:
  rest_name: cloudnetworkinterface
  resource_name: cloudnetworkinterfaces
  entity_name: CloudNetworkInterface
  package: pcn
  group: prisma/infrastructure
  description: Manages the set of network interfaces that are associated with endpoints.
  get:
    description: Retrieves the network interface with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the network interface with the given ID.
  delete:
    description: Deletes the network interface with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@prismabase'

# Indexes
indexes:
- - nativeID
- - namespace
  - nativeID

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

  - name: subnets
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
