package middleware

import (
	"log"
	"rhmn-coffe/shared/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) fiber.Handler
}

type authMiddleware struct {
	jwtService service.JwtService
}

type AuthHeader struct {
	AuthorizationHeader string
}

func (a *authMiddleware) RequireToken(roles ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		tokenHeader := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenHeader == "" {
			log.Println("RequireToken: Missing token")
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		claims, err := a.jwtService.ValidateToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken: Error parsing token: %v \n", err)
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		ctx.Locals("employee", claims.UserId)

		role := claims.Role
		if role == "" {
			log.Println("RequireToken: Missing role in token")
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		if !isValidRole(role, roles) {
			log.Println("RequireToken: Invalid role")
			return ctx.Status(fiber.StatusForbidden).SendString("Forbidden")
		}

		return ctx.Next()
	}
}

func isValidRole(userRole string, validRoles []string) bool {
	for _, role := range validRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
