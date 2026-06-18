package manageUserHandler

import (
	manageUserService "view_lab/src/controllers/labview/manageUser/service"
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
)

// GetUsers เรียก service ดึง users และ build response
func GetUsers() (manageUserUtils.GetUsersResponse, error) {
	users, err := manageUserService.GetAllUsers()
	if err != nil {
		return manageUserUtils.GetUsersResponse{}, err
	}

	if users == nil {
		users = []manageUserUtils.UserItem{}
	}

	return manageUserUtils.GetUsersResponse{
		Success: true,
		Total:   len(users),
		Users:   users,
	}, nil
}
