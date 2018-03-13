{
    "attributes": [
        {
            "description": "Propagated indicates if the audit profile is propagated.",
            "exposed": true,
            "filterable": true,
            "name": "propagated",
            "stored": true,
            "type": "boolean"
        },
        {
            "description": "Rules is the list of audit policy rules associated with this policy.",
            "exposed": true,
            "name": "rules",
            "stored": true,
            "subtype": "audit_profile_rule_list",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "ap"
        ],
        "create": true,
        "delete": true,
        "get": true,
        "update": true,
        "description": "AuditProfile is an audit policy that consists of a set of audit rules. An audit policy will determine that types of events that must be captured in the kernel.",
        "entity_name": "AuditProfile",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "package": "squall",
        "resource_name": "auditprofiles",
        "rest_name": "auditprofile"
    }
}