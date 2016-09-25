package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// NotificationIdentity represents the Identity of the object
var NotificationIdentity = elemental.Identity{
	Name:     "notification",
	Category: "notifications",
}

// NotificationsList represents a list of Notifications
type NotificationsList []*Notification

// Notification represents the model of a notification
type Notification struct {
	// Created is the time when the notification was created
	Created string `json:"Created" cql:"created,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-"`

	// Deleted is the time when the notification was deleted
	Deleted string `json:"deleted" cql:"deleted,omitempty"`

	// Limits the amount of results in the "LayersIntroducingVulnerability" property on New and Old vulnerabilities
	Limit string `json:"limit" cql:"limit,omitempty"`

	// Name is the name of the notification
	Name string `json:"name" cql:"name,omitempty"`

	// New is the new layers that introduced vulnerability
	New interface{} `json:"new" cql:"new,omitempty"`

	// NextPage is the next page number
	NextPage string `json:"nextPage" cql:"nextpage,omitempty"`

	// Norified is the time when the notification was sent
	Notified string `json:"notified" cql:"notified,omitempty"`

	// Old is the old layers that introduced vulnerability
	Old interface{} `json:"old" cql:"old,omitempty"`

	// Page is the page number
	Page string `json:"page" cql:"page,omitempty"`
}

// NewNotification returns a new *Notification
func NewNotification() *Notification {

	return &Notification{}
}

// Identity returns the Identity of the object.
func (o *Notification) Identity() elemental.Identity {

	return NotificationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Notification) Identifier() string {

	return o.ID
}

func (o *Notification) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Notification) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *Notification) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Notification) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NotificationAttributesMap[name]
}

// NotificationAttributesMap represents the map of attribute for Notification.
var NotificationAttributesMap = map[string]elemental.AttributeSpecification{
	"Created": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "Created",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "deleted",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Limit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "limit",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"New": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "new",
		Orderable:      true,
		Stored:         true,
		SubType:        "VulnerabilityWithLayers",
		Type:           "object",
	},
	"NextPage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "nextPage",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Notified": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "notified",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Old": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Name:           "old",
		Orderable:      true,
		Stored:         true,
		SubType:        "VulnerabilityWithLayers",
		Type:           "object",
	},
	"Page": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "page",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
