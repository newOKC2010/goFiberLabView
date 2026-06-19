package mophAlert

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	handlerMophAlert "view_lab/src/controllers/alert/moph/handler"
	mophAlertUtils "view_lab/src/controllers/alert/moph/utils"
	loadEnv "view_lab/src/loadenv"
)

// returnError - helper สำหรับ return error response
func returnError(msg string, code int) mophAlertUtils.AlertResponse {
	return mophAlertUtils.AlertResponse{Message: msg, MessageCode: code}
}

// SendAlert - ส่งข้อความแจ้งเตือนผ่าน API
func SendAlert(payload mophAlertUtils.FlexAlertPayload) mophAlertUtils.AlertResponse {
	env := loadEnv.LoadMOPHAlert()
	apiURL := env.URL
	if apiURL == "" {
		apiURL = "https://api.example.com/alert"
	}

	// เตรียม JSON และสร้าง request
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("❌ JSON error: %v", err)
		return returnError("ส่งแจ้งเตือนไม่สำเร็จ", 500)
	}

	method := env.Method
	if method == "" {
		method = "POST"
	}

	req, err := http.NewRequest(method, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("❌ Request error: %v", err)
		return returnError("ส่งแจ้งเตือนไม่สำเร็จ", 500)
	}

	// ตั้งค่า headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("client-key", env.ClientID)
	req.Header.Set("secret-key", env.SecretKey)

	// ส่ง request - ปิด SSL verification สำหรับ MOPH API
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("❌ HTTP error: %v", err)
		return returnError("ส่งแจ้งเตือนไม่สำเร็จ", 500)
	}
	defer resp.Body.Close()

	// อ่านและ decode response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("❌ Read error: %v", err)
		return returnError("ส่งแจ้งเตือนไม่สำเร็จ", 500)
	}

	var result mophAlertUtils.AlertResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Printf("❌ JSON decode error: %v", err)
		return returnError("ส่งแจ้งเตือนไม่สำเร็จ", 500)
	}

	// แสดงผลลัพธ์
	log.Printf("📤 ส่งแจ้งเตือนไปยัง CID: %s", strings.Join(payload.CID, ", "))
	log.Printf("📊 สถานะ: %d - %s", result.MessageCode, result.Message)

	return result
}

func SendMophOTP(cid, otpCode, fullName string) mophAlertUtils.AlertResponse {
	flexMsg := handlerMophAlert.CreateOTPFlexMessage(otpCode, fullName)
	payload := mophAlertUtils.FlexAlertPayload{
		CID:      []string{cid},
		Messages: []mophAlertUtils.FlexMessage{flexMsg},
	}
	return SendAlert(payload)
}
