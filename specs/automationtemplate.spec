# Model
model:
  rest_name: automationtemplate
  resource_name: automationtemplates
  entity_name: AutomationTemplate
  package: sephiroth
  group: integration/automation
  description: Templates that ca be used in automations.
  aliases:
  - autotmpl
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@described'
  - '@named'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: entitlements
    description: Entitlements contains the entitlements needed for executing the function.
    type: external
    exposed: true
    subtype: _automation_entitlements

  - name: function
    description: Function contains the code.
    type: string
    exposed: true

  - name: key
    description: Key contains the unique identifier key for the template.
    type: string
    exposed: true

  - name: kind
    description: Kind represents the kind of template.
    type: enum
    exposed: true
    allowed_choices:
    - Action
    - Condition
    default_value: Condition

  - name: parameters
    description: Parameters contains the computed parameters.
    type: external
    exposed: true
    subtype: map[string]interface{}

  - name: steps
    description: Steps contains all the steps with parameters.
    type: refList
    exposed: true
    subtype: uistep
    extensions:
      refMode: pointer
