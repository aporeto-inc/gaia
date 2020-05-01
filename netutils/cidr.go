package netutils

import (
	"fmt"
	"net"
	"strings"
)

// ipMaskContains is used to check if a mask's range contains another mask's range
func ipMaskContains(mask1 net.IPMask, mask2 net.IPMask) bool {
	ones1, size1 := mask1.Size()
	ones2, size2 := mask2.Size()
	if size1 != size2 {
		return false // Dont compare IPv4 and IPv6
	}
	return ones1 <= ones2
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

// parseCIDRs converts a list of string to list of net.IPNet in v4 and v6. Returns an error if it wasnt able to parse a CIDR
func parseCIDRs(cidrs []string) (ipNetsV4 []*net.IPNet, ipNetsV6 []*net.IPNet, _ error) {

	for _, s := range cidrs {
		ip, network, err := net.ParseCIDR(s)
		if err != nil {
			return ipNetsV4, ipNetsV6, fmt.Errorf("%s is not a valid CIDR", s)
		}
		if ipv4 := ip.To4(); ipv4 != nil {
			ipNetsV4 = append(ipNetsV4, network)
		} else if ipv6 := ip.To16(); ipv6 != nil {
			ipNetsV6 = append(ipNetsV6, network)
		}
	}
	return ipNetsV4, ipNetsV6, nil
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
	prefixesV4, prefixesV6, err := parseCIDRs(regular)
	if err != nil {
		return err
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, s := range cidrs {
		if !strings.HasPrefix(s, "!") {
			continue
		}
		c := strings.TrimPrefix(s, "!")
		ip, network, err := net.ParseCIDR(c)
		if err != nil {
			return fmt.Errorf("%s is not a valid CIDR", c)
		}
		if ipv4 := ip.To4(); ipv4 != nil {
			if !prefixIsContained(prefixesV4, network) {
				return fmt.Errorf("%s is not contained in any CIDR", s)
			}
		} else if ipv6 := ip.To16(); ipv6 != nil {
			if !prefixIsContained(prefixesV6, network) {
				return fmt.Errorf("%s is not contained in any CIDR", s)
			}
		}
	}

	return nil
}
