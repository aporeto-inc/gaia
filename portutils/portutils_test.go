package portutils

import (
	"reflect"
	"testing"
)

func Test_ConvertToPortsRange(t *testing.T) {
	type args struct {
		ports string
	}
	tests := []struct {
		name    string
		args    args
		want    *PortsRange
		wantErr bool
	}{
		{
			"valid range",
			args{"80:443"},
			&PortsRange{
				FromPort: 80,
				ToPort:   443,
			},
			false,
		},
		{
			"missing port",
			args{"80:"},
			nil,
			true,
		},
		{
			"invalid left bound port",
			args{"something:80"},
			nil,
			true,
		},
		{
			"invalid right bound port",
			args{"80:something"},
			nil,
			true,
		},
		{
			"left bound greather than right bound",
			args{"443:80"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToPortsRange(tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToPortsRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToPortsRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertToPortsList(t *testing.T) {
	type args struct {
		ports string
	}
	tests := []struct {
		name    string
		args    args
		want    *PortsList
		wantErr bool
	}{
		{
			"single port",
			args{"80"},
			&PortsList{
				80,
			},
			false,
		},
		{
			"mutliple ports",
			args{"80,443,1200"},
			nil,
			true,
		},
		{
			"empty ports",
			args{""},
			nil,
			true,
		},
		{
			"invalid single port",
			args{"something"},
			nil,
			true,
		},
		{
			"missing port",
			args{"80,"},
			nil,
			true,
		},
		{
			"multiple port with invalid one",
			args{"80,something"},
			nil,
			true,
		},
		{
			"multiple port with existing one",
			args{"80,80"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToPortsList(tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToPortsList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToPortsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertToSinglePort(t *testing.T) {
	type args struct {
		ports string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"single port",
			args{"80"},
			80,
			false,
		},
		{
			"empty ports",
			args{""},
			-1,
			true,
		},
		{
			"invalid single port",
			args{"something"},
			-1,
			true,
		},
		{
			"single port lower than 1",
			args{"0"},
			-1,
			true,
		},
		{
			"single port greater than 65535",
			args{"65536"},
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToSinglePort(tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToSinglePort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToSinglePort() = %v, want %v", got, tt.want)
			}
		})
	}
}
