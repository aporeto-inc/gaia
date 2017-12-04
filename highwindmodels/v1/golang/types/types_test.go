package types

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Copy(t *testing.T) {
	Convey("Given a parameter, it's copy should be similar", t, func() {
		parameter := &ServiceParameter{
			Name:            "name",
			Description:     "description",
			LongDescription: "long-description",
			Key:             "key",
			Value:           "value",
			Env:             "ENV",
			Type:            ServiceParameterTypeString,
			AllowedValues:   []interface{}{"A"},
			DefaultValue:    "default-value",
			MountPath:       "/demo",
			Backend:         ServiceParameterBackendLocalConfigMap,
			Optional:        true,
		}

		copy := parameter.Copy()
		So(copy.Name, ShouldEqual, parameter.Name)
		So(copy.Description, ShouldEqual, parameter.Description)
		So(copy.LongDescription, ShouldEqual, parameter.LongDescription)
		So(copy.Key, ShouldEqual, parameter.Key)
		So(copy.Value, ShouldEqual, parameter.Value)
		So(copy.Env, ShouldEqual, parameter.Env)
		So(copy.Type, ShouldEqual, parameter.Type)
		So(copy.AllowedValues, ShouldResemble, parameter.AllowedValues)
		So(copy.DefaultValue, ShouldEqual, parameter.DefaultValue)
		So(copy.MountPath, ShouldEqual, parameter.MountPath)
		So(copy.Backend, ShouldEqual, parameter.Backend)
		So(copy.Optional, ShouldEqual, parameter.Optional)
	})
}

func Test_StringParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: "A",
			Type:  ServiceParameterTypeString,
		}
		So(parameter.ValueToString(), ShouldEqual, "A")
	})
}

func Test_BoolParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: true,
			Type:  ServiceParameterTypeBool,
		}
		So(parameter.ValueToString(), ShouldEqual, "true")
	})
}

func Test_IntParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: 10,
			Type:  ServiceParameterTypeInt,
		}
		So(parameter.ValueToString(), ShouldEqual, "10")
	})
}

func Test_FloatParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: 1.2,
			Type:  ServiceParameterTypeFloat,
		}
		So(parameter.ValueToString(), ShouldEqual, "1.2")
	})
}

func Test_DurationParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: "30s",
			Type:  ServiceParameterTypeDuration,
		}
		So(parameter.ValueToString(), ShouldEqual, "30s")
	})
}

func Test_StringSliceParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: []interface{}{"A", "B"},
			Type:  ServiceParameterTypeStringSlice,
		}
		So(parameter.ValueToString(), ShouldEqual, "A B")
	})
}

func Test_IntSliceParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: []interface{}{10, 20},
			Type:  ServiceParameterTypeIntSlice,
		}
		So(parameter.ValueToString(), ShouldEqual, "10 20")
	})
}

func Test_FloatSliceParameterValueToString(t *testing.T) {
	Convey("Given a parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: []interface{}{1.2, 2.2},
			Type:  ServiceParameterTypeFloatSlice,
		}
		So(parameter.ValueToString(), ShouldEqual, "1.2 2.2")
	})
}

func Test_EnumParameterValueToString(t *testing.T) {
	Convey("Given a string parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:          "DemoParam",
			Value:         "A",
			Type:          ServiceParameterTypeEmum,
			AllowedValues: []interface{}{"A", "B", "C"},
		}
		So(parameter.ValueToString(), ShouldEqual, "A")
	})

	Convey("Given an int parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:          "DemoParam",
			Value:         10,
			Type:          ServiceParameterTypeEmum,
			AllowedValues: []interface{}{10, 11, 12},
		}
		So(parameter.ValueToString(), ShouldEqual, "10")
	})

	Convey("Given a bool parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:          "DemoParam",
			Value:         true,
			Type:          ServiceParameterTypeEmum,
			AllowedValues: []interface{}{true, false},
		}
		So(parameter.ValueToString(), ShouldEqual, "true")
	})

	Convey("Given a float parameter with a correct value, it should return a string", t, func() {
		parameter := &ServiceParameter{
			Name:          "DemoParam",
			Value:         2.1,
			Type:          ServiceParameterTypeEmum,
			AllowedValues: []interface{}{2.1, 2.2, 2.3},
		}
		So(parameter.ValueToString(), ShouldEqual, "2.1")
	})
}
