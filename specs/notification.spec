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
            "description": "Deleted is the time when the notification was deleted",
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
            "name": "deletedAt",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
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
            "description": "Limit is the number of layers returned in notification",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "limit",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
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
            "description": "Name is the name of the notification",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "name",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
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
            "description": "New is the new layers that introduced vulnerability",
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
            "name": "new",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": "vulnerabilitylayers",
            "transient": false,
            "type": "external",
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
            "description": "NextPage is the next page number",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "nextPage",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
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
            "description": "NorifiedAt is the time when the notification was sent",
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
            "name": "notifiedAt",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
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
            "description": "Old is the old layers that introduced vulnerability",
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
            "name": "old",
            "orderable": true,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": "vulnerabilitylayers",
            "transient": false,
            "type": "external",
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
            "description": "Page is the page number",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "page",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": null,
        "delete": false,
        "description": null,
        "entity_name": "Notification",
        "extends": [
            "@base",
            "@identifiable-pk-stored"
        ],
        "get": true,
        "package": "Infrastructure",
        "resource_name": "notifications",
        "rest_name": "notification",
        "root": null,
        "update": false
    }
}