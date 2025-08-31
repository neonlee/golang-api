package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	Controllers "petApi/internal/controllers"
	"petApi/internal/models"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEmployeeRepository é um mock do repositório para testes
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) UpdateEmployee(id int, employee models.Employee) (models.Employee, error) {
	args := m.Called(id, employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) GetEmployee(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) GetEmployees() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Create(employee models.Employee) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) DeleteEmployee(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestControllersEmployee_Update(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo)

	router := setupRouter()
	router.PUT("/employee/:id", controller.Update)

	// Teste de sucesso
	t.Run("Update employee successfully", func(t *testing.T) {
		employeeId := 1
		expectedEmployee := models.Employee{
			Id:   employeeId,
			Nome: "John Updated",
		}

		mockRepo.On("UpdateEmployee", employeeId, mock.AnythingOfType("models.Employee")).
			Return(expectedEmployee, nil)

		employeeJSON, _ := json.Marshal(expectedEmployee)
		req, _ := http.NewRequest("PUT", "/employee/"+strconv.Itoa(employeeId), bytes.NewBuffer(employeeJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Employee
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, expectedEmployee.Id, response.Id)
		assert.Equal(t, expectedEmployee.Nome, response.Nome)
		mockRepo.AssertExpectations(t)
	})

	// Teste com Id inválido
	t.Run("Update with invalid Id", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/employee/invalid", bytes.NewBuffer([]byte(`{}`)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Teste com JSON inválido
	t.Run("Update with invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest("PUT", "/employee/1", bytes.NewBuffer([]byte(`invalid json`)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestControllersEmployee_Get(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo)

	router := setupRouter()
	router.GET("/employee/:id", controller.Get)

	// Teste de sucesso
	t.Run("Get employee successfully", func(t *testing.T) {
		employeeId := 1
		expectedEmployee := models.Employee{
			Id:   employeeId,
			Nome: "John Doe",
		}

		mockRepo.On("GetEmployee", employeeId).Return(expectedEmployee, nil)

		req, _ := http.NewRequest("GET", "/employee/"+strconv.Itoa(employeeId), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Employee
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, expectedEmployee.Id, response.Id)
		assert.Equal(t, expectedEmployee.Nome, response.Nome)
		mockRepo.AssertExpectations(t)
	})

	// Teste com Id inválido
	t.Run("Get with invalid Id", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/employee/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestControllersEmployee_GetEmployee(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo)

	router := setupRouter()
	router.GET("/employees", controller.GetEmployee)

	t.Run("Get all employees successfully", func(t *testing.T) {
		expectedEmployees := []models.Employee{
			{Id: 1, Nome: "John Doe"},
			{Id: 2, Nome: "Jane Smith"},
		}

		mockRepo.On("GetEmployees").Return(expectedEmployees, nil)

		req, _ := http.NewRequest("GET", "/employees", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Employee
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Len(t, response, 2)
		assert.Equal(t, expectedEmployees[0].Nome, response[0].Nome)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get employees with repository error", func(t *testing.T) {
		mockRepo.On("GetEmployees").Return([]models.Employee{}, assert.AnError)

		req, _ := http.NewRequest("GET", "/employees", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestControllersEmployee_Create(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo)

	router := setupRouter()
	router.POST("/employee", controller.Create)

	t.Run("Create employee successfully", func(t *testing.T) {
		newEmployee := models.Employee{
			Nome: "New Employee",
		}
		createdEmployee := models.Employee{
			Id:   1,
			Nome: "New Employee",
		}

		mockRepo.On("Create", newEmployee).Return(createdEmployee, nil)

		employeeJSON, _ := json.Marshal(newEmployee)
		req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(employeeJSON))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Employee
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, createdEmployee.Id, response.Id)
		assert.Equal(t, createdEmployee.Nome, response.Nome)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create with invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer([]byte(`invalid json`)))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestControllersEmployee_Delete(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo)

	router := setupRouter()
	router.DELETE("/employee/:id", controller.Delete)

	t.Run("Delete employee successfully", func(t *testing.T) {
		employeeId := 1

		mockRepo.On("DeleteEmployee", employeeId).Return(true, nil)

		req, _ := http.NewRequest("DELETE", "/employee/"+strconv.Itoa(employeeId), nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response bool
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.True(t, response)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete with invalid Id", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/employee/invalid", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
