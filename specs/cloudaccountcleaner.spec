# Model
model:
  rest_name: cloudaccountcleaner
  resource_name: cloudaccountcleaner
  entity_name: CloudAccountCleaner
  package: yeul
  group: pcn/infrastructure
  description: |-
    Used for garbage collection of all objects in an account that have not been
    updated since the provided time.

# Attributes
attributes:
  v1:
  - name: date
    description: The date after which objects must be cleaned.
    type: time
    exposed: true
    required: true
    example_value: 2021-03-16 09:30:04 -0700 PDT
