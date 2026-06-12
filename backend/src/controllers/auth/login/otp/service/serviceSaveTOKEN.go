package serviceLogin

import (
	"database/sql"
	"time"

	loginHandler "view_lab/src/controllers/auth/login/otp/handler"
	conn "view_lab/src/database/connection"
	loadEnv "view_lab/src/loadenv"
)

func SaveTokenToDB(_ *sql.DB, userViewLabID int64, token string) error {
	jwtConfig := loadEnv.LoadJWT()
	expiresInSeconds := loginHandler.GetExpiresInSeconds(jwtConfig.ExpireIn, 3600)
	expiresAt := time.Now().Add(time.Second * time.Duration(expiresInSeconds))

	_, err := conn.DB.Exec(
		`INSERT INTO tokens_view_lab (user_view_lab_id, token, expires_at, login_last)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (user_view_lab_id)
		 DO UPDATE SET token = $2, expires_at = $3, login_last = $4`,
		userViewLabID, token, expiresAt, time.Now(),
	)

	return err
}
