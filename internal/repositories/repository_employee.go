package repositories

import (
	"petApi/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	connection *gorm.DB
}

func NewEmployeeRepository(connection *gorm.DB) EmployeeRepository {
	return EmployeeRepository{connection: connection}
}

func (r *EmployeeRepository) Create(user models.Employee) (*models.Employee, error) {
	err := r.connection.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &user, nil
}

func (r *EmployeeRepository) GetEmployees() (*[]models.Employee, error) {
	var clientes []models.Employee

	err := r.connection.
		Preload("Pets").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *EmployeeRepository) GetEmployee(id int) (*models.Employee, error) {
	var client models.Employee
	var user models.Users
	err := r.connection.Preload("TipoAcesso.Modulos").
		Where("id = ?", id).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *EmployeeRepository) UpdateEmployee(id int, cliente models.Employee) (*models.Employee, error) {
	// Aplica o update direto pelo ID
	err := r.connection.Model(&models.Employee{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedEmployee models.Employee
	err = r.connection.First(&updatedEmployee, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedEmployee, nil
}

func (r *EmployeeRepository) DeleteEmployee(id int) (bool, error) {
	err := r.connection.Delete(&models.Employee{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
