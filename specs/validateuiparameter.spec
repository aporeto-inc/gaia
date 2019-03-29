# Model
model:
  rest_name: validateuiparameter
  resource_name: validateuiparameters
  entity_name: ValidateUIParameter
  package: ignis
  group: core/workflow
  description: This api validates a list of UIParameters.
  aliases:
  - validparam

# Attributes
attributes:
  v1:
  - name: errors
    description: Errors contains the list of errors.
    type: external
    exposed: true
    subtype: map[string]string

  - name: parameters
    description: List of parameters to validate.
    type: refList
    exposed: true
    subtype: uiparameter
    stored: true
    extensions:
      refMode: pointer

  - name: values
    description: Values contains the computed values.
    type: external
    exposed: true
    subtype: map[string]interface{}
