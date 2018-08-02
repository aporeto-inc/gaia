# Model
model:
  rest_name: filepath
  resource_name: filepaths
  entity_name: FilePath
  package: squall
  description: |-
    A File Path represents a random path to a file or a folder. They can be used in
    aFile Access Policiesin order to allow Processing Units to access them, using
    various modes (read, write, execute). You will need to use the File Paths tags
    to set some policies. A good example would bevolume=web or file=/etc/passwd.
  aliases:
  - fp
  - fps
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

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
    filterable: true

  - name: server
    description: server is the server name/ID/IP associated with the file path.
    type: string
    exposed: true
    stored: true
    creation_only: true
    filterable: true
