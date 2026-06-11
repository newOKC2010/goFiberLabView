package handlerEmailAlert

import (
	"fmt"
	"time"
)

func CreateOTPEmailHTML(otpCode, fullName string) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	currentTime := time.Now().In(loc).Format("02/01/2006 15:04:05")
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: 40px auto; background: white; border-radius: 10px; overflow: hidden; box-shadow: 0 2px 10px rgba(0,0,0,0.1);">
        
        <!-- Header -->
        <div style="background: linear-gradient(135deg, #43a047 0%%, #2e7d32 100%%); padding: 30px; text-align: center;">
            <h1 style="margin: 0; color: white; font-size: 24px;">🌿 ระบบลงข้อมูลทรัพยากร</h1>
            <p style="margin: 10px 0 0 0; color: rgba(255,255,255,0.9); font-size: 14px;">Ecobase System</p>
        </div>
        
        <!-- Body -->
        <div style="padding: 40px 30px;">
            <p style="margin: 0 0 20px 0; color: #333; font-size: 16px;">สวัสดีคุณ <strong>%s</strong></p>
            
            <p style="margin: 0 0 30px 0; color: #666; font-size: 14px; line-height: 1.6;">
                คุณได้ทำการร้องขอรหัส OTP สำหรับเข้าใช้งานระบบลงข้อมูลทรัพยากร<br>
                กรุณาใช้รหัสด้านล่างเพื่อยืนยันตัวตนและเข้าสู่ระบบ
            </p>
            
            <!-- OTP Box -->
            <div style="background: #f1fff4; border: 2px dashed #6ee778; border-radius: 8px; padding: 30px; text-align: center; margin: 30px 0;">
                <p style="margin: 0 0 10px 0; color: #888; font-size: 12px; text-transform: uppercase; letter-spacing: 1px;">รหัส OTP ของคุณ</p>
                <div style="font-size: 36px; font-weight: bold; color: #2e7d32; letter-spacing: 8px; font-family: 'Courier New', monospace;">
                    %s
                </div>
            </div>
            
            <!-- Warning -->
            <div style="background: #fff3cd; border-left: 4px solid #ffc107; padding: 15px; margin: 20px 0; border-radius: 4px;">
                <p style="margin: 0; color: #856404; font-size: 13px;">
                    ⏱️ <strong>รหัสนี้มีอายุ 2 นาที</strong><br>
                    🕐 ส่งเมื่อ: <strong>%s</strong><br>
                    หากหมดอายุ กรุณาขอรหัสใหม่อีกครั้ง
                </p>
            </div>
            
            <div style="background: #e8f5e9; border-left: 4px solid #43a047; padding: 15px; margin: 20px 0; border-radius: 4px;">
                <p style="margin: 0; color: #1b5e20; font-size: 13px;">
                    ⚠️ <strong>คำเตือน:</strong> กรุณาอย่าแชร์รหัสนี้กับผู้อื่น<br>
                    เจ้าหน้าที่จะไม่ขอรหัส OTP จากคุณทางใดๆ ทั้งสิ้น
                </p>
            </div>
        </div>
        
        <!-- Footer -->
        <div style="background: #f8f9fa; padding: 20px 30px; text-align: center; border-top: 1px solid #e9ecef;">
            <p style="margin: 0; color: #6c757d; font-size: 12px;">
                © 2025 ระบบลงข้อมูลทรัพยากร<br>
                โรงพยาบาล | ระบบอัตโนมัติ - กรุณาอย่าตอบกลับอีเมลนี้
            </p>
        </div>
        
    </div>
</body>
</html>
`, fullName, otpCode, currentTime)
}
