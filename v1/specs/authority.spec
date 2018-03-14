{
    "attributes": [
        {
            "autogenerated": true,
            "description": "ID is the identitfier of the Authority.",
            "exposed": true,
            "format": "free",
            "identifier": true,
            "name": "ID",
            "read_only": true,
            "stored": true,
            "type": "string"
        },
        {
            "autogenerated": true,
            "description": "PEM encoded certificate data.",
            "exposed": true,
            "format": "free",
            "name": "certificate",
            "read_only": true,
            "stored": true,
            "type": "string"
        },
        {
            "creation_only": true,
            "description": "CommonName contains the common name of the CA.",
            "exposed": true,
            "format": "free",
            "name": "commonName",
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "creation_only": true,
            "description": "Date of expiration of the authority.",
            "exposed": true,
            "name": "expirationDate",
            "stored": true,
            "type": "time"
        },
        {
            "autogenerated": true,
            "description": "Encrypted private key of the Authority.",
            "format": "free",
            "name": "key",
            "stored": true,
            "type": "string"
        },
        {
            "autogenerated": true,
            "description": "serialNumber of the certificate",
            "exposed": true,
            "format": "free",
            "name": "serialNumber",
            "read_only": true,
            "stored": true,
            "type": "string"
        }
    ],
    "model": {
        "aliases": [
            "ca"
        ],
        "delete": true,
        "description": "Authority represents a certificate authority",
        "entity_name": "Authority",
        "package": "barret",
        "resource_name": "authorities",
        "rest_name": "authority",
        "private": true
    }
}
