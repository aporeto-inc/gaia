package netutils

import (
	"fmt"
	"net"
	"strings"
)

var (
	opInclude = "include" // default operation to include cidr in networks
	opExclude = "exclude" // not include cidr when cidr string has the ! prefix
)

// cidr represents the cidr network to be included or excluded
type cidr struct {
	// op is the opreation to perform with the cidr: include or exclude
	op string

	// ipNet is the IPNet object that contains ip and mask of the cidr
	ipNet *net.IPNet

	// str is the original cidr string
	str string
}

// prefixIsContained is used to check if an cidr is contained in a list of CIDRs
func prefixIsContained(cidrs []*cidr, c *cidr) bool {

	for _, pc := range cidrs {
		if pc.op == opExclude || !pc.ipNet.Contains(c.ipNet.IP) {
			continue
		}
		ones1, size1 := pc.ipNet.Mask.Size()
		ones2, size2 := c.ipNet.Mask.Size()
		if size1 != size2 {
			continue
		}
		if ones1 <= ones2 {
			return true
		}
	}

	return false
}

// hasDuplicates is used to check if a list of cidr has more than one network equals to the given network
func hasDuplicates(cidrs []*cidr, net *net.IPNet) bool {

	var count int
	for _, pc := range cidrs {
		if pc.ipNet.String() == net.String() {
			count++
		}
	}

	return count > 1
}

// parseCIDR converts the given string to cidr. Returns an error if it wasnt able to parse a CIDR
func parseCIDR(s string) (*cidr, error) {

	c := &cidr{op: opInclude, str: s}
	_, network, err := net.ParseCIDR(strings.TrimPrefix(s, "!"))
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid CIDR", s)
	}
	c.ipNet = network
	if strings.HasPrefix(s, "!") {
		c.op = opExclude
	}

	return c, nil
}

// ValidateCIDRs validates that the list of string provided as a set is a valid CIDR set
func ValidateCIDRs(ss []string) error {

	cidrs := []*cidr{}
	for _, s := range ss {
		cidr, err := parseCIDR(s)
		if err != nil {
			return err
		}
		cidrs = append(cidrs, cidr)
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrs {
		if hasDuplicates(cidrs, c.ipNet) {
			return fmt.Errorf("CIDR subnet parsed from %s is duplicated", c.str)
		}
		if c.op == opExclude && !prefixIsContained(cidrs, c) {
			return fmt.Errorf("%s is not contained in any CIDR", c.str)
		}
	}

	return nil
}
