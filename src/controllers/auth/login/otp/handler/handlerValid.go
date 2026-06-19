package handlerLogin

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(strings.TrimSpace(email))
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func IsValidOTP(otp string) bool {
	otp = strings.TrimSpace(otp)
	if len(otp) != 6 {
		return false
	}
	for _, c := range otp {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func ValidateRequestOTP(body map[string]interface{}) (string, string) {
	if len(body) != 1 {
		return "", "ส่งข้อมูลได้เฉพาะ email เท่านั้น"
	}

	email, ok := body["email"].(string)
	if !ok {
		return "", "กรุณาระบุอีเมล"
	}

	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return "", "กรุณาระบุอีเมล"
	}

	if !emailRegex.MatchString(email) {
		return "", "รูปแบบอีเมลไม่ถูกต้อง"
	}

	return email, ""
}

func ValidateVerifyOTP(body map[string]interface{}) (string, string, string) {
	if len(body) != 2 {
		return "", "", "ข้อมูลได้เฉพาะ email และ otp เท่านั้น"
	}

	email, ok := body["email"].(string)
	if !ok {
		return "", "", "กรุณาระบุอีเมล"
	}

	otp, ok := body["otp"].(string)
	if !ok {
		return "", "", "กรุณาระบุรหัส OTP"
	}

	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return "", "", "กรุณาระบุอีเมล"
	}

	if !emailRegex.MatchString(email) {
		return "", "", "รูปแบบอีเมลไม่ถูกต้อง"
	}

	otp = strings.TrimSpace(otp)
	if otp == "" {
		return "", "", "กรุณาระบุรหัส OTP"
	}

	if !IsValidOTP(otp) {
		return "", "", "รหัส OTP ไม่ถูกต้อง (ต้องเป็นตัวเลข 6 หลัก)"
	}

	return email, otp, ""
}
