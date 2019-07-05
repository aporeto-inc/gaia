package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
func (o AccountsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAccountsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAccount)
	}

	return out
}

// Version returns the version of the content.
func (o AccountsList) Version() int {

	return 1
}

// Account represents the model of a account
type Account struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Enable or disable two-factor authentication.
	OTPEnabled bool `json:"OTPEnabled" msgpack:"OTPEnabled" bson:"otpenabled" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64-encoded QR code for setting up two-factor authentication.
	OTPQRCode string `json:"OTPQRCode" msgpack:"OTPQRCode" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the two-factor authentication secret.
	OTPSecret string `json:"-" msgpack:"-" bson:"otpsecret" mapstructure:"-,omitempty"`

	// Holds the SSH certificate authority used by the account namespace.
	SSHCA string `json:"SSHCA" msgpack:"SSHCA" bson:"sshca" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the account namespace.
	SSHCARenew bool `json:"SSHCARenew" msgpack:"SSHCARenew" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	// Defines if the account holder should have access to the system.
	AccessEnabled bool `json:"accessEnabled" msgpack:"accessEnabled" bson:"accessenabled" mapstructure:"accessEnabled,omitempty"`

	// Contains the expiration date of the activation token.
	ActivationExpiration time.Time `json:"-" msgpack:"-" bson:"activationexpiration" mapstructure:"-,omitempty"`

	// Contains the activation token.
	ActivationToken string `json:"activationToken,omitempty" msgpack:"activationToken,omitempty" bson:"activationtoken" mapstructure:"activationToken,omitempty"`

	// Holds the ID of the associated API authorization.
	AssociatedAPIAuthPolicyID string `json:"-" msgpack:"-" bson:"associatedapiauthpolicyid" mapstructure:"-,omitempty"`

	// Contains a map of associated AWS enforcer policies.
	AssociatedAWSPolicies map[string]string `json:"-" msgpack:"-" bson:"associatedawspolicies" mapstructure:"-,omitempty"`

	// Holds the ID of the associated billing customer.
	AssociatedBillingID string `json:"associatedBillingID" msgpack:"associatedBillingID" bson:"associatedbillingid" mapstructure:"associatedBillingID,omitempty"`

	// Contains the ID of the associated namespace.
	AssociatedNamespaceID string `json:"-" msgpack:"-" bson:"associatednamespaceid" mapstructure:"-,omitempty"`

	// Contains the plan key associated with this account.
	AssociatedPlanKey string `json:"associatedPlanKey" msgpack:"associatedPlanKey" bson:"associatedplankey" mapstructure:"associatedPlanKey,omitempty"`

	// Contains a map of the associated quotas.
	AssociatedQuotaPolicies map[string]string `json:"-" msgpack:"-" bson:"associatedquotapolicies" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company string `json:"company" msgpack:"company" bson:"company" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email string `json:"email" msgpack:"email" bson:"email" mapstructure:"email,omitempty"`

	// First name of the account user.
	FirstName string `json:"firstName" msgpack:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// Last name of the account user.
	LastName string `json:"lastName" msgpack:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// Name of the account.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Password for the account.
	Password string `json:"password" msgpack:"password" bson:"password" mapstructure:"password,omitempty"`

	// Contains the completely automated public Turing test (CAPTCHA)
	// validation if reCAPTCHA is enabled.
	ReCAPTCHAKey string `json:"reCAPTCHAKey" msgpack:"reCAPTCHAKey" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// Contains the expiration time for resetting the password.
	ResetPasswordExpiration time.Time `json:"-" msgpack:"-" bson:"resetpasswordexpiration" mapstructure:"-,omitempty"`

	// Contains the token to use for resetting password.
	ResetPasswordToken string `json:"-" msgpack:"-" bson:"resetpasswordtoken" mapstructure:"-,omitempty"`

	// Status of the account.
	Status AccountStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAccount returns a new *Account
func NewAccount() *Account {

	return &Account{
		ModelVersion:            1,
		AssociatedAWSPolicies:   map[string]string{},
		AssociatedQuotaPolicies: map[string]string{},
		Status:                  AccountStatusPending,
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

// BleveType implements the bleve.Classifier Interface.
func (o *Account) BleveType() string {

	return "account"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Account) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Account) Doc() string {

	return `Allows you to view and manage basic information about your account like
your name, password, and whether or not two-factor authentication is enabled.`
}

func (o *Account) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Account) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Account) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Account) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Account) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
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
			OTPEnabled:                &o.OTPEnabled,
			OTPQRCode:                 &o.OTPQRCode,
			OTPSecret:                 &o.OTPSecret,
			SSHCA:                     &o.SSHCA,
			SSHCARenew:                &o.SSHCARenew,
			AccessEnabled:             &o.AccessEnabled,
			ActivationExpiration:      &o.ActivationExpiration,
			ActivationToken:           &o.ActivationToken,
			AssociatedAPIAuthPolicyID: &o.AssociatedAPIAuthPolicyID,
			AssociatedAWSPolicies:     &o.AssociatedAWSPolicies,
			AssociatedBillingID:       &o.AssociatedBillingID,
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
		case "OTPEnabled":
			sp.OTPEnabled = &(o.OTPEnabled)
		case "OTPQRCode":
			sp.OTPQRCode = &(o.OTPQRCode)
		case "OTPSecret":
			sp.OTPSecret = &(o.OTPSecret)
		case "SSHCA":
			sp.SSHCA = &(o.SSHCA)
		case "SSHCARenew":
			sp.SSHCARenew = &(o.SSHCARenew)
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
	if so.OTPEnabled != nil {
		o.OTPEnabled = *so.OTPEnabled
	}
	if so.OTPQRCode != nil {
		o.OTPQRCode = *so.OTPQRCode
	}
	if so.OTPSecret != nil {
		o.OTPSecret = *so.OTPSecret
	}
	if so.SSHCA != nil {
		o.SSHCA = *so.SSHCA
	}
	if so.SSHCARenew != nil {
		o.SSHCARenew = *so.SSHCARenew
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

	if err := elemental.ValidateRequiredString("email", o.Email); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, `must not contain any '*' or '='`, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Active", "Disabled", "Invited", "Pending"}, true); err != nil {
		errors = errors.Append(err)
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
	case "OTPEnabled":
		return o.OTPEnabled
	case "OTPQRCode":
		return o.OTPQRCode
	case "OTPSecret":
		return o.OTPSecret
	case "SSHCA":
		return o.SSHCA
	case "SSHCARenew":
		return o.SSHCARenew
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
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"OTPEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Enable or disable two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"OTPQRCode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64-encoded QR code for setting up two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"OTPSecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the two-factor authentication secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"SSHCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `Holds the SSH certificate authority used by the account namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"SSHCARenew": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		Type:           "boolean",
	},
	"AccessEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `Defines if the account holder should have access to the system.`,
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
		Description:    `Contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ActivationToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAPIAuthPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `Holds the ID of the associated API authorization.`,
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedAWSPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `Contains a map of associated AWS enforcer policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"AssociatedBillingID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `Holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedNamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `Contains the ID of the associated namespace.`,
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedPlanKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedPlanKey",
		CreationOnly:   true,
		Description:    `Contains the plan key associated with this account.`,
		Exposed:        true,
		Name:           "associatedPlanKey",
		Stored:         true,
		Type:           "string",
	},
	"AssociatedQuotaPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `Contains a map of the associated quotas.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "map[string]string",
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
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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
		Description:    `First name of the account user.`,
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
		Description:    `Last name of the account user.`,
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
		Description: `Contains the completely automated public Turing test (CAPTCHA)
validation if reCAPTCHA is enabled.`,
		Exposed: true,
		Name:    "reCAPTCHAKey",
		Type:    "string",
	},
	"ResetPasswordExpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `Contains the expiration time for resetting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"ResetPasswordToken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `Contains the token to use for resetting password.`,
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// AccountLowerCaseAttributesMap represents the map of attribute for Account.
var AccountLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"otpenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPEnabled",
		Description:    `Enable or disable two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"otpqrcode": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "OTPQRCode",
		Description:    `Returns the base64-encoded QR code for setting up two-factor authentication.`,
		Exposed:        true,
		Name:           "OTPQRCode",
		ReadOnly:       true,
		Transient:      true,
		Type:           "string",
	},
	"otpsecret": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "OTPSecret",
		Description:    `Stores the two-factor authentication secret.`,
		Name:           "OTPSecret",
		Stored:         true,
		Type:           "string",
	},
	"sshca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SSHCA",
		Description:    `Holds the SSH certificate authority used by the account namespace.`,
		Exposed:        true,
		Name:           "SSHCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"sshcarenew": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SSHCARenew",
		Description:    `Set to ` + "`" + `true` + "`" + ` to renew the SSH certificate authority of the account namespace.`,
		Exposed:        true,
		Name:           "SSHCARenew",
		Type:           "boolean",
	},
	"accessenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccessEnabled",
		Description:    `Defines if the account holder should have access to the system.`,
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
		Description:    `Contains the expiration date of the activation token.`,
		Name:           "activationExpiration",
		Stored:         true,
		Type:           "time",
	},
	"activationtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ActivationToken",
		Description:    `Contains the activation token.`,
		Exposed:        true,
		Name:           "activationToken",
		Stored:         true,
		Type:           "string",
	},
	"associatedapiauthpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAPIAuthPolicyID",
		Description:    `Holds the ID of the associated API authorization.`,
		Name:           "associatedAPIAuthPolicyID",
		Stored:         true,
		Type:           "string",
	},
	"associatedawspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedAWSPolicies",
		Description:    `Contains a map of associated AWS enforcer policies.`,
		Name:           "associatedAWSPolicies",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"associatedbillingid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedBillingID",
		Description:    `Holds the ID of the associated billing customer.`,
		Exposed:        true,
		Name:           "associatedBillingID",
		Stored:         true,
		Type:           "string",
	},
	"associatednamespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedNamespaceID",
		Description:    `Contains the ID of the associated namespace.`,
		Name:           "associatedNamespaceID",
		Stored:         true,
		Type:           "string",
	},
	"associatedplankey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedPlanKey",
		CreationOnly:   true,
		Description:    `Contains the plan key associated with this account.`,
		Exposed:        true,
		Name:           "associatedPlanKey",
		Stored:         true,
		Type:           "string",
	},
	"associatedquotapolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedQuotaPolicies",
		Description:    `Contains a map of the associated quotas.`,
		Name:           "associatedQuotaPolicies",
		Stored:         true,
		SubType:        "map[string]string",
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
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
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
		Description:    `First name of the account user.`,
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
		Description:    `Last name of the account user.`,
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
		Description: `Contains the completely automated public Turing test (CAPTCHA)
validation if reCAPTCHA is enabled.`,
		Exposed: true,
		Name:    "reCAPTCHAKey",
		Type:    "string",
	},
	"resetpasswordexpiration": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordExpiration",
		Description:    `Contains the expiration time for resetting the password.`,
		Name:           "resetPasswordExpiration",
		Stored:         true,
		Type:           "time",
	},
	"resetpasswordtoken": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ResetPasswordToken",
		Description:    `Contains the token to use for resetting password.`,
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
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
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Enable or disable two-factor authentication.
	OTPEnabled *bool `json:"OTPEnabled,omitempty" msgpack:"OTPEnabled,omitempty" bson:"otpenabled,omitempty" mapstructure:"OTPEnabled,omitempty"`

	// Returns the base64-encoded QR code for setting up two-factor authentication.
	OTPQRCode *string `json:"OTPQRCode,omitempty" msgpack:"OTPQRCode,omitempty" bson:"-" mapstructure:"OTPQRCode,omitempty"`

	// Stores the two-factor authentication secret.
	OTPSecret *string `json:"-" msgpack:"-" bson:"otpsecret,omitempty" mapstructure:"-,omitempty"`

	// Holds the SSH certificate authority used by the account namespace.
	SSHCA *string `json:"SSHCA,omitempty" msgpack:"SSHCA,omitempty" bson:"sshca,omitempty" mapstructure:"SSHCA,omitempty"`

	// Set to `true` to renew the SSH certificate authority of the account namespace.
	SSHCARenew *bool `json:"SSHCARenew,omitempty" msgpack:"SSHCARenew,omitempty" bson:"-" mapstructure:"SSHCARenew,omitempty"`

	// Defines if the account holder should have access to the system.
	AccessEnabled *bool `json:"accessEnabled,omitempty" msgpack:"accessEnabled,omitempty" bson:"accessenabled,omitempty" mapstructure:"accessEnabled,omitempty"`

	// Contains the expiration date of the activation token.
	ActivationExpiration *time.Time `json:"-" msgpack:"-" bson:"activationexpiration,omitempty" mapstructure:"-,omitempty"`

	// Contains the activation token.
	ActivationToken *string `json:"activationToken,omitempty" msgpack:"activationToken,omitempty" bson:"activationtoken,omitempty" mapstructure:"activationToken,omitempty"`

	// Holds the ID of the associated API authorization.
	AssociatedAPIAuthPolicyID *string `json:"-" msgpack:"-" bson:"associatedapiauthpolicyid,omitempty" mapstructure:"-,omitempty"`

	// Contains a map of associated AWS enforcer policies.
	AssociatedAWSPolicies *map[string]string `json:"-" msgpack:"-" bson:"associatedawspolicies,omitempty" mapstructure:"-,omitempty"`

	// Holds the ID of the associated billing customer.
	AssociatedBillingID *string `json:"associatedBillingID,omitempty" msgpack:"associatedBillingID,omitempty" bson:"associatedbillingid,omitempty" mapstructure:"associatedBillingID,omitempty"`

	// Contains the ID of the associated namespace.
	AssociatedNamespaceID *string `json:"-" msgpack:"-" bson:"associatednamespaceid,omitempty" mapstructure:"-,omitempty"`

	// Contains the plan key associated with this account.
	AssociatedPlanKey *string `json:"associatedPlanKey,omitempty" msgpack:"associatedPlanKey,omitempty" bson:"associatedplankey,omitempty" mapstructure:"associatedPlanKey,omitempty"`

	// Contains a map of the associated quotas.
	AssociatedQuotaPolicies *map[string]string `json:"-" msgpack:"-" bson:"associatedquotapolicies,omitempty" mapstructure:"-,omitempty"`

	// Company of the account user.
	Company *string `json:"company,omitempty" msgpack:"company,omitempty" bson:"company,omitempty" mapstructure:"company,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Email of the account holder.
	Email *string `json:"email,omitempty" msgpack:"email,omitempty" bson:"email,omitempty" mapstructure:"email,omitempty"`

	// First name of the account user.
	FirstName *string `json:"firstName,omitempty" msgpack:"firstName,omitempty" bson:"firstname,omitempty" mapstructure:"firstName,omitempty"`

	// Last name of the account user.
	LastName *string `json:"lastName,omitempty" msgpack:"lastName,omitempty" bson:"lastname,omitempty" mapstructure:"lastName,omitempty"`

	// Name of the account.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Password for the account.
	Password *string `json:"password,omitempty" msgpack:"password,omitempty" bson:"password,omitempty" mapstructure:"password,omitempty"`

	// Contains the completely automated public Turing test (CAPTCHA)
	// validation if reCAPTCHA is enabled.
	ReCAPTCHAKey *string `json:"reCAPTCHAKey,omitempty" msgpack:"reCAPTCHAKey,omitempty" bson:"-" mapstructure:"reCAPTCHAKey,omitempty"`

	// Contains the expiration time for resetting the password.
	ResetPasswordExpiration *time.Time `json:"-" msgpack:"-" bson:"resetpasswordexpiration,omitempty" mapstructure:"-,omitempty"`

	// Contains the token to use for resetting password.
	ResetPasswordToken *string `json:"-" msgpack:"-" bson:"resetpasswordtoken,omitempty" mapstructure:"-,omitempty"`

	// Status of the account.
	Status *AccountStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
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
	if o.OTPEnabled != nil {
		out.OTPEnabled = *o.OTPEnabled
	}
	if o.OTPQRCode != nil {
		out.OTPQRCode = *o.OTPQRCode
	}
	if o.OTPSecret != nil {
		out.OTPSecret = *o.OTPSecret
	}
	if o.SSHCA != nil {
		out.SSHCA = *o.SSHCA
	}
	if o.SSHCARenew != nil {
		out.SSHCARenew = *o.SSHCARenew
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

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAccount) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAccount) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAccount) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAccount) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
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
