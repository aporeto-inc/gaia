package types

import (
	"fmt"
	"regexp"
)

// ProcessingUnitService is a network service that the processing unit listens to.
type ProcessingUnitService struct {
	Protocol uint8
	Ports    string
}

// ProcessingUnitServicesList is a list of ProcessingUnitServices
type ProcessingUnitServicesList []*ProcessingUnitService

// Validate will validate the types in the ProcessingUnitService. Uses the same
// regular expression as the main API in external services. Do not touch unless
// you know what you are doing.
func (p ProcessingUnitServicesList) Validate() error {
	reg := regexp.MustCompile("^([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535)(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))?$")
	for _, pu := range p {
		if !reg.Match([]byte(pu.Ports)) {
			return fmt.Errorf("Invalid port or port pair: %s", pu.Ports)
		}
	}
	return nil
}
