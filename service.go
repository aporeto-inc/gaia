package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ServiceTLSTypeValue represents the possible values for attribute "TLSType".
type ServiceTLSTypeValue string

const (
	// ServiceTLSTypeAporeto represents the value Aporeto.
	ServiceTLSTypeAporeto ServiceTLSTypeValue = "Aporeto"

	// ServiceTLSTypeExternal represents the value External.
	ServiceTLSTypeExternal ServiceTLSTypeValue = "External"

	// ServiceTLSTypeLetsEncrypt represents the value LetsEncrypt.
	ServiceTLSTypeLetsEncrypt ServiceTLSTypeValue = "LetsEncrypt"

	// ServiceTLSTypeNone represents the value None.
	ServiceTLSTypeNone ServiceTLSTypeValue = "None"
)

// ServiceAuthorizationTypeValue represents the possible values for attribute "authorizationType".
type ServiceAuthorizationTypeValue string

const (
	// ServiceAuthorizationTypeJWT represents the value JWT.
	ServiceAuthorizationTypeJWT ServiceAuthorizationTypeValue = "JWT"

	// ServiceAuthorizationTypeMTLS represents the value MTLS.
	ServiceAuthorizationTypeMTLS ServiceAuthorizationTypeValue = "MTLS"

	// ServiceAuthorizationTypeNone represents the value None.
	ServiceAuthorizationTypeNone ServiceAuthorizationTypeValue = "None"

	// ServiceAuthorizationTypeOIDC represents the value OIDC.
	ServiceAuthorizationTypeOIDC ServiceAuthorizationTypeValue = "OIDC"
)

// ServiceTypeValue represents the possible values for attribute "type".
type ServiceTypeValue string

const (
	// ServiceTypeHTTP represents the value HTTP.
	ServiceTypeHTTP ServiceTypeValue = "HTTP"

	// ServiceTypeKubernetesSecrets represents the value KubernetesSecrets.
	ServiceTypeKubernetesSecrets ServiceTypeValue = "KubernetesSecrets"

	// ServiceTypeTCP represents the value TCP.
	ServiceTypeTCP ServiceTypeValue = "TCP"

	// ServiceTypeVaultSecrets represents the value VaultSecrets.
	ServiceTypeVaultSecrets ServiceTypeValue = "VaultSecrets"
)

// ServiceIdentity represents the Identity of the object.
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
	Package:  "squall",
	Private:  false,
}

// ServicesList represents a list of Services
type ServicesList []*Service

// Identity returns the identity of the objects in the list.
func (o ServicesList) Identity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the ServicesList.
func (o ServicesList) Copy() elemental.Identifiables {

	copy := append(ServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServicesList.
func (o ServicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Service))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the ServicesList converted to SparseServicesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ServicesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o ServicesList) Version() int {

	return 1
}

// Service represents the model of a service
type Service struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// IPs is the list of IP addresses where the service can be accessed.
	// This is an optional attribute and is only required if no host names are
	// provided.
	// The system will automatically resolve IP addresses from host names otherwise.
	IPs []string `json:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// PEM encoded certificate that will be used to validate user JWT in HTTP requests.
	// This is an optional field, needed only if the `+"`"+`authorizationType`+"`"+`
	// is set to `+"`"+`JWT`+"`"+`.
	JWTSigningCertificate string `json:"JWTSigningCertificate" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

	// PEM encoded Certificate Authority to use to verify client
	// certificates. This only applies if `+"`"+`authorizationType`+"`"+` is set to
	// `+"`"+`MTLS`+"`"+`. If it is not set, Aporeto's Public Signing Certificate Authority will
	// be used.
	MTLSCertificateAuthority string `json:"MTLSCertificateAuthority" bson:"mtlscertificateauthority" mapstructure:"MTLSCertificateAuthority,omitempty"`

	// This is an advanced setting. Optional OIDC callback URL. If you don't set it,
	// Aporeto will autodiscover it. It will be
	// `+"`"+`https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback`+"`"+`.
	OIDCCallbackURL string `json:"OIDCCallbackURL" bson:"oidccallbackurl" mapstructure:"OIDCCallbackURL,omitempty"`

	// OIDC Client ID. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCClientID string `json:"OIDCClientID" bson:"oidcclientid" mapstructure:"OIDCClientID,omitempty"`

	// OIDC Client Secret. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCClientSecret string `json:"OIDCClientSecret" bson:"oidcclientsecret" mapstructure:"OIDCClientSecret,omitempty"`

	// OIDC Provider URL. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCProviderURL string `json:"OIDCProviderURL" bson:"oidcproviderurl" mapstructure:"OIDCProviderURL,omitempty"`

	// Configures the scopes you want to add to the OIDC provider. Only has effect if
	// `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCScopes []string `json:"OIDCScopes" bson:"oidcscopes" mapstructure:"OIDCScopes,omitempty"`

	// PEM encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `+"`"+`TLSType`+"`"+` is set to `+"`"+`External`+"`"+`.
	TLSCertificate string `json:"TLSCertificate" bson:"tlscertificate" mapstructure:"TLSCertificate,omitempty"`

	// PEM encoded certificate key associated to `+"`"+`TLSCertificate`+"`"+`. Only has effect and
	// required if `+"`"+`TLSType`+"`"+` is set to `+"`"+`External`+"`"+`.
	TLSCertificateKey string `json:"TLSCertificateKey" bson:"tlscertificatekey" mapstructure:"TLSCertificateKey,omitempty"`

	// Set how to provide a server certificate to the service.
	//
	// - `+"`"+`Aporeto`+"`"+`: Generate a certificate issued from Aporeto public CA.
	// - `+"`"+`LetsEncrypt`+"`"+`: Issue a certificate from letsencrypt.
	// - `+"`"+`External`+"`"+`: : Let you define your own certificate and key to use.
	// - `+"`"+`None`+"`"+`: : TLS is disabled (not recommended).
	TLSType ServiceTLSTypeValue `json:"TLSType" bson:"tlstype" mapstructure:"TLSType,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags []string `json:"-" bson:"allapitags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags []string `json:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// AuthorizationType defines the user authorization type that should be used.
	//
	// - `+"`"+`None`+"`"+`: No auhtorization.
	// - `+"`"+`JWT`+"`"+`:  Configures a simple JWT verification from the HTTP `+"`"+`Auhorization`+"`"+`
	// Header
	// - `+"`"+`OIDC`+"`"+`: Configures OIDC authorization. You must then set `+"`"+`OIDCClientID`+"`"+`,
	// `+"`"+`OIDCClientSecret`+"`"+`, OIDCProviderURL`+"`"+`.
	// - `+"`"+`MTLS`+"`"+`: Configures Client Certificate authorization. Then you can optionaly
	// `+"`"+`MTLSCertificateAuthority`+"`"+` otherwise Aporeto Public Signing Certificate will be
	// used.
	AuthorizationType ServiceAuthorizationTypeValue `json:"authorizationType" bson:"authorizationtype" mapstructure:"authorizationType,omitempty"`

	// Defines a list of mappings between claims and
	// HTTP headers. When these mappings are defined, the enforcer will copy the
	// values of the claims to the corresponding HTTP headers.
	ClaimsToHTTPHeaderMappings []*ClaimMapping `json:"claimsToHTTPHeaderMappings" bson:"claimstohttpheadermappings" mapstructure:"claimsToHTTPHeaderMappings,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Endpoints is a read only attribute that actually resolves the API
	// endpoints that the service is exposing. Only valid during policy rendering.
	Endpoints []*Endpoint `json:"endpoints" bson:"-" mapstructure:"endpoints,omitempty"`

	// ExposedAPIs contains a tag expression that will determine which
	// APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
	// similar specifications for other L7 protocols.
	ExposedAPIs [][]string `json:"exposedAPIs" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// ExposedPort is the port that the service can be accessed. Note that
	// this is different from the Port attribute that describes the port that the
	// service is actually listening. For example if a load balancer is used, the
	// ExposedPort is the port that the load balancer is listening for the service,
	// whereas the port that the implementation is listening can be different.
	ExposedPort int `json:"exposedPort" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
	// the enforcer has to initiate a TLS session in order to forrward traffic to the
	// service.
	ExposedServiceIsTLS bool `json:"exposedServiceIsTLS" bson:"exposedserviceistls" mapstructure:"exposedServiceIsTLS,omitempty"`

	// External is a boolean that indicates if this is an external service.
	External bool `json:"external" bson:"external" mapstructure:"external,omitempty"`

	// Hosts are the names that the service can be accessed with.
	Hosts []string `json:"hosts" bson:"hosts" mapstructure:"hosts,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Port is the port that the implementation of the service is listening to and
	// it can be different than the exposedPorts describing the service. This is needed
	// for port mapping use cases where there is private and public ports.
	Port int `json:"port" bson:"port" mapstructure:"port,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// PublicApplicationPort is a new virtual port that the service can
	// be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
	// application path, you might want to declare a new port where the enforcer
	// listens for TLS. However, the application does not need to be modified and
	// the enforcer will map the traffic to the correct application port. This useful
	// when an application is being accessed from a public network.
	PublicApplicationPort int `json:"publicApplicationPort" bson:"publicapplicationport" mapstructure:"publicApplicationPort,omitempty"`

	// If this is set, the user will be redirected to that URL in case of any
	// authorization failure to let you chance to provide a nice message to the user.
	// The query parameter `+"`"+`?failure_message=<message>`+"`"+` will be added to that url
	// explaining the possible reasons of the failure.
	RedirectURLOnAuthorizationFailure string `json:"redirectURLOnAuthorizationFailure" bson:"redirecturlonauthorizationfailure" mapstructure:"redirectURLOnAuthorizationFailure,omitempty"`

	// Selectors contains the tag expression that an a processing unit
	// must match in order to implement this particular service.
	Selectors [][]string `json:"selectors" bson:"selectors" mapstructure:"selectors,omitempty"`

	// PEM encoded Certificate Authorities to trust when additional hops are needed. It
	// must be set if the service must reach a Service marked as `+"`"+`external`+"`"+` or must go
	// through an additional TLS termination point like a L7 Load Balancer.
	TrustedCertificateAuthorities string `json:"trustedCertificateAuthorities" bson:"trustedcertificateauthorities" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type is the type of the service.
	Type ServiceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion:               1,
		Mutex:                      &sync.Mutex{},
		AllAPITags:                 []string{},
		Annotations:                map[string][]string{},
		AllServiceTags:             []string{},
		AssociatedTags:             []string{},
		ExposedAPIs:                [][]string{},
		ExposedServiceIsTLS:        false,
		External:                   false,
		Hosts:                      []string{},
		ClaimsToHTTPHeaderMappings: []*ClaimMapping{},
		Endpoints:                  []*Endpoint{},
		AuthorizationType:          ServiceAuthorizationTypeNone,
		OIDCScopes:                 []string{},
		Selectors:                  [][]string{},
		Type:                       ServiceTypeHTTP,
		TLSType:                    ServiceTLSTypeAporeto,
		Metadata:                   []string{},
		IPs:                        []string{},
		NormalizedTags:             []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Service) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Service) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Service) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Service) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Service) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Service) Doc() string {
	return `A Service defines a generic service object at L4 or L7 that encapsulates the
description of a micro-service. A service exposes APIs and can be implemented
through third party entities (such as a cloud provider) or through  processing
units.`
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Service) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Service) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *Service) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *Service) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Service) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Service) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Service) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Service) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Service) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Service) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *Service) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *Service) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *Service) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Service) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Service) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Service) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Service) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Service) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Service) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Service) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Service) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Service) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Service) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Service) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Service) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Service) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Service) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Service) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Service) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseService{
			ID:                                &o.ID,
			IPs:                               &o.IPs,
			JWTSigningCertificate:             &o.JWTSigningCertificate,
			MTLSCertificateAuthority:          &o.MTLSCertificateAuthority,
			OIDCCallbackURL:                   &o.OIDCCallbackURL,
			OIDCClientID:                      &o.OIDCClientID,
			OIDCClientSecret:                  &o.OIDCClientSecret,
			OIDCProviderURL:                   &o.OIDCProviderURL,
			OIDCScopes:                        &o.OIDCScopes,
			TLSCertificate:                    &o.TLSCertificate,
			TLSCertificateKey:                 &o.TLSCertificateKey,
			TLSType:                           &o.TLSType,
			AllAPITags:                        &o.AllAPITags,
			AllServiceTags:                    &o.AllServiceTags,
			Annotations:                       &o.Annotations,
			Archived:                          &o.Archived,
			AssociatedTags:                    &o.AssociatedTags,
			AuthorizationType:                 &o.AuthorizationType,
			ClaimsToHTTPHeaderMappings:        &o.ClaimsToHTTPHeaderMappings,
			CreateTime:                        &o.CreateTime,
			Description:                       &o.Description,
			Disabled:                          &o.Disabled,
			Endpoints:                         &o.Endpoints,
			ExposedAPIs:                       &o.ExposedAPIs,
			ExposedPort:                       &o.ExposedPort,
			ExposedServiceIsTLS:               &o.ExposedServiceIsTLS,
			External:                          &o.External,
			Hosts:                             &o.Hosts,
			Metadata:                          &o.Metadata,
			Name:                              &o.Name,
			Namespace:                         &o.Namespace,
			NormalizedTags:                    &o.NormalizedTags,
			Port:                              &o.Port,
			Protected:                         &o.Protected,
			PublicApplicationPort:             &o.PublicApplicationPort,
			RedirectURLOnAuthorizationFailure: &o.RedirectURLOnAuthorizationFailure,
			Selectors:                         &o.Selectors,
			TrustedCertificateAuthorities:     &o.TrustedCertificateAuthorities,
			Type:                              &o.Type,
			UpdateTime:                        &o.UpdateTime,
			ZHash:                             &o.ZHash,
			Zone:                              &o.Zone,
		}
	}

	sp := &SparseService{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "IPs":
			sp.IPs = &(o.IPs)
		case "JWTSigningCertificate":
			sp.JWTSigningCertificate = &(o.JWTSigningCertificate)
		case "MTLSCertificateAuthority":
			sp.MTLSCertificateAuthority = &(o.MTLSCertificateAuthority)
		case "OIDCCallbackURL":
			sp.OIDCCallbackURL = &(o.OIDCCallbackURL)
		case "OIDCClientID":
			sp.OIDCClientID = &(o.OIDCClientID)
		case "OIDCClientSecret":
			sp.OIDCClientSecret = &(o.OIDCClientSecret)
		case "OIDCProviderURL":
			sp.OIDCProviderURL = &(o.OIDCProviderURL)
		case "OIDCScopes":
			sp.OIDCScopes = &(o.OIDCScopes)
		case "TLSCertificate":
			sp.TLSCertificate = &(o.TLSCertificate)
		case "TLSCertificateKey":
			sp.TLSCertificateKey = &(o.TLSCertificateKey)
		case "TLSType":
			sp.TLSType = &(o.TLSType)
		case "allAPITags":
			sp.AllAPITags = &(o.AllAPITags)
		case "allServiceTags":
			sp.AllServiceTags = &(o.AllServiceTags)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "archived":
			sp.Archived = &(o.Archived)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "authorizationType":
			sp.AuthorizationType = &(o.AuthorizationType)
		case "claimsToHTTPHeaderMappings":
			sp.ClaimsToHTTPHeaderMappings = &(o.ClaimsToHTTPHeaderMappings)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "endpoints":
			sp.Endpoints = &(o.Endpoints)
		case "exposedAPIs":
			sp.ExposedAPIs = &(o.ExposedAPIs)
		case "exposedPort":
			sp.ExposedPort = &(o.ExposedPort)
		case "exposedServiceIsTLS":
			sp.ExposedServiceIsTLS = &(o.ExposedServiceIsTLS)
		case "external":
			sp.External = &(o.External)
		case "hosts":
			sp.Hosts = &(o.Hosts)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "port":
			sp.Port = &(o.Port)
		case "protected":
			sp.Protected = &(o.Protected)
		case "publicApplicationPort":
			sp.PublicApplicationPort = &(o.PublicApplicationPort)
		case "redirectURLOnAuthorizationFailure":
			sp.RedirectURLOnAuthorizationFailure = &(o.RedirectURLOnAuthorizationFailure)
		case "selectors":
			sp.Selectors = &(o.Selectors)
		case "trustedCertificateAuthorities":
			sp.TrustedCertificateAuthorities = &(o.TrustedCertificateAuthorities)
		case "type":
			sp.Type = &(o.Type)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseService to the object.
func (o *Service) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseService)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.IPs != nil {
		o.IPs = *so.IPs
	}
	if so.JWTSigningCertificate != nil {
		o.JWTSigningCertificate = *so.JWTSigningCertificate
	}
	if so.MTLSCertificateAuthority != nil {
		o.MTLSCertificateAuthority = *so.MTLSCertificateAuthority
	}
	if so.OIDCCallbackURL != nil {
		o.OIDCCallbackURL = *so.OIDCCallbackURL
	}
	if so.OIDCClientID != nil {
		o.OIDCClientID = *so.OIDCClientID
	}
	if so.OIDCClientSecret != nil {
		o.OIDCClientSecret = *so.OIDCClientSecret
	}
	if so.OIDCProviderURL != nil {
		o.OIDCProviderURL = *so.OIDCProviderURL
	}
	if so.OIDCScopes != nil {
		o.OIDCScopes = *so.OIDCScopes
	}
	if so.TLSCertificate != nil {
		o.TLSCertificate = *so.TLSCertificate
	}
	if so.TLSCertificateKey != nil {
		o.TLSCertificateKey = *so.TLSCertificateKey
	}
	if so.TLSType != nil {
		o.TLSType = *so.TLSType
	}
	if so.AllAPITags != nil {
		o.AllAPITags = *so.AllAPITags
	}
	if so.AllServiceTags != nil {
		o.AllServiceTags = *so.AllServiceTags
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.AuthorizationType != nil {
		o.AuthorizationType = *so.AuthorizationType
	}
	if so.ClaimsToHTTPHeaderMappings != nil {
		o.ClaimsToHTTPHeaderMappings = *so.ClaimsToHTTPHeaderMappings
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Endpoints != nil {
		o.Endpoints = *so.Endpoints
	}
	if so.ExposedAPIs != nil {
		o.ExposedAPIs = *so.ExposedAPIs
	}
	if so.ExposedPort != nil {
		o.ExposedPort = *so.ExposedPort
	}
	if so.ExposedServiceIsTLS != nil {
		o.ExposedServiceIsTLS = *so.ExposedServiceIsTLS
	}
	if so.External != nil {
		o.External = *so.External
	}
	if so.Hosts != nil {
		o.Hosts = *so.Hosts
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Port != nil {
		o.Port = *so.Port
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.PublicApplicationPort != nil {
		o.PublicApplicationPort = *so.PublicApplicationPort
	}
	if so.RedirectURLOnAuthorizationFailure != nil {
		o.RedirectURLOnAuthorizationFailure = *so.RedirectURLOnAuthorizationFailure
	}
	if so.Selectors != nil {
		o.Selectors = *so.Selectors
	}
	if so.TrustedCertificateAuthorities != nil {
		o.TrustedCertificateAuthorities = *so.TrustedCertificateAuthorities
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Service.
func (o *Service) DeepCopy() *Service {

	if o == nil {
		return nil
	}

	out := &Service{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Service.
func (o *Service) DeepCopyInto(out *Service) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Service: %s", err))
	}

	*out = *target.(*Service)
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("TLSType", string(o.TLSType), []string{"Aporeto", "LetsEncrypt", "External", "None"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("authorizationType", string(o.AuthorizationType), []string{"None", "JWT", "OIDC", "MTLS"}, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.ClaimsToHTTPHeaderMappings {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.Endpoints {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateRequiredInt("exposedPort", o.ExposedPort); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("exposedPort", o.ExposedPort, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("port", o.Port); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("port", o.Port, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("publicApplicationPort", o.PublicApplicationPort, int(65535), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"}, false); err != nil {
		errors = append(errors, err)
	}

	// Custom object validation.
	if err := ValidateServiceEntity(o); err != nil {
		switch e := err.(type) {
		case elemental.Errors:
			errors = append(errors, e...)
		case elemental.Error:
			errors = append(errors, e)
		default:
			errors = append(errors, e)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Service) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Service) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServiceAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Service) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "IPs":
		return o.IPs
	case "JWTSigningCertificate":
		return o.JWTSigningCertificate
	case "MTLSCertificateAuthority":
		return o.MTLSCertificateAuthority
	case "OIDCCallbackURL":
		return o.OIDCCallbackURL
	case "OIDCClientID":
		return o.OIDCClientID
	case "OIDCClientSecret":
		return o.OIDCClientSecret
	case "OIDCProviderURL":
		return o.OIDCProviderURL
	case "OIDCScopes":
		return o.OIDCScopes
	case "TLSCertificate":
		return o.TLSCertificate
	case "TLSCertificateKey":
		return o.TLSCertificateKey
	case "TLSType":
		return o.TLSType
	case "allAPITags":
		return o.AllAPITags
	case "allServiceTags":
		return o.AllServiceTags
	case "annotations":
		return o.Annotations
	case "archived":
		return o.Archived
	case "associatedTags":
		return o.AssociatedTags
	case "authorizationType":
		return o.AuthorizationType
	case "claimsToHTTPHeaderMappings":
		return o.ClaimsToHTTPHeaderMappings
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "endpoints":
		return o.Endpoints
	case "exposedAPIs":
		return o.ExposedAPIs
	case "exposedPort":
		return o.ExposedPort
	case "exposedServiceIsTLS":
		return o.ExposedServiceIsTLS
	case "external":
		return o.External
	case "hosts":
		return o.Hosts
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "port":
		return o.Port
	case "protected":
		return o.Protected
	case "publicApplicationPort":
		return o.PublicApplicationPort
	case "redirectURLOnAuthorizationFailure":
		return o.RedirectURLOnAuthorizationFailure
	case "selectors":
		return o.Selectors
	case "trustedCertificateAuthorities":
		return o.TrustedCertificateAuthorities
	case "type":
		return o.Type
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ServiceAttributesMap represents the map of attribute for Service.
var ServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"IPs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"JWTSigningCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `PEM encoded certificate that will be used to validate user JWT in HTTP requests.
This is an optional field, needed only if the ` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `JWT` + "`" + `.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"MTLSCertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MTLSCertificateAuthority",
		Description: `PEM encoded Certificate Authority to use to verify client
certificates. This only applies if ` + "`" + `authorizationType` + "`" + ` is set to
` + "`" + `MTLS` + "`" + `. If it is not set, Aporeto's Public Signing Certificate Authority will
be used.`,
		Exposed: true,
		Name:    "MTLSCertificateAuthority",
		Stored:  true,
		Type:    "string",
	},
	"OIDCCallbackURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCCallbackURL",
		Description: `This is an advanced setting. Optional OIDC callback URL. If you don't set it,
Aporeto will autodiscover it. It will be
` + "`" + `https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCCallbackURL",
		Stored:  true,
		Type:    "string",
	},
	"OIDCClientID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientID",
		Description:    `OIDC Client ID. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientID",
		Stored:         true,
		Type:           "string",
	},
	"OIDCClientSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientSecret",
		Description:    `OIDC Client Secret. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientSecret",
		Stored:         true,
		Type:           "string",
	},
	"OIDCProviderURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCProviderURL",
		Description:    `OIDC Provider URL. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCProviderURL",
		Stored:         true,
		Type:           "string",
	},
	"OIDCScopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCScopes",
		Description: `Configures the scopes you want to add to the OIDC provider. Only has effect if
` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCScopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"TLSCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificate",
		Description: `PEM encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificate",
		Stored:  true,
		Type:    "string",
	},
	"TLSCertificateKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificateKey",
		Description: `PEM encoded certificate key associated to ` + "`" + `TLSCertificate` + "`" + `. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificateKey",
		Stored:  true,
		Type:    "string",
	},
	"TLSType": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "LetsEncrypt", "External", "None"},
		ConvertedName:  "TLSType",
		DefaultValue:   ServiceTLSTypeAporeto,
		Description: `Set how to provide a server certificate to the service.

- ` + "`" + `Aporeto` + "`" + `: Generate a certificate issued from Aporeto public CA.
- ` + "`" + `LetsEncrypt` + "`" + `: Issue a certificate from letsencrypt.
- ` + "`" + `External` + "`" + `: : Let you define your own certificate and key to use.
- ` + "`" + `None` + "`" + `: : TLS is disabled (not recommended).`,
		Exposed: true,
		Name:    "TLSType",
		Stored:  true,
		Type:    "enum",
	},
	"AllAPITags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AllServiceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"Archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"AuthorizationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"None", "JWT", "OIDC", "MTLS"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `AuthorizationType defines the user authorization type that should be used.

- ` + "`" + `None` + "`" + `: No auhtorization.
- ` + "`" + `JWT` + "`" + `:  Configures a simple JWT verification from the HTTP ` + "`" + `Auhorization` + "`" + `
Header
- ` + "`" + `OIDC` + "`" + `: Configures OIDC authorization. You must then set ` + "`" + `OIDCClientID` + "`" + `,
` + "`" + `OIDCClientSecret` + "`" + `, OIDCProviderURL` + "`" + `.
- ` + "`" + `MTLS` + "`" + `: Configures Client Certificate authorization. Then you can optionaly
` + "`" + `MTLSCertificateAuthority` + "`" + ` otherwise Aporeto Public Signing Certificate will be
used.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"ClaimsToHTTPHeaderMappings": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClaimsToHTTPHeaderMappings",
		Description: `Defines a list of mappings between claims and
HTTP headers. When these mappings are defined, the enforcer will copy the
values of the claims to the corresponding HTTP headers.`,
		Exposed: true,
		Name:    "claimsToHTTPHeaderMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
	},
	"ExposedAPIs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `ExposedAPIs contains a tag expression that will determine which
APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
similar specifications for other L7 protocols.`,
		Exposed: true,
		Name:    "exposedAPIs",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"ExposedPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `ExposedPort is the port that the service can be accessed. Note that
this is different from the Port attribute that describes the port that the
service is actually listening. For example if a load balancer is used, the
ExposedPort is the port that the load balancer is listening for the service,
whereas the port that the implementation is listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"ExposedServiceIsTLS": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
the enforcer has to initiate a TLS session in order to forrward traffic to the
service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"External": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "External",
		Description:    `External is a boolean that indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Hosts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `Hosts are the names that the service can be accessed with.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description: `Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PublicApplicationPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PublicApplicationPort",
		Description: `PublicApplicationPort is a new virtual port that the service can
be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
application path, you might want to declare a new port where the enforcer
listens for TLS. However, the application does not need to be modified and
the enforcer will map the traffic to the correct application port. This useful
when an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"RedirectURLOnAuthorizationFailure": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURLOnAuthorizationFailure",
		Description: `If this is set, the user will be redirected to that URL in case of any
authorization failure to let you chance to provide a nice message to the user.
The query parameter ` + "`" + `?failure_message=<message>` + "`" + ` will be added to that url
explaining the possible reasons of the failure.`,
		Exposed: true,
		Name:    "redirectURLOnAuthorizationFailure",
		Stored:  true,
		Type:    "string",
	},
	"Selectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.`,
		Exposed: true,
		Name:    "selectors",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"TrustedCertificateAuthorities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM encoded Certificate Authorities to trust when additional hops are needed. It
must be set if the service must reach a Service marked as ` + "`" + `external` + "`" + ` or must go
through an additional TLS termination point like a L7 Load Balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type is the type of the service.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// ServiceLowerCaseAttributesMap represents the map of attribute for Service.
var ServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ips": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPs",
		Description: `IPs is the list of IP addresses where the service can be accessed.
This is an optional attribute and is only required if no host names are
provided.
The system will automatically resolve IP addresses from host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"jwtsigningcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `PEM encoded certificate that will be used to validate user JWT in HTTP requests.
This is an optional field, needed only if the ` + "`" + `authorizationType` + "`" + `
is set to ` + "`" + `JWT` + "`" + `.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"mtlscertificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MTLSCertificateAuthority",
		Description: `PEM encoded Certificate Authority to use to verify client
certificates. This only applies if ` + "`" + `authorizationType` + "`" + ` is set to
` + "`" + `MTLS` + "`" + `. If it is not set, Aporeto's Public Signing Certificate Authority will
be used.`,
		Exposed: true,
		Name:    "MTLSCertificateAuthority",
		Stored:  true,
		Type:    "string",
	},
	"oidccallbackurl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCCallbackURL",
		Description: `This is an advanced setting. Optional OIDC callback URL. If you don't set it,
Aporeto will autodiscover it. It will be
` + "`" + `https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCCallbackURL",
		Stored:  true,
		Type:    "string",
	},
	"oidcclientid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientID",
		Description:    `OIDC Client ID. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientID",
		Stored:         true,
		Type:           "string",
	},
	"oidcclientsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCClientSecret",
		Description:    `OIDC Client Secret. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCClientSecret",
		Stored:         true,
		Type:           "string",
	},
	"oidcproviderurl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCProviderURL",
		Description:    `OIDC Provider URL. Only has effect if the ` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed:        true,
		Name:           "OIDCProviderURL",
		Stored:         true,
		Type:           "string",
	},
	"oidcscopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OIDCScopes",
		Description: `Configures the scopes you want to add to the OIDC provider. Only has effect if
` + "`" + `authorizationType` + "`" + ` is set to ` + "`" + `OIDC` + "`" + `.`,
		Exposed: true,
		Name:    "OIDCScopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"tlscertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificate",
		Description: `PEM encoded certificate to expose to the clients for TLS. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificate",
		Stored:  true,
		Type:    "string",
	},
	"tlscertificatekey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TLSCertificateKey",
		Description: `PEM encoded certificate key associated to ` + "`" + `TLSCertificate` + "`" + `. Only has effect and
required if ` + "`" + `TLSType` + "`" + ` is set to ` + "`" + `External` + "`" + `.`,
		Exposed: true,
		Name:    "TLSCertificateKey",
		Stored:  true,
		Type:    "string",
	},
	"tlstype": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "LetsEncrypt", "External", "None"},
		ConvertedName:  "TLSType",
		DefaultValue:   ServiceTLSTypeAporeto,
		Description: `Set how to provide a server certificate to the service.

- ` + "`" + `Aporeto` + "`" + `: Generate a certificate issued from Aporeto public CA.
- ` + "`" + `LetsEncrypt` + "`" + `: Issue a certificate from letsencrypt.
- ` + "`" + `External` + "`" + `: : Let you define your own certificate and key to use.
- ` + "`" + `None` + "`" + `: : TLS is disabled (not recommended).`,
		Exposed: true,
		Name:    "TLSType",
		Stored:  true,
		Type:    "enum",
	},
	"allapitags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"allservicetags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"archived": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Archived",
		Description:    `Archived defines if the object is archived.`,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"authorizationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"None", "JWT", "OIDC", "MTLS"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `AuthorizationType defines the user authorization type that should be used.

- ` + "`" + `None` + "`" + `: No auhtorization.
- ` + "`" + `JWT` + "`" + `:  Configures a simple JWT verification from the HTTP ` + "`" + `Auhorization` + "`" + `
Header
- ` + "`" + `OIDC` + "`" + `: Configures OIDC authorization. You must then set ` + "`" + `OIDCClientID` + "`" + `,
` + "`" + `OIDCClientSecret` + "`" + `, OIDCProviderURL` + "`" + `.
- ` + "`" + `MTLS` + "`" + `: Configures Client Certificate authorization. Then you can optionaly
` + "`" + `MTLSCertificateAuthority` + "`" + ` otherwise Aporeto Public Signing Certificate will be
used.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"claimstohttpheadermappings": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ClaimsToHTTPHeaderMappings",
		Description: `Defines a list of mappings between claims and
HTTP headers. When these mappings are defined, the enforcer will copy the
values of the claims to the corresponding HTTP headers.`,
		Exposed: true,
		Name:    "claimsToHTTPHeaderMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "endpoint",
		Type:     "refList",
	},
	"exposedapis": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIs",
		Description: `ExposedAPIs contains a tag expression that will determine which
APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
similar specifications for other L7 protocols.`,
		Exposed: true,
		Name:    "exposedAPIs",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"exposedport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedPort",
		Description: `ExposedPort is the port that the service can be accessed. Note that
this is different from the Port attribute that describes the port that the
service is actually listening. For example if a load balancer is used, the
ExposedPort is the port that the load balancer is listening for the service,
whereas the port that the implementation is listening can be different.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "exposedPort",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"exposedserviceistls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServiceIsTLS",
		Description: `ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
the enforcer has to initiate a TLS session in order to forrward traffic to the
service.`,
		Exposed:    true,
		Filterable: true,
		Name:       "exposedServiceIsTLS",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"external": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "External",
		Description:    `External is a boolean that indicates if this is an external service.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "external",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"hosts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Hosts",
		Description:    `Hosts are the names that the service can be accessed with.`,
		Exposed:        true,
		Name:           "hosts",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"port": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Port",
		Description: `Port is the port that the implementation of the service is listening to and
it can be different than the exposedPorts describing the service. This is needed
for port mapping use cases where there is private and public ports.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "port",
		Required: true,
		Stored:   true,
		Type:     "integer",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"publicapplicationport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PublicApplicationPort",
		Description: `PublicApplicationPort is a new virtual port that the service can
be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
application path, you might want to declare a new port where the enforcer
listens for TLS. However, the application does not need to be modified and
the enforcer will map the traffic to the correct application port. This useful
when an application is being accessed from a public network.`,
		Exposed:  true,
		MaxValue: 65535,
		Name:     "publicApplicationPort",
		Stored:   true,
		Type:     "integer",
	},
	"redirecturlonauthorizationfailure": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURLOnAuthorizationFailure",
		Description: `If this is set, the user will be redirected to that URL in case of any
authorization failure to let you chance to provide a nice message to the user.
The query parameter ` + "`" + `?failure_message=<message>` + "`" + ` will be added to that url
explaining the possible reasons of the failure.`,
		Exposed: true,
		Name:    "redirectURLOnAuthorizationFailure",
		Stored:  true,
		Type:    "string",
	},
	"selectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Selectors",
		Description: `Selectors contains the tag expression that an a processing unit
must match in order to implement this particular service.`,
		Exposed: true,
		Name:    "selectors",
		Stored:  true,
		SubType: "[][]string",
		Type:    "external",
	},
	"trustedcertificateauthorities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCertificateAuthorities",
		Description: `PEM encoded Certificate Authorities to trust when additional hops are needed. It
must be set if the service must reach a Service marked as ` + "`" + `external` + "`" + ` or must go
through an additional TLS termination point like a L7 Load Balancer.`,
		Exposed: true,
		Name:    "trustedCertificateAuthorities",
		Stored:  true,
		Type:    "string",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"HTTP", "TCP", "KubernetesSecrets", "VaultSecrets"},
		ConvertedName:  "Type",
		DefaultValue:   ServiceTypeHTTP,
		Description:    `Type is the type of the service.`,
		Exposed:        true,
		Name:           "type",
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseServicesList represents a list of SparseServices
type SparseServicesList []*SparseService

// Identity returns the identity of the objects in the list.
func (o SparseServicesList) Identity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the SparseServicesList.
func (o SparseServicesList) Copy() elemental.Identifiables {

	copy := append(SparseServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseServicesList.
func (o SparseServicesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseService))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseServicesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseServicesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseServicesList converted to ServicesList.
func (o SparseServicesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseServicesList) Version() int {

	return 1
}

// SparseService represents the sparse version of a service.
type SparseService struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// IPs is the list of IP addresses where the service can be accessed.
	// This is an optional attribute and is only required if no host names are
	// provided.
	// The system will automatically resolve IP addresses from host names otherwise.
	IPs *[]string `json:"IPs,omitempty" bson:"ips" mapstructure:"IPs,omitempty"`

	// PEM encoded certificate that will be used to validate user JWT in HTTP requests.
	// This is an optional field, needed only if the `+"`"+`authorizationType`+"`"+`
	// is set to `+"`"+`JWT`+"`"+`.
	JWTSigningCertificate *string `json:"JWTSigningCertificate,omitempty" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

	// PEM encoded Certificate Authority to use to verify client
	// certificates. This only applies if `+"`"+`authorizationType`+"`"+` is set to
	// `+"`"+`MTLS`+"`"+`. If it is not set, Aporeto's Public Signing Certificate Authority will
	// be used.
	MTLSCertificateAuthority *string `json:"MTLSCertificateAuthority,omitempty" bson:"mtlscertificateauthority" mapstructure:"MTLSCertificateAuthority,omitempty"`

	// This is an advanced setting. Optional OIDC callback URL. If you don't set it,
	// Aporeto will autodiscover it. It will be
	// `+"`"+`https://<hosts[0]|IPs[0]>/.aporeto/oidc/callback`+"`"+`.
	OIDCCallbackURL *string `json:"OIDCCallbackURL,omitempty" bson:"oidccallbackurl" mapstructure:"OIDCCallbackURL,omitempty"`

	// OIDC Client ID. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCClientID *string `json:"OIDCClientID,omitempty" bson:"oidcclientid" mapstructure:"OIDCClientID,omitempty"`

	// OIDC Client Secret. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCClientSecret *string `json:"OIDCClientSecret,omitempty" bson:"oidcclientsecret" mapstructure:"OIDCClientSecret,omitempty"`

	// OIDC Provider URL. Only has effect if the `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCProviderURL *string `json:"OIDCProviderURL,omitempty" bson:"oidcproviderurl" mapstructure:"OIDCProviderURL,omitempty"`

	// Configures the scopes you want to add to the OIDC provider. Only has effect if
	// `+"`"+`authorizationType`+"`"+` is set to `+"`"+`OIDC`+"`"+`.
	OIDCScopes *[]string `json:"OIDCScopes,omitempty" bson:"oidcscopes" mapstructure:"OIDCScopes,omitempty"`

	// PEM encoded certificate to expose to the clients for TLS. Only has effect and
	// required if `+"`"+`TLSType`+"`"+` is set to `+"`"+`External`+"`"+`.
	TLSCertificate *string `json:"TLSCertificate,omitempty" bson:"tlscertificate" mapstructure:"TLSCertificate,omitempty"`

	// PEM encoded certificate key associated to `+"`"+`TLSCertificate`+"`"+`. Only has effect and
	// required if `+"`"+`TLSType`+"`"+` is set to `+"`"+`External`+"`"+`.
	TLSCertificateKey *string `json:"TLSCertificateKey,omitempty" bson:"tlscertificatekey" mapstructure:"TLSCertificateKey,omitempty"`

	// Set how to provide a server certificate to the service.
	//
	// - `+"`"+`Aporeto`+"`"+`: Generate a certificate issued from Aporeto public CA.
	// - `+"`"+`LetsEncrypt`+"`"+`: Issue a certificate from letsencrypt.
	// - `+"`"+`External`+"`"+`: : Let you define your own certificate and key to use.
	// - `+"`"+`None`+"`"+`: : TLS is disabled (not recommended).
	TLSType *ServiceTLSTypeValue `json:"TLSType,omitempty" bson:"tlstype" mapstructure:"TLSType,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags *[]string `json:"-" bson:"allapitags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags *[]string `json:"-" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived *bool `json:"-" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// AuthorizationType defines the user authorization type that should be used.
	//
	// - `+"`"+`None`+"`"+`: No auhtorization.
	// - `+"`"+`JWT`+"`"+`:  Configures a simple JWT verification from the HTTP `+"`"+`Auhorization`+"`"+`
	// Header
	// - `+"`"+`OIDC`+"`"+`: Configures OIDC authorization. You must then set `+"`"+`OIDCClientID`+"`"+`,
	// `+"`"+`OIDCClientSecret`+"`"+`, OIDCProviderURL`+"`"+`.
	// - `+"`"+`MTLS`+"`"+`: Configures Client Certificate authorization. Then you can optionaly
	// `+"`"+`MTLSCertificateAuthority`+"`"+` otherwise Aporeto Public Signing Certificate will be
	// used.
	AuthorizationType *ServiceAuthorizationTypeValue `json:"authorizationType,omitempty" bson:"authorizationtype" mapstructure:"authorizationType,omitempty"`

	// Defines a list of mappings between claims and
	// HTTP headers. When these mappings are defined, the enforcer will copy the
	// values of the claims to the corresponding HTTP headers.
	ClaimsToHTTPHeaderMappings *[]*ClaimMapping `json:"claimsToHTTPHeaderMappings,omitempty" bson:"claimstohttpheadermappings" mapstructure:"claimsToHTTPHeaderMappings,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Endpoints is a read only attribute that actually resolves the API
	// endpoints that the service is exposing. Only valid during policy rendering.
	Endpoints *[]*Endpoint `json:"endpoints,omitempty" bson:"-" mapstructure:"endpoints,omitempty"`

	// ExposedAPIs contains a tag expression that will determine which
	// APIs a service is exposing. The APIs can be defined as the RESTAPISpec or
	// similar specifications for other L7 protocols.
	ExposedAPIs *[][]string `json:"exposedAPIs,omitempty" bson:"exposedapis" mapstructure:"exposedAPIs,omitempty"`

	// ExposedPort is the port that the service can be accessed. Note that
	// this is different from the Port attribute that describes the port that the
	// service is actually listening. For example if a load balancer is used, the
	// ExposedPort is the port that the load balancer is listening for the service,
	// whereas the port that the implementation is listening can be different.
	ExposedPort *int `json:"exposedPort,omitempty" bson:"exposedport" mapstructure:"exposedPort,omitempty"`

	// ExposedServiceIsTLS indicates that the exposed service is TLS. This means that
	// the enforcer has to initiate a TLS session in order to forrward traffic to the
	// service.
	ExposedServiceIsTLS *bool `json:"exposedServiceIsTLS,omitempty" bson:"exposedserviceistls" mapstructure:"exposedServiceIsTLS,omitempty"`

	// External is a boolean that indicates if this is an external service.
	External *bool `json:"external,omitempty" bson:"external" mapstructure:"external,omitempty"`

	// Hosts are the names that the service can be accessed with.
	Hosts *[]string `json:"hosts,omitempty" bson:"hosts" mapstructure:"hosts,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Port is the port that the implementation of the service is listening to and
	// it can be different than the exposedPorts describing the service. This is needed
	// for port mapping use cases where there is private and public ports.
	Port *int `json:"port,omitempty" bson:"port" mapstructure:"port,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// PublicApplicationPort is a new virtual port that the service can
	// be accessed, using HTTPs. Since the enforcer transparently inserts TLS in the
	// application path, you might want to declare a new port where the enforcer
	// listens for TLS. However, the application does not need to be modified and
	// the enforcer will map the traffic to the correct application port. This useful
	// when an application is being accessed from a public network.
	PublicApplicationPort *int `json:"publicApplicationPort,omitempty" bson:"publicapplicationport" mapstructure:"publicApplicationPort,omitempty"`

	// If this is set, the user will be redirected to that URL in case of any
	// authorization failure to let you chance to provide a nice message to the user.
	// The query parameter `+"`"+`?failure_message=<message>`+"`"+` will be added to that url
	// explaining the possible reasons of the failure.
	RedirectURLOnAuthorizationFailure *string `json:"redirectURLOnAuthorizationFailure,omitempty" bson:"redirecturlonauthorizationfailure" mapstructure:"redirectURLOnAuthorizationFailure,omitempty"`

	// Selectors contains the tag expression that an a processing unit
	// must match in order to implement this particular service.
	Selectors *[][]string `json:"selectors,omitempty" bson:"selectors" mapstructure:"selectors,omitempty"`

	// PEM encoded Certificate Authorities to trust when additional hops are needed. It
	// must be set if the service must reach a Service marked as `+"`"+`external`+"`"+` or must go
	// through an additional TLS termination point like a L7 Load Balancer.
	TrustedCertificateAuthorities *string `json:"trustedCertificateAuthorities,omitempty" bson:"trustedcertificateauthorities" mapstructure:"trustedCertificateAuthorities,omitempty"`

	// Type is the type of the service.
	Type *ServiceTypeValue `json:"type,omitempty" bson:"type" mapstructure:"type,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseService returns a new  SparseService.
func NewSparseService() *SparseService {
	return &SparseService{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseService) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseService) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseService) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseService) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseService) ToPlain() elemental.PlainIdentifiable {

	out := NewService()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.IPs != nil {
		out.IPs = *o.IPs
	}
	if o.JWTSigningCertificate != nil {
		out.JWTSigningCertificate = *o.JWTSigningCertificate
	}
	if o.MTLSCertificateAuthority != nil {
		out.MTLSCertificateAuthority = *o.MTLSCertificateAuthority
	}
	if o.OIDCCallbackURL != nil {
		out.OIDCCallbackURL = *o.OIDCCallbackURL
	}
	if o.OIDCClientID != nil {
		out.OIDCClientID = *o.OIDCClientID
	}
	if o.OIDCClientSecret != nil {
		out.OIDCClientSecret = *o.OIDCClientSecret
	}
	if o.OIDCProviderURL != nil {
		out.OIDCProviderURL = *o.OIDCProviderURL
	}
	if o.OIDCScopes != nil {
		out.OIDCScopes = *o.OIDCScopes
	}
	if o.TLSCertificate != nil {
		out.TLSCertificate = *o.TLSCertificate
	}
	if o.TLSCertificateKey != nil {
		out.TLSCertificateKey = *o.TLSCertificateKey
	}
	if o.TLSType != nil {
		out.TLSType = *o.TLSType
	}
	if o.AllAPITags != nil {
		out.AllAPITags = *o.AllAPITags
	}
	if o.AllServiceTags != nil {
		out.AllServiceTags = *o.AllServiceTags
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.AuthorizationType != nil {
		out.AuthorizationType = *o.AuthorizationType
	}
	if o.ClaimsToHTTPHeaderMappings != nil {
		out.ClaimsToHTTPHeaderMappings = *o.ClaimsToHTTPHeaderMappings
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Endpoints != nil {
		out.Endpoints = *o.Endpoints
	}
	if o.ExposedAPIs != nil {
		out.ExposedAPIs = *o.ExposedAPIs
	}
	if o.ExposedPort != nil {
		out.ExposedPort = *o.ExposedPort
	}
	if o.ExposedServiceIsTLS != nil {
		out.ExposedServiceIsTLS = *o.ExposedServiceIsTLS
	}
	if o.External != nil {
		out.External = *o.External
	}
	if o.Hosts != nil {
		out.Hosts = *o.Hosts
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Port != nil {
		out.Port = *o.Port
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.PublicApplicationPort != nil {
		out.PublicApplicationPort = *o.PublicApplicationPort
	}
	if o.RedirectURLOnAuthorizationFailure != nil {
		out.RedirectURLOnAuthorizationFailure = *o.RedirectURLOnAuthorizationFailure
	}
	if o.Selectors != nil {
		out.Selectors = *o.Selectors
	}
	if o.TrustedCertificateAuthorities != nil {
		out.TrustedCertificateAuthorities = *o.TrustedCertificateAuthorities
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseService) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseService) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetArchived returns the Archived of the receiver.
func (o *SparseService) GetArchived() bool {

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseService) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseService) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseService) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseService) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseService) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseService) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseService) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseService) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseService) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseService) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseService) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseService) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseService) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseService) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseService) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseService) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseService) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseService) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseService) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseService) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseService) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseService) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseService) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseService) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseService.
func (o *SparseService) DeepCopy() *SparseService {

	if o == nil {
		return nil
	}

	out := &SparseService{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseService.
func (o *SparseService) DeepCopyInto(out *SparseService) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseService: %s", err))
	}

	*out = *target.(*SparseService)
}
