{
    "attributes": [
        {
            "description": "Identities contains the list of identity names where the quota will be applied.",
            "exposed": true,
            "name": "identities",
            "required": true,
            "stored": true,
            "subtype": "string",
            "type": "list"
        },
        {
            "description": "Quota contains the maximum number of object matching the policy subject that can be created.",
            "exposed": true,
            "name": "quota",
            "type": "integer"
        },
        {
            "description": "TargetNamespace contains the base namespace from where the count will be done.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "targetNamespace",
            "required": true,
            "stored": true,
            "type": "string"
        }
    ],
    "model": {
        "aliases": [
            "quota",
            "quotas",
            "quotapol",
            "quotapols"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "Quotas Policies allows to set quotas on the number of objects that can be created in a namespace.",
        "entity_name": "QuotaPolicy",
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
        "resource_name": "quotapolicies",
        "rest_name": "quotapolicy"
    }
}
