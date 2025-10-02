package requests

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}
