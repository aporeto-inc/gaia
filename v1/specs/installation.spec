{
    "attributes": [
        {
            "description": "ID represents the identifier of the installation.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "ID",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "AccountName that should be installed.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "accountName",
            "orderable": true,
            "stored": true,
            "type": "string"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "description": "Installation represents an installation for a given account",
        "entity_name": "Installation",
        "package": "highwind",
        "resource_name": "installations",
        "rest_name": "installation"
    }
}
