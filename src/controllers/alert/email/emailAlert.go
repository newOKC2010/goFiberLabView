package emailAlertSend

import (
	"net/smtp"

	handlerEmailAlert "view_lab/src/controllers/alert/email/handler"
	loadEnv "view_lab/src/loadenv"
)

func SendEmailSMTP(to, subject, body string) error {
	emailConfig := loadEnv.LoadEmail()

	from := emailConfig.Email
	password := emailConfig.Password

	smtpHost := "mail.privateemail.com"
	smtpPort := "587"

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	addr := smtpHost + ":" + smtpPort

	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

func SendOTPEmail(to, otpCode, fullName string) error {
	subject := "� รหัส OTP สำหรับเข้าใช้งานระบบดูผล lab"
	body := handlerEmailAlert.CreateOTPEmailHTML(otpCode, fullName)
	return SendEmailSMTP(to, subject, body)
}
