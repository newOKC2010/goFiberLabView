package loadEnv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func LoadEmail() EmailConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return EmailConfig{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("EMAIL_PASSWORD"),
	}
}

func LoadMOPHAlert() MOPHAlertConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return MOPHAlertConfig{
		URL:       os.Getenv("MOPH_ALERT_URL"),
		ClientID:  os.Getenv("MOPH_CLIENT_ID"),
		SecretKey: os.Getenv("MOPH_SECRET_KEY"),
	}
}

func LoadDBconnec() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv("DB_URL")
}

func LoadPort() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv("PORT")
}

func LoadCreateModel() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return os.Getenv("CREATE_MODEL")
}
