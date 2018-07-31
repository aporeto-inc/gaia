# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  description: root object.
  get:
    description: Retrieves the object with the given ID.
  extends:
  - '@identifiable-nopk-nostored'
  root: true

# Relations
relations:
- rest_name: account
  get:
    description: |-
      Retrieves all accounts. This is a private API that can only be done by the
      system.
  create:
    description: Creates a new Account.

- rest_name: accountcheck
  get:
    description: Verifies an account credentials.

- rest_name: activate
  get:
    description: Activates a pending account.

- rest_name: activity
  get:
    description: Retrieves the list of activity logs.

- rest_name: alarm
  get:
    description: Retrieves all the alarms.
  create:
    description: Creates a new alarm.

- rest_name: apiauthorizationpolicy
  get:
    description: Retrieves the list of API authorization policies.
  create:
    description: Creates a new API authorization policies.

- rest_name: apicheck
  create:
    description: Verifies the authorizations on various identities for a given token.

- rest_name: restapispec
  get:
    description: Retrieves the list of REST API specifications.
  create:
    description: Creates a new REST API specification.

- rest_name: service
  get:
    description: Retrieves the list of Services.
  create:
    description: Creates a new Service.

- rest_name: auditprofile
  get:
    description: Retrieves the list of audit profiles.
  create:
    description: Creates a new audit profile.

- rest_name: auth
  get:
    description: Verify the validity of a token.

- rest_name: authority
  create:
    description: Creates a new certificate authority.

- rest_name: automation
  get:
    description: Retrieves the list of Automations.
  create:
    description: Creates a new Automation.

- rest_name: automationtemplate
  get:
    description: Retrieves the list of automation templates.

- rest_name: app
  get:
    description: Retrieves the list of apps.

- rest_name: awsaccount
  get:
    description: Retrieves the list of aws account bindings.
  create:
    description: Creates a new aws account binding.

- rest_name: awsregister
  create:
    description: Creates a new aws registration for billing.

- rest_name: awsapigateway
  get:
    description: create an AWS API Gateway.
  create:
    description: Manages the AWS API Gateway.

- rest_name: certificate
  get:
    description: Retrieves the list of existing user certificates.
  create:
    description: Creates a new user certificate.

- rest_name: dependencymap
  get:
    description: Retrieves the dependencymap of a namespace.

- rest_name: email
  create:
    description: Sends an email.

- rest_name: enforcer
  get:
    description: Retrieves the list of enforcers.
  create:
    description: Creates a new enforcer.

- rest_name: enforcerprofile
  get:
    description: Retrieves the list of enforcer profiles.
  create:
    description: Creates a new enforcer profile.

- rest_name: enforcerprofilemappingpolicy
  get:
    description: Retrieves the list of enforcer profile mapping policies.
  create:
    description: Creates a new enforcer profile mapping policies.

- rest_name: eventlog
  get:
    description: Retrieves the eventlogs for one or multiple entities.
  create:
    description: Creates a new eventlog for a particular entity.

- rest_name: export
  create:
    description: Exports all policies and related object of a namespace.

- rest_name: externalaccess
  get:
    description: Retrieves the list of external access according to parameters.

- rest_name: externalservice
  get:
    description: Retrieves the list of external services.
  create:
    description: Creates a new external service.

- rest_name: externalnetwork
  get:
    description: Retrieves the list of external network.
  create:
    description: Creates a new external network.

- rest_name: fileaccess
  get:
    description: Retrieves the list of file access according to parameters.

- rest_name: fileaccesspolicy
  get:
    description: Retrieves the list of file access policies.
  create:
    description: Creates a new file access policies.

- rest_name: filepath
  get:
    description: Retrieves the list of file path.
  create:
    description: Create a new file path.

- rest_name: flowstatistic
  get:
    description: Retrieves the flow statistics according to parameters.

- rest_name: hookpolicy
  get:
    description: Retrieves the list of hook policies.
  create:
    description: Creates a new hook policy.

- rest_name: import
  create:
    description: Imports data from a previous export.

- rest_name: installation
  get:
    description: Retrieves the list of installations.
  create:
    description: Creates a new installation.

- rest_name: isolationprofile
  get:
    description: Retrieves the list of isolation profiles.
  create:
    description: Creates a new isolation profile.

- rest_name: issue
  create:
    description: Issues a new token.

- rest_name: jaegerbatch
  create:
    description: Sends a jaeger tracing batch.

- rest_name: kubernetescluster
  get:
    description: Retrieves the list of kubernetes clusters.
  create:
    description: Creates a new kubernetes cluster.

- rest_name: k8scluster
  get:
    description: Retrieves the list of kubernetes clusters.
  create:
    description: Creates a new kubernetes cluster.

- rest_name: message
  get:
    description: Retrieves the list of messages.
  create:
    description: Creates a new message.

- rest_name: namespace
  get:
    description: Retrieves the list of namespaces.
  create:
    description: Creates a new namespace.

- rest_name: namespacemappingpolicy
  get:
    description: Retrieves the list namespace mapping policies.
  create:
    description: Creates a new namespace mapping policy.

- rest_name: networkaccesspolicy
  get:
    description: Retrieves the list of network access policies.
  create:
    description: Creates a new network access policy.

- rest_name: passwordreset
  get:
    description: Sends a link to the account email to reset the password.
  create:
    description: Resets the password for an account using the provided link.

- rest_name: plan
  get:
    description: Retrieves the list of plans.

- rest_name: policy
  get:
    description: Retrieves the list of policy primitives.

- rest_name: policyrenderer
  create:
    description: Render a policy of a given type for a given set of tags.

- rest_name: processingunit
  get:
    description: Retrieves the list of processing units.
  create:
    description: Creates a new processing unit.

- rest_name: processingunitpolicy
  get:
    description: Retrieves the list of processing unit policies.
  create:
    description: Creates a new processing unit policy.

- rest_name: quotacheck
  create:
    description: Verifies if the quota is exceeded for a particular object.

- rest_name: quotapolicy
  get:
    description: Retrieves the list of quota policies.
  create:
    description: Creates a new quota policy.

- rest_name: remoteprocessor
  create:
    description: This should be be here.

- rest_name: renderedpolicy
  create:
    description: Render a policy for a processing unit.

- rest_name: report
  create:
    description: Create a statistics report.

- rest_name: flowreport
  create:
    description: Create a flow statistics report.

- rest_name: revocation
  get:
    description: Verify the revocation of a certificate according to parameters.

- rest_name: role
  get:
    description: Retrieves the list of existing roles.

- rest_name: installedapp
  get:
    description: Retrieves the list of installed apps.
  create:
    description: Installs a new app.

- rest_name: statsquery
  get:
    description: Retrieves statistics information based on parameters.

- rest_name: servicedependency
  get:
    description: Retrieves the list of service dependencies.
  create:
    description: Creates a new service dependency.

- rest_name: suggestedpolicy
  get:
    description: Retrieves a list of network policy suggestion.

- rest_name: tabulation
  get:
    description: Retrieves tabulated informations based on parameters.

- rest_name: tag
  get:
    description: Retrieves the list of existing tags in the system.

- rest_name: taginject
  create:
    description: Internal api to inject tags.

- rest_name: token
  create:
    description: Creates a new token.

- rest_name: tokenscopepolicy
  get:
    description: Retrieves the list of token scope policies.
  create:
    description: Creates a new token scope policy.

- rest_name: vulnerability
  get:
    description: Retrieves the list of vulnerabilities.
  create:
    description: Creates a new vulnerability.

- rest_name: x509certificate
  get:
    description: Retrieves a X509 certificates.
  create:
    description: Creates a new x509 certificate.

- rest_name: x509certificatecheck
  get:
    description: Verifies if a x509 certificate is valid.

- rest_name: squalltag
  get:
    description: Retrieves a computed list of tags from squall for caching.
