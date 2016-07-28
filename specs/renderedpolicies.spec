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
            "description": "EgressPolicies lists all the egress policies attached to ProcessingUnit",
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
            "name": "egressPolicies",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": "rendered_policy",
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
            "description": "IngressPolicies lists all the ingress policies attached to ProcessingUnit",
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
            "name": "ingressPolicies",
            "orderable": true,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": "rendered_policy",
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
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Identifier of the ProcessingUnit",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": "free",
            "getter": false,
            "identifier": true,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "processingUnitID",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "setter": false,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": true,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "create": null,
        "delete": null,
        "description": "RenderedPolicies attached to the given set of Subjects.",
        "entity_name": "RenderedPolicies",
        "extends": [],
        "get": true,
        "package": "Policies",
        "resource_name": "renderedpolicies",
        "rest_name": "renderedpolicies",
        "root": null,
        "update": null
    }
}