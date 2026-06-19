package middleware

import "github.com/golang-jwt/jwt/v5"

type UserViewLabInfo struct {
	ID     int64  `db:"id" json:"user_view_lab_id"`
	CID    string `db:"cid" json:"cid"`
	Email  string `db:"email" json:"email"`
	Role   string `db:"role" json:"role"`
	Status bool   `db:"status" json:"-"`
}

type Response struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	User    *UserViewLabInfo `json:"user"`
}

type JWTdecode struct {
	UserViewLabID int64  `json:"user_view_lab_id"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	jwt.RegisteredClaims
}
