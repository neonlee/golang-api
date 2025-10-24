package repositories

import (
	"fmt"
	"petApi/internal/models"

	"gorm.io/gorm"
)

type FuncionariosRepository interface {
	Create(funcionario models.Funcionarios) (*models.Funcionarios, error)
	GetAllFuncionarios(empresaID uint) (*[]models.Funcionarios, error)
	GetFuncionarios(id int, empresaID uint) (*models.Funcionarios, error)
	UpdateFuncionarios(id int, funcionario *models.Funcionarios) error
	DeleteFuncionarios(id int) (bool, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewFuncionariosRepository(connection *gorm.DB) FuncionariosRepository {
	return &employeeRepository{db: connection}
}

func (r *employeeRepository) Create(employee models.Funcionarios) (*models.Funcionarios, error) {
	var usuario models.Usuarios
	err := r.db.Preload("Empresa").First(&usuario, employee.UsuarioID).Error
	if err != nil {
		return nil, fmt.Errorf("usuario com ID %d não existe", employee.UsuarioID)
	}

	err = r.db.Create(&employee).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &employee, nil
}

func (r *employeeRepository) GetAllFuncionarios(empresaID uint) (*[]models.Funcionarios, error) {
	var funcionarios []models.Funcionarios

	err := r.db.
		Where("empresa_id = ?", empresaID).
		Find(&funcionarios).Error

	return &funcionarios, err
}

func (r *employeeRepository) GetFuncionarios(id int, empresaID uint) (*models.Funcionarios, error) {
	var client models.Funcionarios
	err := r.db.
		Where("id = ?", id).
		Where("empresa_id = ?", empresaID).
		First(&client).Error

	return &client, err
}

func (r *employeeRepository) UpdateFuncionarios(id int, funcionario *models.Funcionarios) error {
	// Atualiza diretamente usando Where
	result := r.db.Model(&models.Funcionarios{}).Where("id = ?", id).Updates(funcionario)
	return result.Error
}

func (r *employeeRepository) DeleteFuncionarios(id int) (bool, error) {
	err := r.db.Delete(&models.Funcionarios{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
