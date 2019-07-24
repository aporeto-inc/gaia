# Model
model:
  rest_name: hostservice
  resource_name: hostservices
  entity_name: HostService
  package: squall
  group: policy/hosts
  description: Represents services that a host must expose and protect.
  aliases:
  - hostsrv
  - hostsrvs
  get:
    description: Retrieves the host service with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the host service with the given ID.
  delete:
    description: Deletes the host service with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@archivable'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@timeable'
  validations:
  - $hostservices

# Attributes
attributes:
  v1:
  - name: hostModeEnabled
    description: |-
      Forces the corresponding enforcers to enable host protection. When `true`, all
      incoming and outgoing flows will be monitored. Flows will be allowed if and only
      if a network policy has been created to allow the flow. The option applies to
      all
      enforcers to which the host service is mapped.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: services
    description: |-
      Lists all protocols and ports a service is running. A service entry can be
      defined
      by a protocol and port `(tcp/80)`, or range of protocol/port pairs
      `(udp/80:100)`.
      If no protocol is provided, it is assumed to be TCP. Only `tcp` and `udp`
      protocols
      are allowed.
    type: list
    exposed: true
    subtype: string
    stored: true
