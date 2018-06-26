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
  descriptions:
    get: Returns the list of external services that are targets of service dependency.
  get: true

- rest_name: processingunit
  descriptions:
    get: Returns the list of Processing Units that depend on an service.
  get: true
