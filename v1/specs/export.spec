{
    "attributes": [
        {
            "autogenerated": true,
            "description": "APIVersion of the api used for the exported data",
            "exposed": true,
            "name": "APIVersion",
            "read_only": true,
            "type": "integer"
        },
        {
            "autogenerated": true,
            "description": "List of all exported audit profiles.",
            "exposed": true,
            "name": "auditProfiles",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of exported external services.",
            "exposed": true,
            "name": "externalServices",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of exported file access policies.",
            "exposed": true,
            "name": "fileAccessPolicies",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of exported file paths.",
            "exposed": true,
            "name": "filePaths",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of all exported isolation profiles.",
            "exposed": true,
            "name": "isolationProfiles",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of exported network policies",
            "exposed": true,
            "name": "networkAccessPolicies",
            "read_only": true,
            "subtype": "exported_data_content",
            "type": "external"
        },
        {
            "autogenerated": true,
            "description": "List of all exported processingUnitPolicies.",
            "exposed": true,
            "name": "processingUnitPolicies",
            "required": true,
            "subtype": "exported_data_content",
            "type": "external"
        }
    ],
    "model": {
        "entity_name": "Export",
        "package": "squall",
        "resource_name": "export",
        "rest_name": "export"
    }
}