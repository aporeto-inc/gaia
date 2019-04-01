package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
	Package:  "vince",
	Private:  false,
}

// AccountsList represents a list of Accounts
type AccountsList []*Account

// Identity returns the identity of the objects in the list.
func (o AccountsList) Identity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the AccountsList.
func (o AccountsList) Copy() elemental.Identifiables {

	copy := append(AccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccountsList.
func (o AccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Account))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccountsList converted to SparseAccountsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccountsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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
	ActivationToken string `json:"activationToken,omitempty" bson:"activationtoken" mapstructure:"activationToken,omitempty"`

	// AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.
	AssociatedAPIAuthPolicyID string `json:"-" bson:"associatedapiauthpolicyid" mapstructure:"-,omitempty"`

	// AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.
	AssociatedAWSPolicies map[string]string `json:"-" bson:"associatedawspolicies" mapstructure:"-,omitempty"`

	// associatedBillingID holds the ID of the associated billing customer.
	AssociatedBillingID string `json:"associatedBillingID" bson:"associatedbillingid" mapstructure:"associatedBillingID,omitempty"`

	// associatedGCPPolicies contains a map of associated GCP Enforcerd Policies.
	AssociatedGCPPolicies map[string]string `json:"-" bson:"associatedgcppolicies" mapstructure:"-,omitempty"`

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

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		ModelVersion:             1,
		AssociatedQuotaPolicies:  map[string]string{},
		AssociatedPlanKey:        "aporeto.plan.free",
		AssociatedGCPPolicies:    map[string]string{},
		AssociatedAWSPolicies:    map[string]string{},
		LDAPSubjectKey:           "uid",
		LDAPIgnoredKeys:          []string{},
		LDAPConnSecurityProtocol: AccountLDAPConnSecurityProtocolInbandTLS,
		LDAPBindSearchFilter:     "uid={USERNAME}",
		Status:                   AccountStatusPending,
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

// GetZHash returns the ZHash of the receiver.
func (o *Account) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Account) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Account) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Account) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Account) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccount{
			ID:                        &o.ID,
			LDAPAddress:               &o.LDAPAddress,
			LDAPBaseDN:                &o.LDAPBaseDN,
			LDAPBindDN:                &o.LDAPBindDN,
			LDAPBindPassword:          &o.LDAPBindPassword,
			LDAPBindSearchFilter:      &o.LDAPBindSearchFilter,
			LDAPCertificateAuthority:  &o.LDAPCertificateAuthority,
			LDAPConnSecurityProtocol:  &o.LDAPConnSecurityProtocol,
			LDAPEnabled:               &o.LDAPEnabled,
			LDAPIgnoredKeys:           &o.LDAPIgnoredKeys,
			LDAPSubjectKey:            &o.LDAPSubjectKey,
			OTPEnabled:                &o.OTPEnabled,
			OTPQRCode:                 &o.OTPQRCode,
			OTPSecret:                 &o.OTPSecret,
			AccessEnabled:             &o.AccessEnabled,
			ActivationExpiration:      &o.ActivationExpiration,
			ActivationToken:           &o.ActivationToken,
			AssociatedAPIAuthPolicyID: &o.AssociatedAPIAuthPolicyID,
			AssociatedAWSPolicies:     &o.AssociatedAWSPolicies,
			AssociatedBillingID:       &o.AssociatedBillingID,
			AssociatedGCPPolicies:     &o.AssociatedGCPPolicies,
			AssociatedNamespaceID:     &o.AssociatedNamespaceID,
			AssociatedPlanKey:         &o.AssociatedPlanKey,
			AssociatedQuotaPolicies:   &o.AssociatedQuotaPolicies,
			Company:                   &o.Company,
			CreateTime:                &o.CreateTime,
			Email:                     &o.Email,
			FirstName:                 &o.FirstName,
			LastName:                  &o.LastName,
			Name:                      &o.Name,
			Password:                  &o.Password,
			ReCAPTCHAKey:              &o.ReCAPTCHAKey,
			ResetPasswordExpiration:   &o.ResetPasswordExpiration,
			ResetPasswordToken:        &o.ResetPasswordToken,
			Status:                    &o.Status,
			UpdateTime:                &o.UpdateTime,
			ZHash:                     &o.ZHash,
			Zone:                      &o.Zone,
		}
	}

	sp := &SparseAccount{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "LDAPAddress":
			sp.LDAPAddress = &(o.LDAPAddress)
		case "LDAPBaseDN":
			sp.LDAPBaseDN = &(o.LDAPBaseDN)
		case "LDAPBindDN":
			sp.LDAPBindDN = &(o.LDAPBindDN)
		case "LDAPBindPassword":
			sp.LDAPBindPassword = &(o.LDAPBindPassword)
		case "LDAPBindSearchFilter":
			sp.LDAPBindSearchFilter = &(o.LDAPBindSearchFilter)
		case "LDAPCertificateAuthority":
			sp.LDAPCertificateAuthority = &(o.LDAPCertificateAuthority)
		case "LDAPConnSecurityProtocol":
			sp.LDAPConnSecurityProtocol = &(o.LDAPConnSecurityProtocol)
		case "LDAPEnabled":
			sp.LDAPEnabled = &(o.LDAPEnabled)
		case "LDAPIgnoredKeys":
			sp.LDAPIgnoredKeys = &(o.LDAPIgnoredKeys)
		case "LDAPSubjectKey":
			sp.LDAPSubjectKey = &(o.LDAPSubjectKey)
		case "OTPEnabled":
			sp.OTPEnabled = &(o.OTPEnabled)
		case "OTPQRCode":
			sp.OTPQRCode = &(o.OTPQRCode)
		case "OTPSecret":
			sp.OTPSecret = &(o.OTPSecret)
		case "accessEnabled":
			sp.AccessEnabled = &(o.AccessEnabled)
		case "activationExpiration":
			sp.ActivationExpiration = &(o.ActivationExpiration)
		case "activationToken":
			sp.ActivationToken = &(o.ActivationToken)
		case "associatedAPIAuthPolicyID":
			sp.AssociatedAPIAuthPolicyID = &(o.AssociatedAPIAuthPolicyID)
		case "associatedAWSPolicies":
			sp.AssociatedAWSPolicies = &(o.AssociatedAWSPolicies)
		case "associatedBillingID":
			sp.AssociatedBillingID = &(o.AssociatedBillingID)
		case "associatedGCPPolicies":
			sp.AssociatedGCPPolicies = &(o.AssociatedGCPPolicies)
		case "associatedNamespaceID":
			sp.AssociatedNamespaceID = &(o.AssociatedNamespaceID)
		case "associatedPlanKey":
			sp.AssociatedPlanKey = &(o.AssociatedPlanKey)
		case "associatedQuotaPolicies":
			sp.AssociatedQuotaPolicies = &(o.AssociatedQuotaPolicies)
		case "company":
			sp.Company = &(o.Company)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "email":
			sp.Email = &(o.Email)
		case "firstName":
			sp.FirstName = &(o.FirstName)
		case "lastName":
			sp.LastName = &(o.LastName)
		case "name":
			sp.Name = &(o.Name)
		case "password":
			sp.Password = &(o.Password)
		case "reCAPTCHAKey":
			sp.ReCAPTCHAKey = &(o.ReCAPTCHAKey)
		case "resetPasswordExpiration":
			sp.ResetPasswordExpiration = &(o.ResetPasswordExpiration)
		case "resetPasswordToken":
			sp.ResetPasswordToken = &(o.ResetPasswordToken)
		case "status":
			sp.Status = &(o.Status)
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

// Patch apply the non nil value of a *SparseAccount to the object.
func (o *Account) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAccount)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.LDAPAddress != nil {
		o.LDAPAddress = *so.LDAPAddress
	}
	if so.LDAPBaseDN != nil {
		o.LDAPBaseDN = *so.LDAPBaseDN
	}
	if so.LDAPBindDN != nil {
		o.LDAPBindDN = *so.LDAPBindDN
	}
	if so.LDAPBindPassword != nil {
		o.LDAPBindPassword = *so.LDAPBindPassword
	}
	if so.LDAPBindSearchFilter != nil {
		o.LDAPBindSearchFilter = *so.LDAPBindSearchFilter
	}
	if so.LDAPCertificateAuthority != nil {
		o.LDAPCertificateAuthority = *so.LDAPCertificateAuthority
	}
	if so.LDAPConnSecurityProtocol != nil {
		o.LDAPConnSecurityProtocol = *so.LDAPConnSecurityProtocol
	}
	if so.LDAPEnabled != nil {
		o.LDAPEnabled = *so.LDAPEnabled
	}
	if so.LDAPIgnoredKeys != nil {
		o.LDAPIgnoredKeys = *so.LDAPIgnoredKeys
	}
	if so.LDAPSubjectKey != nil {
		o.LDAPSubjectKey = *so.LDAPSubjectKey
	}
	if so.OTPEnabled != nil {
		o.OTPEnabled = *so.OTPEnabled
	}
	if so.OTPQRCode != nil {
		o.OTPQRCode = *so.OTPQRCode
	}
	if so.OTPSecret != nil {
		o.OTPSecret = *so.OTPSecret
	}
	if so.AccessEnabled != nil {
		o.AccessEnabled = *so.AccessEnabled
	}
	if so.ActivationExpiration != nil {
		o.ActivationExpiration = *so.ActivationExpiration
	}
	if so.ActivationToken != nil {
		o.ActivationToken = *so.ActivationToken
	}
	if so.AssociatedAPIAuthPolicyID != nil {
		o.AssociatedAPIAuthPolicyID = *so.AssociatedAPIAuthPolicyID
	}
	if so.AssociatedAWSPolicies != nil {
		o.AssociatedAWSPolicies = *so.AssociatedAWSPolicies
	}
	if so.AssociatedBillingID != nil {
		o.AssociatedBillingID = *so.AssociatedBillingID
	}
	if so.AssociatedGCPPolicies != nil {
		o.AssociatedGCPPolicies = *so.AssociatedGCPPolicies
	}
	if so.AssociatedNamespaceID != nil {
		o.AssociatedNamespaceID = *so.AssociatedNamespaceID
	}
	if so.AssociatedPlanKey != nil {
		o.AssociatedPlanKey = *so.AssociatedPlanKey
	}
	if so.AssociatedQuotaPolicies != nil {
		o.AssociatedQuotaPolicies = *so.AssociatedQuotaPolicies
	}
	if so.Company != nil {
		o.Company = *so.Company
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Email != nil {
		o.Email = *so.Email
	}
	if so.FirstName != nil {
		o.FirstName = *so.FirstName
	}
	if so.LastName != nil {
		o.LastName = *so.LastName
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Password != nil {
		o.Password = *so.Password
	}
	if so.ReCAPTCHAKey != nil {
		o.ReCAPTCHAKey = *so.ReCAPTCHAKey
	}
	if so.ResetPasswordExpiration != nil {
		o.ResetPasswordExpiration = *so.ResetPasswordExpiration
	}
	if so.ResetPasswordToken != nil {
		o.ResetPasswordToken = *so.ResetPasswordToken
	}
	if so.Status != nil {
		o.Status = *so.Status
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

// DeepCopy returns a deep copy if the Account.
func (o *Account) DeepCopy() *Account {

	if o == nil {
		return nil
	}

	out := &Account{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Account.
func (o *Account) DeepCopyInto(out *Account) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Account: %s", err))
	}

	*out = *target.(*Account)
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

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, `must not contain any '*' or '='`, true); err != nil {
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Account) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "LDAPAddress":
		return o.LDAPAddress
	case "LDAPBaseDN":
		return o.LDAPBaseDN
	case "LDAPBindDN":
		return o.LDAPBindDN
	case "LDAPBindPassword":
		return o.LDAPBindPassword
	case "LDAPBindSearchFilter":
		return o.LDAPBindSearchFilter
	case "LDAPCertificateAuthority":
		return o.LDAPCertificateAuthority
	case "LDAPConnSecurityProtocol":
		return o.LDAPConnSecurityProtocol
	case "LDAPEnabled":
		return o.LDAPEnabled
	case "LDAPIgnoredKeys":
		return o.LDAPIgnoredKeys
	case "LDAPSubjectKey":
		return o.LDAPSubjectKey
	case "OTPEnabled":
		return o.OTPEnabled
	case "OTPQRCode":
		return o.OTPQRCode
	case "OTPSecret":
		return o.OTPSecret
	case "accessEnabled":
		return o.AccessEnabled
	case "activationExpiration":
		return o.ActivationExpiration
	case "activationToken":
		return o.ActivationToken
	case "associatedAPIAuthPolicyID":
		return o.AssociatedAPIAuthPolicyID
	case "associatedAWSPolicies":
		return o.AssociatedAWSPolicies
	case "associatedBillingID":
		return o.AssociatedBillingID
	case "associatedGCPPolicies":
		return o.AssociatedGCPPolicies
	case "associatedNamespaceID":
		return o.AssociatedNamespaceID
	case "associatedPlanKey":
		return o.AssociatedPlanKey
	case "associatedQuotaPolicies":
		return o.AssociatedQuotaPolicies
	case "company":
		return o.Company
	case "createTime":
		return o.CreateTime
	case "email":
		return o.Email
	case "firstName":
		return o.FirstName
	case "lastName":
		return o.LastName
	case "name":
		return o.Name
	case "password":
		return o.Password
	case "reCAPTCHAKey":
		return o.ReCAPTCHAKey
	case "resetPasswordExpiration":
		return o.ResetPasswordExpiration
	case "resetPasswordToken":
		return o.ResetPasswordToken
	case "status":
		return o.Status
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
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
		Exposed:   true,
		Name:      "LDAPBindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"LDAPCertificateAuthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPCertificateAuthority",
		Description: `LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:   true,
		Name:      "LDAPCertificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		Exposed:   true,
		Name:      "LDAPIgnoredKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"LDAPSubjectKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPSubjectKey",
		DefaultValue:   "uid",
		Description: `LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:   true,
		Name:      "LDAPSubjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"OTPSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the 2 factor secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"AccessEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
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
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
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
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"AssociatedBillingID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `associatedBillingID holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedGCPPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedGCPPolicies",
		Description:    `associatedGCPPolicies contains a map of associated GCP Enforcerd Policies.`,
		Name:           "associatedGCPPolicies",
		Stored:         true,
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"AssociatedNamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
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
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"Company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
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
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
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

// AccountLowerCaseAttributesMap represents the map of attribute for Account.
var AccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ldapaddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPAddress",
		Description:    `LDAPAddress holds the account authentication account's private ldap server.`,
		Exposed:        true,
		Filterable:     true,
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
		Exposed:   true,
		Name:      "LDAPBindSearchFilter",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"ldapcertificateauthority": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPCertificateAuthority",
		Description: `LDAPCertificateAuthority contains the optional certificate author ity that will
be used to connect to the LDAP server. It is not needed if the TLS certificate
of the LDAP is issued from a public truster CA.`,
		Exposed:   true,
		Name:      "LDAPCertificateAuthority",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		Exposed:   true,
		Name:      "LDAPIgnoredKeys",
		Orderable: true,
		Stored:    true,
		SubType:   "string",
		Type:      "list",
	},
	"ldapsubjectkey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LDAPSubjectKey",
		DefaultValue:   "uid",
		Description: `LDAPSubjectKey holds key to be used to populate the subject. If you want to
use the user as a subject, for Windows based systems you may use
'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
also use any alternate key.`,
		Exposed:   true,
		Name:      "LDAPSubjectKey",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"otpsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the 2 factor secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"accessenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `AccessEnabled defines if the account holder should have access to the systems.`,
		Exposed:        true,
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
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"associatedapiauthpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.`,
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
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"associatedbillingid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `associatedBillingID holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"associatedgcppolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedGCPPolicies",
		Description:    `associatedGCPPolicies contains a map of associated GCP Enforcerd Policies.`,
		Name:           "associatedGCPPolicies",
		Stored:         true,
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"associatednamespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `AssociatedNamespaceID contains the ID of the associated namespace.`,
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
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"company": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Company",
		Description:    `Company of the account user.`,
		Exposed:        true,
		Filterable:     true,
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
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
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

// SparseAccountsList represents a list of SparseAccounts
type SparseAccountsList []*SparseAccount

// Identity returns the identity of the objects in the list.
func (o SparseAccountsList) Identity() elemental.Identity {

	return AccountIdentity
}

// Copy returns a pointer to a copy the SparseAccountsList.
func (o SparseAccountsList) Copy() elemental.Identifiables {

	copy := append(SparseAccountsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccountsList.
func (o SparseAccountsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccountsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccount))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccountsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccountsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccountsList converted to AccountsList.
func (o SparseAccountsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccountsList) Version() int {

	return 1
}

// SparseAccount represents the sparse version of a account.
type SparseAccount struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// LDAPAddress holds the account authentication account's private ldap server.
	LDAPAddress *string `json:"LDAPAddress,omitempty" bson:"ldapaddress" mapstructure:"LDAPAddress,omitempty"`

	// LDAPBaseDN holds the base DN to use to ldap queries.
	LDAPBaseDN *string `json:"LDAPBaseDN,omitempty" bson:"ldapbasedn" mapstructure:"LDAPBaseDN,omitempty"`

	// LDAPBindDN holds the account's internal LDAP bind dn for querying auth.
	LDAPBindDN *string `json:"LDAPBindDN,omitempty" bson:"ldapbinddn" mapstructure:"LDAPBindDN,omitempty"`

	// LDAPBindPassword holds the password to the LDAPBindDN.
	LDAPBindPassword *string `json:"LDAPBindPassword,omitempty" bson:"ldapbindpassword" mapstructure:"LDAPBindPassword,omitempty"`

	// LDAPBindSearchFilter holds filter to be used to uniquely search a user. For
	// Windows based systems, value may be 'sAMAccountName={USERNAME}'. For Linux and
	// other systems, value may be 'uid={USERNAME}'.
	LDAPBindSearchFilter *string `json:"LDAPBindSearchFilter,omitempty" bson:"ldapbindsearchfilter" mapstructure:"LDAPBindSearchFilter,omitempty"`

	// LDAPCertificateAuthority contains the optional certificate author ity that will
	// be used to connect to the LDAP server. It is not needed if the TLS certificate
	// of the LDAP is issued from a public truster CA.
	LDAPCertificateAuthority *string `json:"LDAPCertificateAuthority,omitempty" bson:"ldapcertificateauthority" mapstructure:"LDAPCertificateAuthority,omitempty"`

	// LDAPConnProtocol holds the connection type for the LDAP provider.
	LDAPConnSecurityProtocol *AccountLDAPConnSecurityProtocolValue `json:"LDAPConnSecurityProtocol,omitempty" bson:"ldapconnsecurityprotocol" mapstructure:"LDAPConnSecurityProtocol,omitempty"`

	// LDAPEnabled triggers if the account uses it's own LDAP for authentication.
	LDAPEnabled *bool `json:"LDAPEnabled,omitempty" bson:"ldapenabled" mapstructure:"LDAPEnabled,omitempty"`

	// LDAPIgnoredKeys holds a list of keys that must not be imported into Aporeto
	// authorization system.
	LDAPIgnoredKeys *[]string `json:"LDAPIgnoredKeys,omitempty" bson:"ldapignoredkeys" mapstructure:"LDAPIgnoredKeys,omitempty"`

	// LDAPSubjectKey holds key to be used to populate the subject. If you want to
	// use the user as a subject, for Windows based systems you may use
	// 'sAMAccountName' and for Linux and other systems, value may be 'uid'. You can
	// also use any alternate key.
	LDAPSubjectKey *string `json:"LDAPSubjectKey,omitempty" bson:"ldapsubjectkey" mapstructure:"LDAPSubjectKey,omitempty"`

	// Set to enable or disable two factor authentication.
	OTPEnabled *bool `json:"OTPEnabled,omitempty" bson:"otpenabled" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64 encoded QRCode for setting up 2 factor auth.
	OTPQRCode *string `json:"OTPQRCode,omitempty" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the 2 factor secret.
	OTPSecret *string `json:"-,omitempty" bson:"otpsecret" mapstructure:"-,omitempty"`

	// AccessEnabled defines if the account holder should have access to the systems.
	AccessEnabled *bool `json:"accessEnabled,omitempty" bson:"accessenabled" mapstructure:"accessEnabled,omitempty"`

	// ActivationExpiration contains the expiration date of the activation token.
	ActivationExpiration *time.Time `json:"-,omitempty" bson:"activationexpiration" mapstructure:"-,omitempty"`

	// ActivationToken contains the activation token.
	ActivationToken *string `json:"activationToken,omitempty" bson:"activationtoken" mapstructure:"activationToken,omitempty"`

	// AssociatedAPIAuthPolicyID holds the ID of the associated API auth policy.
	AssociatedAPIAuthPolicyID *string `json:"-,omitempty" bson:"associatedapiauthpolicyid" mapstructure:"-,omitempty"`

	// AssociatedAWSPolicies contains a map of associated AWS Enforcerd Policies.
	AssociatedAWSPolicies *map[string]string `json:"-,omitempty" bson:"associatedawspolicies" mapstructure:"-,omitempty"`

	// associatedBillingID holds the ID of the associated billing customer.
	AssociatedBillingID *string `json:"associatedBillingID,omitempty" bson:"associatedbillingid" mapstructure:"associatedBillingID,omitempty"`

	// associatedGCPPolicies contains a map of associated GCP Enforcerd Policies.
	AssociatedGCPPolicies *map[string]string `json:"-,omitempty" bson:"associatedgcppolicies" mapstructure:"-,omitempty"`

	// AssociatedNamespaceID contains the ID of the associated namespace.
	AssociatedNamespaceID *string `json:"-,omitempty" bson:"associatednamespaceid" mapstructure:"-,omitempty"`

	// AssociatedPlanKey contains the plan key that is associated to this account.
	AssociatedPlanKey *string `json:"associatedPlanKey,omitempty" bson:"associatedplankey" mapstructure:"associatedPlanKey,omitempty"`

	// AssociatedQuotaPolicies contains a mapping to the associated quota pollicies.
	AssociatedQuotaPolicies *map[string]string `json:"-,omitempty" bson:"associatedquotapolicies" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company *string `json:"company,omitempty" bson:"company" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email *string `json:"email,omitempty" bson:"email" mapstructure:"email,omitempty"`

	// First Name of the account user.
	FirstName *string `json:"firstName,omitempty" bson:"firstname" mapstructure:"firstName,omitempty"`

	// Last Name of the account user.
	LastName *string `json:"lastName,omitempty" bson:"lastname" mapstructure:"lastName,omitempty"`

	// Name of the account.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Password for the account.
	Password *string `json:"password,omitempty" bson:"password" mapstructure:"password,omitempty"`

	// ReCAPTCHAKey contains the capcha validation if reCAPTCH is enabled.
	ReCAPTCHAKey *string `json:"reCAPTCHAKey,omitempty" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// ResetPasswordExpiration contains the expiration time for reseting the password.
	ResetPasswordExpiration *time.Time `json:"-,omitempty" bson:"resetpasswordexpiration" mapstructure:"-,omitempty"`

	// ResetPasswordToken contains the token to use for resetting password.
	ResetPasswordToken *string `json:"-,omitempty" bson:"resetpasswordtoken" mapstructure:"-,omitempty"`

	// Status of the account.
	Status *AccountStatusValue `json:"status,omitempty" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-,omitempty" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-,omitempty" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAccount returns a new  SparseAccount.
func NewSparseAccount() *SparseAccount {
	return &SparseAccount{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccount) Identity() elemental.Identity {

	return AccountIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccount) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccount) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAccount) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccount) ToPlain() elemental.PlainIdentifiable {

	out := NewAccount()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.LDAPAddress != nil {
		out.LDAPAddress = *o.LDAPAddress
	}
	if o.LDAPBaseDN != nil {
		out.LDAPBaseDN = *o.LDAPBaseDN
	}
	if o.LDAPBindDN != nil {
		out.LDAPBindDN = *o.LDAPBindDN
	}
	if o.LDAPBindPassword != nil {
		out.LDAPBindPassword = *o.LDAPBindPassword
	}
	if o.LDAPBindSearchFilter != nil {
		out.LDAPBindSearchFilter = *o.LDAPBindSearchFilter
	}
	if o.LDAPCertificateAuthority != nil {
		out.LDAPCertificateAuthority = *o.LDAPCertificateAuthority
	}
	if o.LDAPConnSecurityProtocol != nil {
		out.LDAPConnSecurityProtocol = *o.LDAPConnSecurityProtocol
	}
	if o.LDAPEnabled != nil {
		out.LDAPEnabled = *o.LDAPEnabled
	}
	if o.LDAPIgnoredKeys != nil {
		out.LDAPIgnoredKeys = *o.LDAPIgnoredKeys
	}
	if o.LDAPSubjectKey != nil {
		out.LDAPSubjectKey = *o.LDAPSubjectKey
	}
	if o.OTPEnabled != nil {
		out.OTPEnabled = *o.OTPEnabled
	}
	if o.OTPQRCode != nil {
		out.OTPQRCode = *o.OTPQRCode
	}
	if o.OTPSecret != nil {
		out.OTPSecret = *o.OTPSecret
	}
	if o.AccessEnabled != nil {
		out.AccessEnabled = *o.AccessEnabled
	}
	if o.ActivationExpiration != nil {
		out.ActivationExpiration = *o.ActivationExpiration
	}
	if o.ActivationToken != nil {
		out.ActivationToken = *o.ActivationToken
	}
	if o.AssociatedAPIAuthPolicyID != nil {
		out.AssociatedAPIAuthPolicyID = *o.AssociatedAPIAuthPolicyID
	}
	if o.AssociatedAWSPolicies != nil {
		out.AssociatedAWSPolicies = *o.AssociatedAWSPolicies
	}
	if o.AssociatedBillingID != nil {
		out.AssociatedBillingID = *o.AssociatedBillingID
	}
	if o.AssociatedGCPPolicies != nil {
		out.AssociatedGCPPolicies = *o.AssociatedGCPPolicies
	}
	if o.AssociatedNamespaceID != nil {
		out.AssociatedNamespaceID = *o.AssociatedNamespaceID
	}
	if o.AssociatedPlanKey != nil {
		out.AssociatedPlanKey = *o.AssociatedPlanKey
	}
	if o.AssociatedQuotaPolicies != nil {
		out.AssociatedQuotaPolicies = *o.AssociatedQuotaPolicies
	}
	if o.Company != nil {
		out.Company = *o.Company
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Email != nil {
		out.Email = *o.Email
	}
	if o.FirstName != nil {
		out.FirstName = *o.FirstName
	}
	if o.LastName != nil {
		out.LastName = *o.LastName
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Password != nil {
		out.Password = *o.Password
	}
	if o.ReCAPTCHAKey != nil {
		out.ReCAPTCHAKey = *o.ReCAPTCHAKey
	}
	if o.ResetPasswordExpiration != nil {
		out.ResetPasswordExpiration = *o.ResetPasswordExpiration
	}
	if o.ResetPasswordToken != nil {
		out.ResetPasswordToken = *o.ResetPasswordToken
	}
	if o.Status != nil {
		out.Status = *o.Status
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

// GetZHash returns the ZHash of the receiver.
func (o *SparseAccount) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAccount) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAccount) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAccount) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAccount.
func (o *SparseAccount) DeepCopy() *SparseAccount {

	if o == nil {
		return nil
	}

	out := &SparseAccount{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccount.
func (o *SparseAccount) DeepCopyInto(out *SparseAccount) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccount: %s", err))
	}

	*out = *target.(*SparseAccount)
}
