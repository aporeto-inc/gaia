# Model
model:
  rest_name: graphpolicyinfo
  resource_name: graphpolicyinfos
  entity_name: GraphPolicyInfo
  package: jenova
  group: visualization/depmaps
  description: Represents policy information.
  detached: true

# Attributes
attributes:
  v1:
  - name: count
    description: Number of times the policy has been hit.
    type: integer
    exposed: true

  - name: namespace
    description: Namespace of the policy.
    type: string
    exposed: true
