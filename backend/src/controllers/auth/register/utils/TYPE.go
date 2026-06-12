package registerUtils

type RegisterRequest struct {
	CID          string `json:"cid" validate:"required"`
	FullName     string `json:"full_name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	FacilityType string `json:"facility_type" validate:"required"`
	FacilityCode string `json:"facility_code" validate:"required"`
	FacilityName string `json:"facility_name" validate:"required"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
