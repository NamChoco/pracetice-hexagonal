package controllers

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/NamChoco/pracetice-hexagonal/internal/adapter/fiber/models"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

type QAController struct {
    service port.QAService
}

func NewQAController(service port.QAService) *QAController {
    return &QAController{service: service}
}

func (c *QAController) AskQuestion(ctx *fiber.Ctx) error {
    user := ctx.Locals("user").(*domain.UserSession)
    
    req := new(models.AskQuestionRequest)
    if err := ctx.BodyParser(req); err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    if req.Content == "" {
        return ctx.Status(400).JSON(fiber.Map{"error": "content is required"})
    }

    q, err := c.service.AskQuestion(user.UserID, req.Content)
    if err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(q)
}

func (c *QAController) GetQuestions(ctx *fiber.Ctx) error {
    qs, err := c.service.GetQuestions()
    if err != nil {
        return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(qs)
}

func (c *QAController) AnswerQuestion(ctx *fiber.Ctx) error {
    user := ctx.Locals("user").(*domain.UserSession)
    if user.Role != domain.UserRoleAdmin {
        return ctx.Status(403).JSON(fiber.Map{"error": "admin access required"})
    }

    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": "invalid id"})
    }

    req := new(models.AnswerQuestionRequest)
    if err := ctx.BodyParser(req); err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    if req.Answer == "" {
        return ctx.Status(400).JSON(fiber.Map{"error": "answer is required"})
    }

    if err := c.service.AnswerQuestion(uint(id), req.Answer); err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(fiber.Map{"message": "answer saved successfully"})
}