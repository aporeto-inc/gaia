# Model
model:
  rest_name: appcredential
  resource_name: appcredentials
  entity_name: AppCredential
  package: cactuar
  group: policy/authorization
  description: Create an app credential.
  aliases:
  - appcred
  - appcreds
  get:
    description: Retrieves the app credential with the given ID.
  update:
    description: Updates the app credential with the given ID.
  delete:
    description: Deletes the app credential with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@timeable'
  - '@named'
  - '@disabled'

# Attributes
attributes:
  v1:
  - name: CSR
    description: |-
      Contains a PEM-encoded certificate signing request (CSR). It can
      only be set during a renew.

      - The CN **MUST** be `app:credential:<appcred-id>:<appcred-name>`
      - The O **MUST** be the namespace of the app credential

      If you send anything else, the signing request will be rejected.
    type: string
    exposed: true
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

  - name: authorizedSubnets
    description: |-
      If set, the app credential will only be valid if the request comes from one
      the declared subnets.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: certificate
    description: The string representation of the certificate used by the app credential.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: certificateSN
    description: Link to the certificate created for this app credential.
    type: string
    stored: true

  - name: credentials
    description: The app credential data.
    type: ref
    exposed: true
    subtype: credential
    read_only: true
    autogenerated: true
    orderable: true
    transient: true
    extensions:
      noInit: true
      refMode: pointer

  - name: email
    description: The email address that will receive a copy of the app credential.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: maxIssuedTokenValidity
    description: |-
      If set, this will limit the maximum validity of the token issued from this app
      credential. This information will be embeded into the delivered certificate and
      cannot be changed once set. In order to change it, you need to renew the
      certificate.
    type: string
    exposed: true
    stored: true
    validations:
    - $optionalTimeDuration

  - name: parentIDs
    description: |-
      Contains the ID of the parent app credential if this is a derived app
      credential.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    autogenerated: true

  - name: roles
    description: List of roles to give the app credential.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - '@auth:role=enforcer'
    - '@auth:role=kubesquall'
