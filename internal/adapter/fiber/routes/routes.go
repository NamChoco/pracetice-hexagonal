package routes

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/controllers"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/middleware"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/googleoauth"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/sqlite"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/service"
)

func RegisterAllRoutes(app *fiber.App, db *gorm.DB) {
    // Initialize repositories
    userRepo := sqlite.NewUserRepository(db)
    qaRepo := sqlite.NewQARepository(db)

    // Initialize OAuth client
    oauthClient := googleoauth.NewOAuthClient()

    // Initialize services
    authService := service.NewAuthService(userRepo, oauthClient)
    qaService := service.NewQAService(qaRepo)

    // Initialize controllers
    authController := controllers.NewAuthController(authService)
    qaController := controllers.NewQAController(qaService)

    // Home route (like working example)
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("OK: /auth/google/login → sign in with Google")
    })

    // Auth routes (public)
    auth := app.Group("/auth")
    auth.Get("/google/login", authController.GoogleLogin)
    auth.Get("/google/callback", authController.GoogleCallback)
    
    // Protected auth routes
    authProtected := auth.Use(middleware.AuthMiddleware(authService))
    authProtected.Get("/profile", authController.GetProfile)
    authProtected.Post("/logout", authController.Logout)

    // QA routes (protected)
    qa := app.Group("/qa", middleware.AuthMiddleware(authService))
    qa.Post("/", qaController.AskQuestion)
    qa.Get("/", qaController.GetQuestions)
    qa.Put("/:id/answer", qaController.AnswerQuestion) // Only admins can answer
    
    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
            "message": "Q&A System with OAuth 2.0 is running",
        })
    })
}