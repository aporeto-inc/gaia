# Model
model:
  rest_name: dnsreport
  resource_name: dnsreports
  entity_name: DNSReport
  package: zack
  group: policy/dns
  description: |-
    A DNSReport is used to report a DNS lookup that is happening on
    behalf of a processing unit. If the DNS server is on the standard udp port 53
    then enforcer is able to proxy the DNS traffic and make a report. The report
    indicate whether or not the lookup was successful.

# Attributes
attributes:
  v1:
  - name: action
    description: Action of the DNS request.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - Accept
    - Reject
    example_value: Accept

  - name: enforcerID
    description: ID of the enforcer.
    type: string
    exposed: true

  - name: enforcerNamespace
    description: Namespace of the enforcer.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: nameLookup
    description: name looked up by PU.
    type: string
    exposed: true
    required: true
    example_value: www.google.com

  - name: processingUnitID
    description: ID of the PU.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx

  - name: processingUnitNamespace
    description: Namespace of the PU.
    type: string
    exposed: true
    required: true
    example_value: /my/namespace

  - name: reason
    description: |-
      This field is only set when the lookup fails. It specifies the reason for the
      failure.
    type: string
    exposed: true

  - name: sourceIP
    description: Type of the source.
    type: string
    exposed: true
    required: true
    example_value: 10.0.0.1

  - name: timestamp
    description: Time and date of the log.
    type: time
    exposed: true

  - name: value
    description: Number of times the client saw this activity.
    type: integer
    exposed: true
    required: true
    example_value: 1
