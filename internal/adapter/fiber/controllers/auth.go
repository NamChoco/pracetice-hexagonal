package controllers

import (
    "crypto/rand"
    "encoding/base64"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/models"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

type AuthController struct {
    authService port.AuthService
}

func NewAuthController(authService port.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (c *AuthController) GoogleLogin(ctx *fiber.Ctx) error {
    state := randomState()
    
    // Set state in cookie for validation (like working example)
    ctx.Cookie(&fiber.Cookie{
        Name:     "oauth_state",
        Value:    state,
        HTTPOnly: true,
        SameSite: "Lax",
        Expires:  time.Now().Add(10 * time.Minute),
    })
    
    url := c.authService.GetAuthURL(state)
    
    log.Printf("🔐 Generated OAuth URL with state: %s", state[:8]+"...")
    
    // Redirect like working example (not JSON response)
    return ctx.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (c *AuthController) GoogleCallback(ctx *fiber.Ctx) error {
    receivedState := ctx.Query("state")
    cookieState := ctx.Cookies("oauth_state")
    
    log.Printf("📞 Callback received:")
    log.Printf("  - State from URL: %s", func() string {
        if len(receivedState) > 10 { return receivedState[:10] + "..." }
        return receivedState
    }())
    log.Printf("  - State from Cookie: %s", func() string {
        if len(cookieState) > 10 { return cookieState[:10] + "..." }
        return cookieState
    }())
    
    // Validate state like working example
    if receivedState == "" || receivedState != cookieState {
        log.Printf("❌ State validation failed")
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid state"})
    }
    
    code := ctx.Query("code")
    if code == "" {
        log.Printf("❌ Missing authorization code")
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing code"})
    }
    
    log.Printf("✅ State validated successfully")
    
    // Clear the state cookie
    ctx.Cookie(&fiber.Cookie{
        Name:     "oauth_state",
        Value:    "",
        HTTPOnly: true,
        Expires:  time.Now().Add(-1 * time.Hour),
    })
    
    user, err := c.authService.HandleCallback(ctx.Context(), code)
    if err != nil {
        log.Printf("❌ OAuth callback failed: %s", err.Error())
        return ctx.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
    }

    log.Printf("✅ User authenticated: %s (%s)", user.Name, user.Email)

    token, err := c.authService.GenerateJWT(user)
    if err != nil {
        log.Printf("❌ JWT generation failed: %s", err.Error())
        return ctx.Status(500).JSON(fiber.Map{"error": "failed to generate token"})
    }

    log.Printf("🎫 JWT token generated successfully")

    response := models.LoginResponse{
        Token: token,
    }
    response.User.ID = user.ID
    response.User.Email = user.Email
    response.User.Name = user.Name
    response.User.Picture = user.Picture
    response.User.Role = string(user.Role)

    return ctx.JSON(response)
}

func (c *AuthController) GetProfile(ctx *fiber.Ctx) error {
    user := ctx.Locals("user").(*domain.UserSession)
    log.Printf("👤 Profile requested for: %s", user.Email)
    return ctx.JSON(user)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
    log.Printf("🚪 User logged out")
    return ctx.JSON(fiber.Map{"message": "logged out successfully"})
}

// Generate random state like working example
func randomState() string {
    b := make([]byte, 24)
    _, _ = rand.Read(b)
    return base64.RawURLEncoding.EncodeToString(b)
}