{
    "attributes": [
        {
            "description": "mappedNamespace is the mapped namespace",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "mappedNamespace",
            "orderable": true,
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Subject is the subject.",
            "exposed": true,
            "name": "subject",
            "orderable": true,
            "subtype": "policies_list",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "nspolicy",
            "nspolicies",
            "nsmap",
            "nsmaps"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "A Namespace Mapping Policy defines in which namespace aProcessing Unit should be placed when it is created, based on its tags. When an Aporeto Agent creates a new Processing Unit, the system will place it in its own namespace if no matching Namespace Mapping Policy can be found. If one match is found, then the Processing will be bumped down to the namespace declared in the policy. If it finds in that child namespace another matching Namespace Mapping Policy, then the Processing Unit will be bumped down again, until it reach a namespace with no matching policies.  This is very useful to dispatch processes and containers into a particular namespace, based on a lot of factor. You can put in place a quarantine namespace that will grab all Processing Units with too much vulnerabilities for instances.",
        "entity_name": "NamespaceMappingPolicy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-nopk-nostored",
            "@metadatable",
            "@named"
        ],
        "package": "squall",
        "resource_name": "namespacemappingpolicies",
        "rest_name": "namespacemappingpolicy"
    }
}
