package port

import (
    "context"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
)

// Secondary Port - Repository
type UserRepository interface {
    CreateUser(user *domain.User) error
    FindByGoogleID(googleID string) (*domain.User, error)
    FindByEmail(email string) (*domain.User, error)
    FindByID(id uint) (*domain.User, error)
    UpdateUser(user *domain.User) error
}

// Primary Port - Service
type AuthService interface {
    GetAuthURL(state string) string
    HandleCallback(ctx context.Context, code string) (*domain.User, error)
    GenerateJWT(user *domain.User) (string, error)
    ValidateJWT(token string) (*domain.UserSession, error)
    CreateOrUpdateUser(googleUser domain.GoogleUser) (*domain.User, error)
}