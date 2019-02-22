# Model
model:
  rest_name: policyttl
  resource_name: policyttls
  entity_name: PolicyTTL
  package: squall
  group: core/policy
  description: |-
    This is an unexposed api that defines a helper document allowing to handle
    pushes on objects that are deleted by TTL.
  private: true
  extends:
  - '@identifiable-stored'

# Attributes
attributes:
  v1:
  - name: expirationTime
    description: Time when the policy must be deleted.
    type: time
    stored: true
