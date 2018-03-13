{
    "attributes": [
        {
            "description": "SourceNamespace contains the original namespace of the updated object.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "sourceNamespace",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Type contains the policy type that is affected.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "type",
            "orderable": true,
            "stored": true,
            "type": "string"
        }
    ],
    "model": {
        "description": "PolicyRefresh is sent to client when as a push event when a policy refresh is needed on their side.",
        "entity_name": "PolicyRefresh",
        "package": "squall",
        "resource_name": "policyrefreshs",
        "rest_name": "policyrefresh"
    }
}