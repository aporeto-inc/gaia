# Model
model:
  rest_name: quotapolicy
  resource_name: quotapolicies
  entity_name: QuotaPolicy
  package: squall
  description: |-
    Quotas Policies allows to set quotas on the number of objects that can be
    created in a namespace.
  aliases:
  - quota
  - quotas
  - quotapol
  - quotapols
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

# Attributes
attributes:
  v1:
  - name: identities
    description: Identities contains the list of identity names where the quota will
      be applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    required: true
    example_value:
    - processingunit
    - enforcer

  - name: quota
    description: |-
      Quota contains the maximum number of object matching the policy subject that can
      be created.
    type: integer
    exposed: true

  - name: targetNamespace
    description: TargetNamespace contains the base namespace from where the count
      will be done.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/namespace
    filterable: true
    format: free
