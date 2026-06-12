package provider

// HealthIDToken struct สำหรับ Health ID token response
type HealthIDToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccountID   string `json:"account_id"`
}

// ProviderIDToken struct สำหรับ Provider ID token response
type ProviderIDToken struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountID      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}

// Organization struct สำหรับข้อมูลองค์กร
type Organization struct {
	BusinessID   string  `json:"business_id,omitempty"`
	HCode        string  `json:"hcode,omitempty"`
	HNameTH      string  `json:"hname_th,omitempty"`
	HNameENG     string  `json:"hname_eng,omitempty"`
	Position     string  `json:"position,omitempty"`
	PositionType string  `json:"position_type,omitempty"`
	IsHRAdmin    bool    `json:"is_hr_admin,omitempty"`
	IsDirector   bool    `json:"is_director,omitempty"`
	Address      Address `json:"address,omitempty"`
}

// Address struct สำหรับที่อยู่
type Address struct {
	Address     string `json:"address,omitempty"`
	Province    string `json:"province,omitempty"`
	District    string `json:"district,omitempty"`
	SubDistrict string `json:"sub_district,omitempty"`
	ZipCode     string `json:"zip_code,omitempty"`
}

// ProviderProfile struct สำหรับข้อมูล profile
type ProviderProfile struct {
	AccountID    string         `json:"account_id"`
	ProviderID   string         `json:"provider_id"`
	HashCID      string         `json:"hash_cid,omitempty"`
	TitleTH      string         `json:"title_th,omitempty"`
	TitleEN      string         `json:"title_en,omitempty"`
	NameTH       string         `json:"name_th"`
	NameENG      string         `json:"name_eng,omitempty"`
	FirstnameTH  string         `json:"firstname_th,omitempty"`
	LastnameTH   string         `json:"lastname_th,omitempty"`
	FirstnameEN  string         `json:"firstname_en,omitempty"`
	LastnameEN   string         `json:"lastname_en,omitempty"`
	Organization []Organization `json:"organization,omitempty"`
}
