# Model
model:
  rest_name: dependencymap
  resource_name: dependencymaps
  entity_name: DependencyMap
  package: jenova
  group: visualization/depmaps
  description: |-
    Returns a data structure representing the graph of all processing units
    and their connections in a particular namespace, in a given time window. To pass
    the time window you can use the query parameters `startAbsolute`, `endAbsolute`,
    `startRelative`, `endRelative`.

    For example:

    `/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000`.
  aliases:
  - depmaps
  - depmap

# Attributes
attributes:
  v1:
  - name: edges
    description: The edges of the map.
    type: refMap
    exposed: true
    subtype: graphedge
    read_only: true
    extensions:
      refMode: pointer

  - name: groups
    description: Provides information about the group values.
    type: refMap
    exposed: true
    subtype: graphgroup
    read_only: true
    extensions:
      refMode: pointer

  - name: nodes
    description: Refers to the nodes of the map.
    type: refMap
    exposed: true
    subtype: graphnode
    read_only: true
    extensions:
      refMode: pointer

  - name: viewSuggestions
    description: Provides suggested views based on relevant tags.
    type: list
    exposed: true
    subtype: string
    read_only: true
