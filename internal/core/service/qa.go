package service

import (
    "errors"
    "time"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

type qaService struct {
    repo port.QARepository
}

func NewQAService(repo port.QARepository) port.QAService {
    return &qaService{repo: repo}
}

func (s *qaService) AskQuestion(userID uint, content string) (*domain.Question, error) {
    if content == "" {
        return nil, errors.New("question cannot be empty")
    }

    q := &domain.Question{
        Content:   content,
        UserID:    userID,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    
    if err := s.repo.SaveQuestion(q); err != nil {
        return nil, err
    }
    return q, nil
}

func (s *qaService) GetQuestions() ([]domain.Question, error) {
    return s.repo.FindAll()
}

func (s *qaService) AnswerQuestion(id uint, answer string) error {
    if answer == "" {
        return errors.New("answer cannot be empty")
    }
    return s.repo.UpdateAnswer(id, answer)
}