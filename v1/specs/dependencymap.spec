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
      "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000"
  aliases:
  - depmaps
  - depmap
  extends:
  - '@identifiable-nopk-nostored'

# Attributes
attributes:
- name: edges
  description: edges are the edges of the map
  type: external
  exposed: true
  subtype: graphedges_map
  required: true
  read_only: true

- name: groups
  description: Groups provide information about the group values
  type: external
  exposed: true
  subtype: graphgroups_map
  required: true
  read_only: true

- name: nodes
  description: nodes refers to the nodes of the map
  type: external
  exposed: true
  subtype: graphnodes_map
  required: true
  read_only: true

- name: viewSuggestions
  description: viewSuggestions provides suggestion of views based on relevant tags.
  type: external
  exposed: true
  subtype: view_suggestions
  required: true
  read_only: true
