package googleoauth

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"

    "github.com/NamChoco/pracetice-hexagonal/internal/core/domain"
    "github.com/NamChoco/pracetice-hexagonal/internal/core/port"
)

type Client struct {
    cfg *oauth2.Config
}

func NewOAuthClient() port.OAuthClient {
    baseURL := os.Getenv("BASE_URL")
    clientID := os.Getenv("GOOGLE_CLIENT_ID")
    clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

    cfg := &oauth2.Config{
        ClientID:     clientID,
        ClientSecret: clientSecret,
        RedirectURL:  fmt.Sprintf("%s/auth/google/callback", baseURL),
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }
    return &Client{cfg: cfg}
}

func (c *Client) LoginURL(state string) string {
    return c.cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (c *Client) Exchange(ctx context.Context, code string) (domain.GoogleUser, error) {
    tok, err := c.cfg.Exchange(ctx, code)
    if err != nil {
        return domain.GoogleUser{}, err
    }

    httpClient := c.cfg.Client(ctx, tok)
    resp, err := httpClient.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        return domain.GoogleUser{}, err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return domain.GoogleUser{}, fmt.Errorf("userinfo status: %s", resp.Status)
    }

    var payload struct {
        ID      string `json:"id"`
        Email   string `json:"email"`
        Name    string `json:"name"`
        Picture string `json:"picture"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
        return domain.GoogleUser{}, err
    }
    
    return domain.GoogleUser{
        GoogleID: payload.ID,
        Name:     payload.Name,
        Email:    payload.Email,
        Picture:  payload.Picture,
    }, nil
}