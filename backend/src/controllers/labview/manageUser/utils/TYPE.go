package manageUserUtils

import "time"

type UserItem struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUsersResponse struct {
	Success bool       `json:"success"`
	Total   int        `json:"total"`
	Users   []UserItem `json:"users"`
}

type UpdateStatusRequest struct {
	ID     int64 `json:"id"`
	Status *bool `json:"status"`
}

type UserNotifyInfo struct {
	FullName string
	Email    string
	CID      string
	Status   bool
}
