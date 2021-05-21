# Model
model:
  rest_name: cloudnetworkinterface
  resource_name: cloudnetworkinterfaces
  entity_name: CloudNetworkInterface
  package: yeul
  group: pcn/infrastructure
  description: Manages the set of network interfaces that are associated with endpoints.
  get:
    description: Retrieves the network interface with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the network interface with the given ID.
  delete:
    description: Deletes the network interface with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: parameters
    description: Cloud network interface related parameters.
    type: ref
    exposed: true
    subtype: cloudinterfacedata
    stored: true
    filterable: true
    extensions:
      refMode: pointer
