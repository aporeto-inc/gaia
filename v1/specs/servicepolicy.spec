# Model
model:
  rest_name: servicepolicy
  resource_name: servicepolicies
  entity_name: ServicePolicy
  package: squall
  description: |-
    A ServicePolicy allows to declare what Services one or a group of processing
    units
    are exposed or consuming.
  aliases:
  - srvpol
  get: true
  update: true
  delete: true
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-nopk-nostored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@fallback'

# Attributes
attributes:
  v1:
  - name: FQDN
    description: |-
      FQDN is the fully qualified domain name of the the servers behind the services
      matched by object. FQND must match the host part
      of the URI that is used to call a service. It will be used for automatically
      generating service certificates for internal services.
    type: string
    exposed: true
    stored: true
    filterable: true
    format: free
    orderable: true

  - name: IPList
    description: |-
      IPList is the list of IP addresses or subnets of the servers behind the services
      matched by object. This is an optional field and it can be
      automatically populated at runtime by the enforcers if DNS resolution is
      available.
    type: external
    exposed: true
    subtype: ip_list
    stored: true

  - name: action
    description: |-
      Action tells if the processing units matched by the subject are
      consuming or exposing the services matched by object.
    type: enum
    exposed: true
    stored: true
    required: true
    allowed_choices:
    - Consume
    - Expose
    default_value: Expose
    filterable: true
    orderable: true

  - name: object
    description: |-
      Objects contains the tag expression that an object must match in order to
      trigger the hook.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
    example_value:
    - - package=p1

  - name: subject
    description: |-
      subject contains the tag expression that an a processing unit
      must match in order to
      trigger the hook.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    required: true
    example_value:
    - - $identity=processingunit
