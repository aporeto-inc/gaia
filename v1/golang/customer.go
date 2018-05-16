package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// CustomerProviderValue represents the possible values for attribute "provider".
type CustomerProviderValue string

const (
	// CustomerProviderAWS represents the value AWS.
	CustomerProviderAWS CustomerProviderValue = "AWS"

	// CustomerProviderAporeto represents the value Aporeto.
	CustomerProviderAporeto CustomerProviderValue = "Aporeto"
)

// CustomerStateValue represents the possible values for attribute "state".
type CustomerStateValue string

const (
	// CustomerStateSubscribeFailed represents the value SubscribeFailed.
	CustomerStateSubscribeFailed CustomerStateValue = "SubscribeFailed"

	// CustomerStateSubscribePending represents the value SubscribePending.
	CustomerStateSubscribePending CustomerStateValue = "SubscribePending"

	// CustomerStateSubscribeSuccess represents the value SubscribeSuccess.
	CustomerStateSubscribeSuccess CustomerStateValue = "SubscribeSuccess"

	// CustomerStateUnsubscribePending represents the value UnsubscribePending.
	CustomerStateUnsubscribePending CustomerStateValue = "UnsubscribePending"

	// CustomerStateUnsubscribeSuccess represents the value UnsubscribeSuccess.
	CustomerStateUnsubscribeSuccess CustomerStateValue = "UnsubscribeSuccess"
)

// CustomerIdentity represents the Identity of the object.
var CustomerIdentity = elemental.Identity{
	Name:     "customer",
	Category: "customers",
	Private:  true,
}

// CustomersList represents a list of Customers
type CustomersList []*Customer

// ContentIdentity returns the identity of the objects in the list.
func (o CustomersList) ContentIdentity() elemental.Identity {

	return CustomerIdentity
}

// Copy returns a pointer to a copy the CustomersList.
func (o CustomersList) Copy() elemental.ContentIdentifiable {

	copy := append(CustomersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CustomersList.
func (o CustomersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(CustomersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Customer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CustomersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CustomersList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o CustomersList) Version() int {

	return 1
}

// Customer represents the model of a customer
type Customer struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Provider holds the name of the provider to be billed for this service.
	Provider CustomerProviderValue `json:"provider" bson:"provider" mapstructure:"provider,omitempty"`

	// ProviderCustomerID holds the customer id as used by the provider for this
	// customer to enable provider billing.
	ProviderCustomerID string `json:"providerCustomerID" bson:"providercustomerid" mapstructure:"providerCustomerID,omitempty"`

	// State holds the status of the customer with the provider.
	State CustomerStateValue `json:"state" bson:"state" mapstructure:"state,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewCustomer returns a new *Customer
func NewCustomer() *Customer {

	return &Customer{
		ModelVersion: 1,
		Provider:     "Aporeto",
		State:        "SubscribePending",
	}
}

// Identity returns the Identity of the object.
func (o *Customer) Identity() elemental.Identity {

	return CustomerIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Customer) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Customer) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Customer) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Customer) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Customer) Doc() string {
	return `This api allows to view and manage basic information about customer profile for
billing purposes.`
}

func (o *Customer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Customer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("provider", string(o.Provider), []string{"Aporeto", "AWS"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("state", string(o.State), []string{"SubscribePending", "SubscribeFailed", "SubscribeSuccess", "UnsubscribePending", "UnsubscribeSuccess"}, false); err != nil {
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
func (*Customer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CustomerAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CustomerLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Customer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CustomerAttributesMap
}

// CustomerAttributesMap represents the map of attribute for Customer.
var CustomerAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Provider": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "Provider",
		DefaultValue:   CustomerProviderAporeto,
		Description:    `Provider holds the name of the provider to be billed for this service.`,
		Exposed:        true,
		Name:           "provider",
		Stored:         true,
		Type:           "enum",
	},
	"ProviderCustomerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProviderCustomerID",
		Description: `ProviderCustomerID holds the customer id as used by the provider for this
customer to enable provider billing.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "providerCustomerID",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"State": elemental.AttributeSpecification{
		AllowedChoices: []string{"SubscribePending", "SubscribeFailed", "SubscribeSuccess", "UnsubscribePending", "UnsubscribeSuccess"},
		ConvertedName:  "State",
		DefaultValue:   CustomerStateSubscribePending,
		Description:    `State holds the status of the customer with the provider.`,
		Exposed:        true,
		Name:           "state",
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

// CustomerLowerCaseAttributesMap represents the map of attribute for Customer.
var CustomerLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"provider": elemental.AttributeSpecification{
		AllowedChoices: []string{"Aporeto", "AWS"},
		ConvertedName:  "Provider",
		DefaultValue:   CustomerProviderAporeto,
		Description:    `Provider holds the name of the provider to be billed for this service.`,
		Exposed:        true,
		Name:           "provider",
		Stored:         true,
		Type:           "enum",
	},
	"providercustomerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProviderCustomerID",
		Description: `ProviderCustomerID holds the customer id as used by the provider for this
customer to enable provider billing.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "providerCustomerID",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"state": elemental.AttributeSpecification{
		AllowedChoices: []string{"SubscribePending", "SubscribeFailed", "SubscribeSuccess", "UnsubscribePending", "UnsubscribeSuccess"},
		ConvertedName:  "State",
		DefaultValue:   CustomerStateSubscribePending,
		Description:    `State holds the status of the customer with the provider.`,
		Exposed:        true,
		Name:           "state",
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
