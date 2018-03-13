{
    "attributes": [
        {
            "autogenerated": true,
            "description": "Authorized contains the results of the check.",
            "exposed": true,
            "name": "authorized",
            "read_only": true,
            "subtype": "authorized_identities",
            "type": "external"
        },
        {
            "description": "Namespace is the namespace to use to check the api authentication.",
            "exposed": true,
            "format": "free",
            "name": "namespace",
            "required": true,
            "type": "string"
        },
        {
            "allowed_choices": [
                "Create",
                "Delete",
                "Info",
                "Patch",
                "Retrieve",
                "RetrieveMany",
                "Update"
            ],
            "description": "Operation is the operation you want to check.",
            "exposed": true,
            "filterable": true,
            "name": "operation",
            "orderable": true,
            "stored": true,
            "type": "enum"
        },
        {
            "description": "TargetIdentities contains the list of identities you want to check the authorization.",
            "exposed": true,
            "name": "targetIdentities",
            "required": true,
            "subtype": "identity_list",
            "type": "external"
        },
        {
            "description": "Token is the token to use to check api authentication",
            "exposed": true,
            "format": "free",
            "name": "token",
            "required": true,
            "type": "string"
        }
    ],
    "model": {
        "description": "This API allows to verify is a client identitied by his token is allowed to do some operations on some apis. For example, it allows third party system to impersonate a user and ensure a proxfied request should be allowed.",
        "entity_name": "APICheck",
        "package": "squall",
        "resource_name": "apichecks",
        "rest_name": "apicheck"
    }
}