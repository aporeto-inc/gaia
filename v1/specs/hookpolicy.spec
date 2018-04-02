# Model
model:
  rest_name: hookpolicy
  resource_name: hookpolicies
  entity_name: HookPolicy
  package: squall
  description: |-
    Hook allows to to define hooks to the write operations in squall. Hooks are sent
    to an external Rufus server that will do the processing and eventually return a
    modified version of the object before we save it.
  aliases:
  - hook
  - hooks
  - hookpol
  - hookpols
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: |-
      CertificateAuthority contains the pem block of the certificate authority used by
      the remote endpoint.
    type: string
    exposed: true
    stored: true
    required: true
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
    format: free
    orderable: true

  - name: clientCertificate
    description: |-
      ClientCertificate contains the client certificate that will be used to connect
      to the remote endoint.
    type: string
    exposed: true
    stored: true
    required: true
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
    format: free
    orderable: true

  - name: clientCertificateKey
    description: ClientCertificateKey contains the key associated to the clientCertificate.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: |-
      -----BEGIN EC PRIVATE KEY-----
      MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
      AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
      F9eXFkofGX3UgRtsHe123456789xQ1naSw==
      -----END EC PRIVATE KEY-----
    format: free
    orderable: true
    secret: true

  - name: endpoint
    description: Endpoint contains the full address of the remote processor endoint.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: https://hooks.hookserver.com/remoteprocessors
    filterable: true
    format: free
    orderable: true

  - name: mode
    description: Mode define the type of the hook.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Both
    - Post
    - Pre
    default_value: Pre
    filterable: true
    orderable: true

  - name: subject
    description: |-
      Subject contains the tag expression that an object must match in order to
      trigger the hook.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
    example_value:
    - - $identity=processingunit
