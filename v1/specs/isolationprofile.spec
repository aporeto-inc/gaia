{
    "attributes": [
        {
            "description": "CapabilitiesActions identifies the capabilities that should be added or removed from the processing unit.",
            "exposed": true,
            "filterable": true,
            "name": "capabilitiesActions",
            "orderable": true,
            "stored": true,
            "subtype": "cap_map",
            "type": "external"
        },
        {
            "description": "DefaultAction is the default action applied to all syscalls of this profile. Default is \"Allow\".",
            "exposed": true,
            "filterable": true,
            "name": "defaultSyscallAction",
            "orderable": true,
            "stored": true,
            "subtype": "syscall_action",
            "type": "external"
        },
        {
            "description": "SyscallRules is a list of syscall rules that identify actions for particular syscalls.",
            "exposed": true,
            "filterable": true,
            "name": "syscallRules",
            "orderable": true,
            "stored": true,
            "subtype": "syscall_rules",
            "type": "external"
        },
        {
            "description": "TargetArchitectures is the target processor architectures where this profile can be applied. Default all.",
            "exposed": true,
            "filterable": true,
            "name": "targetArchitectures",
            "orderable": true,
            "stored": true,
            "subtype": "arch_list",
            "type": "external"
        }
    ],
    "model": {
        "aliases": [
            "ip"
        ],
        "create": true,
        "delete": true,
        "get": true,
        "update": true,
        "entity_name": "IsolationProfile",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "package": "squall",
        "resource_name": "isolationprofiles",
        "rest_name": "isolationprofile"
    }
}