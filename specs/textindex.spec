# Model
model:
  rest_name: textindex
  resource_name: textindexes
  entity_name: TextIndex
  package: gogole
  group: core/search
  description: Internal storage for full text search.
  private: true
  extends:
  - '@identifiable-stored'
  - '@zoned'

# Indexes
indexes:
- - date
- - objectNamespace
- - objectIdentity
- - objectID
- - objectNamespace
  - objectIdentity
  - objectID

# Attributes
attributes:
  v1:
  - name: data
    description: raw text data that are indexed.
    type: string
    stored: true

  - name: date
    description: Date of the last indexing.
    type: time
    stored: true

  - name: objectID
    description: Contains the ID of the match.
    type: string
    stored: true

  - name: objectIdentity
    description: Contains the identity of the match.
    type: string
    stored: true

  - name: objectNamespace
    description: Contains the namespace of the match.
    type: string
    stored: true

  - name: score
    description: Contains a match score.
    type: float
    stored: true
