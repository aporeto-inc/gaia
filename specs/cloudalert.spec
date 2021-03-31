# Model
model:
  rest_name: cloudalert
  resource_name: cloudalerts
  entity_name: CloudAlert
  package: yeul
  group: pcn/infrastructure
  description: Creates a Prisma Cloud policy and corresponding alert rules.
  get:
    description: Retrieves the Prisma Cloud policy with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the Prisma Cloud policy with the given ID.
  delete:
    description: Deletes the the Prisma Cloud policy with the given ID.
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: cloudpolicies
    description: The list of policies that apply to this alert.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: notifications
    description: Type of notifications.
    type: list
    exposed: true
    subtype: string
    stored: true

  - name: targetSelector
    description: |-
      Selector of namespaces where this alert rule must apply. If empty it applies to
      current namespace.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true
    validations:
    - $tagsExpression
