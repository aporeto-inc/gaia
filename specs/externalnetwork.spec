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
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
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
    validations:
    - $networks

  - name: ports
    description: List of single ports or range (xx:yy).
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    default_value:
    - 1:65535
    validations:
    - $ports

  - name: protocols
    description: List of protocols (tcp, udp, or protocol number).
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    default_value:
    - tcp
    validations:
    - $protocols
