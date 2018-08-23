# Model
model:
  rest_name: plan
  resource_name: plans
  entity_name: Plan
  package: vince
  description: Plan contains the various billing plans available.
  get:
    description: Retrieves the object with the given ID.

# Attributes
attributes:
  v1:
  - name: description
    description: Description contains the description of the Plan.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: enforcersQuota
    description: EnforcerQuota contains the maximum number of enforcers available
      in the Plan.
    type: integer
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: key
    description: Key contains the key identifier of the Plan.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: name
    description: Name contains the name of the Plan.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: policiesQuota
    description: PoliciesQuota contains the maximum number of policies available in
      the Plan.
    type: integer
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: processingUnitsQuota
    description: ProcessingUnitsQuota contains the maximum PUs available in the Plan.
    type: integer
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
