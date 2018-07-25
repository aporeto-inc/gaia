# Model
model:
  rest_name: externalnetwork
  resource_name: externalnetworks
  entity_name: ExternalNetwork
  package: squall
  description: |-
    An External Network represents a random network or ip that is not managed by the
    system. They can be used in Network Access Policies in order to allow traffic
    from or to the declared network or IP, using the provided protocol and port or
    ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
    as address, and 1-65000 for the ports. You will need to use the External
    Services tags to set some policies.
  aliases:
  - extnet
  - extnets
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
  - name: entries
    description: List of CIDRs or domain name.
    type: list
    exposed: true
    subtype: string
    stored: true
    format: free

  - name: port
    description: |-
      Port refers to network port which could be a single number or 100:2000 to
      represent a range of ports.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$
    default_value: 1:65535
    filterable: true

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
