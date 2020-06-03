package x509extensions

import "encoding/asn1"

var (
	// x509ExtensionTags represents an extension to encode tags. The contents are a
	// JSON encoded slice of strings. Each string is in the form key=value.
	x509ExtensionTags asn1.ObjectIdentifier
	// x509ExtensionController represents an extension to encode a URI of a controller.
	x509ExtensionController asn1.ObjectIdentifier
)

func init() {
	x509ExtensionTags = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1}
	x509ExtensionController = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 2}
}

// clone an asn1.ObjectIdentifier
func clone(in asn1.ObjectIdentifier) asn1.ObjectIdentifier {
	clone := make([]int, len(in))
	copy(clone, in)
	return clone
}

// Tags returns the OID for tags extensions
func Tags() asn1.ObjectIdentifier {
	return clone(x509ExtensionTags)
}

// Controller returns the OID for controller extensions
func Controller() asn1.ObjectIdentifier {
	return clone(x509ExtensionController)
}
