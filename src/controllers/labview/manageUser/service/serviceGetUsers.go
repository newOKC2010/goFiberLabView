package manageUserService

import (
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
	conn "view_lab/src/database/connection"
)

// GetAllUsers ดึงรายชื่อผู้ใช้ทั้งหมดจาก DB (เฉพาะ role user)
func GetAllUsers() ([]manageUserUtils.UserItem, error) {
	rows, err := conn.DB.Query(
		`SELECT id, full_name, status, created_at
		 FROM user_view_lab
		 WHERE role = 'user'
		 ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []manageUserUtils.UserItem
	for rows.Next() {
		var u manageUserUtils.UserItem
		if err := rows.Scan(&u.ID, &u.FullName, &u.Status, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
