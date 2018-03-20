# Model
model:
  rest_name: x509certificatecheck
  resource_name: x509certificatechecks
  entity_name: X509CertificateCheck
  package: barret
  description: '[nodoc]'
  private: true
  get: true

# Attributes
attributes:
- name: ID
  description: ID contains the certificate serialNumber
  type: string
  exposed: true
  required: true
  format: free
  identifier: true
