attributes:
- description: edges are the edges of the map
  exposed: true
  name: edges
  read_only: true
  required: true
  subtype: graphedges_map
  type: external
- description: Groups provide information about the group values
  exposed: true
  name: groups
  read_only: true
  required: true
  subtype: graphgroups_map
  type: external
- description: nodes refers to the nodes of the map
  exposed: true
  name: nodes
  read_only: true
  required: true
  subtype: graphnodes_map
  type: external
- description: viewSuggestions provides suggestion of views based on relevant tags.
  exposed: true
  name: viewSuggestions
  read_only: true
  required: true
  subtype: view_suggestions
  type: external
model:
  aliases:
  - depmaps
  - depmap
  description: This api returns a data structure representing the graph of all processing
    units and their connections in a particular namespace, in a given time window.
    To pass the time window you can use the query parameters "startAbsolute", "endAbsolute",
    "startRelative", "endRelative".  For example "https://squall.aporeto.com/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000"
  entity_name: DependencyMap
  extends:
  - '@identifiable-nopk-nostored'
  package: jenova
  resource_name: dependencymaps
  rest_name: dependencymap
