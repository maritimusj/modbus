package modbus

import "fmt"

type Error interface {
	error
	Temporary() bool // Is the error temporary?
}

type TcpError struct {
	err string
}

func NewModbusProtocolError(err string, args ...interface{}) *TcpError {
	return &TcpError{
		err: fmt.Sprintf(err, args...),
	}
}

func (e *TcpError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.err
}

func (e *TcpError) Temporary() bool {
	return true
}
