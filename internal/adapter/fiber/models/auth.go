package models

type LoginResponse struct {
    Token string `json:"token"`
    User  struct {
        ID      uint   `json:"id"`
        Email   string `json:"email"`
        Name    string `json:"name"`
        Picture string `json:"picture"`
        Role    string `json:"role"`
    } `json:"user"`
}
