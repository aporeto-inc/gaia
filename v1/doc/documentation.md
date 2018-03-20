<!-- markdownlint-disable MD024 MD025 -->

# Gaia API Documentation

> Version: 1

| Object                                                        | Description                                                                         |
| -                                                             | -                                                                                   |
| [Account](#account)                                           | This api allows to view and manage basic information about your account like        |
| [AccountCheck](#accountcheck)                                 | [nodoc]                                                                             |
| [Activate](#activate)                                         | Used to activate a pending account                                                  |
| [Activity](#activity)                                         | Contains all the activity log that happened in a namespace. All successful or       |
| [Alarm](#alarm)                                               | An alarm represents an event requiring attention.                                   |
| [APIAuthorizationPolicy](#apiauthorizationpolicy)             | An API Authorization Policy defines what kind of operations a user of a system      |
| [APICheck](#apicheck)                                         | This API allows to verify is a client identitied by his token is allowed to do      |
| [APIService](#apiservice)                                     | APIService descibes a L4/L7 service and the corresponding implementation. It        |
| [AuditProfile](#auditprofile)                                 | AuditProfile is an audit policy that consists of a set of audit rules. An audit     |
| [Auth](#auth)                                                 | This API verifies if the given token is valid or not.                               |
| [Automation](#automation)                                     | An automation needs documentation.                                                  |
| [AutomationTemplate](#automationtemplate)                     | Templates that ca be used in automations                                            |
| [AvailableService](#availableservice)                         | AvailableService represents a service that is available for launching               |
| [AWSAccount](#awsaccount)                                     | Allows to bind an AWS account to your Aporeto account to allow auto registration... |
| [Category](#category)                                         | Category allows to categorized services                                             |
| [Certificate](#certificate)                                   | A User represents the owner of some certificates.                                   |
| [DependencyMap](#dependencymap)                               | This api returns a data structure representing the graph of all processing units... |
| [Enforcer](#enforcer)                                         | An Enforcer Profile contains a configuration for a Enforcer. It contains various... |
| [EnforcerProfile](#enforcerprofile)                           | Allows to create reusable configuration profile for your enforcers. Enforcer        |
| [EnforcerProfileMappingPolicy](#enforcerprofilemappingpolicy) | A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used... |
| [Export](#export)                                             | Export the policies and related objects in a given namespace.                       |
| [ExternalAccess](#externalaccess)                             | ExternalAccess allows to retrieve connection from or to an external service         |
| [ExternalService](#externalservice)                           | An External Service represents a random network or ip that is not managed by the... |
| [FileAccess](#fileaccess)                                     | Returns file access statistics on a particular processing unit.                     |
| [FileAccessPolicy](#fileaccesspolicy)                         | A File Access Policy allows Processing Units to access various folder and files.... |
| [FilePath](#filepath)                                         | A File Path represents a random path to a file or a folder. They can be used in     |
| [FlowStatistic](#flowstatistic)                               | Returns network access statistics on a particular processing unit or group of       |
| [HookPolicy](#hookpolicy)                                     | Hook allows to to define hooks to the write operations in squall. Hooks are sent... |
| [Import](#import)                                             | Imports an export of policies and related objects into the namespace.               |
| [Installation](#installation)                                 | Installation represents an installation for a given account                         |
| [IsolationProfile](#isolationprofile)                         | An IsolationProfile needs documentation                                             |
| [Issue](#issue)                                               | This API issues a new token according to given data.                                |
| [Jaegerbatch](#jaegerbatch)                                   | A jaegerbatch is a batch of jaeger spans. This is used by external service to       |
| [KubernetesCluster](#kubernetescluster)                       | Create a remote Kubernetes Cluster integration.                                     |
| [Log](#log)                                                   | Retrieves the log of a deployed app.                                                |
| [Message](#message)                                           | The Message API allows to post public messages that will be visible through all     |
| [Namespace](#namespace)                                       | A Namespace represents the core organizational unit of the system. All objects      |
| [NamespaceMappingPolicy](#namespacemappingpolicy)             | A Namespace Mapping Policy defines in which namespace a Processing Unit should      |
| [NetworkAccessPolicy](#networkaccesspolicy)                   | Allows to define networking policies to allow or prevent processing units           |
| [PasswordReset](#passwordreset)                               | Used to reset an account password.                                                  |
| [Plan](#plan)                                                 | Plan contains the various billing plans available                                   |
| [Poke](#poke)                                                 | When available, poke can be used to update various information about the parent.... |
| [Policy](#policy)                                             | [nodoc]                                                                             |
| [PolicyRefresh](#policyrefresh)                               | PolicyRefresh is sent to client when as a push event when a policy refresh is       |
| [PolicyRule](#policyrule)                                     | PolicyRule is an internal policy resolution API. Services can use this API to       |
| [ProcessingUnit](#processingunit)                             | A Processing Unit reprents anything that can compute. It can be a Docker            |
| [ProcessingUnitPolicy](#processingunitpolicy)                 | A ProcessingUnitPolicies needs a better description.                                |
| [QuotaPolicy](#quotapolicy)                                   | Quotas Policies allows to set quotas on the number of objects that can be           |
| [RemoteProcessor](#remoteprocessor)                           | Hook to integrate an Aporeto service.                                               |
| [RenderedPolicy](#renderedpolicy)                             | Retrieve the aggregated policies applied to a particular processing unit.           |
| [Report](#report)                                             | Post a new statistics report.                                                       |
| [Revocation](#revocation)                                     | Used to revoke a certificate                                                        |
| [Role](#role)                                                 | Roles returns the available roles that can be used with API Authorization           |
| [Root](#root)                                                 | [nodoc]                                                                             |
| [Service](#service)                                           | Service represents a service that can be launched                                   |
| [StatsQuery](#statsquery)                                     | StatsQuery is a generic API to retrieve time series data stored by the Aporeto      |
| [SuggestedPolicy](#suggestedpolicy)                           | Allows to get policy suggestions                                                    |
| [SystemCall](#systemcall)                                     | [nodoc]                                                                             |
| [Tabulation](#tabulation)                                     | Tabulate API allows you to retrieve a custom table view for any identity using      |
| [Tag](#tag)                                                   | A tag is a string in the form of "key=value" that can applied to all objects in     |
| [TokenScopePolicy](#tokenscopepolicy)                         | The TokenScopePolicy defines a set of policies that allow customization of the      |
| [Trigger](#trigger)                                           | Trigger can be used to remotely trigger an automation.                              |
| [Vulnerability](#vulnerability)                               | A vulnerabily represents a particular CVE                                           |

# Account

> Operations: `GET` `UPDATE` `DELETE`

This api allows to view and manage basic information about your account like
your name, password, enable 2 factor authentication.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPAddress (string)`

LDAPAddress holds the account authentication account's private ldap server.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPBaseDN (string)`

LDAPBaseDN holds the base DN to use to ldap queries.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPBindDN (string)`

LDAPBindDN holds the account's internal LDAP bind dn for querying auth.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPBindPassword (string)`

LDAPBindPassword holds the password to the LDAPBindDN.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPCertificateAuthority (string)`

LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `LDAPEnabled (boolean)`

LDAPEnabled triggers if the account uses it's own LDAP for authentication.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `OTPEnabled (boolean)`

Set to enable or disable two factor authentication.

| Characteristics | Value |
| -               | -:    |

### `OTPQRCode (string)`

Returns the base64 encoded QRCode for setting up 2 factor auth.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `accessEnabled (boolean)`

AccessEnabled defines if the account holder should have access to the systems.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `associatedPlanKey (string)`

AssociatedPlanKey contains the plan key that is associated to this account.

| Characteristics | Value               |
| -               | -:                  |
| Default         | `aporeto.plan.free` |
| Autogenerated   | `true`              |
| Read only       | `true`              |

### `company (string)`

Company of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `email (string)`

Email of the account holder.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `firstName (string)`

First Name of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `lastName (string)`

Last Name of the account user.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

Name of the account.

| Characteristics | Value          |
| -               | -:             |
| Format          | `/^[^\*\=]*$/` |
| Required        | `true`         |
| Creation only   | `true`         |
| Orderable       | `true`         |
| Filterable      | `true`         |

### `password (string)`

Password for the account.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `reCAPTCHAKey (string)`

ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `status (enum)`

Status of the account.

| Characteristics | Value                                |
| -               | -:                                   |
| Allowed Value   | `Active, Disabled, Invited, Pending` |
| Default         | `Pending`                            |
| Autogenerated   | `true`                               |
| Read only       | `true`                               |
| Orderable       | `true`                               |
| Filterable      | `true`                               |

### `updateTime (time)`

Last update date of the object

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# AccountCheck

> Operations:

[nodoc]

# Activate

> Operations:

Used to activate a pending account

## Attributes

### `token (string)`

Token contains the activation token

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

# Activity

> Operations: `GET`

Contains all the activity log that happened in a namespace. All successful or
failed actions will be available, and eventual errors as well as the claims of
the user who triggered the actiions. This log is capped and only keeps the last
50k entries by default.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `claims (external:raw_data)`

Claims of the user who performed the operation.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `data (external:raw_data)`

Data of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `date (time)`

Date of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `error (external:raw_data)`

Error contains the eventual error.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `message (string)`

Message of the notification.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace of the notification.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `operation (string)`

Operation describe what kind of operation the notification represents.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

### `originalData (external:raw_data)`

OriginalData contains the eventual original data of the object that has been
modified.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `source (string)`

Source contains meta information about the source.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `targetIdentity (string)`

TargetIdentity is the Identity of the related object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# Alarm

> Operations: `GET` `UPDATE` `DELETE`

An alarm represents an event requiring attention.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `content (string)`

Content of the alarm.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `data (external:alarm_data)`

Data represent user data related to the alams

| Characteristics | Value |
| -               | -:    |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `kind (string)`

Kind identifies the kind of alarms. If two alarms are created with the same
identifier, then only the occurrence will be incremented.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `occurrences (external:alarm_occurrences)`

Number of time this alarm have been seen.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Creation only   | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `status (enum)`

Status of the alarm

| Characteristics | Value                          |
| -               | -:                             |
| Allowed Value   | `Acknowledged, Open, Resolved` |
| Default         | `Open`                         |
| Orderable       | `true`                         |
| Filterable      | `true`                         |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# APIAuthorizationPolicy

> Operations: `GET` `UPDATE` `DELETE`

An API Authorization Policy defines what kind of operations a user of a system
can do in a namespace. The operations can be any combination of GET, POST, PUT,
DELETE,PATCH or HEAD. By default, an API Authorization Policy will only give
permissions in the context of the current namespace but you can make it
propagate to all the child namespaces. It is also possible restrict permissions
to apply only on a particular subset of the apis by setting the target
identities.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `authorizedIdentities (external:identity_list)`

AuthorizedIdentities defines the list of api identities the policy applies to.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `authorizedNamespace (string)`

AuthorizedNamespace defines on what namespace the policy applies.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject is the subject.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# APICheck

> Operations:

This API allows to verify is a client identitied by his token is allowed to do
some operations on some apis. For example, it allows third party system to
impersonate a user and ensure a proxfied request should be allowed.

## Attributes

### `authorized (external:authorized_identities)`

Authorized contains the results of the check.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `namespace (string)`

Namespace is the namespace to use to check the api authentication.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `operation (enum)`

Operation is the operation you want to check.

| Characteristics | Value                                                         |
| -               | -:                                                            |
| Allowed Value   | `Create, Delete, Info, Patch, Retrieve, RetrieveMany, Update` |
| Orderable       | `true`                                                        |
| Filterable      | `true`                                                        |

### `targetIdentities (external:identity_list)`

TargetIdentities contains the list of identities you want to check the
authorization.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `token (string)`

Token is the token to use to check api authentication

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

# APIService

> Operations: `GET` `UPDATE` `DELETE`

APIService descibes a L4/L7 service and the corresponding implementation. It
allows users to define their services, the APIs that they expose, the
implementation of the service. These definitions can be used by network policy
in order to define advanced controls based on the APIs.

## Attributes

### `FQDN (string)`

FQDN is the fully qualified domain name of the service. It is required for
external API services. It can be deduced from a service discovery system in
advanced environments.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `IPList (external:ip_list)`

IPList is the list of ip address or subnets of the service if available.

| Characteristics | Value |
| -               | -:    |

### `JWTSigningCertificate (string)`

JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.

| Characteristics | Value |
| -               | -:    |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `exposedAPIs (external:exposed_api_list)`

ExposedAPIs is a list of API endpoints that are exposed for the service.

| Characteristics | Value |
| -               | -:    |

### `external (boolean)`

External is a boolean that indicates if this is an external service.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |
| Filterable      | `true`  |

### `externalServiceCA (string)`

ExternalServiceCA is the certificate authority that the service is using. This
is needed for external API services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.

| Characteristics | Value |
| -               | -:    |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `networkProtocol (integer)`

NetworkProtocol is the network protocol of the service.

| Characteristics | Value  |
| -               | -:     |
| Default         | `6`    |
| Min length      | `1`    |
| Max length      | `255`  |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `ports (external:port_list)`

Ports is a list of ports for the service. Ports are either exact match, or a
range portMin:portMax.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `runtimeSelectors (external:target_tags)`

RuntimeSelectors is a list of tag selectors that identifies that Processing
Units that will implement this service.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `type (enum)`

Type is the type of the service (HTTP, TCP, etc). More types will be added to
the system.

| Characteristics | Value           |
| -               | -:              |
| Allowed Value   | `HTTP, L3, TCP` |
| Default         | `L3`            |
| Required        | `true`          |
| Orderable       | `true`          |
| Filterable      | `true`          |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# AuditProfile

> Operations: `GET` `UPDATE` `DELETE`

AuditProfile is an audit policy that consists of a set of audit rules. An audit
policy will determine that types of events that must be captured in the kernel.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagated (boolean)`

Propagated indicates if the audit profile is propagated.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `rules (external:audit_profile_rule_list)`

Rules is the list of audit policy rules associated with this policy.

| Characteristics | Value |
| -               | -:    |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Auth

> Operations:

This API verifies if the given token is valid or not.

## Attributes

### `claims (external:claims)`

Claims are the claims.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# Automation

> Operations: `GET` `UPDATE` `DELETE`

An automation needs documentation.

## Relations

Relation with [Trigger](#triggers)

- `GET  /automations/id/triggers`
- `POST /automations/id/triggers`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `actions (list)`

Action contains the code that will be executed if the condition is met.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `condition (string)`

Condition contains the code that will be executed to decide if any action should
be taken.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `entitlements (external:automation_entitlements)`

Entitlements declares which operations are allowed on which identities.

| Characteristics | Value |
| -               | -:    |

### `errors (list)`

Error contains the eventual error of the last run.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `events (external:automation_events)`

Events contains the identity and operation an event must have to trigger the
automation

| Characteristics | Value |
| -               | -:    |

### `lastExecTime (time)`

LastExecTime holds the last successful execution tine.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `parameters (external:automation_parameters)`

Parameters are passed to the functions.

| Characteristics | Value |
| -               | -:    |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `schedule (string)`

Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.

| Characteristics | Value |
| -               | -:    |

### `stdout (string)`

Stdout contains the last run standard output.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `token (string)`

Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Creation only   | `true` |

### `tokenRenew (boolean)`

If set to true a new token will be issued, and the previous one invalidated.

| Characteristics | Value |
| -               | -:    |

### `trigger (enum)`

Trigger controls when the automation should be triggered.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Event, RemoteCall, Time` |
| Default         | `Time`                    |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# AutomationTemplate

> Operations: `GET`

Templates that ca be used in automations

## Attributes

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `entitlements (external:automation_entitlements)`

Entitlements contains the entitlements needed for executing the function.

| Characteristics | Value |
| -               | -:    |

### `function (string)`

Function contains the code.

| Characteristics | Value |
| -               | -:    |

### `key (string)`

Key contains the unique identifier key for the template.

| Characteristics | Value |
| -               | -:    |

### `kind (enum)`

Kind represents the kind of template.

| Characteristics | Value               |
| -               | -:                  |
| Allowed Value   | `Action, Condition` |
| Default         | `Condition`         |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `parameters (external:automation_template_parameters)`

Parameters contains the parameter description of the function.

| Characteristics | Value |
| -               | -:    |

# AvailableService

> Operations:

AvailableService represents a service that is available for launching

## Attributes

### `beta (boolean)`

Beta indicates if the service is in a beta version.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

### `categoryID (string)`

CategoryID of the service.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Filterable      | `true` |

### `description (string)`

Description of the service

| Characteristics | Value |
| -               | -:    |

### `icon (string)`

Icon contains a base64 image for the available service.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

### `longDescription (string)`

LongDescription contains a more detailed description of the service.

| Characteristics | Value |
| -               | -:    |

### `name (string)`

Name of the Service

| Characteristics | Value |
| -               | -:    |

### `parameters (external:service_parameters)`

Parameters of the service the user can or has to specify

| Characteristics | Value |
| -               | -:    |

### `title (string)`

Title represents the title of the service.

| Characteristics | Value |
| -               | -:    |

# AWSAccount

> Operations: `GET` `DELETE`

Allows to bind an AWS account to your Aporeto account to allow auto registration
of enforcers running on EC2

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `accessKeyID (string)`

AccessKeyID contains the aws access key ID. This is used to retrieve your
account id, and it is not stored.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

### `accessToken (string)`

AccessToken contains your aws access token. It is used to retrieve your account
id, and it not stored.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `accountID (string)`

accountID contains your verified accound id.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `parentID (string)`

ParentID contains the parent Vince account ID.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

### `parentName (string)`

ParentName contains the name of the Vince parent Account.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

### `region (string)`

Region contains your the region where your AWS account is located

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `secretAccessKey (string)`

secretAccessKey contains the secret key. It is used to retrieve your account id,
and it is not stored.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

### `updateTime (time)`

Last update date of the object

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# Category

> Operations:

Category allows to categorized services

## Attributes

### `ID (string)`

ID is the identifier of the category.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `description (string)`

Description is the desription of the category.

| Characteristics | Value |
| -               | -:    |

### `name (string)`

Name of the category.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

# Certificate

> Operations: `GET` `UPDATE` `DELETE`

A User represents the owner of some certificates.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `admin (boolean)`

Admin determines if the certificate must be added to the admin list.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `commonName (string)`

CommonName (CN) for the user certificate

| Characteristics | Value  |
| -               | -:     |
| Max length      | `64`   |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `data (string)`

Certificate provides a certificate for the user

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `email (string)`

e-mail address of the user

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `expirationDate (time)`

CertificateExpirationDate indicates the expiration day for the certificate.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `key (string)`

CertificateKey provides the key for the user. Only available at create or update
time.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `name (string)`

Name of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `organizationalUnits (list)`

OrganizationalUnits attribute for the generated certificates

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `parentID (string)`

ParentID holds the parent account ID.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `passphrase (string)`

Passphrase to use for the generated p12.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `serialNumber (string)`

SerialNumber of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `status (enum)`

CertificateStatus provides the status of the certificate. Update with RENEW to
get a new certificate.

| Characteristics | Value            |
| -               | -:               |
| Allowed Value   | `Revoked, Valid` |
| Default         | `Valid`          |
| Orderable       | `true`           |
| Filterable      | `true`           |

### `updateTime (time)`

Last update date of the object

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# DependencyMap

> Operations:

This api returns a data structure representing the graph of all processing units
and their connections in a particular namespace, in a given time window. To pass
the time window you can use the query parameters 'startAbsolute', 'endAbsolute',
'startRelative', 'endRelative'.

For example
  "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000"

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `edges (external:graphedges_map)`

edges are the edges of the map

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Read only       | `true` |

### `groups (external:graphgroups_map)`

Groups provide information about the group values

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Read only       | `true` |

### `nodes (external:graphnodes_map)`

nodes refers to the nodes of the map

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Read only       | `true` |

### `viewSuggestions (external:view_suggestions)`

viewSuggestions provides suggestion of views based on relevant tags.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Read only       | `true` |

# Enforcer

> Operations: `GET` `UPDATE` `DELETE`

An Enforcer Profile contains a configuration for a Enforcer. It contains various
parameters, like what network should not policeds, what processing units should
be ignored based on their tags and so on. It also contains more advanced
parameters to fine tune the Agent. A Enforcer will decide what profile to apply
using aEnforcer Profile Mapping Policy. This policy will decide according the
Enforcer's tags what profile to use. If an Enforcer tags are matching more than
a single policy, it will refuse to start. Some parameters will be applied
directly to a running agent, some will need to restart it.

## Relations

Relation with [EnforcerProfile](#enforcerprofiles)

- `GET  /enforcers/id/enforcerprofiles`

Relation with [Poke](#poke)

- `GET  /enforcers/id/poke`

## Attributes

### `FQDN (string)`

FQDN contains the fqdn of the server where the enforcer is running.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `certificate (string)`

Certificate is the certificate of the enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `certificateExpirationDate (time)`

CertificateExpirationDate is the expiration date of the certificate.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `certificateKey (string)`

CertificateKey is the secret key of the enforcer. Returned only when a enforcer
is created or the certificate is updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `certificateRequest (string)`

CertificateRequest can be used to generate a certificate from that CSR instead
of letting the server generate your private key for you.

| Characteristics | Value |
| -               | -:    |

### `certificateRequestEnabled (boolean)`

If set during creation,the server will not initially generate a certificate. In
that case you have to provide a valid CSR through certificateRequest during an
update.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `collectInfo (boolean)`

CollectInfo indicates to the enforcer it needs to collect information.

| Characteristics | Value |
| -               | -:    |

### `collectedInfo (external:collected_info)`

CollectedInfo represents the latest info collected by the enforcer.

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `currentVersion (string)`

CurrentVersion holds the enforcerd binary version that is currently associated
to this object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `enforcerProfileID (string)`

Contains the ID of the profile used by the instance of enforcerd.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `lastCollectionTime (time)`

LastCollectionTime represents the date and time when the info have been
collected.

| Characteristics | Value |
| -               | -:    |

### `lastSyncTime (time)`

LastSyncTime holds the last heart beat time.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `localCA (string)`

LocalCA contains the initial chain of trust for the enforcer. This valud is only
given when you retrieve a single enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `operationalStatus (enum)`

OperationalStatus tells the status of the enforcer.

| Characteristics | Value                                           |
| -               | -:                                              |
| Allowed Value   | `Connected, Disconnected, Initialized, Unknown` |
| Default         | `Initialized`                                   |
| Autogenerated   | `true`                                          |
| Read only       | `true`                                          |
| Filterable      | `true`                                          |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `publicToken (string)`

PublicToken is the public token of the server that will be included in the
datapath and its signed by the private CA.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `updateAvailable (boolean)`

Tells if the the version of enforcerd is outdated and should be updated.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# EnforcerProfile

> Operations: `GET` `UPDATE` `DELETE`

Allows to create reusable configuration profile for your enforcers. Enforcer
Profiles contains various startup information that can (for some) be updated
live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
Mapping Policy.

## Relations

Relation with [AuditProfile](#auditprofiles)

- `GET  /enforcerprofiles/id/auditprofiles`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `IPTablesMarkValue (integer)`

IptablesMarkValue is the mark value to be used in an iptables implementation.

| Characteristics | Value   |
| -               | -:      |
| Default         | `1000`  |
| Max length      | `65000` |
| Orderable       | `true`  |
| Filterable      | `true`  |

### `PUBookkeepingInterval (string)`

PUBookkeepingInterval configures how often the PU will be synchronized.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `15m`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

### `PUHeartbeatInterval (string)`

PUHeartbeatInterval configures the heart beat interval.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `5s`              |
| Orderable       | `true`            |
| Filterable      | `true`            |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `auditProfileSelectors (external:audit_profile_selector)`

AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

### `auditProfiles (external:audit_profiles)`

AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `auditSocketBufferSize (integer)`

AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.

| Characteristics | Value    |
| -               | -:       |
| Default         | `16384`  |
| Max length      | `262144` |
| Orderable       | `true`   |
| Filterable      | `true`   |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `dockerSocketAddress (string)`

DockerSocketAddress is the address of the docker daemon.

| Characteristics | Value                  |
| -               | -:                     |
| Default         | `/var/run/docker.sock` |
| Orderable       | `true`                 |
| Filterable      | `true`                 |

### `dockerSocketType (enum)`

DockerSocketType is the type of socket to use to talk to the docker daemon.

| Characteristics | Value       |
| -               | -:          |
| Allowed Value   | `TCP, Unix` |
| Default         | `Unix`      |
| Orderable       | `true`      |
| Filterable      | `true`      |

### `excludedInterfaces (external:excluded_interfaces_list)`

ExcludedInterfaces is a list of interfaces that must be excluded.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `excludedNetworks (external:excluded_networks_list)`

ExcludedNetworks is the list of networks that must be excluded for this
enforcer.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `hostServices (external:host_services_list)`

HostServices is a list of services that must be activated by default to all
enforcers matching this profile.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `ignoreExpression (external:policies_list)`

IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.

| Characteristics | Value |
| -               | -:    |

### `kubernetesSupportEnabled (boolean)`

KubernetesSupportEnabled enables kubernetes mode for the enforcer.

| Characteristics | Value   |
| -               | -:      |
| Default         | `false` |
| Orderable       | `true`  |
| Filterable      | `true`  |

### `linuxProcessesSupportEnabled (boolean)`

LinuxProcessesSupportEnabled configures support for Linux processes.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |

### `metadataExtractor (enum)`

Select which metadata extractor to use to process new processing units.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Docker, ECS, Kubernetes` |
| Default         | `Docker`                  |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `policySynchronizationInterval (string)`

PolicySynchronizationInterval configures how often the policy will be
resynchronized.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |
| Default         | `10m`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `proxyListenAddress (string)`

ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
(example: unix:/var/run/aporeto.sock)

| Characteristics | Value                                                                                                                                     |
| -               | -:                                                                                                                                        |
| Format          | `/^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix:(/[^/]{1,16}){1,5}/?)$/` |
| Default         | `unix:/var/run/aporeto.sock`                                                                                                              |
| Orderable       | `true`                                                                                                                                    |
| Filterable      | `true`                                                                                                                                    |

### `receiverNumberOfQueues (integer)`

ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |
| Filterable      | `true` |

### `receiverQueue (integer)`

ReceiverQueue is the base queue number for traffic from the network.

| Characteristics | Value  |
| -               | -:     |
| Default         | `0`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

### `receiverQueueSize (integer)`

ReceiverQueueSize is the queue size of the receiver

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `5000` |
| Orderable       | `true` |
| Filterable      | `true` |

### `remoteEnforcerEnabled (boolean)`

RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.

| Characteristics | Value  |
| -               | -:     |
| Default         | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `targetNetworks (external:target_networks_list)`

TargetNetworks is the list of networks that authorization should be applied.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `transmitterNumberOfQueues (integer)`

TransmitterNumberOfQueues is the number of queues for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `16`   |
| Orderable       | `true` |
| Filterable      | `true` |

### `transmitterQueue (integer)`

TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue

| Characteristics | Value  |
| -               | -:     |
| Default         | `4`    |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

### `transmitterQueueSize (integer)`

TransmitterQueueSize is the size of the queue for application traffic.

| Characteristics | Value  |
| -               | -:     |
| Default         | `500`  |
| Min length      | `1`    |
| Max length      | `1000` |
| Orderable       | `true` |
| Filterable      | `true` |

### `trustedCAs (external:trusted_cas_list)`

List of trusted CA. If empty the main chain of trust will be used.

| Characteristics | Value |
| -               | -:    |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# EnforcerProfileMappingPolicy

> Operations: `GET` `UPDATE` `DELETE`

A Enforcer Profile Mapping Policy will tell what Enforcer Profile should be used
by and Aporeto Agent based on the Enforcer that have been used during the
registration. The policy can also be propagated down to the child namespace.

## Relations

Relation with [EnforcerProfile](#enforcerprofiles)

- `GET  /enforcerprofilemappingpolicies/id/enforcerprofiles`

Relation with [Enforcer](#enforcers)

- `GET  /enforcerprofilemappingpolicies/id/enforcers`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `object (external:policies_list)`

Object is the list of tags to use to find a enforcer profile.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject is the subject of the policy.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Export

> Operations:

Export the policies and related objects in a given namespace.

## Attributes

### `APIVersion (integer)`

APIVersion of the api used for the exported data

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `auditProfiles (external:exported_data_content)`

List of all exported audit profiles.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `externalServices (external:exported_data_content)`

List of exported external services.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `fileAccessPolicies (external:exported_data_content)`

List of exported file access policies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `filePaths (external:exported_data_content)`

List of exported file paths.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `isolationProfiles (external:exported_data_content)`

List of all exported isolation profiles.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `networkAccessPolicies (external:exported_data_content)`

List of exported network policies

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `processingUnitPolicies (external:exported_data_content)`

List of all exported processingUnitPolicies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Required        | `true` |

# ExternalAccess

> Operations:

ExternalAccess allows to retrieve connection from or to an external service

## Attributes

### `IPRecords (external:ip_records)`

IPRecords refers to a list of IPRecord that contains the IP information

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# ExternalService

> Operations: `GET` `UPDATE` `DELETE`

An External Service represents a random network or ip that is not managed by the
system. They can be used in Network Access Policies in order to allow traffic
from or to the declared network or IP, using the provided protocol and port or
ports range. If you want to describe the Internet (ie. anywhere), use 0.0.0.0/0
as address, and 1-65000 for the ports. You will need to use the External
Services tags to set some policies.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `loadbalancerAddresses (external:addresses_list)`

LoadbalancerAddresses represents the list of adresses of the external services
of type LoadBalancer.

| Characteristics | Value |
| -               | -:    |

### `loadbalancerPortsMapping (external:portmapping_list)`

LoadbalancerPortsMapping is the list of ports mapped by an extenral service of
type load balancer.

| Characteristics | Value |
| -               | -:    |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `network (string)`

Network refers to either CIDR or domain name

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `port (string)`

Port refers to network port which could be a single number or 100:2000 to
represent a range of ports

| Characteristics | Value                                                                                                                                                                                                            |
| -               | -:                                                                                                                                                                                                               |
| Format          | `/^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$/` |
| Default         | `1:65535`                                                                                                                                                                                                        |
| Filterable      | `true`                                                                                                                                                                                                           |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protocol (string)`

Protocol refers to network protocol like TCP/UDP or the number of the protocol.

| Characteristics | Value                                                                  |
| -               | -:                                                                     |
| Format          | `/^(TCP|UDP|tcp|udp|[1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/` |
| Required        | `true`                                                                 |
| Filterable      | `true`                                                                 |

### `type (enum)`

Type represents the type of external service.

| Characteristics | Value                                        |
| -               | -:                                           |
| Allowed Value   | `LoadBalancerHTTP, LoadBalancerTCP, Network` |
| Default         | `Network`                                    |
| Orderable       | `true`                                       |
| Filterable      | `true`                                       |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# FileAccess

> Operations:

Returns file access statistics on a particular processing unit.

## Attributes

### `action (string)`

Action tells if the access has been allowed or not.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `count (integer)`

Count tells how many times the file has been accessed.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `host (string)`

Host is the host that served the accessed file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `mode (enum)`

Mode is the mode of the accessed file.

| Characteristics | Value                    |
| -               | -:                       |
| Allowed Value   | `Read, ReadWrite, Write` |
| Autogenerated   | `true`                   |
| Read only       | `true`                   |

### `path (string)`

Path is the path of the accessed file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protocol (string)`

Protocol is the protocol used to access the file.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# FileAccessPolicy

> Operations: `GET` `UPDATE` `DELETE`

A File Access Policy allows Processing Units to access various folder and files.
It will use the tags of a File Path to know what is the path of the file or
folder to allow access to. You can allow the Processing Unit to have any
combination of read, write or execute.

When a Processing Unit is Docker container, then it will police the volumes
mount. executewon''t have any effect.

File path are not supported yet for standard Linux processes.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `allowsExecute (boolean)`

AllowsExecute allows to execute the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `allowsRead (boolean)`

AllowsRead allows to read the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `allowsWrite (boolean)`

AllowsWrite allows to write the files.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `encryptionEnabled (boolean)`

EncryptionEnabled will enable the automatic encryption

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `logsEnabled (boolean)`

LogsEnabled will enable logging when this policy is used.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `object (external:policies_list)`

Object is the object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject is the subject of the policy

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# FilePath

> Operations: `GET` `UPDATE` `DELETE`

A File Path represents a random path to a file or a folder. They can be used in
aFile Access Policiesin order to allow Processing Units to access them, using
various modes (read, write, execute). You will need to use the File Paths tags
to set some policies. A good example would bevolume=web or file=/etc/passwd.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `filepath (string)`

FilePath refer to the file mount path

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `server (string)`

server is the server name/ID/IP associated with the file path

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# FlowStatistic

> Operations:

Returns network access statistics on a particular processing unit or group of
processing units based on their tags.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `dataPoints (external:datapoints_list)`

DataPoints is a list of time/value pairs that represent the flow events over
time.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `destinationIDs (external:flowstatistic_origin_list)`

DestinationIDs is the IDs of the destination.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `destinationTags (external:selectors_list)`

DestinationTags contains the tags used to identify destination

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `metric (enum)`

Metric is the kind of metric the statistic represents.

| Characteristics | Value          |
| -               | -:             |
| Allowed Value   | `Flows, Ports` |
| Default         | `Flows`        |
| Autogenerated   | `true`         |
| Read only       | `true`         |

### `mode (enum)`

Mode defines if the metric is for accepted or rejected flows.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Accepted, Any, Rejected` |
| Default         | `Accepted`                |
| Autogenerated   | `true`                    |
| Read only       | `true`                    |

### `sourceIDs (external:flowstatistic_origin_list)`

SourceIDs is the sources of the stats.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `sourceTags (external:selectors_list)`

SourceTags contains the tags used to identify the source.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `type (enum)`

Type is the type of representation

| Characteristics | Value                |
| -               | -:                   |
| Allowed Value   | `Repartition, Serie` |
| Default         | `Serie`              |
| Autogenerated   | `true`               |
| Read only       | `true`               |

### `userIdentifier (string)`

UserIdentifier can be set by the user as a query parameter. It will be returned
in the FlowStatistic object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Orderable       | `true` |

# HookPolicy

> Operations: `GET` `UPDATE` `DELETE`

Hook allows to to define hooks to the write operations in squall. Hooks are sent
to an external Rufus server that will do the processing and eventually return a
modified version of the object before we save it.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `certificateAuthority (string)`

CertificateAuthority contains the pem block of the certificate authority used by
the remote endpoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

### `clientCertificate (string)`

ClientCertificate contains the client certificate that will be used to connect
to the remote endoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

### `clientCertificateKey (string)`

ClientCertificateKey contains the key associated to the clientCertificate.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `endpoint (string)`

Endpoint contains the full address of the remote processor endoint.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `mode (enum)`

Mode define the type of the hook.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Both, Post, Pre` |
| Default         | `Pre`             |
| Orderable       | `true`            |
| Filterable      | `true`            |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject contains the tag expression that an object must match in order to
trigger the hook.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Import

> Operations:

Imports an export of policies and related objects into the namespace.

## Attributes

### `data (external:exported_data)`

The data to import.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `mode (enum)`

How to import the data.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Append, Replace` |
| Default         | `Replace`         |
| Required        | `true`            |

# Installation

> Operations: `GET` `UPDATE` `DELETE`

Installation represents an installation for a given account

## Attributes

### `ID (string)`

ID represents the identifier of the installation.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `accountName (string)`

AccountName that should be installed.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

# IsolationProfile

> Operations: `GET` `UPDATE` `DELETE`

An IsolationProfile needs documentation

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `capabilitiesActions (external:cap_map)`

CapabilitiesActions identifies the capabilities that should be added or removed
from the processing unit.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `defaultSyscallAction (external:syscall_action)`

DefaultAction is the default action applied to all syscalls of this profile.
Default is "Allow".

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `syscallRules (external:syscall_rules)`

SyscallRules is a list of syscall rules that identify actions for particular
syscalls.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `targetArchitectures (external:arch_list)`

TargetArchitectures is the target processor architectures where this profile can
be applied. Default all.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Issue

> Operations:

This API issues a new token according to given data.

## Attributes

### `data (string)`

Data contains additional data. The value depends on the issuer type.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata)`

Metadata contains various additional information. Meaning depends on the realm.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `realm (enum)`

Realm is the realm

| Characteristics | Value                                                                              |
| -               | -:                                                                                 |
| Allowed Value   | `AWSIdentityDocument, Certificate, Facebook, Github, Google, LDAP, Twitter, Vince` |
| Required        | `true`                                                                             |

### `token (string)`

Token is the token to use for the registration.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `validity (string)`

Validity configures the max validity time for a token. If it is bigger than the
configured max validity, it will be capped.

| Characteristics | Value                                                                                                             |
| -               | -:                                                                                                                |
| Format          | `/^([0-9]+h[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+m[0-9]+s|[0-9]+h[0-9]+s|[0-9]+h[0-9]+m|[0-9]+s|[0-9]+h|[0-9]+m)$/` |
| Default         | `24h`                                                                                                             |
| Orderable       | `true`                                                                                                            |
| Filterable      | `true`                                                                                                            |

# Jaegerbatch

> Operations:

A jaegerbatch is a batch of jaeger spans. This is used by external service to
post jaeger span in our private jaeger services

## Attributes

### `batch (external:jaeger_batch)`

Represent an jaeger batch

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

# KubernetesCluster

> Operations: `GET` `UPDATE` `DELETE`

Create a remote Kubernetes Cluster integration.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `createTime (time)`

Creation date of the object.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `kubernetesDefinitions (string)`

base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

The name of your cluster

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `parentID (string)`

ID of the parent account.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `regenerate (boolean)`

Regenerates the k8s files and certificates.

| Characteristics | Value |
| -               | -:    |

### `targetNamespace (string)`

The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `targetNetworks (external:target_networks_list)`

List of target networks [deprecated]

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

Last update date of the object

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# Log

> Operations:

Retrieves the log of a deployed app.

## Attributes

### `data (external:logs)`

Data contains all logs data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# Message

> Operations: `GET` `UPDATE` `DELETE`

The Message API allows to post public messages that will be visible through all
children namespaces.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `expirationTime (time)`

expirationTime is the time after which the message will be deleted.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `level (enum)`

Level defines how the message is important.

| Characteristics | Value                   |
| -               | -:                      |
| Allowed Value   | `Danger, Info, Warning` |
| Default         | `Info`                  |
| Orderable       | `true`                  |
| Filterable      | `true`                  |

### `local (boolean)`

If local is set, the message will only be visible in the current namespace.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `notifyByEmail (boolean)`

If enabled, the message will be sent to the email associated in namespaces
annotations.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `validity (string)`

Validity set using golang time duration, when the message will be automatically
deleted.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

# Namespace

> Operations: `GET` `UPDATE` `DELETE`

A Namespace represents the core organizational unit of the system. All objects
always exists in a single namespace. A Namespace can also have child namespaces.
They can be used to split the system into organizations, business units,
applications, services or any combination you like.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `localCA (string)`

LocalCA holds the eventual certificate authority used by this namespace.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `localCAEnabled (boolean)`

LocalCAEnabled defines if the namespace should use a local Certificate
Authority. Switching it off and on again will regenerate a new CA.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the namespace.

| Characteristics | Value                 |
| -               | -:                    |
| Format          | `/^[a-zA-Z0-9-_/]+$/` |
| Required        | `true`                |
| Creation only   | `true`                |
| Orderable       | `true`                |
| Filterable      | `true`                |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# NamespaceMappingPolicy

> Operations: `GET` `UPDATE` `DELETE`

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

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `mappedNamespace (string)`

mappedNamespace is the mapped namespace

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject is the subject.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# NetworkAccessPolicy

> Operations: `GET` `UPDATE` `DELETE`

Allows to define networking policies to allow or prevent processing units
identitied by their tags to talk to other processing units or external services
(also identified by their tags).

## Relations

Relation with [ExternalService](#externalservices)

- `GET  /networkaccesspolicies/id/externalservices`

Relation with [ProcessingUnit](#processingunits)

- `GET  /networkaccesspolicies/id/processingunits`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `allowsTraffic (boolean)`

AllowsTraffic if true, the flow will be accepted. Otherwise other actions like
"logs" can still be done, but the traffic will be rejected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `destinationPorts (external:ports_list)`

DestinationPorts contains the list of allowed ports and ranges.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `encryptionEnabled (boolean)`

EncryptionEnabled defines if the flow has to be encrypted.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `logsEnabled (boolean)`

LogsEnabled defines if the flow has to be logged.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `object (external:policies_list)`

Object of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `observationEnabled (boolean)`

If set to true, the flow will be in observation mode.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `observedTrafficAction (enum)`

If observationEnabled is set to true, this will be the final action taken on the
packets.

| Characteristics | Value             |
| -               | -:                |
| Allowed Value   | `Apply, Continue` |
| Default         | `Continue`        |
| Orderable       | `true`            |
| Filterable      | `true`            |

### `passthrough (external:policies_list)`

List of tags expressions to match the list of entity to pass the flow through.

| Characteristics | Value |
| -               | -:    |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject of the policy.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# PasswordReset

> Operations:

Used to reset an account password.

## Attributes

### `password (string)`

Password contains the new password.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `token (string)`

Token contains the reset password token

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

# Plan

> Operations: `GET`

Plan contains the various billing plans available

## Relations

Relation with [Plan](#plans)

- `GET  /plans/id/plans`

## Attributes

### `description (string)`

Description contains the description of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `enforcersQuota (integer)`

EnforcerQuota contains the maximum number of enforcers available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `key (string)`

Key contains the key identifier of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `name (string)`

Name contains the name of the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `policiesQuota (integer)`

PoliciesQuota contains the maximum number of policies available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `processingUnitsQuota (integer)`

ProcessingUnitsQuota contains the maximum PUs available in the Plan.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# Poke

> Operations:

When available, poke can be used to update various information about the parent.
For instance, for enforcers, poke will be use as the heartbeat.

# Policy

> Operations: `GET` `DELETE`

[nodoc]

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `action (external:actions_list)`

Action defines set of actions that must be enforced when a dependency is met.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `object (external:policies_list)`

Object represents set of entities that another entity depends on. As subjects,
objects are identified as logical operations on tags when a policy is defined.

| Characteristics | Value |
| -               | -:    |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `relation (external:relations_list)`

Relation describes the required operation to be performed between subjects and
objects

| Characteristics | Value |
| -               | -:    |

### `subject (external:policies_list)`

Subject represent sets of entities that will have a dependency other entities.
Subjects are defined as logical operations on tags. Logical operations can
includes AND/OR

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `type (enum)`

Type of the policy

| Characteristics | Value                                                                                                                  |
| -               | -:                                                                                                                     |
| Allowed Value   | `APIAuthorization, EnforcerProfile, File, Hook, NamespaceMapping, Network, ProcessingUnit, Quota, Syscall, TokenScope` |
| Required        | `true`                                                                                                                 |
| Creation only   | `true`                                                                                                                 |
| Filterable      | `true`                                                                                                                 |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# PolicyRefresh

> Operations:

PolicyRefresh is sent to client when as a push event when a policy refresh is
needed on their side.

## Attributes

### `sourceNamespace (string)`

SourceNamespace contains the original namespace of the updated object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `type (string)`

Type contains the policy type that is affected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

# PolicyRule

> Operations: `GET`

PolicyRule is an internal policy resolution API. Services can use this API to
retrieve a policy resolution.

## Attributes

### `APIServices (external:api_services_entities)`

APIServices provides the APIServices of this policy rule.

| Characteristics | Value |
| -               | -:    |

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `action (external:actions_list)`

Action defines set of actions that must be enforced when a dependency is met.

| Characteristics | Value |
| -               | -:    |

### `enforcerProfiles (external:enforcerprofiles_list)`

EnforcerProfiles provides the information about the server profile.

| Characteristics | Value |
| -               | -:    |

### `externalServices (external:network_entities)`

Policy target networks.

| Characteristics | Value |
| -               | -:    |

### `filePaths (external:file_entities)`

Policy target networks.

| Characteristics | Value |
| -               | -:    |

### `isolationProfiles (external:isolation_profile_entities)`

IsolationProfiles are the isolation profiles of the rule.

| Characteristics | Value |
| -               | -:    |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespaces (external:namespace_entities)`

Policy target networks.

| Characteristics | Value |
| -               | -:    |

### `passthroughExternalServices (external:network_entities)`

List of external services the policy mandate to pass through before reaching the
destination.

| Characteristics | Value |
| -               | -:    |

### `propagated (boolean)`

Propagated indicates if the policy is propagated.

| Characteristics | Value |
| -               | -:    |

### `relation (external:relations_list)`

Relation describes the required operation to be performed between subjects and
objects

| Characteristics | Value |
| -               | -:    |

### `tagClauses (external:target_tags)`

Policy target tags

| Characteristics | Value |
| -               | -:    |

# ProcessingUnit

> Operations: `GET` `UPDATE` `DELETE`

A Processing Unit reprents anything that can compute. It can be a Docker
container, or a simple Unix process. They are created, updated and deleted by
the system as they come and go. You can only modify its tags.  Processing Units
use Network Access Policies to define which other Processing Units or External
Services they can communicate with and File Access Policies to define what File
Paths they can use.

## Relations

Relation with [APIService](#apiservices)

- `GET  /processingunits/id/apiservices`

Relation with [FileAccess](#fileaccesses)

- `GET  /processingunits/id/fileaccesses`

Relation with [RenderedPolicy](#renderedpolicies)

- `GET  /processingunits/id/renderedpolicies`

Relation with [Vulnerability](#vulnerabilities)

- `GET  /processingunits/id/vulnerabilities`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `enforcerID (string)`

EnforcerID is the ID of the enforcer associated with the processing unit

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

### `lastSyncTime (time)`

LastSyncTime is the time when the policy was last resolved

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Orderable       | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `nativeContextID (string)`

NativeContextID is the Docker UUID or service PID

| Characteristics | Value  |
| -               | -:     |
| Filterable      | `true` |

### `networkServices (external:processing_unit_services_list)`

NetworkServices is the list of services that this processing unit has declared
that it will be listening to. This can happen either with an activation command
or by exposing the ports in a container manifest.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `operationalStatus (enum)`

OperationalStatus of the processing unit

| Characteristics | Value                                               |
| -               | -:                                                  |
| Allowed Value   | `Initialized, Paused, Running, Stopped, Terminated` |
| Default         | `Initialized`                                       |
| Filterable      | `true`                                              |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `type (enum)`

Type of the container ecosystem

| Characteristics | Value                             |
| -               | -:                                |
| Allowed Value   | `Docker, LinuxService, RKT, User` |
| Required        | `true`                            |
| Creation only   | `true`                            |
| Filterable      | `true`                            |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# ProcessingUnitPolicy

> Operations: `GET` `UPDATE` `DELETE`

A ProcessingUnitPolicies needs a better description.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `action (enum)`

Action determines the action to take while enforcing the isolation profile.

| Characteristics | Value                                                    |
| -               | -:                                                       |
| Allowed Value   | `Delete, Enforce, LogCompliance, Reject, Snapshot, Stop` |
| Orderable       | `true`                                                   |
| Filterable      | `true`                                                   |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `isolationProfileSelector (external:policies_list)`

IsolationProfileSelector are the profiles that must be applied when this policy
matches. Only applies to Enforce and LogCompliance actions.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject defines the tag selectors that identitfy the processing units to which
this policy applies.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# QuotaPolicy

> Operations: `GET` `UPDATE` `DELETE`

Quotas Policies allows to set quotas on the number of objects that can be
created in a namespace.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `identities (list)`

Identities contains the list of identity names where the quota will be applied.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `quota (integer)`

Quota contains the maximum number of object matching the policy subject that can
be created.

| Characteristics | Value |
| -               | -:    |

### `targetNamespace (string)`

TargetNamespace contains the base namespace from where the count will be done.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# RemoteProcessor

> Operations:

Hook to integrate an Aporeto service.

## Attributes

### `claims (list)`

Represents the claims of the currently managed object

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `input (external:raw_json)`

Represents data received from the service

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `mode (enum)`

Node defines the type of the hook

| Characteristics | Value       |
| -               | -:          |
| Allowed Value   | `Post, Pre` |
| Required        | `true`      |

### `namespace (string)`

Represents the current namespace

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `operation (external:elemental_operation)`

Define the operation that is currently handled by the service

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

### `output (external:elemental_identitifable)`

Returns the OutputData filled with the processor information

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `requestID (string)`

RequestID gives the id of the request coming from the main server.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `targetIdentity (string)`

Represents the Identity name of the managed object

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |

# RenderedPolicy

> Operations:

Retrieve the aggregated policies applied to a particular processing unit.

## Attributes

### `certificate (string)`

Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |

### `egressPolicies (external:rendered_policy)`

EgressPolicies lists all the egress policies attached to ProcessingUnit

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `exposedAPIServices (external:api_services_entities)`

ExposedAPIServices is the list of services that this processing unit is
implementing.

| Characteristics | Value |
| -               | -:    |

### `ingressPolicies (external:rendered_policy)`

IngressPolicies lists all the ingress policies attached to ProcessingUnit

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `matchingTags (list)`

MatchingTags contains the list of tags that matched the policies.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `processingUnit (external:processingunit)`

Can be set during a POST operation to render a policy on a Processing Unit that
has not been created yet.

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |

### `processingUnitID (string)`

Identifier of the ProcessingUnit

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `profile (external:trust_profile)`

Profile is the trust profile of the processing unit that should be used during
all communications.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `scopes (external:scopes_list)`

Scopes is the set of scopes granted to this Processing Unit that it has to
present in HTTP requests.

| Characteristics | Value |
| -               | -:    |

# Report

> Operations:

Post a new statistics report.

## Attributes

### `fields (external:tsdb_fields)`

TSDB Fields to set for the report.

| Characteristics | Value |
| -               | -:    |

### `kind (enum)`

Kind contains the kind of report.

| Characteristics | Value                                                              |
| -               | -:                                                                 |
| Allowed Value   | `Audit, Enforcer, FileAccess, Flow, ProcessingUnit, Syscall, User` |

### `tags (external:tags_map)`

Tags contains the tags associated to the data point.

| Characteristics | Value |
| -               | -:    |

### `timestamp (time)`

Timestamp contains the time for the report.

| Characteristics | Value |
| -               | -:    |

### `value (float)`

Value contains the value for the report.

| Characteristics | Value |
| -               | -:    |

# Revocation

> Operations: `UPDATE`

Used to revoke a certificate

## Attributes

### `expirationDate (time)`

Contains the certificate expiration date. This will be used to clean up revoked
certificates that have expired.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `revokeDate (time)`

Set time from when the certificate will be revoked.

| Characteristics | Value |
| -               | -:    |

### `serialNumber (string)`

SerialNumber of the revoked certificate.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `subject (string)`

Subject of the certificate related to the revocation.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

# Role

> Operations:

Roles returns the available roles that can be used with API Authorization
Policies.

## Attributes

### `authorizations (external:authorization_map)`

Authorizations of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `description (string)`

Description is the description of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `key (string)`

Key is the of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `name (string)`

Name of the role.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# Root

> Operations: `GET`

[nodoc]

## Relations

Relation with [Account](#accounts)

- `GET  /root/id/accounts`
- `POST /root/id/accounts`

Relation with [AccountCheck](#accountchecks)

- `GET  /root/id/accountchecks`

Relation with [Activate](#activate)

- `GET  /root/id/activate`

Relation with [Activity](#activities)

- `GET  /root/id/activities`

Relation with [Alarm](#alarms)

- `GET  /root/id/alarms`
- `POST /root/id/alarms`

Relation with [APIAuthorizationPolicy](#apiauthorizationpolicies)

- `GET  /root/id/apiauthorizationpolicies`
- `POST /root/id/apiauthorizationpolicies`

Relation with [APICheck](#apichecks)

- `POST /root/id/apichecks`

Relation with [APIService](#apiservices)

- `GET  /root/id/apiservices`
- `POST /root/id/apiservices`

Relation with [AuditProfile](#auditprofiles)

- `GET  /root/id/auditprofiles`
- `POST /root/id/auditprofiles`

Relation with [Auth](#auth)

- `GET  /root/id/auth`

Relation with [Authority](#authorities)

- `POST /root/id/authorities`

Relation with [Automation](#automations)

- `GET  /root/id/automations`
- `POST /root/id/automations`

Relation with [AutomationTemplate](#automationtemplates)

- `GET  /root/id/automationtemplates`

Relation with [AvailableService](#availableservices)

- `GET  /root/id/availableservices`

Relation with [AWSAccount](#awsaccounts)

- `GET  /root/id/awsaccounts`
- `POST /root/id/awsaccounts`

Relation with [Certificate](#certificates)

- `GET  /root/id/certificates`
- `POST /root/id/certificates`

Relation with [DependencyMap](#dependencymaps)

- `GET  /root/id/dependencymaps`

Relation with [Email](#emails)

- `POST /root/id/emails`

Relation with [Enforcer](#enforcers)

- `GET  /root/id/enforcers`
- `POST /root/id/enforcers`

Relation with [EnforcerProfile](#enforcerprofiles)

- `GET  /root/id/enforcerprofiles`
- `POST /root/id/enforcerprofiles`

Relation with [EnforcerProfileMappingPolicy](#enforcerprofilemappingpolicies)

- `GET  /root/id/enforcerprofilemappingpolicies`
- `POST /root/id/enforcerprofilemappingpolicies`

Relation with [Export](#export)

- `GET  /root/id/export`
- `POST /root/id/export`

Relation with [ExternalAccess](#externalaccesses)

- `GET  /root/id/externalaccesses`

Relation with [ExternalService](#externalservices)

- `GET  /root/id/externalservices`
- `POST /root/id/externalservices`

Relation with [FileAccess](#fileaccesses)

- `GET  /root/id/fileaccesses`

Relation with [FileAccessPolicy](#fileaccesspolicies)

- `GET  /root/id/fileaccesspolicies`
- `POST /root/id/fileaccesspolicies`

Relation with [FilePath](#filepaths)

- `GET  /root/id/filepaths`
- `POST /root/id/filepaths`

Relation with [FlowStatistic](#flowstatistics)

- `GET  /root/id/flowstatistics`

Relation with [HookPolicy](#hookpolicies)

- `GET  /root/id/hookpolicies`
- `POST /root/id/hookpolicies`

Relation with [Import](#import)

- `POST /root/id/import`

Relation with [Installation](#installations)

- `GET  /root/id/installations`
- `POST /root/id/installations`

Relation with [IsolationProfile](#isolationprofiles)

- `GET  /root/id/isolationprofiles`
- `POST /root/id/isolationprofiles`

Relation with [Issue](#issue)

- `POST /root/id/issue`

Relation with [KubernetesCluster](#kubernetesclusters)

- `GET  /root/id/kubernetesclusters`
- `POST /root/id/kubernetesclusters`

Relation with [Message](#messages)

- `GET  /root/id/messages`
- `POST /root/id/messages`

Relation with [Namespace](#namespaces)

- `GET  /root/id/namespaces`
- `POST /root/id/namespaces`

Relation with [NamespaceMappingPolicy](#namespacemappingpolicies)

- `GET  /root/id/namespacemappingpolicies`
- `POST /root/id/namespacemappingpolicies`

Relation with [NetworkAccessPolicy](#networkaccesspolicies)

- `GET  /root/id/networkaccesspolicies`
- `POST /root/id/networkaccesspolicies`

Relation with [PasswordReset](#passwordreset)

- `GET  /root/id/passwordreset`
- `POST /root/id/passwordreset`

Relation with [Plan](#plans)

- `GET  /root/id/plans`

Relation with [Policy](#policies)

- `GET  /root/id/policies`

Relation with [PolicyRule](#policyrules)

- `GET  /root/id/policyrules`
- `POST /root/id/policyrules`

Relation with [ProcessingUnit](#processingunits)

- `GET  /root/id/processingunits`
- `POST /root/id/processingunits`

Relation with [ProcessingUnitPolicy](#processingunitpolicies)

- `GET  /root/id/processingunitpolicies`
- `POST /root/id/processingunitpolicies`

Relation with [QuotaPolicy](#quotapolicies)

- `GET  /root/id/quotapolicies`
- `POST /root/id/quotapolicies`

Relation with [RemoteProcessor](#remoteprocessors)

- `POST /root/id/remoteprocessors`

Relation with [RenderedPolicy](#renderedpolicies)

- `POST /root/id/renderedpolicies`

Relation with [Report](#reports)

- `POST /root/id/reports`

Relation with [Revocation](#revocations)

- `GET  /root/id/revocations`

Relation with [Role](#roles)

- `GET  /root/id/roles`

Relation with [Service](#services)

- `GET  /root/id/services`
- `POST /root/id/services`

Relation with [StatsQuery](#statsqueries)

- `GET  /root/id/statsqueries`

Relation with [SuggestedPolicy](#suggestedpolicies)

- `GET  /root/id/suggestedpolicies`

Relation with [SystemCall](#systemcalls)

- `GET  /root/id/systemcalls`
- `POST /root/id/systemcalls`

Relation with [Tabulation](#tabulations)

- `GET  /root/id/tabulations`

Relation with [Tag](#tags)

- `GET  /root/id/tags`

Relation with [Token](#tokens)

- `POST /root/id/tokens`

Relation with [TokenScopePolicy](#tokenscopepolicies)

- `GET  /root/id/tokenscopepolicies`
- `POST /root/id/tokenscopepolicies`

Relation with [Vulnerability](#vulnerabilities)

- `GET  /root/id/vulnerabilities`
- `POST /root/id/vulnerabilities`

Relation with [X509Certificate](#x509certificates)

- `GET  /root/id/x509certificates`
- `POST /root/id/x509certificates`

Relation with [X509CertificateCheck](#x509certificatechecks)

- `GET  /root/id/x509certificatechecks`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

# Service

> Operations: `GET` `UPDATE` `DELETE`

Service represents a service that can be launched

## Relations

Relation with [Log](#logs)

- `GET  /services/id/logs`

## Attributes

### `ID (string)`

ID of the service

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Filterable      | `true` |

### `accountName (string)`

AccountName represents the vince account name

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |

### `categoryID (string)`

CategoryID of the service.

| Characteristics | Value  |
| -               | -:     |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

Name of the service

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace in which the service in running.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `parameters (external:service_parameters)`

Parameters is a list of parameters to start the service

| Characteristics | Value |
| -               | -:    |

### `replicas (integer)`

Replicas represents the number of replicas for the service.

| Characteristics | Value |
| -               | -:    |

### `status (enum)`

Status of the service.

| Characteristics | Value                     |
| -               | -:                        |
| Allowed Value   | `Error, Pending, Running` |
| Default         | `Pending`                 |
| Read only       | `true`                    |
| Orderable       | `true`                    |
| Filterable      | `true`                    |

# StatsQuery

> Operations:

StatsQuery is a generic API to retrieve time series data stored by the Aporeto
system. The API allows different types of queries that are all protected within
the namespace of the user.

## Attributes

### `results (external:time_series_results)`

Results contains the result of the query.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# SuggestedPolicy

> Operations:

Allows to get policy suggestions

## Attributes

### `networkAccessPolicies (external:network_access_policies_list)`

List of suggested network access policies

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

# SystemCall

> Operations: `GET` `UPDATE` `DELETE`

[nodoc]

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Tabulation

> Operations:

Tabulate API allows you to retrieve a custom table view for any identity using
any tags you like as columns.

## Attributes

### `headers (list)`

Headers contains the requests headers that matched.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `rows (external:tabulated_data)`

Rows contains the tabulated data.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `targetIdentity (string)`

TargetIdentity contains the requested target identity.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

# Tag

> Operations:

A tag is a string in the form of "key=value" that can applied to all objects in
the system. They are used for policy resolution. Tags starting by a "$" are
derived from the property of an object (for instance an object with ID set to
xxx and a name set to "the name" will be tagged by default with "$name=the name"
and "$id=xxx"). Tags starting with an "@" have been generated by an external
system.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `count (integer)`

Count represents the number of time the tag is used.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `namespace (string)`

Namespace represents the namespace of the counted tag.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `value (string)`

Value represents the value of the tag.

| Characteristics | Value                                                     |
| -               | -:                                                        |
| Format          | `/^[\w\d\*\$\+\.:,|@<>/-]+=[= \w\d\*\$\+\.:,|@~<>#/-]+$/` |
| Required        | `true`                                                    |
| Creation only   | `true`                                                    |

# TokenScopePolicy

> Operations: `GET` `UPDATE` `DELETE`

The TokenScopePolicy defines a set of policies that allow customization of the
authorization tokens issued by the Aporeto service. This allows Aporeto
generated tokens to be used by external applications.

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `activeDuration (string)`

ActiveDuration defines for how long the policy will be active according to the
activeSchedule.

| Characteristics | Value             |
| -               | -:                |
| Format          | `/^[0-9]+[smh]$/` |

### `activeSchedule (external:cron_expression)`

ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.

| Characteristics | Value |
| -               | -:    |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `assignedScopes (external:tags_list)`

AssignedScopes is the the list of scopes that the policiy will assigns.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `disabled (boolean)`

Disabled defines if the propert is disabled.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `metadata (external:metadata_list)`

Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.

| Characteristics | Value  |
| -               | -:     |
| Creation only   | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `propagate (boolean)`

Propagate will propagate the policy to all of its children.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `propagationHidden (boolean)`

If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `subject (external:policies_list)`

Subject defines the selection criteria that this policy must match on identiy
and scope request information.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

# Trigger

> Operations:

Trigger can be used to remotely trigger an automation.

## Attributes

# Vulnerability

> Operations: `GET`

A vulnerabily represents a particular CVE

## Relations

Relation with [ProcessingUnit](#processingunits)

- `GET  /vulnerabilities/id/processingunits`

## Attributes

### `ID (string)`

ID is the identifier of the object.

| Characteristics | Value  |
| -               | -:     |
| Identifier      | `true` |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `annotations (external:annotations)`

Annotation stores additional information about an entity

| Characteristics | Value |
| -               | -:    |

### `associatedTags (external:tags_list)`

AssociatedTags are the list of tags attached to an entity

| Characteristics | Value |
| -               | -:    |

### `createTime (time)`

CreatedTime is the time at which the object was created

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |

### `description (string)`

Description is the description of the object.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `link (string)`

Link is the URL that refers to the vulnerability

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Required        | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `name (string)`

Name is the name of the entity

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `namespace (string)`

Namespace tag attached to an entity

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Creation only   | `true` |
| Orderable       | `true` |
| Filterable      | `true` |

### `normalizedTags (external:tags_list)`

NormalizedTags contains the list of normalized tags of the entities

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |

### `protected (boolean)`

Protected defines if the object is protected.

| Characteristics | Value  |
| -               | -:     |
| Orderable       | `true` |
| Filterable      | `true` |

### `severity (external:vulnerability_level)`

Severity refers to the security vulnerability level

| Characteristics | Value  |
| -               | -:     |
| Required        | `true` |
| Creation only   | `true` |
| Filterable      | `true` |

### `updateTime (time)`

UpdateTime is the time at which an entity was updated.

| Characteristics | Value  |
| -               | -:     |
| Autogenerated   | `true` |
| Read only       | `true` |
| Orderable       | `true` |
