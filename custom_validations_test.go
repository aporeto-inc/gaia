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
			true,
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
