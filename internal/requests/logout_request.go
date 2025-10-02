package requests

type LogoutRequest struct {
	Token string `json:"token" binding:"required"`
}
