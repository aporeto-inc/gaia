# Model
model:
  rest_name: filepath
  resource_name: filepaths
  entity_name: FilePath
  package: squall
  group: policy/files
  description: |-
    A File Path represents a random path to a file or a folder. They can be used in
    aFile Access Policiesin order to allow Processing Units to access them, using
    various modes (read, write, execute). You will need to use the File Paths tags
    to set some policies. A good example would bevolume=web or file=/etc/passwd.
  aliases:
  - fp
  - fps
  indexes:
  - - :shard
    - zone
    - zHash
  - - namespace
  - - namespace
    - name
  - - namespace
    - normalizedTags
  - - archived
  get:
    description: Retrieves the object with the given ID.
    global_parameters:
    - $archivable
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
  - '@metadatable'
  - '@named'
  - '@zonable'
  - '@archivable'

# Attributes
attributes:
  v1:
  - name: filepath
    description: FilePath refer to the file mount path.
    type: string
    exposed: true
    stored: true
    required: true
    example_value: /etc/passwd

  - name: server
    description: server is the server name/ID/IP associated with the file path.
    type: string
    exposed: true
    stored: true
    creation_only: true
