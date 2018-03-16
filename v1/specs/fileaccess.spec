attributes:
- autogenerated: true
  description: Action tells if the access has been allowed or not.
  exposed: true
  format: free
  name: action
  read_only: true
  type: string
- autogenerated: true
  description: Count tells how many times the file has been accessed.
  exposed: true
  name: count
  read_only: true
  type: integer
- autogenerated: true
  description: Host is the host that served the accessed file.
  exposed: true
  format: free
  name: host
  read_only: true
  type: string
- allowed_choices:
  - Read
  - ReadWrite
  - Write
  autogenerated: true
  description: Mode is the mode of the accessed file.
  exposed: true
  name: mode
  read_only: true
  type: enum
- autogenerated: true
  description: Path is the path of the accessed file.
  exposed: true
  format: free
  name: path
  read_only: true
  type: string
- autogenerated: true
  description: Protocol is the protocol used to access the file.
  exposed: true
  format: free
  name: protocol
  read_only: true
  type: string
model:
  description: Returns file access statistics on a particular processing unit.
  entity_name: FileAccess
  package: jenova
  resource_name: fileaccesses
  rest_name: fileaccess
