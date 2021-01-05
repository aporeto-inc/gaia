# Model
model:
  rest_name: endpointdata
  resource_name: endpointdata
  entity_name: EndpointData
  package: pcn
  group: prisma/infrastructure
  description: Parameters associated with a cloud endpoint.
  detached: true

# Attributes
attributes:
  v1:
  - name: VPCAttached
    description: |-
      Indicates that the endpoint is directly attached to the VPC. In this case the
      attachedInterfaces is empty. In general this is only valid for endpoint type
      Gateway and Peering Connection.
    type: boolean
    exposed: true
    stored: true

  - name: VPCAttachments
    description: The list of VPCs that this endpoint is directly attached to.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: associatedRouteTables
    description: |-
      List of route tables associated with this endpoint. Depending on cloud provider
      it can apply in some gateways.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: attachedInterfaces
    description: |-
      A list of interfaces attached with the endpoint. In some cases endpoints can
      have more than one interface.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - eni-12344
    - eni-33333

  - name: forwardingEnabled
    description: |-
      If the endpoint has multiple connections and forwarding can be enabled between
      them.
    type: boolean
    exposed: true
    stored: true
    default_value: true
    example_value: false

  - name: type
    description: Type of the endpoint.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Instance
    - LoadBalancer
    - PeeringConnection
    - Service
    - Gateway
    - TransitGateway
    example_value: Instance
