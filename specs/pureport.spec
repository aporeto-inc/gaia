# Model
model:
  rest_name: pureport
  resource_name: pureports
  entity_name: PUReport
  package: zack
  description: Post a new pu statistics report.

# Attributes
attributes:
  v1:
  - name: PU
    description: PU to report.
    type: external
    exposed: true
    subtype: processingunit
    required: true
    example_value: ""

  - name: timestamp
    description: Date of the report.
    type: time
    exposed: true
    required: true
    example_value: "2018-06-14T23:10:46.420397985Z"
