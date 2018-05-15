# Model
model:
  rest_name: message
  resource_name: messages
  entity_name: Message
  package: squall
  description: |-
    The Message API allows to post public messages that will be visible through all
    children namespaces.
  aliases:
  - mess
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: expirationTime
    description: expirationTime is the time after which the message will be deleted.
    type: time
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: level
    description: Level defines how the message is important.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Danger
    - Info
    - Warning
    default_value: Info
    filterable: true
    orderable: true

  - name: local
    description: If local is set, the message will only be visible in the current
      namespace.
    type: boolean
    exposed: true
    stored: true
    filterable: true
    orderable: true

  - name: notifyByEmail
    description: |-
      If enabled, the message will be sent to the email associated in namespaces
      annotations.
    type: boolean
    exposed: true
    creation_only: true
    filterable: true

  - name: validity
    description: |-
      Validity set using golang time duration, when the message will be automatically
      deleted.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    format: free
