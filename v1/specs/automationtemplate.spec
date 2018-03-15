{
    "attributes": [
        {
            "description": "Entitlements contains the entitlements needed for executing the function.",
            "exposed": true,
            "name": "entitlements",
            "subtype": "automation_entitlements",
            "type": "external"
        },
        {
            "description": "Function contains the code.",
            "exposed": true,
            "format": "free",
            "name": "function",
            "type": "string"
        },
        {
            "description": "Key contais the unique identifier key for the template.",
            "exposed": true,
            "format": "free",
            "name": "key",
            "type": "string"
        },
        {
            "allowed_choices": [
                "Action",
                "Condition"
            ],
            "default_value": "Condition",
            "description": "Kind represents the kind of template.",
            "exposed": true,
            "name": "kind",
            "type": "enum"
        },
        {
            "description": "Parameters contains the parameter description of the function.",
            "exposed": true,
            "name": "parameters",
            "subtype": "automation_template_parameters",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "autotmpl"
        ],
        "get": true,
        "description": "Templates that ca be used in automations",
        "entity_name": "AutomationTemplate",
        "extends": [
            "@described",
            "@named"
        ],
        "package": "sephiroth",
        "resource_name": "automationtemplates",
        "rest_name": "automationtemplate"
    }
}
