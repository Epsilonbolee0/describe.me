package postgres

import (
	"gorm.io/plugin/dbresolver"
)

// Option -.
type Option func(resolver *dbresolver.DBResolver)

// MaxIdleConnections -.
func MaxIdleConnections(conns int) Option {
	return func(resolver *dbresolver.DBResolver) {
		resolver.SetMaxIdleConns(conns)
	}
}

// MaxOpenConnections -.
func MaxOpenConnections(conns int) Option {
	return func(resolver *dbresolver.DBResolver) {
		resolver.SetMaxOpenConns(conns)
	}
}
