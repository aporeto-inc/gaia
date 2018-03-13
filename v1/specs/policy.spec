{
    "attributes": [
        {
            "description": "Action defines set of actions that must be enforced when a dependency is met.",
            "exposed": true,
            "name": "action",
            "required": true,
            "stored": true,
            "subtype": "actions_list",
            "type": "external"
        },
        {
            "description": "This is a set of all object tags for matching in the DB",
            "name": "allObjectTags",
            "required": true,
            "stored": true,
            "subtype": "tags_list",
            "type": "external"
        },
        {
            "description": "This is a set of all subject tags for matching in the DB",
            "name": "allSubjectTags",
            "required": true,
            "stored": true,
            "subtype": "tags_list",
            "type": "external"
        },
        {
            "description": "Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.",
            "exposed": true,
            "name": "object",
            "stored": true,
            "subtype": "policies_list",
            "type": "external"
        },
        {
            "description": "Relation describes the required operation to be performed between subjects and objects",
            "exposed": true,
            "name": "relation",
            "stored": true,
            "subtype": "relations_list",
            "type": "external"
        },
        {
            "description": "Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR",
            "exposed": true,
            "name": "subject",
            "required": true,
            "stored": true,
            "subtype": "policies_list",
            "type": "external"
        },
        {
            "allowed_choices": [
                "APIAuthorization",
                "EnforcerProfile",
                "File",
                "Hook",
                "NamespaceMapping",
                "Network",
                "ProcessingUnit",
                "Quota",
                "Syscall",
                "TokenScope"
            ],
            "creation_only": true,
            "description": "Type of the policy",
            "exposed": true,
            "filterable": true,
            "name": "type",
            "primary_key": true,
            "required": true,
            "stored": true,
            "type": "enum"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "description": "[nodoc]",
        "entity_name": "Policy",
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named",
            "@propagated",
            "@schedulable"
        ],
        "package": "squall",
        "resource_name": "policies",
        "rest_name": "policy"
    }
}