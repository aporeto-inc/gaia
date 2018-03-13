{
    "attributes": [
        {
            "autogenerated": true,
            "description": "Headers contains the requests headers that matched.",
            "exposed": true,
            "name": "headers",
            "read_only": true,
            "subtype": "string",
            "type": "list"
        },
        {
            "autogenerated": true,
            "description": "Rows contains the tabulated data.",
            "exposed": true,
            "name": "rows",
            "read_only": true,
            "subtype": "tabulated_data",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "TargetIdentity contains the requested target identity.",
            "exposed": true,
            "format": "free",
            "name": "targetIdentity",
            "read_only": true,
            "type": "string"
        }
    ],
    "model": {
        "aliases": [
            "table",
            "tables",
            "tabs",
            "tab"
        ],
        "description": "Tabulate API allows you to retrieve a custom table view for any identity using any tags you like as columns.",
        "entity_name": "Tabulation",
        "package": "squall",
        "resource_name": "tabulations",
        "rest_name": "tabulation"
    }
}