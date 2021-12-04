package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/p12s/library-rest-api/library/pkg/grpcClient"
	"github.com/p12s/library-rest-api/library/pkg/models"
	"github.com/p12s/library-rest-api/library/pkg/queueClient"
	"github.com/p12s/library-rest-api/library/pkg/service"
	mock_service "github.com/p12s/library-rest-api/library/pkg/service/mocks"
)

const DEFAULT_ENTITY_ID = 1

func TestHandler_signUp(t *testing.T) {

	type mockBehavior func(s *mock_service.MockAuthorization, user models.User)

	tests := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Can sign up with correct input",
			inputBody: `{"name": "Ivan", "username": "ivan", "password": "qwerty"}`,
			inputUser: models.User{
				Name:     "Ivan",
				Username: "ivan",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Can't sign up with input without name",
			inputBody:           `{"username": "ivan", "password": "qwerty"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user models.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Can't sign up with input without username",
			inputBody:           `{"password": "qwerty"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user models.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Can't sign up with input without password",
			inputBody:           `{}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, user models.User) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can return error response if service failure",
			inputBody: `{"name": "Ivan", "username": "ivan", "password": "qwerty"}`,
			inputUser: models.User{
				Name:     "Ivan",
				Username: "ivan",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failure"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			tt.mockBehavior(auth, tt.inputUser)

			services := &service.Service{Authorization: auth}

			var grpcLogger *grpcClient.Client
			var queueLogger *queueClient.QueueClient

			handler := NewHandler(services, grpcLogger, queueLogger)
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/auth/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-up",
				bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {

	type mockBehavior func(s *mock_service.MockAuthorization, input signInInput)

	tests := []struct {
		name                string
		inputBody           string
		input               signInInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Can sign in with correct input",
			inputBody: `{"username": "ivan", "password": "qwerty"}`,
			input: signInInput{
				Username: "ivan",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, input signInInput) {
				s.EXPECT().GenerateToken(input.Username, input.Password).Return("TOKEN", nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"token":"TOKEN"}`,
		},
		{
			name:                "Can't sign in with input without username",
			inputBody:           `{"password": "qwerty"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, input signInInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:                "Can't sign in with input without password",
			inputBody:           `{"username": "ivan"}`,
			mockBehavior:        func(s *mock_service.MockAuthorization, input signInInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can return error response if service failure",
			inputBody: `{"username": "ivan", "password": "qwerty"}`,
			input: signInInput{
				Username: "ivan",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, input signInInput) {
				s.EXPECT().GenerateToken(input.Username, input.Password).Return("", errors.New("service failure"))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			auth := mock_service.NewMockAuthorization(ctrl)
			tt.mockBehavior(auth, tt.input)

			services := &service.Service{Authorization: auth}
			var grpcLogger *grpcClient.Client
			var queueLogger *queueClient.QueueClient

			handler := NewHandler(services, grpcLogger, queueLogger)
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/auth/sign-in", handler.signIn)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-in",
				bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}
}
