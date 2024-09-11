package schema1

import (
	"crypto/x509"

	"github.com/docker/libtrust"
	"github.com/sirupsen/logrus"
)

// Verify verifies the signature of the signed manifest returning the public
// keys used during signing.
//
// Deprecated: Docker Image Manifest v2, Schema 1 is deprecated since 2015.
// Use Docker Image Manifest v2, Schema 2, or the OCI Image Specification.
func Verify(sm *SignedManifest) ([]libtrust.PublicKey, error) {
	js, err := libtrust.ParsePrettySignature(sm.all, "signatures")
	if err != nil {
		logrus.WithField("err", err).Debugf("(*SignedManifest).Verify")
		return nil, err
	}

	return js.Verify()
}

// VerifyChains verifies the signature of the signed manifest against the
// certificate pool returning the list of verified chains. Signatures without
// an x509 chain are not checked.
//
// Deprecated: Docker Image Manifest v2, Schema 1 is deprecated since 2015.
// Use Docker Image Manifest v2, Schema 2, or the OCI Image Specification.
func VerifyChains(sm *SignedManifest, ca *x509.CertPool) ([][]*x509.Certificate, error) {
	js, err := libtrust.ParsePrettySignature(sm.all, "signatures")
	if err != nil {
		return nil, err
	}

	return js.VerifyChains(ca)
}
