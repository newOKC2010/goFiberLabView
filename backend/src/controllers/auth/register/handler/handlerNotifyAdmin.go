package handlerRegister

import (
	"fmt"
	"log"

	emailAlert "view_lab/src/controllers/alert/email"
	mophAlert "view_lab/src/controllers/alert/moph"
	handlerMophAlert "view_lab/src/controllers/alert/moph/handler"
	mophAlertUtils "view_lab/src/controllers/alert/moph/utils"
	modelAuth "view_lab/src/database/models/auth"
)

// NotifyAdminsNewRegistration ส่งแจ้งเตือนไปยัง admin เมื่อมีการลงทะเบียนใหม่
func NotifyAdminsNewRegistration(admins []*modelAuth.UserViewLab, newUser *modelAuth.UserViewLab) {
	message := fmt.Sprintf(
		"มีผู้ใช้ใหม่ลงทะเบียนรอการอนุมัติ\n\nชื่อ: %s\nอีเมล: %s\nสถานพยาบาล: %s (%s)",
		newUser.FullName,
		newUser.Email,
		getStringValue(newUser.FacilityName),
		getStringValue(newUser.FacilityCode),
	)

	for _, admin := range admins {
		// ส่ง Email
		if admin.Email != "" {
			go func(adminEmail, adminName string) {
				subject := "🔔 แจ้งเตือน: มีผู้ใช้ใหม่ลงทะเบียน"
				err := emailAlert.SendEmailSMTP(adminEmail, subject, message)
				if err != nil {
					log.Printf("❌ ส่ง Email แจ้งเตือนไปยัง %s ล้มเหลว: %v", adminEmail, err)
				} else {
					log.Printf("✅ ส่ง Email แจ้งเตือนไปยัง %s สำเร็จ", adminEmail)
				}
			}(admin.Email, admin.FullName)
		}

		// ส่ง MOPH Alert (LINE)
		if admin.CID != "" {
			go func(adminCID, adminName string) {
				flexMsg := handlerMophAlert.CreateRegisterNotificationFlexMessage(
					newUser.FullName,
					newUser.Email,
					getStringValue(newUser.FacilityName),
					getStringValue(newUser.FacilityCode),
				)
				payload := mophAlertUtils.FlexAlertPayload{
					CID:      []string{adminCID},
					Messages: []mophAlertUtils.FlexMessage{flexMsg},
				}
				result := mophAlert.SendAlert(payload)
				if result.MessageCode == 200 {
					log.Printf("✅ ส่ง MOPH Alert แจ้งเตือนไปยัง %s สำเร็จ", adminName)
				} else {
					log.Printf("❌ ส่ง MOPH Alert แจ้งเตือนไปยัง %s ล้มเหลว: %s", adminName, result.Message)
				}
			}(admin.CID, admin.FullName)
		}
	}
}

func getStringValue(s *string) string {
	if s == nil {
		return "-"
	}
	return *s
}
