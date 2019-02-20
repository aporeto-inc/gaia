# Model
model:
  rest_name: dependencymap
  resource_name: dependencymaps
  entity_name: DependencyMap
  package: jenova
  group: visualization/depmaps
  description: |-
    This api returns a data structure representing the graph of all processing units
    and their connections in a particular namespace, in a given time window. To pass
    the time window you can use the query parameters 'startAbsolute', 'endAbsolute',
    'startRelative', 'endRelative'.

    For example
      "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000".
  aliases:
  - depmaps
  - depmap

# Attributes
attributes:
  v1:
  - name: claims
    description: claims represents a user or a script that have accessed an api.
    type: external
    exposed: true
    subtype: map[string][]string
    read_only: true

  - name: edges
    description: edges are the edges of the map.
    type: refMap
    exposed: true
    subtype: graphedge
    read_only: true
    extensions:
      refMode: pointer

  - name: groups
    description: Groups provide information about the group values.
    type: refMap
    exposed: true
    subtype: graphgroup
    read_only: true
    extensions:
      refMode: pointer

  - name: nodes
    description: nodes refers to the nodes of the map.
    type: refMap
    exposed: true
    subtype: graphnode
    read_only: true
    extensions:
      refMode: pointer

  - name: viewSuggestions
    description: viewSuggestions provides suggestion of views based on relevant tags.
    type: list
    exposed: true
    subtype: string
    read_only: true
