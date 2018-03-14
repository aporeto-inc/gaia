{
    "attributes": [
        {
            "description": "LoadbalancerAddresses represents the list of adresses of the external services of type LoadBalancer.",
            "exposed": true,
            "name": "loadbalancerAddresses",
            "stored": true,
            "subtype": "addresses_list",
            "type": "external"
        },
        {
            "description": "LoadbalancerPortsMapping is the list of ports mapped by an extenral service of type load balancer. ",
            "exposed": true,
            "name": "loadbalancerPortsMapping",
            "stored": true,
            "subtype": "portmapping_list",
            "type": "external"
        },
        {
            "description": "Network refers to either CIDR or domain name",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "network",
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "allowed_chars": "^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$",
            "default_value": "1:65535",
            "description": "Port refers to network port which could be a single number or 100:2000 to represent a range of ports",
            "exposed": true,
            "filterable": true,
            "name": "port",
            "stored": true,
            "type": "string"
        },
        {
            "allowed_chars": "^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$",
            "description": "Protocol refers to network protocol like TCP/UDP or the number of the protocol.",
            "exposed": true,
            "filterable": true,
            "name": "protocol",
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "allowed_choices": [
                "LoadBalancerHTTP",
                "LoadBalancerTCP",
                "Network"
            ],
            "default_value": "Network",
            "description": "Type represents the type of external service.",
            "exposed": true,
            "filterable": true,
            "name": "type",
            "orderable": true,
            "stored": true,
            "type": "enum"
        }
    ],
    "model": {
        "aliases": [
            "extsrv",
            "extsrvs"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "An External Service represents a random network or ip that is not managed by the system. They can be used in Network Access Policies in order to allow traffic from or to the declared network or IP, using the provided protocol and port or ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0 as address, and 1-65000 for the ports. You will need to use the External Services tags to set some policies.",
        "entity_name": "ExternalService",
        "extends": [
            "@archivable",
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "package": "squall",
        "resource_name": "externalservices",
        "rest_name": "externalservice"
    }
}
