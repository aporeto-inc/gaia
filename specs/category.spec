# Model
model:
  rest_name: category
  resource_name: categories
  entity_name: Category
  package: highwind
  group: integration/app
  description: Allows you to categorize services.
  extends:
  - '@identifiable-stored'
  - '@named'
  - '@described'

# Indexes
indexes:
- - :no-inherit
