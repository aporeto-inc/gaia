{
    "attributes": [
        {
            "description": "Attachments is a list of attachments to send",
            "exposed": true,
            "name": "attachments",
            "subtype": "list_attachments",
            "type": "external"
        },
        {
            "description": "Bcc represents email that should be in copy but hidden ",
            "exposed": true,
            "name": "bcc",
            "subtype": "list_emails",
            "type": "external"
        },
        {
            "description": "Cc represents the addresses that should be in copy",
            "exposed": true,
            "name": "cc",
            "subtype": "list_emails",
            "type": "external"
        },
        {
            "description": "Content of the email to send",
            "exposed": true,
            "format": "free",
            "name": "content",
            "type": "string"
        },
        {
            "description": "From represents the sender of the email",
            "exposed": true,
            "format": "email",
            "name": "from",
            "required": true,
            "type": "string"
        },
        {
            "description": "Subject represents the subject of the email",
            "exposed": true,
            "format": "free",
            "name": "subject",
            "type": "string"
        },
        {
            "description": "To represents receivers of the email ",
            "exposed": true,
            "name": "to",
            "subtype": "list_emails",
            "type": "external"
        },
        {
            "allowed_choices": [
                "HTML",
                "Plain"
            ],
            "default_value": "Plain",
            "description": "Type represents the type of the content.",
            "exposed": true,
            "name": "type",
            "type": "enum"
        }
    ],
    "model": {
        "description": "Email is a message that can be send via email",
        "entity_name": "Email",
        "package": "yuffie",
        "resource_name": "emails",
        "rest_name": "email",
        "private": true
    }
}