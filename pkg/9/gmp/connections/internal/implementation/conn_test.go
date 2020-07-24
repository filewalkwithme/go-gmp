package implementation

import (
	"encoding/xml"
	"net"
	"sync"
	"testing"
)

func TestSetRawConn(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	var wg sync.WaitGroup
	var nRead int
	var errRead error

	wg.Add(1)
	go func() {
		buf := make([]byte, 1024)
		nRead, errRead = c2.Read(buf)
		wg.Done()
	}()

	n, err := conn.rawConn.Write([]byte("0123456789"))
	if err != nil {
		t.Fatalf("Unexpected error during write: %s", err)
	}

	if n != 10 {
		t.Fatalf("Expected to write 10 bytes, wrote: %d", n)
	}

	wg.Wait()

	if errRead != nil {
		t.Fatalf("Unexpected error during read: %s", err)
	}

	if nRead != 10 {
		t.Fatalf("Expected to read 10 bytes, got: %d", nRead)
	}

	errClose := conn.Close()
	if errClose != nil {
		t.Fatalf("Unexpected error during Close: %s", errClose)
	}
}

func TestPerformRequest(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	go func() {
		buf2 := make([]byte, 150000)
		nRead, _ := c2.Read(buf2)
		c2.Write(buf2[:nRead])
	}()

	buf := make([]byte, 50000)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 255)
	}

	bufRead, err := conn.performRequest(buf)
	if err != nil {
		t.Fatalf("Unexpected error during write: %s", err)
	}

	if len(bufRead) != 50000 {
		t.Fatalf("Expected to write 50000 bytes, wrote: %d", len(bufRead))
	}

	for i := 0; i < 50000; i++ {
		if bufRead[i] != byte(i%255) {
			t.Fatalf("Incorrect value at position %d: %d", i, bufRead[i])
		}
	}
}

func TestPerformRequestWriteFail(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	c1.Close()
	go func() {
		buf2 := make([]byte, 150000)
		nRead, _ := c2.Read(buf2)
		c2.Write(buf2[:nRead])
	}()

	buf := make([]byte, 50000)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 255)
	}

	_, err := conn.performRequest(buf)
	expectedError := "io: read/write on closed pipe"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during performRequest.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}

func TestPerformRequestReadFail(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	go func() {
		buf2 := make([]byte, 150000)
		c2.Read(buf2)
		c2.Close()
	}()

	buf := make([]byte, 50000)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 255)
	}

	_, err := conn.performRequest(buf)
	expectedError := "EOF"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during performRequest.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}

func TestExecute(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

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

func TestExecuteXMLMarshallFail(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	go func() {
		buf2 := make([]byte, 150000)
		nRead, _ := c2.Read(buf2)
		c2.Write(buf2[:nRead])
	}()

	cmd := func() {}
	response := &struct {
		Foo string `xml:"foo"`
	}{}

	err := conn.Execute(cmd, response)
	expectedError := "xml: unsupported type: func()"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during Execute.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}

func TestExecuteXMLUnmarshallFail(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

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
	response := struct {
		Foo string `xml:"foo"`
	}{}

	err := conn.Execute(cmd, response)
	expectedError := "non-pointer passed to Unmarshal"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during Execute.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}

func TestExecutePerformRequestFail(t *testing.T) {
	c1, c2 := net.Pipe()

	conn := Connection{}
	conn.SetRawConn(c1)

	c1.Close()
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
	expectedError := "io: read/write on closed pipe"
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Unexpected error during Execute.\nExpected: %s\n     Got: %s", expectedError, err)
	}
}
