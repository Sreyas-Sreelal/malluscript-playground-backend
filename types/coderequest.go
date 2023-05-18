package types

type CodeRequest struct {
	Code  string `json:"code" binding:"required"`
	Input string `json:"input"`
}
