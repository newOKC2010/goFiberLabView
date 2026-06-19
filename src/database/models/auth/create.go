package modelAuth

import "time"

type UserRole string

const (
	RoleUser       UserRole = "user"
	RoleAdmin      UserRole = "admin"
	RoleSuperAdmin UserRole = "super_admin"
)

type UserViewLab struct {
	ID           int64
	CID          string
	HashCID      string
	Role         UserRole
	Status       bool
	FullName     string
	Email        string
	FacilityType *string
	FacilityCode *string
	FacilityName *string
	OtpCode      *string
	OtpExpiresAt *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type TokensViewLab struct {
	ID            int64
	UserViewLabID int64
	Token         string
	ExpiresAt     time.Time
	LoginLast     time.Time
	UserViewLab   *UserViewLab
}
