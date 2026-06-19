package manageUserHandler

import (
	"fmt"

	manageUserService "view_lab/src/controllers/labview/manageUser/service"
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
)

// UpdateStatus validate และ update status ของ user
func UpdateStatus(userID int64, status *bool) (*manageUserUtils.UserNotifyInfo, error) {
	if status == nil {
		return nil, fmt.Errorf("กรุณาระบุ status")
	}
	return manageUserService.UpdateUserStatus(userID, *status)
}
