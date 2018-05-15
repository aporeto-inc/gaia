# Model
model:
  rest_name: policyrefresh
  resource_name: policyrefreshs
  entity_name: PolicyRefresh
  package: squall
  description: |-
    PolicyRefresh is sent to client when as a push event when a policy refresh is
    needed on their side.

# Attributes
attributes:
  v1:
  - name: sourceNamespace
    description: SourceNamespace contains the original namespace of the updated object.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: type
    description: Type contains the policy type that is affected.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true
