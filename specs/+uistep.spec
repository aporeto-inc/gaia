# Model
model:
  rest_name: uistep
  resource_name: uisteps
  entity_name: UIStep
  package: highwind
  group: integration/app
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

  - name: parameters
    description: List of parameters for this step.
    type: refList
    exposed: true
    subtype: uiparameter
    stored: true
    extensions:
      refMode: pointer
