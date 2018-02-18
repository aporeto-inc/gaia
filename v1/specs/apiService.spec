{
    "attributes": [],
    "children": [],
    "model": {
        "aliases": [],
        "create": null,
        "delete": true,
        "description": "APIService descibes a L4/L7 service and the corresponding implementation. It allows users to define their services, the APIs that they expose, the implementation of the service. These definitions can be used by network policy in order to define advanced controls based on the APIs.",
        "entity_name": "Apiservice",
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
        "resource_name": "apiServices",
        "rest_name": "apiService",
        "root": null,
        "update": true
    }
}