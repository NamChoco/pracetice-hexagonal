package models

type AskQuestionRequest struct {
    Content string `json:"content" validate:"required"`
}

type AnswerQuestionRequest struct {
    Answer string `json:"answer" validate:"required"`
}