# Model
model:
  rest_name: cachedflowrecord
  resource_name: cachedflowrecords
  entity_name: CachedFlowRecord
  package: zack
  group: policy/networking
  description: Post a new cached flow record.
  extends:
  - '@flowedge'

# Attributes
attributes:
  v1:
  - name: destinationController
    description: Identifier of the destination controller.
    type: string
    exposed: true
    example_value: api.east.acme.com

  - name: destinationID
    description: ID of the destination. May be a processing unit ID or an enforcer-local ID.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: destinationIP
    description: Destination IP address.
    type: string
    exposed: true

  - name: destinationIsTemporary
    description: Indicates if the destination endpoint is an enforcer-local processing unit.
    type: boolean
    exposed: true

  - name: destinationPort
    description: Port of the destination.
    type: integer
    exposed: true

  - name: sourceController
    description: Identifier of the source controller.
    type: string
    exposed: true
    example_value: api.west.acme.com

  - name: sourceID
    description: ID of the source. May be a processing unit ID or an enforcer-local ID.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true

  - name: sourceIsTemporary
    description: Indicates if the source endpoint is an enforcer-local processing unit.
    type: boolean
    exposed: true
