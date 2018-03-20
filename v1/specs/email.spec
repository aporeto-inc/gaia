# Model
model:
  rest_name: email
  resource_name: emails
  entity_name: Email
  package: yuffie
  description: Email is a message that can be send via email
  private: true

# Attributes
attributes:
- name: attachments
  description: Attachments is a list of attachments to send
  type: external
  exposed: true
  subtype: list_attachments

- name: bcc
  description: Bcc represents email that should be in copy but hidden.
  type: external
  exposed: true
  subtype: list_emails

- name: cc
  description: Cc represents the addresses that should be in copy
  type: external
  exposed: true
  subtype: list_emails

- name: content
  description: Content of the email to send
  type: string
  exposed: true
  format: free

- name: from
  description: From represents the sender of the email
  type: string
  exposed: true
  required: true
  format: email

- name: subject
  description: Subject represents the subject of the email
  type: string
  exposed: true
  format: free

- name: to
  description: To represents receivers of the email
  type: external
  exposed: true
  subtype: list_emails

- name: type
  description: Type represents the type of the content.
  type: enum
  exposed: true
  allowed_choices:
  - HTML
  - Plain
  default_value: Plain
