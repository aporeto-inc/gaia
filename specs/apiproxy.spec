# Model
model:
  rest_name: apiproxy
  resource_name: apiproxies
  entity_name: APIProxy
  package: relm
  group: integration/apiproxy
  description: |-
    Represents information needed to register and interact with an application's
    remote endpoint.
  aliases:
  - apiprox
  - apiproxs
  get:
    description: Retrieves the API proxy with the given ID.
  update:
    description: Updates the API proxy with the given ID.
  delete:
    description: Deletes the API proxy with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@timeable'
  validations:
  - $apiproxyEntity

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: |-
      Contains the PEM block of the certificate authority used by the
      remote endpoint.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBbjCCARSgAwIBAgIRANRbvVzTzBZOvMCb8BiKCLowCgYIKoZIzj0EAwIwJjEN
      MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNTE4
      NDgwN1oXDTI3MTEyNDE4NDgwN1owJjENMAsGA1UEChMEQWNtZTEVMBMGA1UEAxMM
      QWNtZSBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJ/80HR51+vau
      7XH7zS7b8ABA0e/TdBOg1NznbnXdXil1tDvWloWuH5+/bbaiEg54wksJHFXaukw8
      jhTLU7zT56MjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
      KoZIzj0EAwIDSAAwRQIhALwAZh2KLFFC1qfb5CqFHExlXS0PUltax9PvQCN9P0vl
      AiBl7/st9u/JpERjJgirxJxOgKNlV6pq9ti75EfQtZZcQA==
      -----END CERTIFICATE-----
    orderable: true
    validations:
    - $pem

  - name: clientCertificate
    description: |-
      Contains the client certificate that will be used to connect
      to the remote endpoint. If provided, the private key associated with this
      certificate must also be configured.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBczCCARigAwIBAgIRALD3Vz81Pq10g7n4eAkOsCYwCgYIKoZIzj0EAwIwJjEN
      MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNzA2
      NTM1MloXDTI3MTEyNjA2NTM1MlowGDEWMBQGA1UEAxMNY2xhaXJlLWNsaWVudDBZ
      MBMGByqGSM49AgEGCCqGSM49AwEHA0IABOmzPJj+t25T148eQH5gVrZ7nHwckF5O
      evJQ3CjSEMesjZ/u7cW8IBfXlxZKHxl91IEbbB3svci4c8pycUNZ2kujNTAzMA4G
      A1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
      MAoGCCqGSM49BAMCA0kAMEYCIQCjAAmkQpTua0HR4q6jnePaFBp/JMXwTXTxzbV6
      peGbBQIhAP+1OR8GFnn2PlacwHqWXHwkvy6CLPVikvgtwEdB6jH8
      -----END CERTIFICATE-----
    orderable: true
    validations:
    - $pem

  - name: clientCertificateKey
    description: |-
      Contains the key associated with the `clientCertificate`. It must be
      provided only when `clientCertificate` has been configured.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN EC PRIVATE KEY-----
      MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
      AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
      F9eXFkofGX3UgRtsHe123456789xQ1naSw==
      -----END EC PRIVATE KEY-----
    orderable: true
    secret: true
    transient: true
    encrypted: true
    validations:
    - $pem

  - name: endpoint
    description: Contains the full address of the remote api endpoint.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://api.remoteserver.com/remoteroute
    orderable: true
    validations:
    - $httpsURL

  - name: operation
    description: Defines the operation that is currently handled by the service.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - GET
    - PATCH
    - POST
    - PUT
    - DELETE
    default_value: GET

# Relations
relations:
- rest_name: call
  get:
    description: |-
      Allows a system to send a remote request to the API proxy based on the
      operation attribute.
  create:
    description: |-
      Allows a system to send a remote request to the API proxy based on the
      operation attribute.
