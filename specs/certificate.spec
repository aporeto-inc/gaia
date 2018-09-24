# Model
model:
  rest_name: certificate
  resource_name: certificates
  entity_name: Certificate
  package: vince
  description: A User represents the owner of some certificates.
  indexes:
  - - commonName
  - - parentID
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: admin
    description: Admin determines if the certificate must be added to the admin list.
    type: boolean
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: commonName
    description: CommonName (CN) for the user certificate.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: john doe
    filterable: true
    max_length: 64
    orderable: true

  - name: data
    description: Certificate provides a certificate for the user.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: email
    description: e-mail address of the user.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: john@doe.com
    filterable: true
    orderable: true

  - name: expirationDate
    description: CertificateExpirationDate indicates the expiration day for the certificate.
    type: time
    exposed: true
    stored: true
    creation_only: true
    orderable: true

  - name: key
    description: |-
      CertificateKey provides the key for the user. Only available at create or update
      time.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: name
    description: Name of the certificate.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: john.doe
    filterable: true
    orderable: true

  - name: organizationalUnits
    description: OrganizationalUnits attribute for the generated certificates.
    type: list
    exposed: true
    subtype: string
    stored: true
    creation_only: true

  - name: parentID
    description: ParentID holds the parent account ID.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    filterable: true
    orderable: true

  - name: passphrase
    description: Passphrase to use for the generated p12.
    type: string
    exposed: true
    creation_only: true
    orderable: true

  - name: serialNumber
    description: SerialNumber of the certificate.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    orderable: true

  - name: status
    description: |-
      CertificateStatus provides the status of the certificate. Update with RENEW to
      get a new certificate.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Revoked
    - Valid
    default_value: Valid
    filterable: true
    orderable: true
