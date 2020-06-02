# Model
model:
  rest_name: certificates
  resource_name: certificates
  entity_name: Certificates
  package: squall
  group: core/namespace
  description: |-
    Can be used to retrieve or renew the local and SSH certificate authorities of
    the namespace.

# Attributes
attributes:
  v1:
  - name: SSHCA
    description: The SSH certificate authority used by the namespace.
    type: string
    exposed: true
    transient: true

  - name: SSHCARenew
    description: Set to `true` to renew the SSH certificate authority of the namespace.
    type: boolean
    exposed: true
    transient: true

  - name: localCA
    description: The certificate authority used by the namespace.
    type: string
    exposed: true
    transient: true

  - name: localCARenew
    description: Set to `true` to renew the local certificate authority of the namespace.
    type: boolean
    exposed: true
    transient: true
