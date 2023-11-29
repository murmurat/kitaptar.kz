package handler

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
	mock_service "github.com/murat96k/kitaptar.kz/internal/kitaptar/service/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_createAuthor(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, author api.AuthorRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"
	authorID := "5e442e63-be98-493e-9943-58f708e2f1df"
	testTable := []struct {
		name                 string
		inputBody            string
		inputAuthor          api.AuthorRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstname":"Test_author", "lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputAuthor: api.AuthorRequest{
				Firstname:   "Test_author",
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, author api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateAuthor(gomock.Any(), &author).Return(authorID, nil)

			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"5e442e63-be98-493e-9943-58f708e2f1df"}`,
		},
		{
			name:      "Wrong input (Missing firstname)",
			inputBody: `{"lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputAuthor: api.AuthorRequest{
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, author api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateAuthor(gomock.Any(), &author).Return(authorID, nil)

			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"5e442e63-be98-493e-9943-58f708e2f1df"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"firstname":"Test_author", "lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputAuthor: api.AuthorRequest{
				Firstname:   "Test_author",
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, author api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateAuthor(gomock.Any(), &author).Return("", errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputAuthor)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/authors", bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_updateAuthor(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, authorID string, author *api.AuthorRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputID              string
		inputAuthor          *api.AuthorRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK (Update all data)",
			inputBody: `{"firstname":"Test_author", "lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputID:   "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputAuthor: &api.AuthorRequest{
				Firstname:   "Test_author",
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, authorID string, author *api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateAuthor(gomock.Any(), authorID, author).Return(nil)

			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"firstname":"Test_author","lastname":"Test_author","image_path":"test_url","about_author":"test_description"}`,
		},
		{
			name:      "Input with missing firstname",
			inputBody: `{"lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputID:   "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputAuthor: &api.AuthorRequest{
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, authorID string, author *api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateAuthor(gomock.Any(), authorID, author).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"firstname":"","lastname":"Test_author","image_path":"test_url","about_author":"test_description"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"firstname":"Test_author", "lastname":"Test_author", "image_path":"test_url", "about_author":"test_description"}`,
			inputID:   "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputAuthor: &api.AuthorRequest{
				Firstname:   "Test_author",
				Lastname:    "Test_author",
				ImagePath:   "test_url",
				AboutAuthor: "test_description",
			},
			mockBehavior: func(s *mock_service.MockService, authorID string, author *api.AuthorRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateAuthor(gomock.Any(), authorID, author).Return(errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputID, testCase.inputAuthor)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", strings.TrimSpace(fmt.Sprintf("/authors/%s", testCase.inputID)), bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getAuthorById(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, authorID string)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		authorID             string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "OK",
			authorID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, authorID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAuthorById(gomock.Any(), authorID).Return(&entity.Author{
					Id:          uuid.MustParse("9634d0d0-ba9f-4516-9459-83eb58ebdb86"),
					Firstname:   "Meiirzhan",
					Lastname:    "Uristemov",
					ImagePath:   "meiir_path",
					AboutAuthor: "Meiirzhan is very good programmer",
					CreatedAt:   time.Time{},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"author_id":"9634d0d0-ba9f-4516-9459-83eb58ebdb86","firstname":"Meiirzhan","lastname":"Uristemov","image_path":"meiir_path","about_author":"Meiirzhan is very good programmer","CreatedAt":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:     "Service error",
			authorID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, authorID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAuthorById(gomock.Any(), authorID).Return(nil, errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.authorID)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()

			request := httptest.NewRequest("GET", strings.TrimSpace(fmt.Sprintf("/authors/%s", testCase.authorID)), bytes.NewBufferString(testCase.authorID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())
		})
	}
}

func TestHandler_deleteAuthor(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, authorID string)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Delete author by id",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, authorID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteAuthor(gomock.Any(), authorID).Return(nil)

			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"Author deleted"}`,
		},
		{
			name:    "Empty author id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, authorID string) {

			},
			expectedStatusCode:   404,
			expectedResponseBody: `404 page not found`,
		},
		{
			name:    "Service error",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, authorID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteAuthor(gomock.Any(), authorID).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:    "Request by not authorizing user",
			inputID: "lalksdmvklasndvklsaklv",
			mockBehavior: func(s *mock_service.MockService, authorID string) {
				s.EXPECT().VerifyToken("token").Return("", errors.New("error"))
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
			testCase.mockBehavior(mockService, testCase.inputID)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", fmt.Sprintf("/authors/%s", testCase.inputID), bytes.NewBufferString(testCase.inputID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getAllAuthors(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAllAuthors(gomock.Any()).Return([]entity.Author{{
					Id:          uuid.MustParse("9634d0d0-ba9f-4516-9459-83eb58ebdb86"),
					Firstname:   "Meiirzhan",
					Lastname:    "Uristemov",
					ImagePath:   "meiir_path",
					AboutAuthor: "Meiirzhan is very good programmer",
					CreatedAt:   time.Time{},
				}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"author_id":"9634d0d0-ba9f-4516-9459-83eb58ebdb86","firstname":"Meiirzhan","lastname":"Uristemov","image_path":"meiir_path","about_author":"Meiirzhan is very good programmer","CreatedAt":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name: "Service error",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAllAuthors(gomock.Any()).Return([]entity.Author{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name: "Request by not authorizing user",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return("", errors.New("error"))
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
			testCase.mockBehavior(mockService)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/authors", nil)

			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
