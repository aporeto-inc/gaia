# Model
model:
  rest_name: awsapigateway
  resource_name: awsapigateways
  entity_name: AWSAPIGateway
  package: goldrush
  group: none
  description: managed API decisions for the AWS API Gateway.
  get:
    description: Retrieves the object with the given ID.
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

# Attributes
attributes:
  v1:
  - name: APIID
    description: API ID as defined on AWS for the API that handled this request.
    type: string
    exposed: true

  - name: accountID
    description: the account ID for the gateway managing this request.
    type: string
    exposed: true

  - name: authorized
    description: The policy decision for this API flow.
    type: boolean
    exposed: true
    read_only: true
    orderable: true

  - name: method
    description: API method that handled this request.
    type: string
    exposed: true

  - name: namespaceID
    description: Link to the cluster namespace where the AWS API gateway is defined.
    type: string
    exposed: true

  - name: resource
    description: API resource that handled this request.
    type: string
    exposed: true

  - name: sourceIP
    description: the client ip for this request.
    type: string
    exposed: true

  - name: stage
    description: the stage name as defined on AWS for the API that handled this request.
    type: string
    exposed: true

  - name: token
    description: the JWT token that was optionally attached to this request.
    type: string
    exposed: true
