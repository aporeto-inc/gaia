# Model
model:
  rest_name: hostservicemappingpolicy
  resource_name: hostservicemappingpolicies
  entity_name: HostServiceMappingPolicy
  package: squall
  group: policy/hosts
  description: |-
    Host service mapping allows you to map host services to the defenders which should
    implement them. You must map host services to one or more defenders for the host 
    services to have any effect.
  aliases:
  - hostsrvmappol
  - hostsrvmappols
  get:
    description: Retrieves the mapping with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the mapping with the given ID.
  delete:
    description: Deletes the mapping with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: object
    description: |-
      A tag or tag expression identifying the host service(s) to be mapped.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      A tag or tag expression identifying the defender(s) that should implement
      the specified host service(s).
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: enforcer
  get:
    description: Returns the list of defenders that are affected by this mapping.

- rest_name: hostservice
  get:
    description: Returns the list of host services that are referenced by this mapping.
