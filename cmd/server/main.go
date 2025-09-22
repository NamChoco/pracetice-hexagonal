package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/routes"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/sqlite"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using system environment variables")
    }

    // Validate required environment variables
    requiredEnvVars := []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "BASE_URL", "JWT_SECRET"}
    for _, envVar := range requiredEnvVars {
        if os.Getenv(envVar) == "" {
            log.Fatalf("Required environment variable %s is not set", envVar)
        }
    }

    // Setup DB
    db, err := sqlite.InitDB("qa.db")
    if err != nil {
        log.Fatal(err)
    }

    // Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }
            return ctx.Status(code).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })

    // Middleware
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:3000,http://127.0.0.1:3000",
        AllowCredentials: true,
        AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
    }))

    // Register routes
    routes.RegisterAllRoutes(app, db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("Server running on http://localhost:%s", port)
    log.Println("OAuth endpoints:")
    log.Println("  - Login: GET /auth/google/login")
    log.Println("  - Callback: GET /auth/google/callback")
    log.Println("  - Profile: GET /auth/profile")
    log.Println("  - Logout: POST /auth/logout")
    log.Println("QA endpoints:")
    log.Println("  - Ask Question: POST /qa/")
    log.Println("  - Get Questions: GET /qa/")
    log.Println("  - Answer Question: PUT /qa/:id/answer")
    
    if err := app.Listen(":" + port); err != nil {
        log.Fatal(err)
    }
}