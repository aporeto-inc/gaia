package types

import (
	"fmt"
	"strings"

	"go.aporeto.io/gaia/portutils"
)

// ProcessingUnitService is a network service that the processing unit listens to.
type ProcessingUnitService struct {
	Protocol uint8  `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`
	Ports    string `json:"ports" bson:"ports" mapstructure:"ports,omitempty"`
}

// ProcessingUnitServicesList is a list of ProcessingUnitServices
type ProcessingUnitServicesList []*ProcessingUnitService

// Validate validates a list of processing unit services.
func (p ProcessingUnitServicesList) Validate() error {

	_, _, err := p.ValidateWithoutOverlap(&portutils.PortsList{}, []*portutils.PortsRange{})
	return err
}

// ValidateWithoutOverlap validates a list of processing unit services has no overlap with any given parameter.
func (p ProcessingUnitServicesList) ValidateWithoutOverlap(cachePortsList *portutils.PortsList, cacheRanges []*portutils.PortsRange) (*portutils.PortsList, []*portutils.PortsRange, error) {

	for _, pu := range p {

		ports := pu.Ports

		// Range port
		if strings.Contains(ports, ":") {

			pr, err := portutils.ConvertToPortsRange(ports)
			if err != nil {
				return nil, nil, err
			}

			if pr.HasOverlapWithPortsRanges(cacheRanges) {
				return nil, nil, fmt.Errorf("Port range overlaps with another range")
			}

			if pr.HasOverlapWithPortsList(cachePortsList) {
				return nil, nil, fmt.Errorf("Port range overlaps with another port")
			}

			cacheRanges = append(cacheRanges, pr)
			for _, c := range cacheRanges {
				fmt.Println(c.FromPort, c.ToPort)
			}

			continue
		}

		// Single & Multiple ports
		pl, err := portutils.ConvertToPortsList(ports)
		if err != nil {
			return nil, nil, err
		}

		if pl.HasOverlapWithPortsList(cachePortsList) {
			return nil, nil, fmt.Errorf("Port overlaps with another port")
		}

		if pl.HasOverlapWithPortsRanges(cacheRanges) {
			return nil, nil, fmt.Errorf("Port overlaps with another port range")
		}

		*cachePortsList = append(*cachePortsList, *pl...)
	}

	return cachePortsList, cacheRanges, nil
}
