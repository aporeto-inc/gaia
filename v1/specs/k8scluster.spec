# Model
model:
  rest_name: k8scluster
  resource_name: k8sclusters
  entity_name: K8SCluster
  package: cactuar
  description: Create a remote Kubernetes Cluster integration.
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: APIAuthorizationPolicyID
    description: Link to the API authorization policy.
    type: string
    stored: true
    format: free

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
    filterable: true
    orderable: true

  - name: certificateSN
    description: Link to the certificate created for this cluster.
    type: string
    stored: true
    format: free

  - name: kubernetesDefinitions
    description: |-
      base64 of the .tar.gz file that contains all the .YAMLs files needed to create
      the aporeto side on your kubernetes Cluster.
    type: string
    exposed: true
    read_only: true
    format: free
    orderable: true

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
    filterable: true
    orderable: true

  - name: regenerate
    description: Regenerates the k8s files and certificates.
    type: boolean
    exposed: true
