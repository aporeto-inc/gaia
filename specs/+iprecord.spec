# Model
model:
  rest_name: iprecord
  resource_name: iprecords
  entity_name: IPRecord
  package: jenova
  description: Represents an IP access.
  detached: true

# Attributes
attributes:
  v1:
  - name: hits
    description: Number of time the port was accessed.
    type: integer
    exposed: true

  - name: protocol
    description: Protocol used.
    type: integer
    exposed: true

  - name: timestamp
    description: Date of the last access on that port.
    type: time
    exposed: true
