package repositories

import (
	"fmt"
	"petApi/internal/models"

	"gorm.io/gorm"
)

type EmployeesRepository interface {
	Create(user models.Employees) (*models.Employees, error)
	GetAllEmployees() (*[]models.Employees, error)
	GetEmployees(id int) (*models.Employees, error)
	UpdateEmployees(id int, cliente models.Employees) (*models.Employees, error)
	DeleteEmployees(id int) (bool, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeesRepository(connection *gorm.DB) EmployeesRepository {
	return &employeeRepository{db: connection}
}

func (r *employeeRepository) Create(employee models.Employees) (*models.Employees, error) {
	var user models.Usuarios
	if err := r.db.First(&user, employee.UserID).Error; err != nil {
		return nil, fmt.Errorf("categoria com ID %d não existe", employee.UserID)
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	// query.Close()
	return &employee, nil
}

func (r *employeeRepository) GetAllEmployees() (*[]models.Employees, error) {
	var clientes []models.Employees

	err := r.db.
		Preload("User").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *employeeRepository) GetEmployees(id int) (*models.Employees, error) {
	var client models.Employees
	err := r.db.Preload("User").
		Where("id = ?", id).
		First(&client).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *employeeRepository) UpdateEmployees(id int, cliente models.Employees) (*models.Employees, error) {
	// Aplica o update direto pelo ID
	err := r.db.Model(&models.Employees{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedEmployees models.Employees
	err = r.db.First(&updatedEmployees, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedEmployees, nil
}

func (r *employeeRepository) DeleteEmployees(id int) (bool, error) {
	err := r.db.Delete(&models.Employees{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
