# Model
model:
  rest_name: email
  resource_name: emails
  entity_name: Email
  package: yuffie
  description: Email is a message that can be send via email.
  private: true

# Attributes
attributes:
  v1:
  - name: attachments
    description: Attachments is a list of attachments to send.
    type: external
    exposed: true
    subtype: map_of_string_of_strings

  - name: bcc
    description: Bcc represents email that should be in copy but hidden.
    type: list
    exposed: true
    subtype: string

  - name: cc
    description: Cc represents the addresses that should be in copy.
    type: list
    exposed: true
    subtype: string

  - name: content
    description: Content of the email to send.
    type: string
    exposed: true

  - name: from
    description: |-
      From represents the sender of the email. If not set, the default sender will be
      used.
    type: string
    exposed: true
    example_value: sender@server.com

  - name: subject
    description: Subject represents the subject of the email.
    type: string
    exposed: true

  - name: to
    description: To represents receivers of the email.
    type: list
    exposed: true
    subtype: string

  - name: type
    description: Type represents the type of the content.
    type: enum
    exposed: true
    allowed_choices:
    - HTML
    - Plain
    default_value: Plain
