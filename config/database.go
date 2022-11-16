package config

import "fmt"

type database struct {
	host     string
	port     int
	username string
	password string
	dbName   string
}

func (d *database) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		d.host, d.port, d.username, d.password, d.dbName)
}
