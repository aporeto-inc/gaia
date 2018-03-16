attributes:
- description: 'AuthorizedIdentities defines the list of api identities the policy
    applies to. '
  exposed: true
  name: authorizedIdentities
  required: true
  subtype: identity_list
  type: external
- description: AuthorizedNamespace defines on what namespace the policy applies.
  exposed: true
  filterable: true
  format: free
  name: authorizedNamespace
  required: true
  type: string
- description: Subject is the subject.
  exposed: true
  name: subject
  orderable: true
  subtype: policies_list
  type: external
model:
  aliases:
  - apiauth
  - apiauths
  delete: true
  get: true
  update: true
  description: An API Authorization Policy defines what kind of operations a user
    of a system can do in a namespace. The operations can be any combination of GET,
    POST, PUT, DELETE,PATCH or HEAD. By default, an API Authorization Policy will
    only give permissions in the context of the current namespace but you can make
    it propagate to all the child namespaces.  It is also possible restrict permissions
    to apply only on a particular subset of the apis by setting the target identities.
  entity_name: APIAuthorizationPolicy
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@schedulable'
  package: squall
  resource_name: apiauthorizationpolicies
  rest_name: apiauthorizationpolicy
