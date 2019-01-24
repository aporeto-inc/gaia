# Model
model:
  rest_name: appparameter
  resource_name: appparameters
  entity_name: AppParameter
  package: highwind
  description: Represents a parameter that can be passed to an app.
  detached: true

# Attributes
attributes:
  v1:
  - name: advanced
    description: Defines if the parameter is an advanced one.
    type: boolean
    exposed: true

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

  - name: description
    description: Description of the paramerter.
    type: string
    exposed: true

  - name: key
    description: Key identifying the parameter.
    type: string
    exposed: true

  - name: longDescription
    description: Long explanation of the parameter.
    type: string
    exposed: true

  - name: name
    description: Name of the paramerter.
    type: string
    exposed: true

  - name: optional
    description: Defines if the parameter is optional.
    type: boolean
    exposed: true

  - name: type
    description: The type of the parameter.
    type: enum
    exposed: true
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

  - name: value
    description: Value of the parameter.
    type: object
    exposed: true
