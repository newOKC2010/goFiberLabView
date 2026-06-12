package middleware

import (
	"database/sql"
	"time"

	modelAuth "view_lab/src/database/models/auth"
)

func VerifyToken(db *sql.DB, userID int64, token string) error {
	var exists bool
	err := db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM tokens_view_lab 
		 WHERE user_view_lab_id = $1 AND token = $2 AND expires_at > $3)`,
		userID, token, time.Now(),
	).Scan(&exists)

	if err != nil || !exists {
		return err
	}
	return nil
}

func GetUserByID(db *sql.DB, userID int64) (*UserViewLabInfo, error) {
	user := &modelAuth.UserViewLab{}
	var role string

	err := db.QueryRow(
		`SELECT id, email, role, status FROM user_view_lab WHERE id = $1`,
		userID,
	).Scan(&user.ID, &user.Email, &role, &user.Status)

	if err != nil {
		return nil, err
	}

	userInfo := &UserViewLabInfo{
		ID:     user.ID,
		Email:  user.Email,
		Role:   role,
		Status: user.Status,
	}

	return userInfo, nil
}
