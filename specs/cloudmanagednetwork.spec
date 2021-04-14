# Model
model:
  rest_name: cloudmanagednetwork
  resource_name: cloudmanagednetworks
  entity_name: CloudManagedNetwork
  package: yeul
  group: pcn/infrastructure
  description: |-
    A cloud managed network represents a set of enterprise subnets that can be used
    in policies.
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
    - $propagatable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@zoned'
  - '@migratable'
  - '@namespaced'
  - '@identifiable-stored'
  - '@prismabase'
  - '@timeable'

# Indexes
indexes:
- - namespace
  - type
- - type
- - key

# Attributes
attributes:
  v1:
  - name: entries
    description: List of CIDRs.
    type: list
    exposed: true
    subtype: string
    stored: true
    validations:
    - $optionalcidrs

  - name: key
    description: Internal unique key for a resource to guarantee no overlaps at write.
    type: string
    stored: true

  - name: type
    description: The type of cloud managed network.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Enterprise
    - AWSPrefixLists
    - AWSElasticIPs
    - GCP
    - Custom
    default_value: Enterprise
