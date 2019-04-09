# Model
model:
  rest_name: authority
  resource_name: authorities
  entity_name: Authority
  package: barret
  group: internal/x509
  description: Authority represents a certificate authority.
  aliases:
  - ca
  private: true
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering

# Indexes
indexes:
- - :shard
  - $hashed:serialNumber
- - :unique
  - serialNumber
- - commonName

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the identitfier of the Authority.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    identifier: true

  - name: certificate
    description: PEM encoded certificate data.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: commonName
    description: CommonName contains the common name of the CA.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: my ca

  - name: expirationDate
    description: Date of expiration of the authority.
    type: time
    exposed: true
    stored: true
    creation_only: true

  - name: key
    description: Encrypted private key of the Authority.
    type: string
    stored: true
    autogenerated: true

  - name: serialNumber
    description: serialNumber of the certificate.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
