# Model
model:
  rest_name: processingunitrefresh
  resource_name: processingunitrefreshs
  entity_name: ProcessingUnitRefresh
  package: gaga
  group: core/policy
  description: |-
    ProcessingUnitRefresh is sent to client when a poke has been triggered using the
    parameter `?notify=true`. This is used by instances of enforcerd to be notify an
    external change on the processing unit must be processed.

# Attributes
attributes:
  v1:
  - name: ID
    description: ID contains the original ID of the Processing Unit.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: namespace
    description: Namespace contains the original namespace of the Processing Unit.
    type: string
    exposed: true
    stored: true
    filterable: true
    orderable: true
