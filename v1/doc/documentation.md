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
| [AuditProfile](#auditprofile)                                 | AuditProfile is an audit policy that consists of a set of audit rules. An audit     |
| [Auth](#auth)                                                 | This API verifies if the given token is valid or not.                               |
| [Automation](#automation)                                     | An automation needs documentation.                                                  |
| [AutomationTemplate](#automationtemplate)                     | Templates that ca be used in automations.                                           |
| [AWSAccount](#awsaccount)                                     | Allows to bind an AWS account to your Aporeto account to allow auto registration... |
| [Category](#category)                                         | Category allows to categorized services.                                            |
| [Certificate](#certificate)                                   | A User represents the owner of some certificates.                                   |
| [DependencyMap](#dependencymap)                               | This api returns a data structure representing the graph of all processing units... |
| [Enforcer](#enforcer)                                         | An Enforcer Profile contains a configuration for a Enforcer. It contains various... |
| [EnforcerProfile](#enforcerprofile)                           | Allows to create reusable configuration profile for your enforcers. Enforcer        |
| [EnforcerProfileMappingPolicy](#enforcerprofilemappingpolicy) | A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used... |
| [Export](#export)                                             | Export the policies and related objects in a given namespace.                       |
| [ExternalAccess](#externalaccess)                             | ExternalAccess allows to retrieve connection from or to an external service.        |
| [ExternalService](#externalservice)                           | An External Service represents a random network or ip that is not managed by the... |
| [FileAccess](#fileaccess)                                     | Returns file access statistics on a particular processing unit.                     |
| [FileAccessPolicy](#fileaccesspolicy)                         | A File Access Policy allows Processing Units to access various folder and files.... |
| [FilePath](#filepath)                                         | A File Path represents a random path to a file or a folder. They can be used in     |
| [FlowStatistic](#flowstatistic)                               | Returns network access statistics on a particular processing unit or group of       |
| [HookPolicy](#hookpolicy)                                     | Hook allows to to define hooks to the write operations in squall. Hooks are sent... |
| [Import](#import)                                             | Imports an export of policies and related objects into the namespace.               |
| [Installation](#installation)                                 | Installation represents an installation for a given account.                        |
| [InstalledApp](#installedapp)                                 | InstalledApps represents an installed application.                                  |
| [IsolationProfile](#isolationprofile)                         | An IsolationProfile needs documentation.                                            |
| [Issue](#issue)                                               | This API issues a new token according to given data.                                |
| [Jaegerbatch](#jaegerbatch)                                   | A jaegerbatch is a batch of jaeger spans. This is used by external service to       |
| [KubernetesCluster](#kubernetescluster)                       | Create a remote Kubernetes Cluster integration.                                     |
| [Log](#log)                                                   | Retrieves the log of a deployed app.                                                |
| [Message](#message)                                           | The Message API allows to post public messages that will be visible through all     |
| [Namespace](#namespace)                                       | A Namespace represents the core organizational unit of the system. All objects      |
| [NamespaceMappingPolicy](#namespacemappingpolicy)             | A Namespace Mapping Policy defines in which namespace a Processing Unit should      |
| [NetworkAccessPolicy](#networkaccesspolicy)                   | Allows to define networking policies to allow or prevent processing units           |
| [PasswordReset](#passwordreset)                               | Used to reset an account password.                                                  |
| [Plan](#plan)                                                 | Plan contains the various billing plans available.                                  |
| [Poke](#poke)                                                 | When available, poke can be used to update various information about the parent.... |
| [Policy](#policy)                                             | Policy represents the policy primitive used by all aporeto policies.                |
| [PolicyRefresh](#policyrefresh)                               | PolicyRefresh is sent to client when as a push event when a policy refresh is       |
| [PolicyRule](#policyrule)                                     | PolicyRule is an internal policy resolution API. Services can use this API to       |
| [ProcessingUnit](#processingunit)                             | A Processing Unit reprents anything that can compute. It can be a Docker            |
| [ProcessingUnitPolicy](#processingunitpolicy)                 | A ProcessingUnitPolicies needs a better description.                                |
| [QuotaPolicy](#quotapolicy)                                   | Quotas Policies allows to set quotas on the number of objects that can be           |
| [RemoteProcessor](#remoteprocessor)                           | Hook to integrate an Aporeto service.                                               |
| [RenderedPolicy](#renderedpolicy)                             | Retrieve the aggregated policies applied to a particular processing unit.           |
| [Report](#report)                                             | Post a new statistics report.                                                       |
| [RESTAPISpec](#restapispec)                                   | RESTAPISpec descibes the REST APIs exposed by a service. These APIs                 |
| [Role](#role)                                                 | Roles returns the available roles that can be used with API Authorization           |
| [Root](#root)                                                 | root object.                                                                        |
| [Service](#service)                                           | A Service defines a generic service object at L4 or L7 that encapsulates the        |
| [StatsQuery](#statsquery)                                     | StatsQuery is a generic API to retrieve time series data stored by the Aporeto      |
| [SuggestedPolicy](#suggestedpolicy)                           | Allows to get policy suggestions.                                                   |
| [Tabulation](#tabulation)                                     | Tabulate API allows you to retrieve a custom table view for any identity using      |
| [Tag](#tag)                                                   | A tag is a string in the form of "key=value" that can applied to all objects in     |
| [TokenScopePolicy](#tokenscopepolicy)                         | The TokenScopePolicy defines a set of policies that allow customization of the      |
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

| Method   | URL         | Description |
| -:       | -           | -           |
| `GET`    | `/accounts` | Retrieves all accounts. This is a private API that can only be done by the
system.    |
| `POST`   | `/accounts`     | Creates a new Account.                       |
| `DELETE` | `/accounts/:id` | Deletes the `account` with the given `:id`.  |
| `GET`    | `/accounts/:id` | Retrieve the `account` with the given `:id`. |
| `PUT`    | `/accounts/:id` | Updates the `account` with the given `:id`.  |

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
| Filterable      | `true` |

#### `LDAPBindSearchFilter (string)`

LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
other systems, value may be 'uid={USERNAME}'.

| Characteristics | Value            |
| -               | -:               |
| Default         | `uid={USERNAME}` |
| Orderable       | `true`           |
| Filterable      | `true`           |

#### `LDAPCertificateAuthority (string)`

LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPConnSecurityProtocol (enum)`

LDAPConnProtocol holds the connection type for the LDAP provider.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `TLS, InbandTLS` |
| Default         | `InbandTLS`      |
| Orderable       | `true`           |
| Filterable      | `true`           |

#### `LDAPEnabled (boolean)`

LDAPEnabled triggers if the account uses it's own LDAP for authentication.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPIgnoredKeys (external:ignore_list)`

LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `LDAPSubjectKey (string)`

LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.

| Characteristics | Value  |
| -               | -:     |
| Default         | `uid`  |
| Orderable       | `true` |
| Filterable      | `true` |

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
| Filterable      | `true` |

#### `activationToken (string)`

ActivationToken contains the activation token.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

#### `associatedPlanKey (string)`

AssociatedPlanKey contains the plan key that is associated to this account.

| Characteristics | Value               |
| -               | -:                  |
| Default         | `aporeto.plan.free` |
| Autogenerated   | `true`              |
| Read only       | `true`              |

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
| Filterable      | `true` |

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
| Filterable      | `true` |

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
| Default         | `Pending`                            |
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
| Filterable      | `true` |

## Activate

Used to activate a pending account.

### Relations

| Method | URL         | Description                  |
| -:     | -           | -                            |
| `GET`  | `/activate` | Activates a pending account. |

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

| Method | URL               | Description                                   |
| -:     | -                 | -                                             |
| `GET`  | `/activities`     | Retrieves the list of activity logs.          |
| `GET`  | `/activities/:id` | Retrieve the `activity` with the given `:id`. |

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

#### `claims (external:raw_data)`

Claims of the user who performed the operation.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `data (external:raw_data)`

Data of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `date (time)`

Date of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `error (external:raw_data)`

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
| Filterable      | `true` |

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

#### `originalData (external:raw_data)`

OriginalData contains the eventual original data of the object that has been
modified.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

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

| Method   | URL           | Description                                |
| -:       | -             | -                                          |
| `GET`    | `/alarms`     | Retrieves all the alarms.                  |
| `POST`   | `/alarms`     | Creates a new alarm.                       |
| `DELETE` | `/alarms/:id` | Deletes the `alarm` with the given `:id`.  |
| `GET`    | `/alarms/:id` | Retrieve the `alarm` with the given `:id`. |
| `PUT`    | `/alarms/:id` | Updates the `alarm` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `data (external:alarm_data)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `occurrences (external:alarm_occurrences)`

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
| Filterable      | `true` |

#### `status (enum)`

Status of the alarm.

| Characteristics | Value                          |
| -               | -:                             |
| Allowed Value   | `Acknowledged, Open, Resolved` |
| Default         | `Open`                         |
| Orderable       | `true`                         |
| Filterable      | `true`                         |

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

| Method   | URL                             | Description                                                 |
| -:       | -                               | -                                                           |
| `GET`    | `/apiauthorizationpolicies`     | Retrieves the list of API authorization policies.           |
| `POST`   | `/apiauthorizationpolicies`     | Creates a new API authorization policies.                   |
| `DELETE` | `/apiauthorizationpolicies/:id` | Deletes the `apiauthorizationpolicy` with the given `:id`.  |
| `GET`    | `/apiauthorizationpolicies/:id` | Retrieve the `apiauthorizationpolicy` with the given `:id`. |
| `PUT`    | `/apiauthorizationpolicies/:id` | Updates the `apiauthorizationpolicy` with the given `:id`.  |

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

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity.

#### `authorizedIdentities (external:identity_list)`

AuthorizedIdentities defines the list of api identities the policy applies to.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `authorizedNamespace (string)`

AuthorizedNamespace defines on what namespace the policy applies.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

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

| Method | URL          | Description                                                         |
| -:     | -            | -                                                                   |
| `POST` | `/apichecks` | Verfies the authorizations on various identities for a given token. |

### Attributes

#### `authorized (external:authorized_identities)`

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

#### `targetIdentities (external:identity_list)`

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

| Method | URL     | Description                 |
| -:     | -       | -                           |
| `GET`  | `/apps` | Retrieves the list of apps. |

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
| Filterable      | `true` |

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

#### `parameters (external:app_parameters)`

Parameters of the app the user can or has to specify.

#### `title (string)`

Title represents the title of the app.

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

| Method   | URL                                   | Description                                                    |
| -:       | -                                     | -                                                              |
| `GET`    | `/auditprofiles`                      | Retrieves the list of audit profiles.                          |
| `POST`   | `/auditprofiles`                      | Creates a new audit profile.                                   |
| `DELETE` | `/auditprofiles/:id`                  | Deletes the `auditprofile` with the given `:id`.               |
| `GET`    | `/auditprofiles/:id`                  | Retrieve the `auditprofile` with the given `:id`.              |
| `PUT`    | `/auditprofiles/:id`                  | Updates the `auditprofile` with the given `:id`.               |
| `GET`    | `/enforcerprofiles/:id/auditprofiles` | Returns the list of AuditProfiles used by an enforcer profile. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `propagated (boolean)`

Propagated indicates if the audit profile is propagated.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `rules (external:audit_profile_rule_list)`

Rules is the list of audit policy rules associated with this policy.

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Auth

This API verifies if the given token is valid or not.

### Relations

| Method | URL     | Description                     |
| -:     | -       | -                               |
| `GET`  | `/auth` | Verify the validity of a token. |

### Attributes

#### `claims (external:claims)`

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

| Method         | URL                         | Description                                     |
| -:             | -                           | -                                               |
| `GET`          | `/automations`              | Retrieves the list of Automations.              |
| `POST`         | `/automations`              | Creates a new Automation.                       |
| `DELETE`       | `/automations/:id`          | Deletes the `automation` with the given `:id`.  |
| `GET`          | `/automations/:id`          | Retrieve the `automation` with the given `:id`. |
| `PUT`          | `/automations/:id`          | Updates the `automation` with the given `:id`.  |
| `GET`          | `/automations/:id/triggers` | Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`. |
| `POST`         | `/automations/:id/triggers` | Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `entitlements (external:automation_entitlements)`

Entitlements declares which operations are allowed on which identities.

#### `errors (list)`

Error contains the eventual error of the last run.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `events (external:automation_events)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `parameters (external:automation_parameters)`

Parameters are passed to the functions.

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

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
| Default         | `Time`                    |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

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

| Method | URL                        | Description                                             |
| -:     | -                          | -                                                       |
| `GET`  | `/automationtemplates`     | Retrieves the list of automation templates.             |
| `GET`  | `/automationtemplates/:id` | Retrieve the `automationtemplate` with the given `:id`. |

### Attributes

#### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `1024` |
| Orderable       | `true` |

#### `entitlements (external:automation_entitlements)`

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
| Default         | `Condition`         |

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `parameters (external:automation_template_parameters)`

Parameters contains the parameter description of the function.

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

| Method   | URL                | Description                                     |
| -:       | -                  | -                                               |
| `GET`    | `/awsaccounts`     | Retrieves the list of aws account bindings.     |
| `POST`   | `/awsaccounts`     | Creates a new aws account binding.              |
| `DELETE` | `/awsaccounts/:id` | Deletes the `awsaccount` with the given `:id`.  |
| `GET`    | `/awsaccounts/:id` | Retrieve the `awsaccount` with the given `:id`. |

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
| Filterable      | `true` |

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
| Filterable      | `true` |

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

| Method   | URL                 | Description                                       |
| -:       | -                   | -                                                 |
| `GET`    | `/certificates`     | Retrieves the list of existing user certificates. |
| `POST`   | `/certificates`     | Creates a new user certificate.                   |
| `DELETE` | `/certificates/:id` | Deletes the `certificate` with the given `:id`.   |
| `GET`    | `/certificates/:id` | Retrieve the `certificate` with the given `:id`.  |
| `PUT`    | `/certificates/:id` | Updates the `certificate` with the given `:id`.   |

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
| Filterable      | `true` |

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
| Filterable      | `true` |

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
| Filterable      | `true` |

#### `serialNumber (string)`

SerialNumber of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `status (enum)`

CertificateStatus provides the status of the certificate. Update with RENEW to
get a new certificate.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `Revoked, Valid` |
| Default         | `Valid`          |
| Orderable       | `true`           |
| Filterable      | `true`           |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

## DependencyMap

This api returns a data structure representing the graph of all processing units
and their connections in a particular namespace, in a given time window. To pass
the time window you can use the query parameters 'startAbsolute', 'endAbsolute',
'startRelative', 'endRelative'.

For example
  "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000".

### Relations

| Method | URL               | Description                                 |
| -:     | -                 | -                                           |
| `GET`  | `/dependencymaps` | Retrieves the dependencymap of a namespace. |

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

#### `claims (external:graphclaims_map)`

claims represents a user or a script that have accessed an api.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `edges (external:graphedges_map)`

edges are the edges of the map.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `groups (external:graphgroups_map)`

Groups provide information about the group values.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `nodes (external:graphnodes_map)`

nodes refers to the nodes of the map.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `viewSuggestions (external:view_suggestions)`

viewSuggestions provides suggestion of views based on relevant tags.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

## Enforcer

An Enforcer Profile contains a configuration for a Enforcer. It contains various
parameters, like what network should not policeds, what processing units should
be ignored based on their tags and so on. It also contains more advanced
parameters to fine tune the Agent. A Enforcer will decide what profile to apply
using aEnforcer Profile Mapping Policy. This policy will decide according the
Enforcer's tags what profile to use. If an Enforcer tags are matching more than
a single policy, it will refuse to start. Some parameters will be applied
directly to a running agent, some will need to restart it.

### Example

```json
{
  "FQDN": "server1.domain.com",
  "name": "the name"
}
```

### Relations

| Method   | URL                                             | Description                                                                      |
| -:       | -                                               | -                                                                                |
| `GET`    | `/enforcers`                                    | Retrieves the list of enforcers.                                                 |
| `POST`   | `/enforcers`                                    | Creates a new enforcer.                                                          |
| `DELETE` | `/enforcers/:id`                                | Deletes the `enforcer` with the given `:id`.                                     |
| `GET`    | `/enforcers/:id`                                | Retrieve the `enforcer` with the given `:id`.                                    |
| `PUT`    | `/enforcers/:id`                                | Updates the `enforcer` with the given `:id`.                                     |
| `GET`    | `/enforcerprofilemappingpolicies/:id/enforcers` | Returns the list of enforcers affected by an enforcer profile mapping policy.    |
| `GET`    | `/enforcers/:id/enforcerprofiles`               | Returns the enforcer profile that must be used by an enforcer.                   |
| `GET`    | `/enforcers/:id/poke`                           | Sends a poke empty object. This is used to ensure an enforcer is up and running. |

### Attributes

#### `FQDN (string)`

FQDN contains the fqdn of the server where the enforcer is running.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

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

#### `collectedInfo (external:collected_info)`

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
| Filterable      | `true` |

#### `lastCollectionTime (time)`

LastCollectionTime represents the date and time when the info have been
collected.

#### `lastSyncTime (time)`

LastSyncTime holds the last heart beat time.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `localCA (string)`

LocalCA contains the initial chain of trust for the enforcer. This valud is only
given when you retrieve a single enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `operationalStatus (enum)`

OperationalStatus tells the status of the enforcer.

| Characteristics | Value                                           |
| -               | -:                                              |
| Allowed Value   | `Connected, Disconnected, Initialized, Unknown` |
| Default         | `Initialized`                                   |
| Autogenerated   | `true`                                          |
| Read only       | `true`                                          |
| Filterable      | `true`                                          |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `publicToken (string)`

PublicToken is the public token of the server that will be included in the
datapath and its signed by the private CA.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `updateAvailable (boolean)`

Tells if the the version of enforcerd is outdated and should be updated.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

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

| Method   | URL                                                    | Description                                          |
| -:       | -                                                      | -                                                    |
| `GET`    | `/enforcerprofiles`                                    | Retrieves the list of enforcer profiles.             |
| `POST`   | `/enforcerprofiles`                                    | Creates a new enforcer profile.                      |
| `DELETE` | `/enforcerprofiles/:id`                                | Deletes the `enforcerprofile` with the given `:id`.  |
| `GET`    | `/enforcerprofiles/:id`                                | Retrieve the `enforcerprofile` with the given `:id`. |
| `PUT`    | `/enforcerprofiles/:id`                                | Updates the `enforcerprofile` with the given `:id`.  |
| `GET`    | `/enforcerprofilemappingpolicies/:id/enforcerprofiles` | Returns the list of enforcer profiles that an enforcer profile mapping policy
matches.   |
| `GET`    | `/enforcers/:id/enforcerprofiles`     | Returns the enforcer profile that must be used by an enforcer. |
| `GET`    | `/enforcerprofiles/:id/auditprofiles` | Returns the list of AuditProfiles used by an enforcer profile. |

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
| Filterable      | `true`  |

#### `PUBookkeepingInterval (string)`

PUBookkeepingInterval configures how often the PU will be synchronized.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `15m`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

#### `PUHeartbeatInterval (string)`

PUHeartbeatInterval configures the heart beat interval.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `5s`              |
| Orderable       | `true`            |
| Filterable      | `true`            |

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity.

#### `auditProfileSelectors (external:audit_profile_selector)`

AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `auditProfiles (external:audit_profiles)`

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
| Filterable      | `true`   |

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
| Default         | `unix:///var/run/docker.sock`                                                                                                               |
| Orderable       | `true`                                                                                                                                      |
| Filterable      | `true`                                                                                                                                      |

#### `excludedInterfaces (external:excluded_interfaces_list)`

ExcludedInterfaces is a list of interfaces that must be excluded.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `excludedNetworks (external:excluded_networks_list)`

ExcludedNetworks is the list of networks that must be excluded for this
enforcer.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `hostServices (external:host_services_list)`

HostServices is a list of services that must be activated by default to all
enforcers matching this profile.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `ignoreExpression (external:policies_list)`

IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.

#### `kubernetesMetadataExtractor (enum)`

Select which metadata extractor to use to process new processing units from
Kubernetes.

| Characteristics | Value                                  |
| -               | -:                                     |
| Allowed Value   | `KubeSquall, PodAtomic, PodContainers` |
| Default         | `KubeSquall`                           |
| Orderable       | `true`                                 |
| Filterable      | `true`                                 |

#### `kubernetesSupportEnabled (boolean)`

KubernetesSupportEnabled enables kubernetes mode for the enforcer.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |
| Filterable      | `true`  |

#### `linuxProcessesSupportEnabled (boolean)`

LinuxProcessesSupportEnabled configures support for Linux processes.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |

#### `metadataExtractor (enum)`

Select which metadata extractor to use to process new processing units.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Docker, ECS, Kubernetes` |
| Default         | `Docker`                  |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

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

#### `normalizedTags (external:tags_list)`

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
| Default         | `10m`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `proxyListenAddress (string)`

ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
example:
  unix:///var/run/aporeto.sock.

| Characteristics | Value                                                                                                                                       |
| -               | -:                                                                                                                                          |
| Format          | `/^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix://(/[^/]{1,16}){1,5}/?)$/` |
| Default         | `unix:///var/run/aporeto.sock`                                                                                                              |
| Orderable       | `true`                                                                                                                                      |
| Filterable      | `true`                                                                                                                                      |

#### `receiverNumberOfQueues (integer)`

ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |
| Filterable      | `true` |

#### `receiverQueue (integer)`

ReceiverQueue is the base queue number for traffic from the network.

| Characteristics | Value  |
| -               | -:     |
| Default         | `0`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `receiverQueueSize (integer)`

ReceiverQueueSize is the queue size of the receiver.

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `5000` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `remoteEnforcerEnabled (boolean)`

RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `targetNetworks (external:target_networks_list)`

TargetNetworks is the list of networks that authorization should be applied.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `transmitterNumberOfQueues (integer)`

TransmitterNumberOfQueues is the number of queues for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |
| Filterable      | `true` |

#### `transmitterQueue (integer)`

TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `transmitterQueueSize (integer)`

TransmitterQueueSize is the size of the queue for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `trustedCAs (external:trusted_cas_list)`

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

| Method   | URL                                                    | Description                                                       |
| -:       | -                                                      | -                                                                 |
| `GET`    | `/enforcerprofilemappingpolicies`                      | Retrieves the list of enforcer profile mapping policies.          |
| `POST`   | `/enforcerprofilemappingpolicies`                      | Creates a new enforcer profile mapping policies.                  |
| `DELETE` | `/enforcerprofilemappingpolicies/:id`                  | Deletes the `enforcerprofilemappingpolicy` with the given `:id`.  |
| `GET`    | `/enforcerprofilemappingpolicies/:id`                  | Retrieve the `enforcerprofilemappingpolicy` with the given `:id`. |
| `PUT`    | `/enforcerprofilemappingpolicies/:id`                  | Updates the `enforcerprofilemappingpolicy` with the given `:id`.  |
| `GET`    | `/enforcerprofilemappingpolicies/:id/enforcerprofiles` | Returns the list of enforcer profiles that an enforcer profile mapping policy
matches.   |
| `GET`    | `/enforcerprofilemappingpolicies/:id/enforcers` | Returns the list of enforcers affected by an enforcer profile mapping policy. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:policies_list)`

Object is the list of tags to use to find a enforcer profile.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

Subject is the subject of the policy.

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

## Export

Export the policies and related objects in a given namespace.

### Relations

| Method | URL       | Description                                             |
| -:     | -         | -                                                       |
| `GET`  | `/export` | Exports all policies and related object of a namespace. |

### Attributes

#### `APIVersion (integer)`

APIVersion of the api used for the exported data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `RESTAPISpecs (external:exported_data_content)`

List of all exported RESTAPISpecs.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

#### `auditProfiles (external:exported_data_content)`

List of all exported audit profiles.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `externalServices (external:exported_data_content)`

List of exported external services.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `fileAccessPolicies (external:exported_data_content)`

List of exported file access policies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `filePaths (external:exported_data_content)`

List of exported file paths.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `isolationProfiles (external:exported_data_content)`

List of all exported isolation profiles.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `networkAccessPolicies (external:exported_data_content)`

List of exported network policies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Omit if empty   | `true` |

#### `processingUnitPolicies (external:exported_data_content)`

List of all exported processingUnitPolicies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

#### `services (external:exported_data_content)`

List of all exported services.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

#### `tokenScopePolicies (external:exported_data_content)`

List of all exported tokenScopePolicies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Omit if empty   | `true` |

## ExternalAccess

ExternalAccess allows to retrieve connection from or to an external service.

### Relations

| Method | URL                 | Description                                                    |
| -:     | -                   | -                                                              |
| `GET`  | `/externalaccesses` | Retrieves the list of external access according to parameters. |

### Attributes

#### `IPRecords (external:ip_records)`

IPRecords refers to a list of IPRecord that contains the IP information.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## ExternalService

An External Service represents a random network or ip that is not managed by the
system. They can be used in Network Access Policies in order to allow traffic
from or to the declared network or IP, using the provided protocol and port or
ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
as address, and 1-65000 for the ports. You will need to use the External
Services tags to set some policies.

### Example

```json
{
  "name": "the name",
  "network": "0.0.0.0/0",
  "protocol": "TCP"
}
```

### Relations

| Method   | URL                                           | Description                                                                |
| -:       | -                                             | -                                                                          |
| `GET`    | `/externalservices`                           | Retrieves the list of external services.                                   |
| `POST`   | `/externalservices`                           | Creates a new external service.                                            |
| `DELETE` | `/externalservices/:id`                       | Deletes the `externalservice` with the given `:id`.                        |
| `GET`    | `/externalservices/:id`                       | Retrieve the `externalservice` with the given `:id`.                       |
| `PUT`    | `/externalservices/:id`                       | Updates the `externalservice` with the given `:id`.                        |
| `GET`    | `/networkaccesspolicies/:id/externalservices` | Returns the list of external services affected by a network access policy. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `loadbalancerAddresses (external:addresses_list)`

LoadbalancerAddresses represents the list of adresses of the external services
of type LoadBalancer.

#### `loadbalancerPortsMapping (external:portmapping_list)`

LoadbalancerPortsMapping is the list of ports mapped by an extenral service of
type load balancer.

#### `metadata (external:metadata_list)`

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

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

#### `normalizedTags (external:tags_list)`

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
| Default         | `1:65535`                                                                                                                                                                                                        |
| Filterable      | `true`                                                                                                                                                                                                           |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protocol (string)`

Protocol refers to network protocol like TCP/UDP or the number of the protocol.

| Characteristics | Value                                                                  |
| -               | -:                                                                     |
| Format          | `/^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/` |
| Required        | `true`                                                                 |
| Filterable      | `true`                                                                 |

#### `type (enum)`

Type represents the type of external service.

| Characteristics | Value                                        |
| -               | -:                                           |
| Allowed Value   | `LoadBalancerHTTP, LoadBalancerTCP, Network` |
| Default         | `Network`                                    |
| Orderable       | `true`                                       |
| Filterable      | `true`                                       |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## FileAccess

Returns file access statistics on a particular processing unit.

### Relations

| Method | URL                                 | Description                                                |
| -:     | -                                   | -                                                          |
| `GET`  | `/fileaccesses`                     | Retrieves the list of file access according to parameters. |
| `GET`  | `/processingunits/:id/fileaccesses` | Retrieves the file accesses done by the processing unit.   |

### Attributes

#### `action (string)`

Action tells if the access has been allowed or not.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `count (integer)`

Count tells how many times the file has been accessed.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `host (string)`

Host is the host that served the accessed file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `mode (enum)`

Mode is the mode of the accessed file.

| Characteristics | Value                    |
| -               | -:                       |
| Allowed Value   | `Read, ReadWrite, Write` |
| Autogenerated   | `true`                   |
| Read only       | `true`                   |

#### `path (string)`

Path is the path of the accessed file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `protocol (string)`

Protocol is the protocol used to access the file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

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

| Method   | URL                       | Description                                           |
| -:       | -                         | -                                                     |
| `GET`    | `/fileaccesspolicies`     | Retrieves the list of file access policies.           |
| `POST`   | `/fileaccesspolicies`     | Creates a new file access policies.                   |
| `DELETE` | `/fileaccesspolicies/:id` | Deletes the `fileaccesspolicy` with the given `:id`.  |
| `GET`    | `/fileaccesspolicies/:id` | Retrieve the `fileaccesspolicy` with the given `:id`. |
| `PUT`    | `/fileaccesspolicies/:id` | Updates the `fileaccesspolicy` with the given `:id`.  |

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

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `allowsExecute (boolean)`

AllowsExecute allows to execute the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `allowsRead (boolean)`

AllowsRead allows to read the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `allowsWrite (boolean)`

AllowsWrite allows to write the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `encryptionEnabled (boolean)`

EncryptionEnabled will enable the automatic encryption.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `logsEnabled (boolean)`

LogsEnabled will enable logging when this policy is used.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:policies_list)`

Object is the object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

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

| Method   | URL              | Description                                   |
| -:       | -                | -                                             |
| `GET`    | `/filepaths`     | Retrieves the list of file path.              |
| `POST`   | `/filepaths`     | Create a new file path.                       |
| `DELETE` | `/filepaths/:id` | Deletes the `filepath` with the given `:id`.  |
| `GET`    | `/filepaths/:id` | Retrieve the `filepath` with the given `:id`. |
| `PUT`    | `/filepaths/:id` | Updates the `filepath` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `server (string)`

server is the server name/ID/IP associated with the file path.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## FlowStatistic

Returns network access statistics on a particular processing unit or group of
processing units based on their tags.

### Relations

| Method | URL               | Description                                            |
| -:     | -                 | -                                                      |
| `GET`  | `/flowstatistics` | Retrieves the flow statistics according to parameters. |

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

#### `dataPoints (external:datapoints_list)`

DataPoints is a list of time/value pairs that represent the flow events over
time.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `destinationIDs (external:flowstatistic_origin_list)`

DestinationIDs is the IDs of the destination.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `destinationTags (external:selectors_list)`

DestinationTags contains the tags used to identify destination.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `metric (enum)`

Metric is the kind of metric the statistic represents.

| Characteristics | Value          |
| -               | -:             |
| Allowed Value   | `Flows, Ports` |
| Default         | `Flows`        |
| Autogenerated   | `true`         |
| Read only       | `true`         |

#### `mode (enum)`

Mode defines if the metric is for accepted or rejected flows.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Accepted, Any, Rejected` |
| Default         | `Accepted`                |
| Autogenerated   | `true`                    |
| Read only       | `true`                    |

#### `sourceIDs (external:flowstatistic_origin_list)`

SourceIDs is the sources of the stats.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `sourceTags (external:selectors_list)`

SourceTags contains the tags used to identify the source.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `type (enum)`

Type is the type of representation.

| Characteristics | Value                |
| -               | -:                   |
| Allowed Value   | `Repartition, Serie` |
| Default         | `Serie`              |
| Autogenerated   | `true`               |
| Read only       | `true`               |

#### `userIdentifier (string)`

UserIdentifier can be set by the user as a query parameter. It will be returned
in the FlowStatistic object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Orderable       | `true` |

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

| Method   | URL                 | Description                                     |
| -:       | -                   | -                                               |
| `GET`    | `/hookpolicies`     | Retrieves the list of hook policies.            |
| `POST`   | `/hookpolicies`     | Creates a new hook policy.                      |
| `DELETE` | `/hookpolicies/:id` | Deletes the `hookpolicy` with the given `:id`.  |
| `GET`    | `/hookpolicies/:id` | Retrieve the `hookpolicy` with the given `:id`. |
| `PUT`    | `/hookpolicies/:id` | Updates the `hookpolicy` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `endpoint (string)`

Endpoint contains the full address of the remote processor endoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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
| Default         | `Pre`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

Subject contains the tag expression that an object must match in order to
trigger the hook.

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

## Import

Imports an export of policies and related objects into the namespace.

### Example

```json
{
  "data": "previous output of export"
}
```

### Relations

| Method | URL       | Description                          |
| -:     | -         | -                                    |
| `POST` | `/import` | Imports data from a previous export. |

### Attributes

#### `data (external:exported_data)`

The data to import.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `mode (enum)`

How to import the data.

| Characteristics | Value                                 |
| -               | -:                                    |
| Allowed Value   | `Append, ReplacePartial, ReplaceFull` |
| Default         | `Append`                              |
| Required        | `true`                                |

## Installation

Installation represents an installation for a given account.

### Relations

| Method   | URL                  | Description                                       |
| -:       | -                    | -                                                 |
| `GET`    | `/installations`     | Retrieves the list of installations.              |
| `POST`   | `/installations`     | Creates a new installation.                       |
| `DELETE` | `/installations/:id` | Deletes the `installation` with the given `:id`.  |
| `GET`    | `/installations/:id` | Retrieve the `installation` with the given `:id`. |
| `PUT`    | `/installations/:id` | Updates the `installation` with the given `:id`.  |

### Attributes

#### `ID (string)`

ID represents the identifier of the installation.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `accountName (string)`

AccountName that should be installed.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

## InstalledApp

InstalledApps represents an installed application.

### Relations

| Method   | URL                       | Description                                       |
| -:       | -                         | -                                                 |
| `GET`    | `/installedapps`          | Retrieves the list of installed apps.             |
| `POST`   | `/installedapps`          | Installs a new app.                               |
| `DELETE` | `/installedapps/:id`      | Deletes the `installedapp` with the given `:id`.  |
| `GET`    | `/installedapps/:id`      | Retrieve the `installedapp` with the given `:id`. |
| `PUT`    | `/installedapps/:id`      | Updates the `installedapp` with the given `:id`.  |
| `GET`    | `/installedapps/:id/logs` | Returns the logs for a app.                       |

### Attributes

#### `ID (string)`

ID of the installed app.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

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
| Filterable      | `true` |

#### `name (string)`

Name of the installed app.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespace (string)`

Namespace in which the app is running.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `parameters (external:app_parameters)`

Parameters is a list of parameters to start the app.

#### `replicas (integer)`

Replicas represents the number of replicas for the app.

#### `status (enum)`

Status of the app.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Error, Pending, Running` |
| Default         | `Pending`                 |
| Read only       | `true`                    |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

## IsolationProfile

An IsolationProfile needs documentation.

### Example

```json
{
  "name": "the name"
}
```

### Relations

| Method   | URL                      | Description                                           |
| -:       | -                        | -                                                     |
| `GET`    | `/isolationprofiles`     | Retrieves the list of isolation profiles.             |
| `POST`   | `/isolationprofiles`     | Creates a new isolation profile.                      |
| `DELETE` | `/isolationprofiles/:id` | Deletes the `isolationprofile` with the given `:id`.  |
| `GET`    | `/isolationprofiles/:id` | Retrieve the `isolationprofile` with the given `:id`. |
| `PUT`    | `/isolationprofiles/:id` | Updates the `isolationprofile` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity.

#### `capabilitiesActions (external:cap_map)`

CapabilitiesActions identifies the capabilities that should be added or removed
from the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `createTime (time)`

CreatedTime is the time at which the object was created.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

#### `defaultSyscallAction (external:syscall_action)`

DefaultAction is the default action applied to all syscalls of this profile.
Default is "Allow".

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

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `syscallRules (external:syscall_rules)`

SyscallRules is a list of syscall rules that identify actions for particular
syscalls.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `targetArchitectures (external:arch_list)`

TargetArchitectures is the target processor architectures where this profile can
be applied. Default all.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

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

| Method | URL      | Description         |
| -:     | -        | -                   |
| `POST` | `/issue` | Issues a new token. |

### Attributes

#### `data (string)`

Data contains additional data. The value depends on the issuer type.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata)`

Metadata contains various additional information. Meaning depends on the realm.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `realm (enum)`

Realm is the authentication realm.

| Characteristics | Value                                                                              |
| -               | -:                                                                                 |
| Allowed Value   | `AWSIdentityDocument, Certificate, Facebook, Github, Google, LDAP, Twitter, Vince` |
| Required        | `true`                                                                             |

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
| Default         | `24h`                                                                                                             |
| Orderable       | `true`                                                                                                            |
| Filterable      | `true`                                                                                                            |

## Jaegerbatch

A jaegerbatch is a batch of jaeger spans. This is used by external service to
post jaeger span in our private jaeger services.

### Attributes

#### `batch (external:jaeger_batch)`

Represents a jaeger batch.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

## KubernetesCluster

Create a remote Kubernetes Cluster integration.

### Relations

| Method   | URL                       | Description                                            |
| -:       | -                         | -                                                      |
| `GET`    | `/kubernetesclusters`     | Retrieves the list of kubernetes clusters.             |
| `POST`   | `/kubernetesclusters`     | Creates a new kubernetes cluster.                      |
| `DELETE` | `/kubernetesclusters/:id` | Deletes the `kubernetescluster` with the given `:id`.  |
| `GET`    | `/kubernetesclusters/:id` | Retrieve the `kubernetescluster` with the given `:id`. |
| `PUT`    | `/kubernetesclusters/:id` | Updates the `kubernetescluster` with the given `:id`.  |

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
| Filterable      | `true` |

#### `kubernetesDefinitions (string)`

base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `name (string)`

The name of your cluster.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `parentID (string)`

ID of the parent account.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `regenerate (boolean)`

Regenerates the k8s files and certificates.

#### `targetNamespace (string)`

The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `targetNetworks (external:target_networks_list)`

> This attribute is deprecated

List of target networks.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

Last update date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

## Log

Retrieves the log of a deployed app.

### Relations

| Method | URL                       | Description                 |
| -:     | -                         | -                           |
| `GET`  | `/installedapps/:id/logs` | Returns the logs for a app. |

### Attributes

#### `data (external:logs)`

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

| Method   | URL             | Description                                  |
| -:       | -               | -                                            |
| `GET`    | `/messages`     | Retrieves the list of messages.              |
| `POST`   | `/messages`     | Creates a new message.                       |
| `DELETE` | `/messages/:id` | Deletes the `message` with the given `:id`.  |
| `GET`    | `/messages/:id` | Retrieve the `message` with the given `:id`. |
| `PUT`    | `/messages/:id` | Updates the `message` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `level (enum)`

Level defines how the message is important.

| Characteristics | Value                   |
| -               | -:                      |
| Allowed Value   | `Danger, Info, Warning` |
| Default         | `Info`                  |
| Orderable       | `true`                  |
| Filterable      | `true`                  |

#### `local (boolean)`

If local is set, the message will only be visible in the current namespace.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

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

| Method   | URL               | Description                                    |
| -:       | -                 | -                                              |
| `GET`    | `/namespaces`     | Retrieves the list of namespaces.              |
| `POST`   | `/namespaces`     | Creates a new namespace.                       |
| `DELETE` | `/namespaces/:id` | Deletes the `namespace` with the given `:id`.  |
| `GET`    | `/namespaces/:id` | Retrieve the `namespace` with the given `:id`. |
| `PUT`    | `/namespaces/:id` | Updates the `namespace` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## NamespaceMappingPolicy

A Namespace Mapping Policy defines in which namespace a Processing Unit should
be placed when it is created, based on its tags.

When an Aporeto Agent creates a new Processing Unit, the system will place it in
its own namespace if no matching Namespace Mapping Policy can be found. If one
match is found, then the Processing will be bumped down to the namespace
declared in the policy. If it finds in that child namespace another matching
Namespace Mapping Policy, then the Processing Unit will be bumped down again,
until it reach a namespace with no matching policies.

This is very useful to dispatch processes and containers into a particular
namespace, based on a lot of factor.

You can put in place a quarantine namespace that will grab all Processing Units
with too much vulnerabilities for instances.

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

| Method   | URL                             | Description                                                 |
| -:       | -                               | -                                                           |
| `GET`    | `/namespacemappingpolicies`     | Retrieves the list namespace mapping policies.              |
| `POST`   | `/namespacemappingpolicies`     | Creates a new namespace mapping policy.                     |
| `DELETE` | `/namespacemappingpolicies/:id` | Deletes the `namespacemappingpolicy` with the given `:id`.  |
| `GET`    | `/namespacemappingpolicies/:id` | Retrieve the `namespacemappingpolicy` with the given `:id`. |
| `PUT`    | `/namespacemappingpolicies/:id` | Updates the `namespacemappingpolicy` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `mappedNamespace (string)`

mappedNamespace is the mapped namespace.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `subject (external:policies_list)`

Subject is the subject.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
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

| Method   | URL                                           | Description                                                                |
| -:       | -                                             | -                                                                          |
| `GET`    | `/networkaccesspolicies`                      | Retrieves the list of network access policies.                             |
| `POST`   | `/networkaccesspolicies`                      | Creates a new network access policy.                                       |
| `DELETE` | `/networkaccesspolicies/:id`                  | Deletes the `networkaccesspolicy` with the given `:id`.                    |
| `GET`    | `/networkaccesspolicies/:id`                  | Retrieve the `networkaccesspolicy` with the given `:id`.                   |
| `PUT`    | `/networkaccesspolicies/:id`                  | Updates the `networkaccesspolicy` with the given `:id`.                    |
| `GET`    | `/networkaccesspolicies/:id/externalservices` | Returns the list of external services affected by a network access policy. |
| `GET`    | `/networkaccesspolicies/:id/processingunits`  | Returns the list of Processing Units affected by a network access policy.  |

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
| Default         | `Allow`                   |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `destinationPorts (external:ports_list)`

DestinationPorts contains the list of allowed ports and ranges.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `encryptionEnabled (boolean)`

EncryptionEnabled defines if the flow has to be encrypted.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `logsEnabled (boolean)`

LogsEnabled defines if the flow has to be logged.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:policies_list)`

Object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

#### `observationEnabled (boolean)`

If set to true, the flow will be in observation mode.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `observedTrafficAction (enum)`

If observationEnabled is set to true, this will be the final action taken on the
packets.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Apply, Continue` |
| Default         | `Continue`        |
| Orderable       | `true`            |
| Filterable      | `true`            |

#### `passthrough (external:policies_list)`

List of tags expressions to match the list of entity to pass the flow through.

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

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

| Method | URL              | Description                                                 |
| -:     | -                | -                                                           |
| `GET`  | `/passwordreset` | Sends a link to the account email to reset the password.    |
| `POST` | `/passwordreset` | Resets the password for an account using the provided link. |

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

| Method | URL          | Description                               |
| -:     | -            | -                                         |
| `GET`  | `/plans`     | Retrieves the list of plans.              |
| `GET`  | `/plans/:id` | Retrieve the `plan` with the given `:id`. |

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

| Method | URL                   | Description                                                                      |
| -:     | -                     | -                                                                                |
| `GET`  | `/enforcers/:id/poke` | Sends a poke empty object. This is used to ensure an enforcer is up and running. |

## Policy

Policy represents the policy primitive used by all aporeto policies.

### Example

```json
{
  "name": "the name"
}
```

### Relations

| Method   | URL             | Description                                 |
| -:       | -               | -                                           |
| `GET`    | `/policies`     | Retrieves the list of policy primitives.    |
| `DELETE` | `/policies/:id` | Deletes the `policy` with the given `:id`.  |
| `GET`    | `/policies/:id` | Retrieve the `policy` with the given `:id`. |

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

#### `action (external:actions_list)`

Action defines set of actions that must be enforced when a dependency is met.

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `object (external:policies_list)`

Object represents set of entities that another entity depends on. As subjects,
objects are identified as logical operations on tags when a policy is defined.

#### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `relation (external:relations_list)`

Relation describes the required operation to be performed between subjects and
objects.

#### `subject (external:policies_list)`

Subject represent sets of entities that will have a dependency other entities.
Subjects are defined as logical operations on tags. Logical operations can
includes AND/OR.

#### `type (enum)`

Type of the policy.

| Characteristics | Value                                                                                                                           |
| -               | -:                                                                                                                              |
| Allowed Value   | `APIAuthorization, EnforcerProfile, File, Hook, NamespaceMapping, Network, ProcessingUnit, Quota, Service, Syscall, TokenScope` |
| Creation only   | `true`                                                                                                                          |
| Filterable      | `true`                                                                                                                          |

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

| Method | URL                | Description                                     |
| -:     | -                  | -                                               |
| `GET`  | `/policyrules/:id` | Retrieve the `policyrule` with the given `:id`. |

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

#### `action (external:actions_list)`

Action defines set of actions that must be enforced when a dependency is met.

#### `enforcerProfiles (external:enforcerprofiles_list)`

EnforcerProfiles provides the information about the server profile.

#### `externalServices (external:network_entities)`

Policy target networks.

#### `filePaths (external:file_entities)`

Policy target file paths.

#### `isolationProfiles (external:isolation_profile_entities)`

IsolationProfiles are the isolation profiles of the rule.

#### `name (string)`

Name is the name of the entity.

| Characteristics | Value  |
| -               | -:     |
| Max length      | `256`  |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

#### `namespaces (external:namespace_entities)`

Policy target namespaces.

#### `passthroughExternalServices (external:network_entities)`

List of external services the policy mandate to pass through before reaching the
destination.

#### `policyNamespace (string)`

PolicyNamespace is the namespace of the policy that created this rule.

#### `propagated (boolean)`

Propagated indicates if the policy is propagated.

#### `relation (external:relations_list)`

Relation describes the required operation to be performed between subjects and
objects.

#### `services (external:api_services_entities)`

Services provides the services of this policy rule.

#### `tagClauses (external:target_tags)`

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
  "name": "the name"
}
```

### Relations

| Method   | URL                                          | Description                                                               |
| -:       | -                                            | -                                                                         |
| `GET`    | `/processingunits`                           | Retrieves the list of processing units.                                   |
| `POST`   | `/processingunits`                           | Creates a new processing unit.                                            |
| `DELETE` | `/processingunits/:id`                       | Deletes the `processingunit` with the given `:id`.                        |
| `GET`    | `/processingunits/:id`                       | Retrieve the `processingunit` with the given `:id`.                       |
| `PUT`    | `/processingunits/:id`                       | Updates the `processingunit` with the given `:id`.                        |
| `GET`    | `/networkaccesspolicies/:id/processingunits` | Returns the list of Processing Units affected by a network access policy. |
| `GET`    | `/services/:id/processingunits`              | Retrieves the Processing Units that implement this service.               |
| `GET`    | `/vulnerabilities/:id/processingunits`       | Retrieves the processing units affected by the a vulnerabily.             |
| `GET`    | `/processingunits/:id/fileaccesses`          | Retrieves the file accesses done by the processing unit.                  |
| `GET`    | `/processingunits/:id/renderedpolicies`      | Retrieves the policies for the processing unit.                           |
| `GET`    | `/processingunits/:id/services`              | Retrieves the services used by a processing unit.                         |
| `GET`    | `/processingunits/:id/vulnerabilities`       | Retrieves the vulnerabilities affecting the processing unit.              |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `enforcerID (string)`

EnforcerID is the ID of the enforcer associated with the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `lastSyncTime (time)`

LastSyncTime is the time when the policy was last resolved.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Orderable       | `true` |

#### `metadata (external:metadata_list)`

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

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

#### `networkServices (external:processing_unit_services_list)`

NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `normalizedTags (external:tags_list)`

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
| Default         | `Initialized`                                       |
| Filterable      | `true`                                              |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `type (enum)`

Type of the container ecosystem.

| Characteristics | Value                             |
| -               | -:                                |
| Allowed Value   | `Docker, LinuxService, RKT, User` |
| Default         | `Docker`                          |
| Required        | `true`                            |
| Creation only   | `true`                            |
| Filterable      | `true`                            |

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

| Method   | URL                           | Description                                               |
| -:       | -                             | -                                                         |
| `GET`    | `/processingunitpolicies`     | Retrieves the list of processing unit policies.           |
| `POST`   | `/processingunitpolicies`     | Creates a new processing unit policy.                     |
| `DELETE` | `/processingunitpolicies/:id` | Deletes the `processingunitpolicy` with the given `:id`.  |
| `GET`    | `/processingunitpolicies/:id` | Retrieve the `processingunitpolicy` with the given `:id`. |
| `PUT`    | `/processingunitpolicies/:id` | Updates the `processingunitpolicy` with the given `:id`.  |

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
| Filterable      | `true`                                                   |

#### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `isolationProfileSelector (external:policies_list)`

IsolationProfileSelector are the profiles that must be applied when this policy
matches. Only applies to Enforce and LogCompliance actions.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

Subject defines the tag selectors that identitfy the processing units to which
this policy applies.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

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

| Method   | URL                  | Description                                      |
| -:       | -                    | -                                                |
| `GET`    | `/quotapolicies`     | Retrieves the list of quota policies.            |
| `POST`   | `/quotapolicies`     | Creates a new quota policy.                      |
| `DELETE` | `/quotapolicies/:id` | Deletes the `quotapolicy` with the given `:id`.  |
| `GET`    | `/quotapolicies/:id` | Retrieve the `quotapolicy` with the given `:id`. |
| `PUT`    | `/quotapolicies/:id` | Updates the `quotapolicy` with the given `:id`.  |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `identities (list)`

Identities contains the list of identity names where the quota will be applied.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `quota (integer)`

Quota contains the maximum number of object matching the policy subject that can
be created.

#### `targetNamespace (string)`

TargetNamespace contains the base namespace from where the count will be done.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

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
  "namespace": "/my/namespace",
  "operation": "create",
  "targetIdentity": "processingunit"
}
```

### Relations

| Method | URL                 | Description             |
| -:     | -                   | -                       |
| `POST` | `/remoteprocessors` | This should be be here. |

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
| Default         | `Pre`       |
| Required        | `true`      |

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

#### `output (external:elemental_identitifable)`

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
| Filterable      | `true` |

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

| Method | URL                                     | Description                                     |
| -:     | -                                       | -                                               |
| `POST` | `/renderedpolicies`                     | Render a policy for a processing unit.          |
| `GET`  | `/processingunits/:id/renderedpolicies` | Retrieves the policies for the processing unit. |

### Attributes

#### `certificate (string)`

Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `egressPolicies (external:rendered_policy)`

EgressPolicies lists all the egress policies attached to processing unit.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `exposedServices (external:api_services_entities)`

ExposedServices is the list of services that this processing unit is
implementing.

#### `ingressPolicies (external:rendered_policy)`

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

#### `processingUnit (external:processingunit)`

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

#### `profile (external:trust_profile)`

Profile is the trust profile of the processing unit that should be used during
all communications.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `scopes (external:scopes_list)`

Scopes is the set of scopes granted to this processing unit that it has to
present in HTTP requests.

## Report

Post a new statistics report.

### Relations

| Method | URL        | Description                 |
| -:     | -          | -                           |
| `POST` | `/reports` | Create a statistics report. |

### Attributes

#### `fields (external:tsdb_fields)`

TSDB Fields to set for the report.

#### `kind (enum)`

Kind contains the kind of report.

| Characteristics | Value                                                                |
| -               | -:                                                                   |
| Allowed Value   | `Audit, Enforcer, FileAccess, Flow, ProcessingUnit, Syscall, Claims` |

#### `tags (external:tags_map)`

Tags contains the tags associated to the data point.

#### `timestamp (time)`

Timestamp contains the time for the report.

#### `value (float)`

Value contains the value for the report.

## RESTAPISpec

RESTAPISpec descibes the REST APIs exposed by a service. These APIs
can be associated with one or more services.

### Example

```json
{
  "name": "the name"
}
```

### Relations

| Method   | URL                          | Description                                      |
| -:       | -                            | -                                                |
| `GET`    | `/restapispecs`              | Retrieves the list of REST API specifications.   |
| `POST`   | `/restapispecs`              | Creates a new REST API specification.            |
| `DELETE` | `/restapispecs/:id`          | Deletes the `restapispec` with the given `:id`.  |
| `GET`    | `/restapispecs/:id`          | Retrieve the `restapispec` with the given `:id`. |
| `PUT`    | `/restapispecs/:id`          | Updates the `restapispec` with the given `:id`.  |
| `GET`    | `/services/:id/restapispecs` | Retrieves the REST APIs exposed by this service. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `endpoints (external:exposed_api_list)`

EndPoints is a list of API endpoints that are exposed for the service.

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Role

Roles returns the available roles that can be used with API Authorization
Policies.

### Relations

| Method | URL      | Description                           |
| -:     | -        | -                                     |
| `GET`  | `/roles` | Retrieves the list of existing roles. |

### Attributes

#### `authorizations (external:authorization_map)`

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
  "exposedAPIs": [
    [
      "package=p1"
    ]
  ],
  "exposedPort": 443,
  "name": "the name",
  "selectors": [
    [
      "$identity=processingunit"
    ]
  ]
}
```

### Relations

| Method   | URL                             | Description                                                 |
| -:       | -                               | -                                                           |
| `GET`    | `/services`                     | Retrieves the list of Services.                             |
| `POST`   | `/services`                     | Creates a new Service.                                      |
| `DELETE` | `/services/:id`                 | Deletes the `service` with the given `:id`.                 |
| `GET`    | `/services/:id`                 | Retrieve the `service` with the given `:id`.                |
| `PUT`    | `/services/:id`                 | Updates the `service` with the given `:id`.                 |
| `GET`    | `/processingunits/:id/services` | Retrieves the services used by a processing unit.           |
| `GET`    | `/services/:id/processingunits` | Retrieves the Processing Units that implement this service. |
| `GET`    | `/services/:id/restapispecs`    | Retrieves the REST APIs exposed by this service.            |

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

#### `IPs (external:ip_list)`

IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from  host names otherwise.

#### `JWTSigningCertificate (string)`

JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `endpoints (external:exposed_api_list)`

Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

#### `exposedAPIs (external:policies_list)`

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
| Min length      | `1`     |
| Max length      | `65535` |
| Required        | `true`  |

#### `external (boolean)`

External is a boolean that indicates if this is an external service.

| Characteristics | Value             |
| -               | -:                |
| Default         | `%!s(bool=false)` |
| Orderable       | `true`            |
| Filterable      | `true`            |

#### `hosts (list)`

Hosts are the names that the service can be accessed with.

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

#### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `port (integer)`

Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.

| Characteristics | Value         |
| -               | -:            |
| Default         | `%!s(int=80)` |
| Min length      | `1`           |
| Max length      | `65535`       |
| Required        | `true`        |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `selectors (external:policies_list)`

Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.

#### `serviceCA (string)`

ServiceCA  is the certificate authority that the service is using. This
is needed for external services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.

#### `type (enum)`

Type is the type of the service.

| Characteristics | Value       |
| -               | -:          |
| Allowed Value   | `HTTP, TCP` |
| Default         | `HTTP`      |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## StatsQuery

StatsQuery is a generic API to retrieve time series data stored by the Aporeto
system. The API allows different types of queries that are all protected within
the namespace of the user.

### Relations

| Method | URL             | Description                                           |
| -:     | -               | -                                                     |
| `GET`  | `/statsqueries` | Retrieves statistics information based on parameters. |

### Attributes

#### `results (external:time_series_results)`

Results contains the result of the query.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

## SuggestedPolicy

Allows to get policy suggestions.

### Relations

| Method | URL                  | Description                                    |
| -:     | -                    | -                                              |
| `GET`  | `/suggestedpolicies` | Retrieves a list of network policy suggestion. |

### Attributes

#### `networkAccessPolicies (external:network_access_policies_list)`

List of suggested network access policies.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

## Tabulation

Tabulate API allows you to retrieve a custom table view for any identity using
any tags you like as columns.

### Relations

| Method | URL            | Description                                           |
| -:     | -              | -                                                     |
| `GET`  | `/tabulations` | Retrieves tabulated informations based on parameters. |

### Attributes

#### `headers (list)`

Headers contains the requests headers that matched.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

#### `rows (external:tabulated_data)`

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

| Method | URL     | Description                                        |
| -:     | -       | -                                                  |
| `GET`  | `/tags` | Retrieves the list of existing tags in the system. |

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

| Method   | URL                       | Description                                           |
| -:       | -                         | -                                                     |
| `GET`    | `/tokenscopepolicies`     | Retrieves the list of token scope policies.           |
| `POST`   | `/tokenscopepolicies`     | Creates a new token scope policy.                     |
| `DELETE` | `/tokenscopepolicies/:id` | Deletes the `tokenscopepolicy` with the given `:id`.  |
| `GET`    | `/tokenscopepolicies/:id` | Retrieve the `tokenscopepolicy` with the given `:id`. |
| `PUT`    | `/tokenscopepolicies/:id` | Updates the `tokenscopepolicy` with the given `:id`.  |

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

#### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `assignedScopes (external:tags_list)`

AssignedScopes is the the list of scopes that the policiy will assigns.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `associatedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `fallback (boolean)`

Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `metadata (external:metadata_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `subject (external:policies_list)`

Subject defines the selection criteria that this policy must match on identiy
and scope request information.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

## Trigger

Trigger can be used to remotely trigger an automation.

### Relations

| Method         | URL                         | Description |
| -:             | -                           | -           |
| `GET`          | `/automations/:id/triggers` | Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`. |
| `POST`         | `/automations/:id/triggers` | Allows a system to trigger the automation if its `triggerType` property is set
to `RemoteCall`. |

## Vulnerability

A vulnerabily represents a particular CVE.

### Example

```json
{
  "link": "https://cve.com/CVE-1234",
  "name": "the name",
  "severity": 3
}
```

### Relations

| Method | URL                                    | Description                                                   |
| -:     | -                                      | -                                                             |
| `GET`  | `/vulnerabilities`                     | Retrieves the list of vulnerabilities.                        |
| `POST` | `/vulnerabilities`                     | Creates a new vulnerability.                                  |
| `GET`  | `/vulnerabilities/:id`                 | Retrieve the `vulnerability` with the given `:id`.            |
| `GET`  | `/processingunits/:id/vulnerabilities` | Retrieves the vulnerabilities affecting the processing unit.  |
| `GET`  | `/vulnerabilities/:id/processingunits` | Retrieves the processing units affected by the a vulnerabily. |

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

#### `annotations (external:annotations)`

Annotation stores additional information about an entity.

#### `associatedTags (external:tags_list)`

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

#### `normalizedTags (external:tags_list)`

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
| Filterable      | `true` |

#### `severity (external:vulnerability_level)`

Severity refers to the security vulnerability level.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Filterable      | `true` |

#### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
