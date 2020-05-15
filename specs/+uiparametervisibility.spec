# Model
model:
  rest_name: uiparametervisibility
  resource_name: uiparametervisibilities
  entity_name: UIParameterVisibility
  package: ignis
  group: core/workflow
  description: Represents a visibility condition for a [UIParameter](#uiparameter).
  detached: true

# Attributes
attributes:
  v1:
  - name: key
    description: Key holding the value to compare.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: enableThing

  - name: operator
    description: Operator to apply.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Equal
    - NotEqual
    - GreaterThan
    - LesserThan
    - Defined
    - Undefined
    - Match
    - NotMatch

  - name: value
    description: Values that must match the key.
    type: object
    exposed: true
    stored: true
    required: true
    example_value: true
