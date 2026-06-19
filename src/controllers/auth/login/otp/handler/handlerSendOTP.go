package handlerLogin

import (
	"log"

	emailAlert "view_lab/src/controllers/alert/email"
	mophAlert "view_lab/src/controllers/alert/moph"
	modelAuth "view_lab/src/database/models/auth"
)

func SendOTPToUser(user *modelAuth.UserViewLab, otpCode string) {

	if user.CID != "" {
		go func() {
			result := mophAlert.SendMophOTP(user.CID, otpCode, user.FullName)
			if result.MessageCode == 200 {
				log.Printf("✅ ส่ง OTP ผ่าน MOPH สำเร็จ: %s", user.CID)
			} else {
				log.Printf("❌ ส่ง OTP ผ่าน MOPH ล้มเหลว: %s", result.Message)
			}
		}()
	}
	if user.Email != "" {
		go func() {
			err := emailAlert.SendOTPEmail(user.Email, otpCode, user.FullName)
			if err != nil {
				log.Printf("❌ ส่ง OTP ผ่าน Email ล้มเหลว: %v", err)
			} else {
				log.Printf("✅ ส่ง OTP ผ่าน Email สำเร็จ: %s", user.Email)
			}
		}()
	}
}
