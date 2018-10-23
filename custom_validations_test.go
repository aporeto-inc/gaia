package gaia

import (
	"testing"
)

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
			if err := ValidateNetwork(tt.args.attribute, tt.args.cidr); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNetwork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNetworkList(t *testing.T) {
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
			if err := ValidateNetworkList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateNetworkList() error = %v, wantErr %v", err, tt.wantErr)
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
			"valid string all",
			args{
				"proto",
				"all",
			},
			false,
		},
		{
			"valid string ALL",
			args{
				"proto",
				"ALL",
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
			true,
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

func TestValidateHostServicesList(t *testing.T) {
	type args struct {
		attribute    string
		hostservices []*HostService
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"alphanumeric name",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "ssh22",
					},
				},
			},
			false,
		},
		{
			"name with underscore",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "proxy_ssh",
					},
				},
			},
			false,
		},
		{
			"empty name",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "",
					},
				},
			},
			true,
		},
		{
			"name with space",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "Host Service",
					},
				},
			},
			true,
		},
		{
			"name contains hyphen",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "proxy-ssh",
					},
				},
			},
			true,
		},
		{
			"name with 12 characters",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "ansible_proxy_ssh",
					},
				},
			},
			true,
		},
		{
			"invalid name",
			args{
				"hostservices",
				[]*HostService{
					&HostService{
						Name: "###InvalidName!",
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateHostServicesList(tt.args.attribute, tt.args.hostservices); (err != nil) != tt.wantErr {
				t.Errorf("TestValidateHostServicesList() error = %v, wantErr %v", err, tt.wantErr)
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
		// Invalid CIDR
		{
			"valid target network",
			&EnforcerProfile{
				Name:           "Valid target network",
				TargetNetworks: []string{"0.0.0.0/0"},
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
					OIDCProviderURL:   "http://url.com",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateServiceEntity(tt.args.service); (err != nil) != tt.wantErr {
				t.Errorf("ValidateServiceEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
