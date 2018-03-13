{
    "attributes": [
        {
            "description": "Represents the claims of the currently managed object",
            "exposed": true,
            "name": "claims",
            "required": true,
            "subtype": "string",
            "type": "list"
        },
        {
            "description": "Represents data received from the service",
            "exposed": true,
            "name": "input",
            "required": true,
            "subtype": "raw_json",
            "type": "external"
        },
        {
            "allowed_choices": [
                "Post",
                "Pre"
            ],
            "description": "Node defines the type of the hook",
            "exposed": true,
            "name": "mode",
            "required": true,
            "type": "enum"
        },
        {
            "description": "Represents the current namespace",
            "exposed": true,
            "format": "free",
            "name": "namespace",
            "required": true,
            "type": "string"
        },
        {
            "description": "Define the operation that is currently handled by the service",
            "exposed": true,
            "name": "operation",
            "required": true,
            "subtype": "elemental_operation",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "Returns the OutputData filled with the processor information",
            "exposed": true,
            "name": "output",
            "read_only": true,
            "subtype": "elemental_identitifable",
            "type": "external"
        },
        {
            "description": "RequestID gives the id of the request coming from the main server.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "requestID",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Represents the Identity name of the managed object",
            "exposed": true,
            "format": "free",
            "name": "targetIdentity",
            "required": true,
            "type": "string"
        }
    ],
    "model": {
        "aliases": [
            "hks",
            "hk"
        ],
        "description": "Hook to integrate an Aporeto service.",
        "entity_name": "RemoteProcessor",
        "package": "rufus",
        "resource_name": "remoteprocessors",
        "rest_name": "remoteprocessor"
    }
}