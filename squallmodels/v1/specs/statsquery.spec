{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": null,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Fields is a list of fields that can be returned by the query. Similar to \"select fields ... \". Default is \"value\".",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "fields",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "tags_list",
            "transient": null,
            "type": "external",
            "unique": null,
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
            "description": "GroupBy is a list of tags to group by the results.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "groupBy",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "tags_list",
            "transient": null,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "audit",
                "file",
                "flow"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": "Flow",
            "deprecated": null,
            "description": "Measurement is the time series measurement of interest. Options are PUs, flows, file, audit.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "measurement",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Results are the results of the query.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "results",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": null,
            "type": "object",
            "unique": null,
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
            "description": "Tags is a list of tags that will be returned by the query. Corresponding of \"select tags ...\" statement.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "tags",
            "orderable": false,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "tags_list",
            "transient": null,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "Where is the list of where clauses as key/value pairs on allowed tags.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "where",
            "orderable": false,
            "primary_key": null,
            "read_only": false,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": false,
            "subtype": "tags_list",
            "transient": null,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [
            "ts"
        ],
        "create": true,
        "delete": false,
        "description": "TimeSeriesData is a generic API to retrieve time series data stored by the Aporeto system. The API allows different types of queries that are all protected within the namespace of the user.",
        "entity_name": "StatsQuery",
        "extends": [],
        "get": false,
        "package": "\nSquall API",
        "resource_name": "statsqueries",
        "rest_name": "statsquery",
        "root": null,
        "update": false
    }
}
