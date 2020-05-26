# Model
model:
  rest_name: fileaccessreport
  resource_name: fileaccessreports
  entity_name: FileAccessReport
  package: zack
  group: policy/files
  description: Post a new file access report.
  extends:
  - '@identifiable-stored'
  - '@zoned'
  - '@migratable'

# Attributes
attributes:
  v1:
  - name: action
    description: Action taken.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    - Limit
    example_value: Accepted

  - name: host
    description: Host storing the file.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: localhost

  - name: mode
    description: Mode of file access.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: rxw

  - name: path
    description: Path of the file.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: /etc/passwd

  - name: processingUnitID
    description: ID of the processing unit.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: processingUnitNamespace
    description: Namespace of the processing unit.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
