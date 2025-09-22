package sqlite

import (
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
    "gorm.io/gorm"
)

type qaRepository struct {
    db *gorm.DB
}

func NewQARepository(db *gorm.DB) port.QARepository {
    db.AutoMigrate(&domain.Question{})
    return &qaRepository{db: db}
}

func (r *qaRepository) SaveQuestion(q *domain.Question) error {
    return r.db.Create(q).Error
}

func (r *qaRepository) FindAll() ([]domain.Question, error) {
    var questions []domain.Question
    err := r.db.Preload("User").Find(&questions).Error
    return questions, err
}

func (r *qaRepository) FindByID(id uint) (*domain.Question, error) {
    var q domain.Question
    if err := r.db.Preload("User").First(&q, id).Error; err != nil {
        return nil, err
    }
    return &q, nil
}

func (r *qaRepository) UpdateAnswer(id uint, answer string) error {
    return r.db.Model(&domain.Question{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "answer":     answer,
            "updated_at": "now()",
        }).Error
}