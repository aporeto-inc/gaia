package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Comment represents the model of a comment
type Comment struct {
	// The claims of the author.
	Claims []string `json:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// The content of the comment.
	Content string `json:"content" bson:"content" mapstructure:"content,omitempty"`

	// The date of the comment.
	Date time.Time `json:"date" bson:"date" mapstructure:"date,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewComment returns a new *Comment
func NewComment() *Comment {

	return &Comment{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
		Claims:       []string{},
	}
}

// DeepCopy returns a deep copy if the Comment.
func (o *Comment) DeepCopy() *Comment {

	if o == nil {
		return nil
	}

	out := &Comment{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Comment.
func (o *Comment) DeepCopyInto(out *Comment) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Comment: %s", err))
	}

	*out = *target.(*Comment)
}

// Validate valides the current information stored into the structure.
func (o *Comment) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
