package connections

import (
	"crypto/tls"
	"encoding/xml"
	"log"
	"net"
	"testing"
)

var certPem = []byte(`-----BEGIN CERTIFICATE-----
MIIB/jCCAaigAwIBAgIJAOMDItx1P3oFMA0GCSqGSIb3DQEBCwUAMFkxCzAJBgNV
BAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX
aWRnaXRzIFB0eSBMdGQxEjAQBgNVBAMMCWxvY2FsaG9zdDAeFw0yMDA3MjMyMjIw
MjNaFw00NzEyMDkyMjIwMjNaMFkxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21l
LVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxEjAQBgNV
BAMMCWxvY2FsaG9zdDBcMA0GCSqGSIb3DQEBAQUAA0sAMEgCQQDNi9Y0Hnq0yFFV
e5XOEKyETjBr1AWuEe39qRN/cae+LNg0oEllqfiC2IyKqYnFyUM9ggWS+tyLeoSc
mL91zFODAgMBAAGjUzBRMB0GA1UdDgQWBBQIqHYmwur3ZB2ZYcTwnFA4LhA0RTAf
BgNVHSMEGDAWgBQIqHYmwur3ZB2ZYcTwnFA4LhA0RTAPBgNVHRMBAf8EBTADAQH/
MA0GCSqGSIb3DQEBCwUAA0EAVx29PTXB+UqdtxKyf7g8UGzW8pXbZghR5pV/Wux/
cE2YLFey9ZwyHfmS1Fkd75wrRb2yfZSvFD/nMWNT1wOuuQ==
-----END CERTIFICATE-----`)

var keyPem = []byte(`-----BEGIN PRIVATE KEY-----
MIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAzYvWNB56tMhRVXuV
zhCshE4wa9QFrhHt/akTf3GnvizYNKBJZan4gtiMiqmJxclDPYIFkvrci3qEnJi/
dcxTgwIDAQABAkB+b2sGykzekW3+cDPY+saz58i/Oz93MM49P1igB1CQmLj71zGt
jCCNlfPlNtzqIFuBzPb+IJk+stoMXp16d9XpAiEA5koJRlSOEMSRxBfqyFu3rKwn
dAdUnD3y6ltkY+aEN80CIQDkfp1NRZbVeaofyhA8SvnSPzqFFEqm5ySfNN/WBL3I
jwIgGuNiGdgdjk+lRWQVgbdTxtGC+cUxV9zT1BE/s3pizbECIQCVb00XFTTxRlGB
2tfFZs99tkZidIPiJfcofB8LzCwGdwIgIpI2tLScYVIzHfzfzn0U5mflhF5W6WIF
BZo0rFh9Ir0=
-----END PRIVATE KEY-----`)

func newLocalListener(t testing.TB) net.Listener {
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS11,
		MaxVersion:         tls.VersionTLS11,
	})
	if err != nil {
		ln, err = tls.Listen("tcp6", "[::1]:0", nil)
	}
	if err != nil {
		t.Fatal(err)
	}
	return ln
}

func newLocalListenerProtocolNotSupported(t testing.TB) net.Listener {
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionSSL30,
		MaxVersion:         tls.VersionSSL30,
	})
	if err != nil {
		ln, err = tls.Listen("tcp6", "[::1]:0", nil)
	}
	if err != nil {
		t.Fatal(err)
	}
	return ln
}

func TestNewTLSConnection(t *testing.T) {
	ln := newLocalListener(t)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				t.Fatalf("Unexpected error during Accept: %s", err)
			}
			buf2 := make([]byte, 150000)
			nRead, _ := conn.Read(buf2)
			conn.Write(buf2[:nRead])
		}
	}()

	conn, err := NewTLSConnection(ln.Addr().String(), true)
	if err != nil {
		t.Fatalf("Unexpected error during NewTLSConnection: %s", err)
	}

	expectedValue := "0123456789"
	cmd := &struct {
		XMLName xml.Name `xml:"cmd"`
		Foo     string   `xml:"foo"`
	}{
		Foo: expectedValue,
	}
	response := &struct {
		Foo string `xml:"foo"`
	}{}

	err = conn.Execute(cmd, response)
	if err != nil {
		t.Fatalf("Unexpected error during Execute: %s", err)
	}

	if response.Foo != expectedValue {
		t.Fatalf("Unexpected response value: %s\nExpected: %s", response.Foo, expectedValue)
	}
}

func TestNewTLSConnectionFail(t *testing.T) {
	ln := newLocalListenerProtocolNotSupported(t)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				t.Fatalf("Unexpected error during Accept: %s", err)
			}
			buf2 := make([]byte, 150000)
			nRead, _ := conn.Read(buf2)
			conn.Write(buf2[:nRead])
		}
	}()

	_, err := NewTLSConnection(ln.Addr().String(), true)

	expectedError := "remote error: tls: protocol version not supported"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during Execute.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}
