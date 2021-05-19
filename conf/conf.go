package conf

import (
	"log"
	"os"
)

const (
	CEMETERY_PARK_PORT    = "CEMETERY_PARK_PORT"
	CEMETERY_PARK_DB_PATH = "CEMETERY_PARK_DB_PATH"
)

//Configuration app configuration
type Configuration struct {
	Port         string
	ClientFolder string
	Sqlite       SqliteConfig
}

//SqliteConfig sqlite configuration
type SqliteConfig struct {
	DbPath string
}

var configuration *Configuration

//Get get configuration
func Get() *Configuration {
	if configuration == nil {
		config, err := load()
		handleError(err)
		configuration = config
	}
	return configuration
}

//Load configuration
func load() (*Configuration, error) {
	var err error
	port := os.Getenv(CEMETERY_PARK_PORT)
	// handleError(err)

	dbPath := os.Getenv(CEMETERY_PARK_DB_PATH)

	return &Configuration{
		Port: port,
		Sqlite: SqliteConfig{
			DbPath: dbPath,
		},
	}, err
}

//Process error
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
