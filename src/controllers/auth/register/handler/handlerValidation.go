package handlerRegister

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidateCID(cid string) error {
	if len(cid) != 13 {
		return fmt.Errorf("เลขบัตรประชาชนต้องมี 13 หลัก")
	}
	if !regexp.MustCompile(`^\d{13}$`).MatchString(cid) {
		return fmt.Errorf("เลขบัตรประชาชนต้องเป็นตัวเลขเท่านั้น")
	}
	return nil
}

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("กรุณากรอกอีเมล")
	}
	pattern := `^[a-zA-Z0-9._@#]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched := regexp.MustCompile(pattern).MatchString(email)
	if !matched {
		return fmt.Errorf("รูปแบบอีเมลไม่ถูกต้อง (ใช้ได้เฉพาะ a-z A-Z 0-9 . _ @ #)")
	}
	return nil
}

func ValidateFullName(fullName string) error {
	if fullName == "" {
		return fmt.Errorf("กรุณากรอกชื่อ-นามสกุล")
	}
	fullName = strings.TrimSpace(fullName)
	matched, err := regexp.MatchString(`^[\p{Thai}\s]+$`, fullName)
	if err != nil || !matched {
		return fmt.Errorf("ชื่อ-นามสกุลต้องเป็นภาษาไทยเท่านั้น")
	}
	return nil
}
