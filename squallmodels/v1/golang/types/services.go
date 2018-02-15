package types

import (
	"fmt"
	"regexp"
)

var (
	reg = regexp.MustCompile("^([0-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$")
)

// ProcessingUnitService is a network service that the processing unit listens to.
type ProcessingUnitService struct {
	Protocol uint8  `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`
	Ports    string `json:"ports" bson:"ports" mapstructure:"ports,omitempty"`
}

// ProcessingUnitServicesList is a list of ProcessingUnitServices
type ProcessingUnitServicesList []*ProcessingUnitService

// Validate will validate the types in the ProcessingUnitService. Uses the same
// regular expression as the main API in external services. Do not touch unless
// you know what you are doing.
func (p ProcessingUnitServicesList) Validate() error {
	for _, pu := range p {
		if !reg.Match([]byte(pu.Ports)) {
			return fmt.Errorf("Invalid port or port pair: %s", pu.Ports)
		}
	}
	return nil
}
