package types

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_StringParameterValidation(t *testing.T) {
	Convey("Given a parameter that is valid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: "demo-value",
			Type:  ServiceParameterTypeString,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with no value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: nil,
			Type:  ServiceParameterTypeString,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})

	Convey("Given an optional parameter with no value, It should be valid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    nil,
			Type:     ServiceParameterTypeString,
			Optional: true,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with a non string value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    10,
			Type:     ServiceParameterTypeString,
			Optional: true,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})
}

func Test_IntParameterValidation(t *testing.T) {
	Convey("Given a parameter that is valid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: 12,
			Type:  ServiceParameterTypeInt,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with no value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: nil,
			Type:  ServiceParameterTypeInt,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})

	Convey("Given an optional parameter with no value, It should be valid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    nil,
			Type:     ServiceParameterTypeInt,
			Optional: true,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with a non string value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    "10",
			Type:     ServiceParameterTypeInt,
			Optional: true,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})
}

func Test_BoolParameterValidation(t *testing.T) {
	Convey("Given a parameter that is valid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: true,
			Type:  ServiceParameterTypeBool,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with no value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: nil,
			Type:  ServiceParameterTypeBool,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})

	Convey("Given an optional parameter with no value, It should be valid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    nil,
			Type:     ServiceParameterTypeBool,
			Optional: true,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with a non string value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    10,
			Type:     ServiceParameterTypeBool,
			Optional: true,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})
}

func Test_DurationParameterValidation(t *testing.T) {
	Convey("Given a parameter that is valid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: "15s",
			Type:  ServiceParameterTypeDuration,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a parameter that is invalid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: "sixHours",
			Type:  ServiceParameterTypeDuration,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})

	Convey("Given a required parameter with no value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:  "DemoParam",
			Value: nil,
			Type:  ServiceParameterTypeDuration,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})

	Convey("Given an optional parameter with no value, It should be valid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    nil,
			Type:     ServiceParameterTypeDuration,
			Optional: true,
		}
		So(parameter.Validate(), ShouldBeNil)
	})

	Convey("Given a required parameter with a non string value, It should be invalid", t, func() {
		parameter := &ServiceParameter{
			Name:     "DemoParam",
			Value:    10,
			Type:     ServiceParameterTypeDuration,
			Optional: true,
		}
		So(parameter.Validate(), ShouldNotBeNil)
	})
}
