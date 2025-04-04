package requests

type CreateFriendshipRequest struct {
	FristEmail string `json:"frist_email" binding:"required,email"`
	SecondEmail string `json:"second_email" binding:"required,email"`
}