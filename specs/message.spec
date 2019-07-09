# Model
model:
  rest_name: message
  resource_name: messages
  entity_name: Message
  package: squall
  group: core/monitoring
  description: |-
    Allows you to post public messages that will be visible through all
    children namespaces.
  aliases:
  - mess
  get:
    description: Retrieves the message with the given ID.
    global_parameters:
    - $propagatable
  update:
    description: Updates the message with the given ID.
  delete:
    description: Deletes the message with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'
  - '@propagated'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: expirationTime
    description: The time after which the message will be deleted.
    type: time
    exposed: true
    stored: true
    orderable: true

  - name: level
    description: Importance of the message.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Danger
    - Info
    - Warning
    default_value: Info
    orderable: true

  - name: validity
    description: |-
      Sets when the message will be automatically deleted using
      [Golang duration syntax](https://golang.org/pkg/time/#example_Duration).
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
