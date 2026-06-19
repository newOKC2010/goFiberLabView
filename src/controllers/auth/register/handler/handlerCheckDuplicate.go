package handlerRegister

import (
	"fmt"

	conn "view_lab/src/database/connection"
)

func CheckCIDExists(cid string) (bool, error) {
	var count int
	err := conn.DB.QueryRow(
		`SELECT COUNT(1) FROM user_view_lab WHERE cid = $1`, cid,
	).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("ตรวจสอบเลขบัตรประชาชนไม่สำเร็จ: %w", err)
	}
	return count > 0, nil
}

func CheckEmailExists(email string) (bool, error) {
	var count int
	err := conn.DB.QueryRow(
		`SELECT COUNT(1) FROM user_view_lab WHERE LOWER(email) = LOWER($1)`, email,
	).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("ตรวจสอบอีเมลไม่สำเร็จ: %w", err)
	}
	return count > 0, nil
}
