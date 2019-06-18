package gaia

import (
	"net/http"
	"reflect"
	"testing"

	"go.aporeto.io/elemental"
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

func TestValidateOptionalNetworkList(t *testing.T) {
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
			if err := ValidateOptionalNetworkList(tt.args.attribute, tt.args.networks); (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptionalNetworkList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateHookPolicy(t *testing.T) {
	testCases := []struct {
		scenario      string
		model         *HookPolicy
		expectedError bool
	}{
		{
			scenario: "should NOT return an error if both clientCertificateKey and clientCertificate are absent",
			model: &HookPolicy{
				ClientCertificate:    "",
				ClientCertificateKey: "",
			},
			expectedError: false,
		},
		{
			scenario: "should NOT return an error if both clientCertificateKey and clientCertificate are provided",
			model: &HookPolicy{
				ClientCertificate: `-----BEGIN CERTIFICATE-----
MIIBczCCARigAwIBAgIRALD3Vz81Pq10g7n4eAkOsCYwCgYIKoZIzj0EAwIwJjEN
MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNzA2
NTM1MloXDTI3MTEyNjA2NTM1MlowGDEWMBQGA1UEAxMNY2xhaXJlLWNsaWVudDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABOmzPJj+t25T148eQH5gVrZ7nHwckF5O
evJQ3CjSEMesjZ/u7cW8IBfXlxZKHxl91IEbbB3svci4c8pycUNZ2kujNTAzMA4G
A1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
MAoGCCqGSM49BAMCA0kAMEYCIQCjAAmkQpTua0HR4q6jnePaFBp/JMXwTXTxzbV6
peGbBQIhAP+1OR8GFnn2PlacwHqWXHwkvy6CLPVikvgtwEdB6jH8
-----END CERTIFICATE-----`,
				ClientCertificateKey: `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
F9eXFkofGX3UgRtsHe123456789xQ1naSw==
-----END EC PRIVATE KEY-----`,
			},
			expectedError: false,
		},
		{
			scenario: "should return an error if clientCertificate is provided, but clientCertificateKey is absent",
			model: &HookPolicy{
				ClientCertificate: `-----BEGIN CERTIFICATE-----
MIIBczCCARigAwIBAgIRALD3Vz81Pq10g7n4eAkOsCYwCgYIKoZIzj0EAwIwJjEN
MAsGA1UEChMEQWNtZTEVMBMGA1UEAxMMQWNtZSBSb290IENBMB4XDTE4MDExNzA2
NTM1MloXDTI3MTEyNjA2NTM1MlowGDEWMBQGA1UEAxMNY2xhaXJlLWNsaWVudDBZ
MBMGByqGSM49AgEGCCqGSM49AwEHA0IABOmzPJj+t25T148eQH5gVrZ7nHwckF5O
evJQ3CjSEMesjZ/u7cW8IBfXlxZKHxl91IEbbB3svci4c8pycUNZ2kujNTAzMA4G
A1UdDwEB/wQEAwIHgDATBgNVHSUEDDAKBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
MAoGCCqGSM49BAMCA0kAMEYCIQCjAAmkQpTua0HR4q6jnePaFBp/JMXwTXTxzbV6
peGbBQIhAP+1OR8GFnn2PlacwHqWXHwkvy6CLPVikvgtwEdB6jH8
-----END CERTIFICATE-----`,
				ClientCertificateKey: "",
			},
			expectedError: true,
		},
		{
			scenario: "should return an error if clientCertificateKey is provided, but clientCertificate is absent",
			model: &HookPolicy{
				ClientCertificate: "",
				ClientCertificateKey: `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIGOXJI/123456789oamOu4tQAIKFdbyvkIJg9GME0mHzoAoGCCqGSM49
AwEHoUQDQgAE6bM8mP123456789AfmBWtnucfByQXk568lDcKNIQx6yNn+7txbwg
F9eXFkofGX3UgRtsHe123456789xQ1naSw==
-----END EC PRIVATE KEY-----`,
			},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.scenario, func(t *testing.T) {
			err := ValidateHookPolicy(tc.model)
			if (err != nil) != tc.expectedError {
				t.Fatalf("expected to get an error: %t - received error: %+v", tc.expectedError, err)
			}

			if elemErr, ok := err.(elemental.Error); tc.expectedError {
				if !ok {
					t.Fatalf("expected an elemental error to be returned, but got an error of type: %s", reflect.TypeOf(err))
				}

				if elemErr.Code != http.StatusUnprocessableEntity {
					t.Errorf("expected elemental error code to be 422, but got %d", elemErr.Code)
				}

				if elemErr.Title != "Validation Error" {
					t.Errorf("expected elemental error code to be 'Validation Error', but got %s", elemErr.Title)
				}
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
