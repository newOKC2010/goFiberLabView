package middleware

import (
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthGuards(db *sql.DB, allowedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := authVerify(c, db, allowedRoles)
		if !auth.Success {
			code := fiber.StatusUnauthorized
			if strings.Contains(auth.Message, "ไม่มีสิทธิ์") {
				code = fiber.StatusForbidden
			}
			return c.Status(code).JSON(fiber.Map{"success": false, "message": auth.Message})
		}
		c.Locals("user_view_lab", auth.User)
		return c.Next()
	}
}

func authVerify(c *fiber.Ctx, db *sql.DB, allowedRoles []string) Response {
	token := extractToken(c)
	if token == "" {
		return Response{Success: false, Message: "ไม่พบ token"}
	}

	claims, err := validateJWT(token)
	if err != nil {
		return Response{Success: false, Message: "token ไม่ถูกต้องหรือหมดอายุ"}
	}

	if err := VerifyToken(db, claims.UserViewLabID, token); err != nil {
		return Response{Success: false, Message: "token หมดอายุหรือไม่ถูกต้อง"}
	}

	user, err := GetUserByID(db, claims.UserViewLabID)
	if err != nil {
		return Response{Success: false, Message: "ไม่พบผู้ใช้งาน"}
	}

	if !user.Status {
		return Response{Success: false, Message: "บัญชีผู้ใช้ถูกระงับ"}
	}

	errMsg := checkRole(user.Role, allowedRoles)
	if errMsg != "" {
		return Response{Success: false, Message: errMsg}
	}

	return Response{Success: true, User: user}
}
