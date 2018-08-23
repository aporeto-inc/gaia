# Model
model:
  rest_name: renderedpolicy
  resource_name: renderedpolicies
  entity_name: RenderedPolicy
  package: squall
  description: Retrieve the aggregated policies applied to a particular processing
    unit.
  aliases:
  - rpol
  - rpols

# Attributes
attributes:
  v1:
  - name: certificate
    description: |-
      Certificate is the certificate associated with this PU. It will identify the PU
      to any internal or external services.
    type: string
    exposed: true
    read_only: true

  - name: dependendServices
    description: DependendServices is the list of services that this processing unit
      depends on.
    type: external
    exposed: true
    subtype: api_services_entities

  - name: egressPolicies
    description: EgressPolicies lists all the egress policies attached to processing
      unit.
    type: external
    exposed: true
    subtype: rendered_policy
    read_only: true
    autogenerated: true

  - name: exposedServices
    description: |-
      ExposedServices is the list of services that this processing unit is
      implementing.
    type: external
    exposed: true
    subtype: api_services_entities

  - name: hashedTags
    description: hashedTags contains the list of tags that matched the policies and
      their hashes.
    type: external
    exposed: true
    subtype: hashed_tags
    read_only: true
    autogenerated: true

  - name: ingressPolicies
    description: IngressPolicies lists all the ingress policies attached to processing
      unit.
    type: external
    exposed: true
    subtype: rendered_policy
    read_only: true
    autogenerated: true

  - name: matchingTags
    description: MatchingTags contains the list of tags that matched the policies.
    type: list
    exposed: true
    subtype: string
    read_only: true
    autogenerated: true

  - name: processingUnit
    description: |-
      Can be set during a POST operation to render a policy on a Processing Unit that
      has not been created yet.
    type: external
    exposed: true
    subtype: processingunit
    required: true
    creation_only: true
    example_value: |-
      {
        "name": "pu",
        "type": "Docker",
        "normalizedTags": [
          "a=a",
          "b=b"
        ]
      }

  - name: processingUnitID
    description: Identifier of the processing unit.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: profile
    description: |-
      Profile is the trust profile of the processing unit that should be used during
      all communications.
    type: external
    exposed: true
    subtype: trust_profile
    read_only: true
    autogenerated: true

  - name: scopes
    description: |-
      Scopes is the set of scopes granted to this processing unit that it has to
      present in HTTP requests.
    type: external
    exposed: true
    subtype: scopes_list
    stored: true
