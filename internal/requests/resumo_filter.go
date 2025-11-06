package requests

type ResumoFilter struct {
	Periodo string `json:"periodo" binding:"required"`
}
