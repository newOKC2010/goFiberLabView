package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	conn "view_lab/src/database/connection"
	modelAuth "view_lab/src/database/models/auth"
	loadEnv "view_lab/src/loadenv"
)

// GetUserByHashCID ดึงข้อมูล user ด้วย hash_cid สำหรับ provider login
func getUserByHashCID(hashCID string) (*modelAuth.UserViewLab, error) {
	var user modelAuth.UserViewLab
	var role string
	err := conn.DB.QueryRow(
		`SELECT id, hash_cid, full_name, email, role, status
		 FROM user_view_lab 
		 WHERE hash_cid = $1`,
		hashCID,
	).Scan(&user.ID, &user.HashCID, &user.FullName, &user.Email, &role, &user.Status)

	if err != nil {
		return nil, err
	}
	user.Role = modelAuth.UserRole(role)
	return &user, nil
}

// getAuthorizationUrl สร้าง OAuth authorization URL
func getAuthorizationUrl(redirectURI string) (string, error) {
	oauthConfig := loadEnv.LoadOauth() // ใช้ config ที่มีอยู่

	if oauthConfig.HealthID.ClientID == "" {
		return "", fmt.Errorf("client ID is required")
	}

	params := url.Values{}
	params.Add("client_id", oauthConfig.HealthID.ClientID)
	params.Add("redirect_uri", redirectURI)
	params.Add("response_type", "code")

	return fmt.Sprintf("%s/oauth/redirect?%s", oauthConfig.HealthID.BaseURL, params.Encode()), nil
}

// getAccessTokenHealthId ขอ Access Token จาก Health ID
func getAccessTokenHealthId(code, redirectURI string) (*HealthIDToken, error) {
	oauthConfig := loadEnv.LoadOauth()

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("client_id", oauthConfig.HealthID.ClientID)
	data.Set("client_secret", oauthConfig.HealthID.ClientSecret)

	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/token", oauthConfig.HealthID.BaseURL),
		"application/x-www-form-urlencoded",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request token: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Data HealthIDToken `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &result.Data, nil
}

// getProviderIDToken ขอ Provider ID Token
func getProviderIDToken(healthIdToken string) (*ProviderIDToken, error) {
	oauthConfig := loadEnv.LoadOauth()

	payload := map[string]interface{}{
		"client_id":  oauthConfig.ProviderID.ClientID,
		"secret_key": oauthConfig.ProviderID.ClientSecret,
		"token_by":   "Health ID",
		"token":      healthIdToken,
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/services/token", oauthConfig.ProviderID.BaseURL),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request provider token: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Status  int             `json:"status"`
		Data    ProviderIDToken `json:"data"`
		Message string          `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Status != 200 {
		return nil, fmt.Errorf("failed to get provider token: %s", result.Message)
	}

	return &result.Data, nil
}

// getProviderProfile ขอข้อมูล Profile
func getProviderProfile(providerToken string) (*ProviderProfile, error) {
	oauthConfig := loadEnv.LoadOauth()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/services/profile", oauthConfig.ProviderID.BaseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+providerToken)
	req.Header.Set("client-id", oauthConfig.ProviderID.ClientID)
	req.Header.Set("secret-key", oauthConfig.ProviderID.ClientSecret)

	q := req.URL.Query()
	q.Add("moph_center_token", "1")
	q.Add("moph_idp_permission", "1")
	q.Add("position_type", "1")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request profile: %v", err)
	}
	defer resp.Body.Close()

	var result struct {
		Status  int             `json:"status"`
		Data    ProviderProfile `json:"data"`
		Message string          `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Status != 200 {
		return nil, fmt.Errorf("failed to get profile: %s", result.Message)
	}

	return &result.Data, nil
}
