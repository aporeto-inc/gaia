package netutils

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

var (
	opInclude = "include"
	opExclude = "exclude"
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
		if !pc.ipNet.Contains(c.ipNet.IP) {
			continue
		}
		ones1, size1 := pc.ipNet.Mask.Size()
		ones2, size2 := c.ipNet.Mask.Size()
		if size1 != size2 {
			continue
		}
		if ones1 < ones2 {
			return true
		}
	}

	return false
}

// hasDuplicates is used to check if a list of cidr already has the same subnet of given cidr
func hasDuplicates(cidrs []*cidr, c *cidr) bool {

	for _, pc := range cidrs {
		if net.IP.Equal(pc.ipNet.IP, c.ipNet.IP) && bytes.Equal(pc.ipNet.Mask, c.ipNet.Mask) {
			return true
		}
	}

	return false
}

// parseCIDRs converts a list of string to list of cidr. Returns an error if it wasnt able to parse a CIDR
func parseCIDRs(addresses []string) ([]*cidr, error) {

	cidrs := []*cidr{}
	for _, s := range addresses {
		// parse address string into cidr
		c := &cidr{op: opInclude, str: s}
		if strings.HasPrefix(s, "!") {
			c.op = opExclude
			c.str = strings.TrimPrefix(s, "!")
		}
		_, network, err := net.ParseCIDR(c.str)
		if err != nil {
			return nil, fmt.Errorf("%s is not a valid CIDR", s)
		}
		c.ipNet = network

		// validate subnet uniqness
		if hasDuplicates(cidrs, c) {
			return nil, fmt.Errorf("CIDR subnet converted from %s is duplicated", c.str)
		}

		cidrs = append(cidrs, c)
	}
	return cidrs, nil
}

// ValidateCIDRs validates that the list of string provided as a set is valid CIDR set
func ValidateCIDRs(addresses []string) error {

	cidrs, err := parseCIDRs(addresses)
	if err != nil {
		return err
	}

	// Parse and validate all not CIDRs are included in regular CIDRs
	for _, c := range cidrs {
		if c.op == opInclude {
			continue
		}
		if !prefixIsContained(cidrs, c) {
			return fmt.Errorf("prefix %s is not contained in any CIDR", c.str)
		}
	}

	return nil
}
