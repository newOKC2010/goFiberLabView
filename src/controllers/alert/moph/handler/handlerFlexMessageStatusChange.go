package handlerMophAlert

import (
	mophAlertUtils "view_lab/src/controllers/alert/moph/utils"
)

// CreateStatusChangeFlexMessage สร้าง Flex Message แจ้งเตือน user เมื่อ status เปลี่ยน
func CreateStatusChangeFlexMessage(fullName string, status bool) mophAlertUtils.FlexMessage {
	var headerColor, statusText, bodyText string
	if status {
		headerColor = "#1DB446"
		statusText = "✅ บัญชีได้รับการอนุมัติแล้ว"
		bodyText = "บัญชีของท่านได้รับการอนุมัติแล้ว จาก ระบบ View Lab สามารถเข้าใช้งานระบบได้"
	} else {
		headerColor = "#E53935"
		statusText = "⛔ บัญชีถูกระงับการใช้งาน"
		bodyText = "บัญชีของท่านถูกระงับการใช้งาน จาก ระบบ View Lab กรุณาติดต่อผู้ดูแลระบบ"
	}

	bubble := mophAlertUtils.FlexBubble{
		Type: "bubble",
		Header: &mophAlertUtils.FlexBox{
			Type:   "box",
			Layout: "vertical",
			Contents: []interface{}{
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   statusText,
					Weight: "bold",
					Color:  headerColor,
					Size:   "md",
				},
			},
		},
		Body: &mophAlertUtils.FlexBox{
			Type:   "box",
			Layout: "vertical",
			Contents: []interface{}{
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   "เรียน " + fullName,
					Weight: "bold",
					Size:   "sm",
					Wrap:   true,
				},
				mophAlertUtils.FlexText{
					Type:   "text",
					Text:   bodyText,
					Wrap:   true,
					Size:   "sm",
					Color:  "#666666",
					Margin: "md",
				},
			},
		},
	}

	altText := "แจ้งเตือน: บัญชีของท่านถูกระงับการใช้งาน"
	if status {
		altText = "แจ้งเตือน: บัญชีของท่านได้รับการอนุมัติแล้ว"
	}

	return mophAlertUtils.FlexMessage{
		Type:     "flex",
		AltText:  altText,
		Contents: bubble,
	}
}
