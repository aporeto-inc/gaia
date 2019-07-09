# Model
model:
  rest_name: fileaccesspolicy
  resource_name: fileaccesspolicies
  entity_name: FileAccessPolicy
  package: squall
  group: policy/files
  description: |-
    A file access policy allows processing units to access various folder and files.
    It will use the tags of a file path to know what is the path of the file or
    folder to allow access to. You can allow the processing unit to have any
    combination of read, write, or execute.

    When a processing unit is a Docker container, then it will police the volumes.
    Mount and execute won't have any effect.

    File paths are not supported yet for standard Linux processes.
  get:
    description: Retrieves the policy with the given ID.
  update:
    description: Updates the policy with the given ID.
  delete:
    description: Deletes the policy with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@namespaced'
  - '@described'
  - '@disabled'
  - '@identifiable-not-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'
  - '@schedulable'
  - '@timeable'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: allowsExecute
    description: Allows files to be executed.
    type: boolean
    exposed: true
    orderable: true

  - name: allowsRead
    description: Allows files to be read.
    type: boolean
    exposed: true
    orderable: true

  - name: allowsWrite
    description: Allows files to be written.
    type: boolean
    exposed: true
    orderable: true

  - name: encryptionEnabled
    description: Set to `true` to enable automatic encryption.
    type: boolean
    exposed: true
    orderable: true

  - name: expirationTime
    description: If set the policy will be automatically deleted after the given time.
    type: time
    exposed: true
    stored: true
    getter: true
    setter: true

  - name: logsEnabled
    description: A value of `true` enables logging.
    type: boolean
    exposed: true
    orderable: true

  - name: object
    description: The object of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

  - name: subject
    description: The subject of the policy.
    type: external
    exposed: true
    subtype: '[][]string'
    orderable: true
    validations:
    - $tagsExpression

# Relations
relations:
- rest_name: filepath
  get:
    description: Returns the list of file paths that match the policy.

- rest_name: processingunit
  get:
    description: Returns the list of processing units that match the policy.
