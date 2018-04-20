# Model
model:
  rest_name: dependencymap
  resource_name: dependencymaps
  entity_name: DependencyMap
  package: jenova
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
  extends:
  - '@identifiable-nopk-nostored'

# Attributes
attributes:
  v1:
  - name: claims
    description: claims represents a user or a script that have accessed an api.
    type: external
    exposed: true
    subtype: graphclaims_map
    read_only: true

  - name: edges
    description: edges are the edges of the map.
    type: external
    exposed: true
    subtype: graphedges_map
    read_only: true

  - name: groups
    description: Groups provide information about the group values.
    type: external
    exposed: true
    subtype: graphgroups_map
    read_only: true

  - name: nodes
    description: nodes refers to the nodes of the map.
    type: external
    exposed: true
    subtype: graphnodes_map
    read_only: true

  - name: viewSuggestions
    description: viewSuggestions provides suggestion of views based on relevant tags.
    type: external
    exposed: true
    subtype: view_suggestions
    read_only: true
