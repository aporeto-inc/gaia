# Model
model:
  rest_name: awsregister
  resource_name: awsregister
  entity_name: AWSRegister
  package: bill
  group: none
  description: This endpoint allows AWS customers to register with Palo Alto Networks-hosted Segment Consoles for billing.
  extends:
  - '@identifiable-stored'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: provider
    description: Token Provided by AWS.
    type: string
    exposed: true
