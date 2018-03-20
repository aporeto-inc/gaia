# Model
model:
  rest_name: kubernetescluster
  resource_name: kubernetesclusters
  entity_name: KubernetesCluster
  package: vince
  description: Create a remote Kubernetes Cluster integration.
  get: true
  update: true
  delete: true
  extends:
  - '@identifiable-pk-stored'
  - '@timeable'

# Attributes
attributes:
- name: APIAuthorizationPolicyID
  description: Link to the API authorization policy.
  type: string
  stored: true
  format: free

- name: certificateID
  description: Link to the certificate created in Vince for this cluster.
  type: string
  stored: true
  format: free

- name: kubernetesDefinitions
  description: |-
    base64 of the .tar.gz file that contains all the .YAMLs files needed to create
    the aporeto side on your kubernetes Cluster
  type: string
  exposed: true
  read_only: true
  filterable: true
  format: free
  orderable: true

- name: name
  description: The name of your cluster
  type: string
  exposed: true
  stored: true
  filterable: true
  format: free
  orderable: true

- name: namespaceID
  description: Link to your namespace.
  type: string
  stored: true
  format: free

- name: parentID
  description: ID of the parent account.
  type: string
  exposed: true
  stored: true
  read_only: true
  filterable: true
  format: free
  orderable: true

- name: regenerate
  description: Regenerates the k8s files and certificates.
  type: boolean
  exposed: true

- name: targetNamespace
  description: |-
    The namespace in which the Kubernetes specific namespace will be created. By
    default your account namespace.
  type: string
  exposed: true
  stored: true
  filterable: true
  format: free
  orderable: true

- name: targetNetworks
  description: List of target networks [deprecated]
  type: external
  exposed: true
  subtype: target_networks_list
  stored: true
  filterable: true
  orderable: true
