# Model
model:
  rest_name: automationtemplateparameter
  resource_name: automationtemplateparameters
  entity_name: AutomationTemplateParameter
  package: sephiroth
  description: Represents an Automation template parameter.
  extends:
  - '@named'
  - '@described'
  detached: true

# Attributes
attributes:
  v1:
  - name: allowedChoices
    description: Set the possible values for the parameter.
    type: external
    exposed: true
    subtype: map_of_string_of_objects

  - name: defaultValue
    description: Default value of the parameter.
    type: object
    exposed: true

  - name: description
    description: Name of the parameter.
    type: string
    exposed: true

  - name: position
    description: Preferred position for the parameter.
    type: integer
    exposed: true

  - name: required
    description: Set if the parameter must be set.
    type: boolean
    exposed: true

  - name: type
    description: Type of the parameter.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Array
    - Boolean
    - Enum
    - Filter
    - Float
    - Integer
    - Object
    - String
    example_value: String
