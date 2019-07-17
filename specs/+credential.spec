# Model
model:
  rest_name: credential
  resource_name: credentials
  entity_name: Credential
  package: cactuar
  group: policy/authorization
  description: Represents an app credential.
  detached: true

# Attributes
attributes:
  v1:
  - name: APIURL
    description: The URL of the Aporeto API.
    type: string
    exposed: true

  - name: ID
    description: The ID of the app credential.
    type: string
    exposed: true

  - name: certificate
    description: The base64-encoded certificate.
    type: string
    exposed: true

  - name: certificateAuthority
    description: The base64-encoded certificate authority.
    type: string
    exposed: true

  - name: certificateKey
    description: The base64-encoded certificate key.
    type: string
    exposed: true

  - name: name
    description: The name of the app credential.
    type: string
    exposed: true

  - name: namespace
    description: The namespace of the app credential.
    type: string
    exposed: true
