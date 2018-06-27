# Model
model:
  rest_name: automationtemplate
  resource_name: automationtemplates
  entity_name: AutomationTemplate
  package: sephiroth
  description: Templates that ca be used in automations.
  aliases:
  - autotmpl
  get: true
  extends:
  - '@described'
  - '@named'

# Attributes
attributes:
  v1:
  - name: entitlements
    description: Entitlements contains the entitlements needed for executing the function.
    type: external
    exposed: true
    subtype: automation_entitlements

  - name: function
    description: Function contains the code.
    type: string
    exposed: true
    format: free

  - name: key
    description: Key contains the unique identifier key for the template.
    type: string
    exposed: true
    format: free

  - name: kind
    description: Kind represents the kind of template.
    type: enum
    exposed: true
    allowed_choices:
    - Action
    - Condition
    default_value: Condition

  - name: parameters
    description: Parameters contains the parameter description of the function.
    type: external
    exposed: true
    subtype: automation_template_parameters
