package util

type OutlineResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}