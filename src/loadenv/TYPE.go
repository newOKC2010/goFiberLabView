package loadEnv

type EmailConfig struct {
	Email    string
	Password string
}

type MOPHAlertConfig struct {
	URL       string
	Method    string
	ClientID  string
	SecretKey string
}

type CORS struct {
	Origins     []string
	Credentials bool
	Methods     []string
	Headers     []string
}

type JWT struct {
	Secret   string
	ExpireIn string
}

// Oauth Config
type LoadOauthConfig struct {
	HealthID    HealthIDConfig
	ProviderID  ProviderIDConfig
	Redirect    RedirectConfig
	FrontendURL string
}

type HealthIDConfig struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
}

type ProviderIDConfig struct {
	BaseURL      string
	ClientID     string
	ClientSecret string
}

type RedirectConfig struct {
	CallbackURL string
}
