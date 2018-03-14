{
    "attributes": [
        {
            "description": "Password contains the new password.",
            "exposed": true,
            "format": "free",
            "name": "password",
            "required": true,
            "type": "string"
        },
        {
            "description": "Token contains the reset password token",
            "exposed": true,
            "format": "free",
            "name": "token",
            "required": true,
            "type": "string"
        }
    ],
    "model": {
        "entity_name": "PasswordReset",
        "package": "vince",
        "resource_name": "passwordreset",
        "rest_name": "passwordreset"
    }
}
