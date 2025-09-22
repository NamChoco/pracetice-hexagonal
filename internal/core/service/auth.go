package service

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

type authService struct {
    userRepo    port.UserRepository
    oauthClient port.OAuthClient
}

func NewAuthService(userRepo port.UserRepository, oauthClient port.OAuthClient) port.AuthService {
    return &authService{
        userRepo:    userRepo,
        oauthClient: oauthClient,
    }
}

func (s *authService) GetAuthURL(state string) string {
    return s.oauthClient.LoginURL(state)
}

func (s *authService) HandleCallback(ctx context.Context, code string) (*domain.User, error) {
    googleUser, err := s.oauthClient.Exchange(ctx, code)
    if err != nil {
        return nil, fmt.Errorf("failed to exchange code for token: %w", err)
    }

    return s.CreateOrUpdateUser(googleUser)
}

func (s *authService) CreateOrUpdateUser(googleUser domain.GoogleUser) (*domain.User, error) {
    user, err := s.userRepo.FindByGoogleID(googleUser.GoogleID)
    if err == nil {
        // Update existing user
        user.Email = googleUser.Email
        user.Name = googleUser.Name
        user.Picture = googleUser.Picture
        user.UpdatedAt = time.Now()
        if err := s.userRepo.UpdateUser(user); err != nil {
            return nil, fmt.Errorf("failed to update user: %w", err)
        }
        return user, nil
    }

    // Create new user
    user = &domain.User{
        GoogleID:  googleUser.GoogleID,
        Email:     googleUser.Email,
        Name:      googleUser.Name,
        Picture:   googleUser.Picture,
        Role:      domain.UserRoleUser,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    if err := s.userRepo.CreateUser(user); err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return user, nil
}

func (s *authService) GenerateJWT(user *domain.User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "name":    user.Name,
        "role":    user.Role,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (s *authService) ValidateJWT(tokenString string) (*domain.UserSession, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        return nil, fmt.Errorf("failed to parse token: %w", err)
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    userID, ok := claims["user_id"].(float64)
    if !ok {
        return nil, fmt.Errorf("invalid user_id in token")
    }

    email, ok := claims["email"].(string)
    if !ok {
        return nil, fmt.Errorf("invalid email in token")
    }

    name, ok := claims["name"].(string)
    if !ok {
        return nil, fmt.Errorf("invalid name in token")
    }

    role, ok := claims["role"].(string)
    if !ok {
        return nil, fmt.Errorf("invalid role in token")
    }

    exp, ok := claims["exp"].(float64)
    if !ok {
        return nil, fmt.Errorf("invalid exp in token")
    }

    return &domain.UserSession{
        UserID:    uint(userID),
        Email:     email,
        Name:      name,
        Role:      domain.UserRole(role),
        ExpiresAt: time.Unix(int64(exp), 0),
    }, nil
}