package resultLab

import (
	"log"
	"regexp"

	"github.com/gofiber/fiber/v2"

	resultLabHandler "view_lab/src/controllers/labview/resultLab/handler"
	resultLabService "view_lab/src/controllers/labview/resultLab/service"
	resultLabUtils "view_lab/src/controllers/labview/resultLab/utils"
)

var cidRegex = regexp.MustCompile(`^\d{13}$`)

func GetLabResults(c *fiber.Ctx) error {
	var req resultLabUtils.LabResultRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "รูปแบบข้อมูลไม่ถูกต้อง กรุณาส่งข้อมูลเป็น JSON",
		})
	}

	// ตรวจสอบ CID
	if req.CID == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "กรุณาระบุ cid",
		})
	}
	if !cidRegex.MatchString(req.CID) {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "cid ต้องเป็นตัวเลข 13 หลักเท่านั้น (ได้รับ: " + req.CID + ")",
		})
	}

	// ตรวจสอบว่ากรอกวันที่มาครบ
	if req.StartDate == "" || req.EndDate == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "กรุณาระบุ start_date และ end_date (รูปแบบ YYYY-MM-DD เช่น 2567-01-01)",
		})
	}

	// แปลงวันที่
	startDate, err := resultLabHandler.ParseDateString(req.StartDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "start_date \"" + req.StartDate + "\" ไม่ถูกต้อง รูปแบบที่รองรับ: YYYY-MM-DD (พ.ศ. หรือ ค.ศ.)",
		})
	}

	endDate, err := resultLabHandler.ParseDateString(req.EndDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "end_date \"" + req.EndDate + "\" ไม่ถูกต้อง รูปแบบที่รองรับ: YYYY-MM-DD (พ.ศ. หรือ ค.ศ.)",
		})
	}

	if endDate.Before(startDate) {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "end_date (" + req.EndDate + ") ต้องไม่น้อยกว่า start_date (" + req.StartDate + ")",
		})
	}

	log.Printf("🔍 GetLabResults CID=%s, start=%s, end=%s",
		req.CID, startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	results, err := resultLabService.GetLabResults(req.CID, startDate, endDate)
	if err != nil {
		log.Printf("❌ GetLabResults error CID=%s: %v", req.CID, err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "เกิดข้อผิดพลาดในการดึงข้อมูล",
		})
	}

	if len(results) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "ไม่พบข้อมูลผล lab สำหรับ CID นี้ในช่วงวันที่ที่ระบุ",
		})
	}

	ptName := ""
	if len(results) > 0 {
		ptName = results[0].PtName
	}

	return c.JSON(resultLabUtils.LabResultResponse{
		Success: true,
		Total:   len(results),
		PtName:  ptName,
		Results: results,
	})
}
