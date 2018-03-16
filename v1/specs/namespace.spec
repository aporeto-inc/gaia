attributes:
- description: AssociatedLocalCAID holds the remote ID of the certificate authority
    to use.
  format: free
  name: associatedLocalCAID
  read_only: true
  secret: true
  stored: true
  type: string
- autogenerated: true
  description: LocalCA holds the eventual certificate authority used by this namespace.
  exposed: true
  format: free
  name: localCA
  read_only: true
  stored: true
  type: string
- description: LocalCAEnabled defines if the namespace should use a local Certificate
    Authority. Switching it off and on again will regenerate a new CA.
  exposed: true
  filterable: true
  name: localCAEnabled
  orderable: true
  stored: true
  type: boolean
- allowed_chars: ^[a-zA-Z0-9-_/]+$
  creation_only: true
  description: Name is the name of the namespace.
  exposed: true
  filterable: true
  format: free
  getter: true
  index: true
  name: name
  orderable: true
  primary_key: true
  required: true
  stored: true
  type: string
model:
  aliases:
  - ns
  delete: true
  get: true
  update: true
  description: A Namespace represents the core organizational unit of the system.
    All objects always exists in a single namespace. A Namespace can also have child
    namespaces. They can be used to split the system into organizations, business
    units, applications, services or any combination you like.
  entity_name: Namespace
  extends:
  - '@base'
  - '@described'
  - '@identifiable-nopk-stored'
  - '@metadatable'
  package: squall
  resource_name: namespaces
  rest_name: namespace
