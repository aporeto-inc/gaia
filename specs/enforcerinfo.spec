# Model
model:
  rest_name: enforcerinfo
  resource_name: enforcerinfo
  entity_name: EnforcerInfo
  package: squall
  group: core/enforcer
  description: Post a new enforcer information.

# Attributes
attributes:
  v1:
  - name: collectedInfo
    description: Represents the latest information collected by the enforcer.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: enforcerInfoID
    description: |-
      EnforcerInfoID is the ID of the enforcer information. EnforcerInfoID is used to
      aggergate the multipart requests.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx
