{
    "attributes": [
        {
            "description": "AllowsTraffic if true, the flow will be accepted. Otherwise other actions like \"logs\" can still be done, but the traffic will be rejected.",
            "exposed": true,
            "filterable": true,
            "name": "allowsTraffic",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "DestinationPorts contains the list of allowed ports and ranges.",
            "exposed": true,
            "filterable": true,
            "name": "destinationPorts",
            "orderable": true,
            "subtype": "ports_list",
            "type": "external"
        },
        {
            "description": "EncryptionEnabled defines if the flow has to be encrypted.",
            "exposed": true,
            "filterable": true,
            "name": "encryptionEnabled",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "LogsEnabled defines if the flow has to be logged.",
            "exposed": true,
            "filterable": true,
            "name": "logsEnabled",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "Object of the policy.",
            "exposed": true,
            "name": "object",
            "orderable": true,
            "subtype": "policies_list",
            "type": "external"
        },
        {
            "description": "If set to true, the flow will be in observation mode.",
            "exposed": true,
            "filterable": true,
            "name": "observationEnabled",
            "orderable": true,
            "type": "boolean"
        },
        {
            "allowed_choices": [
                "Apply",
                "Continue"
            ],
            "default_value": "Continue",
            "description": "If observationEnabled is set to true, this will be the final action taken on the packets.",
            "exposed": true,
            "filterable": true,
            "name": "observedTrafficAction",
            "orderable": true,
            "type": "enum"
        },
        {
            "description": "List of tags expressions to match the list of entity to pass the flow through.",
            "exposed": true,
            "name": "passthrough",
            "subtype": "policies_list",
            "type": "external"
        },
        {
            "description": "Subject of the policy.",
            "exposed": true,
            "name": "subject",
            "orderable": true,
            "subtype": "policies_list",
            "type": "external"
        }
    ],
    "children": [
        {
            "rest_name": "externalservice",
            "get": true
        },
        {
            "rest_name": "processingunit",
            "get": true
        }
    ],
    "model": {
        "aliases": [
            "netpol",
            "netpols"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "Allows to define networking policies to allow or prevent processing units identitied by their tags to talk to other processing units or external services (also identified by their tags).",
        "entity_name": "NetworkAccessPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated",
            "@schedulable"
        ],
        "package": "squall",
        "resource_name": "networkaccesspolicies",
        "rest_name": "networkaccesspolicy"
    }
}
