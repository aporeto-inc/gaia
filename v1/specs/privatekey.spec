attributes:
- description: ID is the internal ID of the key.
  format: free
  identifier: true
  name: ID
  primary_key: true
  stored: true
  type: string
- description: CertificateSerialNumber represents the certificate serial number associated
    to this key.
  format: free
  name: certificateSerialNumber
  stored: true
  type: string
- creation_only: true
  description: Data contains the privateKey data.
  format: free
  name: data
  stored: true
  type: string
model:
  delete: true
  get: true
  update: true
  description: Internal representation of an private key
  entity_name: PrivateKey
  package: barret
  resource_name: privatekeys
  rest_name: privatekey
  private: true
