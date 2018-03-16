attributes:
- description: expirationTime is the time after which the message will be deleted.
  exposed: true
  filterable: true
  name: expirationTime
  orderable: true
  stored: true
  type: time
- allowed_choices:
  - Danger
  - Info
  - Warning
  default_value: Info
  description: Level defines how the message is important.
  exposed: true
  filterable: true
  name: level
  orderable: true
  stored: true
  type: enum
- description: If local is set, the message will only be visible in the current namespace.
  exposed: true
  filterable: true
  name: local
  orderable: true
  stored: true
  type: boolean
- creation_only: true
  description: If enabled, the message will be sent to the email associated in namespaces
    annotations.
  exposed: true
  filterable: true
  name: notifyByEmail
  type: boolean
- allowed_chars: ^[0-9]+[smh]$
  description: Validity set using golang time duration, when the message will be automatically
    deleted.
  exposed: true
  format: free
  name: validity
  stored: true
  type: string
model:
  aliases:
  - mess
  delete: true
  get: true
  update: true
  description: The Message API allows to post public messages that will be visible
    through all children namespaces
  entity_name: Message
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@named'
  package: squall
  resource_name: messages
  rest_name: message
