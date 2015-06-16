package mysql

import (
	"crypto/tls"
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
	"io"

	"../adt"
)

func NewRegistry() imysql.Registry {
	return &Registry{}
}

type Registry struct{}

func (r *Registry) DeregisterLocalFile(filePath string) {
	mysql.DeregisterLocalFile(filePath)
}

func (r *Registry) DeregisterReaderHandler(name string) {
	mysql.DeregisterReaderHandler(name)
}

func (r *Registry) DeregisterTLSConfig(key string) {
	mysql.DeregisterTLSConfig(key)
}

func (r *Registry) RegisterDial(net string, dial mysql.DialFunc) {
	mysql.RegisterDial(net, dial)
}

func (r *Registry) RegisterLocalFile(filePath string) {
	mysql.RegisterLocalFile(filePath)
}

func (r *Registry) RegisterReaderHandler(name string, handler func() io.Reader) {
	mysql.RegisterReaderHandler(name, handler)
}

func (r *Registry) RegisterTLSConfig(key string, config *tls.Config) error {
	return mysql.RegisterTLSConfig(key, config)
}

func (r *Registry) SetLogger(logger mysql.Logger) error {
	return mysql.SetLogger(logger)
}

func NewMySQLDriver() imysql.MySQLDriver {
	return &MySQLDriver{&mysql.MySQLDriver{}}
}

type MySQLDriver struct {
	driver *mysql.MySQLDriver
}

func (m *MySQLDriver) Open(dsn string) (driver.Conn, error) {
	return m.driver.Open(dsn)
}
