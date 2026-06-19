package manageUserService

import (
	"database/sql"
	"fmt"
	"time"

	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
	conn "view_lab/src/database/connection"
)

// UpdateUserStatus อัพเดท status ของ user ตาม id (เฉพาะ role user)
func UpdateUserStatus(userID int64, status bool) (*manageUserUtils.UserNotifyInfo, error) {
	// ดึงข้อมูล user พร้อมตรวจสอบ role
	var info manageUserUtils.UserNotifyInfo
	var role string
	err := conn.DB.QueryRow(
		`SELECT full_name, email, cid, role FROM user_view_lab WHERE id = $1`,
		userID,
	).Scan(&info.FullName, &info.Email, &info.CID, &role)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ไม่พบผู้ใช้ id=%d", userID)
		}
		return nil, err
	}

	// ห้าม update status ของ admin/super_admin
	if role == "admin" || role == "super_admin" {
		return nil, fmt.Errorf("ไม่สามารถแก้ไข status ของ %s ได้", role)
	}

	result, err := conn.DB.Exec(
		`UPDATE user_view_lab SET status = $1, updated_at = $2 WHERE id = $3 AND role = 'user'`,
		status, time.Now(), userID,
	)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, fmt.Errorf("อัพเดทไม่สำเร็จ")
	}

	info.Status = status
	return &info, nil
}
