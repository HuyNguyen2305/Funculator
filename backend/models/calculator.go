package models

type CalculateRequest struct {
	Expression string `json:"expression" binding:"required"`
}

type CalculateResponse struct {
	Expression string  `json:"expression"`
	Result     float64 `json:"result"`
}
