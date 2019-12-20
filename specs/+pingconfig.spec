# Model
model:
  rest_name: pingconfig
  resource_name: pingconfig
  entity_name: PingConfig
  package: squall
  group: core/enforcer
  description: Represents the ping configuration to apply to a processing unit.
  detached: true

# Attributes
attributes:
  v1:
  - name: network
    description: Destination network to test.
    type: string
    exposed: true
    stored: true

  - name: ports
    description: Destination port(s) to test.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $ports

  - name: requests
    description: Number of requests to make on one call.
    type: integer
    exposed: true
    stored: true

  - name: type
    description: Ping type.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - None
    - AporetoIdentity
    - CustomIdentity
    - AporetoIdentityPassthrough
    default_value: None
