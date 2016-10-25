{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Address provides the current IP address of the server after its initialized.",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "address",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Domain refers to the discovered domain name of the server",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "domain",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "AWS",
                "GCP",
                "Private"
            ],
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Environment describes where the server will be running.",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "environment",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "CONNECTED",
                "INITIALIZED",
                "UNKNOWN"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": "UNKNOWN",
            "deprecated": false,
            "description": "Operational status of the server",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "operationalStatus",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": true,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "certificate",
            "update": false
        }
    ],
    "model": {
        "create": false,
        "delete": true,
        "description": "Server runs all the containers.",
        "entity_name": "Server",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@named"
        ],
        "get": true,
        "package": "Infrastructure",
        "resource_name": "servers",
        "rest_name": "server",
        "root": null,
        "update": true
    }
}