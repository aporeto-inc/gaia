# Model
model:
  rest_name: prismacloudconfiguration
  resource_name: prismacloudconfiguration
  entity_name: PrismaCloudConfiguration
  package: karl
  group: core/tenant
  description: Holds the various Prisma Cloud configuration for a namespace.
  aliases:
  - pcc
  update:
    description: Updates the Prisma Cloud configuration with the given namespace ID.
  delete:
    description: Deletes the Prisma Cloud configuration with the given namespace ID.
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
