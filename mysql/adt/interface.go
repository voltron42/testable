package imysql

import (
	"crypto/tls"
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
	"io"
)

type Registry interface {
	DeregisterLocalFile(filePath string)
	DeregisterReaderHandler(name string)
	DeregisterTLSConfig(key string)
	RegisterDial(net string, dial mysql.DialFunc)
	RegisterLocalFile(filePath string)
	RegisterReaderHandler(name string, handler func() io.Reader)
	RegisterTLSConfig(key string, config *tls.Config) error
	SetLogger(logger mysql.Logger) error
}

type MySQLDriver interface {
	Open(dsn string) (driver.Conn, error)
}
