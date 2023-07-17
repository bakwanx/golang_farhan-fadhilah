package controller

import (
	"belajar-go-echo/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func GetAllUsersTest(t *testing.T) {
	e := echo.New()

	mockDb := new(MockDB)

	expectedUsers := []model.User{}

	mockDb.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{})

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := GetAllUsers(mockDb)
	err := handler(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse := echo.Map{
		"data": expectedUsers,
	}
	jsonString, err := toJSONString(expectedResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	mockDb.AssertCalled(t, "Find", mock.Anything, mock.Anything)

}

func toJSONString(v interface{}) (string, error) {
	// Mengonversi objek menjadi JSON string
	jsonString, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}
