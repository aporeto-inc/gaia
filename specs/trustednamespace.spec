# Model
model:
  rest_name: trustednamespace
  resource_name: trustedamespaces
  entity_name: TrustedNamespace
  package: squall
  group: policy/enforcerconfig
  description: |-
    This objects allows to declare trust between namespaces that cryptographically
    isolated. The namespaces can be local or served by different Aporeto platforms.
  aliases: []
  get:
    description: Retrieve the trusted namespace with the given ID.
  update:
    description: Update the trusted namespace with the given ID.
  delete:
    description: Delete the trusted namespace with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: certificateAuthority
    description: Contains the PEM block of the certificate authority trusted namespace.
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
    validations:
    - $capem
