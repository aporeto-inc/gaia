package netutils

import (
	"net"
	"testing"
)

func Test_prefixIsContained(t *testing.T) {
	type args struct {
		prefixes []string
		ip       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic test",
			args: args{
				prefixes: []string{"10.10.0.0/16", "0.0.0.0/0"},
				ip:       "10.10.10.10/32",
			},
			want: true,
		},
		{
			name: "basic test failure",
			args: args{
				prefixes: []string{},
				ip:       "20.20.20.20/32",
			},
			want: false,
		},
		{
			name: "multiple match test",
			args: args{
				prefixes: []string{"10.10.10.0/24", "10.10.0.0/16", "0.0.0.0/0"},
				ip:       "10.10.10.10/32",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prefixes, err := parseCIDRs(tt.args.prefixes)
			if err != nil {
				t.Errorf("err in test: %v", err)
			}
			ip, _, err := net.ParseCIDR(tt.args.ip)
			if err != nil {
				t.Errorf("err in test: %v", err)
			}
			if got := prefixIsContained(prefixes, ip); got != tt.want {
				t.Errorf("prefixIsContained() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ValidateCIDRs(t *testing.T) {
	type args struct {
		cidrs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "basic test",
			args: args{
				[]string{"10.10.10.10/32"},
			},
			wantErr: false,
		},
		{
			name: "basic inclusion test",
			args: args{
				[]string{"10.10.10.0/24", "!10.10.10.10/32"},
			},
			wantErr: false,
		},
		{
			name: "basic failure test",
			args: args{
				[]string{"!10.10.10.10/32"},
			},
			wantErr: true,
		},
		{
			name: "recursive test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8", "0.0.0.0/0"},
			},
			wantErr: false,
		},
		{
			name: "recursive fail test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8"},
			},
			wantErr: true,
		},
		{
			name: "recursive multiple not test",
			args: args{
				[]string{"!10.10.10.10/32", "!10.10.10.0/24", "!10.10.0.0/16", "!10.0.0.0/8", "0.0.0.0/0"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCIDRs(tt.args.cidrs); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCIDRs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
