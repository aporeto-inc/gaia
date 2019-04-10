# Model
model:
  rest_name: pkixname
  resource_name: pkixnames
  entity_name: PKIXName
  package: barret
  group: internal/x509
  description: Represents a PKIX.Name.
  detached: true

# Attributes
attributes:
  v1:
  - name: commonName
    description: Represents the CommonName field.
    type: string
    exposed: true

  - name: country
    description: Represents the Country field.
    type: list
    exposed: true
    subtype: string

  - name: locality
    description: Represents the Locality field.
    type: list
    exposed: true
    subtype: string

  - name: organization
    description: Represents the Organization field.
    type: list
    exposed: true
    subtype: string

  - name: organizationalUnit
    description: Represents the OrganizationalUnit field.
    type: list
    exposed: true
    subtype: string

  - name: postalCode
    description: Represents the PostalCode field.
    type: list
    exposed: true
    subtype: string

  - name: province
    description: Represents the Province field.
    type: list
    exposed: true
    subtype: string

  - name: streetAddress
    description: Represents the StreetAddress field.
    type: list
    exposed: true
    subtype: string
