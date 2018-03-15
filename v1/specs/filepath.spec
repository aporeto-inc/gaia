{
    "attributes": [
        {
            "description": "FilePath refer to the file mount path",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "filepath",
            "required": true,
            "stored": true,
            "type": "string"
        },
        {
            "creation_only": true,
            "description": "server is the server name/ID/IP associated with the file path",
            "exposed": true,
            "filterable": true,
            "format": "free",
            "name": "server",
            "required": true,
            "stored": true,
            "type": "string"
        }
    ],
    "model": {
        "aliases": [
            "fp",
            "fps"
        ],
        "delete": true,
        "get": true,
        "update": true,
        "description": "A File Path represents a random path to a file or a folder. They can be used in aFile Access Policiesin order to allow Processing Units to access them, using various modes (read, write, execute). You will need to use the File Paths tags to set some policies. A good example would bevolume=web or file=/etc/passwd.",
        "entity_name": "FilePath",
        "extends": [
            "@base",
            "@described",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named"
        ],
        "package": "squall",
        "resource_name": "filepaths",
        "rest_name": "filepath"
    }
}
