package connections

import (
	"crypto/tls"

	"github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp"
	"github.com/filewalkwithme/go-openvas-gmp/pkg/9/gmp/connections/internal/implementation"
)

// NewTLSConnection returns an instance of `gmp.Connection` that uses a TLS connection as underlying trasport to communicate with Openvas GVMD.
// The `address` parameter refers to the host and port where Openvas GVMD is listening.
// The `insecure` parameter allows the client to accept invalid certificates (ex: self-signed)
func NewTLSConnection(address string, insecure bool) (gmp.Connection, error) {
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: insecure,
	})
	if err != nil {
		return nil, err
	}

	c := &implementation.Connection{}
	c.SetRawConn(conn)

	return c, nil
}
