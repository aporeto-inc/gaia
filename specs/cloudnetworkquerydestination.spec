# Model
model:
  rest_name: cloudnetworkquerydestination
  resource_name: cloudnetworkquerydestinations
  entity_name: CloudNetworkQueryDestination
  package: yeul
  group: pcn/infrastructure
  description: Returns the set of discovered destinations and the associated verdicts.
  detached: true

# Attributes
attributes:
  v1:
  - name: indirect
    description: Returns true if this is an indirect path through an midlebox.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true

  - name: indirectNode
    description: Returns the native ID of the indirect node.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: reachable
    description: Returns true if the destination is reachable through routing.
    type: boolean
    exposed: true
    read_only: true
    autogenerated: true

  - name: type
    description: Returns the type of the destination.
    type: enum
    exposed: true
    read_only: true
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
    read_only: true
    autogenerated: true
