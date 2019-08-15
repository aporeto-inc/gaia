# Model
model:
  rest_name: dnsreport
  resource_name: dnsreports
  entity_name: DNSReport
  package: zack
  group: policy/networking
  description: Post a new dns request report.

# Attributes
attributes:
  v1:
  - name: sourceNamespace
    description: Namespace of the source.
    type: string
    exposed: true
    example_value: /my/namespace

  - name: sourceID
    description: ID of the source.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true

  - name: nameLookup
    description: name looked up by PU.
    type: string
    exposed: true
    required: true
    example_value: www.google.com

  - name: result
    description: Result reports whether dns request succeeded or failed.
    type: boolean
    exposed: true
    required: true
    example_value: true

  - name: error
    description: If the result is false, error reports the reason of the dns failure.
    type: string
    exposed: true

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true
