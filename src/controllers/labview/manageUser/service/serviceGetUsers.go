package manageUserService

import (
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
	conn "view_lab/src/database/connection"
)

// GetAllUsers ดึงรายชื่อผู้ใช้แบบ offset pagination (เฉพาะ role user)
func GetAllUsers(page, pageSize int, search string) ([]manageUserUtils.UserItem, int, error) {
	offset := (page - 1) * pageSize
	searchParam := "%" + search + "%"

	var totalCount int
	err := conn.DB.QueryRow(
		`SELECT COUNT(*) FROM user_view_lab WHERE role = 'user' AND full_name ILIKE $1`,
		searchParam,
	).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	rows, err := conn.DB.Query(
		`SELECT id, full_name, facility_name, status, created_at
		 FROM user_view_lab
		 WHERE role = 'user' AND full_name ILIKE $1
		 ORDER BY created_at DESC
		 LIMIT $2 OFFSET $3`,
		searchParam, pageSize, offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []manageUserUtils.UserItem
	for rows.Next() {
		var u manageUserUtils.UserItem
		if err := rows.Scan(&u.ID, &u.FullName, &u.FacilityName, &u.Status, &u.CreatedAt); err != nil {
			return nil, 0, err
		}
		users = append(users, u)
	}

	return users, totalCount, nil
}
