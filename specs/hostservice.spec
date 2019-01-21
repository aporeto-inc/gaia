# Model
model:
  rest_name: hostservice
  resource_name: hostservices
  entity_name: HostService
  package: squall
  description: Represents a set of services that a host must expose and protect.
  aliases:
  - hostsrv
  - hostsrvs
  indexes:
  - - :shard
    - zone
    - zhash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
  - - namespace
    - archived
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@archivable'
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: networkonly
    description: networkonly indicate the host service is of type network only.
    type: boolean
    exposed: true
    stored: true
    read_only: true
    default_value: true
    deprecated: true

  - name: services
    description: Services lists all protocols and ports a service is running.
    type: refList
    exposed: true
    subtype: processingunitservice
    stored: true
    orderable: true
    extensions:
      refMode: pointer
