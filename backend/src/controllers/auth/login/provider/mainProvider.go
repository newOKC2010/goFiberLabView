package provider

import (
	"fmt"
	"net/url"
	"time"

	conn "view_lab/src/database/connection"
	modelAuth "view_lab/src/database/models/auth"
	loadEnv "view_lab/src/loadenv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	FrontendURL string
)

func init() {
	oauthConfig := loadEnv.LoadOauth()
	FrontendURL = oauthConfig.FrontendURL
}

func generateJWTToken(user modelAuth.UserViewLab) (string, error) {
	jwtConfig := loadEnv.LoadJWT()
	expireIn := 24 * time.Hour

	claims := jwt.MapClaims{
		"user_view_lab_id": user.ID,
		"email":            user.Email,
		"role":             string(user.Role),
		"exp":              time.Now().Add(expireIn).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

func saveTokenToDB(userID int64, token string) error {
	_, err := conn.DB.Exec(
		`INSERT INTO tokens_view_lab (user_view_lab_id, token, expires_at, login_last)
		 VALUES ($1, $2, $3, $4)
		 ON CONFLICT (user_view_lab_id)
		 DO UPDATE SET token = $2, expires_at = $3, login_last = $4`,
		userID, token, time.Now().Add(24*time.Hour), time.Now(),
	)
	return err
}

// ProviderLogin เริ่ม OAuth flow
func ProviderLogin(c *fiber.Ctx) error {
	// สร้าง callback URL สำหรับ OAuth
	callbackURL := fmt.Sprintf("http://%s/provider/callback", c.Get("Host"))

	// สร้าง authorization URL
	authUrl, err := getAuthorizationUrl(callbackURL)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "เกิดข้อผิดพลาดในการสร้าง authorization URL"})
	}

	return c.Redirect(authUrl)
}

// ProviderCallback handle OAuth callback
func ProviderCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("ไม่พบ code จากระบบ"))
		return c.Redirect(redirectUrl)
	}

	// สร้าง callback URL เดียวกับที่ส่งไป
	callbackURL := fmt.Sprintf("http://%s/provider/callback", c.Get("Host"))

	// ขอ Health ID Token
	healthIdToken, err := getAccessTokenHealthId(code, callbackURL)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("เกิดข้อผิดพลาดในการขอ Access Token"))
		return c.Redirect(redirectUrl)
	}

	// ขอ Provider ID Token
	providerIdToken, err := getProviderIDToken(healthIdToken.AccessToken)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("เกิดข้อผิดพลาดในการขอ Provider Token"))
		return c.Redirect(redirectUrl)
	}

	// ขอข้อมูล Profile
	profile, err := getProviderProfile(providerIdToken.AccessToken)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("เกิดข้อผิดพลาดในการขอข้อมูล Profile"))
		return c.Redirect(redirectUrl)
	}

	// ตรวจสอบ hash_cid
	if profile.HashCID == "" {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("ไม่พบข้อมูล กรุณาลงทะเบียนสมาชิกก่อนใช้งาน หรือไม่พบข้อมูลจากระบบ"))
		return c.Redirect(redirectUrl)
	}

	// ตรวจสอบ user ในฐานข้อมูล
	user, err := getUserByHashCID(profile.HashCID)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("ไม่พบข้อมูล กรุณาลงทะเบียนสมาชิกก่อนใช้งาน หรือไม่พบข้อมูลจากระบบ"))
		return c.Redirect(redirectUrl)
	}

	// ตรวจสอบสถานะ user
	if !user.Status {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("บัญชีผู้ใช้ถูกระงับอยู่ กรุณาติดต่อ admin"))
		return c.Redirect(redirectUrl)
	}

	// สร้าง JWT Token
	token, err := generateJWTToken(*user)
	if err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("เกิดข้อผิดพลาดในการสร้าง token"))
		return c.Redirect(redirectUrl)
	}

	// บันทึก token ลงฐานข้อมูล
	if err := saveTokenToDB(user.ID, token); err != nil {
		redirectUrl := fmt.Sprintf("%s?success=false&message=%s",
			FrontendURL, url.QueryEscape("เกิดข้อผิดพลาดในการบันทึก token"))
		return c.Redirect(redirectUrl)
	}

	// Redirect ไป frontend พร้อมข้อมูล
	successMessage := url.QueryEscape("เข้าสู่ระบบสำเร็จ ด้วย provider สำเร็จ")
	redirectUrl := fmt.Sprintf("%s?success=true&token=%s&message=%s",
		FrontendURL, token, successMessage)

	return c.Redirect(redirectUrl)
}
