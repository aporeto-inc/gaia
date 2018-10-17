# Model
model:
  rest_name: credential
  resource_name: credentials
  entity_name: Credential
  package: cactuar
  description: Represents an application credential data.
  detached: true

# Attributes
attributes:
  v1:
  - name: ID
    description: The ID of app credential.
    type: string
    exposed: true

  - name: certificate
    description: The certificate data encoded in base64.
    type: string
    exposed: true

  - name: certificateAuthority
    description: The certificate authority data encoded in base64.
    type: string
    exposed: true

  - name: certificateKey
    description: The certificate key data encoded in base64.
    type: string
    exposed: true

  - name: name
    description: The name of app credential.
    type: string
    exposed: true

  - name: namespace
    description: The namespace of app credential.
    type: string
    exposed: true
