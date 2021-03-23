# Model
model:
  rest_name: cloudnetworkquery
  resource_name: cloudnetworkqueries
  entity_name: CloudNetworkQuery
  package: yeul
  group: pcn/infrastructure
  description: Provides the parameters for an effective network permissions query.
  get:
    description: Retrieves the cloud query with the given ID.
    global_parameters:
    - $filtering
  update:
    description: Updates the cloud query with the given ID.
  delete:
    description: Deletes the the cloud query with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@zoned'
  - '@migratable'
  - '@base'
  - '@namespaced'
  - '@described'
  - '@identifiable-stored'
  - '@named'

# Attributes
attributes:
  v1:
  - name: destinationIP
    description: The destination IP of a trace route request. Might not always be
      an endpoint.
    type: string
    exposed: true
    stored: true

  - name: destinationPorts
    description: The destination port or ports that should be used for the trace route
      command.
    type: external
    exposed: true
    subtype: _portlist
    stored: true

  - name: destinationProtocol
    description: The destination protocol that should be used for the trace route
      commands.
    type: integer
    exposed: true
    stored: true
    default_value: -1
    max_value: 255
    min_value: -1

  - name: destinationSelector
    description: A filter for selecting destinations for the query.
    type: ref
    exposed: true
    subtype: cloudnetworkqueryfilter
    stored: true
    extensions:
      refMode: pointer

  - name: excludeEnterpriseIPs
    description: |-
      If set, the evaluation will exclude enterprise IPs from the effective
      permissions.
    type: boolean
    exposed: true
    stored: true

  - name: rawRQL
    description: The RQL string for this query as a reference.
    type: string
    stored: true

  - name: sourceIP
    description: The source IP of a trace route request. Might not be always and endpoint.
    type: string
    exposed: true
    stored: true

  - name: sourceSelector
    description: A filter for selecting the sources of the request.
    type: ref
    exposed: true
    subtype: cloudnetworkqueryfilter
    stored: true
    extensions:
      refMode: pointer

  - name: type
    description: Indicates the type of results that should be provided by the query.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Summary
    - CompressedGraph
    - FullGraph
    default_value: Summary

# Relations
relations:
- rest_name: cloudgraph
  get:
    description: Initiates a calculation of the query and retrieves the results in
      CloudGraph.

- rest_name: cloudpolicy
  get:
    description: Retrieves the policies associated with this query.
  create:
    description: Creates a policy associated with this query.
