package manageUserHandler

import (
	"log"

	emailAlert "view_lab/src/controllers/alert/email"
	mophAlert "view_lab/src/controllers/alert/moph"
	handlerMophAlert "view_lab/src/controllers/alert/moph/handler"
	mophAlertUtils "view_lab/src/controllers/alert/moph/utils"
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
)

// NotifyUserStatusChange ส่งแจ้งเตือนไปยัง user เมื่อ status เปลี่ยน
func NotifyUserStatusChange(user *manageUserUtils.UserNotifyInfo) {
	// ส่ง Email
	if user.Email != "" {
		go func() {
			subject := "แจ้งเตือน: สถานะบัญชีของท่านเปลี่ยนแปลง"
			var body string
			if user.Status {
				body = "เรียนคุณ " + user.FullName + "\n\nบัญชีของท่านได้รับการอนุมัติแล้ว จาก ระบบ View Lab สามารถเข้าใช้งานระบบได้"
			} else {
				body = "เรียนคุณ " + user.FullName + "\n\nบัญชีของท่านถูกระงับการใช้งาน จาก ระบบ View Lab กรุณาติดต่อผู้ดูแลระบบ"
			}
			if err := emailAlert.SendEmailSMTP(user.Email, subject, body); err != nil {
				log.Printf("❌ ส่ง Email แจ้งเตือน status ไปยัง %s ล้มเหลว: %v", user.Email, err)
			} else {
				log.Printf("✅ ส่ง Email แจ้งเตือน status ไปยัง %s สำเร็จ", user.Email)
			}
		}()
	}

	// ส่ง MOPH Alert (LINE)
	if user.CID != "" {
		go func() {
			flexMsg := handlerMophAlert.CreateStatusChangeFlexMessage(user.FullName, user.Status)
			payload := mophAlertUtils.FlexAlertPayload{
				CID:      []string{user.CID},
				Messages: []mophAlertUtils.FlexMessage{flexMsg},
			}
			result := mophAlert.SendAlert(payload)
			if result.MessageCode == 200 {
				log.Printf("✅ ส่ง MOPH Alert แจ้งเตือน status ไปยัง %s สำเร็จ", user.FullName)
			} else {
				log.Printf("❌ ส่ง MOPH Alert แจ้งเตือน status ไปยัง %s ล้มเหลว: %s", user.FullName, result.Message)
			}
		}()
	}
}
