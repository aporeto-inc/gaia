# Model
model:
  rest_name: externalservice
  resource_name: externalservices
  entity_name: ExternalService
  package: squall
  group: deprecated
  description: This API is deprecated in favor of externalnetworks.
  aliases:
  - extsrv
  - extsrvs
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
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: network
    description: Network refers to either CIDR or domain name.
    type: string
    exposed: true
    stored: true
    example_value: 0.0.0.0/0

  - name: port
    description: |-
      Port refers to network port which could be a single number or 100:2000 to
      represent a range of ports.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$
    allowed_chars_message: must be a valid single number between 1 and 65535 or port
      range made or two number separated by a colon
    default_value: 1:65535

  - name: protocol
    description: Protocol refers to network protocol like TCP/UDP or the number of
      the protocol.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$
    allowed_chars_message: must be a valid IANA protocol number or TCP or UDP
    example_value: TCP
