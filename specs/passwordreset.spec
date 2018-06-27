# Model
model:
  rest_name: passwordreset
  resource_name: passwordreset
  entity_name: PasswordReset
  package: vince
  description: Used to reset an account password.

# Attributes
attributes:
  v1:
  - name: password
    description: Password contains the new password.
    type: string
    exposed: true
    required: true
    example_value: NewPassword123@
    format: free

  - name: token
    description: Token contains the reset password token.
    type: string
    exposed: true
    required: true
    example_value: valid.jwt.token
    format: free
