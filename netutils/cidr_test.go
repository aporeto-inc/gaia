package netutils

import (
	"testing"
)

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
			name: "basic failure test",
			args: args{
				[]string{"!10.10.10.10/32"},
			},
			wantErr: true,
		},
		{
			name: "basic format failure test",
			args: args{
				[]string{"!10.10.10.10/ss"},
			},
			wantErr: true,
		},
		{
			name: "basic format failure test 2",
			args: args{
				[]string{"!10.10.10/24"},
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
			name: "basic duplication test",
			args: args{
				[]string{"0.0.0.0/0", "10.10.10.0/16", "!10.10.10.0/16"},
			},
			wantErr: true,
		},
		{
			name: "duplication in inclusion test",
			args: args{
				[]string{"0.0.0.0/0", "10.10.10.0/16", "10.10.10.0/16"},
			},
			wantErr: true,
		},
		{
			name: "duplication in exclusion test",
			args: args{
				[]string{"0.0.0.0/0", "!10.10.10.0/16", "10.10.10.0/16"},
			},
			wantErr: true,
		},
		{
			name: "duplication in inclusion with different IP test",
			args: args{
				[]string{"0.0.0.0/0", "10.10.10.2/16", "10.10.10.4/16"},
			},
			wantErr: true,
		},
		{
			name: "duplication in exclusion with different IP test",
			args: args{
				[]string{"0.0.0.0/0", "!10.10.10.2/16", "10.10.10.4/16"},
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
			name: "basic format failure test IPv6",
			args: args{
				[]string{"2001:db8::/ss"},
			},
			wantErr: true,
		},
		{
			name: "basic format failure test 2 IPv6",
			args: args{
				[]string{"2001:db8:/64"},
			},
			wantErr: true,
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
			name: "basic duplication test IPv6",
			args: args{
				[]string{"::/0", "2001:db8::/64", "!2001:db8::/64"},
			},
			wantErr: true,
		},
		{
			name: "duplication test in inclusion IPv6",
			args: args{
				[]string{"::/0", "2001:db8::/64", "2001:db8::/64"},
			},
			wantErr: true,
		},
		{
			name: "duplication test in exclusion IPv6",
			args: args{
				[]string{"::/0", "!2001:db8::/64", "!2001:db8::/64"},
			},
			wantErr: true,
		},
		{
			name: "duplication test in inclusion with different IP IPv6",
			args: args{
				[]string{"::/0", "2001:db8::8888/64", "2001:db8::/64"},
			},
			wantErr: true,
		},
		{
			name: "duplication test in inclusion with different IP IPv6",
			args: args{
				[]string{"::/0", "!2001:db8::8888/64", "!2001:db8::/64"},
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
			name: "both IPv4 and IPv6 IPv4 with IPv4 not contained failure test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8", "2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16", "::/0"},
			},
			wantErr: true,
		},
		{
			name: "both IPv4 and IPv6 IPv4 with IPv6 not contained failure test",
			args: args{
				[]string{"10.10.10.10/32", "!10.10.10.0/24", "10.10.0.0/16", "!10.0.0.0/8", "0.0.0.0/0", "2001:db8::/128", "!2001:db8::/64", "2001:db8::/32", "!2001:db8::/16"},
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
