# Model
model:
  rest_name: category
  resource_name: categories
  entity_name: Category
  package: highwind
  description: Category allows to categorized services

# Attributes
attributes:
- name: ID
  description: ID is the identifier of the category.
  type: string
  exposed: true
  stored: true
  filterable: true
  format: free
  identifier: true
  orderable: true
  primary_key: true

- name: description
  description: Description is the desription of the category.
  type: string
  exposed: true
  format: free

- name: name
  description: Name of the category.
  type: string
  exposed: true
  stored: true
  filterable: true
  format: free
  orderable: true
