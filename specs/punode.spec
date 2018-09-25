# Model
model:
  rest_name: punode
  resource_name: punodes
  entity_name: PUNode
  package: squall
  description: Internal scaling down of a pu for dep map representation.
  private: true

# Attributes
attributes:
  v1:
  - name: ID
    description: Identifier of the pu.
    type: string
    exposed: true

  - name: enforcementStatus
    description: Enforcement status of the pu.
    type: string
    exposed: true

  - name: image
    description: Image of the pu.
    type: string
    exposed: true

  - name: lastUpdate
    description: Last update of the pu.
    type: time
    exposed: true

  - name: name
    description: Name of the pu.
    type: string
    exposed: true

  - name: namespace
    description: Namespace of the pu.
    type: string
    exposed: true

  - name: status
    description: Status of the pu.
    type: string
    exposed: true

  - name: tags
    description: Tags of the pu.
    type: external
    exposed: true
    subtype: tags_map

  - name: type
    description: Type of the pu.
    type: string
    exposed: true
