# Model
model:
  rest_name: ldapprovider
  resource_name: ldapproviders
  entity_name: LDAPProvider
  package: cactuar
  group: core/authentication
  description: |-
    Allows to declare a generic LDAP provider that can be used in exchange
    for a Midgard token.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@identifiable-stored'
  - '@timeable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: address
    description: Address holds the account authentication account's private LDAP
      server.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: ldap.company.com
    filterable: true
    orderable: true

  - name: baseDN
    description: BaseDN holds the base DN to use to LDAP queries.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: dc=universe,dc=io
    filterable: true
    orderable: true

  - name: bindDN
    description: BindDN holds the account's internal LDAP bind DN for querying
      auth.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: cn=readonly,dc=universe,dc=io
    filterable: true
    orderable: true

  - name: bindPassword
    description: BindPassword holds the password to the LDAP bind DN.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: s3cr3t
    orderable: true

  - name: bindSearchFilter
    description: |-
      BindSearchFilter holds filter to be used to uniquely search a user. For
      Windows based systems, value may be `sAMAccountName={USERNAME}`. For Linux and
      other systems, value may be `uid={USERNAME}`.
    type: string
    exposed: true
    stored: true
    default_value: uid={USERNAME}
    orderable: true

  - name: certificateAuthority
    description: |-
      CertificateAuthority contains the optional certificate authority that will
      be used to connect to the LDAP server. It is not needed if the TLS certificate
      of the LDAP is issued from a public truster CA.
    type: string
    exposed: true
    stored: true
    example_value: |-
      -----BEGIN CERTIFICATE-----
      MIIBPzCB5qADAgECAhEAwbx3c+QW24ePXyD94geytzAKBggqhkjOPQQDAjAPMQ0w
      CwYDVQQDEwR0b3RvMB4XDTE5MDIyMjIzNDA1MFoXDTI4MTIzMTIzNDA1MFowDzEN
      MAsGA1UEAxMEdG90bzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJi6CwRDeKks
      Xb3pDEslmFGR7k9Aeh5RK+XmdqKKPGb3NQWEFPGolnqOR34iVuf7KSxTuzaaVWfu
      XEa94faUQEqjIzAhMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMBAf8EBTADAQH/MAoG
      CCqGSM49BAMCA0gAMEUCIQD+nL9RF9EvQXHyYuJ31Lz9yWd9hsK91stnpAs890gS
      /AIgQIKjBBpiyQNZZWso5H04qke9QYMVPegiQQufFFBj32c=
      -----END CERTIFICATE-----
    orderable: true

  - name: connSecurityProtocol
    description: ConnSecurityProtocol holds the connection type for the LDAP provider.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - TLS
    - InbandTLS
    default_value: InbandTLS
    filterable: true
    orderable: true

  - name: default
    description: |-
      If set, this will be the default LDAP provider. There can be only one default
      provider in your account. When logging in with LDAP, if no provider name is
      given, the default will be used.
    type: boolean
    exposed: true
    stored: true

  - name: ignoredKeys
    description: |-
      IgnoredKeys holds a list of keys that must not be imported into Aporeto
      authorization system.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: subjectKey
    description: |-
      SubjectKey holds key to be used to populate the subject. If you want to
      use the user as a subject, for Windows based systems you may use
      'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
      also use any alternate key.
    type: string
    exposed: true
    stored: true
    default_value: uid
    orderable: true
