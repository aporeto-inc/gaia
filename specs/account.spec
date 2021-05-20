# Model
model:
  rest_name: account
  resource_name: accounts
  entity_name: Account
  package: vince
  group: core/account
  description: |-
    Allows you to view and manage basic information about your account like
    your name, password, and whether or not two-factor authentication is enabled.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@identifiable-stored'
  - '@timeable'

# Indexes
indexes:
- - name
- - email
- - activationToken
- - resetPasswordToken

# Attributes
attributes:
  v1:
  - name: OTPEnabled
    description: Enable or disable two-factor authentication.
    type: boolean
    exposed: true
    stored: true

  - name: OTPQRCode
    description: Returns the base64-encoded QR code for setting up two-factor authentication.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: OTPSecret
    description: Stores the two-factor authentication secret.
    type: string
    stored: true

  - name: SSHCA
    description: Holds the SSH certificate authority used by the account namespace.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: SSHCARenew
    description: Set to `true` to renew the SSH certificate authority of the account
      namespace.
    type: boolean
    exposed: true

  - name: accessEnabled
    description: Defines if the account holder should have access to the system.
    type: boolean
    exposed: true
    stored: true
    orderable: true

  - name: activationExpiration
    description: Contains the expiration date of the activation token.
    type: time
    stored: true
    autogenerated: true

  - name: activationToken
    description: Contains the activation token.
    type: string
    exposed: true
    stored: true
    autogenerated: true
    omit_empty: true

  - name: associatedAPIAuthPolicyID
    description: Holds the ID of the associated API authorization.
    type: string
    stored: true

  - name: associatedAWSPolicies
    description: Contains a map of associated AWS enforcer policies.
    type: external
    subtype: map[string]string
    stored: true

  - name: associatedNamespaceID
    description: Contains the ID of the associated namespace.
    type: string
    stored: true

  - name: associatedPlanKey
    description: Contains the plan key associated with this account.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: associatedQuotaPolicies
    description: Contains a map of the associated quotas.
    type: external
    subtype: map[string]string
    stored: true

  - name: company
    description: Company of the account user.
    type: string
    exposed: true
    stored: true
    example_value: Acme
    filterable: true
    orderable: true

  - name: email
    description: Email of the account holder.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: user@acme.com
    filterable: true
    orderable: true

  - name: failedAuth
    description: Internally keeps track of the number of failed attempt.
    type: integer
    stored: true

  - name: failedTime
    description: Internally keeps track of the time of the last failed attempt.
    type: time
    stored: true

  - name: firstName
    description: First name of the account user.
    type: string
    exposed: true
    stored: true
    example_value: John
    filterable: true
    orderable: true

  - name: lastName
    description: Last name of the account user.
    type: string
    exposed: true
    stored: true
    example_value: Doe
    filterable: true
    orderable: true

  - name: localCA
    description: The certificate authority used by this namespace.
    type: string
    exposed: true
    read_only: true
    autogenerated: true
    transient: true

  - name: localCARenew
    description: Set to `true` to renew the local certificate authority of the account
      namespace.
    type: boolean
    exposed: true

  - name: name
    description: Name of the account.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    allowed_chars: ^[^\*\=]*$
    allowed_chars_message: must not contain any '*' or '='
    example_value: acme
    filterable: true
    orderable: true

  - name: newPassword
    description: |-
      New password for the account. If set the previous password must be given through
      the property `password`.
    type: string
    exposed: true

  - name: password
    description: Password for the account.
    type: string
    exposed: true
    stored: true

  - name: reCAPTCHAKey
    description: |-
      Contains the completely automated public Turing test (CAPTCHA)
      validation if reCAPTCHA is enabled.
    type: string
    exposed: true
    creation_only: true

  - name: resetPasswordExpiration
    description: Contains the expiration time for resetting the password.
    type: time
    stored: true
    autogenerated: true

  - name: resetPasswordToken
    description: Contains the token to use for resetting password.
    type: string
    stored: true
    autogenerated: true

  - name: status
    description: Status of the account.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Active
    - Disabled
    - Invited
    - Pending
    autogenerated: true
    default_value: Pending
    filterable: true
    orderable: true
