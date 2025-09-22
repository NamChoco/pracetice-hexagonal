package sqlite

import (
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
    "gorm.io/gorm"
)

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
    db.AutoMigrate(&domain.User{})
    return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByGoogleID(googleID string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("google_id = ?", googleID).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) FindByID(id uint) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) error {
    return r.db.Save(user).Error
}