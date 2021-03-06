# Model
model:
  rest_name: issueservicetoken
  resource_name: issueservicetokens
  entity_name: IssueServiceToken
  package: barret
  group: internal/servicetokens
  description: |-
    This is an internal API. Services can call this API to issue service tokens on
    behalf of users.
  private: true

# Attributes
attributes:
  v1:
  - name: audience
    description: Audience is the valid audience for this token.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: my-service

  - name: namespaceID
    description: The ID of the corresponding namespace.
    type: string
    exposed: true
    creation_only: true
    example_value: 12341234

  - name: serviceClaims
    description: |-
      ServiceClaims is a list of service claims that have been validated provided as
      key/value pairs. If the same key is provided multiple times it will be converted
      to an array. The claims  will appear under the Data section of the token.
    type: list
    exposed: true
    subtype: string
    creation_only: true
    example_value:
    - org=acme
    - sub=test

  - name: signingKeyID
    description: SigningKeyID holds the ID of the private certificate to use to sign the token.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: 2739128

  - name: subject
    description: The subject claims of the token.
    type: string
    exposed: true
    required: true
    creation_only: true
    example_value: joe@acme.com

  - name: token
    description: Token contains the generated token.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: userClaims
    description: |-
      UserClaims is a list of user claims that have been validated provided as
      key/value pairs. If the same key is provided multiple times it will be converted
      to an array. The claims  will appear under the Data section of the token.
    type: list
    exposed: true
    subtype: string
    creation_only: true
    example_value:
    - org=acme
    - sub=test

  - name: validity
    description: Validity contains the token validity duration.
    type: string
    exposed: true
    required: true
    creation_only: true
    default_value: 15m
    validations:
    - $timeDuration
