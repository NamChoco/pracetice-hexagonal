package port

import "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"

// Secondary Port
type QARepository interface {
	SaveQuestion(q *domain.Question) error
	FindAll() ([]domain.Question, error)
	FindByID(id uint) (*domain.Question, error)
	UpdateAnswer(id uint, answer string) error
}

// Primary Port
type QAService interface {
	AskQuestion(content string) (*domain.Question, error)
	GetQuestions() ([]domain.Question, error)
	AnswerQuestion(id uint, answer string) error
}
