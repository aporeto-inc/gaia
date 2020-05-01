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
				prefixes: []string{"0.0.0.0/0"},
				ip:       "0.0.0.0/0",
			},
			want: true,
		},
		{
			name: "basic test 2",
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
		{
			name: "cidr not contained test failure",
			args: args{
				prefixes: []string{"10.10.0.0/16"},
				ip:       "10.10.0.0/8",
			},
			want: false,
		},
		{
			name: "basic test IPv6",
			args: args{
				prefixes: []string{"::/0"},
				ip:       "::/0",
			},
			want: true,
		},
		{
			name: "basic test IPv6 2",
			args: args{
				prefixes: []string{"2001:db8::/64", "::/0"},
				ip:       "2001:db8::/128",
			},
			want: true,
		},
		{
			name: "basic test failure IPv6",
			args: args{
				prefixes: []string{},
				ip:       "2001:db8::/128",
			},
			want: false,
		},
		{
			name: "multiple match test IPv6",
			args: args{
				prefixes: []string{"2001:db8::/64", "2001:db8::/32", "::/0"},
				ip:       "2001:db8::/128",
			},
			want: true,
		},
		{
			name: "cidr not contained test failure IPv6",
			args: args{
				prefixes: []string{"2001:db8::/64"},
				ip:       "2001:db8::/32",
			},
			want: false,
		},
		{
			name: "mix of both IPv4 and IPv6 test",
			args: args{
				prefixes: []string{"2001:db8::/64", "2001:db8::/32", "::/0", "10.10.10.0/24", "10.10.0.0/16", "0.0.0.0/0"},
				ip:       "2001:db8::/128",
			},
			want: true,
		},
		{
			name: "IPv6 not contained in IPv4 cidrs failure test",
			args: args{
				prefixes: []string{"10.10.10.0/24", "10.10.0.0/16", "0.0.0.0/0"},
				ip:       "2001:db8::/128",
			},
			want: false,
		},
		{
			name: "IPv4 not contained in IPv6 cidrs failure test",
			args: args{
				prefixes: []string{"2001:db8::/64", "2001:db8::/32", "::/0"},
				ip:       "10.10.10.10/32",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prefixesV4, prefixesV6, err := parseCIDRs(tt.args.prefixes)
			if err != nil {
				t.Errorf("err in test: %v", err)
			}
			ip, network, err := net.ParseCIDR(tt.args.ip)
			if err != nil {
				t.Errorf("err in test: %v", err)
			}

			if ipv4 := ip.To4(); ipv4 != nil {
				if got := prefixIsContained(prefixesV4, network); got != tt.want {
					t.Errorf("prefixIsContained() = %v, want %v", got, tt.want)
				}
			} else if ipv6 := ip.To16(); ipv6 != nil {
				if got := prefixIsContained(prefixesV6, network); got != tt.want {
					t.Errorf("prefixIsContained() = %v, want %v", got, tt.want)
				}
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
			name: "basic cidr not contained failure test",
			args: args{
				[]string{"10.10.10.0/16", "!10.10.10.10/8"},
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
		{
			name: "basic test IPv6",
			args: args{
				[]string{"2001:db8::/128"},
			},
			wantErr: false,
		},
		{
			name: "basic inclusion test IPv6",
			args: args{
				[]string{"2001:db8::/64", "!2001:db8::/128"},
			},
			wantErr: false,
		},
		{
			name: "basic failure test IPv6",
			args: args{
				[]string{"!2001:db8::/128"},
			},
			wantErr: true,
		},
		{
			name: "basic cidr not contained failure test IPv6",
			args: args{
				[]string{"2001:db8::/64", "!2001:db8::/32"},
			},
			wantErr: true,
		},
		{
			name: "recursive test IPv6",
			args: args{
				[]string{"2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16", "::/0"},
			},
			wantErr: false,
		},
		{
			name: "recursive fail test IPv6",
			args: args{
				[]string{"2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16"},
			},
			wantErr: true,
		},
		{
			name: "recursive multiple not test IPv6",
			args: args{
				[]string{"!2001:db8::/128", "!2001:db8::/64", "!2001:db8::/32", "!2001:db8::/16", "::/0"},
			},
			wantErr: false,
		},
		{
			name: "both IPv4 and IPv6 test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8", "0.0.0.0/0", "2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16", "::/0"},
			},
			wantErr: false,
		},
		{
			name: "both IPv4 and IPv6 IPv4 not contained failure test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8", "2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16", "::/0"},
			},
			wantErr: true,
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
