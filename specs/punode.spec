# Model
model:
  rest_name: punode
  resource_name: punodes
  entity_name: PUNode
  package: squall
  group: visualization/depmaps
  description: Internal scaling down of a pu for dep map representation.
  private: true
  extends:
  - '@identifiable-not-stored'

# Attributes
attributes:
  v1:
  - name: enforcementStatus
    description: Enforcement status of the pu.
    type: string
    exposed: true

  - name: images
    description: Images of the pu.
    type: list
    exposed: true
    subtype: string

  - name: lastPokeTime
    description: Last poke time of the pu.
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
    type: list
    exposed: true
    subtype: string

  - name: type
    description: Type of the pu.
    type: string
    exposed: true

  - name: unreachable
    description: If true, the pu is unreachable.
    type: boolean
    exposed: true
