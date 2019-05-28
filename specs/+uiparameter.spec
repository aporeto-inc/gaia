# Model
model:
  rest_name: uiparameter
  resource_name: uiparameters
  entity_name: UIParameter
  package: ignis
  group: core/workflow
  description: Represents a parameter that will be shown in the UI.
  detached: true

# Attributes
attributes:
  v1:
  - name: advanced
    description: Defines if the parameter is an advanced one.
    type: boolean
    exposed: true
    stored: true

  - name: allowedChoices
    description: allowedChoices lists all the choices in case of an enum.
    type: external
    exposed: true
    subtype: map[string]string
    stored: true

  - name: allowedValues
    description: List of values that can be used.
    type: list
    exposed: true
    subtype: object
    stored: true

  - name: defaultValue
    description: Default value of the parameter.
    type: object
    exposed: true
    stored: true

  - name: description
    description: Description of the paramerter.
    type: string
    exposed: true
    stored: true

  - name: key
    description: Key identifying the parameter.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: unique_key

  - name: longDescription
    description: Long explanation of the parameter.
    type: string
    exposed: true
    stored: true

  - name: name
    description: Name of the paramerter.
    type: string
    exposed: true
    stored: true

  - name: optional
    description: Defines if the parameter is optional.
    type: boolean
    exposed: true
    stored: true

  - name: type
    description: The type of the parameter.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Boolean
    - Duration
    - Enum
    - IntegerSlice
    - Integer
    - Float
    - FloatSlice
    - Password
    - String
    - StringSlice
    - CVSSThreshold
    - JSON
    - TagsExpression
    example_value: String

  - name: validationFunction
    description: ValidationFunction represents the function to validate the parameter.
    type: string
    exposed: true
    stored: true

  - name: value
    description: Value of the parameter.
    type: object
    exposed: true
    stored: true
    deprecated: true

  - name: visibilityCondition
    description: |-
      List of ors of ands of uiparametervisibility that must be verified for the
      parameter to be displayed to the user.
    type: external
    exposed: true
    subtype: uiparametersexpression
    stored: true
