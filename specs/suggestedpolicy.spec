# Model
model:
  rest_name: suggestedpolicy
  resource_name: suggestedpolicies
  entity_name: SuggestedPolicy
  package: jenova
  description: Allows to get policy suggestions.
  aliases:
  - sugpol
  - sugpols
  - sugg
  - suggs

# Attributes
attributes:
  v1:
  - name: networkAccessPolicies
    description: List of suggested network access policies.
    type: refList
    exposed: true
    subtype: networkaccesspolicy
    stored: true
    orderable: true
