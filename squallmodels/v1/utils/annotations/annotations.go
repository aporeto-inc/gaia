package annotations

// AnnotationSetter is an interface that allows to set annotations
type AnnotationSetter interface {
	SetAnnotations(map[string][]string)
}

// AnnotationsGetter is an interface that allows to get a annotations
type AnnotationsGetter interface { // nolint
	GetAnnotations() map[string][]string
}
