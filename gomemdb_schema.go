package gaia

/*
import "github.com/hashicorp/go-memdb"

var schema = &memdb.DBSchema{
    Tables: map[string]*memdb.TableSchema{

        "account":  &memdb.TableSchema{
            Name: "account",
            Indexes: map[string]*memdb.IndexSchema{

                "name": &memdb.IndexSchema{
                    Name:    "name",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "email": &memdb.IndexSchema{
                    Name:    "email",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "Email"},
                },
                "activationToken": &memdb.IndexSchema{
                    Name:    "activationToken",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "ActivationToken"},
                },
                "resetPasswordToken": &memdb.IndexSchema{
                    Name:    "resetPasswordToken",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "ResetPasswordToken"},
                },
            },
        },

        "accountcheck":  &memdb.TableSchema{
            Name: "accountcheck",
        },

        "activate":  &memdb.TableSchema{
            Name: "activate",
        },

        "activity":  &memdb.TableSchema{
            Name: "activity",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-date": &memdb.IndexSchema{
                    Name:    "namespace-date",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Date"},
                },
                "namespace-operation": &memdb.IndexSchema{
                    Name:    "namespace-operation",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Operation"},
                },
            },
        },

        "alarm":  &memdb.TableSchema{
            Name: "alarm",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-kind": &memdb.IndexSchema{
                    Name:    "namespace-kind",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Kind"},
                },
            },
        },

        "apiauthorizationpolicy":  &memdb.TableSchema{
            Name: "apiauthorizationpolicy",
        },

        "apicheck":  &memdb.TableSchema{
            Name: "apicheck",
        },

        "app":  &memdb.TableSchema{
            Name: "app",
        },

        "appcredential":  &memdb.TableSchema{
            Name: "appcredential",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "auditprofile":  &memdb.TableSchema{
            Name: "auditprofile",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "auditreport":  &memdb.TableSchema{
            Name: "auditreport",
        },

        "auth":  &memdb.TableSchema{
            Name: "auth",
        },

        "authority":  &memdb.TableSchema{
            Name: "authority",
            Indexes: map[string]*memdb.IndexSchema{

                "$hashed:serialNumber": &memdb.IndexSchema{
                    Name:    "$hashed:serialNumber",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "SerialNumber"},
                },
                "serialNumber": &memdb.IndexSchema{
                    Name:    "serialNumber",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "SerialNumber"},
                },
                "commonName": &memdb.IndexSchema{
                    Name:    "commonName",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "CommonName"},
                },
            },
        },

        "automation":  &memdb.TableSchema{
            Name: "automation",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "automationtemplate":  &memdb.TableSchema{
            Name: "automationtemplate",
        },

        "awsaccount":  &memdb.TableSchema{
            Name: "awsaccount",
        },

        "awsapigateway":  &memdb.TableSchema{
            Name: "awsapigateway",
        },

        "awsregister":  &memdb.TableSchema{
            Name: "awsregister",
        },

        "category":  &memdb.TableSchema{
            Name: "category",
        },

        "certificate":  &memdb.TableSchema{
            Name: "certificate",
            Indexes: map[string]*memdb.IndexSchema{

                "commonName": &memdb.IndexSchema{
                    Name:    "commonName",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "CommonName"},
                },
                "parentID-commonName": &memdb.IndexSchema{
                    Name:    "parentID-commonName",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "CommonName"},
                },
            },
        },

        "claimmapping":  &memdb.TableSchema{
            Name: "claimmapping",
        },

        "credential":  &memdb.TableSchema{
            Name: "credential",
        },

        "customer":  &memdb.TableSchema{
            Name: "customer",
            Indexes: map[string]*memdb.IndexSchema{

                "providerCustomerID": &memdb.IndexSchema{
                    Name:    "providerCustomerID",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "ProviderCustomerID"},
                },
            },
        },

        "dependencymap":  &memdb.TableSchema{
            Name: "dependencymap",
        },

        "email":  &memdb.TableSchema{
            Name: "email",
        },

        "endpoint":  &memdb.TableSchema{
            Name: "endpoint",
        },

        "enforcer":  &memdb.TableSchema{
            Name: "enforcer",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "enforcerprofile":  &memdb.TableSchema{
            Name: "enforcerprofile",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "enforcerprofilemappingpolicy":  &memdb.TableSchema{
            Name: "enforcerprofilemappingpolicy",
        },

        "enforcerreport":  &memdb.TableSchema{
            Name: "enforcerreport",
        },

        "eventlog":  &memdb.TableSchema{
            Name: "eventlog",
        },

        "export":  &memdb.TableSchema{
            Name: "export",
        },

        "externalnetwork":  &memdb.TableSchema{
            Name: "externalnetwork",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
                "namespace-archived": &memdb.IndexSchema{
                    Name:    "namespace-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
            },
        },

        "externalservice":  &memdb.TableSchema{
            Name: "externalservice",
        },

        "fileaccesspolicy":  &memdb.TableSchema{
            Name: "fileaccesspolicy",
        },

        "fileaccessreport":  &memdb.TableSchema{
            Name: "fileaccessreport",
        },

        "filepath":  &memdb.TableSchema{
            Name: "filepath",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
                "archived": &memdb.IndexSchema{
                    Name:    "archived",
                    Unique:  false,
                },
            },
        },

        "flowreport":  &memdb.TableSchema{
            Name: "flowreport",
        },

        "graphedge":  &memdb.TableSchema{
            Name: "graphedge",
        },

        "graphgroup":  &memdb.TableSchema{
            Name: "graphgroup",
        },

        "graphnode":  &memdb.TableSchema{
            Name: "graphnode",
        },

        "graphpolicyinfo":  &memdb.TableSchema{
            Name: "graphpolicyinfo",
        },

        "hookpolicy":  &memdb.TableSchema{
            Name: "hookpolicy",
        },

        "hostservice":  &memdb.TableSchema{
            Name: "hostservice",
        },

        "httpresourcespec":  &memdb.TableSchema{
            Name: "httpresourcespec",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-archived": &memdb.IndexSchema{
                    Name:    "namespace-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "import":  &memdb.TableSchema{
            Name: "import",
        },

        "installedapp":  &memdb.TableSchema{
            Name: "installedapp",
            Indexes: map[string]*memdb.IndexSchema{

                "accountName-name": &memdb.IndexSchema{
                    Name:    "accountName-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "invoice":  &memdb.TableSchema{
            Name: "invoice",
        },

        "invoicerecord":  &memdb.TableSchema{
            Name: "invoicerecord",
        },

        "ipinfo":  &memdb.TableSchema{
            Name: "ipinfo",
        },

        "isolationprofile":  &memdb.TableSchema{
            Name: "isolationprofile",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "issue":  &memdb.TableSchema{
            Name: "issue",
        },

        "jaegerbatch":  &memdb.TableSchema{
            Name: "jaegerbatch",
        },

        "k8scluster":  &memdb.TableSchema{
            Name: "k8scluster",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
            },
        },

        "log":  &memdb.TableSchema{
            Name: "log",
        },

        "message":  &memdb.TableSchema{
            Name: "message",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "namespace":  &memdb.TableSchema{
            Name: "namespace",
            Indexes: map[string]*memdb.IndexSchema{

                "$hashed:name": &memdb.IndexSchema{
                    Name:    "$hashed:name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "name": &memdb.IndexSchema{
                    Name:    "name",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "namespacemappingpolicy":  &memdb.TableSchema{
            Name: "namespacemappingpolicy",
        },

        "networkaccesspolicy":  &memdb.TableSchema{
            Name: "networkaccesspolicy",
        },

        "oidcprovider":  &memdb.TableSchema{
            Name: "oidcprovider",
            Indexes: map[string]*memdb.IndexSchema{

                "parentID-name": &memdb.IndexSchema{
                    Name:    "parentID-name",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "passwordreset":  &memdb.TableSchema{
            Name: "passwordreset",
        },

        "plan":  &memdb.TableSchema{
            Name: "plan",
        },

        "poke":  &memdb.TableSchema{
            Name: "poke",
        },

        "policy":  &memdb.TableSchema{
            Name: "policy",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-type": &memdb.IndexSchema{
                    Name:    "namespace-type",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Type"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
                "namespace-type-allObjectTags": &memdb.IndexSchema{
                    Name:    "namespace-type-allObjectTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllObjectTags"},
                        },
                    },
                },
                "namespace-type-allSubjectTags": &memdb.IndexSchema{
                    Name:    "namespace-type-allSubjectTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllSubjectTags"},
                        },
                    },
                },
                "namespace-type-allObjectTags-disabled": &memdb.IndexSchema{
                    Name:    "namespace-type-allObjectTags-disabled",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllObjectTags"},
                        },
                    },
                },
                "namespace-type-allSubjectTags-disabled": &memdb.IndexSchema{
                    Name:    "namespace-type-allSubjectTags-disabled",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllSubjectTags"},
                        },
                    },
                },
                "namespace-type-allObjectTags-propagate": &memdb.IndexSchema{
                    Name:    "namespace-type-allObjectTags-propagate",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllObjectTags"},
                        },
                    },
                },
                "namespace-type-allSubjectTags-propagate": &memdb.IndexSchema{
                    Name:    "namespace-type-allSubjectTags-propagate",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Type"},
&memdb.StringSliceFieldIndex{Field: "AllSubjectTags"},
                        },
                    },
                },
                "namespace-fallback": &memdb.IndexSchema{
                    Name:    "namespace-fallback",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
            },
        },

        "policyrefresh":  &memdb.TableSchema{
            Name: "policyrefresh",
        },

        "policyrenderer":  &memdb.TableSchema{
            Name: "policyrenderer",
        },

        "policyrule":  &memdb.TableSchema{
            Name: "policyrule",
        },

        "privatekey":  &memdb.TableSchema{
            Name: "privatekey",
        },

        "processingunit":  &memdb.TableSchema{
            Name: "processingunit",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-archived": &memdb.IndexSchema{
                    Name:    "namespace-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-operationalStatus-archived": &memdb.IndexSchema{
                    Name:    "namespace-operationalStatus-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "OperationalStatus"},
                },
                "namespace-normalizedTags-archived": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags-archived",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
&memdb.StringFieldIndex{Field: "Namespace"},
                        },
                    },
                },
            },
        },

        "processingunitpolicy":  &memdb.TableSchema{
            Name: "processingunitpolicy",
        },

        "processingunitservice":  &memdb.TableSchema{
            Name: "processingunitservice",
        },

        "punode":  &memdb.TableSchema{
            Name: "punode",
        },

        "quotacheck":  &memdb.TableSchema{
            Name: "quotacheck",
        },

        "quotapolicy":  &memdb.TableSchema{
            Name: "quotapolicy",
        },

        "remoteprocessor":  &memdb.TableSchema{
            Name: "remoteprocessor",
        },

        "renderedpolicy":  &memdb.TableSchema{
            Name: "renderedpolicy",
        },

        "report":  &memdb.TableSchema{
            Name: "report",
        },

        "restapispec":  &memdb.TableSchema{
            Name: "restapispec",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-archived": &memdb.IndexSchema{
                    Name:    "namespace-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
&memdb.StringFieldIndex{Field: "Namespace"},
                        },
                    },
                },
            },
        },

        "revocation":  &memdb.TableSchema{
            Name: "revocation",
            Indexes: map[string]*memdb.IndexSchema{

                "$hashed:serialNumber": &memdb.IndexSchema{
                    Name:    "$hashed:serialNumber",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "SerialNumber"},
                },
                "serialNumber": &memdb.IndexSchema{
                    Name:    "serialNumber",
                    Unique:  true,
                    Indexer: &memdb.StringFieldIndex{Field: "SerialNumber"},
                },
            },
        },

        "role":  &memdb.TableSchema{
            Name: "role",
        },

        "root":  &memdb.TableSchema{
            Name: "root",
        },

        "service":  &memdb.TableSchema{
            Name: "service",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
                "namespace-archived": &memdb.IndexSchema{
                    Name:    "namespace-archived",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-normalizedTags": &memdb.IndexSchema{
                    Name:    "namespace-normalizedTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "NormalizedTags"},
                        },
                    },
                },
                "allAPITags": &memdb.IndexSchema{
                    Name:    "allAPITags",
                    Unique:  false,
                    Indexer: &memdb.StringSliceFieldIndex{Field: "AllAPITags"},
                },
                "namespace-allAPITags": &memdb.IndexSchema{
                    Name:    "namespace-allAPITags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "AllAPITags"},
                        },
                    },
                },
                "allServiceTags": &memdb.IndexSchema{
                    Name:    "allServiceTags",
                    Unique:  false,
                    Indexer: &memdb.StringSliceFieldIndex{Field: "AllServiceTags"},
                },
                "namespace-allServiceTags": &memdb.IndexSchema{
                    Name:    "namespace-allServiceTags",
                    Unique:  false,
                    Indexer: &memdb.CompoundIndex{
                        Indexes: []memdb.Indexer{
                            &memdb.StringFieldIndex{Field: "Namespace"},
&memdb.StringSliceFieldIndex{Field: "AllServiceTags"},
                        },
                    },
                },
            },
        },

        "servicedependency":  &memdb.TableSchema{
            Name: "servicedependency",
        },

        "squalltag":  &memdb.TableSchema{
            Name: "squalltag",
        },

        "statsinfo":  &memdb.TableSchema{
            Name: "statsinfo",
        },

        "statsquery":  &memdb.TableSchema{
            Name: "statsquery",
        },

        "suggestedpolicy":  &memdb.TableSchema{
            Name: "suggestedpolicy",
        },

        "tabulation":  &memdb.TableSchema{
            Name: "tabulation",
        },

        "tag":  &memdb.TableSchema{
            Name: "tag",
        },

        "taginject":  &memdb.TableSchema{
            Name: "taginject",
        },

        "tagvalue":  &memdb.TableSchema{
            Name: "tagvalue",
        },

        "token":  &memdb.TableSchema{
            Name: "token",
        },

        "tokenscopepolicy":  &memdb.TableSchema{
            Name: "tokenscopepolicy",
        },

        "trigger":  &memdb.TableSchema{
            Name: "trigger",
        },

        "vulnerability":  &memdb.TableSchema{
            Name: "vulnerability",
            Indexes: map[string]*memdb.IndexSchema{

                "zone-zHash": &memdb.IndexSchema{
                    Name:    "zone-zHash",
                    Unique:  false,
                    Indexer: &memdb.UintFieldIndex{Field: "ZHash"},
                },
                "namespace": &memdb.IndexSchema{
                    Name:    "namespace",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Namespace"},
                },
                "namespace-name": &memdb.IndexSchema{
                    Name:    "namespace-name",
                    Unique:  false,
                    Indexer: &memdb.StringFieldIndex{Field: "Name"},
                },
            },
        },

        "x509certificate":  &memdb.TableSchema{
            Name: "x509certificate",
        },

        "x509certificatecheck":  &memdb.TableSchema{
            Name: "x509certificatecheck",
        },

    },
}
*/
