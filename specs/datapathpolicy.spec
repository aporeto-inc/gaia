# Model
model:
  rest_name: datapathpolicy
  resource_name: datapathpolicies
  entity_name: DatapathPolicy
  package: squall
  group: policy/access
  description: The enforcer policy that controls the datapath for processing units.
  aliases:
  - dppol
  - dppols
  get:
    description: Retrieves the policy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the policy with the given ID.
  delete:
    description: Deletes the policy with the given ID.
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
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: datapathType
    description: |-
      The datapath type that processing units selected by `subject` should
      implement.

      - `Aporeto` (default): Aporeto is managing and handling the datapath.
      - `EnvoyAuthorizer`: Aporeto is serving envoy compatible gRPC APIs
      for every processing unit that for example can be used by an envoy
      proxy to use the Aporeto PKI and implement Aporeto network access
      policies. NOTE: The enforcer is not going to own the datapath in this
      example. It is merely providing an authorizer API.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Aporeto
    - EnvoyAuthorizer
    default_value: Aporeto

  - name: subject
    description: |-
      Contains the tag expression the tags need to match for the policy to
      apply.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    example_value:
    - - $identity=processingunit
    orderable: true
    validations:
    - $tagsExpression
