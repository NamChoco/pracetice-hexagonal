package port

import (
    "context"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
)

// Secondary Port - OAuth Client
type OAuthClient interface {
    LoginURL(state string) string
    Exchange(ctx context.Context, code string) (domain.GoogleUser, error)
}