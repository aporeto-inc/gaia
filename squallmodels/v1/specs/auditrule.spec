{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": [
                "32-bit",
                "64-bit"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "64-bit",
            "deprecated": null,
            "description": "Architecture is the processor architecture that this rule applies to. It can be 32-bit or 64-bit.",
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
            "name": "architecture",
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
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "GroupName is the name of the group that this rule must be associated with.",
            "exposed": true,
            "filterable": true,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": 31,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "groupName",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "FILE",
                "SYSCALL"
            ],
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "SYSCALL",
            "deprecated": null,
            "description": "RuleType is the type of the audit rule and it can be SYSCALL or FILE.",
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
            "name": "ruleType",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": true,
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
            "autogenerated": null,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "SysCalls is the list of system calls that the rule applies to. It is only valid if ruleType is SYSCALL.",
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
            "name": "sysCalls",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": "AuditRule describes an audit rule that must be applied to an enforer to detect anomalous events. ",
        "entity_name": "AuditRule",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "get": true,
        "package": null,
        "resource_name": "auditrules",
        "rest_name": "auditrule",
        "root": null,
        "update": true
    }
}