package loadEnv

import (
	"log"
	"os"
	"strconv"
	"strings"

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

func LoadCORS() CORS {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	credentials, _ := strconv.ParseBool(os.Getenv("CORS_CREDENTIALS"))
	return CORS{
		Origins:     strings.Split(os.Getenv("CORS_ORIGINS"), ","),
		Credentials: credentials,
		Methods:     strings.Split(os.Getenv("CORS_METHODS"), ","),
		Headers:     strings.Split(os.Getenv("CORS_HEADERS"), ","),
	}
}

func LoadJWT() JWT {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	return JWT{
		Secret:   os.Getenv("JWT_SECRET"),
		ExpireIn: os.Getenv("JWT_EXPIRES_IN"),
	}
}

func LoadOauth() LoadOauthConfig {
	godotenv.Load()
	return LoadOauthConfig{
		HealthID: HealthIDConfig{
			BaseURL:      os.Getenv("HEALTH_ID_URL"),
			ClientID:     os.Getenv("HEALTH_ID_CLIENT_ID"),
			ClientSecret: os.Getenv("HEALTH_ID_CLIENT_SECRET"),
		},
		ProviderID: ProviderIDConfig{
			BaseURL:      os.Getenv("PROVIDER_ID_URL"),
			ClientID:     os.Getenv("PROVIDER_ID_CLIENT_ID"),
			ClientSecret: os.Getenv("PROVIDER_ID_CLIENT_SECRET"),
		},
		Redirect: RedirectConfig{
			CallbackURL: os.Getenv("CALL_BACK_URL"),
		},
		FrontendURL: os.Getenv("FRONTEND_URL"),
	}
}
