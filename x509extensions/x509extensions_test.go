package x509extensions

import (
	"encoding/asn1"
	"reflect"
	"testing"
)

func Test_clone(t *testing.T) {
	type args struct {
		in asn1.ObjectIdentifier
	}
	tests := []struct {
		name string
		args args
		want asn1.ObjectIdentifier
	}{
		{
			name: "sanity",
			args: args{
				in: asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1},
			},
			want: asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clone(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() = %v, want %v", got, tt.want)
			}
			// Modify got and make sure input was not clobbered.
			got[0] = 0
			if !reflect.DeepEqual(tt.args.in, tt.want) {
				t.Errorf("clone() clobber detected = %v, want %v", tt.args.in, tt.want)
			}
		})
	}
}

func TestIdentityTags(t *testing.T) {
	tests := []struct {
		name string
		want asn1.ObjectIdentifier
	}{
		{
			name: "sanity",
			want: asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdentityTags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController(t *testing.T) {
	tests := []struct {
		name string
		want asn1.ObjectIdentifier
	}{
		{
			name: "sanity",
			want: asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Controller(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumIssuedTokenValidity(t *testing.T) {
	tests := []struct {
		name string
		want asn1.ObjectIdentifier
	}{
		{
			name: "sanity",
			want: asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumIssuedTokenValidity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaximumIssuedTokenValidity() = %v, want %v", got, tt.want)
			}
		})
	}
}
