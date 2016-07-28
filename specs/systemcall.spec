{
    "attributes": [],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": false,
            "get": false,
            "relationship": "member",
            "rest_name": "tag",
            "update": true
        }
    ],
    "model": {
        "create": false,
        "delete": true,
        "description": "SystemCall represents a system call.",
        "entity_name": "SystemCall",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@named"
        ],
        "get": true,
        "package": "Policies",
        "resource_name": "systemcalls",
        "rest_name": "systemcall",
        "root": null,
        "update": true
    }
}