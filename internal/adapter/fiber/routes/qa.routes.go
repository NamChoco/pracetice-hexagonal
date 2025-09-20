package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/controllers"
	"github.com/NamChoco/pracetice-hexagonal/internal/adapter/sqlite"
	"github.com/NamChoco/pracetice-hexagonal/internal/core/service"
)

func RegisterQARoutes(app *fiber.App, db *gorm.DB) {
	repo := sqlite.NewQARepository(db)
	svc := service.NewQAService(repo)
	ctrl := controllers.NewQAController(svc)

	api := app.Group("/qa")
	api.Post("/", ctrl.AskQuestion)
	api.Get("/", ctrl.GetQuestions)
	api.Put("/:id/answer", ctrl.AnswerQuestion)
}
