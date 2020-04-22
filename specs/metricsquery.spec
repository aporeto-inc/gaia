# Model
model:
  rest_name: metricsquery
  resource_name: metricsqueries
  entity_name: MetricsQuery
  package: jenova
  group: visualization/metricsquery
  description: |-
    Prometheus compatible endpoint to evaluate an expression query over a range of
    time. This can be used to retrieve back Aporeto specific metrics for a given
    namespace. All queries are protected within the namespace of the caller.
  aliases:
  - mq
