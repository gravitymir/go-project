package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel            string `envconfig:"LOG_LEVEL"`
	PgURL               string `envconfig:"PG_URL"`
	PgMigrationsPath    string `envconfig:"PG_MIGRATIONS_PATH"`
	MysqlAddr           string `envconfig:"MYSQL_ADDR"`
	MysqlUser           string `envconfig:"MYSQL_USER"`
	MysqlPassword       string `envconfig:"MYSQL_PASSWORD"`
	MysqlDB             string `envconfig:"MYSQL_DB"`
	MysqlMigrationsPath string `envconfig:"MYSQL_MIGRATIONS_PATH"`
	HTTPAddr            string `envconfig:"HTTP_ADDR"`
	GCBucket            string `envconfig:"GC_BUCKET"`
	FilePath            string `envconfig:"FILE_PATH"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}

		cb, err := json.MarshalIndent(config, "", "  ")

		cb = bytes.ReplaceAll(cb, []byte(`\u003c`), []byte("<"))
		cb = bytes.ReplaceAll(cb, []byte(`\u003e`), []byte(">"))
		cb = bytes.ReplaceAll(cb, []byte(`\u0026`), []byte("&"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(cb))
	})
	return &config
}
