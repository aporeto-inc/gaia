package namespaces

import (
	"fmt"
	"testing"

	"github.com/aporeto-inc/gaia/v1/golang"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/manipulate"
	"github.com/aporeto-inc/manipulate/maniptest"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNamespaces_NamespaceValidation(t *testing.T) {

	Convey("Given I have a good namespace name", t, func() {

		ns := "asdasdasdasd"

		Convey("When I validate it", func() {

			errs := ValidateNamespaceStrings(ns)

			Convey("Then the errs should be nil", func() {
				So(errs, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a namespace name containing a *", t, func() {

		ns := "asdas*dasdasd"

		Convey("When I validate it", func() {

			errs := ValidateNamespaceStrings(ns)

			Convey("Then the errs should be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a namespace name containing a =", t, func() {

		ns := "asdasd=asdasd"

		Convey("When I validate it", func() {

			errs := ValidateNamespaceStrings(ns)

			Convey("Then the errs should be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a namespace name containing a * and a =", t, func() {

		ns := "as*dasd=asdasd"

		Convey("When I validate it", func() {

			errs := ValidateNamespaceStrings(ns)

			Convey("Then the errs should be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a namespace name containing a  /", t, func() {

		ns := "asdasd/sdasd"

		Convey("When I validate it", func() {

			errs := ValidateNamespaceStrings(ns)

			Convey("Then the errs should be nil", func() {
				So(errs, ShouldNotBeNil)
			})
		})
	})
}

func TestNamespaces_NamespaceAncestorsNames(t *testing.T) {

	Convey("Given I have a namespace", t, func() {

		ns := "/hello/world/wesh/ta/vu"

		Convey("When I call NamespaceAncestorsNames", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("Then the array should have 5 elements", func() {
				So(len(nss), ShouldEqual, 5)
			})

			Convey("Then the first namespace should be correct", func() {
				So(nss[0], ShouldEqual, "/hello/world/wesh/ta")
			})

			Convey("Then the second namespace should be correct", func() {
				So(nss[1], ShouldEqual, "/hello/world/wesh")
			})

			Convey("Then the third namespace should be correct", func() {
				So(nss[2], ShouldEqual, "/hello/world")
			})

			Convey("Then the fourth namespace should be correct", func() {
				So(nss[3], ShouldEqual, "/hello")
			})

			Convey("Then the fifth namespace should be correct", func() {
				So(nss[4], ShouldEqual, "/")
			})
		})
	})

	Convey("Given I have a / namespace", t, func() {

		ns := "/"

		Convey("When I call NamespaceAncestorsNames", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("Then the array should have 0 elements", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})
	})
}

func TestNamespaces_ParentNamespaceFromString(t *testing.T) {

	Convey("Given I have a namespace", t, func() {

		ns := "/hello/world"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "/hello")
			})
		})
	})

	Convey("Given I have a / namespace", t, func() {

		ns := "/"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "")
			})
		})
	})

	Convey("Given I have a bad namespace", t, func() {

		ns := "asdasdasd"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "")
			})
		})
	})
}

func TestNamespaces_DescendentsOfNamespace(t *testing.T) {

	Convey("Given I have a namespace and a manipulator", t, func() {

		ns := gaia.NewNamespace()
		ns.Name = "/"
		ns.ID = "0"
		m := maniptest.NewTestManipulator()

		Convey("When I call DescendentsOfNamespace", func() {

			m.MockRetrieveMany(t, func(ctx *manipulate.Context, dest elemental.ContentIdentifiable) error {

				switch ctx.Filter.Values()[0][0].(string) {

				case "/":
					*dest.(*gaia.NamespacesList) = append(*dest.(*gaia.NamespacesList),
						&gaia.Namespace{Name: "/1", Namespace: ctx.Namespace, ID: "a"},
						&gaia.Namespace{Name: "/2", Namespace: ctx.Namespace, ID: "b"},
					)

				case "/1":
					*dest.(*gaia.NamespacesList) = append(*dest.(*gaia.NamespacesList),
						&gaia.Namespace{Name: "/1/1", Namespace: ctx.Namespace, ID: "c"},
					)

				case "/2":
					*dest.(*gaia.NamespacesList) = append(*dest.(*gaia.NamespacesList),
						&gaia.Namespace{Name: "/2/1", Namespace: ctx.Namespace, ID: "d"},
						&gaia.Namespace{Name: "/2/2", Namespace: ctx.Namespace, ID: "e"},
					)
				}

				_ = dest // shut the linter up
				return nil
			})

			nss, err := DescendentsOfNamespace(m, ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then nss should should have 6 elements", func() {
				So(len(nss), ShouldEqual, 6)
			})

			Convey("Then the namespace should be correct", func() {
				So(nss[0].Name, ShouldEqual, "/")
				So(nss[1].Name, ShouldEqual, "/1")
				So(nss[2].Name, ShouldEqual, "/1/1")
				So(nss[3].Name, ShouldEqual, "/2")
				So(nss[4].Name, ShouldEqual, "/2/1")
				So(nss[5].Name, ShouldEqual, "/2/2")
			})
		})

		Convey("When I call DescendentsOfNamespace but the manipulator returns an error", func() {

			m.MockRetrieveMany(t, func(ctx *manipulate.Context, dest elemental.ContentIdentifiable) error {
				return fmt.Errorf("oops")
			})

			nss, err := DescendentsOfNamespace(m, ns)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then nss should should have 0 elements", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})

		Convey("When I call DescendentsOfNamespace but the manipulator returns an error later in recursion", func() {

			m.MockRetrieveMany(t, func(ctx *manipulate.Context, dest elemental.ContentIdentifiable) error {

				if ctx.Namespace == "/0/1" {
					*dest.(*gaia.NamespacesList) = append(*dest.(*gaia.NamespacesList),
						&gaia.Namespace{Name: "/0/1/1", Namespace: ctx.Namespace, ID: "a"},
						&gaia.Namespace{Name: "/0/1/2", Namespace: ctx.Namespace, ID: "b"},
					)
					return nil
				}

				return fmt.Errorf("oops")
			})

			nss, err := DescendentsOfNamespace(m, ns)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then nss should should have 0 elements", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})
	})
}

func TestNamespace_IsNamespaceChildrenOfNamespace(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/b", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a/b")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /z", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/c", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/b/c", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceChildrenOfNamespace("", "")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestNamespace_IsNamespaceParentOfNamespace(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceParentOfNamespace on /a/b", func() {

			ok := IsNamespaceParentOfNamespace("/a/b", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a", func() {

			ok := IsNamespaceParentOfNamespace("/a", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /", func() {

			ok := IsNamespaceParentOfNamespace("/", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /z", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a/c", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a/b/c", func() {

			ok := IsNamespaceParentOfNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I check if /aa/b is a children of /a", t, func() {

		ok := IsNamespaceChildrenOfNamespace("/aa/b", "/a")

		Convey("Then ok should be false", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("Given I check if /a is a children of /aa/b", t, func() {

		ok := IsNamespaceParentOfNamespace("/a", "/aa/b")

		Convey("Then ok should be false", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("Given I have an empty namespace", t, func() {
		ns := ""

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/a/b")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestNamespace_IsNamespaceRelatedToNamesapce(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/b", func() {

			ok := IsNamespaceRelatedToNamespace("/a/b", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a", func() {

			ok := IsNamespaceRelatedToNamespace("/a", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /", func() {

			ok := IsNamespaceRelatedToNamespace("/", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /z", func() {

			ok := IsNamespaceRelatedToNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/c", func() {

			ok := IsNamespaceRelatedToNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/b/c", func() {

			ok := IsNamespaceRelatedToNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceRelatedToNamespace("", "")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
