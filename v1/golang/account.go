package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// AccountLDAPConnSecurityProtocolValue represents the possible values for attribute "LDAPConnSecurityProtocol".
type AccountLDAPConnSecurityProtocolValue string

const (
	// AccountLDAPConnSecurityProtocolInbandTLS represents the value InbandTLS.
	AccountLDAPConnSecurityProtocolInbandTLS AccountLDAPConnSecurityProtocolValue = "InbandTLS"

	// AccountLDAPConnSecurityProtocolTLS represents the value TLS.
	AccountLDAPConnSecurityProtocolTLS AccountLDAPConnSecurityProtocolValue = "TLS"
)

// AccountStatusValue represents the possible values for attribute "status".
type AccountStatusValue string

const (
	// AccountStatusActive represents the value Active.
	AccountStatusActive AccountStatusValue = "Active"

	// AccountStatusDisabled represents the value Disabled.
	AccountStatusDisabled AccountStatusValue = "Disabled"

	// AccountStatusInvited represents the value Invited.
	AccountStatusInvited AccountStatusValue = "Invited"

	// AccountStatusPending represents the value Pending.
	AccountStatusPending AccountStatusValue = "Pending"
)

// AccountIdentity represents the Identity of the object.
var AccountIdentity = elemental.Identity{
	Name:     "account",
	Category: "accounts",
	Private:  false,
}

// AccountsList represents a list of Accounts
type AccountsList []*Account

// ContentIdentity returns the identity of the objects in the list.
func (o AccountsList) ContentIdentity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the AccountsList.
func (o AccountsList) Copy() elemental.ContentIdentifiable {

	copy := append(AccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccountsList.
func (o AccountsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Account))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AccountsList) Version() int {

	return 1
}

// Account represents the model of a account
type Account struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// LDAPAddress holds the account authentication account's private ldap server.
	LDAPAddress string `json:"LDAPAddress" bson:"ldapaddress" mapstructure:"LDAPAddress,omitempty"`

	// LDAPBaseDN holds the base DN to use to ldap queries.
	LDAPBaseDN string `json:"LDAPBaseDN" bson:"ldapbasedn" mapstructure:"LDAPBaseDN,omitempty"`

	// LDAPBindDN holds the account's internal LDAP bind dn for querying auth.
	LDAPBindDN string `json:"LDAPBindDN" bson:"ldapbinddn" mapstructure:"LDAPBindDN,omitempty"`

	// LDAPBindPassword holds the password to the LDAPBindDN.
	LDAPBindPassword string `json:"LDAPBindPassword" bson:"ldapbindpassword" mapstructure:"LDAPBindPassword,omitempty"`

	// LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
	// Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
	// other systems, value may be 'uid={USERNAME}'.
	LDAPBindSearchFilter string `json:"LDAPBindSearchFilter" bson:"ldapbindsearchfilter" mapstructure:"LDAPBindSearchFilter,omitempty"`

	// LDAPCertificateAuthority contains the optional certificate author ity that will
	// be used to connect to the LDAP server. It is not needed if the TLS certificate
	// of the LDAP is issued from a public truster CA.
	LDAPCertificateAuthority string `json:"LDAPCertificateAuthority" bson:"ldapcertificateauthority" mapstructure:"LDAPCertificateAuthority,omitempty"`

	// LDAPConnProtocol holds the connection type for the LDAP provider.
	LDAPConnSecurityProtocol AccountLDAPConnSecurityProtocolValue `json:"LDAPConnSecurityProtocol" bson:"ldapconnsecurityprotocol" mapstructure:"LDAPConnSecurityProtocol,omitempty"`

	// LDAPEnabled triggers if the account uses it's own LDAP for authentication.
	LDAPEnabled bool `json:"LDAPEnabled" bson:"ldapenabled" mapstructure:"LDAPEnabled,omitempty"`

	// LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
	// authorization system.
	LDAPIgnoredKeys []string `json:"LDAPIgnoredKeys" bson:"ldapignoredkeys" mapstructure:"LDAPIgnoredKeys,omitempty"`

	// LDAPSubjectKey holds key to be used to populate the subject. If you want to
	// use the user as a subject, for Windows based systems you may use
	// 'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
	// also use any alternate key.
	LDAPSubjectKey string `json:"LDAPSubjectKey" bson:"ldapsubjectkey" mapstructure:"LDAPSubjectKey,omitempty"`

	// Set to enable or disable two factor authentication.
	OTPEnabled bool `json:"OTPEnabled" bson:"otpenabled" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64 encoded QRCode for setting up 2 factor auth.
	OTPQRCode string `json:"OTPQRCode" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the 2 factor secret.
	OTPSecret string `json:"-" bson:"otpsecret" mapstructure:"-,omitempty"`

	// AccessEnabled defines if the account holder should have access to the systems.
	AccessEnabled bool `json:"accessEnabled" bson:"accessenabled" mapstructure:"accessEnabled,omitempty"`

	// ActivationExpiration contains the expiration date of the activation token.
	ActivationExpiration time.Time `json:"-" bson:"activationexpiration" mapstructure:"-,omitempty"`

	// ActivationToken contains the activation token.
	ActivationToken string `json:"activationToken,omitempty" bson:"activationtoken" mapstructure:"activationToken,omitempty,omitempty"`

	// AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.
	AssociatedAPIAuthPolicyID string `json:"-" bson:"associatedapiauthpolicyid" mapstructure:"-,omitempty"`

	// AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.
	AssociatedAWSPolicies map[string]string `json:"-" bson:"associatedawspolicies" mapstructure:"-,omitempty"`

	// AssociatedNamespaceID contains the ID of the associated namespace.
	AssociatedNamespaceID string `json:"-" bson:"associatednamespaceid" mapstructure:"-,omitempty"`

	// AssociatedPlanKey contains the plan key that is associated to this account.
	AssociatedPlanKey string `json:"associatedPlanKey" bson:"associatedplankey" mapstructure:"associatedPlanKey,omitempty"`

	// AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.
	AssociatedQuotaPolicies map[string]string `json:"-" bson:"associatedquotapolicies" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company string `json:"company" bson:"company" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email string `json:"email" bson:"email" mapstructure:"email,omitempty"`

	// First Name of the account user.
	FirstName string `json:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// Last Name of the account user.
	LastName string `json:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// Name of the account.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Password for the account.
	Password string `json:"password" bson:"password" mapstructure:"password,omitempty"`

	// ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.
	ReCAPTCHAKey string `json:"reCAPTCHAKey" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// ResetPasswordExpiration contains the expiration time for reseting the password.
	ResetPasswordExpiration time.Time `json:"-" bson:"resetpasswordexpiration" mapstructure:"-,omitempty"`

	// ResetPasswordToken contains the token to use for resetting password.
	ResetPasswordToken string `json:"-" bson:"resetpasswordtoken" mapstructure:"-,omitempty"`

	// Status of the account.
	Status AccountStatusValue `json:"status" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		ModelVersion:             1,
		AssociatedAWSPolicies:    map[string]string{},
		AssociatedPlanKey:        "aporeto.plan.free",
		AssociatedQuotaPolicies:  map[string]string{},
		LDAPBindSearchFilter:     "uid={USERNAME}",
		LDAPConnSecurityProtocol: "InbandTLS",
		LDAPIgnoredKeys:          []string{},
		LDAPSubjectKey:           "uid",
		Status:                   "Pending",
	}
}

// Identity returns the Identity of the object.
func (o *Account) Identity() elemental.Identity {

	return AccountIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Account) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Account) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Account) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Account) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Account) Doc() string {
	return `This api allows to view and manage basic information about your account like
your name, password, enable 2 factor authentication.`
}

func (o *Account) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Account) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("LDAPConnSecurityProtocol", string(o.LDAPConnSecurityProtocol), []string{"TLS", "InbandTLS"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Disabled", "Invited", "Pending"}, true); err != nil {
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
func (*Account) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccountAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccountLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Account) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountAttributesMap
}

// AccountAttributesMap represents the map of attribute for Account.
var AccountAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPAddress",
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBaseDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBaseDN",
		Description:    `LDAPBaseDN holds the base DN to use to ldap queries.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindDN",
		Description:    `LDAPBindDN holds the account's internal LDAP bind dn for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindPassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindPassword",
		Description:    `LDAPBindPassword holds the password to the LDAPBindDN.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LDAPBindSearchFilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
other systems, value may be 'uid={USERNAME}'.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPBindSearchFilter",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"LDAPCertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPCertificateAuthority",
		Description: `LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPCertificateAuthority",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"LDAPConnSecurityProtocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "LDAPConnSecurityProtocol",
		DefaultValue:   AccountLDAPConnSecurityProtocolInbandTLS,
		Description:    `LDAPConnProtocol holds the connection type for the LDAP provider.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPConnSecurityProtocol",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"LDAPEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPEnabled",
		Description:    `LDAPEnabled triggers if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"LDAPIgnoredKeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPIgnoredKeys",
		Description: `LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPIgnoredKeys",
		Orderable:  true,
		Stored:     true,
		SubType:    "ignore_list",
		Type:       "external",
	},
	"LDAPSubjectKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPSubjectKey",
		DefaultValue:   "uid",
		Description: `LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPSubjectKey",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"OTPEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Set to enable or disable two factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"OTPQRCode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64 encoded QRCode for setting up 2 factor auth.`,
		Exposed:        true,
		Format:         "free",
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"OTPSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the 2 factor secret.`,
		Format:         "free",
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"AccessEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ActivationExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationExpiration",
		Description:    `ActivationExpiration contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ActivationToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `ActivationToken contains the activation token.`,
		Exposed:        true,
		Format:         "free",
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
		Format:         "free",
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAWSPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"AssociatedNamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
		Format:         "free",
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedPlanKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AssociatedPlanKey",
		DefaultValue:   "aporeto.plan.free",
		Description:    `AssociatedPlanKey contains the plan key that is associated to this account.`,
		Exposed:        true,
		Format:         "free",
		Name:           "associatedPlanKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedQuotaPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"Company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"FirstName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `First Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LastName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `Last Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password for the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ReCAPTCHAKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReCAPTCHAKey",
		CreationOnly:   true,
		Description:    `ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.`,
		Exposed:        true,
		Format:         "free",
		Name:           "reCAPTCHAKey",
		Type:           "string",
	},
	"ResetPasswordExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `ResetPasswordExpiration contains the expiration time for reseting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ResetPasswordToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `ResetPasswordToken contains the token to use for resetting password.`,
		Format:         "free",
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		DefaultValue:   AccountStatusPending,
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// AccountLowerCaseAttributesMap represents the map of attribute for Account.
var AccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ldapaddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPAddress",
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbasedn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBaseDN",
		Description:    `LDAPBaseDN holds the base DN to use to ldap queries.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBaseDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbinddn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindDN",
		Description:    `LDAPBindDN holds the account's internal LDAP bind dn for querying auth.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindDN",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbindpassword": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindPassword",
		Description:    `LDAPBindPassword holds the password to the LDAPBindDN.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "LDAPBindPassword",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ldapbindsearchfilter": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPBindSearchFilter",
		DefaultValue:   "uid={USERNAME}",
		Description: `LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
other systems, value may be 'uid={USERNAME}'.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPBindSearchFilter",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"ldapcertificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPCertificateAuthority",
		Description: `LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPCertificateAuthority",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"ldapconnsecurityprotocol": elemental.AttributeSpecification{
		AllowedChoices: []string{"TLS", "InbandTLS"},
		ConvertedName:  "LDAPConnSecurityProtocol",
		DefaultValue:   AccountLDAPConnSecurityProtocolInbandTLS,
		Description:    `LDAPConnProtocol holds the connection type for the LDAP provider.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPConnSecurityProtocol",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"ldapenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPEnabled",
		Description:    `LDAPEnabled triggers if the account uses it's own LDAP for authentication.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "LDAPEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ldapignoredkeys": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPIgnoredKeys",
		Description: `LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
authorization system.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPIgnoredKeys",
		Orderable:  true,
		Stored:     true,
		SubType:    "ignore_list",
		Type:       "external",
	},
	"ldapsubjectkey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPSubjectKey",
		DefaultValue:   "uid",
		Description: `LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "LDAPSubjectKey",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"otpenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Set to enable or disable two factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"otpqrcode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64 encoded QRCode for setting up 2 factor auth.`,
		Exposed:        true,
		Format:         "free",
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"otpsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the 2 factor secret.`,
		Format:         "free",
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"accessenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "accessEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"activationexpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationExpiration",
		Description:    `ActivationExpiration contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"activationtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `ActivationToken contains the activation token.`,
		Exposed:        true,
		Format:         "free",
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"associatedapiauthpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
		Format:         "free",
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"associatedawspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"associatednamespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
		Format:         "free",
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"associatedplankey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AssociatedPlanKey",
		DefaultValue:   "aporeto.plan.free",
		Description:    `AssociatedPlanKey contains the plan key that is associated to this account.`,
		Exposed:        true,
		Format:         "free",
		Name:           "associatedPlanKey",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedquotapolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "associated_policies",
		Type:           "external",
	},
	"company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "company",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description:    `Email of the account holder.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "email",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"firstname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FirstName",
		Description:    `First Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "firstName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"lastname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastName",
		Description:    `Last Name of the account user.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "lastName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name of the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"password": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Password",
		Description:    `Password for the account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "password",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"recaptchakey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReCAPTCHAKey",
		CreationOnly:   true,
		Description:    `ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.`,
		Exposed:        true,
		Format:         "free",
		Name:           "reCAPTCHAKey",
		Type:           "string",
	},
	"resetpasswordexpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `ResetPasswordExpiration contains the expiration time for reseting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"resetpasswordtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `ResetPasswordToken contains the token to use for resetting password.`,
		Format:         "free",
		Name:           "resetPasswordToken",
		Stored:         true,
		Type:           "string",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Active", "Disabled", "Invited", "Pending"},
		Autogenerated:  true,
		ConvertedName:  "Status",
		DefaultValue:   AccountStatusPending,
		Description:    `Status of the account.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "enum",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
