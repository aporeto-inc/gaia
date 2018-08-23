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
  - name: IP
    description: Actual IP Address.
    type: string
    exposed: true

  - name: actions
    description: List of actions applied from that IP.
    type: list
    exposed: true
    subtype: string

  - name: city
    description: City of the IP.
    type: string
    exposed: true

  - name: country
    description: Country of the IP.
    type: string
    exposed: true

  - name: destinationPorts
    description: List of ports applied used.
    type: list
    exposed: true
    subtype: string

  - name: hits
    description: Number of hits.
    type: integer
    exposed: true

  - name: hostnames
    description: Eventual list of hostnames associated to the IP.
    type: list
    exposed: true
    subtype: string

  - name: l4Protocol
    description: Protocol used.
    type: string
    exposed: true

  - name: latitude
    description: Latitude of the IP.
    type: float
    exposed: true

  - name: longitude
    description: Longitude of the IP.
    type: float
    exposed: true
