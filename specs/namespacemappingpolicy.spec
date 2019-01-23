# Model
model:
  rest_name: namespacemappingpolicy
  resource_name: namespacemappingpolicies
  entity_name: NamespaceMappingPolicy
  package: squall
  description: "A Namespace Mapping Policy defines in which namespace a Processing
    Unit should\nbe placed when it is created, based on its tags.\n\nWhen an Aporeto
    Agent creates a new Processing Unit, the system will place it in\nits own namespace
    if no matching Namespace Mapping Policy can be found. If one\nmatch is found,
    then the Processing will be bumped down to the namespace\ndeclared in the policy.
    If it finds in that child namespace another matching\nNamespace Mapping Policy,
    then the Processing Unit will be bumped down again,\nuntil it reach a namespace
    with no matching policies.\n\nThis is very useful to dispatch processes and containers
    into a particular\nnamespace, based on a lot of factor. \n\nYou can put in place
    a quarantine namespace that will grab all Processing Units\nwith too much vulnerabilities
    for instances."
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
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: mappedNamespace
    description: mappedNamespace is the mapped namespace.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /blue/namespace
    orderable: true

  - name: subject
    description: Subject is the subject.
    type: external
    exposed: true
    subtype: policies_list
    required: true
    example_value:
    - - color=blue
    orderable: true
