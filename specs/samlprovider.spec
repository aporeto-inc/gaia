# Model
model:
  rest_name: samlprovider
  resource_name: samlproviders
  entity_name: SAMLProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows to declare a generic SAML provider that can be used in
    exchange for a Midgard token.
  get:
    description: Retrieves the provider with the given ID.
  update:
    description: Updates the provider with the given ID.
  delete:
    description: Deletes the provider with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'
  validations:
  - $samlprovider

# Attributes
attributes:
  v1:
  - name: IDPCertificate
    description: Identity Provider Certificate in PEM format.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE REQUEST-----
      MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
      BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
      aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
      9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
      wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
      1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
      WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
      wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
      BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
      KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
      hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
      Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
      ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
      29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
      97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
      -----END CERTIFICATE REQUEST-----

  - name: IDPIssuer
    description: Identity Provider Issuer (also called Entity ID).
    type: string
    exposed: true
    stored: true
    example_value: https://accounts.google.com/o/saml2/idp?idpid=AbDcef123

  - name: IDPMetadata
    description: |-
      Pass some XML data containing the IDP metadata that can be used for automatic
      configuration. If you pass this attribute, every other one will be overwritten
      with the data contained in the metadata file.
    type: string
    exposed: true
    omit_empty: true

  - name: IDPURL
    description: URL of the identity provider.
    type: string
    exposed: true
    stored: true
    example_value: https://accounts.google.com/o/saml2/idp?idpid=AbDcef123

  - name: default
    description: |-
      If set, this will be the default SAML provider. There can be only one default
      provider in your account. When logging in with SAML, if no provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: subjects
    description: List of claims that will provide the subject.
    type: list
    exposed: true
    subtype: string
    stored: true
    example_value:
    - email
    - profile
