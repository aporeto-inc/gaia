package x509extensions

import "encoding/asn1"

var (
	x509ExtensionIdentityTags               asn1.ObjectIdentifier
	x509ExtensionController                 asn1.ObjectIdentifier
	x509ExtensionMaximumIssuedTokenValidity asn1.ObjectIdentifier
)

func init() {
	x509ExtensionIdentityTags = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 1}
	x509ExtensionController = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 2}
	x509ExtensionMaximumIssuedTokenValidity = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 50798, 1, 3}
}

// clone an asn1.ObjectIdentifier
func clone(in asn1.ObjectIdentifier) asn1.ObjectIdentifier {
	clone := make([]int, len(in))
	copy(clone, in)
	return clone
}

// IdentityTags returns the OID for identity tags extensions. The extension
// contents are a JSON encoded slice of strings. Each string is in the form
// key=value.
func IdentityTags() asn1.ObjectIdentifier {
	return clone(x509ExtensionIdentityTags)
}

// Controller returns the OID for controller extensions. The extension allows
// encoding a URI of a controller.
func Controller() asn1.ObjectIdentifier {
	return clone(x509ExtensionController)
}

// MaximumIssuedTokenValidity returns the OID for maximum issued token validity
// extension. The extension allows specifying a time duration. The content is a
// string in the form of '30s', '15m', '1h' or '4d'
func MaximumIssuedTokenValidity() asn1.ObjectIdentifier {
	return clone(x509ExtensionMaximumIssuedTokenValidity)
}
