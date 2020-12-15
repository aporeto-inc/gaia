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
  - '@zoned-monotonic'
  - '@migratable'

# Ordering
default_order:
- :no-inherit
- timestamp

# Indexes
indexes:
- - namespace
  - timestamp

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
    omit_empty: true
    extensions:
      bson_name: a

  - name: host
    description: Host storing the file.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: localhost
    omit_empty: true
    extensions:
      bson_name: b

  - name: mode
    description: Mode of file access.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: rxw
    omit_empty: true
    extensions:
      bson_name: c

  - name: path
    description: Path of the file.
    type: string
    exposed: true
    stored: true
    required: true
    default_value: /etc/passwd
    omit_empty: true
    extensions:
      bson_name: d

  - name: processingUnitID
    description: ID of the processing unit.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: xxx-xxx-xxx-xxx
    omit_empty: true
    extensions:
      bson_name: e

  - name: processingUnitNamespace
    description: Namespace of the processing unit.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /my/ns
    omit_empty: true
    extensions:
      bson_name: f

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    stored: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
    orderable: true
    omit_empty: true
    extensions:
      bson_name: g
