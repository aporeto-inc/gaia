package x509extensions

import (
	"encoding/asn1"
	"reflect"
	"testing"
)

func TestTags(t *testing.T) {
	tests := []struct {
		name    string
		clobber bool
		want    asn1.ObjectIdentifier
	}{
		{
			name:    "sanity",
			clobber: false,
			want:    x509ExtensionTags,
		},
		{
			name:    "check constant",
			clobber: true,
			want:    x509ExtensionTags,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.clobber {
				c := Tags()
				c[0] = 0
			}
			if got := Tags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController(t *testing.T) {
	tests := []struct {
		name    string
		clobber bool
		want    asn1.ObjectIdentifier
	}{
		{
			name:    "sanity",
			clobber: false,
			want:    x509ExtensionController,
		},
		{
			name:    "check constant",
			clobber: true,
			want:    x509ExtensionController,
		},
	}
	for _, tt := range tests {
		if tt.clobber {
			c := Controller()
			c[0] = 0
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Controller(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller() = %v, want %v", got, tt.want)
			}
		})
	}
}
