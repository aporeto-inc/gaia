# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  group: core
  description: root object.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-not-stored'
  root: true

# Relations
relations:
- rest_name: account
  get:
    description: |-
      Retrieves all accounts. This is a private API that can only be done by the
      system.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: name
        description: internal parameters.
        type: string
        example_value: account

      - name: status
        description: internal parameters.
        type: string
        example_value: status
  create:
    description: Creates a new Account.

- rest_name: accountcheck
  get:
    description: Verifies an account credentials.

- rest_name: activate
  get:
    description: Activates a pending account.
    parameters:
      required:
      - - - token
      entries:
      - name: noRedirect
        description: If set, do not redirect the request to the UI.
        type: boolean

      - name: token
        description: Activation token.
        type: string
        example_value: xxx-xxx-xxx-xxx

- rest_name: activity
  get:
    description: Retrieves the list of activity logs.
    global_parameters:
    - $filtering

- rest_name: alarm
  get:
    description: Retrieves all the alarms.
    global_parameters:
    - $filtering
  create:
    description: Creates a new alarm.

- rest_name: apiauthorizationpolicy
  get:
    description: Retrieves the list of API authorization policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new API authorization policies.

- rest_name: apicheck
  create:
    description: Verifies the authorizations on various identities for a given token.

- rest_name: app
  get:
    description: Retrieves the list of apps.
    global_parameters:
    - $filtering

- rest_name: appcredential
  get:
    description: Retrieves the list of application credentials.
    global_parameters:
    - $filtering
  create:
    description: Creates a new application credential.

- rest_name: auditprofile
  get:
    description: Retrieves the list of audit profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new audit profile.

- rest_name: auditprofilemappingpolicy
  get:
    description: Retrieves the list of audit profile mapping policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new audit profile mapping policy.

- rest_name: auditreport
  create:
    description: Create a audit statistics report.

- rest_name: auth
  get:
    description: Verify the validity of a token.
    parameters:
      entries:
      - name: token
        description: token to validate.
        type: string
        example_value: abc.def.ghi

- rest_name: authority
  create:
    description: Creates a new certificate authority.
    global_parameters:
    - $filtering

- rest_name: automation
  get:
    description: Retrieves the list of Automations.
    global_parameters:
    - $filtering
  create:
    description: Creates a new Automation.

- rest_name: automationtemplate
  get:
    description: Retrieves the list of automation templates.

- rest_name: awsaccount
  get:
    description: Retrieves the list of aws account bindings.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: accountid
        description: Id of the account.
        type: string
        example_value: xxx-xxx-xxx-xxx
  create:
    description: Creates a new aws account binding.

- rest_name: awsapigateway
  get:
    description: create an AWS API Gateway.
    global_parameters:
    - $filtering
  create:
    description: Manages the AWS API Gateway.

- rest_name: awsregister
  create:
    description: Creates a new aws registration for billing.

- rest_name: datapathcertificate
  create:
    description: Creates a new Certificate for datapath.

- rest_name: dependencymap
  get:
    description: Retrieves the dependencymap of a namespace.
    global_parameters:
    - $timewindow
    - $flowoffset
    parameters:
      entries:
      - name: tag
        description: Only show objects with the given tags in the dependency map.
        type: string
        multiple: true
        example_value: env=prod

      - name: view
        description: Set the view query for grouping the dependency map.
        type: string
        example_value: $namespace then app

      - name: viewSuggestions
        description: Also return the view suggestions.
        type: boolean

- rest_name: email
  create:
    description: Sends an email.

- rest_name: enforcer
  get:
    description: Retrieves the list of enforcers.
    global_parameters:
    - $filtering
  create:
    description: Creates a new enforcer.

- rest_name: enforcerprofile
  get:
    description: Retrieves the list of enforcer profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new enforcer profile.

- rest_name: enforcerprofilemappingpolicy
  get:
    description: Retrieves the list of enforcer profile mapping policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new enforcer profile mapping policies.

- rest_name: enforcerreport
  create:
    description: Create a enforcer statistics report.

- rest_name: enforcertracereport
  create:
    description: Create an enforcer trace report.

- rest_name: eventlog
  create:
    description: Creates a new eventlog for a particular entity.

- rest_name: export
  create:
    description: Exports all policies and related object of a namespace.
    parameters:
      entries:
      - name: ignoredTags
        description: List of tags to ignore from the export.
        type: string
        multiple: true
        example_value: a=a

- rest_name: externalnetwork
  get:
    description: Retrieves the list of external network.
    global_parameters:
    - $filtering
    - $archivable
    - $propagatable
  create:
    description: Creates a new external network.

- rest_name: fileaccesspolicy
  get:
    description: Retrieves the list of file access policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new file access policies.

- rest_name: fileaccessreport
  create:
    description: Create a file access statistics report.

- rest_name: filepath
  get:
    description: Retrieves the list of file path.
    global_parameters:
    - $filtering
    - $archivable
    - $propagatable
  create:
    description: Create a new file path.

- rest_name: flowreport
  create:
    description: Create a flow statistics report.

- rest_name: hookpolicy
  get:
    description: Retrieves the list of hook policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new hook policy.

- rest_name: hostservice
  get:
    description: Retrieves the list of host services.
    global_parameters:
    - $filtering
    - $propagatable
    - $propagatable
  create:
    description: Creates a new host service.

- rest_name: hostservicemappingpolicy
  get:
    description: Retrieves the list of host service mapping policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new host service mapping policy.

- rest_name: httpresourcespec
  get:
    description: Retrieves the list of HTTP Resource specifications.
    global_parameters:
    - $filtering
    - $propagatable
    - $archivable
  create:
    description: Creates a new HTTP Resource specification.

- rest_name: import
  create:
    description: Imports data from a previous export.

- rest_name: importrequest
  get:
    description: Retrieves the list of import requests.
    global_parameters:
    - $filtering
  create:
    description: Creates a new import request.

- rest_name: installedapp
  get:
    description: Retrieves the list of installed apps.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: tag
        description: List of tags to filter on. This parameter is deprecated.
        type: string
        multiple: true
        example_value: a=a
  create:
    description: Installs a new app.

- rest_name: ipinfo
  get:
    description: Returns information about an IP address given as parameters.
    parameters:
      required:
      - - - ip
      entries:
      - name: ip
        description: List of IPs to resolve.
        type: string
        multiple: true
        example_value: 1.2.3.4

- rest_name: isolationprofile
  get:
    description: Retrieves the list of isolation profiles.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new isolation profile.

- rest_name: issue
  create:
    description: Issues a new token.
    parameters:
      entries:
      - name: token
        description: Token to verify.
        type: string
        example_value: xxx.yyyyyyyy.zzzz

- rest_name: jaegerbatch
  create:
    description: Sends a jaeger tracing batch.

- rest_name: message
  get:
    description: Retrieves the list of messages.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new message.

- rest_name: namespace
  get:
    description: Retrieves the list of namespaces.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: authorized
        description: Returns all namespaces the token bearer has the right to read.
          If set, other parameters like `recursive` or `q` will have no effect.
        type: boolean
  create:
    description: Creates a new namespace.

- rest_name: namespacemappingpolicy
  get:
    description: Retrieves the list namespace mapping policies.
    global_parameters:
    - $filtering
  create:
    description: Creates a new namespace mapping policy.

- rest_name: networkaccesspolicy
  get:
    description: Retrieves the list of network access policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new network access policy.

- rest_name: oidcprovider
  get:
    description: Retrieves the list of the account OIDC provider.
    global_parameters:
    - $filtering
    parameters:
      entries:
      - name: OIDCProviderName
        description: OIDCProvider name.
        type: string
        example_value: okta

      - name: accountID
        description: Aporeto Account ID.
        type: string
        example_value: aporeto
  create:
    description: Creates a new OIDC provider.

- rest_name: packetreport
  create:
    description: Create a packet trace report.

- rest_name: passwordreset
  get:
    description: Sends a link to the account email to reset the password.
    parameters:
      required:
      - - - email
      entries:
      - name: email
        description: Email associated to the account.
        type: string
        example_value: user@domain.com
  create:
    description: Resets the password for an account using the provided link.

- rest_name: plan
  get:
    description: Retrieves the list of plans.

- rest_name: policy
  get:
    description: Retrieves the list of policy primitives.
    global_parameters:
    - $filtering
    - $propagatable

- rest_name: policygraph
  create:
    description: Retrieve a policy graph.

- rest_name: policyrenderer
  create:
    description: Render a policy of a given type for a given set of tags.

- rest_name: processingunit
  get:
    description: Retrieves the list of processing units.
    global_parameters:
    - $filtering
    - $archivable
  create:
    description: Creates a new processing unit.

- rest_name: processingunitpolicy
  get:
    description: Retrieves the list of processing unit policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new processing unit policy.

- rest_name: punode
  get:
    description: Retrieves the pu nodes.
    global_parameters:
    - $timewindow
    - $archivable
    - $filtering

- rest_name: quotacheck
  create:
    description: Verifies if the quota is exceeded for a particular object.

- rest_name: quotapolicy
  get:
    description: Retrieves the list of quota policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new quota policy.

- rest_name: remoteprocessor
  create:
    description: This should be be here.

- rest_name: renderedpolicy
  create:
    description: Render a policy for a processing unit.
    parameters:
      entries:
      - name: csr
        description: CSR to sign.
        type: string
        example_value: |-
          --- BEGIN CSR ---
          xxx-xxx-xxx-xxx
          --- END CSR ---

- rest_name: report
  create:
    description: Create a statistics report.

- rest_name: revocation
  get:
    description: Verify the revocation of a certificate according to parameters.
    global_parameters:
    - $filtering

- rest_name: role
  get:
    description: Retrieves the list of existing roles.

- rest_name: service
  get:
    description: Retrieves the list of Services.
    global_parameters:
    - $filtering
    - $archivable
  create:
    description: Creates a new Service.

- rest_name: servicedependency
  get:
    description: Retrieves the list of service dependencies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new service dependency.

- rest_name: squalltag
  get:
    description: Retrieves a computed list of tags from squall for caching.
    parameters:
      required:
      - - - identity
      entries:
      - name: identity
        description: Search for all the tags used for the this identity.
        type: string
        example_value: processingunit

- rest_name: sshauthority
  create:
    description: Creates a new SSH CA.
  delete:
    description: Deletes an existing SSH CA.

- rest_name: sshauthorizationpolicy
  get:
    description: Retrieves the list of SSH authorization policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new SSH authorization policies.

- rest_name: sshcertificate
  create:
    description: Creates a new SSH Certificate.

- rest_name: sshidentity
  create:
    description: Creates a new SSH Identity Certificate.

- rest_name: statsinfo
  create:
    description: Retrieves information about the content of the stats measurement.

- rest_name: statsquery
  get:
    description: Retrieves statistics information based on parameters.
    global_parameters:
    - $timewindow
    - $offset
    parameters:
      entries:
      - name: field
        description: list of fields to query.
        type: string
        multiple: true
        example_value: value

      - name: function
        description: function to use.
        type: enum
        allowed_choices:
        - count
        - mean
        - median
        - sum
        - first
        - last
        - max
        - min
        - distinct

      - name: groupBy
        description: list of groupBy clauses.
        type: string
        multiple: true
        example_value: namespace

      - name: interval
        description: list of groupBy clauses.
        type: duration

      - name: measurement
        description: Name of the measurement to query.
        type: enum
        allowed_choices:
        - flows
        - audit
        - enforcers
        - files
        - eventlogs
        default_value: flows

      - name: query
        description: Raw InfluxDB query. This query will be applied to an immutable
          subquery that will bound the data to your namespace. You must pass the measurement
          parameter and use `from _dataset_` in your query. The dataset will be bound
          to the timing parameter you pass.
        type: string
        example_value: select a,b from _dataset_ where a=a group by b limit 10

      - name: tag
        description: list of tags to query.
        type: string
        multiple: true
        example_value: '@tag'

      - name: where
        description: list of where clauses. If `query` is given, this where clause
          will be applied to the controlled subquery as additional `and` clauses to
          the internally computed where clause.
        type: string
        multiple: true
        example_value: a=a
  create:
    description: Sends a query on statistical data.
    global_parameters:
    - $timewindow
    - $filtering

- rest_name: suggestedpolicy
  get:
    description: Retrieves a list of network policy suggestion.
    global_parameters:
    - $timewindow
    - $flowoffset
    parameters:
      entries:
      - name: filterAction
        description: Action to take with the filter tags.
        type: enum
        allowed_choices:
        - include
        - exclude

      - name: filterTags
        description: Tags to filter in the policy suggestions.
        type: string
        multiple: true
        example_value: a=a

- rest_name: tabulation
  get:
    description: Retrieves tabulated informations based on parameters.
    global_parameters:
    - $filtering
    parameters:
      required:
      - - - identity
      entries:
      - name: column
        description: Columns you want to see.
        type: string
        multiple: true
        example_value: name

      - name: identity
        description: Identity you want to tabulate.
        type: string
        example_value: enforcer

- rest_name: tag
  get:
    description: Retrieves the list of existing tags in the system.

- rest_name: taginject
  create:
    description: Internal api to inject tags.

- rest_name: tagvalue
  get:
    description: Retrieves the list of existing values for the given tag keys.
    parameters:
      required:
      - - - key
      entries:
      - name: key
        description: Keys of the tag you want to get the values of.
        type: string
        multiple: true
        example_value: $name

- rest_name: token
  create:
    description: Creates a new token.

- rest_name: tokenscopepolicy
  get:
    description: Retrieves the list of token scope policies.
    global_parameters:
    - $filtering
    - $propagatable
  create:
    description: Creates a new token scope policy.

- rest_name: vulnerability
  get:
    description: Retrieves the list of vulnerabilities.
    global_parameters:
    - $filtering
  create:
    description: Creates a new vulnerability.

- rest_name: x509certificate
  create:
    description: Creates a new x509 certificate.

- rest_name: x509certificatecheck
  get:
    description: Verifies if a x509 certificate is valid.
    global_parameters:
    - $filtering
