# Model
model:
  rest_name: externalaccess
  resource_name: externalaccesses
  entity_name: ExternalAccess
  package: jenova
  description: ExternalAccess allows to retrieve connection from or to an external
    service.
  aliases:
  - extacs
  - extac

# Attributes
attributes:
  v1:
  - name: IPRecords
    description: IPRecords refers to a list of IPRecord that contains the IP information.
    type: external
    exposed: true
    subtype: ip_records
    read_only: true
    autogenerated: true
