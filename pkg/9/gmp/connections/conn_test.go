package connections

import (
	"encoding/xml"
	"net"
	"testing"
)

func TestNewConn(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := New(c1)

	go func() {
		buf2 := make([]byte, 150000)
		nRead, _ := c2.Read(buf2)
		c2.Write(buf2[:nRead])
	}()

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

	err := conn.Execute(cmd, response)
	if err != nil {
		t.Fatalf("Unexpected error during Execute: %s", err)
	}

	if response.Foo != expectedValue {
		t.Fatalf("Unexpected response value: %s\nExpected: %s", response.Foo, expectedValue)
	}
}
