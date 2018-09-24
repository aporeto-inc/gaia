# Model
model:
  rest_name: servicedependency
  resource_name: servicedependencies
  entity_name: ServiceDependency
  package: squall
  description: |-
    Allows to define a service dependency where a set of processing units as defined
    by their tags require access to specific services.
  aliases:
  - srvdep
  - srvdeps
  indexes:
  - - namespace
  - - namespace
    - archived
  - - namespace
    - normalizedtags
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
  - '@propagated'
  - '@fallback'
  - '@schedulable'

# Attributes
attributes:
  v1:
  - name: object
    description: Object of the policy.
    type: external
    exposed: true
    subtype: policies_list
    orderable: true

  - name: subject
    description: Subject of the policy.
    type: external
    exposed: true
    subtype: policies_list
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
