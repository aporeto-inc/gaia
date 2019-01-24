# Model
model:
  rest_name: app
  resource_name: apps
  entity_name: App
  package: highwind
  description: App represents an application that can be installed.
  extends:
  - '@described'
  - '@named'
  - '@zonable'

# Attributes
attributes:
  v1:
  - name: beta
    description: Beta indicates if the app is in a beta version.
    type: boolean
    exposed: true
    read_only: true

  - name: categoryID
    description: CategoryID of the app.
    type: string
    exposed: true
    read_only: true

  - name: icon
    description: Icon contains a base64 image for the app.
    type: string
    exposed: true
    read_only: true

  - name: latestVersion
    description: LatestVersion represents the latest version available of the app.
    type: string
    exposed: true

  - name: longDescription
    description: LongDescription contains a more detailed description of the app.
    type: string
    exposed: true

  - name: parameters
    description: Parameters is a list of parameters available for the app.
    type: external
    exposed: true
    subtype: _app_parameters
    stored: true

  - name: title
    description: Title represents the title of the app.
    type: string
    exposed: true
