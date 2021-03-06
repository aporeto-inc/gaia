# Model
model:
  rest_name: issue
  resource_name: issue
  entity_name: Issue
  package: midgard
  group: core/authentication
  description: Issues a new Microsegmentation token according to given data.

# Attributes
attributes:
  v1:
  - name: audience
    description: |-
      If given, the issued token will only be valid for the specified namespace.
      Refer to [JSON Web Token (JWT)RFC
      7519](https://tools.ietf.org/html/rfc7519#section-4.1.3).
      for further information.
    type: string
    exposed: true
    example_value: aud:*:*:/namespace
    validations:
    - $audience

  - name: claims
    description: The claims in the token. It is only set is the parameter `asCookie` is given.
    type: external
    exposed: true
    subtype: _claims
    read_only: true
    autogenerated: true
    omit_empty: true

  - name: data
    description: Contains additional data. The value depends on the issuer type.
    type: string
    exposed: true
    deprecated: true
    orderable: true

  - name: metadata
    description: Contains various additional information. Meaning depends on the `realm`.
    type: external
    exposed: true
    subtype: map[string]interface{}
    example_value:
      vinceAccount: acme
      vinceOTP: 665435
      vincePassword: s3cr3t
    orderable: true

  - name: opaque
    description: Opaque data that will be included in the issued token.
    type: external
    exposed: true
    subtype: map[string]string

  - name: quota
    description: Restricts the number of times the issued token can be used.
    type: integer
    exposed: true

  - name: realm
    description: |-
      The authentication realm. This will define how to verify
      credentials from internal or external source of authentication.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - AWSSecurityToken
    - Certificate
    - Google
    - LDAP
    - Vince
    - GCPIdentityToken
    - AzureIdentityToken
    - OIDC
    - SAML
    - AporetoIdentityToken
    - PCIdentityToken
    example_value: Vince

  - name: restrictedNamespace
    description: |-
      Restricts the namespace where the token can be used.

      For instance, if you have have access to `/namespace` and below, you can
      tell the policy engine that it should restrict further more to
      `/namespace/child`.

      Restricting to a namespace you don't have initially access according to the
      policy engine has no effect and may end up making the token unusable.
    type: string
    exposed: true
    example_value: /namespace

  - name: restrictedNetworks
    description: |-
      Restricts the networks from where the token can be used. This will reduce the
      existing set of authorized networks that normally apply to the token according
      to the policy engine.

      For instance, If you have authorized access from `0.0.0.0/0` (by default) or
      from
      `10.0.0.0/8`, you can ask for a token that will only be valid if used from
      `10.1.0.0/16`.

      Restricting to a network that is not initially authorized by the policy
      engine has no effect and may end up making the token unusable.
    type: list
    exposed: true
    subtype: string
    example_value:
    - 10.0.0.0/8
    - 127.0.0.1/32
    validations:
    - $optionalcidrs

  - name: restrictedPermissions
    description: |-
      Restricts the permissions of token. This will reduce the existing permissions
      that normally apply to the token according to the policy engine.

      For instance, if you have administrative role, you can ask for a token that will
      tell the policy engine to reduce the permission it would have granted to what is
      given defined in the token.

      Restricting to some permissions you don't initially have according to the policy
      engine has no effect and may end up making the token unusable.
    type: list
    exposed: true
    subtype: string
    example_value:
    - '@auth:role=enforcer'
    - namespace,post

  - name: token
    description: The token to use for the registration.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: validity
    description: |-
      Configures the maximum length of validity for a token, using
      [Golang duration syntax](https://golang.org/pkg/time/#example_Duration). If it
      is
      bigger than the configured max validity, it will be capped. Default: `24h`.
    type: string
    exposed: true
    default_value: 24h
    orderable: true
    validations:
    - $timeDuration
