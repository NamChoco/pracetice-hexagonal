package models

type AskQuestionRequest struct {
	Content string `json:"content"`
}

type AnswerQuestionRequest struct {
	Answer string `json:"answer"`
}
