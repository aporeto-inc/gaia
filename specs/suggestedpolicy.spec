# Model
model:
  rest_name: suggestedpolicy
  resource_name: suggestedpolicies
  entity_name: SuggestedPolicy
  package: jenova
  group: visualization/depmaps
  description: Allows you to obtain network policy suggestions.
  aliases:
  - sugpol
  - sugpols
  - sugg
  - suggs

# Attributes
attributes:
  v1:
  - name: networkAccessPolicies
    description: List of suggested network policies.
    type: refList
    exposed: true
    subtype: networkaccesspolicy
    stored: true
    orderable: true
