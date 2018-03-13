{
    "attributes": [
        {
            "description": "TSDB Fields to set for the report.",
            "exposed": true,
            "name": "fields",
            "subtype": "tsdb_fields",
            "type": "external"
        },
        {
            "allowed_choices": [
                "Audit",
                "Enforcer",
                "FileAccess",
                "Flow",
                "ProcessingUnit",
                "Syscall",
                "User"
            ],
            "description": "Kind contains the kind of report.",
            "exposed": true,
            "name": "kind",
            "type": "enum"
        },
        {
            "description": "Tags contains the tags associated to the data point.",
            "exposed": true,
            "name": "tags",
            "subtype": "tags_map",
            "type": "external"
        },
        {
            "description": "Timestamp contains the time for the report.",
            "exposed": true,
            "name": "timestamp",
            "type": "time"
        },
        {
            "description": "Value contains the value for the report.",
            "exposed": true,
            "name": "value",
            "type": "float"
        }
    ],
    "model": {
        "description": "Post a new statistics report.",
        "entity_name": "Report",
        "package": "zack",
        "resource_name": "reports",
        "rest_name": "report"
    }
}