# Model
model:
  rest_name: certificates
  resource_name: certificates
  entity_name: Certificates
  package: squall
  group: core/namespace
  description: |-
    Can be used to retrieve or renew the local certificate authority of the
    namespace.

# Attributes
attributes:
  v1:
  - name: LocalCA
    description: The certificate authority used by the namespace.
    type: string
    exposed: true
    transient: true

  - name: LocalCARenew
    description: Set to `true` to renew the local certificate authority of the namespace.
    type: boolean
    exposed: true
    transient: true
