# Model
model:
  rest_name: namespacemappingpolicy
  resource_name: namespacemappingpolicies
  entity_name: NamespaceMappingPolicy
  package: squall
  group: core/namespace
  description: |-
    A Namespace Mapping Policy defines in which namespace a Processing Unit should
    be placed when it is created, based on its tags.  When an Aporeto Agent creates
    a new Processing Unit, the system will place it in its own namespace if no
    matching Namespace Mapping Policy can be found. If one match is found, then the
    Processing will be bumped down to the namespace declared in the policy. If it
    finds in that child namespace another matching Namespace Mapping Policy, then
    the Processing Unit will be bumped down again, until it reach a namespace with
    no matching policies.  This is very useful to dispatch processes and containers
    into a particular namespace, based on a lot of factor.   You can put in place a
    quarantine namespace that will grab all Processing Units with too much
    vulnerabilities for instances.
  aliases:
  - nspolicy
  - nspolicies
  - nsmap
  - nsmaps
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
  - '@zonable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: mappedNamespace
    description: mappedNamespace is the mapped namespace.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/]+$
    allowed_chars_message: must only contain alpha numerical characters, '-' or '_'
    example_value: /blue/namespace
    orderable: true

  - name: subject
    description: Subject is the subject.
    type: external
    exposed: true
    subtype: '[][]string'
    example_value:
    - - color=blue
    orderable: true
    validations:
    - $tagsExpression
