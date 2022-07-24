package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
)

func LoadEnvFile() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	strDir, err := getFirstFileExists(
		dir,
		path.Join(dir, ".."),
		path.Join(dir, "..", ".."),
		path.Join(dir, "..", "..", ".."),
	)
	if err != nil {
		return
	}
	err = godotenv.Load(strDir)
	if err != nil {
		panic("Error loading .env file")
	}
}

func getFirstFileExists(pathFile ...string) (string, error) {
	var err error
	for _, item := range pathFile {
		item = path.Join(item, ".env")
		if _, err = os.Stat(item); err == nil {
			return item, nil
		}
	}
	return "", err
}

func GetEnvOrDefault(key string, value string) string {
	result := os.Getenv(key)
	if result != "" {
		return result
	}
	return value
}

func IsDevelopment() bool {
	env := GetEnvOrDefault("APP_ENV", "Debug")
	return env == "Debug"
}
