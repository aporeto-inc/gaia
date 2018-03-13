{
    "attributes": [
        {
            "creation_only": true,
            "description": "Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.",
            "exposed": true,
            "filterable": true,
            "getter": true,
            "name": "metadata",
            "setter": true,
            "stored": true,
            "subtype": "metadata_list",
            "type": "external"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "rest_name": "@metadatable"
    }
}