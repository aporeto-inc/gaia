# Model
model:
  rest_name: accessreport
  resource_name: accessreports
  entity_name: AccessReport
  package: zack
  group: policy/access
  description: Represents any access made by the user.

# Attributes
attributes:
  v1:
  - name: action
    description: Action applied to the access.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: claimHash
    description: Hash of the claims used to communicate.
    type: string
    exposed: true

  - name: enforcerID
    description: Identifier of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: processingUnitID
    description: ID of the processing unit of the report.
    type: string
    exposed: true
    example_value: xxx-xxx-xxx-xxx

  - name: processingUnitName
    description: Name of the processing unit of the report.
    type: string
    exposed: true
    example_value: pu1

  - name: processingUnitNamespace
    description: Namespace of the processing unit of the report.
    type: string
    exposed: true
    example_value: /my/ns

  - name: reason
    description: |-
      This field is only set if `action` is set to `Reject`. It specifies the reason
      for the rejection.
    type: string
    exposed: true

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true

  - name: type
    description: Type of the report.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - SSHLogin
    - SSHLogout
    - SudoEnter
    - SudoExit
    example_value: SSHLogin
