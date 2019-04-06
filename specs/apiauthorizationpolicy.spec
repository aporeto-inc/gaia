# Model
model:
  rest_name: apiauthorizationpolicy
  resource_name: apiauthorizationpolicies
  entity_name: APIAuthorizationPolicy
  package: squall
  group: policy/authorization
  description: |-
    An API Authorization Policy defines what kind of operations a user of a system
    can do in a namespace. The operations can be any combination of GET, POST, PUT,
    DELETE,PATCH or HEAD. By default, an API Authorization Policy will only give
    permissions in the context of the current namespace but you can make it
    propagate to all the child namespaces. It is also possible restrict permissions
    to apply only on a particular subset of the apis by setting the target
    identities.
  aliases:
  - apiauth
  - apiauths
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
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@hidden'
  - '@fallback'
  - '@schedulable'
  - '@zonable'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: authorizedIdentities
    description: AuthorizedIdentities defines the list of api identities the policy
      applies to.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - '@auth:role=namespace.editor'

  - name: authorizedNamespace
    description: AuthorizedNamespace defines on what namespace the policy applies.
    type: string
    exposed: true
    required: true
    example_value: /namespace

  - name: authorizedSubnets
    description: |-
      If set, the api authorization will only be valid if the request comes from one
      the declared subnets.
    type: list
    exposed: true
    subtype: string
    validations:
    - $optionalnetworks

  - name: expirationTime
    description: If set the policy will be auto deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: subject
    description: Subject is the subject.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
