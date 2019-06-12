# Model
model:
  rest_name: sshauthorizationpolicy
  resource_name: sshauthorizationpolicies
  entity_name: SSHAuthorizationPolicy
  package: squall
  group: policy/ssh
  description: |-
    An SSHAuthorizationPolicy allows to define the permissions for the owner
    of a SSH Identity Certificate. You can define if a user with some claims
    can connect to an sshd server managed by an instance of enforcerd according
    to its tags, what permissions he has and for how long delivered
    certificates are valid.
  aliases:
  - sshpol
  - sshpols
  get:
    description: Retrieves the SSHAuthorizationPolicy with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the SSHAuthorizationPolicy with the given ID.
  delete:
    description: Deletes the SSHAuthorizationPolicy with the given ID.
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

  - name: extensions
    description: |-
      The list of SSH permissions to apply to SSH certificate. You can check the list
      of standard extensions at
      <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L281>.
    type: list
    exposed: true
    subtype: string

  - name: forceCommand
    description: |-
      If set, this will configure the `force-command` option in the SSH Certificate.
      More info can be found at
      <https://github.com/openssh/openssh-portable/blob/38e83e4f219c752ebb1560633b73f06f0392018b/PROTOCOL.certkeys#L249>.
    type: string
    exposed: true

  - name: object
    description: |-
      Object contains the tag expression matching the enforcers the subject is allowed
      to connect to.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: principals
    description: |-
      You can set some principals that will be applied to delivered certificate. If
      not
      set, the user's claim Subject will be used.
    type: list
    exposed: true
    subtype: string

  - name: subject
    description: |-
      Subject contains the tag expression the authentication claims need to match for
      the policy to apply.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: validity
    description: Set the validity of the delivered SSH certificate.
    type: string
    exposed: true
    default_value: 1h
    validations:
    - $timeDuration
