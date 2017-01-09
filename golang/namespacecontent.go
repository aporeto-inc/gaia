package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// NamespaceContentIdentity represents the Identity of the object
var NamespaceContentIdentity = elemental.Identity{
	Name:     "namespacecontent",
	Category: "namespacecontents",
}

// NamespaceContentsList represents a list of NamespaceContents
type NamespaceContentsList []*NamespaceContent

// NamespaceContent represents the model of a namespacecontent
type NamespaceContent struct {
	// ID of the content
	ContentID string `json:"contentID" cql:"contentid,primarykey,omitempty" bson:"_contentid"`

	// Type of the content
	ContentType string `json:"contentType" cql:"contenttype,primarykey,omitempty" bson:"_contenttype"`

	// name of the namespace
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`
}

// NewNamespaceContent returns a new *NamespaceContent
func NewNamespaceContent() *NamespaceContent {

	return &NamespaceContent{}
}

// Identity returns the Identity of the object.
func (o *NamespaceContent) Identity() elemental.Identity {

	return NamespaceContentIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceContent) Identifier() string {

	return ""
}

func (o *NamespaceContent) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceContent) SetIdentifier(ID string) {

}

// Validate valides the current information stored into the structure.
func (o *NamespaceContent) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("contentID", o.ContentID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("contentType", o.ContentType); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (NamespaceContent) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NamespaceContentAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (NamespaceContent) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceContentAttributesMap
}

// NamespaceContentAttributesMap represents the map of attribute for NamespaceContent.
var NamespaceContentAttributesMap = map[string]elemental.AttributeSpecification{
	"ContentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "contentID",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ContentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "contentType",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
