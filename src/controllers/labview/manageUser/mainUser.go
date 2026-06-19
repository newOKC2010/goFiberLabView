package manageUser

import (
	"log"

	"github.com/gofiber/fiber/v2"

	manageUserHandler "view_lab/src/controllers/labview/manageUser/handler"
	manageUserUtils "view_lab/src/controllers/labview/manageUser/utils"
)

func GetUsers(c *fiber.Ctx) error {
	result, err := manageUserHandler.GetUsers()
	if err != nil {
		log.Printf("❌ GetUsers error: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "เกิดข้อผิดพลาดในการดึงข้อมูลผู้ใช้",
		})
	}

	return c.JSON(result)
}

func UpdateStatus(c *fiber.Ctx) error {
	var req manageUserUtils.UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "ข้อมูลไม่ถูกต้อง",
		})
	}

	if req.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "กรุณาระบุ id",
		})
	}

	userInfo, err := manageUserHandler.UpdateStatus(req.ID, req.Status)
	if err != nil {
		log.Printf("❌ UpdateStatus error userID=%d: %v", req.ID, err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	go manageUserHandler.NotifyUserStatusChange(userInfo)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "อัพเดท status สำเร็จ",
	})
}
