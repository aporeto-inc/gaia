# Model
model:
  rest_name: claimmapping
  resource_name: claimmappings
  entity_name: ClaimMapping
  package: squall
  description: Represents a mapping from a claim name to an HTTP header.
  detached: true

# Attributes
attributes:
  v1:
  - name: claimName
    description: Claim name is the name of the claim that must be mapped to an HTTP
      header.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/*#&@\+\$~:]+$
    example_value: email

  - name: targetHTTPHeader
    description: The target HTTP header where this claim name must be mapped.
    type: string
    exposed: true
    stored: true
    required: true
    allowed_chars: ^[a-zA-Z0-9-_/*#&@\+\$~:]+$
    example_value: X-Username
