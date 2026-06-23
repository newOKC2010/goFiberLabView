package manageUserUtils

import "time"

type UserItem struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type GetUsersQuery struct {
	Page     int    `query:"page"`
	PageSize int    `query:"page_size"`
	Search   string `query:"search"`
}

type GetUsersResponse struct {
	Success     bool       `json:"success"`
	Message     string     `json:"message"`
	Data        []UserItem `json:"data"`
	TotalCount  int        `json:"total_count"`
	TotalPages  int        `json:"total_pages"`
	CurrentPage int        `json:"current_page"`
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
