package middleware

import (
    "strings"
    
    "github.com/gofiber/fiber/v2"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

func AuthMiddleware(authService port.AuthService) fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(401).JSON(fiber.Map{"error": "missing authorization header"})
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        if tokenString == "" {
            return c.Status(401).JSON(fiber.Map{"error": "invalid authorization header"})
        }

        session, err := authService.ValidateJWT(tokenString)
        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
        }

        c.Locals("user", session)
        return c.Next()
    }
}

func AdminMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        user := c.Locals("user")
        if user == nil {
            return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
        }

        session, ok := user.(*domain.UserSession)
        if !ok || session.Role != domain.UserRoleAdmin {
            return c.Status(403).JSON(fiber.Map{"error": "admin access required"})
        }

        return c.Next()
    }
}