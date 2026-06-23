package manageUserHandler

import (
	"math"

	manageUserService "view_lab/src/controllers/labview/manageUser/service"
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
)

// GetUsers เรียก service ดึง users แบบ offset pagination
func GetUsers(page, pageSize int, search string) (manageUserUtils.GetUsersResponse, error) {
	users, totalCount, err := manageUserService.GetAllUsers(page, pageSize, search)
	if err != nil {
		return manageUserUtils.GetUsersResponse{}, err
	}

	if users == nil {
		users = []manageUserUtils.UserItem{}
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	if totalPages < 1 {
		totalPages = 1
	}

	return manageUserUtils.GetUsersResponse{
		Success:     true,
		Message:     "ดึงข้อมูลสำเร็จ",
		Data:        users,
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: page,
	}, nil
}
