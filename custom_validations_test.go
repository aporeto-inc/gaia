package gaia

import (
	"fmt"
	"testing"
)

func TestValidateAPIProxyEntity(t *testing.T) {
	type args struct {
		apiProxy *APIProxy
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid cert assignment",
			args{
				&APIProxy{
					ClientCertificate:    "something",
					ClientCertificateKey: "something",
				},
			},
			false,
		},
		{
			"valid empty assignment",
			args{
				&APIProxy{},
			},
			false,
		},
		{
			"invalid only client cert assignment",
			args{
				&APIProxy{
					ClientCertificate: "something",
				},
			},
			true,
		},
		{
			"invalid only cert key assignment",
			args{
				&APIProxy{
					ClientCertificateKey: "something",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateAPIProxyEntity(tt.args.apiProxy); (err != nil) != tt.wantErr {
				t.Errorf("ValidateAPIProxyEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePortString(t *testing.T) {
	type args struct {
		attribute string
		portExp   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple lower port",
			args{
				"port",
				"1",
			},
			false,
		},
		{
			"simple max port",
			args{
				"port",
				"65535",
			},
			false,
		},
		{
			"port set that is not an int",
			args{
				"port",
				"zero",
			},
			true,
		},
		{
			"port set to 65536",
			args{
				"port",
				"0",
			},
			true,
		},

		// range
		{
			"port range",
			args{
				"port",
				"1:65535",
			},
			false,
		},
		{
			"port range with left bound set to 0",
			args{
				"port",
				"0:65535",
			},
			true,
		},
		{
			"port range with left bound set to 65536",
			args{
				"port",
				"1:65536",
			},
			true,
		},
		{
			"port range with left bound greater than right bound",
			args{
				"port",
				"65535:1",
			},
			true,
		},
		{
			"port range with left bound that is not a int",
			args{
				"port",
				"1:two",
			},
			true,
		},

		// format
		{
			"broken port expression 1",
			args{
				"port",
				"1-3",
			},
			true,
		},
		{
			"broken port expression 2",
			args{
				"port",
				"1-",
			},
			true,
		},
		{
			"broken port expression 3",
			args{
				"port",
				"-1",
			},
			true,
		},
		{
			"broken port expression 4",
			args{
				"port",
				"-",
			},
			true,
		},
		{
			"broken port expression 5",
			args{
				"port",
				"",
			},
			true,
		},
		{
			"broken port expression 6",
			args{
				"port",
				"1:2:3",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePortString(tt.args.attribute, tt.args.portExp); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePortString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServicePorts(t *testing.T) {
	type args struct {
		servicePorts []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid serviceports",
			args{
				[]string{"tcp/80", "udp/80", "icmp", "ISIS"},
			},
			false,
		},
		{
			"valid serviceports with only any",
			args{
				[]string{"aNy"},
			},
			false,
		},
		{
			"Any with other protoocls",
			args{
				[]string{"tcp/80", "udp/80", "any", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid tcp serviceports",
			args{
				[]string{"tcp", "udp/80", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid udp serviceports",
			args{
				[]string{"tcp/80", "udp/", "icmp", "ISIS"},
			},
			true,
		},
		{
			"invalid protocol",
			args{
				[]string{"nope"},
			},
			true,
		},
		{
			"invalid serviceports",
			args{
				[]string{"tcp/80/90"},
			},
			true,
		},
		{
			"serviceports with protocol with ports not supported",
			args{
				[]string{"tcp/90:8000", "udp/90:8000", "icmp/9090"},
			},
			true,
		},
		{
			"one serviceport with protocol with ports not supported",
			args{
				[]string{"isis/9090"},
			},
			true,
		},
		{
			"serviceports with valid port range",
			args{
				[]string{"tcp/90:8000", "udp/90:8000"},
			},
			false,
		}, {
			"serviceports with protocol numbers",
			args{
				[]string{"udp/90:8000", "6/90:8000"},
			},
			true,
		},
		{
			"serviceports with invalid port range",
			args{
				[]string{"tcp/90:8000", "udp/90:800000"},
			},
			true,
		},
		{
			"serviceports with just port range",
			args{
				[]string{"88", "90:8000"},
			},
			true,
		},
		{
			"empty string serviceports",
			args{
				[]string{""},
			},
			true,
		},
		{
			"empty serviceports",
			args{
				[]string{},
			},
			false,
		},
		{
			"nil serviceports",
			args{
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServicePorts("serviceports", tt.args.servicePorts); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServicePorts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServicePort(t *testing.T) {
	type args struct {
		servicePort string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid serviceport",
			args{
				"tcp/80",
			},
			false,
		},
		{
			"valid serviceport any",
			args{
				"aNy",
			},
			false,
		},
		{
			"invalid tcp serviceport",
			args{
				"tcp",
			},
			true,
		},
		{
			"invalid udp serviceport",
			args{
				"udp/",
			},
			true,
		},
		{
			"invalid protocol",
			args{
				"nope",
			},
			true,
		},
		{
			"invalid serviceport",
			args{
				"tcp/80/90",
			},
			true,
		},
		{
			"serviceport with protocol with ports not supported",
			args{
				"icmp/9090",
			},
			true,
		},
		{
			"one serviceport with protocol with ports not supported",
			args{
				"isis/9090",
			},
			true,
		},
		{
			"serviceport with valid port range",
			args{
				"udp/90:8000",
			},
			false,
		}, {
			"serviceport with protocol numbers",
			args{
				"6/90:8000",
			},
			true,
		},
		{
			"serviceport with invalid port range",
			args{
				"udp/90:800000",
			},
			true,
		},
		{
			"serviceport with just port range",
			args{
				"90:8000",
			},
			true,
		},
		{
			"empty serviceport",
			args{
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServicePort("serviceport", tt.args.servicePort); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServicePort() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePortStringList(t *testing.T) {
	type args struct {
		attribute string
		ports     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"correct list",
			args{
				"ports",
				[]string{"443", "555:666", "80"},
			},
			false,
		},
		{
			"incorrect list",
			args{
				"ports",
				[]string{"443", "555:666", "80", "yo"},
			},
			true,
		},
		{
			"empty list",
			args{
				"ports",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePortStringList(tt.args.attribute, tt.args.ports); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePortStringList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNetwork(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"google.com",
			},
			false,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"*.google.com",
			},
			false,
		},

		// invalid CIDR
		{
			"invalid CIDR",
			args{
				"cidr",
				"",
			},
			true,
		},

		// invalid DNn
		{
			"invalid DNS Name",
			args{
				"cidr",
				"google@com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNetworkOrHostname(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNetworkOrHostname() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNetworkOrHostnameList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateNetworkOrHostnameList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNetworkOrHostnameList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDR(t *testing.T) {
	type args struct {
		attribute string
		cidr      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// valid
		{
			"valid CIDR",
			args{
				"cidr",
				"10.0.0.0/8",
			},
			false,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"google.com",
			},
			true,
		},
		{
			"valid DNS name",
			args{
				"cidr",
				"*.google.com",
			},
			true,
		},

		// invalid CIDR
		{
			"invalid CIDR",
			args{
				"cidr",
				"",
			},
			true,
		},

		// invalid DNn
		{
			"invalid DNS Name",
			args{
				"cidr",
				"google@com",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDR(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateCIDR() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCIDRList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDRList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCIDRList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProtocol(t *testing.T) {
	type args struct {
		attribute string
		proto     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid string tcp",
			args{
				"proto",
				"tcp",
			},
			false,
		},
		{
			"valid string TCP",
			args{
				"proto",
				"TCP",
			},
			false,
		},
		{
			"valid string udp",
			args{
				"proto",
				"udp",
			},
			false,
		},
		{
			"valid string UDP",
			args{
				"proto",
				"UDP",
			},
			false,
		},
		{
			"valid string IPIP",
			args{
				"proto",
				"IPIP",
			},
			false,
		},
		{
			"valid string ipip",
			args{
				"proto",
				"IPIP",
			},
			false,
		},
		{
			"valid string icmp",
			args{
				"proto",
				"icmp",
			},
			false,
		},
		{
			"valid string ICMP",
			args{
				"proto",
				"ICMP",
			},
			false,
		},
		{
			"valid string number",
			args{
				"proto",
				"30",
			},
			false,
		},

		// invalid
		{
			"invalid string",
			args{
				"proto",
				"NOT",
			},
			true,
		},
		{
			"invalid number 1",
			args{
				"proto",
				"-1",
			},
			true,
		},
		{
			"invalid number 2",
			args{
				"proto",
				"256",
			},
			true,
		},
		{
			"valid string all",
			args{
				"proto",
				"all",
			},
			true,
		},
		{
			"valid string ALL",
			args{
				"proto",
				"ALL",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProtocol(tt.args.attribute, tt.args.proto); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProtocol() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProtocolList(t *testing.T) {
	type args struct {
		attribute string
		protocols []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"proto",
				[]string{"udp", "UDP", "TCP", "tcp", "3", "43"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"proto",
				[]string{"udp", "UDP", "NOOOOO", "tcp", "3", "43"},
			},
			true,
		},
		{
			"empty list",
			args{
				"proto",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProtocolList(tt.args.attribute, tt.args.protocols); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProtocolList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateEnforcerProfile(t *testing.T) {
	tests := []struct {
		name            string
		enforcerprofile *EnforcerProfile
		wantErr         bool
	}{
		// Target Networks
		{
			"valid target network",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"0.0.0.0/0"},
			},
			false,
		},
		{
			"valid target networks with except condition operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid target network",
			&EnforcerProfile{
				Name:           "Invalid target network",
				TargetNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid target network with multiple NOT operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid target network with NOT operator",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Target UDP Networks
		{
			"valid target UDP network",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0"},
			},
			false,
		},
		{
			"valid target UDP networks with except condition operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid target UDP network",
			&EnforcerProfile{
				Name:              "Invalid target UDP network",
				TargetUDPNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid target UDP network with multiple NOT operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid target UDP network with NOT operator",
			&EnforcerProfile{
				Name:              "Valid target UDP network",
				TargetUDPNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Excluded Networks
		{
			"valid excluded network",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"0.0.0.0/0"},
			},
			false,
		},
		{
			"valid excluded networks with except condition operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"0.0.0.0/0", "!10.0.0.0/8"},
			},
			false,
		},
		{
			"invalid excluded network",
			&EnforcerProfile{
				Name:             "Invalid excluded network",
				ExcludedNetworks: []string{"invalid"},
			},
			true,
		},
		{
			"invalid excluded network with multiple NOT operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"!!10.0.0.0/8"},
			},
			true,
		},
		{
			"valid excluded network with NOT operator",
			&EnforcerProfile{
				Name:             "Valid excluded network",
				ExcludedNetworks: []string{"!10.0.0.0/8"},
			},
			true,
		},
		// Trusted CAs
		{
			"valid trusted CA",
			&EnforcerProfile{
				Name: "Valid CAs",
				TrustedCAs: []string{`-----BEGIN CERTIFICATE-----
MIIBPzCB5aADAgECAhB8oW2KA1BfhUmp6B1TkvcRMAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRlc3QwHhcNMTgxMDAyMjI1MjUwWhcNMjgwODEwMjI1MjUwWjAPMQ0w
CwYDVQQDEwR0ZXN0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsHxO39rdxGtD
8bP5AW2gqTMxgq4w9nyVbbpetS0rtjxwZ5bMgdS4GmqaifjqNGGt+VkK7gFNEyqG
Hq/uTgtJ56MjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSQAwRgIhAKRcwnrREJ6EnHOsiUnDfuNFxxALw4kV/ZyRxl1CJcS+
AiEA0epxATHNzheAa8ZuiPeNQL6DhoKYz3B+41J2vgVlGZY=
-----END CERTIFICATE-----`},
			},
			false,
		},
		{
			"invalid trusted CA",
			&EnforcerProfile{
				Name:       "Invalid CAs",
				TrustedCAs: []string{"invalid ca"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEnforcerProfile(tt.enforcerprofile); (err != nil) != tt.wantErr {
				t.Errorf("TestVValidateEnforcerProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateServiceEntity(t *testing.T) {
	type args struct {
		service *Service
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// OIDC
		{
			"valid OIDC service",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "https://url.com",
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
				},
			},
			false,
		},
		{
			"OIDC service with missing OIDCProviderURL",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
				},
			},
			true,
		},
		{
			"OIDC service with bad OIDCProviderURL - no https",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					OIDCProviderURL:   "http://url.com",
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
				},
			},
			true,
		},
		{
			"OIDC service with bad OIDCProviderURL - no https",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					OIDCProviderURL:   "http://!@#!@#/1231",
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCClientID:      "client id",
					OIDCClientSecret:  "client secret",
				},
			},
			true,
		},
		{
			"OIDC service with missing OIDCClientID",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "http://url.com",
					OIDCClientSecret:  "client secret",
				},
			},
			true,
		},
		{
			"OIDC service with missing OIDCClientSecret",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeOIDC,
					OIDCProviderURL:   "http://url.com",
					OIDCClientID:      "client id",
				},
			},
			true,
		},

		// JWT
		{
			"valid JWT service",
			args{
				&Service{
					Hosts:                 []string{"myservice.com"},
					AuthorizationType:     ServiceAuthorizationTypeJWT,
					JWTSigningCertificate: "---pem---",
				},
			},
			false,
		},
		{
			"JWT service with missing JWTSigningCertificate",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeJWT,
				},
			},
			true,
		},

		// TLS
		{
			"valid service with type TLSType set to Aporeto",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeAporeto,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to LetsEcncrypt",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeLetsEncrypt,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to None",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			false,
		},
		{
			"valid service with type TLSType set to External",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					TLSCertificate:    "---cert---",
					TLSCertificateKey: "---key---",
				},
			},
			false,
		},
		{
			"service with type TLSType set to External with missing cert and key",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
				},
			},
			true,
		},
		{
			"service with type TLSType set to External with missing cert",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					TLSCertificateKey: "---key---",
				},
			},
			true,
		},
		{
			"service with type TLSType set to External with missing key",
			args{
				&Service{
					Hosts:             []string{"myservice.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeExternal,
					TLSCertificate:    "---key---",
				},
			},
			true,
		},

		// Hosts an IP
		{
			"service missing hosts and ips",
			args{
				&Service{
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with invalid hosts ",
			args{
				&Service{
					Hosts:             []string{"@#$@#$.2234242"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with overlapping hosts",
			args{
				&Service{
					Hosts:             []string{"foo.com", "foo.com"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with bad IPs",
			args{
				&Service{
					IPs:               []string{"1234.2.3.3"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with bad subnets",
			args{
				&Service{
					IPs:               []string{"10.1.1.0/48"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with overlapping subnets",
			args{
				&Service{
					IPs:               []string{"10.1.0.0/24", "10.1.0.0/16"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with multiple IPs and multiple error case",
			args{
				&Service{
					IPs:               []string{"10.11.0.1234/24", "10.1.0.0/16", "10.1.0.3/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			true,
		},
		{
			"service with multiple IPs positive case",
			args{
				&Service{
					IPs:               []string{"10.11.0.123/24", "10.1.0.0/16", "10.19.0.33/24"},
					AuthorizationType: ServiceAuthorizationTypeMTLS,
					TLSType:           ServiceTLSTypeNone,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServiceEntity(tt.args.service); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServiceEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHTTPMethods(t *testing.T) {
	type args struct {
		attribute string
		methods   []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty list of methods",
			args{
				attribute: "methods",
				methods:   []string{},
			},
			false,
		},
		{
			"valid list of methods",
			args{
				attribute: "methods",
				methods:   []string{"POST"},
			},
			false,
		},
		{
			"invalid list of methods",
			args{
				attribute: "methods",
				methods:   []string{"BIDULE"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHTTPMethods(tt.args.attribute, tt.args.methods); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHTTPMethods() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHTTPSURL(t *testing.T) {
	type args struct {
		attribute string
		address   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"empty url",
			args{
				attribute: "endpoint",
				address:   "",
			},
			true,
		},
		{
			"valid url",
			args{
				attribute: "endpoint",
				address:   "https://aporeto.com/",
			},
			false,
		},
		{
			"invalid url",
			args{
				attribute: "endpoint",
				address:   "htps:/a",
			},
			true,
		},
		{
			"non https url",
			args{
				attribute: "endpoint",
				address:   "http://hello.com/",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHTTPSURL(tt.args.attribute, tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("ValidateHTTPSURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalNetworkOrHostnameList(t *testing.T) {
	type args struct {
		attribute string
		networks  []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google.com"},
			},
			false,
		},
		{
			"invalid list",
			args{
				"nets",
				[]string{"10.0.0.0/8", "google@com"},
			},
			true,
		},
		{
			"empty list",
			args{
				"nets",
				[]string{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalNetworkOrHostnameList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalNetworkOrHostnameList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAutomation(t *testing.T) {
	testCases := map[string]struct {
		automation    *Automation
		expectedError error
	}{
		"should not return an error if trigger type is not webhook and multiple actions have been defined": {
			automation: &Automation{
				Condition: `
						function when(api, params) {
							return {
								continue: true,
							};
						}
			 	`,
				Trigger: AutomationTriggerRemoteCall,
				Actions: []string{
					"Action 1",
					"Action 2",
				},
			},
			expectedError: nil,
		},
		"should not return an error if trigger type is webhook and one action has been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: []string{
					"Action 1",
				},
			},
			expectedError: nil,
		},
		"should return an error if trigger type is set to webhook and more than one action has been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: []string{
					"Action 1",
					"Action 2",
				},
			},
			expectedError: makeValidationError("actions", fmt.Sprintf("Only one action can be defined if trigger type is set to %q", AutomationTriggerWebhook)),
		},
		"should return an error if trigger type is set to webhook and no actions have been defined": {
			automation: &Automation{
				Trigger: AutomationTriggerWebhook,
				Actions: nil,
			},
			expectedError: makeValidationError("actions", fmt.Sprintf("Exactly one action must be defined if trigger type is set to %q", AutomationTriggerWebhook)),
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			err := ValidateAutomation(tc.automation)
			switch {
			case err != nil && tc.expectedError != nil:
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("\n"+
						"actual error: %+v\n"+
						"did not equal\n"+
						"expected error: %+v", err, tc.expectedError)
				}
			case err != nil && tc.expectedError == nil:
				t.Fatalf("did not expect to get an error, but received: %+v", err)
			case err == nil && tc.expectedError != nil:
				t.Fatalf("expected to get an error, but got nothing")
			}
		})
	}
}

func TestValidateHostService(t *testing.T) {
	tests := []struct {
		name    string
		hs      *HostService
		wantErr bool
	}{
		{
			"valid service",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			false,
		},
		{
			"valid service name < 128",
			&HostService{
				Name:            "asdasdsakljfslahfjashjkdlaksfjaksjfvfsadlkjkljsad",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			false,
		},
		{
			"valid service name > 128",
			&HostService{
				Name:            "asdasdsakljfslahfjashjkdlaksfsdhkjfhsdjkhfsiudfhsdjkfnsdkjfn1342141397asldjasklfjakslfj879797kjasdhjaksjfvfsadlkjkljsaddasfhsajlhdkajshfjksahfjksahdjksahdkjashdkasdhkjsa",
				HostModeEnabled: false,
				Services:        []string{"tcp/80"},
			},
			true,
		},
		{
			"invalid name",
			&HostService{
				Name:            "jboss@!#$!$!",
				HostModeEnabled: false,
				Services:        []string{},
			},
			true,
		},
		{
			"empty name",
			&HostService{
				Name:            "",
				HostModeEnabled: false,
				Services:        []string{},
			},
			true,
		},
		{
			"invalid services",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"foo / 123"},
			},
			true,
		},
		{
			"tcp port overlaps",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"tcp/80:100, tcp/90"},
			},
			true,
		},
		{
			"udp port overlaps",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"udp/80:100, udp/90"},
			},
			true,
		},
		{
			"non overlap with tcp and udp ports",
			&HostService{
				Name:            "jboss",
				HostModeEnabled: false,
				Services:        []string{"udp/90", "tcp/90"},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHostServices(tt.hs); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateHostServicesList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProtoPorts(t *testing.T) {
	tests := []struct {
		name     string
		services []string
		wantErr  bool
	}{
		{
			"valid services",
			[]string{"tcp/80", "udp/80"},
			false,
		},
		{
			"invalid protocol",
			[]string{"tcp/80", "bro/80"},
			true,
		},
		{
			"invalid port format",
			[]string{"tcp/80", "bro/code"},
			true,
		},
		{
			"invalid port number",
			[]string{"tcp/80", "bro/10000"},
			true,
		},
		{
			"nil services",
			nil,
			false,
		},
		{
			"empty services",
			[]string{},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProtoPorts("ports", tt.services); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateProtoPorts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isFQDN(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantValue bool
	}{
		{
			"valid fqdn",
			args{
				"www.google.com",
			},
			true,
		},
		{
			"xip",
			args{
				"1.2.3.4.xip.io",
			},
			true,
		},
		{
			"trailing dot",
			args{
				"www.google.com.",
			},
			true,
		},
		{
			"bad one",
			args{
				"@#$@#$@.#$@#.#@$2",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if isFQDN(tt.args.name) != tt.wantValue {
				t.Errorf("isFQDN() failed at test: %s", tt.name)
			}
		})
	}
}

func TestValidatePEM(t *testing.T) {
	type args struct {
		attribute string
		pemdata   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"nothing set",
			args{
				"pem",
				``,
			},
			false,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----`,
			},
			false,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
`,
			},
			false,
		},
		{
			"invalid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoT ----NOT PEM---- I3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----`,
			},
			true,
		},
		{
			"valid single PEM",
			args{
				"pem",
				`-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5loNahC19hQtExCi64fmI3QCcrEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIBpDCCAUmgAwIBAgIQDbXKAZzk9RjcNSGMsWke1zAKBggqhkjOPQQDAjBGMRAw
DgYDVQQKEwdBcG9yZXRvMQ8wDQYDVQQLEwZhcG9tdXgxITAfBgNVBAMTGEFwb211
eCBQdWJsaWMgU2lnbmluZyBDQTAeFw0xOTAxMjQyMjQ3MjlaFw0yODEyMDIyMjQ3
MjlaMCoxEjAQBgNVBAoTCXNlcGhpcm90aDEUMBIGA1UEAxMLYXV0b21hdGlvbnMw
WTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASxKA9vbyk7FXXlOCi0kTKLVne/mK8o
ZQDPRcehze0EMwTAR5     ----NOT PEM----   crEGH9ycUoITYPgozUwMzAO
BgNVHQ8BAf8EBAMCB4AwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIw
ADAKBggqhkjOPQQDAgNJADBGAiEAm1u2T1vRooIy3rd0BmBSAa6WR6BtHl9nDbGN
1ZM+SgsCIQDu4R6OziiWbRdn50bneZT5qPO+07ALY5m4DG96VyCaQw==
-----END CERTIFICATE-----
`,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePEM(tt.args.attribute, tt.args.pemdata); (err != nil) != tt.wantErr {
				t.Errorf("ValidatePEM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// simple
		{
			"simple tag a=a",
			args{
				"a=a",
			},
			false,
		},
		{
			"value is string",
			args{
				"name=value",
			},
			false,
		},
		{
			"value is number",
			args{
				"name=123.12",
			},
			false,
		},
		{
			"value is has space",
			args{
				"name=test 1",
			},
			false,
		},
		{
			"tag value is utf8 character",
			args{
				"name=utf8-_!@#%&\" (*)+.,/$!:;<>=?{}~",
			},
			false,
		},
		{
			"key contains hyphen and underscore",
			args{
				"internal_name:demo-1234=aporeliable_arvind_@",
			},
			false,
		},
		{
			"tag contains 2 equals",
			args{
				"a=a=b",
			},
			false,
		},
		{
			"tag contains 2 touching equals",
			args{
				"a==b",
			},
			false,
		},
		{
			"complex key",
			args{
				`utf8-_!@#%&"(*)+.,/$!:;<>=?{}~=utf8-_!@#%&" (*)+.,/$!:;<>=?{}~`,
			},
			false,
		},

		// Error
		{
			"just a word",
			args{
				"justaword",
			},
			true,
		},
		{
			"key contains spaces",
			args{
				"the key=value",
			},
			true,
		},
		{
			"key starts with =",
			args{
				"=thekey=value",
			},
			true,
		},

		// tag size
		{
			"tags is exactly 1024 bytes",
			args{
				"l=ah6ooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exactly 1026 bytes",
			args{
				"l=ah6ooPaelaizohakee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			true,
		},
		{
			"tags is exacty 1023 bytes",
			args{
				"l=ahooPaelaizohkee6iephoomueMaekosoothae9ieCheR8Foo7aivoo4Aicohphe0bu8wiM8Uu8eidiwui7doowiashahf7niepheeboh7quicaixuPosia3Jee0dahco0do6Teilae4quah9uequeer6che3ojudi2echee8eic1gaeNgiesh0baosh8oo4Fie1Ou8aithaphah4quae4ooteiKoayai7to6ahf8aex8ieDaiw9ag3aodii8tho3ik8Aquaeshaejah8ain0thax3au8ooqueew2eike2kuiShohk4dohw7soogh0zeibieg0idae6pedooma5Lei1ohhaech5Esh7tashieBoolahshohm5Qua1aiy0uho5eichiev0ohp0iu0Izaizi8aisepoox3dahvai7Sudeek6Ea1Ooshoo7aZoofiesh7ithah0Wee1eechoh1xi8Ohd5xahHae5phoiw7eejengu2Eich1Oonahgh3gi1kaHiaD5eeSh7Goophakeiqu7meevi3phiezooSie2uf7PhooThoo6Mooqu1dood2Zeshohm1CheiGh6ShierahJooshohphi6eephi9aeM1isu0IejahQuea9Io1Ahcahque1ce4od7xaYeejoh2ahtimaoluyahyie3ahYohFoVieGup8eYu0maep4Eeghoh3beisheu9ieloe7ichei6johfahLaeCat0naeque5EoyohjieN2era2AejieCou0sei3aej9Che5javah5equ1ieB3eiliegefahfoo3ieth1iequ7zayie9ocheiSaisiech6nai3Ucieroov6iogoonim4zo1iigai7haiF0Yae4cev7Phe0boir3eiteuga8Phuayei6Oj5ohNgog9aeThiewoowoh7shair9wa4Oongairangaa0baemaemae1saefutechah6iev0Aex3aevizeath2ep3hooceaxabuchukahP7Iphae6na",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTag("tag", tt.args.tag); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateYAMLString(t *testing.T) {
	type args struct {
		attribute string
		data      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid yaml",
			args{
				"prop",
				`
hello: world
things:
- stuff
- trucs
`,
			},
			false,
		},
		{
			"invalid yaml",
			args{
				"prop",
				`not a yaml`,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateYAMLString(tt.args.attribute, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ValidateYAMLString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateSAMLProvider(t *testing.T) {
	type args struct {
		provider *SAMLProvider
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"with nothing",
			args{&SAMLProvider{}},
			true,
		},
		{
			"with IDP metadata",
			args{&SAMLProvider{
				IDPMetadata: "m",
			}},
			false,
		},
		{
			"with every other field",
			args{&SAMLProvider{
				IDPURL:         "IDPURL",
				IDPIssuer:      "IDPIssuer",
				IDPCertificate: "IDPCertificate",
			}},
			false,
		},
		{
			"with every other field but IDPURL",
			args{&SAMLProvider{
				IDPIssuer:      "IDPIssuer",
				IDPCertificate: "IDPCertificate",
			}},
			true,
		},
		{
			"with every other field but IDPIssuer",
			args{&SAMLProvider{
				IDPURL:         "IDPURL",
				IDPCertificate: "IDPCertificate",
			}},
			true,
		},
		{
			"with every other field but IDPCertificate",
			args{&SAMLProvider{
				IDPURL:    "IDPURL",
				IDPIssuer: "IDPIssuer",
			}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateSAMLProvider(tt.args.provider); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSAMLProvider() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateAPIAuthorizationPolicySubject(t *testing.T) {
	type args struct {
		attribute string
		subject   [][]string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"valid subject",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"missing realm claim",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject line 2 must contain the '@auth:realm' key",
		},
		{
			"2 realm claims",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a", "@auth:realm=vince"},
					{"@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject line 1 must contain only one '@auth:realm' key",
		},
		{
			"single claim line",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:realm=certificate"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject and line should contain at least 2 claims",
		},
		{
			"missing auth prefix claim",
			args{
				"subject",
				[][]string{
					{"@auth:realm=certificate", "@auth:claim=a"},
					{"@auth:claim=a", "@auth:claim=b", "not:good"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims 'not:good' on line 2 must be prefixed by '@auth:'",
		},
		{
			"oidc correct",
			args{
				"subject",
				[][]string{
					{"@auth:realm=oidc", "@auth:claim=a", "@auth:namespace=/a/b"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"oidc missing namespace",
			args{
				"subject",
				[][]string{
					{"@auth:realm=oidc", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: The realm oidc mandates to add the '@auth:namespace' key to prevent potential security side effects",
		},
		{
			"saml correct",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim=a", "@auth:namespace=/a/b"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			false,
			"",
		},
		{
			"saml missing namespace",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim=a"},
					{"@auth:realm=vince", "@auth:claim=a", "@auth:claim=b"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: The realm saml mandates to add the '@auth:namespace' key to prevent potential security side effects",
		},
		{
			"broken tag with no equal",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim"},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims '@auth:claim' on line 1 is an invalid tag",
		},
		{
			"broken tag with no value",
			args{
				"subject",
				[][]string{
					{"@auth:realm=saml", "@auth:claim="},
				},
			},
			true,
			"error 422 (gaia): Validation Error: Subject claims '@auth:claim=' on line 1 has no value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAPIAuthorizationPolicySubject(tt.args.attribute, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAPIAuthorizationPolicySubject() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.wantErrString {
				t.Errorf("ValidateAPIAuthorizationPolicySubject() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateWriteOperations(t *testing.T) {
	type args struct {
		attribute  string
		operations []string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"nil",
			args{
				"thing",
				nil,
			},
			false,
			"",
		},
		{
			"empty",
			args{
				"thing",
				[]string{},
			},
			false,
			"",
		},
		{
			"valid list",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"delete",
				},
			},
			false,
			"",
		},
		{
			"invalid list",
			args{
				"thing",
				[]string{
					"Create",
					"Update",
					"NotDelete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Invalid operation 'Create': must be 'create', 'update' or 'delete'.",
		},
		{
			"duplicate create",
			args{
				"thing",
				[]string{
					"create",
					"create",
					"update",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times.",
		},
		{
			"duplicate update",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"update",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times.",
		},
		{
			"duplicate delete",
			args{
				"thing",
				[]string{
					"create",
					"update",
					"delete",
					"delete",
				},
			},
			true,
			"error 422 (gaia): Validation Error: Must not contain the same operation multiple times.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateWriteOperations(tt.args.attribute, tt.args.operations)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateWriteOperations() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.wantErrString {
				t.Errorf("ValidateWriteOperations() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateIdentity(t *testing.T) {
	type args struct {
		attribute string
		identity  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"processingunit",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"yo",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIdentity(tt.args.attribute, tt.args.identity); (err != nil) != tt.wantErr {
				t.Errorf("ValidateIdentity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProcessingUnitPolicy(t *testing.T) {
	type args struct {
		policy *ProcessingUnitPolicy
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Empty processing unit policy",
			args{
				NewProcessingUnitPolicy(),
			},
			true,
		},
		{
			"processing unit policy with Action and DatapathType set to Default",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionDefault,
					DatapathType: ProcessingUnitPolicyDatapathTypeDefault,
				},
			},
			true,
		},
		{
			"processing unit policy with Action set to Default and DatapathType to something else",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionDefault,
					DatapathType: ProcessingUnitPolicyDatapathTypeAporeto,
				},
			},
			false,
		},
		{
			"processing unit policy with DatapathType set to Default and Action to something else",
			args{
				&ProcessingUnitPolicy{
					Action:       ProcessingUnitPolicyActionEnforce,
					DatapathType: ProcessingUnitPolicyDatapathTypeDefault,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateProcessingUnitPolicy(tt.args.policy); (err != nil) != tt.wantErr {
				t.Errorf("ValidateProcessingUnitPolicy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateStringListNotEmpty(t *testing.T) {
	type args struct {
		attribute string
		slice     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"nil",
			args{
				"attr",
				nil,
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				[]string{},
			},
			true,
		},
		{
			"set",
			args{
				"attr",
				[]string{"a=a"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateStringListNotEmpty(tt.args.attribute, tt.args.slice); (err != nil) != tt.wantErr {
				t.Errorf("ValidateStringListNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCA(t *testing.T) {

	var (
		notCA = `-----BEGIN CERTIFICATE-----
MIIBKzCB0qADAgECAhBaMjFSVhiDqqG1HMQaPNB3MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjEyWhcNMzAwMjE3MTczMjEyWjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAESFKAZYknsQ/P
u847pk7vkAWkh0hlWCGm8ycTqJtREmfcl3MUYarXjJJLxrg5WmM+xRJuuiTudyLR
MeJFDIVRkqMQMA4wDAYDVR0TAQH/BAIwADAKBggqhkjOPQQDAgNIADBFAiEAr5Sn
5NMpnLPc1yy1A3oyX+iHj2jjPRc/FkTK1ozPA04CIAvBYS8+5Hq+HzmryIP0A57k
dGcCFgI8uzkLKtXLq2n0
-----END CERTIFICATE-----`
		CA = `-----BEGIN CERTIFICATE-----
MIIBPjCB5aADAgECAhAumbVWnomSelU1zBRe9qp4MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjU0WhcNMzAwMjE3MTczMjU0WjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUWqBTwhWMrtr
vv2i+nyiLpwZPV+BfJs+W+qpYLxX6A+WqFTNvu/frdm6dfKLc022QIMzcJwsM+Yz
oBYZoyf02KMjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSAAwRQIhAILlr7jJBS3syxM34zLlwvaoRiCFWnVPHNQeSJtk0h7l
AiBZg9mafabskWu9ekgZm50rKkPeqF94tY6R7tuZBUdcVQ==
-----END CERTIFICATE-----`
		brokenCA = `-----BEGIN CERTIFICATE-----
MDIBPjCB5aADAgECAhAumbVWnomSelU1zBRe9qp4MAoGCCqGSM49BAMCMA8xDTAL
BgNVBAMTBHRvdG8wHhcNMjAwNDEwMTczMjU0WhcNMzAwMjE3MTczMjU0WjAPMQ0w
CwYDVQQDEwR0b3RvMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUWqBTwhWMrtr
vv2i+nyiLpwZPV+BfJs+W+qpYLxX6A+WqFTNvu/frdm6dfKLc022QIMzcJwsM+Yz
oBYZoyf02KMjMCEwDgYDVR0PAQH/BAQDAgEGMA8GA1UdEwEB/wQFMAMBAf8wCgYI
KoZIzj0EAwIDSAAwRQIhAILlr7jJBS3syxM34zLlwvaoRiCFWnVPHNQeSFtk0h7l
AiBZg9mafabskWu9ekgZm50rKkPeqF94tY6R7tuZBUdcVQ==
-----END CERTIFICATE-----`
	)
	type args struct {
		attribute string
		pemdata   string
	}
	tests := []struct {
		name          string
		args          args
		wantErr       bool
		wantErrString string
	}{
		{
			"empty",
			args{
				"attr",
				``,
			},
			false, "",
		},
		{
			"invalid pem",
			args{
				"attr",
				`oh no`,
			},
			true, "error 422 (gaia): Validation Error: Unable to decode PEM",
		},
		{
			"valid pem / invalid CA",
			args{
				"attr",
				brokenCA,
			},
			true, `error 422 (gaia): Validation Error: Unable to parse x509 certificate: asn1: structure error: tags don't match (16 vs {class:0 tag:1 length:62 isCompound:false}) {optional:false explicit:false application:false private:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} tbsCertificate @2`,
		},
		{
			"valid pem but not CA",
			args{
				"attr",
				notCA,
			},
			true, "error 422 (gaia): Validation Error: Given x509 certificate is not a CA",
		},

		{
			"everything is fine",
			args{
				"attr",
				CA,
			},
			false, "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCA(tt.args.attribute, tt.args.pemdata)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCA() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.wantErrString {
				t.Errorf("ValidateCA() error = '%v', wantErrString = '%v'", err, tt.wantErrString)
			}
		})
	}
}

func TestValidateTimeDuration(t *testing.T) {
	type args struct {
		attribute string
		duration  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"1m",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"dog",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateTimeDuration(tt.args.attribute, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("ValidateTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptionalTimeDuration(t *testing.T) {
	type args struct {
		attribute string
		duration  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"valid",
			args{
				"attr",
				"1m",
			},
			false,
		},
		{
			"invalid",
			args{
				"attr",
				"dog",
			},
			true,
		},
		{
			"empty",
			args{
				"attr",
				"",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateOptionalTimeDuration(tt.args.attribute, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalTimeDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
