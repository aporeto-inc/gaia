{
    "attributes": [
        {
            "description": "Propagate will propagate the policy to all of its children.",
            "exposed": true,
            "filterable": true,
            "getter": true,
            "name": "propagate",
            "orderable": true,
            "setter": true,
            "stored": true,
            "type": "boolean"
        },
        {
            "description": "If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.",
            "exposed": true,
            "filterable": true,
            "getter": true,
            "name": "propagationHidden",
            "orderable": true,
            "setter": true,
            "stored": true,
            "type": "boolean"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "rest_name": "@propagated"
    }
}