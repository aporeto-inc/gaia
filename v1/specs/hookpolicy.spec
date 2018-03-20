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
- name: certificateAuthority
  description: |-
    CertificateAuthority contains the pem block of the certificate authority used by
    the remote endpoint.
  type: string
  exposed: true
  stored: true
  required: true
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
  format: free
  orderable: true

- name: clientCertificateKey
  description: ClientCertificateKey contains the key associated to the clientCertificate.
  type: string
  exposed: true
  stored: true
  required: true
  format: free
  orderable: true
  secret: true

- name: endpoint
  description: Endpoint contains the full address of the remote processor endoint.
  type: string
  exposed: true
  stored: true
  required: true
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
