attributes:
- description: ID is the identifier of the category.
  exposed: true
  filterable: true
  format: free
  identifier: true
  name: ID
  orderable: true
  primary_key: true
  stored: true
  type: string
- description: Description is the desription of the category.
  exposed: true
  format: free
  name: description
  type: string
- description: Name of the category.
  exposed: true
  filterable: true
  format: free
  name: name
  orderable: true
  stored: true
  type: string
model:
  description: Category allows to categorized services
  entity_name: Category
  package: highwind
  resource_name: categories
  rest_name: category
