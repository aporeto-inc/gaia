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
            "description": "Annotation stores additional information about an entity",
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
            "name": "annotations",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "annotations",
            "transient": false,
            "type": "external",
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
            "description": "AssociatedTags are the list of tags attached to an entity",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "associatedTags",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": true,
            "stored": true,
            "subtype": "tags_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "CreatedTime is the time at which the object was created",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "createTime",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Namespace tag attached to an entity",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": "free",
            "getter": true,
            "identifier": false,
            "index": true,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "namespace",
            "orderable": true,
            "primary_key": true,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": true,
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
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "NormalizedTags contains the list of normalized tags of the entities",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "normalizedTags",
            "orderable": false,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": true,
            "stored": true,
            "subtype": "tags_list",
            "transient": true,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Protected defines if the object is protected.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": null,
            "getter": true,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "protected",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "boolean",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "UpdateTime is the time at which an entity was updated.",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": true,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "updateTime",
            "orderable": true,
            "primary_key": false,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": true,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "time",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": false,
        "delete": false,
        "description": null,
        "entity_name": null,
        "extends": [],
        "get": false,
        "package": null,
        "resource_name": null,
        "rest_name": null,
        "root": false,
        "update": false
    }
}