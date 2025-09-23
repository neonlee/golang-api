package repositories

import (
	"fmt"
	"petApi/internal/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	Create(user models.Employee) (*models.Employee, error)
	GetEmployees() (*[]models.Employee, error)
	GetEmployee(id int) (*models.Employee, error)
	UpdateEmployee(id int, cliente models.Employee) (*models.Employee, error)
	DeleteEmployee(id int) (bool, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(connection *gorm.DB) EmployeeRepository {
	return &employeeRepository{db: connection}
}

func (r *employeeRepository) Create(employee models.Employee) (*models.Employee, error) {
	var user models.Usuario
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

func (r *employeeRepository) GetEmployees() (*[]models.Employee, error) {
	var clientes []models.Employee

	err := r.db.
		Preload("User").
		Find(&clientes).Error
	if err != nil {
		return nil, err
	}

	return &clientes, nil
}

func (r *employeeRepository) GetEmployee(id int) (*models.Employee, error) {
	var client models.Employee
	err := r.db.Preload("User").
		Where("id = ?", id).
		First(&client).Error

	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *employeeRepository) UpdateEmployee(id int, cliente models.Employee) (*models.Employee, error) {
	// Aplica o update direto pelo ID
	err := r.db.Model(&models.Employee{}).Where("id = ?", id).Updates(cliente).Error
	if err != nil {
		return nil, err
	}

	// Busca o cliente atualizado
	var updatedEmployee models.Employee
	err = r.db.First(&updatedEmployee, id).Error
	if err != nil {
		return nil, err
	}

	return &updatedEmployee, nil
}

func (r *employeeRepository) DeleteEmployee(id int) (bool, error) {
	err := r.db.Delete(&models.Employee{}, id).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
