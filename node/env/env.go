package env

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func init() {
	envPath, err := filepath.Abs(".env")
	if err != nil {
		return
	}

	_, err = os.Stat(envPath)
	if err != nil {
		return
	}

	_ = godotenv.Load(envPath)
}
