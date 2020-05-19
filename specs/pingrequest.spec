# Model
model:
  rest_name: pingrequest
  resource_name: pingrequest
  entity_name: PingRequest
  package: guy
  group: core/enforcer
  description: Post a new pu pingrequest.

# Attributes
attributes:
  v1:
  - name: iterations
    description: Number of probes that will be triggered.
    type: integer
    exposed: true
    stored: true
