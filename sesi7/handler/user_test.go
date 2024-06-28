package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	t.Run("error binding", func(t *testing.T) {
		// init gin local service
		ge := gin.Default()
		userHandler := &UserHandler{}
		ge.POST("/users", userHandler.Create)
		// init http test mock
		w := httptest.NewRecorder()
		// Create an example user for testing
		example := map[string]any{
			"name": 1,
			"dob":  "2000-01-01",
		}
		body, _ := json.Marshal(example)
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(string(body)))
		req.Header.Add("Content-Type", "application/json")
		ge.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("error call service create", func(t *testing.T) {
		// create mock service
		userServiceMock := mocks.NewUserServiceInterface(t)
		userServiceMock.On("Create", mock.Anything).Return(errors.New("some errors"))

		// init gin local service
		ge := gin.Default()
		userHandler := &UserHandler{UserService: userServiceMock}
		ge.POST("/users", userHandler.Create)
		// init http test mock
		w := httptest.NewRecorder()
		// Create an example user for testing
		example := map[string]any{
			"name":     "user",
			"password": "mysecretpassword",
			"gender":   "MALE",
			"dob":      "2000-09-02T00:00:00.000Z",
		}
		body, _ := json.Marshal(example)
		req, _ := http.NewRequest("POST", "/users", strings.NewReader(string(body)))
		req.Header.Add("Content-Type", "application/json")
		ge.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
	})
}
