# Model
model:
  rest_name: cloudnetworkquerydestination
  resource_name: cloudnetworkquerydestinations
  entity_name: CloudNetworkQueryDestination
  package: yeul
  group: prisma/infrastructure
  description: Returns the set of discovered destinations and the associated verdicts.
  detached: true

# Attributes
attributes:
  v1:
  - name: reachable
    description: Returns true if the destination is reachable through routing.
    type: boolean
    exposed: true
    autogenerated: true

  - name: type
    description: Returns the type of the destination.
    type: enum
    exposed: true
    allowed_choices:
    - Interface
    - Instance
    - LoadBalancer
    - PublicIP
    autogenerated: true

  - name: verdict
    description: Returns the network security verdict for the destination.
    type: string
    exposed: true
    autogenerated: true
