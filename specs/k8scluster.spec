# Model
model:
  rest_name: k8scluster
  resource_name: k8sclusters
  entity_name: K8SCluster
  package: cactuar
  description: Create a remote Kubernetes Cluster integration.
  indexes:
  - - namespace
  - - namespace
    - normalizedTags
  - - :shard
    - zone
    - namespace
    - id
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
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: APIAuthorizationPolicyID
    description: Link to the API authorization policy.
    type: string
    stored: true

  - name: activationType
    description: Defines the mode of activation on the KubernetesCluster.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - KubeSquall
    - PodAtomic
    - PodContainers
    default_value: KubeSquall
    orderable: true

  - name: adminEmail
    description: |-
      The email address that will receive a copy of the Kubernetes cluster YAMLs
      definition.
    type: string
    exposed: true
    stored: true
    orderable: true

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

  - name: kubernetesDefinitions
    description: |-
      base64 of the .tar.gz file that contains all the .YAMLs files needed to create
      the aporeto side on your kubernetes Cluster.
    type: string
    exposed: true
    read_only: true
    orderable: true

  - name: name
    description: Name is the name of the entity.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    default_order: true
    example_value: the name
    filterable: true
    getter: true
    setter: true
    max_length: 256
    orderable: true

  - name: namespaceID
    description: Link to the cluster namespace.
    type: string
    stored: true

  - name: networkPolicyType
    description: |-
      Defines what type of network policy will be applied on your cluster.
      Kubernetes means that All the Kubernetes policies will be synced to Squall.
      No Policies means that policies are not synced and it's up to the user to create
      consistent policies in Squall.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Kubernetes
    - NoPolicy
    default_value: Kubernetes
    orderable: true

  - name: regenerate
    description: Regenerates the k8s files and certificates.
    type: boolean
    exposed: true
