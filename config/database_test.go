package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConfig(t *testing.T) {
	db := &database{
		host:     "localhost",
		port:     5432,
		username: "postgres",
		password: "Viyim*80",
		dbName:   "restaurant_rivaldo",
	}

	assert.NotNil(t, db)

	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.username, db.password, db.dbName)

	assert.Equal(t, config, db.String())
}
