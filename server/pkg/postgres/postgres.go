// Package postgres implements postgres connection.
package postgres

import (
	"describe.me/config"
	"github.com/harranali/authority"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"time"
)

const (
	_defaultPrepareStmt            = true
	_defaultAllowGlobalUpdate      = true
	_defaultFullSaveAssociation    = true
	_defaultSkipDefaultTransaction = true
	_defaultCreateBatchSize        = 10
)

const (
	_defaultConnMaxIdleTime = time.Hour
	_defaultConnMaxLifetime = 24 * time.Hour

	_defaultMaxIdleConnections = 100
	_defaultMaxOpenConnections = 200
)

// Postgres -.
type Postgres struct {
	Conn *gorm.DB
	Auth *authority.Authority
}

// New -.
func New(dsn config.DSN, opts ...Option) (*Postgres, error) {
	conf := &gorm.Config{
		PrepareStmt:            _defaultPrepareStmt,
		AllowGlobalUpdate:      _defaultAllowGlobalUpdate,
		FullSaveAssociations:   _defaultFullSaveAssociation,
		SkipDefaultTransaction: _defaultSkipDefaultTransaction,
		CreateBatchSize:        _defaultCreateBatchSize,
	}

	resolver := dbresolver.
		Register(dbresolver.Config{Sources: []gorm.Dialector{postgres.Open(dsn.Unauthorized)}}, "unauthorized").
		Register(dbresolver.Config{Sources: []gorm.Dialector{postgres.Open(dsn.Student)}}, "student").
		Register(dbresolver.Config{Sources: []gorm.Dialector{postgres.Open(dsn.Teacher)}}, "teacher").
		SetConnMaxIdleTime(_defaultConnMaxIdleTime).
		SetConnMaxLifetime(_defaultConnMaxLifetime).
		SetMaxIdleConns(_defaultMaxIdleConnections).
		SetMaxOpenConns(_defaultMaxOpenConnections)

	for _, opt := range opts {
		opt(resolver)
	}

	connection, err := gorm.Open(postgres.Open(dsn.Unauthorized), conf)
	if err != nil {
		return nil, err
	}

	auth, err := setupAuthority(connection)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		connection,
		auth,
	}, nil
}

// setupAuthority -.
func setupAuthority(connection *gorm.DB) (*authority.Authority, error) {
	var err error

	auth := authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           connection,
	})

	if err = auth.CreateRole("student"); err != nil {
		return nil, err
	}
	if err = auth.CreateRole("teacher"); err != nil {
		return nil, err
	}

	if err = auth.CreatePermission("read"); err != nil {
		return nil, err
	}
	if err = auth.CreatePermission("write"); err != nil {
		return nil, err
	}

	if err = auth.AssignPermissions("student", []string{"read"}); err != nil {
		return nil, err
	}
	if err = auth.AssignPermissions("teacher", []string{"read", "write"}); err != nil {
		return nil, err
	}

	return auth, nil
}

// Close -.
func (p *Postgres) Close() {
	sqlConnection, err := p.Conn.DB()
	if err != nil {
		sqlConnection.Close()
	}
}
