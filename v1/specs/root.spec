children:
- rest_name: account
  get: true
  create: true
- rest_name: accountcheck
  get: true
- rest_name: activate
  get: true
- rest_name: activity
  get: true
- rest_name: alarm
  get: true
  create: true
- rest_name: apiauthorizationpolicy
  get: true
  create: true
- rest_name: apicheck
  create: true
- rest_name: apiservice
  get: true
  create: true
- rest_name: auditprofile
  get: true
  create: true
- rest_name: auth
  get: true
- rest_name: authority
  create: true
- rest_name: automation
  get: true
  create: true
- rest_name: automationtemplate
  get: true
- rest_name: availableservice
  get: true
- rest_name: awsaccount
  get: true
  create: true
- rest_name: certificate
  get: true
  create: true
- rest_name: dependencymap
  get: true
- rest_name: email
  create: true
- rest_name: enforcer
  get: true
  create: true
- rest_name: enforcerprofile
  get: true
  create: true
- rest_name: enforcerprofilemappingpolicy
  get: true
  create: true
- rest_name: export
  get: true
  create: true
- rest_name: externalaccess
  get: true
- rest_name: externalservice
  get: true
  create: true
- rest_name: fileaccess
  get: true
- rest_name: fileaccesspolicy
  get: true
  create: true
- rest_name: filepath
  get: true
  create: true
- rest_name: flowstatistic
  get: true
- rest_name: hookpolicy
  get: true
  create: true
- rest_name: import
  create: true
- rest_name: installation
  get: true
  create: true
- rest_name: isolationprofile
  get: true
  create: true
- rest_name: issue
  create: true
- rest_name: kubernetescluster
  get: true
  create: true
- rest_name: message
  get: true
  create: true
- rest_name: namespace
  get: true
  create: true
- rest_name: namespacemappingpolicy
  get: true
  create: true
- rest_name: networkaccesspolicy
  get: true
  create: true
- rest_name: passwordreset
  get: true
  create: true
- rest_name: plan
  get: true
- rest_name: policy
  get: true
- rest_name: policyrule
  get: true
  create: true
- rest_name: processingunit
  get: true
  create: true
- rest_name: processingunitpolicy
  get: true
  create: true
- rest_name: quotapolicy
  get: true
  create: true
- rest_name: remoteprocessor
  create: true
- rest_name: renderedpolicy
  create: true
- rest_name: report
  create: true
- rest_name: revocation
  get: true
- rest_name: role
  get: true
- rest_name: service
  get: true
  create: true
- rest_name: statsquery
  get: true
- rest_name: suggestedpolicy
  get: true
- rest_name: systemcall
  get: true
  create: true
- rest_name: tabulation
  get: true
- rest_name: tag
  get: true
- rest_name: token
  create: true
- rest_name: tokenscopepolicy
  get: true
  create: true
- rest_name: vulnerability
  get: true
  create: true
- rest_name: x509certificate
  get: true
  create: true
- rest_name: x509certificatecheck
  get: true
model:
  get: true
  description: '[nodoc]'
  entity_name: Root
  extends:
  - '@identifiable-nopk-nostored'
  root: true
  package: root
  resource_name: root
  rest_name: root
