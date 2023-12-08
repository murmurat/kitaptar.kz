package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/user/entity"
	mock_service "github.com/murat96k/kitaptar.kz/internal/user/service/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_updateUser(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, req api.UpdateUserRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            api.UpdateUserRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Update all data",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(nil)
			},

			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User data updated!"}`,
		},
		{
			name:      "Missing firstname input",
			inputBody: `{"lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				LastName: "Test_user",
				Password: "password",
				Email:    "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User data updated!"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:      "Empty response field",
			inputBody: `{}`,
			inputUser: api.UpdateUserRequest{},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Update user data not provided"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputUser)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/users/", bytes.NewBufferString(testCase.inputBody))

			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getUser(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, userID string)

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Get user by id",
			inputID: "6f50ba79-1820-40c0-9c23-800400575c65",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetUserById(gomock.Any(), userID).Return(&entity.User{
					Id:        uuid.MustParse("6f50ba79-1820-40c0-9c23-800400575c65"),
					FirstName: "test_user",
					LastName:  "test_user",
					Password:  "password",
					Email:     "test_user@gmail.com",
				}, nil)
			},
			expectedStatusCode:   302,
			expectedResponseBody: `{"id":"6f50ba79-1820-40c0-9c23-800400575c65","firstname":"test_user","lastname":"test_user","password":"password","email":"test_user@gmail.com","is_verified":false,"role":"","CreatedAt":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:    "Empty user id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, errors.New("error"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
		{
			name:    "Service error",
			inputID: "6f50ba79-1820-40c0-9c23-800400575c65",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetUserById(gomock.Any(), userID).Return(nil, errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:    "Invalid user id",
			inputID: "lalksdmvklasndvklsaklv",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, errors.New("error"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, "")

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/users/", bytes.NewBufferString(testCase.inputID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_deleteUser(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, userID string)

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Delete user by id",
			inputID: "6f50ba79-1820-40c0-9c23-800400575c65",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteUser(gomock.Any(), userID).Return(nil)

			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User deleted"}`,
		},
		{
			name:    "Empty user id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, errors.New("error"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
		{
			name:    "Service error",
			inputID: "6f50ba79-1820-40c0-9c23-800400575c65",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteUser(gomock.Any(), userID).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:    "Invalid user id",
			inputID: "lalksdmvklasndvklsaklv",
			mockBehavior: func(s *mock_service.MockService, userID string) {
				s.EXPECT().VerifyToken("token").Return(userID, errors.New("error"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, "")

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", "/users/", bytes.NewBufferString(testCase.inputID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
