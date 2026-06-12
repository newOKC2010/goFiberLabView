package loadEnv

type EmailConfig struct {
	Email    string
	Password string
}

type MOPHAlertConfig struct {
	URL       string
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
