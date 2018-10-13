# Model
model:
  rest_name: k8scredential
  resource_name: k8scredentials
  entity_name: K8SCredential
  package: cactuar
  description: Create a credential for a kubernetes cluster.
  aliases:
  - k8scred
  - k8screds
  indexes:
  - - namespace
  - - namespace
    - normalizedTags
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
  - '@identifiable-pk-stored'
  - '@metadatable'

# Attributes
attributes:
  v1:
  - name: certificate
    description: The string representation of the Certificate used by the Kubernetes
      cluster.
    type: string
    exposed: true
    stored: true
    read_only: true

  - name: certificateSN
    description: Link to the certificate created for this cluster.
    type: string
    stored: true

  - name: email
    description: |-
      The email address that will receive a copy of the Kubernetes cluster YAMLs
      definition.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: name
    description: Name is the name of the entity.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    default_order: true
    example_value: my cluster
    filterable: true
    getter: true
    setter: true
    max_length: 256
    orderable: true

  - name: regenerate
    description: Regenerates the k8s files and certificates.
    type: boolean
    exposed: true

  - name: secretDefinition
    description: The secret file to deploy on your cluster.
    type: external
    exposed: true
    subtype: k8s_secret_definition
    read_only: true
    orderable: true
