# Model
model:
  rest_name: externalservice
  resource_name: externalservices
  entity_name: ExternalService
  package: squall
  description: |-
    An External Service represents a random network or ip that is not managed by the
    system. They can be used in Network Access Policies in order to allow traffic
    from or to the declared network or IP, using the provided protocol and port or
    ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
    as address, and 1-65000 for the ports. You will need to use the External
    Services tags to set some policies.
  aliases:
  - extsrv
  - extsrvs
  get: true
  update: true
  delete: true
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: loadbalancerAddresses
    description: |-
      LoadbalancerAddresses represents the list of adresses of the external services
      of type LoadBalancer.
    type: external
    exposed: true
    subtype: addresses_list
    stored: true

  - name: loadbalancerPortsMapping
    description: |-
      LoadbalancerPortsMapping is the list of ports mapped by an extenral service of
      type load balancer.
    type: external
    exposed: true
    subtype: portmapping_list
    stored: true

  - name: network
    description: |-
      Network is a comma separated list of networks (CIDRs or IP addresses
      or subnets, where this external service is defined.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 0.0.0.0/0
    filterable: true
    format: free

  - name: port
    description: |-
      Port refers to network port which could be a comma separated list
      of single numbers or 100:2000 to represent a range of ports.
    type: string
    exposed: true
    stored: true
    default_value: 1:65535
    filterable: true
    format: free

  - name: protocol
    description: Protocol refers to network protocol like TCP/UDP or the number of
      the protocol.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
    example_value: TCP
    filterable: true

  - name: type
    description: Type represents the type of external service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - LoadBalancerHTTP
    - LoadBalancerTCP
    - Network
    default_value: Network
    filterable: true
    orderable: true
