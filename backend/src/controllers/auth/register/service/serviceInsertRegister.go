package serviceRegister

import (
	"fmt"

	registerUtils "view_lab/src/controllers/auth/register/utils"
	conn "view_lab/src/database/connection"
	modelAuth "view_lab/src/database/models/auth"
	loadEnv "view_lab/src/loadenv"
)

func IsAdminCID(cid string) bool {
	for _, c := range loadEnv.LoadAdminCIDs() {
		if c == cid {
			return true
		}
	}
	return false
}

func CreateUser(cid, fullName, email, facilityType, facilityCode, facilityName string) (*modelAuth.UserViewLab, error) {
	hashCID := registerUtils.HashCID(cid)

	role := modelAuth.RoleUser
	status := false
	if IsAdminCID(cid) {
		role = modelAuth.RoleAdmin
		status = true
	}

	user := &modelAuth.UserViewLab{
		CID:          cid,
		HashCID:      hashCID,
		FullName:     fullName,
		Email:        email,
		Role:         role,
		Status:       status,
		FacilityType: &facilityType,
		FacilityCode: &facilityCode,
		FacilityName: &facilityName,
	}

	_, err := conn.DB.Exec(
		`INSERT INTO user_view_lab (cid, hash_cid, full_name, email, role, status, facility_type, facility_code, facility_name)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		user.CID, user.HashCID, user.FullName, user.Email, string(user.Role), user.Status,
		user.FacilityType, user.FacilityCode, user.FacilityName,
	)
	if err != nil {
		return nil, fmt.Errorf("บันทึกข้อมูลไม่สำเร็จ: %w", err)
	}

	return user, nil
}
