{
    "attributes": [
        {
            "description": "AssignedScopes is the the list of scopes that the policiy will assigns.",
            "exposed": true,
            "filterable": true,
            "name": "assignedScopes",
            "orderable": true,
            "stored": true,
            "subtype": "tags_list",
            "type": "external"
        },
        {
            "description": "Subject defines the selection criteria that this policy must match on identiy and scope request information.",
            "exposed": true,
            "filterable": true,
            "name": "subject",
            "orderable": true,
            "stored": true,
            "subtype": "policies_list",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "tsp"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "The TokenScopePolicy defines a set of policies that allow customization of the authorization tokens issued by the Aporeto service. This allows Aporeto generated tokens to be used by external applications.",
        "entity_name": "TokenScopePolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named",
            "@propagated",
            "@schedulable"
        ],
        "package": "squall",
        "resource_name": "tokenscopepolicies",
        "rest_name": "tokenscopepolicy"
    }
}
