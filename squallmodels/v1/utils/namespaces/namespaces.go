package namespaces

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang"
	"github.com/aporeto-inc/manipulate"
)

// NamespaceSetter is an interface that allows to set a namespace
type NamespaceSetter interface {
	SetNamespace(string)
}

// NamespaceGetter is an interface that allows to get a namespace
type NamespaceGetter interface {
	GetNamespace() string
}

const nsSeparator = "/"

// ParentNamespaceFromString returns the parent namespace of a namespace
// It returns empty it the string is invalid
func ParentNamespaceFromString(namespace string) (string, error) {

	if namespace == nsSeparator {
		return "", nil
	}

	index := strings.LastIndex(namespace, nsSeparator)

	switch index {
	case -1:
		return "", fmt.Errorf("Invalid namespace name")
	case 0:
		return namespace[:index+1], nil
	default:
		return namespace[:index], nil
	}
}

// IsNamespaceRelatedToNamespace returns true if the given namespace is related to the given parent
func IsNamespaceRelatedToNamespace(ns string, parent string) bool {
	return IsNamespaceParentOfNamespace(ns, parent) || IsNamespaceChildrenOfNamespace(ns, parent) || (ns == parent && ns != "" && parent != "")
}

// IsNamespaceParentOfNamespace returns true if the given namespace is a parent of the given parent
func IsNamespaceParentOfNamespace(ns string, child string) bool {

	if len(ns) == 0 {
		return false
	}

	if ns == child {
		return false
	}

	if ns[len(ns)-1] != '/' {
		ns = ns + "/"
	}

	return strings.HasPrefix(child, ns)
}

// IsNamespaceChildrenOfNamespace returns true of the given ns is a children of the given parent.
func IsNamespaceChildrenOfNamespace(ns string, parent string) bool {

	if len(parent) == 0 {
		return false
	}

	if ns == parent {
		return false
	}

	if parent[len(parent)-1] != '/' {
		parent = parent + "/"
	}

	return strings.HasPrefix(ns, parent)
}

// NamespaceAncestorsNames returns the list of fully qualified namespaces
// in the hierarchy of a given namespace. It returns an empty
// array for the root namespace
func NamespaceAncestorsNames(namespace string) []string {

	if namespace == nsSeparator {
		return []string{}
	}

	parts := strings.Split(namespace, nsSeparator)
	sep := nsSeparator
	namespaces := []string{}

	for i := len(parts) - 1; i >= 2; i-- {
		namespaces = append(namespaces, sep+strings.Join(parts[1:i], sep))
	}

	namespaces = append(namespaces, sep)

	return namespaces
}

// DescendentsOfNamespace retrieves the descendents of the given namespace using the given manipulator.
func DescendentsOfNamespace(manipulator manipulate.Manipulator, namespace *squallmodels.Namespace) (squallmodels.NamespacesList, error) {

	out := squallmodels.NamespacesList{namespace}
	namespaces := squallmodels.NamespacesList{}

	mctx := manipulate.NewContextWithFilter(
		manipulate.NewFilterComposer().
			WithKey("namespace").Equals(namespace.Name).
			Done(),
	)

	if err := manipulator.RetrieveMany(mctx, &namespaces); err != nil {
		return nil, err
	}

	for _, n := range namespaces {

		subs, err := DescendentsOfNamespace(manipulator, n)
		if err != nil {
			return nil, err
		}

		out = append(out, subs...)
	}

	return out, nil
}

// AscendentsOfNamespace returns the list of namespace object that are parent of the given namespace.
func AscendentsOfNamespace(manipulator manipulate.Manipulator, namespace *squallmodels.Namespace) (squallmodels.NamespacesList, error) {

	names := NamespaceAncestorsNames(namespace.Name)
	filter := manipulate.NewFilterComposer()
	for i, name := range names {
		if i == 0 {
			filter.WithKey("name").Equals(name)
		} else {
			filter.OrKey("name").Equals(name)
		}
	}

	mctx := manipulate.NewContextWithFilter(filter.Done())
	nss := squallmodels.NamespacesList{}
	if err := manipulator.RetrieveMany(mctx, &nss); err != nil {
		return nil, err
	}

	return nss, nil
}

// NamespaceByName returns the namespace with the given name.
func NamespaceByName(manipulator manipulate.Manipulator, name string) (*squallmodels.Namespace, error) {

	mctx := manipulate.NewContextWithFilter(
		manipulate.NewFilterComposer().
			WithKey("name").Equals(name).
			Done(),
	)

	nslist := squallmodels.NamespacesList{}
	if err := manipulator.RetrieveMany(mctx, &nslist); err != nil {
		return nil, err
	}

	switch len(nslist) {
	case 0:
		return nil, manipulate.NewErrObjectNotFound("Cannot find object with the given ID")
	case 1:
		return nslist[0], nil
	default:
		return nil, fmt.Errorf("Found more than one namespace named %s", name)
	}
}

// ValidateNamespaceStrings validates the name of the given namespaces.
func ValidateNamespaceStrings(namespaces ...string) error {

	errs := elemental.Errors{}

L:
	for _, namespace := range namespaces {

		for _, r := range namespace {
			if r == '/' {
				errs = append(errs, elemental.NewError("Reserved Character", "Namespace name cannot contains a /", "gaia", http.StatusUnprocessableEntity))
				continue L
			}
		}

		ns := &squallmodels.Namespace{Name: namespace}
		if err := ns.Validate(); err != nil {
			if e, ok := err.(elemental.Errors); ok {
				errs = append(errs, e...)
			} else {
				errs = append(errs, e)
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// ValidateNamespaceBump validate that the given namespace is valid for a bump
func ValidateNamespaceBump(m manipulate.Manipulator, policyNamespace string, requestNamespace string) error {

	data := map[string]interface{}{"attribute": "mappedNamespace"}

	if policyNamespace == requestNamespace {
		// Uncomment this once apotests is ready
		// return errors.New("Invalid mapped namespace", "You cannot map a processing unit to the current namespace", http.StatusUnprocessableEntity, data)
		return nil
	}

	if IsNamespaceParentOfNamespace(policyNamespace, requestNamespace) {
		err := elemental.NewError("Invalid mapped namespace", "You cannot map a processing unit to a higher level namespace", "gaia", http.StatusUnprocessableEntity)
		err.Data = data
		return err
	}

	if !IsNamespaceChildrenOfNamespace(policyNamespace, requestNamespace) {
		err := elemental.NewError("Invalid mapped namespace", "You cannot map a processing unit to this level namespace, it needs to be a lower namespace", "gaia", http.StatusUnprocessableEntity)
		err.Data = data
		return err
	}

	mctx := manipulate.NewContextWithFilter(manipulate.NewFilterComposer().WithKey("name").Equals(policyNamespace).Done())
	count, err := m.Count(mctx, squallmodels.NamespaceIdentity)
	if err != nil {
		return err
	}

	if count == 0 {
		err := elemental.NewError("Invalid mapped namespace", "The mapped namespace doesn't exist", "gaia", http.StatusUnprocessableEntity)
		err.Data = data
		return err
	}

	return nil
}
