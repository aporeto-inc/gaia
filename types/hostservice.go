package types

import (
	"fmt"
	"regexp"
)

var regHostServiceName = regexp.MustCompile(`^[a-zA-Z0-9_:.$%/-]{0,64}$`)

// HostService is a service associated with an enforcer profile for a host.
type HostService struct {
	Name        string                     `json:"name" bson:"name" mapstructure:"name,omitempty"`
	NetworkOnly bool                       `json:"networkonly" bson:"networkonly" mapstructure:"networkonly,omitempty"`
	Services    ProcessingUnitServicesList `json:"services" bson:"services" mapstructure:"services,omitempty"`
}

// HostServicesList is a list of ProcessingUnitServices
type HostServicesList []*HostService

// Validate will validate the types in the ProcessingUnitService. Uses the same
// regular expression as the main API in external services. Do not touch unless
// you know what you are doing.
func (h HostServicesList) Validate() error {
	for _, s := range h {
		if len(s.Name) == 0 {
			return fmt.Errorf("Host service names must be specified")
		}

		if !regHostServiceName.MatchString(s.Name) {
			return fmt.Errorf(`Host service name must be less than 64 characters and contains only alphanumeric and the following characters: _, :, ., %%, / and -`)
		}

		if s.Services != nil {
			if err := s.Services.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}
