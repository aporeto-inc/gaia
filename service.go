package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// ServiceAuthorizationTypeValue represents the possible values for attribute "authorizationType".
type ServiceAuthorizationTypeValue string

const (
	// ServiceAuthorizationTypeNone represents the value None.
	ServiceAuthorizationTypeNone ServiceAuthorizationTypeValue = "None"

	// ServiceAuthorizationTypeOIDC represents the value OIDC.
	ServiceAuthorizationTypeOIDC ServiceAuthorizationTypeValue = "OIDC"

	// ServiceAuthorizationTypePKI represents the value PKI.
	ServiceAuthorizationTypePKI ServiceAuthorizationTypeValue = "PKI"
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
	// The system will automatically resolve IP addresses from  host names otherwise.
	IPs types.IPList `json:"IPs" bson:"ips" mapstructure:"IPs,omitempty"`

	// JWTSigningCertificate is a certificate that can be used to validate user JWT in
	// HTTP requests. This is an optional field, needed only if user JWT validation is
	// required for this service. The certificate must be in PEM format.
	JWTSigningCertificate string `json:"JWTSigningCertificate" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

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

	// authorizationClaimMappings defines a list of mappings between incoming and
	// HTTP headers. When these mappings are defined, the enforcer will copy the
	// values of the claims to the corresponding HTTP headers.
	AuthorizationClaimMappings []*ClaimMapping `json:"authorizationClaimMappings" bson:"authorizationclaimmappings" mapstructure:"authorizationClaimMappings,omitempty"`

	// authorizationID is only valid for OIDC authorization and defines the
	// issuer ID of the OAUTH token.
	AuthorizationID string `json:"authorizationID" bson:"authorizationid" mapstructure:"authorizationID,omitempty"`

	// authorizationProvider is only valid for OAUTH authorization and defines the
	// URL to the OAUTH provider that must be used.
	AuthorizationProvider string `json:"authorizationProvider" bson:"authorizationprovider" mapstructure:"authorizationProvider,omitempty"`

	// authorizationSecret is only valid for OIDC authorization and defines the
	// secret that should be used with the OAUTH provider to validate tokens.
	AuthorizationSecret string `json:"authorizationSecret" bson:"authorizationsecret" mapstructure:"authorizationSecret,omitempty"`

	// AuthorizationType defines the user authorization type that should be used.
	// Currently supporting PKI, and OIDC.
	AuthorizationType ServiceAuthorizationTypeValue `json:"authorizationType" bson:"authorizationtype" mapstructure:"authorizationType,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Endpoints is a read only attribute that actually resolves the API
	// endpoints that the service is exposing. Only valid during policy rendering.
	Endpoints types.ExposedAPIList `json:"endpoints" bson:"-" mapstructure:"endpoints,omitempty"`

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

	// RedirectOnFail is a boolean that forces a redirect response if an API request
	// arrives and the user authorization information is not valid. This only applies
	// to HTTP services and it is only send for APIs that are not public.
	RedirectOnFail bool `json:"redirectOnFail" bson:"redirectonfail" mapstructure:"redirectOnFail,omitempty"`

	// RedirectOnNoToken is a boolean that forces a redirect response if an API request
	// arrives and there is no user authorization information. This only applies to
	// HTTP services and it is only send for APIs that are not public.
	RedirectOnNoToken bool `json:"redirectOnNoToken" bson:"redirectonnotoken" mapstructure:"redirectOnNoToken,omitempty"`

	// RedirectURL is the URL that will be send back to the user to
	// redirect for authentication if there is no user authorization information in
	// the API request. URL can be defined if a redirection is requested only.
	RedirectURL string `json:"redirectURL" bson:"redirecturl" mapstructure:"redirectURL,omitempty"`

	// Selectors contains the tag expression that an a processing unit
	// must match in order to implement this particular service.
	Selectors [][]string `json:"selectors" bson:"selectors" mapstructure:"selectors,omitempty"`

	// ServiceCA  is the certificate authority that the service is using. This
	// is needed for external services with private certificate authorities. The
	// field is optional. If provided, this must be a valid PEM CA file.
	ServiceCA string `json:"serviceCA" bson:"serviceca" mapstructure:"serviceCA,omitempty"`

	// Type is the type of the service.
	Type ServiceTypeValue `json:"type" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion:               1,
		AssociatedTags:             []string{},
		AllAPITags:                 []string{},
		AuthorizationClaimMappings: []*ClaimMapping{},
		AllServiceTags:             []string{},
		Annotations:                map[string][]string{},
		AuthorizationType:          ServiceAuthorizationTypeNone,
		Endpoints:                  types.ExposedAPIList{},
		External:                   false,
		RedirectOnFail:             false,
		Metadata:                   []string{},
		NormalizedTags:             []string{},
		IPs:                        types.IPList{},
		RedirectOnNoToken:          false,
		Type:                       ServiceTypeHTTP,
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

// SetAnnotations sets the given Annotations of the receiver.
func (o *Service) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetArchived returns the Archived of the receiver.
func (o *Service) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the given Archived of the receiver.
func (o *Service) SetArchived(archived bool) {

	o.Archived = archived
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Service) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Service) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Service) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Service) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *Service) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *Service) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Service) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *Service) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Service) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Service) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Service) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Service) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Service) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Service) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Service) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Service) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseService{
			ID:                         &o.ID,
			IPs:                        &o.IPs,
			JWTSigningCertificate:      &o.JWTSigningCertificate,
			AllAPITags:                 &o.AllAPITags,
			AllServiceTags:             &o.AllServiceTags,
			Annotations:                &o.Annotations,
			Archived:                   &o.Archived,
			AssociatedTags:             &o.AssociatedTags,
			AuthorizationClaimMappings: &o.AuthorizationClaimMappings,
			AuthorizationID:            &o.AuthorizationID,
			AuthorizationProvider:      &o.AuthorizationProvider,
			AuthorizationSecret:        &o.AuthorizationSecret,
			AuthorizationType:          &o.AuthorizationType,
			CreateTime:                 &o.CreateTime,
			Description:                &o.Description,
			Endpoints:                  &o.Endpoints,
			ExposedAPIs:                &o.ExposedAPIs,
			ExposedPort:                &o.ExposedPort,
			External:                   &o.External,
			Hosts:                      &o.Hosts,
			Metadata:                   &o.Metadata,
			Name:                       &o.Name,
			Namespace:                  &o.Namespace,
			NormalizedTags:             &o.NormalizedTags,
			Port:                       &o.Port,
			Protected:                  &o.Protected,
			PublicApplicationPort:      &o.PublicApplicationPort,
			RedirectOnFail:             &o.RedirectOnFail,
			RedirectOnNoToken:          &o.RedirectOnNoToken,
			RedirectURL:                &o.RedirectURL,
			Selectors:                  &o.Selectors,
			ServiceCA:                  &o.ServiceCA,
			Type:                       &o.Type,
			UpdateTime:                 &o.UpdateTime,
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
		case "authorizationClaimMappings":
			sp.AuthorizationClaimMappings = &(o.AuthorizationClaimMappings)
		case "authorizationID":
			sp.AuthorizationID = &(o.AuthorizationID)
		case "authorizationProvider":
			sp.AuthorizationProvider = &(o.AuthorizationProvider)
		case "authorizationSecret":
			sp.AuthorizationSecret = &(o.AuthorizationSecret)
		case "authorizationType":
			sp.AuthorizationType = &(o.AuthorizationType)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "endpoints":
			sp.Endpoints = &(o.Endpoints)
		case "exposedAPIs":
			sp.ExposedAPIs = &(o.ExposedAPIs)
		case "exposedPort":
			sp.ExposedPort = &(o.ExposedPort)
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
		case "redirectOnFail":
			sp.RedirectOnFail = &(o.RedirectOnFail)
		case "redirectOnNoToken":
			sp.RedirectOnNoToken = &(o.RedirectOnNoToken)
		case "redirectURL":
			sp.RedirectURL = &(o.RedirectURL)
		case "selectors":
			sp.Selectors = &(o.Selectors)
		case "serviceCA":
			sp.ServiceCA = &(o.ServiceCA)
		case "type":
			sp.Type = &(o.Type)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
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
	if so.AuthorizationClaimMappings != nil {
		o.AuthorizationClaimMappings = *so.AuthorizationClaimMappings
	}
	if so.AuthorizationID != nil {
		o.AuthorizationID = *so.AuthorizationID
	}
	if so.AuthorizationProvider != nil {
		o.AuthorizationProvider = *so.AuthorizationProvider
	}
	if so.AuthorizationSecret != nil {
		o.AuthorizationSecret = *so.AuthorizationSecret
	}
	if so.AuthorizationType != nil {
		o.AuthorizationType = *so.AuthorizationType
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
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
	if so.RedirectOnFail != nil {
		o.RedirectOnFail = *so.RedirectOnFail
	}
	if so.RedirectOnNoToken != nil {
		o.RedirectOnNoToken = *so.RedirectOnNoToken
	}
	if so.RedirectURL != nil {
		o.RedirectURL = *so.RedirectURL
	}
	if so.Selectors != nil {
		o.Selectors = *so.Selectors
	}
	if so.ServiceCA != nil {
		o.ServiceCA = *so.ServiceCA
	}
	if so.Type != nil {
		o.Type = *so.Type
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.AuthorizationClaimMappings {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateStringInList("authorizationType", string(o.AuthorizationType), []string{"PKI", "OIDC", "None"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
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
		PrimaryKey:     true,
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
The system will automatically resolve IP addresses from  host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"JWTSigningCertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"AllAPITags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"AllServiceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
		SubType:        "annotations",
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
		SubType:        "tags_list",
		Type:           "external",
	},
	"AuthorizationClaimMappings": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationClaimMappings",
		Description: `authorizationClaimMappings defines a list of mappings between incoming and
HTTP headers. When these mappings are defined, the enforcer will copy the
values of the claims to the corresponding HTTP headers.`,
		Exposed: true,
		Name:    "authorizationClaimMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
	},
	"AuthorizationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationID",
		Description: `authorizationID is only valid for OIDC authorization and defines the
issuer ID of the OAUTH token.`,
		Exposed: true,
		Name:    "authorizationID",
		Stored:  true,
		Type:    "string",
	},
	"AuthorizationProvider": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationProvider",
		Description: `authorizationProvider is only valid for OAUTH authorization and defines the
URL to the OAUTH provider that must be used.`,
		Exposed: true,
		Name:    "authorizationProvider",
		Stored:  true,
		Type:    "string",
	},
	"AuthorizationSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationSecret",
		Description: `authorizationSecret is only valid for OIDC authorization and defines the
secret that should be used with the OAUTH provider to validate tokens.`,
		Exposed: true,
		Name:    "authorizationSecret",
		Stored:  true,
		Type:    "string",
	},
	"AuthorizationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"PKI", "OIDC", "None"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `AuthorizationType defines the user authorization type that should be used.
Currently supporting PKI, and OIDC.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
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
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "exposed_api_list",
		Type:     "external",
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
		SubType: "policies_list",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		CreationOnly:   true,
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
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
	"RedirectOnFail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectOnFail",
		Description: `RedirectOnFail is a boolean that forces a redirect response if an API request
arrives and the user authorization information is not valid. This only applies
to HTTP services and it is only send for APIs that are not public.`,
		Exposed:   true,
		Name:      "redirectOnFail",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"RedirectOnNoToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectOnNoToken",
		Description: `RedirectOnNoToken is a boolean that forces a redirect response if an API request
arrives and there is no user authorization information. This only applies to
HTTP services and it is only send for APIs that are not public.`,
		Exposed:   true,
		Name:      "redirectOnNoToken",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"RedirectURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURL",
		Description: `RedirectURL is the URL that will be send back to the user to
redirect for authentication if there is no user authorization information in
the API request. URL can be defined if a redirection is requested only.`,
		Exposed: true,
		Name:    "redirectURL",
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
		SubType: "policies_list",
		Type:    "external",
	},
	"ServiceCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCA",
		Description: `ServiceCA  is the certificate authority that the service is using. This
is needed for external services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.`,
		Exposed: true,
		Name:    "serviceCA",
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
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		PrimaryKey:     true,
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
The system will automatically resolve IP addresses from  host names otherwise.`,
		Exposed: true,
		Name:    "IPs",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"jwtsigningcertificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "JWTSigningCertificate",
		Description: `JWTSigningCertificate is a certificate that can be used to validate user JWT in
HTTP requests. This is an optional field, needed only if user JWT validation is
required for this service. The certificate must be in PEM format.`,
		Exposed: true,
		Name:    "JWTSigningCertificate",
		Stored:  true,
		Type:    "string",
	},
	"allapitags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllAPITags",
		Description:    `This is a set of all API tags for matching in the DB.`,
		Name:           "allAPITags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"allservicetags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllServiceTags",
		Description:    `This is a set of all selector tags for matching in the DB.`,
		Name:           "allServiceTags",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
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
		SubType:        "annotations",
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
		SubType:        "tags_list",
		Type:           "external",
	},
	"authorizationclaimmappings": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationClaimMappings",
		Description: `authorizationClaimMappings defines a list of mappings between incoming and
HTTP headers. When these mappings are defined, the enforcer will copy the
values of the claims to the corresponding HTTP headers.`,
		Exposed: true,
		Name:    "authorizationClaimMappings",
		Stored:  true,
		SubType: "claimmapping",
		Type:    "refList",
	},
	"authorizationid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationID",
		Description: `authorizationID is only valid for OIDC authorization and defines the
issuer ID of the OAUTH token.`,
		Exposed: true,
		Name:    "authorizationID",
		Stored:  true,
		Type:    "string",
	},
	"authorizationprovider": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationProvider",
		Description: `authorizationProvider is only valid for OAUTH authorization and defines the
URL to the OAUTH provider that must be used.`,
		Exposed: true,
		Name:    "authorizationProvider",
		Stored:  true,
		Type:    "string",
	},
	"authorizationsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuthorizationSecret",
		Description: `authorizationSecret is only valid for OIDC authorization and defines the
secret that should be used with the OAUTH provider to validate tokens.`,
		Exposed: true,
		Name:    "authorizationSecret",
		Stored:  true,
		Type:    "string",
	},
	"authorizationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"PKI", "OIDC", "None"},
		ConvertedName:  "AuthorizationType",
		DefaultValue:   ServiceAuthorizationTypeNone,
		Description: `AuthorizationType defines the user authorization type that should be used.
Currently supporting PKI, and OIDC.`,
		Exposed: true,
		Name:    "authorizationType",
		Stored:  true,
		Type:    "enum",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
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
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"endpoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Endpoints",
		Description: `Endpoints is a read only attribute that actually resolves the API
endpoints that the service is exposing. Only valid during policy rendering.`,
		Exposed:  true,
		Name:     "endpoints",
		ReadOnly: true,
		SubType:  "exposed_api_list",
		Type:     "external",
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
		SubType: "policies_list",
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
		SubType:    "metadata_list",
		Type:       "external",
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
		CreationOnly:   true,
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
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
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
	"redirectonfail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectOnFail",
		Description: `RedirectOnFail is a boolean that forces a redirect response if an API request
arrives and the user authorization information is not valid. This only applies
to HTTP services and it is only send for APIs that are not public.`,
		Exposed:   true,
		Name:      "redirectOnFail",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"redirectonnotoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectOnNoToken",
		Description: `RedirectOnNoToken is a boolean that forces a redirect response if an API request
arrives and there is no user authorization information. This only applies to
HTTP services and it is only send for APIs that are not public.`,
		Exposed:   true,
		Name:      "redirectOnNoToken",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
	},
	"redirecturl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RedirectURL",
		Description: `RedirectURL is the URL that will be send back to the user to
redirect for authentication if there is no user authorization information in
the API request. URL can be defined if a redirection is requested only.`,
		Exposed: true,
		Name:    "redirectURL",
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
		SubType: "policies_list",
		Type:    "external",
	},
	"serviceca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCA",
		Description: `ServiceCA  is the certificate authority that the service is using. This
is needed for external services with private certificate authorities. The
field is optional. If provided, this must be a valid PEM CA file.`,
		Exposed: true,
		Name:    "serviceCA",
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
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
	// The system will automatically resolve IP addresses from  host names otherwise.
	IPs *types.IPList `json:"IPs,omitempty" bson:"ips" mapstructure:"IPs,omitempty"`

	// JWTSigningCertificate is a certificate that can be used to validate user JWT in
	// HTTP requests. This is an optional field, needed only if user JWT validation is
	// required for this service. The certificate must be in PEM format.
	JWTSigningCertificate *string `json:"JWTSigningCertificate,omitempty" bson:"jwtsigningcertificate" mapstructure:"JWTSigningCertificate,omitempty"`

	// This is a set of all API tags for matching in the DB.
	AllAPITags *[]string `json:"-,omitempty" bson:"allapitags" mapstructure:"-,omitempty"`

	// This is a set of all selector tags for matching in the DB.
	AllServiceTags *[]string `json:"-,omitempty" bson:"allservicetags" mapstructure:"-,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// Archived defines if the object is archived.
	Archived *bool `json:"-,omitempty" bson:"archived" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// authorizationClaimMappings defines a list of mappings between incoming and
	// HTTP headers. When these mappings are defined, the enforcer will copy the
	// values of the claims to the corresponding HTTP headers.
	AuthorizationClaimMappings *[]*ClaimMapping `json:"authorizationClaimMappings,omitempty" bson:"authorizationclaimmappings" mapstructure:"authorizationClaimMappings,omitempty"`

	// authorizationID is only valid for OIDC authorization and defines the
	// issuer ID of the OAUTH token.
	AuthorizationID *string `json:"authorizationID,omitempty" bson:"authorizationid" mapstructure:"authorizationID,omitempty"`

	// authorizationProvider is only valid for OAUTH authorization and defines the
	// URL to the OAUTH provider that must be used.
	AuthorizationProvider *string `json:"authorizationProvider,omitempty" bson:"authorizationprovider" mapstructure:"authorizationProvider,omitempty"`

	// authorizationSecret is only valid for OIDC authorization and defines the
	// secret that should be used with the OAUTH provider to validate tokens.
	AuthorizationSecret *string `json:"authorizationSecret,omitempty" bson:"authorizationsecret" mapstructure:"authorizationSecret,omitempty"`

	// AuthorizationType defines the user authorization type that should be used.
	// Currently supporting PKI, and OIDC.
	AuthorizationType *ServiceAuthorizationTypeValue `json:"authorizationType,omitempty" bson:"authorizationtype" mapstructure:"authorizationType,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Endpoints is a read only attribute that actually resolves the API
	// endpoints that the service is exposing. Only valid during policy rendering.
	Endpoints *types.ExposedAPIList `json:"endpoints,omitempty" bson:"-" mapstructure:"endpoints,omitempty"`

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

	// RedirectOnFail is a boolean that forces a redirect response if an API request
	// arrives and the user authorization information is not valid. This only applies
	// to HTTP services and it is only send for APIs that are not public.
	RedirectOnFail *bool `json:"redirectOnFail,omitempty" bson:"redirectonfail" mapstructure:"redirectOnFail,omitempty"`

	// RedirectOnNoToken is a boolean that forces a redirect response if an API request
	// arrives and there is no user authorization information. This only applies to
	// HTTP services and it is only send for APIs that are not public.
	RedirectOnNoToken *bool `json:"redirectOnNoToken,omitempty" bson:"redirectonnotoken" mapstructure:"redirectOnNoToken,omitempty"`

	// RedirectURL is the URL that will be send back to the user to
	// redirect for authentication if there is no user authorization information in
	// the API request. URL can be defined if a redirection is requested only.
	RedirectURL *string `json:"redirectURL,omitempty" bson:"redirecturl" mapstructure:"redirectURL,omitempty"`

	// Selectors contains the tag expression that an a processing unit
	// must match in order to implement this particular service.
	Selectors *[][]string `json:"selectors,omitempty" bson:"selectors" mapstructure:"selectors,omitempty"`

	// ServiceCA  is the certificate authority that the service is using. This
	// is needed for external services with private certificate authorities. The
	// field is optional. If provided, this must be a valid PEM CA file.
	ServiceCA *string `json:"serviceCA,omitempty" bson:"serviceca" mapstructure:"serviceCA,omitempty"`

	// Type is the type of the service.
	Type *ServiceTypeValue `json:"type,omitempty" bson:"type" mapstructure:"type,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
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
	if o.AuthorizationClaimMappings != nil {
		out.AuthorizationClaimMappings = *o.AuthorizationClaimMappings
	}
	if o.AuthorizationID != nil {
		out.AuthorizationID = *o.AuthorizationID
	}
	if o.AuthorizationProvider != nil {
		out.AuthorizationProvider = *o.AuthorizationProvider
	}
	if o.AuthorizationSecret != nil {
		out.AuthorizationSecret = *o.AuthorizationSecret
	}
	if o.AuthorizationType != nil {
		out.AuthorizationType = *o.AuthorizationType
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
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
	if o.RedirectOnFail != nil {
		out.RedirectOnFail = *o.RedirectOnFail
	}
	if o.RedirectOnNoToken != nil {
		out.RedirectOnNoToken = *o.RedirectOnNoToken
	}
	if o.RedirectURL != nil {
		out.RedirectURL = *o.RedirectURL
	}
	if o.Selectors != nil {
		out.Selectors = *o.Selectors
	}
	if o.ServiceCA != nil {
		out.ServiceCA = *o.ServiceCA
	}
	if o.Type != nil {
		out.Type = *o.Type
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}
