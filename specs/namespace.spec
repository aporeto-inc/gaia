# Model
model:
  rest_name: namespace
  resource_name: namespaces
  entity_name: Namespace
  package: squall
  group: core/namespace
  description: |-
    A Namespace represents the core organizational unit of the system. All objects
    always exists in a single namespace. A Namespace can also have child namespaces.
    They can be used to split the system into organizations, business units,
    applications, services or any combination you like.
  aliases:
  - ns
  indexes:
  - - :shard
    - $hashed:name
  - - :unique
    - name
  - - namespace
  - - namespace
    - name
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'

# Attributes
attributes:
  v1:
  - name: SSHCA
    description: SSHCA holds the eventual SSH authority used by this namespace.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: SSHCAEnabled
    description: |-
      If enabled, the a SSH CA will be generated for the namespace. This CA can be
      deployed in SSH server to validate SSH certificates issued by the platform.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: associatedLocalCAID
    description: AssociatedLocalCAID holds the remote ID of the certificate authority
      to use.
    type: string
    stored: true
    read_only: true

  - name: associatedSSHCAID
    description: associatedSSHCAID holds the remote ID of the SSH authority to use.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: localCA
    description: LocalCA holds the eventual certificate authority used by this namespace.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: localCAEnabled
    description: |-
      LocalCAEnabled defines if the namespace should use a local Certificate
      Authority. Switching it off and on again will regenerate a new CA.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: name
    description: Name is the name of the namespace.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_chars: ^[a-zA-Z0-9-_/]+$
    allowed_chars_message: must only contain alpha numerical characters, '-' or '_'
    default_order: true
    example_value: mynamespace
    filterable: true
    getter: true
    orderable: true
    primary_key: true

  - name: networkAccessPolicyTags
    description: |-
      List of tags that will be added to every `or` clause of all network access
      policies in the namespace and its children.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: serviceCertificateValidity
    description: |-
      Determines the validity time of certificates issued in this namespace. Default
      value is 1 hour.
    type: string
    exposed: true
    stored: true
    default_value: 1h
    validations:
    - $timeDuration
