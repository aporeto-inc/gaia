# Model
model:
  rest_name: cnsconfig
  resource_name: cnsconfigs
  entity_name: CNSConfig
  package: karl
  group: core/tenant
  description: Holds the CNS configuration for a namespace.
  aliases:
  - pcc
  get:
    description: Retrieve the CNS configuration with the given ID.
  update:
    description: Updates the CNS configuration with the given ID.
  delete:
    description: Deletes the CNS configuration with the given ID.
  extends:
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: enableNetEffectivePermissions
    description: If `true` net effective permissions feature is enabled.
    type: boolean
    exposed: true
    stored: true

  - name: enableNetworkSecurity
    description: If `true` network security feature is enabled.
    type: boolean
    exposed: true
    stored: true

  - name: key
    description: The unique key of the configuration.
    type: string
    exposed: true
    stored: true
    read_only: true
