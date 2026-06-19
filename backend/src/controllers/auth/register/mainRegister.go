package registerMain

import (
	"database/sql"
	"log"

	handlerRegister "view_lab/src/controllers/auth/register/handler"
	serviceRegister "view_lab/src/controllers/auth/register/service"
	registerUtils "view_lab/src/controllers/auth/register/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(db *sql.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req registerUtils.RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "ข้อมูลไม่ถูกต้อง",
			})
		}

		if err := handlerRegister.ValidateCID(req.CID); err != nil {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if err := handlerRegister.ValidateEmail(req.Email); err != nil {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if err := handlerRegister.ValidateFullName(req.FullName); err != nil {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: err.Error(),
			})
		}

		if req.FacilityType == "" {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "กรุณาระบุประเภทสถานพยาบาล",
			})
		}

		if req.FacilityCode == "" {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "กรุณาระบุรหัสสถานพยาบาล",
			})
		}

		if req.FacilityName == "" {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "กรุณาระบุชื่อสถานพยาบาล",
			})
		}

		cidExists, err := handlerRegister.CheckCIDExists(req.CID)
		if err != nil {
			log.Printf("❌ CheckCIDExists error: %v", err)
			return c.Status(500).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "เกิดข้อผิดพลาดในการตรวจสอบข้อมูล",
			})
		}
		if cidExists {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "เลขบัตรประชาชนนี้มีในระบบแล้ว",
			})
		}

		emailExists, err := handlerRegister.CheckEmailExists(req.Email)
		if err != nil {
			log.Printf("❌ CheckEmailExists error: %v", err)
			return c.Status(500).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "เกิดข้อผิดพลาดในการตรวจสอบข้อมูล",
			})
		}
		if emailExists {
			return c.Status(400).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "อีเมลนี้มีในระบบแล้ว",
			})
		}

		// ตรวจสอบว่ามี admin ในระบบหรือไม่
		admins, err := serviceRegister.GetAdmins()
		if err != nil {
			log.Printf("❌ GetAdmins error: %v", err)
			return c.Status(500).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "เกิดข้อผิดพลาดในการตรวจสอบข้อมูล",
			})
		}
		if len(admins) == 0 && !serviceRegister.IsAdminCID(req.CID) {
			log.Printf("⚠️ พยายามลงทะเบียน แต่ระบบยังไม่มี admin")
			return c.Status(403).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "ระบบยังไม่มี Admin หรือ Super Admin ยังไม่สามารถลงทะเบียนได้ กรุณาติดต่อผู้ดูแลระบบ",
			})
		}

		user, err := serviceRegister.CreateUser(req.CID, req.FullName, req.Email, req.FacilityType, req.FacilityCode, req.FacilityName)
		if err != nil {
			log.Printf("❌ CreateUser error: %v", err)
			return c.Status(500).JSON(registerUtils.RegisterResponse{
				Success: false,
				Message: "ลงทะเบียนไม่สำเร็จ กรุณาลองใหม่อีกครั้ง",
			})
		}

		log.Printf("✅ ลงทะเบียนสำเร็จ: %s - %s", user.CID, user.FullName)

		// ส่งแจ้งเตือนไบยัง admin
		go func() {
			handlerRegister.NotifyAdminsNewRegistration(admins, user)
			log.Printf("📤 ส่งแจ้งเตือนไปยัง %d admin(s)", len(admins))
		}()

		return c.Status(201).JSON(registerUtils.RegisterResponse{
			Success: true,
			Message: "ลงทะเบียนสำเร็จ",
		})
	}
}
