{
    "attributes": [
        {
            "description": "APIServices provides the APIServices of this policy rule.",
            "exposed": true,
            "name": "APIServices",
            "subtype": "api_services_entities",
            "type": "external"
        },
        {
            "description": "Action defines set of actions that must be enforced when a dependency is met.",
            "exposed": true,
            "name": "action",
            "subtype": "actions_list",
            "type": "external"
        },
        {
            "description": "EnforcerProfiles provides the information about the server profile.",
            "exposed": true,
            "name": "enforcerProfiles",
            "subtype": "enforcerprofiles_list",
            "type": "external"
        },
        {
            "description": "Policy target networks ",
            "exposed": true,
            "name": "externalServices",
            "subtype": "network_entities",
            "type": "external"
        },
        {
            "description": "Policy target networks ",
            "exposed": true,
            "name": "filePaths",
            "subtype": "file_entities",
            "type": "external"
        },
        {
            "description": "IsolationProfiles are the isolation profiles of the rule.",
            "exposed": true,
            "name": "isolationProfiles",
            "subtype": "isolation_profile_entities",
            "type": "external"
        },
        {
            "description": "Policy target networks ",
            "exposed": true,
            "name": "namespaces",
            "subtype": "namespace_entities",
            "type": "external"
        },
        {
            "description": "List of external services the policy mandate to pass through before reaching the destination.",
            "exposed": true,
            "name": "passthroughExternalServices",
            "subtype": "network_entities",
            "type": "external"
        },
        {
            "description": "Propagated indicates if the policy is propagated.",
            "exposed": true,
            "name": "propagated",
            "type": "boolean"
        },
        {
            "description": "Relation describes the required operation to be performed between subjects and objects",
            "exposed": true,
            "name": "relation",
            "subtype": "relations_list",
            "type": "external"
        },
        {
            "description": "Policy target tags",
            "exposed": true,
            "name": "tagClauses",
            "subtype": "target_tags",
            "type": "external"
        }
    ],
    "model": {
        "get": true,
        "description": "PolicyRule is an internal policy resolution API. Services can use this API to retrieve a policy resolution. ",
        "entity_name": "PolicyRule",
        "extends": [
            "@identifiable-nopk-nostored",
            "@named"
        ],
        "package": "squall",
        "resource_name": "policyrules",
        "rest_name": "policyrule"
    }
}