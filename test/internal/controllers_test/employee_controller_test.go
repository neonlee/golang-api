package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	Controllers "petApi/internal/controllers"
	"petApi/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMain configuração global
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestEmployeeController_Get_Success(t *testing.T) {
	// Setup
	router, mockRepo, controller := setupTest()

	// Mock expectation - NOTE THE & (pointer)
	expectedEmployee := &models.Employees{
		Id:   1,
		Nome: "John Doe",
		// include other required fields
	}
	mockRepo.On("GetEmployee", int(1)).Return(expectedEmployee, nil)

	// Create router and request
	router.GET("/employee/:id", controller.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employee/1", nil)

	// Execute
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, 200, w.Code)

	var response models.Employees
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, int(1), response.Id)
	assert.Equal(t, "John Doe", response.Nome)

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestEmployeeController_Get_InvalidID(t *testing.T) {
	router, mockRepo, controller := setupTest()

	router.GET("/employee/:id", controller.Get)

	req, _ := http.NewRequest("GET", "/employee/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertNotCalled(t, "GetEmployee")
}
func TestEmployeeController_Create_Success(t *testing.T) {
	router, mockRepo, controller := setupTest()

	testEmployee := createTestEmployee()

	// Mock do repositório
	mockRepo.On("Create", testEmployee).Return(&testEmployee, nil)

	router.POST("/employee", controller.Create)

	// Criar request
	employeeJSON, _ := json.Marshal(testEmployee)
	req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(employeeJSON))
	req.Header.Set("Content-Type", "application/json")

	// Executar request
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verificações
	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Employees
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, testEmployee.Id, response.Id)
	assert.Equal(t, testEmployee.Nome, response.Nome)
	mockRepo.AssertExpectations(t)
}

func TestEmployeeController_Create_InvalidJSON(t *testing.T) {
	router, mockRepo, controller := setupTest()

	router.POST("/employee", controller.Create)

	// JSON inválido
	invalidJSON := `{"name": "John", "email": invalid}`
	req, _ := http.NewRequest("POST", "/employee", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertNotCalled(t, "Create")
}

// MockEmployeeRepository é um mock do repositório
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) GetEmployee(id int) (*models.Employees, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Employees), args.Error(1)
}

func (m *MockEmployeeRepository) GetEmployees() (*[]models.Employees, error) {
	args := m.Called()
	return args.Get(0).(*[]models.Employees), args.Error(1)
}

func (m *MockEmployeeRepository) Create(employee models.Employees) (*models.Employees, error) {
	args := m.Called(employee)
	return args.Get(0).(*models.Employees), args.Error(1)
}

func (m *MockEmployeeRepository) UpdateEmployee(id int, employee models.Employees) (*models.Employees, error) {
	args := m.Called(id, employee)
	return args.Get(0).(*models.Employees), args.Error(1)
}

func (m *MockEmployeeRepository) DeleteEmployee(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

// TestSetup configura o ambiente de teste
func setupTest() (*gin.Engine, *MockEmployeeRepository, *Controllers.EmployeeController) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockEmployeeRepository)
	controller := Controllers.NewEmployeeController(mockRepo) // pass as interface, not pointer to interface

	router := gin.Default()
	return router, mockRepo, controller
}

// createTestEmployee cria um employee de teste
func createTestEmployee() models.Employees {
	return models.Employees{
		Id:   1,
		Nome: "John Doe",
	}
}
func TestEmployeeController_Create_RepositoryError(t *testing.T) {
	router, mockRepo, controller := setupTest()

	testEmployee := createTestEmployee()
	expectedError := errors.New("database error")

	mockRepo.On("Create", testEmployee).Return(models.Employees{}, expectedError)

	router.POST("/employee", controller.Create)

	employeeJSON, _ := json.Marshal(testEmployee)
	req, _ := http.NewRequest("POST", "/employee", bytes.NewBuffer(employeeJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Contains(t, response["error"], "database error")
	mockRepo.AssertExpectations(t)
}
func TestEmployeeController_Get_NotFound(t *testing.T) {
	router, mockRepo, controller := setupTest()

	expectedError := errors.New("employee not found")
	mockRepo.On("GetEmployee", 999).Return(models.Employees{}, expectedError)

	router.GET("/employee/:id", controller.Get)

	req, _ := http.NewRequest("GET", "/employee/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Contains(t, response["error"], "not found")
	mockRepo.AssertExpectations(t)
}

func TestEmployeeController_Update_Success(t *testing.T) {
	router, mockRepo, controller := setupTest()

	testEmployee := createTestEmployee()
	updatedEmployee := testEmployee
	updatedEmployee.Nome = "John Updated"

	mockRepo.On("UpdateEmployee", 1, mock.Anything).Return(updatedEmployee, nil)

	router.PUT("/employee/:id", controller.Update)

	employeeJSON, _ := json.Marshal(updatedEmployee)
	req, _ := http.NewRequest("PUT", "/employee/1", bytes.NewBuffer(employeeJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.Employees
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "John Updated", response.Nome)
	mockRepo.AssertExpectations(t)
}

func TestEmployeeController_Update_InvalidID(t *testing.T) {
	router, mockRepo, controller := setupTest()

	router.PUT("/employee/:id", controller.Update)

	testEmployee := createTestEmployee()
	employeeJSON, _ := json.Marshal(testEmployee)

	req, _ := http.NewRequest("PUT", "/employee/abc", bytes.NewBuffer(employeeJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertNotCalled(t, "UpdateEmployee")
}

func TestEmployeeController_Update_InvalidJSON(t *testing.T) {
	router, mockRepo, controller := setupTest()

	router.PUT("/employee/:id", controller.Update)

	invalidJSON := `{"name": "John", "email": invalid}`
	req, _ := http.NewRequest("PUT", "/employee/1", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertNotCalled(t, "UpdateEmployee")
}
func TestEmployeeController_Delete_Success(t *testing.T) {
	router, mockRepo, controller := setupTest()

	mockRepo.On("DeleteEmployee", 1).Return(true, nil)

	router.DELETE("/employee/:id", controller.Delete)

	req, _ := http.NewRequest("DELETE", "/employee/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]bool
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.True(t, response["deleted"])
	mockRepo.AssertExpectations(t)
}

func TestEmployeeController_Delete_InvalidID(t *testing.T) {
	router, mockRepo, controller := setupTest()

	router.DELETE("/employee/:id", controller.Delete)

	req, _ := http.NewRequest("DELETE", "/employee/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertNotCalled(t, "DeleteEmployee")
}

func TestEmployeeController_Delete_Error(t *testing.T) {
	router, mockRepo, controller := setupTest()

	expectedError := errors.New("delete failed")
	mockRepo.On("DeleteEmployee", 999).Return(false, expectedError)

	router.DELETE("/employee/:id", controller.Delete)

	req, _ := http.NewRequest("DELETE", "/employee/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRepo.AssertExpectations(t)
}
