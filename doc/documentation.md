<!-- markdownlint-disable MD024 MD025 -->

# Gaia API Documentation

> Version: 1

| Object                                                        | Description                                                                         |
| -                                                             | -                                                                                   |
| [Account](#account)                                           | This api allows to view and manage basic information about your account like        |
| [Activate](#activate)                                         | Used to activate a pending account.                                                 |
| [Activity](#activity)                                         | Contains all the activity log that happened in a namespace. All successful or       |
| [Alarm](#alarm)                                               | An alarm represents an event requiring attention.                                   |
| [APIAuthorizationPolicy](#apiauthorizationpolicy)             | An API Authorization Policy defines what kind of operations a user of a system      |
| [APICheck](#apicheck)                                         | This API allows to verify is a client identitied by his token is allowed to do      |
| [App](#app)                                                   | App represents an application that can be installed.                                |
| [AppCredential](#appcredential)                               | Create a credential for an application.                                             |
| [AppParameter](#appparameter)                                 | Represents a parameter that can be passed to an app.                                |
| [AuditProfile](#auditprofile)                                 | AuditProfile is an audit policy that consists of a set of audit rules. An audit     |
| [AuditProfileMappingPolicy](#auditprofilemappingpolicy)       | Defines an audit policy that determine the sets of enforcers that must implement... |
| [AuditReport](#auditreport)                                   | Post a new audit statistics report.                                                 |
| [Auth](#auth)                                                 | This API verifies if the given token is valid or not.                               |
| [Automation](#automation)                                     | An automation needs documentation.                                                  |
| [AutomationTemplate](#automationtemplate)                     | Templates that ca be used in automations.                                           |
| [AutomationTemplateParameter](#automationtemplateparameter)   | Represents an Automation template parameter.                                        |
| [AWSAccount](#awsaccount)                                     | Allows to bind an AWS account to your Aporeto account to allow auto registration... |
| [AWSAPIGateway](#awsapigateway)                               | managed API decisions for the AWS API Gateway.                                      |
| [AWSRegister](#awsregister)                                   | This api allows AWS customer to register with Aporeto SaaS for billing.             |
| [Category](#category)                                         | Category allows to categorized services.                                            |
| [Certificate](#certificate)                                   | A User represents the owner of some certificates.                                   |
| [ClaimMapping](#claimmapping)                                 | Represents a mapping from a claim name to an HTTP header.                           |
| [Credential](#credential)                                     | Represents an application credential data.                                          |
| [DependencyMap](#dependencymap)                               | This api returns a data structure representing the graph of all processing units... |
| [DeprecatedHostService](#deprecatedhostservice)               | Represents a service of the enforcer's host.                                        |
| [Endpoint](#endpoint)                                         | Represents an HTTP endpoint.                                                        |
| [Enforcer](#enforcer)                                         | An Enforcer contains all parameters associated with a registered enforcer. The      |
| [EnforcerProfile](#enforcerprofile)                           | Allows to create reusable configuration profile for your enforcers. Enforcer        |
| [EnforcerProfileMappingPolicy](#enforcerprofilemappingpolicy) | A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used... |
| [EnforcerReport](#enforcerreport)                             | Post a new enforcer statistics report.                                              |
| [EnforcerTraceReport](#enforcertracereport)                   | Post a new enforcer trace that determines how packets are.                          |
| [EventLog](#eventlog)                                         | This api allows to report various event on any objects.                             |
| [Export](#export)                                             | Export the policies and related objects in a given namespace.                       |
| [ExternalNetwork](#externalnetwork)                           | An External Network represents a random network or ip that is not managed by the... |
| [ExternalService](#externalservice)                           | This API is deprecated in favor of externalnetworks.                                |
| [FileAccessPolicy](#fileaccesspolicy)                         | A File Access Policy allows Processing Units to access various folder and files.... |
| [FileAccessReport](#fileaccessreport)                         | Post a new file access statistics report.                                           |
| [FilePath](#filepath)                                         | A File Path represents a random path to a file or a folder. They can be used in     |
| [FlowReport](#flowreport)                                     | Post a new flow statistics report.                                                  |
| [GraphEdge](#graphedge)                                       | Represents an edge from the dependency map.                                         |
| [GraphGroup](#graphgroup)                                     | Represents an group of nodes from the dependency map.                               |
| [GraphNode](#graphnode)                                       | Represents an node from the dependency map.                                         |
| [GraphPolicyInfo](#graphpolicyinfo)                           | Represents a policy information.                                                    |
| [HookPolicy](#hookpolicy)                                     | Hook allows to to define hooks to the write operations in squall. Hooks are sent... |
| [HostService](#hostservice)                                   | Represents a set of services that a host must expose and protect.                   |
| [HostServiceMappingPolicy](#hostservicemappingpolicy)         | Defines a host service mapping policy that provides the relation between            |
| [HTTPResourceSpec](#httpresourcespec)                         | HTTPResourceSpec descibes an HTTP resource exposed by a service. These APIs         |
| [Import](#import)                                             | Imports an export of policies and related objects into the namespace.               |
| [InstalledApp](#installedapp)                                 | InstalledApps represents an installed application.                                  |
| [Invoice](#invoice)                                           | This api allows to view invoices for Aporeto customers.                             |
| [InvoiceRecord](#invoicerecord)                               | This api allows to view detailed records of invoices for Aporeto customers.         |
| [IPInfo](#ipinfo)                                             | This apis allows to resolve information from an IP address.                         |
| [IsolationProfile](#isolationprofile)                         | An IsolationProfile needs documentation.                                            |
| [Issue](#issue)                                               | This API issues a new token according to given data.                                |
| [Jaegerbatch](#jaegerbatch)                                   | A jaegerbatch is a batch of jaeger spans. This is used by external service to       |
| [K8SCluster](#k8scluster)                                     | Create a remote Kubernetes Cluster integration.                                     |
| [Log](#log)                                                   | Retrieves the log of a deployed app.                                                |
| [Message](#message)                                           | The Message API allows to post public messages that will be visible through all     |
| [Namespace](#namespace)                                       | A Namespace represents the core organizational unit of the system. All objects      |
| [NamespaceMappingPolicy](#namespacemappingpolicy)             | A Namespace Mapping Policy defines in which namespace a Processing Unit should      |
| [NetworkAccessPolicy](#networkaccesspolicy)                   | Allows to define networking policies to allow or prevent processing units           |
| [OIDCProvider](#oidcprovider)                                 | Allows to declare a generic OpenID Connect provider that can be used in exchange... |
| [PacketReport](#packetreport)                                 | Post a new packet tracing report.                                                   |
| [PasswordReset](#passwordreset)                               | Used to reset an account password.                                                  |
| [Plan](#plan)                                                 | Plan contains the various billing plans available.                                  |
| [Poke](#poke)                                                 | When available, poke can be used to update various information about the parent.... |
| [Policy](#policy)                                             | Policy represents the policy primitive used by all aporeto policies.                |
| [PolicyRefresh](#policyrefresh)                               | PolicyRefresh is sent to client when as a push event when a policy refresh is       |
| [PolicyRule](#policyrule)                                     | PolicyRule is an internal policy resolution API. Services can use this API to       |
| [ProcessingUnit](#processingunit)                             | A Processing Unit reprents anything that can compute. It can be a Docker            |
| [ProcessingUnitPolicy](#processingunitpolicy)                 | A ProcessingUnitPolicies needs a better description.                                |
| [ProcessingUnitService](#processingunitservice)               | Represents a service attached to a processing unit.                                 |
| [QuotaPolicy](#quotapolicy)                                   | Quotas Policies allows to set quotas on the number of objects that can be           |
| [RemoteProcessor](#remoteprocessor)                           | Hook to integrate an Aporeto service.                                               |
| [RenderedPolicy](#renderedpolicy)                             | Retrieve the aggregated policies applied to a particular processing unit.           |
| [Report](#report)                                             | Post a new statistics report.                                                       |
| [Role](#role)                                                 | Roles returns the available roles that can be used with API Authorization           |
| [Root](#root)                                                 | root object.                                                                        |
| [Service](#service)                                           | A Service defines a generic service object at L4 or L7 that encapsulates the        |
| [ServiceDependency](#servicedependency)                       | Allows to define a service dependency where a set of processing units as defined... |
| [StatsInfo](#statsinfo)                                       | StatsInfo API can be used to list the fields and tags available in a                |
| [StatsQuery](#statsquery)                                     | StatsQuery is a generic API to retrieve time series data stored by the Aporeto      |
| [SuggestedPolicy](#suggestedpolicy)                           | Allows to get policy suggestions.                                                   |
| [Tabulation](#tabulation)                                     | Tabulate API allows you to retrieve a custom table view for any identity using      |
| [Tag](#tag)                                                   | A tag is a string in the form of "key=value" that can applied to all objects in     |
| [TagValue](#tagvalue)                                         | Represents all values associated to tag key.                                        |
| [TimeSeriesQueryResults](#timeseriesqueryresults)             | Represent the results of a stats query.                                             |
| [TimeSeriesRow](#timeseriesrow)                               | Represent a time series row.                                                        |
| [TokenScopePolicy](#tokenscopepolicy)                         | The TokenScopePolicy defines a set of policies that allow customization of the      |
| [TraceMode](#tracemode)                                       | TraceMode is the tracing mode that must be applied to a PU.                         |
| [TraceRecord](#tracerecord)                                   | Represents a single trace record from the enforcer.                                 |
| [Trigger](#trigger)                                           | Trigger can be used to remotely trigger an automation.                              |
| [Vulnerability](#vulnerability)                               | A vulnerabily represents a particular CVE.                                          |

## Account

This api allows to view and manage basic information about your account like
your name, password, enable 2 factor authentication.

### Example

```json
{
  "email": "user@aporeto.com",
  "name": "acme"
}
```

### Relations

#### `GET /accounts`

Retrieves all accounts. This is a private API that can only be done by the
system.

##### Parameters

- `name` (string): internal parameters.
- `status` (string): internal parameters.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /accounts`

Creates a new Account.

#### `DELETE /accounts/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /accounts/:id`

Retrieves the object with the given ID.

#### `PUT /accounts/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPAddress (string)`

LDAPAddress holds the account authentication account's private ldap server.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPBaseDN (string)`

LDAPBaseDN holds the base DN to use to ldap queries.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPBindDN (string)`

LDAPBindDN holds the account's internal LDAP bind dn for querying auth.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPBindPassword (string)`

LDAPBindPassword holds the password to the LDAPBindDN.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `LDAPBindSearchFilter (string)`

LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
other systems, value may be 'uid={USERNAME}'.

| Characteristics | Value              |
| -               | -:                 |
| Default         | `"uid={USERNAME}"` |
| Orderable       | `true`             |

#### `LDAPCertificateAuthority (string)`

LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `LDAPConnSecurityProtocol (enum)`

LDAPConnProtocol holds the connection type for the LDAP provider.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `TLS, InbandTLS` |
| Default         | `"InbandTLS"`    |
| Orderable       | `true`           |
| Filterable      | `true`           |

#### `LDAPEnabled (boolean)`

LDAPEnabled triggers if the account uses it's own LDAP for authentication.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPIgnoredKeys (list)`

LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `LDAPSubjectKey (string)`

LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.

| Characteristics | Value   |
| -               | -:      |
| Default         | `"uid"` |
| Orderable       | `true`  |

#### `OTPEnabled (boolean)`

Set to enable or disable two factor authentication.

#### `OTPQRCode (string)`

Returns the base64 encoded QRCode for setting up 2 factor auth.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `accessEnabled (boolean)`

AccessEnabled defines if the account holder should have access to the systems.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `activationToken (string)`

ActivationToken contains the activation token.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

#### `associatedBillingID (string)`

associatedBillingID holds the ID of the associated billing customer.

#### `associatedPlanKey (string)`

AssociatedPlanKey contains the plan key that is associated to this account.

| Characteristics | Value                 |
| -               | -:                    |
| Default         | `"aporeto.plan.free"` |
| Autogenerated   | `true`                |
| Read only       | `true`                |

#### `company (string)`

Company of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `email (string)`

Email of the account holder.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `firstName (string)`

First Name of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `lastName (string)`

Last Name of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `name (string)`

Name of the account.

| Characteristics | Value          |
| -               | -:             |
| Format          | `/^[^\*\=]*$/` |
| Required        | `true`         |
| Creation only   | `true`         |
| Orderable       | `true`         |
| Filterable      | `true`         |

#### `password (string)`

Password for the account.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `reCAPTCHAKey (string)`

ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `status (enum)`

Status of the account.

| Characteristics | Value                                |
| -               | -:                                   |
| Allowed Value   | `Active, Disabled, Invited, Pending` |
| Default         | `"Pending"`                          |
| Autogenerated   | `true`                               |
| Read only       | `true`                               |
| Orderable       | `true`                               |
| Filterable      | `true`                               |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Activate

Used to activate a pending account.

### Relations

#### `GET /activate`

Activates a pending account.

##### Parameters

- `noRedirect` (boolean): If set, do not redirect the request to the UI.
- `token` (string): Activation token.

##### Mandatory Parameters

`token`

### Attributes

#### `token (string)`

Token contains the activation token.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

## Activity

Contains all the activity log that happened in a namespace. All successful or
failed actions will be available, and eventual errors as well as the claims of
the user who triggered the actiions. This log is capped and only keeps the last
50k entries by default.

### Relations

#### `GET /activities`

Retrieves the list of activity logs.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /activities/:id`

Retrieves the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `claims (object)`

Claims of the user who performed the operation.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `data (object)`

Data of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `date (time)`

Date of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `error (object)`

Error contains the eventual error.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `message (string)`

Message of the notification.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `namespace (string)`

Namespace of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `operation (string)`

Operation describe what kind of operation the notification represents.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

#### `originalData (object)`

OriginalData contains the eventual original data of the object that has been
modified.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `source (string)`

Source contains meta information about the source.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `targetIdentity (string)`

TargetIdentity is the Identity of the related object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

## Alarm

An alarm represents an event requiring attention.

### Example

```json
{
  "content": "This is an alarm",
  "kind": "aporeto.alarm.kind",
  "name": "the name"
}
```

### Relations

#### `GET /alarms`

Retrieves all the alarms.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /alarms`

Creates a new alarm.

#### `DELETE /alarms/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /alarms/:id`

Retrieves the object with the given ID.

#### `PUT /alarms/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `content (string)`

Content of the alarm.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `data (external:list_of_maps_of_string_of_strings)`

Data represent user data related to the alams.

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `kind (string)`

Kind identifies the kind of alarms. If two alarms are created with the same
identifier, then only the occurrence will be incremented.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `occurrences (external:list_of_times)`

Number of time this alarm have been seen.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Creation only   | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `status (enum)`

Status of the alarm.

| Characteristics | Value                          |
| -               | -:                             |
| Allowed Value   | `Acknowledged, Open, Resolved` |
| Default         | `"Open"`                       |
| Orderable       | `true`                         |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## APIAuthorizationPolicy

An API Authorization Policy defines what kind of operations a user of a system
can do in a namespace. The operations can be any combination of GET, POST, PUT,
DELETE,PATCH or HEAD. By default, an API Authorization Policy will only give
permissions in the context of the current namespace but you can make it
propagate to all the child namespaces. It is also possible restrict permissions
to apply only on a particular subset of the apis by setting the target
identities.

### Example

```json
{
  "authorizedIdentities": [
    "@auth:role=namespace.editor"
  ],
  "authorizedNamespace": "/namespace",
  "name": "the name"
}
```

### Relations

#### `GET /apiauthorizationpolicies`

Retrieves the list of API authorization policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /apiauthorizationpolicies`

Creates a new API authorization policies.

#### `DELETE /apiauthorizationpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /apiauthorizationpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /apiauthorizationpolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `authorizedIdentities (list)`

AuthorizedIdentities defines the list of api identities the policy applies to.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `authorizedNamespace (string)`

AuthorizedNamespace defines on what namespace the policy applies.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `authorizedSubnets (list)`

If set, the api authorization will only be valid if the request comes from one
the declared subnets.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject is the subject.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## APICheck

This API allows to verify is a client identitied by his token is allowed to do
some operations on some apis. For example, it allows third party system to
impersonate a user and ensure a proxfied request should be allowed.

### Example

```json
{
  "namespace": "/namespace",
  "targetIdentities": [
    "processingunit",
    "enforcer"
  ],
  "token": "valid.jwt.token"
}
```

### Relations

#### `POST /apichecks`

Verifies the authorizations on various identities for a given token.

### Attributes

#### `authorized (external:map_of_string_of_booleans)`

Authorized contains the results of the check.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `claims (list)`

Claims contains the decoded claims used.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `namespace (string)`

Namespace is the namespace to use to check the api authentication.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `operation (enum)`

Operation is the operation you want to check.

| Characteristics | Value                                                         |
| -               | -:                                                            |
| Allowed Value   | `Create, Delete, Info, Patch, Retrieve, RetrieveMany, Update` |
| Orderable       | `true`                                                        |
| Filterable      | `true`                                                        |

#### `targetIdentities (list)`

TargetIdentities contains the list of identities you want to check the
authorization.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `token (string)`

Token is the token to use to check api authentication.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## App

App represents an application that can be installed.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /apps`

Retrieves the list of apps.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

### Attributes

#### `beta (boolean)`

Beta indicates if the app is in a beta version.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `categoryID (string)`

CategoryID of the app.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `icon (string)`

Icon contains a base64 image for the app.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `latestVersion (string)`

LatestVersion represents the latest version available of the app.

#### `longDescription (string)`

LongDescription contains a more detailed description of the app.

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `parameters (refList)`

Parameters is a list of parameters available for the app.

#### `title (string)`

Title represents the title of the app.

## AppCredential

Create a credential for an application.

### Example

```json
{
  "CSR": "-----BEGIN CERTIFICATE REQUEST-----
MIICvDCCAaQCAQAwdzELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxDzANBgNV
BAcMBkxpbmRvbjEWMBQGA1UECgwNRGlnaUNlcnQgSW5jLjERMA8GA1UECwwIRGln
aUNlcnQxHTAbBgNVBAMMFGV4YW1wbGUuZGlnaWNlcnQuY29tMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8+To7d+2kPWeBv/orU3LVbJwDrSQbeKamCmo
wp5bqDxIwV20zqRb7APUOKYoVEFFOEQs6T6gImnIolhbiH6m4zgZ/CPvWBOkZc+c
1Po2EmvBz+AD5sBdT5kzGQA6NbWyZGldxRthNLOs1efOhdnWFuhI162qmcflgpiI
WDuwq4C9f+YkeJhNn9dF5+owm8cOQmDrV8NNdiTqin8q3qYAHHJRW28glJUCZkTZ
wIaSR6crBQ8TbYNE0dc+Caa3DOIkz1EOsHWzTx+n0zKfqcbgXi4DJx+C1bjptYPR
BPZL8DAeWuA8ebudVT44yEp82G96/Ggcf7F33xMxe0yc+Xa6owIDAQABoAAwDQYJ
KoZIhvcNAQEFBQADggEBAB0kcrFccSmFDmxox0Ne01UIqSsDqHgL+XmHTXJwre6D
hJSZwbvEtOK0G3+dr4Fs11WuUNt5qcLsx5a8uk4G6AKHMzuhLsJ7XZjgmQXGECpY
Q4mC3yT3ZoCGpIXbw+iP3lmEEXgaQL0Tx5LFl/okKbKYwIqNiyKWOMj7ZR/wxWg/
ZDGRs55xuoeLDJ/ZRFf9bI+IaCUd1YrfYcHIl3G87Av+r49YVwqRDT0VDV7uLgqn
29XI1PpVUNCPQGn9p/eX6Qo7vpDaPybRtA2R7XLKjQaF9oXWeCUqy1hvJac9QFO2
97Ob1alpHPoZ7mWiEuJwjBPii6a9M9G30nUo39lBi1w=
-----END CERTIFICATE REQUEST-----",
  "name": "the name",
  "roles": [
    "@auth:role=enforcer",
    "@auth:role=kubesquall"
  ]
}
```

### Relations

#### `GET /appcredentials`

Retrieves the list of application credentials.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /appcredentials`

Creates a new application credential.

#### `DELETE /appcredentials/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /appcredentials/:id`

Retrieves the object with the given ID.

#### `PUT /appcredentials/:id`

Updates the object with the given ID.

### Attributes

#### `CSR (string)`

CSR contains the Certificate Signing Request as a PEM encoded string. It can
only be set during a renew.

- The CN **MUST** be `app:credential:<appcred-id>:<appcred-name>`
- The O **MUST** be the namespace of the appcred

If you send anything else, the signing request will be rejected.

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `certificate (string)`

The string representation of the Certificate used by the application.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `credentials (ref)`

The credential data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `email (string)`

The email address that will receive a copy of the application credentials.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `parentIDs (list)`

Contains the ID of the parent Appcred if this is a derived appcred.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `roles (list)`

List of roles to give the credentials.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AppParameter

Represents a parameter that can be passed to an app.

### Attributes

#### `advanced (boolean)`

Defines if the parameter is an advanced one.

#### `allowedValues (list)`

List of values that can be used.

#### `defaultValue (object)`

Default value of the parameter.

#### `description (string)`

Description of the paramerter.

#### `key (string)`

Key identifying the parameter.

#### `longDescription (string)`

Long explanation of the parameter.

#### `name (string)`

Name of the paramerter.

#### `optional (boolean)`

Defines if the parameter is optional.

#### `type (enum)`

The type of the parameter.

| Characteristics | Value                                                                                                             |
| -               | -:                                                                                                                |
| Allowed Value   | `Boolean, Duration, Enum, IntegerSlice, Integer, Float, FloatSlice, Password, String, StringSlice, CVSSThreshold` |

#### `value (object)`

Value of the parameter.

## AuditProfile

AuditProfile is an audit policy that consists of a set of audit rules. An audit
policy will determine that types of events that must be captured in the kernel.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /auditprofiles`

Retrieves the list of audit profiles.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /auditprofiles`

Creates a new audit profile.

#### `DELETE /auditprofiles/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /auditprofiles/:id`

Retrieves the object with the given ID.

#### `PUT /auditprofiles/:id`

Updates the object with the given ID.

#### `GET /auditprofilemappingpolicies/:id/auditprofiles`

Returns the list of audit profiles that are referred by this policy.

#### `GET /enforcerprofiles/:id/auditprofiles`

Returns the list of AuditProfiles used by an enforcer profile.

#### `GET /enforcers/:id/auditprofiles`

Returns a list of the audit profiles that must be applied to this enforcer.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagated (boolean)`

Propagated indicates if the audit profile is propagated.

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `rules (external:_audit_profile_rule_list)`

Rules is the list of audit policy rules associated with this policy.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AuditProfileMappingPolicy

Defines an audit policy that determine the sets of enforcers that must implement
a specific audit profile.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /auditprofilemappingpolicies`

Retrieves the list of audit profile mapping policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /auditprofilemappingpolicies`

Creates a new audit profile mapping policy.

#### `DELETE /auditprofilemappingpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /auditprofilemappingpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /auditprofilemappingpolicies/:id`

Updates the object with the given ID.

#### `GET /auditprofilemappingpolicies/:id/auditprofiles`

Returns the list of audit profiles that are referred by this policy.

#### `GET /auditprofilemappingpolicies/:id/enforcers`

Returns the list of enforcers that are affected by this poliy.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object of the policy is the selector of the audit profiles that must be applied
based on this policy.

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject of the policy is a selector of the enforcers that must implement the
policy.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AuditReport

Post a new audit statistics report.

### Example

```json
{
  "AUID": "xxx-xxx",
  "CWD": "/etc",
  "EXE": "/bin/ls",
  "a0": "xxx-xxx",
  "a1": "xxx-xxx",
  "a2": "xxx-xxx",
  "a3": "xxx-xxx",
  "arch": "x86_64",
  "auditProfileID": "xxx-xxx-xxx-xxx",
  "auditProfileNamespace": "/my/ns",
  "command": "ls",
  "enforcerID": "xxx-xxx-xxx-xxx",
  "enforcerNamespace": "/my/ns",
  "processingUnitID": "xxx-xxx-xxx-xxx",
  "processingUnitNamespace": "/my/ns",
  "recordType": "Syscall",
  "syscall": "execve",
  "timestamp": "2018-06-14T23:10:46.420397985Z"
}
```

### Relations

#### `POST /auditreports`

Create a audit statistics report.

### Attributes

#### `AUID (string)`

Needs documentation.

#### `CWD (string)`

Command working directory.

#### `EGID (integer)`

Needs documentation.

#### `EUID (integer)`

Needs documentation.

#### `EXE (string)`

Path to the executable.

#### `FSGID (integer)`

Needs documentation.

#### `FSUID (integer)`

Needs documentation.

#### `GID (integer)`

Needs documentation.

#### `PER (integer)`

Needs documentation.

#### `PID (integer)`

PID of the executable.

#### `PPID (integer)`

PID of the parent executable.

#### `SGID (integer)`

Needs documentation.

#### `SUID (integer)`

Needs documentation.

#### `UID (integer)`

Needs documentation.

#### `a0 (string)`

Needs documentation.

#### `a1 (string)`

Needs documentation.

#### `a2 (string)`

Needs documentation.

#### `a3 (string)`

Needs documentation.

#### `arch (string)`

Architecture of the system where the syscall happened.

#### `auditProfileID (string)`

ID the audit profile that triggered the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `auditProfileNamespace (string)`

Namespace the audit profile that triggered the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `command (string)`

Command issued.

#### `enforcerID (string)`

ID of the enforcer reporting.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `enforcerNamespace (string)`

Namespace of the enforcer reporting.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `exit (integer)`

Exit code of the executable.

#### `processingUnitID (string)`

ID of the processing unit originating the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `processingUnitNamespace (string)`

Namespace of the processing unit originating the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `recordType (string)`

Type of record.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `sequence (integer)`

Needs documentation.

#### `success (boolean)`

Tells if the operation has been a success of a failure.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Required        | `true`  |

#### `syscall (string)`

Syscall name.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `timestamp (time)`

Date of the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## Auth

This API verifies if the given token is valid or not.

### Relations

#### `GET /auth`

Verify the validity of a token.

##### Parameters

- `token` (string): token to validate.

### Attributes

#### `claims (external:_claims)`

Claims are the claims.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## Automation

An automation needs documentation.

### Example

```json
{
  "condition": "function when(m, params) { return { continue: true }}",
  "name": "the name"
}
```

### Relations

#### `GET /automations`

Retrieves the list of Automations.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /automations`

Creates a new Automation.

#### `DELETE /automations/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /automations/:id`

Retrieves the object with the given ID.

#### `PUT /automations/:id`

Updates the object with the given ID.

#### `GET /automations/:id/triggers`

Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`.

#### `POST /automations/:id/triggers`

Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `actions (list)`

Action contains the code that will be executed if the condition is met.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `condition (string)`

Condition contains the code that will be executed to decide if any action should
be taken.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `entitlements (external:_automation_entitlements)`

Entitlements declares which operations are allowed on which identities.

#### `errors (list)`

Error contains the eventual error of the last run.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `events (external:_automation_events)`

Events contains the identity and operation an event must have to trigger the
automation.

#### `lastExecTime (time)`

LastExecTime holds the last successful execution tine.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `parameters (external:map_of_string_of_objects)`

Parameters are passed to the functions.

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `schedule (string)`

Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.

#### `stdout (string)`

Stdout contains the last run standard output.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `token (string)`

Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Creation only   | `true` |

#### `tokenRenew (boolean)`

If set to true a new token will be issued, and the previous one invalidated.

#### `trigger (enum)`

Trigger controls when the automation should be triggered.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Event, RemoteCall, Time` |
| Default         | `"Time"`                  |
| Orderable       | `true`                    |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AutomationTemplate

Templates that ca be used in automations.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /automationtemplates`

Retrieves the list of automation templates.

#### `GET /automationtemplates/:id`

Retrieves the object with the given ID.

### Attributes

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `entitlements (external:_automation_entitlements)`

Entitlements contains the entitlements needed for executing the function.

#### `function (string)`

Function contains the code.

#### `key (string)`

Key contains the unique identifier key for the template.

#### `kind (enum)`

Kind represents the kind of template.

| Characteristics | Value               |
| -               | -:                  |
| Allowed Value   | `Action, Condition` |
| Default         | `"Condition"`       |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `parameters (refMap)`

Parameters contains the parameter description of the function.

## AutomationTemplateParameter

Represents an Automation template parameter.

### Example

```json
{
  "name": "the name",
  "type": "String"
}
```

### Attributes

#### `allowedChoices (external:map_of_string_of_objects)`

Set the possible values for the parameter.

#### `defaultValue (object)`

Default value of the parameter.

#### `description (string)`

Name of the parameter.

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `position (integer)`

Preferred position for the parameter.

#### `required (boolean)`

Set if the parameter must be set.

#### `type (enum)`

Type of the parameter.

| Characteristics | Value                                                          |
| -               | -:                                                             |
| Allowed Value   | `Array, Boolean, Enum, Filter, Float, Integer, Object, String` |
| Required        | `true`                                                         |

## AWSAccount

Allows to bind an AWS account to your Aporeto account to allow auto registration
of enforcers running on EC2.

### Example

```json
{
  "accessKeyID": "aws access key id",
  "region": "us-west-2",
  "secretAccessKey": "aws secret access key"
}
```

### Relations

#### `GET /awsaccounts`

Retrieves the list of aws account bindings.

##### Parameters

- `accountid` (string): Id of the account.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /awsaccounts`

Creates a new aws account binding.

#### `DELETE /awsaccounts/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /awsaccounts/:id`

Retrieves the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `accessKeyID (string)`

AccessKeyID contains the aws access key ID. This is used to retrieve your
account id, and it is not stored.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `accessToken (string)`

AccessToken contains your aws access token. It is used to retrieve your account
id, and it not stored.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `accountID (string)`

accountID contains your verified accound id.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `parentID (string)`

ParentID contains the parent Vince account ID.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

#### `parentName (string)`

ParentName contains the name of the Vince parent Account.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

#### `region (string)`

Region contains your the region where your AWS account is located.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `secretAccessKey (string)`

secretAccessKey contains the secret key. It is used to retrieve your account id,
and it is not stored.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AWSAPIGateway

managed API decisions for the AWS API Gateway.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /awsapigateways`

create an AWS API Gateway.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /awsapigateways`

Manages the AWS API Gateway.

#### `DELETE /awsapigateways/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /awsapigateways/:id`

Retrieves the object with the given ID.

#### `PUT /awsapigateways/:id`

Updates the object with the given ID.

### Attributes

#### `APIID (string)`

API ID as defined on AWS for the API that handled this request.

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `accountID (string)`

the account ID for the gateway managing this request.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `authorized (boolean)`

The policy decision for this API flow.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `method (string)`

API method that handled this request.

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespaceID (string)`

Link to the cluster namespace where the AWS API gateway is defined.

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `resource (string)`

API resource that handled this request.

#### `sourceIP (string)`

the client ip for this request.

#### `stage (string)`

the stage name as defined on AWS for the API that handled this request.

#### `token (string)`

the JWT token that was optionally attached to this request.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## AWSRegister

This api allows AWS customer to register with Aporeto SaaS for billing.

### Relations

#### `POST /awsregister`

Creates a new aws registration for billing.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `provider (string)`

Token Provided by AWS.

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Category

Category allows to categorized services.

### Example

```json
{
  "name": "the name"
}
```

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

## Certificate

A User represents the owner of some certificates.

### Example

```json
{
  "commonName": "john doe",
  "email": "john@doe.com",
  "name": "john.doe"
}
```

### Relations

#### `GET /certificates`

Retrieves the list of existing user certificates.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /certificates`

Creates a new user certificate.

#### `DELETE /certificates/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /certificates/:id`

Retrieves the object with the given ID.

#### `PUT /certificates/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `admin (boolean)`

Admin determines if the certificate must be added to the admin list.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `commonName (string)`

CommonName (CN) for the user certificate.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `64`   |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `data (string)`

Certificate provides a certificate for the user.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `email (string)`

e-mail address of the user.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `expirationDate (time)`

CertificateExpirationDate indicates the expiration day for the certificate.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |

#### `key (string)`

CertificateKey provides the key for the user. Only available at create or update
time.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `name (string)`

Name of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `organizationalUnits (list)`

OrganizationalUnits attribute for the generated certificates.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `parentID (string)`

ParentID holds the parent account ID.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `passphrase (string)`

Passphrase to use for the generated p12.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |

#### `serialNumber (string)`

SerialNumber of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `status (enum)`

CertificateStatus provides the status of the certificate. Update with RENEW to
get a new certificate.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `Revoked, Valid` |
| Default         | `"Valid"`        |
| Orderable       | `true`           |
| Filterable      | `true`           |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## ClaimMapping

Represents a mapping from a claim name to an HTTP header.

### Example

```json
{
  "claimName": "email",
  "targetHTTPHeader": "X-Username"
}
```

### Attributes

#### `claimName (string)`

Claim name is the name of the claim that must be mapped to an HTTP header.

| Characteristics | Value                           |
| -               | -:                              |
| Format          | `/^[a-zA-Z0-9-_/*#&@\+\$~:]+$/` |
| Required        | `true`                          |

#### `targetHTTPHeader (string)`

The target HTTP header where this claim name must be mapped.

| Characteristics | Value                           |
| -               | -:                              |
| Format          | `/^[a-zA-Z0-9-_/*#&@\+\$~:]+$/` |
| Required        | `true`                          |

## Credential

Represents an application credential data.

### Attributes

#### `APIURL (string)`

The Aporeto API URL.

#### `ID (string)`

The ID of app credential.

#### `certificate (string)`

The certificate data encoded in base64.

#### `certificateAuthority (string)`

The certificate authority data encoded in base64.

#### `certificateKey (string)`

The certificate key data encoded in base64.

#### `name (string)`

The name of app credential.

#### `namespace (string)`

The namespace of app credential.

## DependencyMap

This api returns a data structure representing the graph of all processing units
and their connections in a particular namespace, in a given time window. To pass
the time window you can use the query parameters 'startAbsolute', 'endAbsolute',
'startRelative', 'endRelative'.

For example
  "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000".

### Relations

#### `GET /dependencymaps`

Retrieves the dependencymap of a namespace.

##### Parameters

- `tag` (string): Only show objects with the given tags in the dependency map.
- `view` (string): Set the view query for grouping the dependency map.
- `viewSuggestions` (boolean): Also return the view suggestions.
- `endAbsolute` (time): Set the absolute end of the time window.
- `endRelative` (duration): Set the relative end of the time window.
- `startAbsolute` (time): Set the absolute start of the time window.
- `startRelative` (duration): Set the relative start of the time window.
- `flowOffset` (duration): Apply an offset to the time window for flows.

##### Mandatory Parameters

(`endRelative`) or (`startRelative`) or (`startRelative` and `endRelative`) or (`startRelative` and `endAbsolute`) or (`startAbsolute` and `endRelative`) or (`startAbsolute` and `endAbsolute`)

### Attributes

#### `claims (external:map_of_string_of_list_of_strings)`

claims represents a user or a script that have accessed an api.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `edges (refMap)`

edges are the edges of the map.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `groups (refMap)`

Groups provide information about the group values.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `nodes (refMap)`

nodes refers to the nodes of the map.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `viewSuggestions (list)`

viewSuggestions provides suggestion of views based on relevant tags.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

## DeprecatedHostService

Represents a service of the enforcer's host.

### Attributes

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `name (string)`

Name of the service.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `networkonly (boolean)`

> This attribute is deprecated

networkonly indicate the host service is of type network only.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |
| Read only       | `true` |

#### `services (refList)`

Services lists all protocols and ports a service is running.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

## Endpoint

Represents an HTTP endpoint.

### Attributes

#### `URI (string)`

URI of the exposed API.

#### `allowedScopes (external:list_of_lists_of_strings)`

AllowedScopes authorized to access the API.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `methods (list)`

methods exposed to access the API.

#### `public (boolean)`

public defines if the api is public or not.

#### `scopes (list)`

> This attribute is deprecated

Scopes is deprecated.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

## Enforcer

An Enforcer contains all parameters associated with a registered enforcer. The
object is mainly by maintained by the enforcers themselves. Users can read the
object in order to understand the current status of the enforcers.

### Example

```json
{
  "FQDN": "server1.domain.com",
  "machineID": "3F23E8DF-C56D-45CF-89B8-A867F3956409",
  "name": "the name"
}
```

### Relations

#### `GET /enforcers`

Retrieves the list of enforcers.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /enforcers`

Creates a new enforcer.

#### `DELETE /enforcers/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Backward compat for enforcer <=v2.2.1. does not have any effect.
- `tag` (string): Backward compat for enforcer <=v2.2.1. does not have any effect.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /enforcers/:id`

Retrieves the object with the given ID.

#### `PUT /enforcers/:id`

Updates the object with the given ID.

#### `GET /auditprofilemappingpolicies/:id/enforcers`

Returns the list of enforcers that are affected by this poliy.

#### `GET /enforcerprofilemappingpolicies/:id/enforcers`

Returns the list of enforcers affected by an enforcer profile mapping policy.

#### `GET /hostservicemappingpolicies/:id/enforcers`

Returns the list of enforcers that are affected by this poliy.

#### `GET /enforcers/:id/auditprofiles`

Returns a list of the audit profiles that must be applied to this enforcer.

#### `GET /enforcers/:id/enforcerprofiles`

Returns the enforcer profile that must be used by an enforcer.

#### `GET /enforcers/:id/hostservices`

Returns a list of the host services policies that apply to this enforcer.

#### `GET /enforcers/:id/poke`

Sends a poke empty object. This is used to ensure an enforcer is up and running.

##### Parameters

- `cpuload` (float): If set, provides the total cpu usage in percentage of vCPUs.
- `memory` (integer): If set, provides the total resident memory used in bytes.
- `processes` (integer): If set, defines the number of current processes.
- `ts` (time): time of report. If not set, local server time will be used.

### Attributes

#### `FQDN (string)`

FQDN contains the fqdn of the server where the enforcer is running.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `certificate (string)`

Certificate is the certificate of the enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `certificateExpirationDate (time)`

CertificateExpirationDate is the expiration date of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `certificateKey (string)`

CertificateKey is the secret key of the enforcer. Returned only when a enforcer
is created or the certificate is updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `certificateRequest (string)`

CertificateRequest can be used to generate a certificate from that CSR instead
of letting the server generate your private key for you.

#### `certificateRequestEnabled (boolean)`

If set during creation,the server will not initially generate a certificate. In
that case you have to provide a valid CSR through certificateRequest during an
update.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `collectInfo (boolean)`

CollectInfo indicates to the enforcer it needs to collect information.

#### `collectedInfo (external:map_of_string_of_strings)`

CollectedInfo represents the latest info collected by the enforcer.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `currentVersion (string)`

CurrentVersion holds the enforcerd binary version that is currently associated
to this object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `enforcerProfileID (string)`

Contains the ID of the profile used by the instance of enforcerd.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `lastCollectionTime (time)`

LastCollectionTime represents the date and time when the info have been
collected.

#### `lastSyncTime (time)`

LastSyncTime holds the last heart beat time.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `localCA (string)`

LocalCA contains the initial chain of trust for the enforcer. This valud is only
given when you retrieve a single enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |

#### `machineID (string)`

MachineID holds a unique identifier for every machine as detected by the
enforcer. It is based on hardware information such as the SMBIOS UUID, MAC
addresses of interfaces or cloud provider IDs.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `operationalStatus (enum)`

OperationalStatus tells the status of the enforcer.

| Characteristics | Value                                                       |
| -               | -:                                                          |
| Allowed Value   | `Registered, Connected, Disconnected, Initialized, Unknown` |
| Default         | `"Registered"`                                              |
| Filterable      | `true`                                                      |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `publicToken (string)`

PublicToken is the public token of the server that will be included in the
datapath and its signed by the private CA.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `startTime (time)`

startTime holds the time this enforcerd was started. This time-stamp is reported
by the enforcer and is is preserved across disconnects.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateAvailable (boolean)`

Tells if the the version of enforcerd is outdated and should be updated.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## EnforcerProfile

Allows to create reusable configuration profile for your enforcers. Enforcer
Profiles contains various startup information that can (for some) be updated
live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
Mapping Policy.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /enforcerprofiles`

Retrieves the list of enforcer profiles.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /enforcerprofiles`

Creates a new enforcer profile.

#### `DELETE /enforcerprofiles/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /enforcerprofiles/:id`

Retrieves the object with the given ID.

#### `PUT /enforcerprofiles/:id`

Updates the object with the given ID.

#### `GET /enforcerprofilemappingpolicies/:id/enforcerprofiles`

Returns the list of enforcer profiles that an enforcer profile mapping policy
matches.

#### `GET /enforcers/:id/enforcerprofiles`

Returns the enforcer profile that must be used by an enforcer.

#### `GET /enforcerprofiles/:id/auditprofiles`

Returns the list of AuditProfiles used by an enforcer profile.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `IPTablesMarkValue (integer)`

IptablesMarkValue is the mark value to be used in an iptables implementation.

| Characteristics | Value   |
| -               | -:      |
| Default         | `1000`  |
| Max length      | `65000` |
| Orderable       | `true`  |

#### `PUBookkeepingInterval (string)`

PUBookkeepingInterval configures how often the PU will be synchronized.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `"15m"`           |
| Orderable       | `true`            |

#### `PUHeartbeatInterval (string)`

PUHeartbeatInterval configures the heart beat interval.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `"5s"`            |
| Orderable       | `true`            |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `applicationProxyPort (integer)`

Port used by aporeto application proxy.

| Characteristics | Value   |
| -               | -:      |
| Default         | `20992` |
| Min length      | `1`     |
| Max length      | `65535` |
| Orderable       | `true`  |

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `auditProfileSelectors (list)`

AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.

#### `auditProfiles (refList)`

AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `auditSocketBufferSize (integer)`

AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.

| Characteristics | Value    |
| -               | -:       |
| Default         | `16384`  |
| Max length      | `262144` |
| Orderable       | `true`   |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `dockerSocketAddress (string)`

DockerSocketAddress is the address of the docker daemon.

| Characteristics | Value                                                                                                                                       |
| -               | -:                                                                                                                                          |
| Format          | `/^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$/` |
| Default         | `"unix:///var/run/docker.sock"`                                                                                                             |
| Orderable       | `true`                                                                                                                                      |

#### `excludedInterfaces (list)`

ExcludedInterfaces is a list of interfaces that must be excluded.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `excludedNetworks (list)`

ExcludedNetworks is the list of networks that must be excluded for this
enforcer.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `hostModeEnabled (boolean)`

hostModeEnabled enables protection of the complete host. When this option is
turned on, all incoming and outgoing flows will be monitored. Flows will
be allowed if and only if a network policy has been created to allow the flow.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |

#### `hostServices (refList)`

HostServices is a list of services that must be activated by default to all
enforcers matching this profile.

#### `ignoreExpression (external:list_of_lists_of_strings)`

IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.

#### `killContainersOnFailure (boolean)`

KillContainersOnFailure will configure the enforcers to kill any containers if
there are policy failures.

#### `kubernetesMetadataExtractor (enum)`

Select which metadata extractor to use to process new processing units from
Kubernetes.

| Characteristics | Value                                  |
| -               | -:                                     |
| Allowed Value   | `KubeSquall, PodAtomic, PodContainers` |
| Default         | `"PodAtomic"`                          |
| Orderable       | `true`                                 |

#### `kubernetesSupportEnabled (boolean)`

KubernetesSupportEnabled enables kubernetes mode for the enforcer.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |

#### `linuxProcessesSupportEnabled (boolean)`

LinuxProcessesSupportEnabled configures support for Linux processes.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `metadataExtractor (enum)`

Select which metadata extractor to use to process new processing units.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Docker, ECS, Kubernetes` |
| Default         | `"Docker"`                |
| Orderable       | `true`                    |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `policySynchronizationInterval (string)`

PolicySynchronizationInterval configures how often the policy will be
resynchronized.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `"10m"`           |
| Orderable       | `true`            |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `proxyListenAddress (string)`

ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
example:
  unix:///var/run/aporeto.sock.

| Characteristics | Value                                                                                                                                       |
| -               | -:                                                                                                                                          |
| Format          | `/^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$/` |
| Default         | `"unix:///var/run/aporeto.sock"`                                                                                                            |
| Orderable       | `true`                                                                                                                                      |

#### `receiverNumberOfQueues (integer)`

ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |

#### `receiverQueue (integer)`

ReceiverQueue is the base queue number for traffic from the network.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1000` |
| Orderable       | `true` |

#### `receiverQueueSize (integer)`

ReceiverQueueSize is the queue size of the receiver.

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `5000` |
| Orderable       | `true` |

#### `remoteEnforcerEnabled (boolean)`

RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |
| Orderable       | `true` |

#### `targetNetworks (list)`

TargetNetworks is the list of networks that authorization should be applied.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `targetUDPNetworks (list)`

TargetUDPNetworks is the list of UDP networks that authorization should be
applied.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `transmitterNumberOfQueues (integer)`

TransmitterNumberOfQueues is the number of queues for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |

#### `transmitterQueue (integer)`

TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |

#### `transmitterQueueSize (integer)`

TransmitterQueueSize is the size of the queue for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |

#### `trustedCAs (list)`

List of trusted CA. If empty the main chain of trust will be used.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## EnforcerProfileMappingPolicy

A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used
by and Aporeto Agent based on the Enforcer that have been used during the
registration. The policy can also be propagated down to the child namespace.

### Example

```json
{
  "name": "the name",
  "object": [
    [
      "a=a",
      "b=b"
    ],
    [
      "c=c"
    ]
  ],
  "subject": [
    [
      "a=a",
      "b=b"
    ],
    [
      "c=c"
    ]
  ]
}
```

### Relations

#### `GET /enforcerprofilemappingpolicies`

Retrieves the list of enforcer profile mapping policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /enforcerprofilemappingpolicies`

Creates a new enforcer profile mapping policies.

#### `DELETE /enforcerprofilemappingpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /enforcerprofilemappingpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /enforcerprofilemappingpolicies/:id`

Updates the object with the given ID.

#### `GET /enforcerprofilemappingpolicies/:id/enforcerprofiles`

Returns the list of enforcer profiles that an enforcer profile mapping policy
matches.

#### `GET /enforcerprofilemappingpolicies/:id/enforcers`

Returns the list of enforcers affected by an enforcer profile mapping policy.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object is the list of tags to use to find a enforcer profile.

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject is the subject of the policy.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## EnforcerReport

Post a new enforcer statistics report.

### Example

```json
{
  "CPULoad": 10,
  "ID": "xxx-xxx-xxx-xxx",
  "memory": 10000,
  "name": "aporeto-enforcerd-xxx",
  "namespace": "/my/ns",
  "processes": 10,
  "timestamp": "2018-06-14T23:10:46.420397985Z"
}
```

### Relations

#### `POST /enforcerreports`

Create a enforcer statistics report.

### Attributes

#### `CPULoad (float)`

Total CPU utilization of the enforcer as a percentage of vCPUs.

#### `ID (string)`

ID of the enforcer to report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `memory (integer)`

Total resident memory used by the enforcer in bytes.

#### `name (string)`

Name of the enforcer to report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `namespace (string)`

Namespace of the enforcer to report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `processes (integer)`

Number of active processes of the enforcer.

#### `timestamp (time)`

Date of the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## EnforcerTraceReport

Post a new enforcer trace that determines how packets are.

### Example

```json
{
  "enforcerID": "xxx-xxx-xxx",
  "enforcerNamespace": "xxx-xxx-xxx",
  "namespace": "xxx-xxx-xxx",
  "puID": "xxx-xxx-xxx"
}
```

### Relations

#### `POST /enforcertracereports`

Create an enforcer trace report.

### Attributes

#### `enforcerID (string)`

EnforcerID of the enforcer where the trace was collected.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `enforcerNamespace (string)`

Namespace of the enforcer where the trace was collected.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `namespace (string)`

Namespace of the PU where the trace was collected.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `puID (string)`

ID of the pu where the trace was collected.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## EventLog

This api allows to report various event on any objects.

### Example

```json
{
  "category": "enforcerd:policy",
  "content": "Unable to activate docker container xyz because abc.",
  "targetID": "xxx-xxx-xxx-xxx",
  "targetIdentity": "processingunit",
  "title": "Error while activating processing unit."
}
```

### Relations

#### `POST /eventlogs`

Creates a new eventlog for a particular entity.

### Attributes

#### `category (string)`

Category of the log.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `content (string)`

Content of the log.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `date (time)`

Creation date of the eventlog.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Creation only   | `true` |
| Orderable       | `true` |

#### `level (enum)`

Represent the level of the log .

| Characteristics | Value                                   |
| -               | -:                                      |
| Allowed Value   | `Debug, Info, Warning, Error, Critical` |
| Default         | `"Info"`                                |
| Creation only   | `true`                                  |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |

#### `opaque (string)`

Opaque data that can attached to the log, for further machine processing.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |

#### `targetID (string)`

ID of the object this eventlog is attached to. The object must be in the same
namespace than the eventlog.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `targetIdentity (string)`

Identity of the object this eventlog is attached to.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `title (string)`

Title of the eventlog.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

## Export

Export the policies and related objects in a given namespace.

### Relations

#### `POST /export`

Exports all policies and related object of a namespace.

##### Parameters

- `ignoredTags` (string): List of tags to ignore from the export.

### Attributes

#### `APIVersion (integer)`

APIVersion of the api used for the exported data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `data (external:map_of_string_of_lists_of_maps_of_string_of_objects)`

List of all exported data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |

#### `identities (list)`

The list of identities to export.

#### `label (string)`

Label allows to define a unique label for this export. When importing the
content of the export, this label will be added as a tag that will be used to
recognize imported object in a later import.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## ExternalNetwork

An External Network represents a random network or ip that is not managed by the
system. They can be used in Network Access Policies in order to allow traffic
from or to the declared network or IP, using the provided protocol and port or
ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
as address, and 1-65000 for the ports. You will need to use the External
Services tags to set some policies.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /externalnetworks`

Retrieves the list of external network.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /externalnetworks`

Creates a new external network.

#### `DELETE /externalnetworks/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /externalnetworks/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /externalnetworks/:id`

Updates the object with the given ID.

#### `GET /networkaccesspolicies/:id/externalnetworks`

Returns the list of external networks affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `entries (list)`

List of CIDRs or domain name.

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `ports (list)`

List of single ports or range (xx:yy).

| Characteristics | Value                       |
| -               | -:                          |
| Default         | `[]interface {}{"1:65535"}` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protocols (list)`

List of protocols (tcp, udp, or protocol number).

| Characteristics | Value                   |
| -               | -:                      |
| Default         | `[]interface {}{"tcp"}` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## ExternalService

This API is deprecated in favor of externalnetworks.

### Example

```json
{
  "name": "the name",
  "network": "0.0.0.0/0",
  "protocol": "TCP"
}
```

### Relations

#### `GET /externalservices`

Retrieves the list of external services.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /externalservices`

Creates a new external service.

#### `DELETE /externalservices/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /externalservices/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /externalservices/:id`

Updates the object with the given ID.

#### `GET /networkaccesspolicies/:id/externalservices`

Returns the list of external services affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `network (string)`

Network refers to either CIDR or domain name.

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `port (string)`

Port refers to network port which could be a single number or 100:2000 to
represent a range of ports.

| Characteristics | Value                                                                                                                                                                                                            |
| -               | -:                                                                                                                                                                                                               |
| Format          | `/^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$/` |
| Default         | `"1:65535"`                                                                                                                                                                                                      |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protocol (string)`

Protocol refers to network protocol like TCP/UDP or the number of the protocol.

| Characteristics | Value                                                                  |
| -               | -:                                                                     |
| Format          | `/^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/` |
| Required        | `true`                                                                 |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## FileAccessPolicy

A File Access Policy allows Processing Units to access various folder and files.
It will use the tags of a File Path to know what is the path of the file or
folder to allow access to. You can allow the Processing Unit to have any
combination of read, write or execute.

When a Processing Unit is Docker container, then it will police the volumes
mount. executewon''t have any effect.

File path are not supported yet for standard Linux processes.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /fileaccesspolicies`

Retrieves the list of file access policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /fileaccesspolicies`

Creates a new file access policies.

#### `DELETE /fileaccesspolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /fileaccesspolicies/:id`

Retrieves the object with the given ID.

#### `PUT /fileaccesspolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `allowsExecute (boolean)`

AllowsExecute allows to execute the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `allowsRead (boolean)`

AllowsRead allows to read the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `allowsWrite (boolean)`

AllowsWrite allows to write the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `encryptionEnabled (boolean)`

EncryptionEnabled will enable the automatic encryption.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `logsEnabled (boolean)`

LogsEnabled will enable logging when this policy is used.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object is the object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject is the subject of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## FileAccessReport

Post a new file access statistics report.

### Example

```json
{
  "action": "Accepted",
  "processingUnitID": "xxx-xxx-xxx-xxx",
  "processingUnitNamespace": "/my/ns",
  "timestamp": "2018-06-14T23:10:46.420397985Z"
}
```

### Relations

#### `POST /fileaccessreports`

Create a file access statistics report.

### Attributes

#### `action (enum)`

Action taken.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `Accept, Reject` |
| Required        | `true`           |

#### `host (string)`

Host of the file.

| Characteristics | Value         |
| -               | -:            |
| Default         | `"localhost"` |
| Required        | `true`        |

#### `mode (string)`

Mode of the file access.

| Characteristics | Value   |
| -               | -:      |
| Default         | `"rxw"` |
| Required        | `true`  |

#### `path (string)`

Path of the file.

| Characteristics | Value           |
| -               | -:              |
| Default         | `"/etc/passwd"` |
| Required        | `true`          |

#### `processingUnitID (string)`

ID of the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `processingUnitNamespace (string)`

Namespace of the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `timestamp (time)`

Date of the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## FilePath

A File Path represents a random path to a file or a folder. They can be used in
aFile Access Policiesin order to allow Processing Units to access them, using
various modes (read, write, execute). You will need to use the File Paths tags
to set some policies. A good example would bevolume=web or file=/etc/passwd.

### Example

```json
{
  "filepath": "/etc/passwd",
  "name": "the name"
}
```

### Relations

#### `GET /filepaths`

Retrieves the list of file path.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /filepaths`

Create a new file path.

#### `DELETE /filepaths/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /filepaths/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /filepaths/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `filepath (string)`

FilePath refer to the file mount path.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `server (string)`

server is the server name/ID/IP associated with the file path.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## FlowReport

Post a new flow statistics report.

### Example

```json
{
  "action": "Accept",
  "destinationID": "xxx-xxx-xxx",
  "destinationNamespace": "/my/namespace",
  "destinationType": "ProcessingUnit",
  "namespace": "/my/namespace",
  "observedPolicyID": "xxx-xxx-xxx",
  "observedPolicyNamespace": "/my/namespace",
  "policyID": "xxx-xxx-xxx",
  "policyNamespace": "/my/namespace",
  "protocol": 6,
  "sourceID": "xxx-xxx-xxx",
  "sourceNamespace": "/my/namespace",
  "sourceType": "ProcessingUnit",
  "value": 1
}
```

### Relations

#### `POST /flowreports`

Create a flow statistics report.

### Attributes

#### `action (enum)`

Action applied to the flow.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `Accept, Reject` |
| Required        | `true`           |

#### `destinationID (string)`

ID of the destination.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `destinationIP (string)`

Type of the destination.

#### `destinationNamespace (string)`

Namespace of the receiver.

#### `destinationPort (integer)`

Port of the destination.

#### `destinationType (enum)`

Type of the source.

| Characteristics | Value                                     |
| -               | -:                                        |
| Allowed Value   | `ProcessingUnit, ExternalNetwork, Claims` |
| Required        | `true`                                    |

#### `dropReason (string)`

This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.

#### `encrypted (boolean)`

Tells is the flow has been encrypted.

#### `namespace (string)`

> This attribute is deprecated

This is here for backward compatibility.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `observed (boolean)`

Tells if the flow is from design mode.

#### `observedAction (enum)`

Action observed on the flow.

| Characteristics | Value                           |
| -               | -:                              |
| Allowed Value   | `Accept, Reject, NotApplicable` |
| Default         | `"NotApplicable"`               |

#### `observedDropReason (string)`

This field is only set if 'observedAction' is set to 'Reject' and specifies the
reason for the rejection.

#### `observedEncrypted (boolean)`

Value of the encryption of the network policy that observed the flow.

#### `observedPolicyID (string)`

ID of the network policy that observed the flow.

#### `observedPolicyNamespace (string)`

Namespace of the network policy that observed the flow.

#### `policyID (string)`

ID of the network policy that accepted the flow.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `policyNamespace (string)`

Namespace of the network policy that accepted the flow.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `protocol (integer)`

protocol number.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `serviceClaimHash (string)`

Hash of the claims used to communicate.

#### `serviceID (string)`

ID of the service.

#### `serviceNamespace (string)`

Service URL accessed.

#### `serviceType (enum)`

ID of the service.

| Characteristics | Value                          |
| -               | -:                             |
| Allowed Value   | `L3, HTTP, TCP, NotApplicable` |
| Default         | `"NotApplicable"`              |

#### `serviceURL (string)`

Service URL accessed.

#### `sourceID (string)`

ID of the source.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `sourceIP (string)`

Type of the source.

#### `sourceNamespace (string)`

Namespace of the receiver.

#### `sourceType (enum)`

Type of the source.

| Characteristics | Value                                     |
| -               | -:                                        |
| Allowed Value   | `ProcessingUnit, ExternalNetwork, Claims` |
| Required        | `true`                                    |

#### `timestamp (time)`

Date of the report.

#### `value (integer)`

Number of flows in the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## GraphEdge

Represents an edge from the dependency map.

### Attributes

#### `ID (string)`

Identifier of the edge.

#### `acceptedFlows (integer)`

Number of accepted flows in the edge.

#### `destinationID (string)`

ID of the destination GraphNode of the edge.

#### `destinationType (enum)`

Type of the destination GraphNode of the edge.

| Characteristics | Value                             |
| -               | -:                                |
| Allowed Value   | `ProcessingUnit, ExternalNetwork` |

#### `encrypted (integer)`

Tells the number of encrypted flows in the edge.

#### `observedAcceptedFlows (integer)`

Number of accepted observed flows.

#### `observedEncrypted (integer)`

Number of encrypted observed flows.

#### `observedPolicyIDs (refMap)`

Information about the observation policies that was hit in the flows
represented by that edge.

#### `observedRejectedFlows (integer)`

Number of rejected observed flows.

#### `observedServiceIDs (external:map_of_string_of_integers)`

Map of ints...

#### `policyIDs (refMap)`

Information about the policies that was hit in the flows represented by that
edge.

#### `rejectedFlows (integer)`

Number of rejected flows in the edge.

#### `serviceIDs (external:map_of_string_of_integers)`

Map of ints...

#### `sourceID (string)`

ID of the source GraphNode of the edge.

#### `sourceType (enum)`

Type of the source GraphNode of the edge.

| Characteristics | Value                             |
| -               | -:                                |
| Allowed Value   | `ProcessingUnit, ExternalNetwork` |

## GraphGroup

Represents an group of nodes from the dependency map.

### Attributes

#### `ID (string)`

Identifier of the group.

#### `color (string)`

Color to use for the group.

#### `match (external:list_of_lists_of_strings)`

List of tag that was used to create this group.

#### `name (string)`

Name of the group.

#### `parentID (string)`

ID of the parent group if any.

## GraphNode

Represents an node from the dependency map.

### Attributes

#### `ID (string)`

Identifier of object represented by the node.

#### `description (string)`

Description of object represented by the node.

#### `enforcementStatus (string)`

Enforcement status of processing unit represented by the node.

#### `groupID (string)`

ID of the group the node is eventually part of.

#### `lastUpdate (time)`

Last update of the node.

#### `name (string)`

Name of object represented by the node.

#### `namespace (string)`

Namespace of object represented by the node.

#### `status (string)`

Status of object represented by the node.

#### `tags (list)`

Tags of object represented by the node.

#### `type (enum)`

Type of object represented by the node.

| Characteristics | Value                                    |
| -               | -:                                       |
| Allowed Value   | `Docker, ExternalNetwork, Volume, Claim` |

## GraphPolicyInfo

Represents a policy information.

### Attributes

#### `count (integer)`

Number of time the policy has been hit.

#### `namespace (string)`

Namespace of the policy.

## HookPolicy

Hook allows to to define hooks to the write operations in squall. Hooks are sent
to an external Rufus server that will do the processing and eventually return a
modified version of the object before we save it.

### Example

```json
{
  "certificateAuthority": "-----BEGIN CERTIFICATE-----
MIIBbjCCARSgAwIBAgIRANRbvVzTzBZOvMCb8BiKCLowCgYIKoZIzj0EAwIwJjEN
MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNTE4
NDgwN1oXDTI3MTEyNDE4NDgwN1owJjENMAsGA1UEChMEQWNtZTEVMBMGA1UEAxMM
QWNtZSBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEJ/80HR51+vau
7XH7zS7b8ABA0e/TdBOg1NznbnXdXil1tDvWloWuH5+/bbaiEg54wksJHFXaukw8
jhTLU7zT56MjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSAAwRQIhALwAZh2KLFFC1qfb5CqFHExlXS0PUltax9PvQCN9P0vl
AiBl7/st9u/JpERjJgirxJxOgKNlV6pq9ti75EfQtZZcQA==
-----END CERTIFICATE-----",
  "clientCertificate": "-----BEGIN CERTIFICATE-----
MIIBczCCARigAwIBAgIRALD3Vz81Pq10g7n4eAkOsCYwCgYIKoZIzj0EAwIwJjEN
MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNzA2
NTM1MloXDTI3MTEyNjA2NTM1MlowGDEWMBQGA1UEAxMNY2xhaXJlLWNsaWVudDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABOmzPJj+t25T148eQH5gVrZ7nHwckF5O
evJQ3CjSEMesjZ/u7cW8IBfXlxZKHxl91IEbbB3svci4c8pycUNZ2kujNTAzMA4G
A1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
MAoGCCqGSM49BAMCA0kAMEYCIQCjAAmkQpTua0HR4q6jnePaFBp/JMXwTXTxzbV6
peGbBQIhAP+1OR8GFnn2PlacwHqWXHwkvy6CLPVikvgtwEdB6jH8
-----END CERTIFICATE-----",
  "clientCertificateKey": "-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
F9eXFkofGX3UgRtsHe123456789xQ1naSw==
-----END EC PRIVATE KEY-----",
  "endpoint": "https://hooks.hookserver.com/remoteprocessors",
  "name": "the name",
  "subject": [
    [
      "$identity=processingunit"
    ]
  ]
}
```

### Relations

#### `GET /hookpolicies`

Retrieves the list of hook policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /hookpolicies`

Creates a new hook policy.

#### `DELETE /hookpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /hookpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /hookpolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `certificateAuthority (string)`

CertificateAuthority contains the pem block of the certificate authority used by
the remote endpoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

#### `clientCertificate (string)`

ClientCertificate contains the client certificate that will be used to connect
to the remote endoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

#### `clientCertificateKey (string)`

ClientCertificateKey contains the key associated to the clientCertificate.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

#### `continueOnError (boolean)`

If set to true and `mode` is in `Pre`, the request will be honored even if
calling the hook fails.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `endpoint (string)`

Endpoint contains the full address of the remote processor endoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `mode (enum)`

Mode define the type of the hook.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Both, Post, Pre` |
| Default         | `"Pre"`           |
| Orderable       | `true`            |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject contains the tag expression that an object must match in order to
trigger the hook.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## HostService

Represents a set of services that a host must expose and protect.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /hostservices`

Retrieves the list of host services.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /hostservices`

Creates a new host service.

#### `DELETE /hostservices/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /hostservices/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /hostservices/:id`

Updates the object with the given ID.

#### `GET /enforcers/:id/hostservices`

Returns a list of the host services policies that apply to this enforcer.

#### `GET /hostservicemappingpolicies/:id/hostservices`

Returns the list of host services that are referred by this policy.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `hostModeEnabled (boolean)`

HostModeEnabled forces the corresponding enforcers to enable complete host
protection. When this option is turned on, all incoming and outgoing flows will
be monitored. Flows will be allowed if and only if a network policy has been
created to allow the flow. The option applies to all enforcers that match the
subject constraints.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `services (list)`

Services lists all protocols and ports a service is running. A service entry can
be defined by a protocol and port '(tcp/80)', or range of protocol/port pairs
'(udp/80:100)'. If no protocol is provided, it is assumed to be TCP. Allowed
protocols are only tcp and udp.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## HostServiceMappingPolicy

Defines a host service mapping policy that provides the relation between
enforcers and host services that they must implement.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /hostservicemappingpolicies`

Retrieves the list of host service mapping policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /hostservicemappingpolicies`

Creates a new host service mapping policy.

#### `DELETE /hostservicemappingpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /hostservicemappingpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /hostservicemappingpolicies/:id`

Updates the object with the given ID.

#### `GET /hostservicemappingpolicies/:id/enforcers`

Returns the list of enforcers that are affected by this poliy.

#### `GET /hostservicemappingpolicies/:id/hostservices`

Returns the list of host services that are referred by this policy.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object of the policy is the selector for the host services that must be applied
to this enforcer.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject of the policy is the selector of the enforcers that the list of host
services must apply to.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## HTTPResourceSpec

HTTPResourceSpec descibes an HTTP resource exposed by a service. These APIs
can be associated with one or more services.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /httpresourcespecs`

Retrieves the list of HTTP Resource specifications.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /httpresourcespecs`

Creates a new HTTP Resource specification.

#### `DELETE /httpresourcespecs/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /httpresourcespecs/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /httpresourcespecs/:id`

Updates the object with the given ID.

#### `GET /services/:id/httpresourcespecs`

Retrieves the HTTP Resource exposed by this service.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `endpoints (refList)`

EndPoints is a list of API endpoints that are exposed for the service.

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Import

Imports an export of policies and related objects into the namespace.

### Example

```json
{
  "data": "previous output of export"
}
```

### Relations

#### `POST /import`

Imports data from a previous export.

### Attributes

#### `data (ref)`

The data to import.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `mode (enum)`

How to import the data. ReplacePartial is deprecated and should be Import. Right
now the API considers both equivalent.

| Characteristics | Value                            |
| -               | -:                               |
| Allowed Value   | `ReplacePartial, Import, Remove` |
| Default         | `"Import"`                       |

## InstalledApp

InstalledApps represents an installed application.

### Relations

#### `GET /installedapps`

Retrieves the list of installed apps.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /installedapps`

Installs a new app.

#### `DELETE /installedapps/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /installedapps/:id`

Retrieves the object with the given ID.

#### `PUT /installedapps/:id`

Updates the object with the given ID.

#### `GET /installedapps/:id/logs`

Returns the logs for a app.

### Attributes

#### `ID (string)`

ID of the installed app.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `accountName (string)`

AccountName represents the vince account name.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `categoryID (string)`

CategoryID of the app.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

#### `currentVersion (string)`

Version of the installed app.

#### `name (string)`

Name of the installed app.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |

#### `namespace (string)`

Namespace in which the app is running.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

#### `parameters (refList)`

Parameters is a list of parameters to start the app.

#### `status (enum)`

Status of the app.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Error, Pending, Running` |
| Default         | `"Pending"`               |
| Read only       | `true`                    |
| Orderable       | `true`                    |

## Invoice

This api allows to view invoices for Aporeto customers.

### Relations

#### `DELETE /invoices/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /invoices/:id`

Retrieves the object with the given ID.

#### `PUT /invoices/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the id of the invoice.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `accountID (string)`

AccountID references the id of the customer that this invoice belongs to.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `billedToProvider (enum)`

BilledToProvider holds the name of the provider that this invoice was billed to.

| Characteristics | Value          |
| -               | -:             |
| Allowed Value   | `Aporeto, AWS` |
| Default         | `"Aporeto"`    |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `endDate (time)`

EndDate holds the end date for this invoice.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `startDate (time)`

StartDate holds the start date for this invoice.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## InvoiceRecord

This api allows to view detailed records of invoices for Aporeto customers.

### Relations

#### `DELETE /invoicerecords/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /invoicerecords/:id`

Retrieves the object with the given ID.

#### `PUT /invoicerecords/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the id of this invoice record.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `invoiceID (string)`

InvoiceID references the id of the invoice that this invoice record provides
details for.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `invoiceRecords (list)`

InvoiceRecords provides details about billing units.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## IPInfo

This apis allows to resolve information from an IP address.

### Relations

#### `GET /ipinfos`

Returns information about an IP address given as parameters.

##### Parameters

- `ip` (string): List of IPs to resolve.

##### Mandatory Parameters

`ip`

### Attributes

#### `IP (string)`

The IP resolved.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `error (string)`

Eventual error that happened during resolution.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `records (external:map_of_string_of_strings)`

List of DNS records associated to that IP.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## IsolationProfile

An IsolationProfile needs documentation.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /isolationprofiles`

Retrieves the list of isolation profiles.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /isolationprofiles`

Creates a new isolation profile.

#### `DELETE /isolationprofiles/:id`

Deletes the object with the given ID.

#### `GET /isolationprofiles/:id`

Retrieves the object with the given ID.

#### `PUT /isolationprofiles/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `capabilitiesActions (external:_cap_map)`

CapabilitiesActions identifies the capabilities that should be added or removed
from the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `defaultSyscallAction (external:_syscall_action)`

DefaultAction is the default action applied to all syscalls of this profile.
Default is "Allow".

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `syscallRules (external:_syscall_rules)`

SyscallRules is a list of syscall rules that identify actions for particular
syscalls.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `targetArchitectures (external:_arch_list)`

TargetArchitectures is the target processor architectures where this profile can
be applied. Default all.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Issue

This API issues a new token according to given data.

### Example

```json
{
  "realm": "Certificate"
}
```

### Relations

#### `POST /issue`

Issues a new token.

##### Parameters

- `token` (string): Token to verify.

### Attributes

#### `audience (string)`

If given, the issued token will only be valid from that namespace declared in
that value.

#### `data (string)`

Data contains additional data. The value depends on the issuer type.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (external:map_of_string_of_objects)`

Metadata contains various additional information. Meaning depends on the realm.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `opaque (external:map_of_string_of_strings)`

Opaque data that will be included in the issued token.

#### `quota (integer)`

Restricts the number of time the issued token should be used.

#### `realm (enum)`

Realm is the authentication realm.

| Characteristics | Value                                                                                                                 |
| -               | -:                                                                                                                    |
| Allowed Value   | `AWSIdentityDocument, AWSSecurityToken, Certificate, Google, LDAP, Vince, GCPIdentityToken, AzureIdentityToken, OIDC` |
| Required        | `true`                                                                                                                |

#### `token (string)`

Token is the token to use for the registration.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `validity (string)`

Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.

| Characteristics | Value                                                                                                             |
| -               | -:                                                                                                                |
| Format          | `/^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$/` |
| Default         | `"24h"`                                                                                                           |
| Orderable       | `true`                                                                                                            |

## Jaegerbatch

A jaegerbatch is a batch of jaeger spans. This is used by external service to
post jaeger span in our private jaeger services.

### Relations

#### `POST /jaegerbatchs`

Sends a jaeger tracing batch.

### Attributes

#### `batch (object)`

Represents a jaeger batch.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

## K8SCluster

Create a remote Kubernetes Cluster integration.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /k8sclusters`

Retrieves the list of kubernetes clusters.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /k8sclusters`

Creates a new kubernetes cluster.

#### `DELETE /k8sclusters/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /k8sclusters/:id`

Retrieves the object with the given ID.

#### `PUT /k8sclusters/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activationType (enum)`

Defines the mode of activation on the KubernetesCluster.

| Characteristics | Value                                  |
| -               | -:                                     |
| Allowed Value   | `KubeSquall, PodAtomic, PodContainers` |
| Default         | `"KubeSquall"`                         |
| Orderable       | `true`                                 |

#### `adminEmail (string)`

The email address that will receive a copy of the Kubernetes cluster YAMLs
definition.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `certificate (string)`

The string representation of the Certificate used by the Kubernetes cluster.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `kubernetesDefinitions (string)`

base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `networkPolicyType (enum)`

Defines what type of network policy will be applied on your cluster.
Kubernetes means that All the Kubernetes policies will be synced to Squall.
No Policies means that policies are not synced and it's up to the user to create
consistent policies in Squall.

| Characteristics | Value                  |
| -               | -:                     |
| Allowed Value   | `Kubernetes, NoPolicy` |
| Default         | `"Kubernetes"`         |
| Orderable       | `true`                 |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `regenerate (boolean)`

Regenerates the k8s files and certificates.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Log

Retrieves the log of a deployed app.

### Relations

#### `GET /installedapps/:id/logs`

Returns the logs for a app.

### Attributes

#### `data (external:map_of_string_of_strings)`

Data contains all logs data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## Message

The Message API allows to post public messages that will be visible through all
children namespaces.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /messages`

Retrieves the list of messages.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /messages`

Creates a new message.

#### `DELETE /messages/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /messages/:id`

Retrieves the object with the given ID.

#### `PUT /messages/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `expirationTime (time)`

expirationTime is the time after which the message will be deleted.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `level (enum)`

Level defines how the message is important.

| Characteristics | Value                   |
| -               | -:                      |
| Allowed Value   | `Danger, Info, Warning` |
| Default         | `"Info"`                |
| Orderable       | `true`                  |

#### `local (boolean)`

If local is set, the message will only be visible in the current namespace.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `notifyByEmail (boolean)`

If enabled, the message will be sent to the email associated in namespaces
annotations.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `validity (string)`

Validity set using golang time duration, when the message will be automatically
deleted.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

## Namespace

A Namespace represents the core organizational unit of the system. All objects
always exists in a single namespace. A Namespace can also have child namespaces.
They can be used to split the system into organizations, business units,
applications, services or any combination you like.

### Example

```json
{
  "name": "mynamespace"
}
```

### Relations

#### `GET /namespaces`

Retrieves the list of namespaces.

##### Parameters

- `authorized` (boolean): Returns all namespaces the token bearer has the right to read. If set, other parameters like `recursive` or `q` will have no effect.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /namespaces`

Creates a new namespace.

#### `DELETE /namespaces/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /namespaces/:id`

Retrieves the object with the given ID.

#### `PUT /namespaces/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `localCA (string)`

LocalCA holds the eventual certificate authority used by this namespace.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `localCAEnabled (boolean)`

LocalCAEnabled defines if the namespace should use a local Certificate
Authority. Switching it off and on again will regenerate a new CA.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the namespace.

| Characteristics | Value                 |
| -               | -:                    |
| Format          | `/^[a-zA-Z0-9-_/]+$/` |
| Required        | `true`                |
| Creation only   | `true`                |
| Orderable       | `true`                |
| Filterable      | `true`                |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `networkAccessPolicyTags (list)`

List of tags that will be added to every `or` clause of all network access
policies in the namespace and its children.

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `serviceCertificateValidity (string)`

Determines the validity time of certificates issued in this namespace. Default
value is 1 hour.

| Characteristics | Value  |
| -               | -:     |
| Default         | `"1h"` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## NamespaceMappingPolicy

A Namespace Mapping Policy defines in which namespace a Processing Unit should
be placed when it is created, based on its tags.  When an Aporeto Agent creates
a new Processing Unit, the system will place it in its own namespace if no
matching Namespace Mapping Policy can be found. If one match is found, then the
Processing will be bumped down to the namespace declared in the policy. If it
finds in that child namespace another matching Namespace Mapping Policy, then
the Processing Unit will be bumped down again, until it reach a namespace with
no matching policies.  This is very useful to dispatch processes and containers
into a particular namespace, based on a lot of factor.   You can put in place a
quarantine namespace that will grab all Processing Units with too much
vulnerabilities for instances.

### Example

```json
{
  "mappedNamespace": "/blue/namespace",
  "name": "the name",
  "subject": [
    [
      "color=blue"
    ]
  ]
}
```

### Relations

#### `GET /namespacemappingpolicies`

Retrieves the list namespace mapping policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /namespacemappingpolicies`

Creates a new namespace mapping policy.

#### `DELETE /namespacemappingpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /namespacemappingpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /namespacemappingpolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `mappedNamespace (string)`

mappedNamespace is the mapped namespace.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject is the subject.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## NetworkAccessPolicy

Allows to define networking policies to allow or prevent processing units
identitied by their tags to talk to other processing units or external services
(also identified by their tags).

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /networkaccesspolicies`

Retrieves the list of network access policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /networkaccesspolicies`

Creates a new network access policy.

#### `DELETE /networkaccesspolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /networkaccesspolicies/:id`

Retrieves the object with the given ID.

##### Parameters

- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `PUT /networkaccesspolicies/:id`

Updates the object with the given ID.

#### `GET /networkaccesspolicies/:id/externalnetworks`

Returns the list of external networks affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /networkaccesspolicies/:id/externalservices`

Returns the list of external services affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /networkaccesspolicies/:id/processingunits`

Returns the list of Processing Units affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /networkaccesspolicies/:id/services`

Returns the list of services affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `action (enum)`

Action defines the action to apply to a flow.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Allow, Reject, Continue` |
| Default         | `"Allow"`                 |
| Orderable       | `true`                    |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `applyPolicyMode (enum)`

applyPolicyMode determines if the policy has to be applied to the
outgoing traffic of a PU or the incoming traffic of a PU or in both directions.
Default is both directions.

| Characteristics | Value                                             |
| -               | -:                                                |
| Allowed Value   | `OutgoingTraffic, IncomingTraffic, Bidirectional` |
| Default         | `"Bidirectional"`                                 |
| Orderable       | `true`                                            |

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `encryptionEnabled (boolean)`

EncryptionEnabled defines if the flow has to be encrypted.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `logsEnabled (boolean)`

LogsEnabled defines if the flow has to be logged.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `negateObject (boolean)`

Setting this to true will invert the object to find what is not matching.

#### `negateSubject (boolean)`

Setting this to true will invert the subject to find what is not matching.

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `observationEnabled (boolean)`

If set to true, the flow will be in observation mode.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `observedTrafficAction (enum)`

If observationEnabled is set to true, this will be the final action taken on the
packets.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Apply, Continue` |
| Default         | `"Continue"`      |
| Orderable       | `true`            |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## OIDCProvider

Allows to declare a generic OpenID Connect provider that can be used in exchange
for a Midgard token.

### Example

```json
{
  "clientID": "xxx-xxx-xxx",
  "clientSecret": "xxx-xxx-xxx",
  "endpoint": "https://accounts.google.com",
  "name": "okta",
  "scopes": [
    "email",
    "profile"
  ]
}
```

### Relations

#### `GET /oidcproviders`

Retrieves the list of the account OIDC provider.

##### Parameters

- `OIDCProviderName` (string): OIDCProvider name.
- `accountID` (string): Aporeto Account ID.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /oidcproviders`

Creates a new OIDC provider.

#### `DELETE /oidcproviders/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /oidcproviders/:id`

Retrieves the object with the given ID.

#### `PUT /oidcproviders/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `clientID (string)`

Unique client ID.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `clientSecret (string)`

Client secret associated with the client ID.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `endpoint (string)`

OIDC information endpoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `name (string)`

Name of the provider.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `parentID (string)`

ParentID contains the parent Vince account ID.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

#### `parentName (string)`

ParentName contains the name of the Vince parent Account.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

#### `scopes (list)`

List of scopes to allow.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## PacketReport

Post a new packet tracing report.

### Example

```json
{
  "destinationPort": 11000,
  "event": "Rcv",
  "length": 94,
  "mark": 123123,
  "namespace": "/my/namespace",
  "packetID": 12333,
  "protocol": 6,
  "puID": "xxx-xxx-xxx",
  "sourcePort": 80,
  "timestamp": "2018-06-14T23:10:46.420397985Z"
}
```

### Relations

#### `POST /packetreports`

Create a packet trace report.

### Attributes

#### `destinationIP (string)`

DestinationIP is the IP address of the destination.

#### `destinationPort (integer)`

DestinationPort is the destination port of a TCP or UDP packet.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65536` |

#### `dropReason (string)`

This field is only set if 'event' is set to 'Dropped' and specifies the reason
for the drop.

#### `encrypt (boolean)`

Encrypt indicates that the packet was encrypted.

#### `event (enum)`

Event is the event that triggered the report.

| Characteristics | Value                            |
| -               | -:                               |
| Allowed Value   | `Received, Transmitted, Dropped` |
| Required        | `true`                           |

#### `namespace (string)`

Namespace of the PU reporting the packet.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

#### `protocol (integer)`

Protocol number.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `255`  |
| Required        | `true` |

#### `puID (string)`

PUID is the ID of the PU reporting the packet.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

#### `sourceIP (string)`

SourceIP is the source IP address of the packet.

#### `sourcePort (integer)`

SourcePort is the source port of the packet.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65536` |

#### `timestamp (time)`

Timestamp is the date of the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `triremePacket (boolean)`

TriremePacket is set if the packet arrived with the Trireme options.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |
| Required        | `true` |

## PasswordReset

Used to reset an account password.

### Example

```json
{
  "password": "NewPassword123@",
  "token": "valid.jwt.token"
}
```

### Relations

#### `GET /passwordreset`

Sends a link to the account email to reset the password.

##### Parameters

- `email` (string): Email associated to the account.

##### Mandatory Parameters

`email`

#### `POST /passwordreset`

Resets the password for an account using the provided link.

### Attributes

#### `password (string)`

Password contains the new password.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `token (string)`

Token contains the reset password token.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## Plan

Plan contains the various billing plans available.

### Relations

#### `GET /plans`

Retrieves the list of plans.

#### `GET /plans/:id`

Retrieves the object with the given ID.

### Attributes

#### `description (string)`

Description contains the description of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `enforcersQuota (integer)`

EnforcerQuota contains the maximum number of enforcers available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `key (string)`

Key contains the key identifier of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `name (string)`

Name contains the name of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `policiesQuota (integer)`

PoliciesQuota contains the maximum number of policies available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `processingUnitsQuota (integer)`

ProcessingUnitsQuota contains the maximum PUs available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## Poke

When available, poke can be used to update various information about the parent.
For instance, for enforcers, poke will be use as the heartbeat.

### Relations

#### `GET /enforcers/:id/poke`

Sends a poke empty object. This is used to ensure an enforcer is up and running.

##### Parameters

- `cpuload` (float): If set, provides the total cpu usage in percentage of vCPUs.
- `memory` (integer): If set, provides the total resident memory used in bytes.
- `processes` (integer): If set, defines the number of current processes.
- `ts` (time): time of report. If not set, local server time will be used.

#### `GET /processingunits/:id/poke`

Sends a poke empty object. This will send a snaphot of the pu to time series
database.

##### Parameters

- `enforcementStatus` (enum): If set, changes the enforcement status of the processing unit alongside with the poke.
- `status` (enum): If set, changes the status of the processing unit alongside with the poke.
- `ts` (time): time of report. If not set, local server time will be used.

## Policy

Policy represents the policy primitive used by all aporeto policies.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /policies`

Retrieves the list of policy primitives.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `DELETE /policies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /policies/:id`

Retrieves the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `action (external:map_of_string_of_maps_of_string_of_objects)`

Action defines set of actions that must be enforced when a dependency is met.

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object represents set of entities that another entity depends on. As subjects,
objects are identified as logical operations on tags when a policy is defined.

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `relation (list)`

Relation describes the required operation to be performed between subjects and
objects.

#### `subject (external:list_of_lists_of_strings)`

Subject represent sets of entities that will have a dependency other entities.
Subjects are defined as logical operations on tags. Logical operations can
includes AND/OR.

#### `type (enum)`

Type of the policy.

| Characteristics | Value                                                                                                                                                                                       |
| -               | -:                                                                                                                                                                                          |
| Allowed Value   | `APIAuthorization, AuditProfileMapping, EnforcerProfile, File, Hook, HostServiceMapping, NamespaceMapping, Network, ProcessingUnit, Quota, Service, ServiceDependency, Syscall, TokenScope` |
| Creation only   | `true`                                                                                                                                                                                      |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## PolicyRefresh

PolicyRefresh is sent to client when as a push event when a policy refresh is
needed on their side.

### Attributes

#### `sourceNamespace (string)`

SourceNamespace contains the original namespace of the updated object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `type (string)`

Type contains the policy type that is affected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

## PolicyRule

PolicyRule is an internal policy resolution API. Services can use this API to
retrieve a policy resolution.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /policyrules/:id`

Retrieves the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `action (external:map_of_string_of_maps_of_string_of_objects)`

Action defines set of actions that must be enforced when a dependency is met.

#### `auditProfiles (refList)`

AuditProfiles provides the audit profiles that must be applied.

#### `enforcerProfiles (refList)`

EnforcerProfiles provides the information about the server profile.

#### `externalNetworks (refList)`

Policy target networks.

#### `externalServices (refList)`

> This attribute is deprecated

Policy target networks.

#### `filePaths (refList)`

Policy target file paths.

#### `hostServices (refList)`

HostServices provides the list of host services that must be instantiated.

#### `isolationProfiles (refList)`

IsolationProfiles are the isolation profiles of the rule.

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespaces (refList)`

Policy target namespaces.

#### `policyNamespace (string)`

PolicyNamespace is the namespace of the policy that created this rule.

#### `policyUpdateTime (time)`

Last time the policy was updated.

#### `propagated (boolean)`

Propagated indicates if the policy is propagated.

#### `relation (list)`

Relation describes the required operation to be performed between subjects and
objects.

#### `services (refList)`

Services provides the services of this policy rule.

#### `tagClauses (external:list_of_lists_of_strings)`

Policy target tags.

## ProcessingUnit

A Processing Unit reprents anything that can compute. It can be a Docker
container, or a simple Unix process. They are created, updated and deleted by
the system as they come and go. You can only modify its tags.  Processing Units
use Network Access Policies to define which other Processing Units or External
Services they can communicate with and File Access Policies to define what File
Paths they can use.

### Example

```json
{
  "name": "the name",
  "type": "Docker"
}
```

### Relations

#### `GET /processingunits`

Retrieves the list of processing units.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /processingunits`

Creates a new processing unit.

#### `DELETE /processingunits/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /processingunits/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /processingunits/:id`

Updates the object with the given ID.

#### `GET /networkaccesspolicies/:id/processingunits`

Returns the list of Processing Units affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /servicedependencies/:id/processingunits`

Returns the list of Processing Units that depend on an service.

#### `GET /services/:id/processingunits`

Retrieves the Processing Units that implement this service.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /vulnerabilities/:id/processingunits`

Retrieves the processing units affected by the a vulnerabily.

#### `GET /processingunits/:id/poke`

Sends a poke empty object. This will send a snaphot of the pu to time series
database.

##### Parameters

- `enforcementStatus` (enum): If set, changes the enforcement status of the processing unit alongside with the poke.
- `status` (enum): If set, changes the status of the processing unit alongside with the poke.
- `ts` (time): time of report. If not set, local server time will be used.

#### `GET /processingunits/:id/renderedpolicies`

Retrieves the policies for the processing unit.

##### Parameters

- `csr` (string): CSR to sign.

#### `GET /processingunits/:id/services`

Retrieves the services used by a processing unit.

#### `GET /processingunits/:id/vulnerabilities`

Retrieves the vulnerabilities affecting the processing unit.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `collectInfo (boolean)`

CollectInfo indicates to the enforcer it needs to collect information for this
PU.

#### `collectedInfo (external:map_of_string_of_strings)`

CollectedInfo represents the latest info collected by the enforcer for this PU.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `enforcementStatus (enum)`

EnforcementStatus communicates the state of the enforcer for that PU.

| Characteristics | Value                      |
| -               | -:                         |
| Allowed Value   | `Active, Failed, Inactive` |
| Default         | `"Inactive"`               |
| Filterable      | `true`                     |

#### `enforcerID (string)`

EnforcerID is the ID of the enforcer associated with the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `enforcerNamespace (string)`

enforcerNamespace is the namespace of the enforcer associated with the
processing unit.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `image (string)`

Docker image, or path to executable.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `lastSyncTime (time)`

LastSyncTime is the time when the policy was last resolved.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `nativeContextID (string)`

NativeContextID is the Docker UUID or service PID.

#### `networkServices (refList)`

NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `operationalStatus (enum)`

OperationalStatus of the processing unit.

| Characteristics | Value                                               |
| -               | -:                                                  |
| Allowed Value   | `Initialized, Paused, Running, Stopped, Terminated` |
| Default         | `"Initialized"`                                     |
| Filterable      | `true`                                              |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `tracing (ref)`

Tracing indicates if this PU must be placed in tracing mode.

#### `type (enum)`

Type of the container ecosystem.

| Characteristics | Value                                         |
| -               | -:                                            |
| Allowed Value   | `Docker, LinuxService, RKT, User, APIGateway` |
| Required        | `true`                                        |
| Creation only   | `true`                                        |
| Filterable      | `true`                                        |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## ProcessingUnitPolicy

A ProcessingUnitPolicies needs a better description.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /processingunitpolicies`

Retrieves the list of processing unit policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /processingunitpolicies`

Creates a new processing unit policy.

#### `DELETE /processingunitpolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /processingunitpolicies/:id`

Retrieves the object with the given ID.

#### `PUT /processingunitpolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `action (enum)`

Action determines the action to take while enforcing the isolation profile.

| Characteristics | Value                                                    |
| -               | -:                                                       |
| Allowed Value   | `Delete, Enforce, LogCompliance, Reject, Snapshot, Stop` |
| Orderable       | `true`                                                   |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `isolationProfileSelector (external:list_of_lists_of_strings)`

IsolationProfileSelector are the profiles that must be applied when this policy
matches. Only applies to Enforce and LogCompliance actions.

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject defines the tag selectors that identitfy the processing units to which
this policy applies.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## ProcessingUnitService

Represents a service attached to a processing unit.

### Attributes

#### `ports (string)`

> This attribute is deprecated

ports contains the list of allowed ports and ranges.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |

#### `protocol (integer)`

Protocol used by the service.

#### `targetPorts (list)`

List of single ports or range (xx:yy).

## QuotaPolicy

Quotas Policies allows to set quotas on the number of objects that can be
created in a namespace.

### Example

```json
{
  "identities": [
    "processingunit",
    "enforcer"
  ],
  "name": "the name",
  "targetNamespace": "/my/namespace"
}
```

### Relations

#### `GET /quotapolicies`

Retrieves the list of quota policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /quotapolicies`

Creates a new quota policy.

#### `DELETE /quotapolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /quotapolicies/:id`

Retrieves the object with the given ID.

#### `PUT /quotapolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `identities (list)`

Identities contains the list of identity names where the quota will be applied.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `quota (integer)`

Quota contains the maximum number of object matching the policy subject that can
be created.

#### `targetNamespace (string)`

TargetNamespace contains the base namespace from where the count will be done.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## RemoteProcessor

Hook to integrate an Aporeto service.

### Example

```json
{
  "claims": [
    "@auth:realm=certificate",
    "@auth:commonname=john"
  ],
  "input": "{
  \"name\": \"hello\",
  \"description\": \"hello\",
}",
  "mode": "Pre",
  "namespace": "/my/namespace",
  "operation": "create",
  "targetIdentity": "processingunit"
}
```

### Relations

#### `POST /remoteprocessors`

This should be be here.

### Attributes

#### `claims (list)`

Represents the claims of the currently managed object.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `input (external:raw_json)`

Represents data received from the service.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `mode (enum)`

Node defines the type of the hook.

| Characteristics | Value       |
| -               | -:          |
| Allowed Value   | `Post, Pre` |

#### `namespace (string)`

Represents the current namespace.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `operation (external:elemental_operation)`

Define the operation that is currently handled by the service.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `output (external:_elemental_identifiable)`

Returns the OutputData filled with the processor information.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `requestID (string)`

RequestID gives the id of the request coming from the main server.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `targetIdentity (string)`

Represents the Identity name of the managed object.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## RenderedPolicy

Retrieve the aggregated policies applied to a particular processing unit.

### Example

```json
{
  "processingUnit": "{
  \"name\": \"pu\",
  \"type\": \"Docker\",
  \"normalizedTags\": [
    \"a=a\",
    \"b=b\"
  ]
}"
}
```

### Relations

#### `POST /renderedpolicies`

Render a policy for a processing unit.

##### Parameters

- `csr` (string): CSR to sign.

#### `GET /processingunits/:id/renderedpolicies`

Retrieves the policies for the processing unit.

##### Parameters

- `csr` (string): CSR to sign.

### Attributes

#### `certificate (string)`

Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `dependendServices (refList)`

DependendServices is the list of services that this processing unit depends on.

#### `egressPolicies (external:_rendered_policy)`

EgressPolicies lists all the egress policies attached to processing unit.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `exposedServices (refList)`

ExposedServices is the list of services that this processing unit is
implementing.

#### `hashedTags (external:map_of_string_of_strings)`

hashedTags contains the list of tags that matched the policies and their hashes.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `ingressPolicies (external:_rendered_policy)`

IngressPolicies lists all the ingress policies attached to processing unit.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `matchingTags (list)`

MatchingTags contains the list of tags that matched the policies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `processingUnit (ref)`

Can be set during a POST operation to render a policy on a Processing Unit that
has not been created yet.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `processingUnitID (string)`

Identifier of the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `profile (external:map_of_string_of_strings)`

Profile is the trust profile of the processing unit that should be used during
all communications.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `scopes (list)`

Scopes is the set of scopes granted to this processing unit that it has to
present in HTTP requests.

## Report

Post a new statistics report.

### Relations

#### `POST /reports`

Create a statistics report.

### Attributes

#### `fields (external:map_of_string_of_objects)`

TSDB Fields to set for the report.

#### `kind (enum)`

Kind contains the kind of report.

| Characteristics | Value                                                                |
| -               | -:                                                                   |
| Allowed Value   | `Audit, Enforcer, FileAccess, Flow, ProcessingUnit, Syscall, Claims` |

#### `tags (external:map_of_string_of_strings)`

Tags contains the tags associated to the data point.

#### `timestamp (time)`

Timestamp contains the time for the report.

#### `value (float)`

Value contains the value for the report.

## Role

Roles returns the available roles that can be used with API Authorization
Policies.

### Relations

#### `GET /roles`

Retrieves the list of existing roles.

### Attributes

#### `authorizations (external:map_of_string_of_list_of_strings)`

Authorizations of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `description (string)`

Description is the description of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `key (string)`

Key is the of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `name (string)`

Name of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## Service

A Service defines a generic service object at L4 or L7 that encapsulates the
description of a micro-service. A service exposes APIs and can be implemented
through third party entities (such as a cloud provider) or through  processing
units.

### Example

```json
{
  "OIDCProviderURL": "https://accounts.google.com",
  "OIDCScopes": [
    "email",
    "profile"
  ],
  "exposedAPIs": [
    [
      "package=p1"
    ]
  ],
  "exposedPort": 443,
  "name": "the name",
  "port": 443,
  "publicApplicationPort": 443,
  "selectors": [
    [
      "$identity=processingunit"
    ]
  ]
}
```

### Relations

#### `GET /services`

Retrieves the list of Services.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `archived` (boolean): Also retrieve the objects that have been archived.

#### `POST /services`

Creates a new Service.

#### `DELETE /services/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /services/:id`

Retrieves the object with the given ID.

##### Parameters

- `archived` (boolean): Also retrieve the objects that have been archived.

#### `PUT /services/:id`

Updates the object with the given ID.

#### `GET /networkaccesspolicies/:id/services`

Returns the list of services affected by a network access policy.

##### Parameters

- `mode` (enum): Matching mode.

#### `GET /processingunits/:id/services`

Retrieves the services used by a processing unit.

#### `GET /servicedependencies/:id/services`

Returns the list of external services that are targets of service dependency.

#### `GET /services/:id/httpresourcespecs`

Retrieves the HTTP Resource exposed by this service.

#### `GET /services/:id/processingunits`

Retrieves the Processing Units that implement this service.

##### Parameters

- `mode` (enum): Matching mode.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `IPs (list)`

IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from host names otherwise.

#### `JWTSigningCertificate (string)`

PEM encoded certificate that will be used to validate user JWT in HTTP requests.
This is an optional field, needed only if the `authorizationType`
is set to `JWT`.

#### `MTLSCertificateAuthority (string)`

PEM encoded Certificate Authority to use to verify client
certificates. This only applies if `authorizationType` is set to
`MTLS`. If it is not set, Aporeto's Public Signing Certificate Authority will
be used.

#### `OIDCCallbackURL (string)`

This is an advanced setting. Optional OIDC callback URL. If you don't set it,
Aporeto will autodiscover it. It will be
`https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback`.

#### `OIDCClientID (string)`

OIDC Client ID. Only has effect if the `authorizationType` is set to `OIDC`.

#### `OIDCClientSecret (string)`

OIDC Client Secret. Only has effect if the `authorizationType` is set to `OIDC`.

#### `OIDCProviderURL (string)`

OIDC Provider URL. Only has effect if the `authorizationType` is set to `OIDC`.

#### `OIDCScopes (list)`

Configures the scopes you want to add to the OIDC provider. Only has effect if
`authorizationType` is set to `OIDC`.

#### `TLSCertificate (string)`

PEM encoded certificate to expose to the clients for TLS. Only has effect and
required if `TLSType` is set to `External`.

#### `TLSCertificateKey (string)`

PEM encoded certificate key associated to `TLSCertificate`. Only has effect and
required if `TLSType` is set to `External`.

#### `TLSType (enum)`

Set how to provide a server certificate to the service.

* `Aporeto`: Generate a certificate issued from Aporeto public CA.
* `LetsEncrypt`: Issue a certificate from letsencrypt.
* `External`: : Let you define your own certificate and key to use.
* `None`: : TLS is disabled (not recommended).

| Characteristics | Value                                  |
| -               | -:                                     |
| Allowed Value   | `Aporeto, LetsEncrypt, External, None` |
| Default         | `"Aporeto"`                            |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `authorizationType (enum)`

AuthorizationType defines the user authorization type that should be used.

* `None`: No auhtorization.
* `JWT`:  Configures a simple JWT verification from the HTTP `Auhorization`
Header
* `OIDC`: Configures OIDC authorization. You must then set `OIDCClientID`,
`OIDCClientSecret`, OIDCProviderURL`.
* `MTLS`: Configures Client Certificate authorization. Then you can optionaly
`MTLSCertificateAuthority` otherwise Aporeto Public Signing Certificate will be
used.

| Characteristics | Value                   |
| -               | -:                      |
| Allowed Value   | `None, JWT, OIDC, MTLS` |
| Default         | `"None"`                |

#### `claimsToHTTPHeaderMappings (refList)`

Defines a list of mappings between claims and
HTTP headers. When these mappings are defined, the enforcer will copy the
values of the claims to the corresponding HTTP headers.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `endpoints (refList)`

Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `exposedAPIs (external:list_of_lists_of_strings)`

ExposedAPIs contains a tag expression that will determine which
APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
similar specifications for other L7 protocols.

#### `exposedPort (integer)`

ExposedPort is the port that the service can be accessed. Note that
this is different from the Port attribute that describes the port that the
service is actually listening. For example if a load balancer is used, the
ExposedPort is the port that the load balancer is listening for the service,
whereas the port that the implementation is listening can be different.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65535` |
| Required        | `true`  |

#### `exposedServiceIsTLS (boolean)`

ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
the enforcer has to initiate a TLS session in order to forrward traffic to the
service.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |
| Filterable      | `true`  |

#### `external (boolean)`

External is a boolean that indicates if this is an external service.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |
| Filterable      | `true`  |

#### `hosts (list)`

Hosts are the names that the service can be accessed with.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `port (integer)`

Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65535` |
| Required        | `true`  |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `publicApplicationPort (integer)`

PublicApplicationPort is a new virtual port that the service can
be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
application path, you might want to declare a new port where the enforcer
listens for TLS. However, the application does not need to be modified and
the enforcer will map the traffic to the correct application port. This useful
when an application is being accessed from a public network.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65535` |

#### `redirectURLOnAuthorizationFailure (string)`

If this is set, the user will be redirected to that URL in case of any
authorization failure to let you chance to provide a nice message to the user.
The query parameter `?failure_message=<message>` will be added to that url
explaining the possible reasons of the failure.

#### `selectors (external:list_of_lists_of_strings)`

Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.

#### `trustedCertificateAuthorities (string)`

PEM encoded Certificate Authorities to trust when additional hops are needed. It
must be set if the service must reach a Service marked as `external` or must go
through an additional TLS termination point like a L7 Load Balancer.

#### `type (enum)`

Type is the type of the service.

| Characteristics | Value                                        |
| -               | -:                                           |
| Allowed Value   | `HTTP, TCP, KubernetesSecrets, VaultSecrets` |
| Default         | `"HTTP"`                                     |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## ServiceDependency

Allows to define a service dependency where a set of processing units as defined
by their tags require access to specific services.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /servicedependencies`

Retrieves the list of service dependencies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /servicedependencies`

Creates a new service dependency.

#### `DELETE /servicedependencies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /servicedependencies/:id`

Retrieves the object with the given ID.

#### `PUT /servicedependencies/:id`

Updates the object with the given ID.

#### `GET /servicedependencies/:id/processingunits`

Returns the list of Processing Units that depend on an service.

#### `GET /servicedependencies/:id/services`

Returns the list of external services that are targets of service dependency.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:list_of_lists_of_strings)`

Object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## StatsInfo

StatsInfo API can be used to list the fields and tags available in a
statistics measurement.

### Relations

#### `POST /statsinfo`

Retrieves information about the content of the stats measurement.

### Attributes

#### `fields (external:map_of_string_of_strings)`

Contains the list of fields. You cannot group by these fields.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `measurement (enum)`

Name of the measurement.

| Characteristics | Value                                       |
| -               | -:                                          |
| Allowed Value   | `Flows, Audit, Enforcers, Files, EventLogs` |
| Default         | `"Flows"`                                   |

#### `tags (list)`

Contains the list of tags. You can group by these tags.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## StatsQuery

StatsQuery is a generic API to retrieve time series data stored by the Aporeto
system. The API allows different types of queries that are all protected within
the namespace of the user.

### Relations

#### `GET /statsqueries`

Retrieves statistics information based on parameters.

##### Parameters

- `field` (string): list of fields to query.
- `function` (enum): function to use.
- `groupBy` (string): list of groupBy clauses.
- `interval` (duration): list of groupBy clauses.
- `measurement` (enum): Name of the measurement to query.
- `query` (string): Raw InfluxDB query. This query will be applied to an immutable subquery that will bound the data to your namespace. You must pass the measurement parameter and use `from _dataset_` in your query. The dataset will be bound to the timing parameter you pass.
- `tag` (string): list of tags to query.
- `where` (string): list of where clauses. If `query` is given, this where clause will be applied to the controlled subquery as additional `and` clauses to the internally computed where clause.
- `endAbsolute` (time): Set the absolute end of the time window.
- `endRelative` (duration): Set the relative end of the time window.
- `startAbsolute` (time): Set the absolute start of the time window.
- `startRelative` (duration): Set the relative start of the time window.
- `limit` (integer): Maximum number of results.
- `offset` (integer): Maximum number of results.

##### Mandatory Parameters

(`endRelative`) or (`startRelative`) or (`startRelative` and `endRelative`) or (`startRelative` and `endAbsolute`) or (`startAbsolute` and `endRelative`) or (`startAbsolute` and `endAbsolute`)

#### `POST /statsqueries`

Sends a query on statistical data.

##### Parameters

- `endAbsolute` (time): Set the absolute end of the time window.
- `endRelative` (duration): Set the relative end of the time window.
- `startAbsolute` (time): Set the absolute start of the time window.
- `startRelative` (duration): Set the relative start of the time window.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

##### Mandatory Parameters

(`endRelative`) or (`startRelative`) or (`startRelative` and `endRelative`) or (`startRelative` and `endAbsolute`) or (`startAbsolute` and `endRelative`) or (`startAbsolute` and `endAbsolute`)

### Attributes

#### `descending (boolean)`

If set, the results will be order by time from the most recent to the oldest.

#### `fields (list)`

List of fields to extract. If you don't pass anything, all available fields will
be returned. It is also possible to use function like `sum(value)`.

#### `filter (string)`

Apply a filter to the query.

#### `groups (list)`

Group results by the provided values. Note that not all fields can be used to
group the results.

#### `limit (integer)`

Limits the number of results. -1 means no limit.

| Characteristics | Value |
| -               | -:    |
| Default         | `-1`  |

#### `measurement (enum)`

Name of the measurement.

| Characteristics | Value                                                                |
| -               | -:                                                                   |
| Allowed Value   | `Flows, Audit, Enforcers, Files, EventLogs, Packets, EnforcerTraces` |
| Default         | `"Flows"`                                                            |

#### `offset (integer)`

Offsets the of results. -1 means no offset.

| Characteristics | Value |
| -               | -:    |
| Default         | `-1`  |

#### `results (refList)`

Results contains the result of the query.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## SuggestedPolicy

Allows to get policy suggestions.

### Relations

#### `GET /suggestedpolicies`

Retrieves a list of network policy suggestion.

##### Parameters

- `filterAction` (enum): Action to take with the filter tags.
- `filterTags` (string): Tags to filter in the policy suggestions.
- `endAbsolute` (time): Set the absolute end of the time window.
- `endRelative` (duration): Set the relative end of the time window.
- `startAbsolute` (time): Set the absolute start of the time window.
- `startRelative` (duration): Set the relative start of the time window.
- `flowOffset` (duration): Apply an offset to the time window for flows.

##### Mandatory Parameters

(`endRelative`) or (`startRelative`) or (`startRelative` and `endRelative`) or (`startRelative` and `endAbsolute`) or (`startAbsolute` and `endRelative`) or (`startAbsolute` and `endAbsolute`)

### Attributes

#### `networkAccessPolicies (refList)`

List of suggested network access policies.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

## Tabulation

Tabulate API allows you to retrieve a custom table view for any identity using
any tags you like as columns.

### Relations

#### `GET /tabulations`

Retrieves tabulated informations based on parameters.

##### Parameters

- `column` (string): Columns you want to see.
- `identity` (string): Identity you want to tabulate.
- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

##### Mandatory Parameters

`identity`

### Attributes

#### `headers (list)`

Headers contains the requests headers that matched.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `rows (external:list_of_lists_of_objects)`

Rows contains the tabulated data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `targetIdentity (string)`

TargetIdentity contains the requested target identity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## Tag

A tag is a string in the form of "key=value" that can applied to all objects in
the system. They are used for policy resolution. Tags starting by a "$" are
derived from the property of an object (for instance an object with ID set to
xxx and a name set to "the name" will be tagged by default with "$name=the name"
and "$id=xxx"). Tags starting with an "@" have been generated by an external
system.

### Example

```json
{
  "value": "key=value"
}
```

### Relations

#### `GET /tags`

Retrieves the list of existing tags in the system.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `count (integer)`

Count represents the number of time the tag is used.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `namespace (string)`

Namespace represents the namespace of the counted tag.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `value (string)`

Value represents the value of the tag.

| Characteristics | Value                                                     |
| -               | -:                                                        |
| Format          | `/^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@~<>#/-]+$/` |
| Required        | `true`                                                    |
| Creation only   | `true`                                                    |

## TagValue

Represents all values associated to tag key.

### Relations

#### `GET /tagvalues`

Retrieves the list of existing values for the given tag keys.

##### Parameters

- `key` (string): Keys of the tag you want to get the values of.

##### Mandatory Parameters

`key`

### Attributes

#### `key (string)`

The requested key.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `values (list)`

List of all values.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## TimeSeriesQueryResults

Represent the results of a stats query.

### Attributes

#### `rows (refList)`

List of rows.

## TimeSeriesRow

Represent a time series row.

### Attributes

#### `columns (list)`

colums of the row.

#### `name (string)`

the name of row.

#### `tags (external:map_of_string_of_strings)`

List of tags.

#### `values (external:list_of_lists_of_objects)`

List of tags.

## TokenScopePolicy

The TokenScopePolicy defines a set of policies that allow customization of the
authorization tokens issued by the Aporeto service. This allows Aporeto
generated tokens to be used by external applications.

### Example

```json
{
  "name": "the name"
}
```

### Relations

#### `GET /tokenscopepolicies`

Retrieves the list of token scope policies.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.
- `propagated` (boolean): Also retrieve the objects that propagate down.

#### `POST /tokenscopepolicies`

Creates a new token scope policy.

#### `DELETE /tokenscopepolicies/:id`

Deletes the object with the given ID.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `GET /tokenscopepolicies/:id`

Retrieves the object with the given ID.

#### `PUT /tokenscopepolicies/:id`

Updates the object with the given ID.

### Attributes

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (string)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `assignedScopes (list)`

AssignedScopes is the the list of scopes that the policiy will assigns.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `metadata (list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `subject (external:list_of_lists_of_strings)`

Subject defines the selection criteria that this policy must match on identiy
and scope request information.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## TraceMode

TraceMode is the tracing mode that must be applied to a PU.

### Attributes

#### `IPTables (boolean)`

IPTables instructs the enforcers to provide an iptables trace for a PU.

#### `applicationConnections (boolean)`

Instructs the enforcer to send records for all
application initiated connections.

#### `interval (string)`

Determines the length of the time interval that the trace must be
enabled.

| Characteristics | Value   |
| -               | -:      |
| Default         | `"10s"` |

#### `networkConnections (boolean)`

Instructs the enforcer to send records for all network
initiated connections.

## TraceRecord

Represents a single trace record from the enforcer.

### Example

```json
{
  "TTL": 64,
  "chain": "PREROUTING",
  "destinationIP": "10.1.1.30",
  "destinationInterface": "en0",
  "destinationPort": 80,
  "length": 98,
  "packetID": 10,
  "protocol": 80,
  "ruleID": 10,
  "sourceIP": "10.1.1.30",
  "sourceInterface": "en0",
  "sourcePort": 80,
  "tableName": "raw",
  "timestamp": "2018-06-14T23:10:46.420397985Z"
}
```

### Attributes

#### `TTL (integer)`

TTL is the TTL value of the packet.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `255`  |
| Required        | `true` |

#### `chain (string)`

Chain is the chain that the trace was collected from.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `destinationIP (string)`

DestinationIP is the destination IP.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `destinationInterface (string)`

DestinationInterface is the destination interface of the packet.

#### `destinationPort (integer)`

DestinationPort is the destination UPD or TCP port of the packet.

| Characteristics | Value   |
| -               | -:      |
| Min length      | `1`     |
| Max length      | `65536` |
| Required        | `true`  |

#### `length (integer)`

Length of the observed packet.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65536` |
| Required        | `true`  |

#### `packetID (integer)`

PacketID is the IP packet header ID.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `protocol (integer)`

Protocol is the protocol of the packets.

| Characteristics | Value   |
| -               | -:      |
| Max length      | `65536` |
| Required        | `true`  |

#### `ruleID (integer)`

ruleID is the priority index of the iptables entry that was hit.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `sourceIP (string)`

SourceIP is the source IP of the packet.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `sourceInterface (string)`

SourceInterface is the source interface of the packet.

#### `sourcePort (integer)`

SourcePort is the source TCP or UDP Port of the packet.

| Characteristics | Value   |
| -               | -:      |
| Min length      | `1`     |
| Max length      | `65536` |
| Required        | `true`  |

#### `tableName (string)`

TableName is the iptable name that the trace was collected.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `timestamp (time)`

Timestamp is the date of the report.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

## Trigger

Trigger can be used to remotely trigger an automation.

### Relations

#### `GET /automations/:id/triggers`

Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`.

#### `POST /automations/:id/triggers`

Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`.

## Vulnerability

A vulnerabily represents a particular CVE.

### Example

```json
{
  "CVSS2Score": 3.2,
  "link": "https://cve.com/CVE-1234",
  "name": "the name",
  "severity": 3
}
```

### Relations

#### `GET /vulnerabilities`

Retrieves the list of vulnerabilities.

##### Parameters

- `q` (string): Filtering query. Consequent `q` parameters will form an or.
- `tag` (string): List of tags to filter on. This parameter is deprecated.

#### `POST /vulnerabilities`

Creates a new vulnerability.

#### `GET /vulnerabilities/:id`

Retrieves the object with the given ID.

#### `GET /processingunits/:id/vulnerabilities`

Retrieves the vulnerabilities affecting the processing unit.

#### `GET /vulnerabilities/:id/processingunits`

Retrieves the processing units affected by the a vulnerabily.

### Attributes

#### `CVSS2Score (float)`

CVSS v2 score.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:map_of_string_of_list_of_strings)`

Annotation stores additional information about an entity.

#### `associatedTags (list)`

AssociatedTags are the list of tags attached to an entity.

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `link (string)`

Link is the URL that refers to the vulnerability.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace tag attached to an entity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `severity (external:_vulnerability_level)`

Severity refers to the security vulnerability level.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
