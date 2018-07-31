# Model
model:
  rest_name: awsapigateway
  resource_name: awsapigateways
  entity_name: AWSAPIGateway
  package: goldrush
  description: managed API decisions for the AWS API Gateway.
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
  - name: APIID
    description: API ID as defined on AWS for the API that handled this request.
    type: string
    exposed: true
    format: free

  - name: accountID
    description: the account ID for the gateway managing this request.
    type: string
    exposed: true
    format: free

  - name: authorized
    description: The policy decision for this API flow.
    type: boolean
    exposed: true
    read_only: true
    format: free
    orderable: true

  - name: method
    description: API method that handled this request.
    type: string
    exposed: true
    format: free

  - name: namespaceID
    description: Link to the cluster namespace where the AWS API gateway is defined.
    type: string
    exposed: true
    format: free

  - name: resource
    description: API resource that handled this request.
    type: string
    exposed: true
    format: free

  - name: sourceIP
    description: the client ip for this request.
    type: string
    exposed: true
    format: free

  - name: stage
    description: the stage name as defined on AWS for the API that handled this request.
    type: string
    exposed: true
    format: free

  - name: token
    description: the JWT token that was optionally attached to this request.
    type: string
    exposed: true
    format: free
