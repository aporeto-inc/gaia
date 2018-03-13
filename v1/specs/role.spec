{
    "attributes": [
        {
            "autogenerated": true,
            "description": "Authorizations of the role.",
            "exposed": true,
            "name": "authorizations",
            "read_only": true,
            "subtype": "authorization_map",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "Description is the description of the role.",
            "exposed": true,
            "format": "free",
            "name": "description",
            "read_only": true,
            "type": "string"
        },
        {
            "autogenerated": true,
            "description": "Key is the of the role.",
            "exposed": true,
            "format": "free",
            "name": "key",
            "read_only": true,
            "type": "string"
        },
        {
            "autogenerated": true,
            "description": "Name of the role.",
            "exposed": true,
            "format": "free",
            "name": "name",
            "read_only": true,
            "type": "string"
        }
    ],
    "model": {
        "description": "Roles returns the available roles that can be used with API Authorization Policies.",
        "entity_name": "Role",
        "package": "squall",
        "resource_name": "roles",
        "rest_name": "role"
    }
}