# Model
model:
  rest_name: servicedependency
  resource_name: servicedependencies
  entity_name: ServiceDependency
  package: squall
  group: policy/services
  description: |-
    Allows to define a service dependency where a set of processing units as defined
    by their tags require access to specific services.
  aliases:
  - srvdep
  - srvdeps
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
  - '@fallback'
  - '@schedulable'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    orderable: true

  - name: subject
    description: Subject of the policy.
    type: external
    exposed: true
    subtype: list_of_lists_of_strings
    orderable: true

# Relations
relations:
- rest_name: service
  get:
    description: Returns the list of external services that are targets of service
      dependency.

- rest_name: processingunit
  get:
    description: Returns the list of Processing Units that depend on an service.
