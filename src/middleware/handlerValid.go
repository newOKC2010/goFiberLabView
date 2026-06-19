package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	loadEnv "view_lab/src/loadenv"
)

func extractToken(c *fiber.Ctx) string {
	auth := c.Get("Authorization")
	if len(auth) > 7 && strings.HasPrefix(auth, "Bearer ") {
		return auth[7:]
	}
	return ""
}

func validateJWT(token string) (*JWTdecode, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &JWTdecode{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(loadEnv.LoadJWT().Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*JWTdecode)
	if !ok || !parsedToken.Valid {
		return nil, err
	}

	return claims, nil
}

func checkRole(usersRole string, allowedRoles []string) string {
	if len(allowedRoles) == 0 {
		return ""
	}

	for _, role := range allowedRoles {
		if role == usersRole {
			return ""
		}
	}

	if len(allowedRoles) == 1 {
		return "ไม่มีสิทธิ์เข้าถึง เฉพาะ " + allowedRoles[0] + " เท่านั้น"
	} else {
		return "ไม่มีสิทธิ์เข้าถึง เฉพาะ " + strings.Join(allowedRoles, ", ") + " เท่านั้น"
	}
}
