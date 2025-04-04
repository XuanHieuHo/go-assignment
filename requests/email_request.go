package requests

type EmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}