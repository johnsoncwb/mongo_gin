package initializer

import "github.com/joho/godotenv"

func LoadEnvFile() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
