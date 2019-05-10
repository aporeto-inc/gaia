# Model
model:
  rest_name: installedapp
  resource_name: installedapps
  entity_name: InstalledApp
  package: highwind
  group: integration/app
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
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@identifiable-stored'
  - '@named'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: appIdentifier
    description: AppIdentifier retains the identifier for the app.
    type: string
    stored: true

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

  - name: deploymentCount
    description: DeploymentCount represents the number of expected deployment for
      this app.
    type: integer
    stored: true
    read_only: true

  - name: parameters
    description: Parameters contains the computed parameters to start the app.
    type: external
    exposed: true
    subtype: map[string]interface{}
    stored: true

  - name: status
    description: Status of the app.
    type: enum
    exposed: true
    stored: true
    read_only: true
    allowed_choices:
    - Unknown
    - Deploying
    - Initializing
    - Running
    - Undeploying
    - Error
    default_value: Unknown
    orderable: true

  - name: statusMessage
    description: Reason for the status of the app.
    type: string
    exposed: true
    stored: true
    read_only: true

# Relations
relations:
- rest_name: log
  get:
    description: Returns the logs for a app.
