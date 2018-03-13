{
    "attributes": [
        {
            "description": "AllowsExecute allows to execute the files.",
            "exposed": true,
            "filterable": true,
            "name": "allowsExecute",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "AllowsRead allows to read the files.",
            "exposed": true,
            "filterable": true,
            "name": "allowsRead",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "AllowsWrite allows to write the files.",
            "exposed": true,
            "filterable": true,
            "name": "allowsWrite",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "EncryptionEnabled will enable the automatic encryption",
            "exposed": true,
            "filterable": true,
            "name": "encryptionEnabled",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "LogsEnabled will enable logging when this policy is used.",
            "exposed": true,
            "filterable": true,
            "name": "logsEnabled",
            "orderable": true,
            "type": "boolean"
        },
        {
            "description": "Object is the object of the policy.",
            "exposed": true,
            "name": "object",
            "orderable": true,
            "subtype": "policies_list",
            "type": "external"
        },
        {
            "description": "Subject is the subject of the policy",
            "exposed": true,
            "name": "subject",
            "orderable": true,
            "subtype": "policies_list",
            "type": "external"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "description": "A File Access Policy allows Processing Units to access various folder and files. It will use the tags of a File Path to know what is the path of the file or folder to allow access to. You can allow the Processing Unit to have any combination of read, write or execute. Note: When a Processing Unit is Docker container, then it will police the volumes mount. executewon't have any effect. Note: File path are not supported yet for standard Linux processes.",
        "entity_name": "FileAccessPolicy",
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
        "resource_name": "fileaccesspolicies",
        "rest_name": "fileaccesspolicy"
    }
}