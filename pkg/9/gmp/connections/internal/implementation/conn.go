package implementation

import (
	"bytes"
	"encoding/xml"
	"net"
)

type Connection struct {
	rawConn net.Conn
}

func (conn *Connection) SetRawConn(rawConn net.Conn) {
	conn.rawConn = rawConn
}

func (conn *Connection) performRequest(buffer []byte) ([]byte, error) {
	_, err := conn.rawConn.Write(buffer)
	if err != nil {
		return nil, err
	}

	var resp bytes.Buffer
	for true {
		buf := make([]byte, 1024)
		n, err := conn.rawConn.Read(buf[:])
		if err != nil {
			return nil, err
		}
		resp.Write(buf[:n])

		if n < 1024 {
			break
		}
	}

	return resp.Bytes(), nil
}

func (conn *Connection) Execute(command interface{}, response interface{}) error {
	cmdBuf, err := xml.Marshal(command)
	if err != nil {
		return err
	}

	resp, err := conn.performRequest(cmdBuf)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(resp, response)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Connection) Close() error {
	return conn.rawConn.Close()
}
