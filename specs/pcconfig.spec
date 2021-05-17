# Model
model:
  rest_name: pcconfig
  resource_name: pcconfig
  entity_name: PCConfig
  package: karl
  group: core/tenant
  description: Holds the Prisma Cloud configuration for a namespace.
  aliases:
  - pcc
  get:
    description: Retrieve the Prisma Cloud configuration with the given ID.
  update:
    description: Updates the Prisma Cloud configuration with the given ID.
  delete:
    description: Deletes the Prisma Cloud configuration with the given ID.
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
