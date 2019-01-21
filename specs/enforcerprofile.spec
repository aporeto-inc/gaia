# Model
model:
  rest_name: enforcerprofile
  resource_name: enforcerprofiles
  entity_name: EnforcerProfile
  package: squall
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
    - zhash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
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
  - '@identifiable-pk-stored'
  - '@named'
  - '@metadatable'
  - '@zonable'
  validations:
  - $enforcerprofile

# Attributes
attributes:
  v1:
  - name: PUBookkeepingInterval
    description: PUBookkeepingInterval configures how often the PU will be synchronized.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 15m
    orderable: true

  - name: PUHeartbeatInterval
    description: PUHeartbeatInterval configures the heart beat interval.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 5s
    orderable: true

  - name: excludedInterfaces
    description: ExcludedInterfaces is a list of interfaces that must be excluded.
    type: external
    exposed: true
    subtype: excluded_interfaces_list
    stored: true
    orderable: true

  - name: excludedNetworks
    description: |-
      ExcludedNetworks is the list of networks that must be excluded for this
      enforcer.
    type: external
    exposed: true
    subtype: excluded_networks_list
    stored: true
    orderable: true

  - name: ignoreExpression
    description: |-
      IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
      docker container started with labels matching the rule.
    type: external
    exposed: true
    subtype: policies_list
    stored: true

  - name: killContainersOnFailure
    description: |-
      KillContainersOnFailure will configure the enforcers to kill any containers if
      there are policy failures.
    type: boolean
    exposed: true
    stored: true

  - name: metadataExtractor
    description: Select which metadata extractor to use to process new processing
      units.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Docker
    - ECS
    - KubernetesPUperPOD
    - KubernetesPUperContainer
    default_value: Docker
    orderable: true

  - name: policySynchronizationInterval
    description: |-
      PolicySynchronizationInterval configures how often the policy will be
      resynchronized.
    type: string
    exposed: true
    stored: true
    allowed_chars: ^[0-9]+[smh]$
    allowed_chars_message: must be a valid duration like <n>s or <n>s or <n>h
    default_value: 10m
    orderable: true

  - name: targetNetworks
    description: TargetNetworks is the list of networks that authorization should
      be applied.
    type: external
    exposed: true
    subtype: target_networks_list
    stored: true
    orderable: true

  - name: targetUDPNetworks
    description: |-
      TargetUDPNetworks is the list of UDP networks that authorization should be
      applied.
    type: external
    exposed: true
    subtype: target_networks_list
    stored: true
    orderable: true

  - name: trustedCAs
    description: List of trusted CA. If empty the main chain of trust will be used.
    type: external
    exposed: true
    subtype: trusted_cas_list
    stored: true
