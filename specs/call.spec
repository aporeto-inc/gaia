# Model
model:
  rest_name: call
  resource_name: calls
  entity_name: Call
  package: relm
  group: integration/apiproxy
  description: Can be used to send a remote request to an API proxy.

# Attributes
attributes:
  v1:
  - name: payload
    description: Contains the remote `POST` payload.
    type: string
    exposed: true
