package serviceLogin

import (
	"database/sql"
	"time"

	loginHandler "view_lab/src/controllers/auth/login/otp/handler"
	conn "view_lab/src/database/connection"
	loadEnv "view_lab/src/loadenv"
)

// UpdateOTPInDB - อัพเดท OTP ลงฐานข้อมูล
func UpdateOTPInDB(_ *sql.DB, cid, otpCode string) error {
	expiresIn := loadEnv.LoadOTPExpiresIn()
	seconds := loginHandler.GetExpiresInSeconds(expiresIn, 300) // default 5 min
	now := time.Now()
	expiresAt := now.Add(time.Duration(seconds) * time.Second)

	_, err := conn.DB.Exec(
		`UPDATE user_view_lab
		 SET otp_code = $1, otp_expires_at = $2, updated_at = $3
		 WHERE cid = $4 AND status = true`,
		otpCode, expiresAt, now, cid,
	)

	return err
}
