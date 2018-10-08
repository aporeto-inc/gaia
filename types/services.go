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

	_, _, err := p.ValidateWithoutOverlap(map[uint8]*portutils.PortsList{}, map[uint8]*portutils.PortsRangeList{})
	return err
}

// ValidateWithoutOverlap validates a list of processing unit services has no overlap with any given parameter.
func (p ProcessingUnitServicesList) ValidateWithoutOverlap(cachePortsList map[uint8]*portutils.PortsList, cacheRanges map[uint8]*portutils.PortsRangeList) (map[uint8]*portutils.PortsList, map[uint8]*portutils.PortsRangeList, error) {

	for _, pu := range p {

		var cpl *portutils.PortsList
		var cpr *portutils.PortsRangeList
		var ok bool

		if cpl, ok = cachePortsList[pu.Protocol]; !ok {
			cpl = &portutils.PortsList{}
			cachePortsList[pu.Protocol] = cpl
		}

		if cpr, ok = cacheRanges[pu.Protocol]; !ok {
			cpr = &portutils.PortsRangeList{}
			cacheRanges[pu.Protocol] = cpr
		}

		ports := pu.Ports

		// Range port
		if strings.Contains(ports, ":") {

			pr, err := portutils.ConvertToPortsRange(ports)
			if err != nil {
				return nil, nil, err
			}

			if pr.HasOverlapWithPortsRanges(cpr) {
				return nil, nil, fmt.Errorf("Port range overlaps with another range")
			}

			if pr.HasOverlapWithPortsList(cpl) {
				return nil, nil, fmt.Errorf("Port range overlaps with another port")
			}

			*cpr = append(*cpr, pr)
			cacheRanges[pu.Protocol] = cpr

			continue
		}

		// Single & Multiple ports
		pl, err := portutils.ConvertToPortsList(ports)
		if err != nil {
			return nil, nil, err
		}

		if pl.HasOverlapWithPortsList(cpl) {
			return nil, nil, fmt.Errorf("Port overlaps with another port")
		}

		if pl.HasOverlapWithPortsRanges(cpr) {
			return nil, nil, fmt.Errorf("Port overlaps with another port range")
		}

		*cpl = append(*cpl, *pl...)
		cachePortsList[pu.Protocol] = cpl
	}

	return cachePortsList, cacheRanges, nil
}
