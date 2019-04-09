# Model
model:
  rest_name: importrequest
  resource_name: importrequests
  entity_name: ImportRequest
  package: vivi
  group: core
  description: |-
    This API allows to send an import request to create objects to a namespace where
    the requester doesn't normally have the permission to do so (other than creating
    import requests).

    The requester must have the permission to create the request in his namespace
    and the target namespace.

    When the request is created, the status is set to `Draft`. The requester can
    edit the content as much as he desires.
    When he's ready to send the request, he must update the status to be
    `Submitted`.
    The request will then be moved to the target namespace.
    At that point nobody can edit the content of the requests other than adding
    comments.

    The requestee will now see the request, and will either

    - Set the status as `Approved`. This will create the objects in the target
      namespace.

    - Set the status as `Rejected`. The request cannot be edited anymore and can be
      deleted.

    - Set the status back as `Draft`. The request will go back to the requester
      namespace so he can make changes. Once the change are ready, the requester
      will set back the status as `Submitted`.

    The `data` format is the same an `Export`.
  aliases:
  - req
  - reqs
  - ireq
  - ireqs
  get:
    description: Retrieve a single existing import request.
  update:
    description: Update an existing import request.
  delete:
    description: Delete an existing import request.
  extends:
  - '@base'
  - '@described'
  - '@timeable'
  - '@identifiable-stored'
  - '@zonable'
  - '@timeable'

# Attributes
attributes:
  v1:
  - name: comment
    description: Post a new comment that will be added to the commentFeed.
    type: string
    exposed: true
    transient: true

  - name: commentFeed
    description: List of comments that have been added to that request.
    type: refList
    exposed: true
    subtype: comment
    stored: true
    read_only: true
    autogenerated: true
    transient: true
    extensions:
      refMode: pointer

  - name: data
    description: The data to import.
    type: external
    exposed: true
    subtype: map[string][]map[string]interface{}
    stored: true
    required: true
    example_value:
      networkaccesspolicies:
      - action: Allow
        description: Allows Acme to access service A
        logsEnabled: true
        name: allow-acme
        object:
        - - $identity=processingunit
          - $namespace=/acme/prod
          - app=query
        subject:
        - - $identity=processingunit
          - app=partner-data

  - name: requesterClaims
    description: |-
      The identity claims of the requester. This will be populated by the control
      plane.
    type: list
    exposed: true
    subtype: string
    stored: true
    read_only: true
    autogenerated: true
    example_value:
    - '@auth:realm=vince'
    - '@auth:account=acme'

  - name: requesterNamespace
    description: |-
      The namespace from the request was created. This will be populated by the
      control plane.
    type: string
    exposed: true
    stored: true
    read_only: true
    autogenerated: true

  - name: status
    description: |-
      The status of the request. The content of data can only be changed when the
      status is set to `Draft` or `ChangeRequested`. When the status is changed to
      `Submitted`, the request will move to the target namespace for validation.
      If the Status is set to `Approved` the data will be created immediately.
      If the status is set to `Rejected` the request cannot be changed anymore and can
      be deleted.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Draft
    - Submitted
    - Approved
    - Rejected
    default_value: Draft

  - name: submittedOnce
    description: Internal field to know if the request has been submit once.
    type: boolean
    stored: true

  - name: targetNamespace
    description: |-
      The namespace where the request will be sent. The requester can set any
      namespace but he needs to have an autorization to post the request in that
      namespace.
    type: string
    exposed: true
    stored: true
    required: true
    creation_only: true
    example_value: /acme/prod
    transient: true
