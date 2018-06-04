# Model
model:
  rest_name: root
  resource_name: root
  entity_name: Root
  package: root
  description: root object.
  get: true
  extends:
  - '@identifiable-nopk-nostored'
  root: true

# Relations
relations:
- rest_name: account
  descriptions:
    create: Creates a new Account.
    get: |-
      Retrieves all accounts. This is a private API that can only be done by the
      system.
  get: true
  create: true

- rest_name: accountcheck
  descriptions:
    get: Verifies an account credentials.
  get: true

- rest_name: activate
  descriptions:
    get: Activates a pending account.
  get: true

- rest_name: activity
  descriptions:
    get: Retrieves the list of activity logs.
  get: true

- rest_name: alarm
  descriptions:
    create: Creates a new alarm.
    get: Retrieves all the alarms.
  get: true
  create: true

- rest_name: apiauthorizationpolicy
  descriptions:
    create: Creates a new API authorization policies.
    get: Retrieves the list of API authorization policies.
  get: true
  create: true

- rest_name: apicheck
  descriptions:
    create: Verfies the authorizations on various identities for a given token.
  create: true

- rest_name: restapispec
  descriptions:
    create: Creates a new REST API specification.
    get: Retrieves the list of REST API specifications.
  get: true
  create: true

- rest_name: service
  descriptions:
    create: Creates a new Service.
    get: Retrieves the list of Services.
  get: true
  create: true

- rest_name: auditprofile
  descriptions:
    create: Creates a new audit profile.
    get: Retrieves the list of audit profiles.
  get: true
  create: true

- rest_name: auth
  descriptions:
    get: Verify the validity of a token.
  get: true

- rest_name: authority
  descriptions:
    create: Creates a new certificate authority.
  create: true

- rest_name: automation
  descriptions:
    create: Creates a new Automation.
    get: Retrieves the list of Automations.
  get: true
  create: true

- rest_name: automationtemplate
  descriptions:
    get: Retrieves the list of automation templates.
  get: true

- rest_name: app
  descriptions:
    get: Retrieves the list of apps.
  get: true

- rest_name: awsaccount
  descriptions:
    create: Creates a new aws account binding.
    get: Retrieves the list of aws account bindings.
  get: true
  create: true

- rest_name: certificate
  descriptions:
    create: Creates a new user certificate.
    get: Retrieves the list of existing user certificates.
  get: true
  create: true

- rest_name: dependencymap
  descriptions:
    get: Retrieves the dependencymap of a namespace.
  get: true

- rest_name: email
  descriptions:
    create: Sends an email.
  create: true

- rest_name: enforcer
  descriptions:
    create: Creates a new enforcer.
    get: Retrieves the list of enforcers.
  get: true
  create: true

- rest_name: enforcerprofile
  descriptions:
    create: Creates a new enforcer profile.
    get: Retrieves the list of enforcer profiles.
  get: true
  create: true

- rest_name: enforcerprofilemappingpolicy
  descriptions:
    create: Creates a new enforcer profile mapping policies.
    get: Retrieves the list of enforcer profile mapping policies.
  get: true
  create: true

- rest_name: export
  descriptions:
    get: Exports all policies and related object of a namespace.
  get: true

- rest_name: externalaccess
  descriptions:
    get: Retrieves the list of external access according to parameters.
  get: true

- rest_name: externalservice
  descriptions:
    create: Creates a new external service.
    get: Retrieves the list of external services.
  get: true
  create: true

- rest_name: fileaccess
  descriptions:
    get: Retrieves the list of file access according to parameters.
  get: true

- rest_name: fileaccesspolicy
  descriptions:
    create: Creates a new file access policies.
    get: Retrieves the list of file access policies.
  get: true
  create: true

- rest_name: filepath
  descriptions:
    create: Create a new file path.
    get: Retrieves the list of file path.
  get: true
  create: true

- rest_name: flowstatistic
  descriptions:
    get: Retrieves the flow statistics according to parameters.
  get: true

- rest_name: hookpolicy
  descriptions:
    create: Creates a new hook policy.
    get: Retrieves the list of hook policies.
  get: true
  create: true

- rest_name: import
  descriptions:
    create: Imports data from a previous export.
  create: true

- rest_name: installation
  descriptions:
    create: Creates a new installation.
    get: Retrieves the list of installations.
  get: true
  create: true

- rest_name: isolationprofile
  descriptions:
    create: Creates a new isolation profile.
    get: Retrieves the list of isolation profiles.
  get: true
  create: true

- rest_name: issue
  descriptions:
    create: Issues a new token.
  create: true

- rest_name: jaegerbatch
  descriptions:
    create: Sends a jaeger tracing batch.
  create: true

- rest_name: kubernetescluster
  descriptions:
    create: Creates a new kubernetes cluster.
    get: Retrieves the list of kubernetes clusters.
  get: true
  create: true

- rest_name: k8scluster
  descriptions:
    create: Creates a new kubernetes cluster.
    get: Retrieves the list of kubernetes clusters.
  get: true
  create: true

- rest_name: message
  descriptions:
    create: Creates a new message.
    get: Retrieves the list of messages.
  get: true
  create: true

- rest_name: namespace
  descriptions:
    create: Creates a new namespace.
    get: Retrieves the list of namespaces.
  get: true
  create: true

- rest_name: namespacemappingpolicy
  descriptions:
    create: Creates a new namespace mapping policy.
    get: Retrieves the list namespace mapping policies.
  get: true
  create: true

- rest_name: networkaccesspolicy
  descriptions:
    create: Creates a new network access policy.
    get: Retrieves the list of network access policies.
  get: true
  create: true

- rest_name: passwordreset
  descriptions:
    create: Resets the password for an account using the provided link.
    get: Sends a link to the account email to reset the password.
  get: true
  create: true

- rest_name: plan
  descriptions:
    get: Retrieves the list of plans.
  get: true

- rest_name: policy
  descriptions:
    get: Retrieves the list of policy primitives.
  get: true

- rest_name: policyrenderer
  descriptions:
    create: Render a policy of a given type for a given set of tags.
  create: true

- rest_name: processingunit
  descriptions:
    create: Creates a new processing unit.
    get: Retrieves the list of processing units.
  get: true
  create: true

- rest_name: processingunitpolicy
  descriptions:
    create: Creates a new processing unit policy.
    get: Retrieves the list of processing unit policies.
  get: true
  create: true

- rest_name: quotacheck
  descriptions:
    create: Verifies if the quota is exceeded for a particular object.
  create: true

- rest_name: quotapolicy
  descriptions:
    create: Creates a new quota policy.
    get: Retrieves the list of quota policies.
  get: true
  create: true

- rest_name: remoteprocessor
  descriptions:
    create: This should be be here.
  create: true

- rest_name: renderedpolicy
  descriptions:
    create: Render a policy for a processing unit.
  create: true

- rest_name: report
  descriptions:
    create: Create a statistics report.
  create: true

- rest_name: revocation
  descriptions:
    get: Verify the revocation of a certificate according to parameters.
  get: true

- rest_name: role
  descriptions:
    get: Retrieves the list of existing roles.
  get: true

- rest_name: installedapp
  descriptions:
    create: Installs a new app.
    get: Retrieves the list of installed apps.
  get: true
  create: true

- rest_name: statsquery
  descriptions:
    get: Retrieves statistics information based on parameters.
  get: true

- rest_name: suggestedpolicy
  descriptions:
    get: Retrieves a list of network policy suggestion.
  get: true

- rest_name: tabulation
  descriptions:
    get: Retrieves tabulated informations based on parameters.
  get: true

- rest_name: tag
  descriptions:
    get: Retrieves the list of existing tags in the system.
  get: true

- rest_name: taginject
  descriptions:
    create: Internal api to inject tags.
  create: true

- rest_name: token
  descriptions:
    create: Creates a new token.
  create: true

- rest_name: tokenscopepolicy
  descriptions:
    create: Creates a new token scope policy.
    get: Retrieves the list of token scope policies.
  get: true
  create: true

- rest_name: vulnerability
  descriptions:
    create: Creates a new vulnerability.
    get: Retrieves the list of vulnerabilities.
  get: true
  create: true

- rest_name: x509certificate
  descriptions:
    create: Creates a new x509 certificate.
    get: Retrieves a X509 certificates.
  get: true
  create: true

- rest_name: x509certificatecheck
  descriptions:
    get: Verifies if a x509 certificate is valid.
  get: true
