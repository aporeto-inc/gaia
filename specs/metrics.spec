# Model
model:
  rest_name: metrics
  resource_name: metrics
  entity_name: Metrics
  package: jenova
  group: visualization/metrics
  description: |-
    Prometheus compatible endpoint to evaluate an expression query over a range of
    time. This can be used to retrieve back Aporeto specific metrics for a given
    namespace. All queries are protected within the namespace of the caller.
  aliases:
  - mq

# Attributes
attributes:
  v1:
  - name: end
    description: End timestamp <rfc3339 | unix_timestamp>.
    type: string
    exposed: true
    example_value: "2015-07-01T20:11:00.781Z"

  - name: query
    description: Contains the remote `POST` payload.
    type: string
    exposed: true
    required: true
    example_value: flows{namespace=~"/mycompany.*"}

  - name: start
    description: Start timestamp <rfc3339 | unix_timestamp>.
    type: string
    exposed: true
    example_value: "2015-07-01T20:11:00.781Z"

  - name: step
    description: Query resolution step width in duration format or float number of
      seconds.
    type: string
    exposed: true
    example_value: 15s
