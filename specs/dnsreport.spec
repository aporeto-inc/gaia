# Model
model:
  rest_name: dnsreport
  resource_name: dnsreports
  entity_name: DNSReport
  package: zack
  group: policy/networking
  description: |-
     `dnsreport` is used to report all the dns lookups that are happening on
     behalf of the PU. If the DNS server is on the standard udp port 53 then
     enforcer is able to proxy the dns traffic and report the dns lookups. It
     also reports whether the lookup was successful or not.

# Attributes
attributes:
  v1:
  - name: processingUnitNamespace
    description: Namespace of the PU.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: namespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace
   
  - name: processingUnitID
    description: ID of the PU.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true

  - name: sourceIP
    description: Type of the source.
    type: string
    required: true
    exposed: true
    example_value: 10.0.0.1

  - name: nameLookup
    description: name looked up by PU.
    type: string
    exposed: true
    required: true
    example_value: www.google.com

  - name: success
    description: Result reports whether dns request succeeded or failed.
    type: boolean
    exposed: true
    required: true
    example_value: true

  - name: error
    description: If the result is false, error reports the reason of the dns failure.
    type: string
    exposed: true

  - name: value
    description: Number of times the client saw this activity.
    type: integer
    exposed: true
    required: true
    example_value: 1

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true