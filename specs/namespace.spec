# Model
model:
  rest_name: namespace
  resource_name: namespaces
  entity_name: Namespace
  package: squall
  description: |-
    A Namespace represents the core organizational unit of the system. All objects
    always exists in a single namespace. A Namespace can also have child namespaces.
    They can be used to split the system into organizations, business units,
    applications, services or any combination you like.
  aliases:
  - ns
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-nopk-stored'
  - '@metadatable'

# Attributes
attributes:
  v1:
  - name: associatedLocalCAID
    description: AssociatedLocalCAID holds the remote ID of the certificate authority
      to use.
    type: string
    stored: true
    read_only: true
    format: free
    secret: true

  - name: localCA
    description: LocalCA holds the eventual certificate authority used by this namespace.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    format: free

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
    example_value: mynamespace
    filterable: true
    format: free
    getter: true
    index: true
    orderable: true
    primary_key: true

  - name: networkAccessPolicyTags
    description: |-
      List of tags that will be added to every `or` clause of all network access
      policies in the namespace and its children.
    type: external
    exposed: true
    subtype: tags_list
    stored: true
