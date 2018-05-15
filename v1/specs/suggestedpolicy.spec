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
    type: external
    exposed: true
    subtype: network_access_policies_list
    stored: true
    filterable: true
    orderable: true
