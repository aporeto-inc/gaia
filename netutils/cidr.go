package netutils

import (
	"fmt"
	"net"
	"strings"
)

// ipMaskContains is used to check if a mask's range contains another mask's range
func ipMaskContains(mask1 net.IPMask, mask2 net.IPMask) bool {
	for i := 0; i < 4; i++ {
		if mask1[i] > mask2[i] {
			return false
		}
	}

	return true
}

// prefixIsContained is used to check if a network is included in a list of CIDRs
func prefixIsContained(cidrs []*net.IPNet, network *net.IPNet) bool {

	for _, c := range cidrs {
		if c.Contains(network.IP) && ipMaskContains(c.Mask, network.Mask) {
			return true
		}
	}

	return false
}

// parseCIDRs converts a list of string to list of net.IPNet. Returns an error if it wasnt able to parse a CIDR
func parseCIDRs(cidrs []string) ([]*net.IPNet, error) {

	prefixes := []*net.IPNet{}
	for _, s := range cidrs {
		_, network, err := net.ParseCIDR(s)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid CIDR", s)
		}
		prefixes = append(prefixes, network)
	}
	return prefixes, nil
}

// ValidateCIDRs validates that the CIDRs provided as a set is valid
func ValidateCIDRs(cidrs []string) error {

	regular := []string{}
	for _, s := range cidrs {
		if strings.HasPrefix(s, "!") {
			continue
		}
		regular = append(regular, s)
	}

	// Get regular prefixes
	prefixes, err := parseCIDRs(regular)
	if err != nil {
		return err
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, s := range cidrs {
		if !strings.HasPrefix(s, "!") {
			continue
		}
		c := strings.TrimPrefix(s, "!")
		_, network, err := net.ParseCIDR(c)
		if err != nil {
			return fmt.Errorf("%s is not a valid CIDR", c)
		}
		if !prefixIsContained(prefixes, network) {
			return fmt.Errorf("%s is not contained in any CIDR", s)
		}
	}

	return nil
}
