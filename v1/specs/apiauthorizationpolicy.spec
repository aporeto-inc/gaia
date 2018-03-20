# Model
model:
  rest_name: apiauthorizationpolicy
  resource_name: apiauthorizationpolicies
  entity_name: APIAuthorizationPolicy
  package: squall
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
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'

# Attributes
attributes:
- name: authorizedIdentities
  description: AuthorizedIdentities defines the list of api identities the policy
    applies to.
  type: external
  exposed: true
  subtype: identity_list
  required: true

- name: authorizedNamespace
  description: AuthorizedNamespace defines on what namespace the policy applies.
  type: string
  exposed: true
  required: true
  filterable: true
  format: free

- name: subject
  description: Subject is the subject.
  type: external
  exposed: true
  subtype: policies_list
  orderable: true
