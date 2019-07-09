# Model
model:
  rest_name: app
  resource_name: apps
  entity_name: App
  package: highwind
  group: integration/app
  description: Represents an application that can be installed.
  extends:
  - '@described'
  - '@named'

# Indexes
indexes:
- - :no-inherit

# Attributes
attributes:
  v1:
  - name: beta
    description: Set to `true` to indicate that the app is in a beta version.
    type: boolean
    exposed: true
    read_only: true

  - name: categoryID
    description: Category ID of the app.
    type: string
    exposed: true
    read_only: true

  - name: icon
    description: Contains a base64-encoded image for the app.
    type: string
    exposed: true
    read_only: true

  - name: latestVersion
    description: Represents the latest version available of the app.
    type: string
    exposed: true

  - name: longDescription
    description: Contains a more detailed description of the app.
    type: string
    exposed: true

  - name: steps
    description: List of steps that contain parameters.
    type: refList
    exposed: true
    subtype: uistep
    stored: true
    extensions:
      refMode: pointer

  - name: title
    description: Represents the title of the app.
    type: string
    exposed: true
