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
    description: Defines if the step is an advanced one.
    type: boolean
    exposed: true
    orderable: true

  - name: description
    description: Description of the step.
    type: string
    exposed: true

  - name: name
    description: Name of the step.
    type: string
    exposed: true

  - name: parameters
    description: List of parameters for this step.
    type: refList
    exposed: true
    subtype: uiparameter
    extensions:
      refMode: pointer
