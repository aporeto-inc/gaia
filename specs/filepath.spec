# Model
model:
  rest_name: filepath
  resource_name: filepaths
  entity_name: FilePath
  package: squall
  group: policy/files
  description: |-
    A file path represents a random path to a file or a folder. They can be used in
    file access policies to allow processing units to access them, using
    various modes (read, write, execute). You will need to use the file paths tags
    to set some policies. A good example would be `volume=web` or `file=/etc/passwd`.
  aliases:
  - fp
  - fps
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
  - '@zoned'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@metadatable'
  - '@named'
  - '@archivable'
  - '@propagated'
  - '@timeable'

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
