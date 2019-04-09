# Model
model:
  rest_name: hostservice
  resource_name: hostservices
  entity_name: HostService
  package: squall
  group: policy/hosts
  description: Represents a set of services that a host must expose and protect.
  aliases:
  - hostsrv
  - hostsrvs
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'
  - '@propagated'
  - '@timeable'
  validations:
  - $hostservices

# Attributes
attributes:
  v1:
  - name: hostModeEnabled
    description: |-
      HostModeEnabled forces the corresponding enforcers to enable complete host
      protection. When this option is turned on, all incoming and outgoing flows will
      be monitored. Flows will be allowed if and only if a network policy has been
      created to allow the flow. The option applies to all enforcers that match the
      subject constraints.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: services
    description: |-
      Services lists all protocols and ports a service is running. A service entry can
      be defined by a protocol and port '(tcp/80)', or range of protocol/port pairs
      '(udp/80:100)'. If no protocol is provided, it is assumed to be TCP. Allowed
      protocols are only tcp and udp.
    type: list
    exposed: true
    subtype: string
    stored: true
