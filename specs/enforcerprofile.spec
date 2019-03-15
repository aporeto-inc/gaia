# Model
model:
  rest_name: enforcerprofile
  resource_name: enforcerprofiles
  entity_name: EnforcerProfile
  package: squall
  group: policy/enforcerconfig
  description: |-
    Allows to create reusable configuration profile for your enforcers. Enforcer
    Profiles contains various startup information that can (for some) be updated
    live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
    Mapping Policy.
  aliases:
  - profile
  - profiles
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $propagatable
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
  - '@named'
  - '@metadatable'
  - '@zonable'
  - '@propagated'
  validations:
  - $enforcerprofile

# Attributes
attributes:
  v1:
  - name: excludedInterfaces
    description: ExcludedInterfaces is a list of interfaces that must be excluded.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: excludedNetworks
    description: |-
      ExcludedNetworks is the list of networks that must be excluded for this
      enforcer.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: ignoreExpression
    description: |-
      IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
      docker container started with labels matching the rule.
    type: external
    exposed: true
    subtype: '[][]string'
    stored: true

  - name: targetNetworks
    description: TargetNetworks is the list of networks that authorization should
      be applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: targetUDPNetworks
    description: |-
      TargetUDPNetworks is the list of UDP networks that authorization should be
      applied.
    type: list
    exposed: true
    subtype: string
    stored: true
    orderable: true

  - name: trustedCAs
    description: List of trusted CA. If empty the main chain of trust will be used.
    type: list
    exposed: true
    subtype: string
    stored: true
