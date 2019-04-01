# Model
model:
  rest_name: revocation
  resource_name: revocations
  entity_name: Revocation
  package: barret
  group: internal/x509
  description: Used to revoke a certificate.
  private: true
  indexes:
  - - :shard
    - $hashed:serialNumber
  - - :unique
    - serialNumber
  update:
    description: Updates the object with the given ID.

# Attributes
attributes:
  v1:
  - name: ID
    description: ID contains the ID of the revocation.
    type: string
    stored: true
    read_only: true
    autogenerated: true
    identifier: true
    primary_key: true

  - name: expirationDate
    description: |-
      Contains the certificate expiration date. This will be used to clean up revoked
      certificates that have expired.
    type: time
    exposed: true
    stored: true
    creation_only: true

  - name: revokeDate
    description: Set time from when the certificate will be revoked.
    type: time
    exposed: true
    stored: true

  - name: serialNumber
    description: SerialNumber of the revoked certificate.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: subject
    description: Subject of the certificate related to the revocation.
    type: string
    exposed: true
    stored: true
    creation_only: true
    validations:
    - $policyExpression
