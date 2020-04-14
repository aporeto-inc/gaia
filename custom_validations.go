package gaia

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/constants"
	"go.aporeto.io/gaia/netutils"
	"go.aporeto.io/gaia/portutils"
	"go.aporeto.io/gaia/protocols"
	"gopkg.in/yaml.v2"
)

// ValidateAPIProxyEntity validates an APIProxy.
func ValidateAPIProxyEntity(apiProxy *APIProxy) error {

	var errs elemental.Errors

	// We only want to check if there is a key on creation as it is a secret which
	// means it will never be exposed outside of the service
	if apiProxy.ID == "" && apiProxy.ClientCertificate != "" && apiProxy.ClientCertificateKey == "" {
		errs = errs.Append(makeValidationError("ClientCertificateKey", "client certificate private key was not provided"))
	}

	if apiProxy.ClientCertificate == "" && apiProxy.ClientCertificateKey != "" {
		errs = errs.Append(makeValidationError("ClientCertificate", "client certificate was not provided"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// ValidatePortString validates a string represents a port or a range of port.
// valid: 443, 443:555
func ValidatePortString(attribute string, portExp string) error {

	// TODO: Use portutils to validate a port
	ports := strings.Split(portExp, ":")
	if len(ports) == 0 || len(ports) > 2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	p1, err := strconv.Atoi(ports[0])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p1 < 1 || p1 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be between 1 and 65535", attribute))
	}

	if len(ports) == 1 {
		return nil
	}

	p2, err := strconv.Atoi(ports[1])
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a port (xxx) or port range (xxx:yyy)", attribute))
	}

	if p2 < 1 || p2 > 65535 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be between 1 and 65535", attribute))
	}

	if p1 >= p2 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' left bound is greater or equal to left bound ", attribute))
	}

	return nil
}

// ValidatePortStringList validates a list of ports.
func ValidatePortStringList(attribute string, ports []string) error {

	for _, port := range ports {
		if err := ValidatePortString(attribute, port); err != nil {
			return err
		}
	}

	return nil
}

var rxDNSName = regexp.MustCompile(`^(\*\.){0,1}([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`)

// ValidateNetwork validates a CIDR.
func ValidateNetwork(attribute string, network string) error {

	if _, _, err := net.ParseCIDR(network); err == nil {
		return nil
	}

	if rxDNSName.MatchString(network) {
		return nil
	}

	return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a CIDR or hostname", attribute))
}

// ValidateNetworkList validates a list of networks.
// The list cannot be empty
func ValidateNetworkList(attribute string, networks []string) error {

	if len(networks) == 0 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must not be empty", attribute))
	}

	return ValidateOptionalNetworkList(attribute, networks)
}

// ValidateOptionalNetworkList validates a list of networks.
// It can be empty/
func ValidateOptionalNetworkList(attribute string, networks []string) error {

	for _, network := range networks {
		if err := ValidateNetwork(attribute, network); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProtocol validates a string represents netwotk a protocol.
func ValidateProtocol(attribute string, proto string) error {

	upperProto := strings.ToUpper(proto)
	if protocols.L4ProtocolNumberFromName(upperProto) != -1 {
		return nil
	}

	p, err := strconv.Atoi(proto)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' valid protocol name or number", attribute))
	}

	if p < 0 || p > 255 {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' protocol number must be between 0 and 255 included", attribute))
	}

	return nil
}

// ValidateProtocolList validates a list of protocols.
func ValidateProtocolList(attribute string, protocols []string) error {

	for _, proto := range protocols {
		if err := ValidateProtocol(attribute, proto); err != nil {
			return err
		}
	}

	return nil
}

// ValidateServiceEntity validates a Service.
func ValidateServiceEntity(service *Service) error {

	var errs elemental.Errors

	switch service.AuthorizationType {

	case ServiceAuthorizationTypeOIDC:

		if service.OIDCProviderURL == "" {
			errs = errs.Append(makeValidationError("OIDCProviderURL", "`OIDCProviderURL` is required when `authorizationType` is set to `OIDC`"))
		}

		if u, err := url.Parse(service.OIDCProviderURL); err != nil || u == nil || u.Scheme != "https" {
			errs = errs.Append(makeValidationError("OIDCProviderURL", "`OIDCProviderURL` must be a valid HTTPS URL: example https://xxx.yyy"))
		}

		if service.OIDCClientID == "" {
			errs = errs.Append(makeValidationError("OIDCClientID", "`OIDCClientID` is required when `authorizationType` is set to `OIDC`"))
		}

		if service.OIDCClientSecret == "" {
			errs = errs.Append(makeValidationError("OIDCClientSecret", "`OIDCClientSecret` is required when `authorizationType` is set to `OIDC`"))
		}

	case ServiceAuthorizationTypeJWT:

		if service.JWTSigningCertificate == "" {
			errs = errs.Append(makeValidationError("JWTSigningCertificate", "`JWTSigningCertificate` is required when `authorizationType` is set to `JWT`"))
		}
	}

	if service.TLSType == ServiceTLSTypeExternal {

		if service.TLSCertificate == "" {
			errs = errs.Append(makeValidationError("TLSCertificate", "`TLSCertificate` is required when `TLSType` is set to `External`"))
		}

		if service.TLSCertificateKey == "" {
			errs = errs.Append(makeValidationError("TLSCertificateKey", "`TLSCertificateKey` is required when `TLSType` is set to `External`"))
		}
	}

	allSubnets := []*net.IPNet{}
	for _, ip := range service.IPs {
		ipNet, err := ipNetFromString(ip)
		if err != nil {
			errs = errs.Append(err)
			continue
		}
		for j := 0; j < len(allSubnets); j++ {
			if allSubnets[j].Contains(ipNet.IP) || ipNet.Contains(allSubnets[j].IP) {
				errs = errs.Append(makeValidationError("IPs", "subnets cannot overlap"))
			}
		}
		allSubnets = append(allSubnets, ipNet)
	}

	allHosts := map[string]bool{}
	for _, name := range service.Hosts {
		if !isFQDN(name) {
			errs = errs.Append(makeValidationError("hosts", "`Hosts` must be a valid hostname or FQDN, compliant with RF952"))
		}
		if _, ok := allHosts[name]; ok {
			errs = errs.Append(makeValidationError("hosts", "`Hosts` must be unique"))
		}
		allHosts[name] = true
	}

	if len(service.Hosts) == 0 && len(service.IPs) == 0 {
		errs = errs.Append(makeValidationError("", "You must set at least one value in `hosts` or `IPs`"))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func makeValidationError(attribute string, message string) elemental.Error {

	err := elemental.NewError(
		"Validation Error",
		message,
		"gaia",
		http.StatusUnprocessableEntity,
	)

	if attribute != "" {
		err.Data = map[string]interface{}{"attribute": attribute}
	}

	return err
}

// ValidateEnforcerProfile validates a an enforcer profile.
func ValidateEnforcerProfile(enforcerProfile *EnforcerProfile) error {

	if err := netutils.ValidateCIDRs(enforcerProfile.TargetNetworks); err != nil {
		return makeValidationError("targetNetworks", err.Error())
	}

	if err := netutils.ValidateCIDRs(enforcerProfile.TargetUDPNetworks); err != nil {
		return makeValidationError("targetUDPNetworks", err.Error())
	}

	if err := netutils.ValidateCIDRs(enforcerProfile.ExcludedNetworks); err != nil {
		return makeValidationError("excludedNetworks", err.Error())
	}

	// Validate trusted CAs
	for i, ca := range enforcerProfile.TrustedCAs {
		rest := []byte(ca)
		var block *pem.Block

		for {
			block, rest = pem.Decode(rest)

			if block == nil || block.Type != "CERTIFICATE" {
				return makeValidationError("trustedCAs", fmt.Sprintf("The CA %d is not a valid CA", i))
			}

			if _, err := x509.ParseCertificate(block.Bytes); err != nil {
				return makeValidationError("trustedCAs", err.Error())
			}

			if len(rest) == 0 {
				break
			}
		}
	}

	return nil
}

// ValidateProcessingUnitPolicy validates a processing unit policy has no action and datapath set to default.
func ValidateProcessingUnitPolicy(policy *ProcessingUnitPolicy) error {

	if policy.Action == ProcessingUnitPolicyActionDefault && policy.DatapathType == ProcessingUnitPolicyDatapathTypeDefault {
		if len(policy.IsolationProfileSelector) == 0 {
			return makeValidationError("datapathType", fmt.Sprintf("Both datapath and action cannot be set to default"))
		}
		return makeValidationError("action", fmt.Sprintf("Both datapath and action cannot be set to default"))
	}

	return nil
}

// ValidateProcessingUnitServicesList validates a list of processing unit services.
func ValidateProcessingUnitServicesList(attribute string, svcs []*ProcessingUnitService) error {

	if _, _, err := ValidateProcessingUnitServicesListWithoutOverlap(svcs, map[int]*portutils.PortsList{}, map[int]*portutils.PortsRangeList{}); err != nil {
		return makeValidationError(attribute, err.Error())
	}
	return nil
}

// ValidateProcessingUnitServicesListWithoutOverlap validates a list of processing unit services has no overlap with any given parameter.
func ValidateProcessingUnitServicesListWithoutOverlap(svcs []*ProcessingUnitService, cachePortsList map[int]*portutils.PortsList, cacheRanges map[int]*portutils.PortsRangeList) (map[int]*portutils.PortsList, map[int]*portutils.PortsRangeList, error) {

	for i, svc := range svcs {

		if svc == nil {
			return nil, nil, fmt.Errorf("nil processingunitservice in list at index %d", i)
		}

		var cpl *portutils.PortsList
		var cpr *portutils.PortsRangeList
		var ok bool

		if cpl, ok = cachePortsList[svc.Protocol]; !ok {
			cpl = &portutils.PortsList{}
			cachePortsList[svc.Protocol] = cpl
		}

		if cpr, ok = cacheRanges[svc.Protocol]; !ok {
			cpr = &portutils.PortsRangeList{}
			cacheRanges[svc.Protocol] = cpr
		}

		targetPorts := svc.TargetPorts

		for _, ports := range targetPorts {
			// Range port
			if strings.Contains(ports, ":") {

				pr, err := portutils.ConvertToPortsRange(ports)
				if err != nil {
					return nil, nil, err
				}

				if pr.HasOverlapWithPortsRanges(cpr) {
					return nil, nil, fmt.Errorf("port range overlaps with another range")
				}

				if pr.HasOverlapWithPortsList(cpl) {
					return nil, nil, fmt.Errorf("port range overlaps with another port")
				}

				*cpr = append(*cpr, pr)
				cacheRanges[svc.Protocol] = cpr

				continue
			}

			// Single & Multiple ports
			pl, err := portutils.ConvertToPortsList(ports)
			if err != nil {
				return nil, nil, err
			}

			if pl.HasOverlapWithPortsList(cpl) {
				return nil, nil, fmt.Errorf("port overlaps with another port")
			}

			if pl.HasOverlapWithPortsRanges(cpr) {
				return nil, nil, fmt.Errorf("port overlaps with another port range")
			}

			*cpl = append(*cpl, *pl...)
			cachePortsList[svc.Protocol] = cpl
		}

	}

	return cachePortsList, cacheRanges, nil
}

// ValidateTimeDuration validates that the time duration provided is compliant
// with the go format.
func ValidateTimeDuration(attribute string, duration string) error {
	_, err := time.ParseDuration(duration)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be valid duration (examaple: 1h or 30s)", attribute))
	}
	return nil
}

// hostname regex from github.com/go-playground/validator
var hostnameRegexRFC1123 = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-\.]+[a-z-Az0-9]$`)

func isFQDN(val string) bool {

	if val == "" {
		return false
	}

	if val[len(val)-1] == '.' {
		val = val[0 : len(val)-1]
	}

	return hostnameRegexRFC1123.MatchString(val)
}

func ipNetFromString(ip string) (*net.IPNet, error) {
	_, ipNet, err := net.ParseCIDR(ip)
	if err != nil {
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			return nil, makeValidationError("IPs", "`IPs` must be a list of valid IPv4 address or CIDR notation")
		}
		ipNet = &net.IPNet{IP: parsedIP, Mask: []byte{0xf, 0xf, 0xf, 0xf}}
	}
	return ipNet, nil
}

// ValidateHTTPMethods validates the attribute methods is a list of HTTP verbs.
func ValidateHTTPMethods(attribute string, methods []string) error {

	for _, m := range methods {
		mu := strings.ToUpper(m)

		if mu != http.MethodPost &&
			mu != http.MethodGet &&
			mu != http.MethodDelete &&
			mu != http.MethodPut &&
			mu != http.MethodHead &&
			mu != http.MethodPatch {

			return makeValidationError(attribute, fmt.Sprintf("invalid HTTP method %s", m))
		}
	}

	return nil
}

// ValidateHTTPSURL validates the URL to make sure it is in a validate format and is https.
func ValidateHTTPSURL(attribute string, address string) error {

	u, err := url.Parse(address)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Attribute '%s' must be a valid HTTPS URL (example: https://aporeto.com/)", attribute))
	}

	if !strings.EqualFold(u.Scheme, "https") || u.Host == "" {
		return makeValidationError(attribute, fmt.Sprintf("Invalid HTTPS URL %s", address))
	}

	return nil
}

// ValidateAutomation validates an automation by checking for the following:
//   - Exactly ONE action MUST be defined if the automation trigger type is set to "Webhook"
func ValidateAutomation(auto *Automation) error {
	switch auto.Trigger {
	case AutomationTriggerWebhook:
		switch len(auto.Actions) {
		case 1:
		case 0:
			return makeValidationError("actions", fmt.Sprintf("Exactly one action must be defined if trigger type is set to %q", AutomationTriggerWebhook))
		default:
			return makeValidationError("actions", fmt.Sprintf("Only one action can be defined if trigger type is set to %q", AutomationTriggerWebhook))
		}
	}

	return nil
}

var regHostServiceName = regexp.MustCompile(`^[a-zA-Z0-9_]{0,127}$`)

// ValidateHostServices validates a host service. Applies to the new API only.
func ValidateHostServices(hs *HostService) error {

	// Constraint on regex is used because the enforcer is using the name as nativeContextID.
	if !regHostServiceName.MatchString(hs.Name) {
		return makeValidationError("name", "Host service name must be less than 128 characters and contains only alphanumeric or _")
	}

	if !hs.HostModeEnabled && len(hs.Services) == 0 {
		return makeValidationError("services", "Host service must have either HostModeEnabled or must declare services")
	}

	if err := ValidateHostServicesNonOverlapPorts(hs.Services); err != nil {
		return makeValidationError("services", err.Error())
	}

	return nil
}

// ValidateProtoPorts validates a list of protocol/ports.
func ValidateProtoPorts(attribute string, services []string) error {

	for _, service := range services {
		if err := ValidateProtoPort(attribute, service); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProtoPort validates protocol/port.
func ValidateProtoPort(attribute string, service string) error {

	portSubString, _, err := portutils.ExtractPortsAndProtocol(service)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("invalid format: %s", err))
	}

	_, err = portutils.ConvertToSinglePort(portSubString)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("invalid port: %s", err))
	}

	return nil
}

// ValidateHostServicesNonOverlapPorts validates a list of processing unit services has no overlap with any given parameter.
func ValidateHostServicesNonOverlapPorts(svcs []string) error {

	udpPorts := portutils.PortsRangeList{}
	tcpPorts := portutils.PortsRangeList{}

	var pr *portutils.PortsRange
	var protocol string
	var err error

	for _, svc := range svcs {

		pr, protocol, err = portutils.ExtractPortsAndProtocolFromHostService(svc)
		if err != nil {
			return err
		}

		switch protocol {
		case protocols.L4ProtocolTCP:
			if pr.HasOverlapWithPortsRanges(&tcpPorts) {
				return fmt.Errorf("host service cannot have overlapping TCP ports")
			}
			tcpPorts = append(tcpPorts, pr)
		case protocols.L4ProtocolUDP:
			if pr.HasOverlapWithPortsRanges(&udpPorts) {
				return fmt.Errorf("host service cannot have overlapping UDP ports")
			}
			udpPorts = append(udpPorts, pr)
		default:
			return fmt.Errorf("host service has invalid format: %s", protocol)
		}

	}

	return nil
}

// ValidateServicePorts validates a list of serviceports.
func ValidateServicePorts(attribute string, servicePorts []string) error {

	for _, servicePort := range servicePorts {

		if strings.EqualFold(servicePort, protocols.ANY) {
			if len(servicePorts) != 1 {
				return makeValidationError(attribute, fmt.Sprintf("Protocol '%s' cannot be used with other protocols", servicePort))
			}
			continue
		}

		if err := ValidateServicePort(attribute, servicePort); err != nil {
			return err
		}
	}

	return nil
}

// ValidateServicePort validates a single serviceport.
func ValidateServicePort(attribute string, servicePort string) error {

	if strings.EqualFold(servicePort, protocols.ANY) {
		return nil
	}

	parts := strings.SplitN(servicePort, "/", 2)
	upperProto := strings.ToUpper(parts[0])
	if protocols.L4ProtocolNumberFromName(upperProto) == -1 {
		return makeValidationError(attribute, fmt.Sprintf("'%s' is not a valid protocol", upperProto))
	}

	if len(parts) == 1 {
		if upperProto == protocols.L4ProtocolTCP || upperProto == protocols.L4ProtocolUDP {
			return makeValidationError(attribute, fmt.Sprintf("protocol '%s' cannot be used without ports", upperProto))
		}
		return nil
	}

	if upperProto != protocols.L4ProtocolTCP && upperProto != protocols.L4ProtocolUDP {
		return makeValidationError(attribute, fmt.Sprintf("protocol '%s' cannot be used with ports", upperProto))
	}

	ports := parts[1]
	return ValidatePortString(attribute, ports)
}

// ValidateAudience validates an audience string.
func ValidateAudience(attribute string, audience string) error {
	// TODO: not liking the idea of importing addedeffect here
	return nil
}

// ValidatePEM validates a string contains a PEM.
func ValidatePEM(attribute string, pemdata string) error {

	if pemdata == "" {
		return nil
	}

	var i int
	var block *pem.Block
	rest := []byte(pemdata)

	for {
		block, rest = pem.Decode(rest)

		if block == nil {
			return makeValidationError(attribute, fmt.Sprintf("Unable to decode PEM number %d", i))
		}

		if len(rest) == 0 {
			return nil
		}
		i++
	}
}

// ValidateCA verifies a string contains a PEM encoding a CA.
func ValidateCA(attribute string, pemdata string) error {

	if pemdata == "" {
		return nil
	}

	var i int
	var block *pem.Block
	rest := []byte(pemdata)

	for {
		block, rest = pem.Decode(rest)

		if block == nil {
			return makeValidationError(attribute, "Unable to decode PEM")
		}

		if len(rest) == 0 {
			break
		}
		i++
	}

	cacert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return makeValidationError(attribute, fmt.Sprintf("Unable to parse x509 certificate: %s", err))
	}

	if !cacert.IsCA {
		return makeValidationError(attribute, "Given x509 certificate is not a CA")
	}

	return nil
}

// Constants to validate tags.
const (
	prefixDynamicTag  = "$"
	prefixExpandedTag = "+"
	prefixMetadata    = "@"
)

// validateTagStrings validates the given string and check if it can be a valid value for a Tag.
func validateTagStrings(attribute string, acceptReservedPrefix bool, strs ...string) error {

	for _, s := range strs {

		if !acceptReservedPrefix && (strings.HasPrefix(s, prefixMetadata) || strings.HasPrefix(s, prefixDynamicTag) || strings.HasPrefix(s, prefixExpandedTag)) {
			return makeValidationError(attribute, fmt.Sprintf("%s starts with an @, a $ or a + that is reserved", s))
		}

		if err := ValidateTag(attribute, s); err != nil {
			return err
		}
	}

	return nil
}

// tagRegex is the regular expression to check the format of a tag.
var tagRegex = regexp.MustCompile(`^[^= ]+=.+`)

// ValidateTag validates a single tag.
func ValidateTag(attribute string, tag string) error {

	if !tagRegex.MatchString(tag) {
		return makeValidationError(attribute, fmt.Sprintf("`%s must contain at least one '=' symbol separating two valid words", tag))
	}

	if len([]byte(tag)) >= 1024 {
		return makeValidationError(attribute, fmt.Sprintf("`%s must be less than 1024 bytes", tag))
	}

	return nil
}

// ValidateTags validates a list of tags are valid. Accepts those with reserved prefix.
func ValidateTags(attribute string, tags []string) error {
	return validateTagStrings(attribute, true, tags...)
}

// ValidateTagsWithoutReservedPrefixes a list of tags are valid. Refuse those with reserved prefix.
func ValidateTagsWithoutReservedPrefixes(attribute string, tags []string) error {
	return validateTagStrings(attribute, false, tags...)
}

// ValidateTagsExpression validates an [][]string is a valid tag expression.
func ValidateTagsExpression(attribute string, expression [][]string) error {
	for _, tags := range expression {
		if err := ValidateTags(attribute, tags); err != nil {
			return err
		}
	}

	return nil
}

// ValidateMetadata validates an []string is a valid list of metadata.
func ValidateMetadata(attribute string, metadata []string) error {

	for _, m := range metadata {

		if !strings.HasPrefix(m, prefixMetadata) {
			return makeValidationError(attribute, fmt.Sprintf("Metadata %s does not starts with an @, a $ or a + that is reserved", m))
		}

		if strings.HasPrefix(m, constants.AuthKey) {
			return makeValidationError(attribute, fmt.Sprintf("Metadata %s is using @auth: which is reserved", m))
		}

		if err := ValidateTag(attribute, m); err != nil {
			return err
		}
	}

	return nil
}

// ValidateYAMLString validates the given data is a correct YAML string.
func ValidateYAMLString(attribute, data string) error {

	if err := yaml.Unmarshal([]byte(data), &map[string]interface{}{}); err != nil {
		return makeValidationError(attribute, fmt.Sprintf("not a valid yaml: %s", err))
	}
	return nil
}

// ValidateSAMLProvider validate the given SAMLProvider
func ValidateSAMLProvider(provider *SAMLProvider) error {

	if provider.IDPMetadata != "" {
		return nil
	}

	if provider.IDPURL == "" {
		return makeValidationError("IDPURL", "IDPURL is required when no IDPMetadata is provided")
	}

	if provider.IDPIssuer == "" {
		return makeValidationError("IDPIssuer", "IDPIssuer is required when no IDPMetadata is provided")
	}

	if provider.IDPCertificate == "" {
		return makeValidationError("IDPCertificate", "IDPCertificate is required when no IDPMetadata is provided")
	}

	return nil
}

// ValidateAPIAuthorizationPolicySubject makes sure api authorization subject is at least secured a bit.
func ValidateAPIAuthorizationPolicySubject(attribute string, subject [][]string) error {

	for i, ands := range subject {

		if len(ands) < 2 {
			return makeValidationError(attribute, "Subject and line should contain at least 2 claims")
		}

		var realmClaims int
		neededAdditionalMandatoryClaims := map[string]string{}

		keys := map[string]struct{}{}

		for _, claim := range ands {

			if !strings.HasPrefix(claim, "@auth:") {
				return makeValidationError(attribute, fmt.Sprintf("Subject claims '%s' on line %d must be prefixed by '@auth:'", claim, i+1))
			}

			parts := strings.SplitN(claim, "=", 2)
			if parts[1] == "" {
				return makeValidationError(attribute, fmt.Sprintf("Subject claims '%s' on line %d has no value", claim, i+1))
			}
			keys[parts[0]] = struct{}{}

			if strings.HasPrefix(claim, "@auth:realm=") {
				realmClaims++

				switch strings.TrimPrefix(claim, "@auth:realm=") {
				case "oidc":
					neededAdditionalMandatoryClaims["@auth:namespace"] = "The realm oidc mandates to add the '@auth:namespace' key to prevent potential security side effects"
				case "saml":
					neededAdditionalMandatoryClaims["@auth:namespace"] = "The realm saml mandates to add the '@auth:namespace' key to prevent potential security side effects"
				default:
				}
			}
		}

		if realmClaims == 0 {
			return makeValidationError(attribute, fmt.Sprintf("Subject line %d must contain the '@auth:realm' key", i+1))
		}

		if realmClaims > 1 {
			return makeValidationError(attribute, fmt.Sprintf("Subject line %d must contain only one '@auth:realm' key", i+1))
		}

		for mkey, msg := range neededAdditionalMandatoryClaims {
			if _, ok := keys[mkey]; !ok {
				return makeValidationError(attribute, msg)
			}
		}
	}

	return nil
}

// ValidateWriteOperations verifies the given []string only contains
// any unique combination Create, Update or Delete.
func ValidateWriteOperations(attribute string, operations []string) error {

	var ncreate, nupdate, ndelete int

	for _, op := range operations {
		switch elemental.Operation(op) {
		case elemental.OperationCreate:
			ncreate++
		case elemental.OperationUpdate:
			nupdate++
		case elemental.OperationDelete:
			ndelete++
		default:
			return makeValidationError(attribute, fmt.Sprintf("Invalid operation '%s': must be 'create', 'update' or 'delete'.", op))
		}
	}

	if ncreate > 1 || nupdate > 1 || ndelete > 1 {
		return makeValidationError(attribute, fmt.Sprintf("Must not contain the same operation multiple times."))
	}

	return nil
}

// ValidateIdentity verifies the given string is a valid gaia identity.
func ValidateIdentity(attribute string, identity string) error {

	i := Manager().IdentityFromAny(identity)
	if i.IsEmpty() {
		return makeValidationError(attribute, fmt.Sprintf("Invalid identity '%s': unknown.", identity))
	}

	return nil
}

// ValidateStringListNotEmpty validates the given string slice is not empty (or nil).
func ValidateStringListNotEmpty(attribute string, slice []string) error {

	if len(slice) == 0 {
		return makeValidationError(attribute, "String list must not be empty")
	}

	return nil
}
