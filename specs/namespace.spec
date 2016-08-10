{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Authenticator for this namespace",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "authenticator",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
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
            "autogenerated": null,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Name is the name of the namespace.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": true,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": true,
            "primary_key": true,
            "read_only": false,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": false,
        "delete": true,
        "description": "Namespace attached to an user.",
        "entity_name": "Namespace",
        "extends": [
            "@base",
            "@described",
            "@identifiable-nopk-stored"
        ],
        "get": true,
        "package": "Infrastructure",
        "resource_name": "namespaces",
        "rest_name": "namespace",
        "root": null,
        "update": true
    }
}