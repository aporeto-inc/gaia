# Model
model:
  rest_name: installation
  resource_name: installations
  entity_name: Installation
  package: highwind
  description: Installation represents an installation for a given account.
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering

# Attributes
attributes:
  v1:
  - name: ID
    description: ID is the identifier of the object.
    type: string
    exposed: true
    stored: true
    identifier: true
    orderable: true
    primary_key: true

  - name: accountName
    description: AccountName that should be installed.
    type: string
    exposed: true
    stored: true
    orderable: true
