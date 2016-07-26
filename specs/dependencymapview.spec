{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Color of the dependency map group",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "color",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Values used by the dependency map group",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "values",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": true,
            "setter": null,
            "stored": true,
            "subtype": "string",
            "transient": false,
            "type": "list",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": null,
        "delete": true,
        "description": "A DependencyMapView is used to store the various for the dependency map using a bunch of selectors.",
        "entity_name": "DependencyMapView",
        "extends": [
            "@base"
        ],
        "get": true,
        "package": "Visualization",
        "resource_name": "dependencymapviews",
        "rest_name": "dependencymapview",
        "root": null,
        "update": true
    }
}