package serviceRegister

import (
	conn "view_lab/src/database/connection"
	modelAuth "view_lab/src/database/models/auth"
)

// GetAdmins ดึงข้อมูล admin และ super_admin ที่ active
func GetAdmins() ([]*modelAuth.UserViewLab, error) {
	rows, err := conn.DB.Query(
		`SELECT id, cid, hash_cid, full_name, email, role, status
		 FROM user_view_lab 
		 WHERE (role = $1 OR role = $2) AND status = true`,
		string(modelAuth.RoleAdmin), string(modelAuth.RoleSuperAdmin),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var admins []*modelAuth.UserViewLab
	for rows.Next() {
		var admin modelAuth.UserViewLab
		var role string

		err := rows.Scan(
			&admin.ID, &admin.CID, &admin.HashCID,
			&admin.FullName, &admin.Email, &role, &admin.Status,
		)
		if err != nil {
			return nil, err
		}
		admin.Role = modelAuth.UserRole(role)
		admins = append(admins, &admin)
	}

	return admins, nil
}
