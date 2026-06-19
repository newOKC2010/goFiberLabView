package handlerMophAlert

import (
	mophAlertUtils "view_lab/src/controllers/alert/moph/utils"
)

// CreateRegisterNotificationFlexMessage สร้าง Flex Message สำหรับแจ้งเตือน admin มีผู้ลงทะเบียนใหม่
func CreateRegisterNotificationFlexMessage(fullName, email, facilityName, facilityCode string) mophAlertUtils.FlexMessage {
	bubble := mophAlertUtils.FlexBubble{
		Type: "bubble",
		Header: &mophAlertUtils.FlexBox{
			Type:   "box",
			Layout: "vertical",
			Contents: []interface{}{
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   "🔔 แจ้งเตือนผู้ใช้ใหม่",
					Weight: "bold",
					Color:  "#1DB446",
					Size:   "lg",
				},
			},
		},
		Body: &mophAlertUtils.FlexBox{
			Type:   "box",
			Layout: "vertical",
			Contents: []interface{}{
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   "มีผู้ใช้ใหม่ลงทะเบียนรอการอนุมัติ",
					Wrap:   true,
					Size:   "sm",
					Color:  "#666666",
					Margin: "md",
				},
				mophAlertUtils.FlexBox{
					Type:    "box",
					Layout:  "vertical",
					Margin:  "lg",
					Spacing: "sm",
					Contents: []interface{}{
						createInfoRow("ชื่อ-นามสกุล", fullName),
						createInfoRow("อีเมล", email),
						createInfoRow("สถานพยาบาล", facilityName),
						createInfoRow("รหัสสถานพยาบาล", facilityCode),
					},
				},
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   "กรุณาตรวจสอบและอนุมัติในระบบ",
					Wrap:   true,
					Size:   "xs",
					Color:  "#aaaaaa",
					Margin: "xl",
				},
			},
		},
	}

	return mophAlertUtils.FlexMessage{
		Type:     "flex",
		AltText:  "🔔 มีผู้ใช้ใหม่ลงทะเบียนรอการอนุมัติ",
		Contents: bubble,
	}
}

func createInfoRow(label, value string) mophAlertUtils.FlexBox {
	return mophAlertUtils.FlexBox{
		Type:    "box",
		Layout:  "baseline",
		Spacing: "sm",
		Contents: []interface{}{
			mophAlertUtils.FlexText{
				Type:  "text",
				Text:  label,
				Color: "#aaaaaa",
				Size:  "sm",
				Flex:  2,
			},
			mophAlertUtils.FlexText{
				Type:  "text",
				Text:  value,
				Wrap:  true,
				Color: "#666666",
				Size:  "sm",
				Flex:  5,
			},
		},
	}
}
