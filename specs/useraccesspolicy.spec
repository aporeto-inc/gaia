# Model
model:
  rest_name: useraccesspolicy
  resource_name: useraccesspolicies
  entity_name: UserAccessPolicy
  package: squall
  group: policy/access
  description: The enforcer policy that controls user access.
  aliases:
  - usrpol
  - usrpols
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
  - name: allowedSudoUsers
    description: Indicates the list of user who can use `sudo` commands.
    type: list
    exposed: true
    subtype: string

  - name: expirationTime
    description: If set the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: object
    description: |-
      Contains the tag expression matching the enforcers the subject is allowed
      to connect to.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: |-
      Contains the tag expression the tags need to match for the policy to
      apply.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression
