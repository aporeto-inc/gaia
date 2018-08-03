# Model
model:
  rest_name: installedapp
  resource_name: installedapps
  entity_name: InstalledApp
  package: highwind
  description: InstalledApps represents an installed application.
  aliases:
  - iapps
  - iapp
  get:
    description: Retrieves the object with the given ID.
  update:
    description: Updates the object with the given ID.
  delete:
    description: Deletes the object with the given ID.

# Attributes
attributes:
  v1:
  - name: ID
    description: ID of the installed app.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true
    identifier: true
    primary_key: true

  - name: accountName
    description: AccountName represents the vince account name.
    type: string
    exposed: true
    stored: true
    creation_only: true

  - name: categoryID
    description: CategoryID of the app.
    type: string
    exposed: true
    stored: true
    read_only: true
    orderable: true

  - name: currentVersion
    description: Version of the installed app.
    type: string
    exposed: true
    stored: true

  - name: data
    description: Data retains all data created to use this service.
    type: external
    subtype: service_data
    stored: true

  - name: k8sIdentifier
    description: K8SIdentifier retains the identifier for kubernetes.
    type: string
    stored: true

  - name: name
    description: Name of the installed app.
    type: string
    exposed: true
    stored: true
    creation_only: true
    orderable: true

  - name: namespace
    description: Namespace in which the app is running.
    type: string
    exposed: true
    stored: true
    orderable: true

  - name: parameters
    description: Parameters is a list of parameters to start the app.
    type: external
    exposed: true
    subtype: app_parameters
    stored: true

  - name: relatedObjects
    description: RelatedObjects retains all objects created to use this app.
    type: external
    subtype: app_relatedobjects
    stored: true

  - name: status
    description: Status of the app.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Error
    - Pending
    - Running
    default_value: Pending
    orderable: true

# Relations
relations:
- rest_name: log
  get:
    description: Returns the logs for a app.
