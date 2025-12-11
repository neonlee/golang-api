package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

// LogRepository interface
type LogRepository interface {
	Create(log *models.LogSistema) error
	GetByID(id uint) (*models.LogSistema, error)
	GetLogsByUsuario(usuarioID uint, limite int) ([]models.LogSistema, error)
	GetLogsByModulo(empresaID uint, modulo string, inicio, fim string) ([]models.LogSistema, error)
	GetLogsErro(empresaID uint, dias int) ([]models.LogSistema, error)
	// GetEst atisticasUso(empresaID uint, periodo string) (*responses.EstatisticasUso, error)
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(log *models.LogSistema) error {
	return r.db.Create(log).Error
}

func (r *logRepository) GetByID(id uint) (*models.LogSistema, error) {
	var log models.LogSistema
	err := r.db.
		Preload("Usuario").
		Preload("Empresa").
		First(&log, id).Error

	return &log, err
}

func (r *logRepository) GetLogsByUsuario(usuarioID uint, limite int) ([]models.LogSistema, error) {
	var logs []models.LogSistema

	err := r.db.
		Where("usuario_id = ?", usuarioID).
		Preload("Empresa").
		Order("created_at DESC").
		Limit(limite).
		Find(&logs).Error

	return logs, err
}

func (r *logRepository) GetLogsByModulo(empresaID uint, modulo string, inicio, fim string) ([]models.LogSistema, error) {
	var logs []models.LogSistema

	err := r.db.
		Where("empresa_id = ? AND modulo = ? AND DATE(created_at) BETWEEN ? AND ?",
			empresaID, modulo, inicio, fim).
		Preload("Usuario").
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

func (r *logRepository) GetLogsErro(empresaID uint, dias int) ([]models.LogSistema, error) {
	var logs []models.LogSistema

	err := r.db.
		Where("empresa_id = ? AND nivel_log = ? AND created_at >= NOW() - INTERVAL '? days'",
			empresaID, "ERROR", dias).
		Preload("Usuario").
		Order("created_at DESC").
		Find(&logs).Error

	return logs, err
}

// func (r *logRepository) GetEstatisticasUso(empresaID uint, periodo string) (*responses.EstatisticasUso, error) {
// 	var estatisticas responses.EstatisticasUso

// 	// Implementar lógica de estatísticas de uso
// 	// Exemplo: total de logs, logs por módulo, etc.

// 	return &estatisticas, nil
// }
