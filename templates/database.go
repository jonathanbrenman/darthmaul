package templates

var DatabaseProvider = `
package database

import "gorm.io/gorm"

type DBProvider interface {
	GetDbClient() *gorm.DB
}
`

var MysqlProvider = `
package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlProvider struct {
	DB *gorm.DB
}

var mysqlClient *mysqlProvider

// Singleton method
func NewMysqlProvider() DBProvider {
	if mysqlClient != nil {
		return mysqlClient
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"user",
		"pass",
		"host",
		"port",
		"dbname",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlClient = &mysqlProvider{db}
	return mysqlClient
}

func (m *mysqlProvider) GetDbClient() *gorm.DB {
	return m.DB
}
`

var PostgresProvider = `
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgProvider struct {
	DB *gorm.DB
}

var pgClient *pgProvider

// Singleton method
func NewPGProvider() DBProvider {
	if pgClient != nil {
		return pgClient
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"host",
		"user",
		"pass",
		"dbname",
		"port",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	pgClient = &pgProvider{db}
	return pgClient
}

func (p *pgProvider) GetDbClient() *gorm.DB {
	return p.DB
}
`

var SqliteProvider = `
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3

type sqliteProvider struct {
	DB *gorm.DB
}

var sqliteClient *sqliteProvider

// Singleton method
func NewsqliteProvider() DBProvider {
	if sqliteClient != nil {
		return sqliteClient
	}
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqliteClient = &sqliteProvider{db}
	return sqliteClient
}

func (p *sqliteProvider) GetDbClient() *gorm.DB {
	return p.DB
}
`