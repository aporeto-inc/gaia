
package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"


// TagCounterIdentity represents the Identity of the object
var TagCounterIdentity = elemental.Identity {
    Name:     "tagcounter",
    Category: "tagcounters",
}

// TagCountersList represents a list of TagCounters
type TagCountersList []*TagCounter

// TagCounter represents the model of a tagcounter
type TagCounter struct {
    // Count represents the number of time the represented tag is used.
    Count int `json:"-" cql:"count,omitempty"`

    // Namespace represents the namespace of the counted tag.
    Namespace string `json:"-" cql:"namespace,omitempty"`

    // Value represents the value of the counted tag.
    Value string `json:"-" cql:"value,omitempty"`

    
}

// NewTagCounter returns a new *TagCounter
func NewTagCounter() *TagCounter {

    return &TagCounter{
        }
}

// Identity returns the Identity of the object.
func (o *TagCounter) Identity() elemental.Identity {

    return TagCounterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *TagCounter) Identifier() string {

    return o.
}

func  (o *TagCounter) String() string {

  return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *TagCounter) SetIdentifier(ID string) {

    o. = ID
}

// Validate valides the current information stored into the structure.
func (o *TagCounter) Validate() elemental.Errors {

    errors := elemental.Errors{}

    if len(errors) > 0 {
        return errors
    }

    return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o TagCounter) SpecificationForAttribute(name string) elemental.AttributeSpecification {

  return TagCounterAttributesMap[name]
}

// TagCounterAttributesMap represents the map of attribute for TagCounter.
var TagCounterAttributesMap = map[string]elemental.AttributeSpecification{
  "Count": elemental.AttributeSpecification{
      AllowedChoices: []string{},
      Name: "count",
      Stored: true,
      Type: "integer",
      },
  "Namespace": elemental.AttributeSpecification{
      AllowedChoices: []string{},
      Format: "free",
      Name: "namespace",
      Stored: true,
      Type: "string",
      },
  "Value": elemental.AttributeSpecification{
      AllowedChoices: []string{},
      Format: "free",
      Name: "value",
      Stored: true,
      Type: "string",
      },
  }