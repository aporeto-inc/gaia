# Model
model:
  rest_name: x509certificatecheck
  resource_name: x509certificatechecks
  entity_name: X509CertificateCheck
  package: barret
  description: Verifies a certificate has not been revoked.
  private: true
  get: true

# Attributes
attributes:
  v1:
  - name: ID
    description: ID contains the certificate serialNumber.
    type: string
    exposed: true
    required: true
    example_value: c155b59c-c04b-430f-b11f-8355a6b7dc48
    format: free
    identifier: true
