# Model
model:
  rest_name: debugbundle
  resource_name: debugbundles
  entity_name: DebugBundle
  package: guy
  group: debug
  description: Represents a file that can be uploaded.
  extends:
  - '@identifiable-not-stored'
  - '@namespaced'

# Attributes
attributes:
  v1:
  - name: enforcerID
    description: The ID of the enforcer.
    type: string
    exposed: true
    read_only: true
