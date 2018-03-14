{
    "attributes": [
        {
            "creation_only": true,
            "description": "Represent an jaeger batch",
            "exposed": true,
            "name": "batch",
            "stored": true,
            "subtype": "jaeger_batch",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "sp"
        ],
        "description": "A jaegerbatch is a batch of jaeger spans. This is used by external service to post jaeger span in our private jaeger services",
        "entity_name": "Jaegerbatch",
        "package": "meister",
        "resource_name": "jaegerbatchs",
        "rest_name": "jaegerbatch"
    }
}
