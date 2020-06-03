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
			want:    asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1},
		},
		{
			name:    "check constant",
			clobber: true,
			want:    asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1},
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
			want:    asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 2},
		},
		{
			name:    "check constant",
			clobber: true,
			want:    asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 2},
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
