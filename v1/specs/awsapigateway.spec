# Model
model:
  rest_name: awsapigateway
  resource_name: awsapigateways
  entity_name: AWSApiGateway
  package: goldrush
  description: managed API decisions for the AWS API Gateway.
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'

# Attributes
attributes:
  v1:
  - name: accountID
    description: the accounf ID for the gateway managing this request.
    type: string
    format: free

  - name: apiID
    description: API ID as defined on AWS for the API that handled this request.
    type: string
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
    format: free

  - name: namespaceID
    description: Link to the cluster namespace where the AWS API gateway is defined.
    type: string
    format: free

  - name: ressource
    description: API ressource that handled this request.
    type: string
    format: free

  - name: ressourceID
    description: API ressource ID that handled this request.
    type: string
    format: free

  - name: sourceIP
    description: the client ip for this request.
    type: string
    format: free

  - name: stage
    description: the stage name as defined on AWS for the API that handled this request.
    type: string
    format: free
