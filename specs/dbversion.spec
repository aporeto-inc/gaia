# Model
model:
  rest_name: dbversion
  resource_name: dbversions
  entity_name: DBVersion
  package: squall
  group: infra/dbversions
  description: A DBVersion shows the versions of each database.
  aliases:
  - dbvers

# Attributes
attributes:
  v1:
  - name: name
    description: Name of the service.
    type: string
    exposed: true

  - name: versionNumber
    description: Version of the database.
    type: integer
    exposed: true
