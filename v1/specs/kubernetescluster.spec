{
    "attributes": [
        {
            "description": "Link to the API authorization policy.",
            "format": "free",
            "name": "APIAuthorizationPolicyID",
            "stored": true,
            "type": "string"
        },
        {
            "description": "Link to the certificate created in Vince for this cluster ",
            "format": "free",
            "name": "certificateID",
            "stored": true,
            "type": "string"
        },
        {
            "description": "base64 of the .tar.gz file that contains all the .YAMLs files needed to create the aporeto side on your kubernetes Cluster",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "kubernetesDefinitions",
            "orderable": true,
            "read_only": true,
            "type": "string"
        },
        {
            "description": "The name of your cluster",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "name",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Link to your namespace ",
            "format": "free",
            "name": "namespaceID",
            "stored": true,
            "type": "string"
        },
        {
            "description": "ID of the parent account.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "parentID",
            "orderable": true,
            "read_only": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "Regenerates the k8s files and certificates.",
            "exposed": true,
            "name": "regenerate",
            "type": "boolean"
        },
        {
            "description": "The namespace in which the Kubernetes specific namespace will be created. By default your account namespace.",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "targetNamespace",
            "orderable": true,
            "stored": true,
            "type": "string"
        },
        {
            "description": "List of target networks [deprecated]",
            "exposed": true,
            "filterable": true,
            "name": "targetNetworks",
            "orderable": true,
            "stored": true,
            "subtype": "target_networks_list",
            "type": "external"
        }
    ],
    "model": {
        "delete": true,
        "get": true,
        "update": true,
        "description": "Create a remote Kubernetes Cluster integration.",
        "entity_name": "KubernetesCluster",
        "extends": [
            "@identifiable-pk-stored",
            "@timeable"
        ],
        "package": "vince",
        "resource_name": "kubernetesclusters",
        "rest_name": "kubernetescluster"
    }
}