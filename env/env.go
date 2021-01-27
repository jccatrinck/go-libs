package env

import (
	"os"

	"github.com/joho/godotenv"
)

var env map[string]string

func Load(filename string) (err error) {
	return godotenv.Load(filename)
}

func Get(key string, def string) (v string) {
	v, exists := os.LookupEnv(key)

	if !exists {
		return def
	}

	return
}
