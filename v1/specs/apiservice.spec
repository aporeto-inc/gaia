{
    "attributes": [],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": "APIService descibes a L4/L7 service and the corresponding implementation. It allows users to define their services, the APIs that they expose, the implementation of the service. These definitions can be used by network policy in order to define advanced controls based on the APIs.",
        "entity_name": "APIservice",
        "exposed": true,
        "extends": [
            "@archivable",
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "get": true,
        "package": null,
        "private": null,
        "resource_name": "apiservices",
        "rest_name": "apiservice",
        "root": null,
        "update": true
    }
}