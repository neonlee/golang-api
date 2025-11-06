package requests

type PeriodoFilter struct {
	Inicio string `json:"inicio" binding:"required,datetime=2006-01-02"`
	Fim    string `json:"fim" binding:"required,datetime=2006-01-02"`
}
