# Model
model:
  rest_name: reportsqueryresults
  resource_name: reportsqueryresults
  entity_name: ReportsQueryResults
  package: jenova
  group: visualization/reportsquery
  description: Represent the results of a reports query.
  detached: true

# Attributes
attributes:
  v1:
  - name: DNSLookupReports
    description: List of DNSLookupReports.
    type: refList
    exposed: true
    subtype: dnslookupreport
    omit_empty: true

  - name: counterReports
    description: List of CounterReports.
    type: refList
    exposed: true
    subtype: counterreport
    omit_empty: true

  - name: enforcerReports
    description: List of EnforcerReports.
    type: refList
    exposed: true
    subtype: enforcerreport
    omit_empty: true

  - name: eventLogs
    description: List of EventLogs.
    type: refList
    exposed: true
    subtype: eventlog
    omit_empty: true

  - name: flowReports
    description: List of FlowReports.
    type: refList
    exposed: true
    subtype: flowreport
    omit_empty: true

  - name: packetReports
    description: List of PacketReports.
    type: refList
    exposed: true
    subtype: packetreport
    omit_empty: true
