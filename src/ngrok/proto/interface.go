package proto

import (
	"github.com/grokker001/ngrok/src/ngrok/conn"
)

type Protocol interface {
	GetName() string
	WrapConn(conn.Conn, interface{}) conn.Conn
}
