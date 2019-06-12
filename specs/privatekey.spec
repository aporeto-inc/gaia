# Model
model:
  rest_name: privatekey
  resource_name: privatekeys
  entity_name: PrivateKey
  package: barret
  group: internal/x509
  description: Internal representation of an private key.
  private: true
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the internal ID of the key.
    type: string
    stored: true
    identifier: true

  - name: certificateSerialNumber
    description: |-
      CertificateSerialNumber represents the certificate serial number associated to
      this key.
    type: string
    stored: true

  - name: data
    description: Data contains the privateKey data.
    type: string
    stored: true
    creation_only: true
