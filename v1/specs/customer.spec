{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "providerCustomerID specifies the customer id as used by the provider for this customer.",
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
            "name": "providerCustomerID",
            "orderable": true,
            "primary_key": null,
            "read_only": true,
            "required": null,
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
                "AWS",
                "Aporeto"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Aporeto",
            "deprecated": null,
            "description": "providerType specifies the cloud provider.",
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
            "name": "providerType",
            "orderable": true,
            "primary_key": null,
            "read_only": true,
            "required": false,
            "secret": null,
            "setter": null,
            "stored": true,
            "subtype": null,
            "transient": null,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [
        {
            "bulk_create": false,
            "bulk_delete": false,
            "bulk_update": false,
            "create": false,
            "delete": false,
            "deprecated": null,
            "get": true,
            "relationship": "child",
            "rest_name": "invoice",
            "update": false
        }
    ],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": null,
        "entity_name": "Customer",
        "exposed": true,
        "extends": [
            "@identifiable-pk-stored",
            "@timeable"
        ],
        "get": true,
        "package": null,
        "private": null,
        "resource_name": "customers",
        "rest_name": "customer",
        "root": null,
        "update": false
    }
}