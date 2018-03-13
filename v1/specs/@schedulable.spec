{
    "attributes": [
        {
            "allowed_chars": "^[0-9]+[smh]$",
            "description": "ActiveDuration defines for how long the policy will be active according to the activeSchedule.",
            "exposed": true,
            "format": "free",
            "getter": true,
            "name": "activeDuration",
            "setter": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "ActiveSchedule defines when the policy should be active using the cron notation. The policy will be active for the given activeDuration.",
            "exposed": true,
            "getter": true,
            "name": "activeSchedule",
            "setter": true,
            "stored": true,
            "subtype": "cron_expression",
            "type": "external"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "rest_name": "@schedulable"
    }
}