{
    "attributes": [
        {
            "description": "CertificateAuthority contains the pem block of the certificate authority used by the remote endpoint.",
            "exposed": true,
            "format": "free",
            "name": "certificateAuthority",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "ClientCertificate contains the client certificate that will be used to connect to the remote endoint.",
            "exposed": true,
            "format": "free",
            "name": "clientCertificate",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "ClientCertificateKey contains the key associated to the clientCertificate.",
            "exposed": true,
            "format": "free",
            "name": "clientCertificateKey",
            "orderable": true,
            "required": true,
            "secret": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Endpoint contains the full address of the remote processor endoint.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "endpoint",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "allowed_choices": [
                "Both",
                "Post",
                "Pre"
            ],
            "default_value": "Pre",
            "description": "Mode define the type of the hook.",
            "exposed": true,
            "filterable": true,
            "name": "mode",
            "orderable": true,
            "stored": true,
            "type": "enum"
        },
        {
            "description": "Subject contains the tag expression that an object must match in order to trigger the hook.",
            "exposed": true,
            "name": "subject",
            "required": true,
            "stored": true,
            "subtype": "policies_list",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "hook",
            "hooks",
            "hookpol",
            "hookpols"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "Hook allows to to define hooks to the write operations in squall. Hooks are sent to an external Rufus server that will do the processing and eventually return a modified version of the object before we save it.",
        "entity_name": "HookPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated"
        ],
        "package": "squall",
        "resource_name": "hookpolicies",
        "rest_name": "hookpolicy"
    }
}
